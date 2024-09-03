package interceptorx

import (
	"context"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type contextKey int

const (
	IP contextKey = iota
	UID
	TELEGRAM_UID
)

func MetadataInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		var err error

		var ip string
		var uid int64
		var telegramUid int64

		ipList := md.Get("gateway-ip")
		if len(ipList) >= 1 {
			ip = ipList[0]
		}

		uidList := md.Get("gateway-uid")
		if len(uidList) >= 1 {
			uid, err = strconv.ParseInt(uidList[0], 10, 64)
			if err != nil {
				logx.Error(err)
			}
		}

		telegramUidList := md.Get("gateway-tg-uid")
		if len(telegramUidList) >= 1 {
			telegramUid, err = strconv.ParseInt(telegramUidList[0], 10, 64)
			if err != nil {
				logx.Error(err)
			}
		}

		ctx = context.WithValue(ctx, IP, ip)
		ctx = context.WithValue(ctx, UID, uid)
		ctx = context.WithValue(ctx, TELEGRAM_UID, telegramUid)
	}

	return handler(ctx, req)
}
