package db

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

func Transaction(ctx context.Context, conn sqlx.SqlConn, fn func(ctx context.Context, session sqlx.Session) error) error {
	return conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}
