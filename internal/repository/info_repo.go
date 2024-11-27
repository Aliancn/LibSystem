package repository

import (
	"LibSystem/internal/model"
	"context"
)

type InfoRepo interface {
	GetInfo(ctx context.Context, day int) ([]model.Info, error)
	AddInfo(ctx context.Context, info model.Info) error
}
