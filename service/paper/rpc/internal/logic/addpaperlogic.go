package logic

import (
	"context"

	"LibSystem/service/paper/rpc/internal/svc"
	"LibSystem/service/paper/rpc/pb/paper_rpc_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPaperLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPaperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPaperLogic {
	return &AddPaperLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddPaperLogic) AddPaper(in *paper_rpc_pb.AddPaperRequest) (*paper_rpc_pb.AddPaperResponse, error) {
	// todo: add your logic here and delete this line
	
	return &paper_rpc_pb.AddPaperResponse{}, nil
}
