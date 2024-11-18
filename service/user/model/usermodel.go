package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		FindAll(ctx context.Context) ([]*User, error)
	}

	customUserModel struct {
		*defaultUserModel
	}
)


// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn, c, opts...),
	}
}

func (m *customUserModel) FindAll(ctx context.Context) ([]*User, error) {
    query := fmt.Sprintf("SELECT %s FROM %s", userRows, m.table)
    var users []*User
    err := m.QueryRowsNoCacheCtx(ctx, &users, query)
    switch err {
    case nil:
        return users, nil
    case sql.ErrNoRows:
        return nil, ErrNotFound
    default:
        return nil, err
    }
}