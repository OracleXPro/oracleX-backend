package identitylogic

import (
	"context"
	"oracleX-backend/pkg/idx"
	"oracleX-backend/pkg/zero-contrib/appctx"
	"oracleX-backend/pkg/zero-contrib/errx"
	"oracleX-backend/pkg/zero-contrib/jwtx"
	"time"

	"github.com/pkg/errors"

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
	DbTelegramUser, err := l.svcCtx.TelegramModel.FindOne(l.ctx, telegramUid)
	if err != nil {
		// register a new user if not found
		if errors.Is(err, sqlx.ErrNotFound) {
			now := time.Now().UTC()
			telegramUser := appctx.GetTelegramUserinfo(l.ctx)
			if err = db.Transaction(l.ctx, l.svcCtx.MySQLConn, func(ctx context.Context, session sqlx.Session) error {
				res, err := l.svcCtx.UserModel.WithSession(session).Insert(ctx, &db.User{
					Id:        idx.ID().Int64(),
					Username:  telegramUser.Username,
					Email:     "",
					Avatar:    "",
					IsDeleted: 0,
					CreatedAt: now,
					UpdatedAt: now,
				})
				if err != nil {
					return errors.Wrap(err, "fail to create user")
				}
				_uid, err := res.LastInsertId()
				if err != nil {
					return errors.Wrap(err, "fail to get last insert id")
				}

				_, err = l.svcCtx.TelegramModel.WithSession(session).Insert(ctx, &db.Telegram{
					Id:        telegramUser.Uid,
					Uid:       _uid,
					Username:  telegramUser.Username,
					FirstName: telegramUser.FirstName,
					LastName:  telegramUser.LastName,
					CreatdAt:  now,
					UpdatedAt: now,
				})
				if err != nil {
					return errors.Wrap(err, "fail to create telegram user")
				}
				uid = _uid
				return nil
			}); err != nil {
				return nil, errx.Error(errx.CodeInternalServerErr, err, errx.MsgInternalServerErr)
			}
		} else {
			return nil, errx.Error(errx.Internal, err, errx.MsgInternalServerErr)
		}
	} else {
		uid = DbTelegramUser.Uid
	}

	accessToken, err := l.svcCtx.JwtManager.Gen(jwtx.User{
		Uid: uid,
	})
	if err != nil {
		return nil, errx.Error(errx.Internal, err, errx.MsgInternalServerErr)
	}

	return &pb.LoginResp{AccessToken: accessToken}, nil
}
