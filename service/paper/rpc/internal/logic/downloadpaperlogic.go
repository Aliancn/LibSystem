package logic

import (
	"context"

	"LibSystem/service/paper/rpc/internal/svc"
	"LibSystem/service/paper/rpc/pb/paper_rpc_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadPaperLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDownloadPaperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadPaperLogic {
	return &DownloadPaperLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DownloadPaperLogic) DownloadPaper(in *paper_rpc_pb.DownloadPaperRequest) (*paper_rpc_pb.DownloadPaperResponse, error) {
	// todo: add your logic here and delete this line

	return &paper_rpc_pb.DownloadPaperResponse{}, nil
}
