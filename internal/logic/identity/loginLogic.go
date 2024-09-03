package identitylogic

import (
	"context"
	"errors"
	"oracleX-backend/pkg/idx"
	"oracleX-backend/pkg/zero-contrib/appctx"
	"oracleX-backend/pkg/zero-contrib/errx"
	"oracleX-backend/pkg/zero-contrib/jwtx"
	"time"

	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"oracleX-backend/api/pb"
	"oracleX-backend/internal/model/db"
	"oracleX-backend/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	if err := in.ValidateAll(); err != nil {
		return nil, errx.Error(errx.InvalidArgument, err, errx.MsgInvalidArgument)
	}

	telegramUid, err := appctx.GetTelegramUid(l.ctx)
	if err != nil {
		return nil, errx.Error(errx.InvalidArgument, err, errx.InvalidArgument)
	}

	var uid int64
	telegramUser, err := l.svcCtx.TelegramModel.FindOne(l.ctx, telegramUid)
	if err != nil {
		// register a new user if not found
		if errors.Is(err, sqlx.ErrNotFound) {
			now := time.Now().UTC()
			telegramUsername, err := appctx.GetTelegramUsername(l.ctx)
			if err != nil {
				return nil, errx.Error(errx.CodeInternalServerErr, err, errx.MsgInternalServerErr)
			}
			res, err := l.svcCtx.UserModel.Insert(l.ctx, &db.User{
				Id:        idx.ID().Int64(),
				Username:  telegramUsername,
				Email:     "",
				Avatar:    "",
				IsDeleted: 0,
				CreatedAt: now,
				UpdatedAt: now,
			})
			if err != nil {
				return nil, errx.Error(errx.CodeInternalServerErr, err, errx.MsgInternalServerErr)
			}
			_uid, err := res.LastInsertId()
			if err != nil {
				return nil, errx.Error(errx.CodeInternalServerErr, err, errx.MsgInternalServerErr)
			}
			uid = _uid
		}
		return nil, errx.Error(errx.Internal, err, errx.MsgInternalServerErr)
	} else {
		uid = telegramUser.Uid
	}

	accessToken, err := l.svcCtx.JwtManager.Gen(jwtx.User{
		Uid: uid,
	})
	if err != nil {
		return nil, errx.Error(errx.Internal, err, errx.MsgInternalServerErr)
	}

	return &pb.LoginResp{AccessToken: accessToken}, nil
}
