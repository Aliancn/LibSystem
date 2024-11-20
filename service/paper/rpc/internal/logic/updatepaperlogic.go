package logic

import (
	"context"
	"database/sql"
	"strconv"

	"LibSystem/service/paper/model"
	"LibSystem/service/paper/rpc/internal/svc"
	"LibSystem/service/paper/rpc/pb/paper_rpc_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePaperLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePaperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePaperLogic {
	return &UpdatePaperLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdatePaperLogic) UpdatePaper(in *paper_rpc_pb.UpdatePaperRequest) (*paper_rpc_pb.UpdatePaperResponse, error) {
	newpaper, err := l.svcCtx.PaperModel.FindOne(l.ctx, in.PaperId)
	if err != nil {
		return nil, err
	}
	if newpaper == nil {
		return nil, model.ErrNotFound
	}
	newpaper.Title = in.Title
	newpaper.Author = in.Author
	newpaper.Department = sql.NullString{String: in.Department, Valid: in.Department != ""}
	year, err := strconv.ParseInt(in.Year, 10, 64)
	if err != nil {
		return nil, err
	}
	newpaper.Year = year
	newpaper.Status = in.Status

	err = l.svcCtx.PaperModel.Update(l.ctx, newpaper)
	if err != nil {
		return nil, err
	}
	return &paper_rpc_pb.UpdatePaperResponse{
		Code: 200,
		Msg:  "success",
	}, nil
}
