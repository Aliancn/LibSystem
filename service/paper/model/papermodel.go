package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ PaperModel = (*customPaperModel)(nil)

type (
	// PaperModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPaperModel.
	PaperModel interface {
		paperModel
		withSession(session sqlx.Session) PaperModel
	}

	customPaperModel struct {
		*defaultPaperModel
	}
)

// NewPaperModel returns a model for the database table.
func NewPaperModel(conn sqlx.SqlConn) PaperModel {
	return &customPaperModel{
		defaultPaperModel: newPaperModel(conn),
	}
}

func (m *customPaperModel) withSession(session sqlx.Session) PaperModel {
	return NewPaperModel(sqlx.NewSqlConnFromSession(session))
}
