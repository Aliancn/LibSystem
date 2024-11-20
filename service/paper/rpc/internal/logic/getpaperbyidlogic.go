package logic

import (
	"context"
	"strconv"

	"LibSystem/service/paper/rpc/internal/svc"
	"LibSystem/service/paper/rpc/pb/paper_rpc_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPaperByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPaperByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPaperByIdLogic {
	return &GetPaperByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPaperByIdLogic) GetPaperById(in *paper_rpc_pb.GetPaperByIdRequest) (*paper_rpc_pb.GetPaperByIdResponse, error) {
	_user, err := l.svcCtx.PaperModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &paper_rpc_pb.GetPaperByIdResponse{
		Code: 200,
		Msg:  "success",
		Data: &paper_rpc_pb.PaperVO{
			PaperId:     _user.Id,
			Title:       _user.Title,
			Author:      _user.Author,
			Department:  _user.Department.String,
			Year:        strconv.FormatInt(_user.Year, 10),
			Status:      _user.Status,
			DownloadNum: _user.DownloadTimes,
			FilePath:    _user.FilePath,
			UploadTime:  _user.CreatedAt.String(),
		},
	}, nil
}
