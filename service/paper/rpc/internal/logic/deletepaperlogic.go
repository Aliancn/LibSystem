package logic

import (
	"context"

	"LibSystem/service/paper/rpc/internal/svc"
	"LibSystem/service/paper/rpc/pb/paper_rpc_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePaperLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletePaperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePaperLogic {
	return &DeletePaperLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeletePaperLogic) DeletePaper(in *paper_rpc_pb.DeletePaperRequest) (*paper_rpc_pb.DeletePaperResponse, error) {
	err := l.svcCtx.PaperModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &paper_rpc_pb.DeletePaperResponse{
		Code: 200,
		Msg:  "删除成功",
	}, nil
}
