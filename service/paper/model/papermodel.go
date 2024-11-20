package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PaperModel = (*customPaperModel)(nil)

type (
	// PaperModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPaperModel.
	PaperModel interface {
		paperModel
		GetAllPaper(ctx context.Context) ([]*Paper, error)
		GetPaperByTitle(ctx context.Context, title string) ([]*Paper, error)
	}

	customPaperModel struct {
		*defaultPaperModel
	}
)

// NewPaperModel returns a model for the database table.
func NewPaperModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) PaperModel {
	return &customPaperModel{
		defaultPaperModel: newPaperModel(conn, c, opts...),
	}
}

func (m *customPaperModel) GetAllPaper(ctx context.Context) ([]*Paper, error) {
	query := fmt.Sprintf("SELECT %s FROM %s", paperRows, m.table)
	var resp []*Paper
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sql.ErrNoRows:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// GetPaperByTitle returns a Paper object by title.
func (m *customPaperModel) GetPaperByTitle(ctx context.Context, title string) ([]*Paper, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE title like ?", paperRows, m.table)
	var resp []*Paper
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, title)
	switch err {
	case nil:
		return resp, nil
	case sql.ErrNoRows:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
