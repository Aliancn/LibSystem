package logic

import (
	"context"
	"strconv"

	"LibSystem/service/paper/rpc/internal/svc"
	"LibSystem/service/paper/rpc/pb/paper_rpc_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPaperByTitleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPaperByTitleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPaperByTitleLogic {
	return &GetPaperByTitleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPaperByTitleLogic) GetPaperByTitle(in *paper_rpc_pb.GetPaperByTitleRequest) (*paper_rpc_pb.GetPaperByTitleResponse, error) {
	_papers, err := l.svcCtx.PaperModel.GetPaperByTitle(l.ctx, in.Title)
	if err != nil {
		return nil, err
	}

	var paperlist paper_rpc_pb.PaperList
	for _, v := range _papers {
		paperlist.PaperList = append(paperlist.PaperList, &paper_rpc_pb.PaperVO{
			PaperId:     v.Id,
			Title:       v.Title,
			Author:      v.Author,
			Department:  v.Department.String,
			Year:        strconv.FormatInt(v.Year, 10),
			Status:      v.Status,
			DownloadNum: v.DownloadTimes,
			FilePath:    v.FilePath,
			UploadTime:  v.CreatedAt.String(),
		})
	}

	return &paper_rpc_pb.GetPaperByTitleResponse{
		Code: 200,
		Msg:  "success",
		Data: &paperlist,
	}, nil
}
