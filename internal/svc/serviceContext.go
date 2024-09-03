package svc

import (
	"oracleX-backend/internal/config"
	"oracleX-backend/internal/model/db"
	"oracleX-backend/pkg/zero-contrib/jwtx"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	JwtManager *jwtx.JWTManager

	BizCache *redis.Redis

	MySQLConn sqlx.SqlConn
	CacheDB   sqlc.CachedConn

	UserModel     db.UserModel
	TelegramModel db.TelegramModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.MysqlDB.DataSource)
	return &ServiceContext{
		Config: c,

		JwtManager: jwtx.NewJwtManager(*c.JWT[0]),
		MySQLConn:  mysqlConn,

		// model
		UserModel:     db.NewUserModel(mysqlConn),
		TelegramModel: db.NewTelegramModel(mysqlConn),
	}
}
