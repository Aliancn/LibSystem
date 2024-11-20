// Code generated by goctl. DO NOT EDIT.
// versions:
//  goctl version: 1.7.3

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	paperFieldNames          = builder.RawFieldNames(&Paper{}, true)
	paperRows                = strings.Join(paperFieldNames, ",")
	paperRowsExpectAutoSet   = strings.Join(stringx.Remove(paperFieldNames, "id", "create_at", "create_time", "created_at", "update_at", "update_time", "updated_at"), ",")
	paperRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(paperFieldNames, "id", "create_at", "create_time", "created_at", "update_at", "update_time", "updated_at"))

	cachePublicPaperIdPrefix = "cache:public:paper:id:"
)

type (
	paperModel interface {
		Insert(ctx context.Context, data *Paper) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Paper, error)
		Update(ctx context.Context, data *Paper) error
		Delete(ctx context.Context, id int64) error
	}

	defaultPaperModel struct {
		sqlc.CachedConn
		table string
	}

	Paper struct {
		Id            int64          `db:"id"`
		Title         string         `db:"title"`
		Author        string         `db:"author"`
		Department    sql.NullString `db:"department"`
		Year          int64          `db:"year"`
		Status        string         `db:"status"`
		DownloadTimes int64          `db:"download_times"`
		FilePath      string         `db:"file_path"`
		CreatedAt     time.Time      `db:"created_at"`
		UpdatedAt     time.Time      `db:"updated_at"`
	}
)

func newPaperModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultPaperModel {
	return &defaultPaperModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      `"public"."paper"`,
	}
}

func (m *defaultPaperModel) Delete(ctx context.Context, id int64) error {
	publicPaperIdKey := fmt.Sprintf("%s%v", cachePublicPaperIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where id = $1", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, publicPaperIdKey)
	return err
}

func (m *defaultPaperModel) FindOne(ctx context.Context, id int64) (*Paper, error) {
	publicPaperIdKey := fmt.Sprintf("%s%v", cachePublicPaperIdPrefix, id)
	var resp Paper
	err := m.QueryRowCtx(ctx, &resp, publicPaperIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where id = $1 limit 1", paperRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultPaperModel) Insert(ctx context.Context, data *Paper) (sql.Result, error) {
	publicPaperIdKey := fmt.Sprintf("%s%v", cachePublicPaperIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3, $4, $5, $6, $7)", m.table, paperRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Title, data.Author, data.Department, data.Year, data.Status, data.DownloadTimes, data.FilePath)
	}, publicPaperIdKey)
	return ret, err
}

func (m *defaultPaperModel) Update(ctx context.Context, data *Paper) error {
	publicPaperIdKey := fmt.Sprintf("%s%v", cachePublicPaperIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = $1", m.table, paperRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Id, data.Title, data.Author, data.Department, data.Year, data.Status, data.DownloadTimes, data.FilePath)
	}, publicPaperIdKey)
	return err
}

func (m *defaultPaperModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cachePublicPaperIdPrefix, primary)
}

func (m *defaultPaperModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", paperRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultPaperModel) tableName() string {
	return m.table
}
