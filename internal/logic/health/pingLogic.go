package healthlogic

import (
	"context"
	"fmt"
	"time"

	"oracleX-backend/api/pb"
	"oracleX-backend/internal/svc"
	"oracleX-backend/pkg/zero-contrib/errx"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *pb.Request) (*pb.Response, error) {
	if err := in.ValidateAll(); err != nil {
		return nil, errx.Error(errx.InvalidArgument, err, errx.MsgInvalidArgument)
	}
	return &pb.Response{Pong: fmt.Sprintf("recv ping=%s, ts=%d", in.Ping, time.Now().UTC())}, nil
}
