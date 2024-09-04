// Code generated by goctl. DO NOT EDIT!
// Source: oracleX.proto

package server

import (
	"context"

	"oracleX-backend/api/pb"
	"oracleX-backend/internal/logic/health"
	"oracleX-backend/internal/svc"
)

type HealthServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedHealthServer
}

func NewHealthServer(svcCtx *svc.ServiceContext) *HealthServer {
	return &HealthServer{
		svcCtx: svcCtx,
	}
}

func (s *HealthServer) Ping(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	l := healthlogic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}
