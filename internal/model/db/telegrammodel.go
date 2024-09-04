package db

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TelegramModel = (*customTelegramModel)(nil)

type (
	// TelegramModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTelegramModel.
	TelegramModel interface {
		telegramModel
		withSession(session sqlx.Session) TelegramModel
		WithSession(session sqlx.Session) TelegramModel
	}

	customTelegramModel struct {
		*defaultTelegramModel
	}
)

// NewTelegramModel returns a model for the database table.
func NewTelegramModel(conn sqlx.SqlConn) TelegramModel {
	return &customTelegramModel{
		defaultTelegramModel: newTelegramModel(conn),
	}
}

func (m *customTelegramModel) withSession(session sqlx.Session) TelegramModel {
	return NewTelegramModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customTelegramModel) WithSession(session sqlx.Session) TelegramModel {
	return NewTelegramModel(sqlx.NewSqlConnFromSession(session))
}
