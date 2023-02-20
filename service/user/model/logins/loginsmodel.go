package logins

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ LoginsModel = (*customLoginsModel)(nil)

type (
	// LoginsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLoginsModel.
	LoginsModel interface {
		loginsModel
	}

	customLoginsModel struct {
		*defaultLoginsModel
	}
)

// NewLoginsModel returns a model for the database table.
func NewLoginsModel(conn sqlx.SqlConn, c cache.CacheConf) LoginsModel {
	return &customLoginsModel{
		defaultLoginsModel: newLoginsModel(conn, c),
	}
}
