package identitylogic

import (
	"context"

	"oracleX-backend/api/pb"
	"oracleX-backend/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginByWalletLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginByWalletLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginByWalletLogic {
	return &LoginByWalletLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginByWalletLogic) LoginByWallet(in *pb.LoginByWalletReq) (*pb.LoginByWalletResp, error) {
	// todo: add your logic here and delete this line

	return &pb.LoginByWalletResp{}, nil
}
