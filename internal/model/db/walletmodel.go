package db

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ WalletModel = (*customWalletModel)(nil)

type (
	// WalletModel is an interface to be customized, add more methods here,
	// and implement the added methods in customWalletModel.
	WalletModel interface {
		walletModel
		withSession(session sqlx.Session) WalletModel
	}

	customWalletModel struct {
		*defaultWalletModel
	}
)

// NewWalletModel returns a model for the database table.
func NewWalletModel(conn sqlx.SqlConn) WalletModel {
	return &customWalletModel{
		defaultWalletModel: newWalletModel(conn),
	}
}

func (m *customWalletModel) withSession(session sqlx.Session) WalletModel {
	return NewWalletModel(sqlx.NewSqlConnFromSession(session))
}
