package repository

import (
	"LibSystem/internal/model"
	"context"
)

type PaperRepo interface {
	// GetByPaperName 根据papername获取试卷信息
	GetByPaperName(ctx context.Context, paperName string) ([]model.Paper, error)
	// GetById 根据id获取试卷信息
	GetById(ctx context.Context, id uint) (*model.Paper, error)
	// Update  修改
	Update(ctx context.Context, paper model.Paper) error
	// Insert 插入数据
	Insert(ctx context.Context, entity model.Paper) error
	// Delete 删除数据
	Delete(ctx context.Context, id uint) error
	// GetAll 获取所有数据
	GetAll(ctx context.Context) ([]model.Paper, error)
	//GetFilePath 获取文件路径
	GetFilePath(ctx context.Context, id uint) (string, error)
}
