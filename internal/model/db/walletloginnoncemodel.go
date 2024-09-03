package db

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ WalletLoginNonceModel = (*customWalletLoginNonceModel)(nil)

type (
	// WalletLoginNonceModel is an interface to be customized, add more methods here,
	// and implement the added methods in customWalletLoginNonceModel.
	WalletLoginNonceModel interface {
		walletLoginNonceModel
		withSession(session sqlx.Session) WalletLoginNonceModel
	}

	customWalletLoginNonceModel struct {
		*defaultWalletLoginNonceModel
	}
)

// NewWalletLoginNonceModel returns a model for the database table.
func NewWalletLoginNonceModel(conn sqlx.SqlConn) WalletLoginNonceModel {
	return &customWalletLoginNonceModel{
		defaultWalletLoginNonceModel: newWalletLoginNonceModel(conn),
	}
}

func (m *customWalletLoginNonceModel) withSession(session sqlx.Session) WalletLoginNonceModel {
	return NewWalletLoginNonceModel(sqlx.NewSqlConnFromSession(session))
}
