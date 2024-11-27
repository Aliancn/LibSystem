package service

import (
	"LibSystem/internal/api/response"
	"LibSystem/internal/model"
	"LibSystem/internal/repository"
	"time"

	"github.com/gin-gonic/gin"
)

type IInfoService interface {
	GetInfo(ctx *gin.Context) (resp response.InfoResponse, err error)
	AddInfo(ctx *gin.Context, paper_id int) (err error)
}

type InfoService struct {
	InfoRepo   repository.InfoRepo
	BookRepo   repository.BookRepo
	UserRepo   repository.UserRepo
	PaperRepo  repository.PaperRepo
	BorrowRepo repository.BorrowRepo
}

func NewInfoService(infoRepo repository.InfoRepo, bookRepo repository.BookRepo, userRepo repository.UserRepo, paperRepo repository.PaperRepo, borrowRepo repository.BorrowRepo) IInfoService {
	return &InfoService{
		InfoRepo:   infoRepo,
		BookRepo:   bookRepo,
		UserRepo:   userRepo,
		PaperRepo:  paperRepo,
		BorrowRepo: borrowRepo,
	}
}

func (is *InfoService) GetInfo(ctx *gin.Context) (response.InfoResponse, error) {
	infoResp, err := is.InfoRepo.GetInfo(ctx, 7)
	if err != nil {
		return response.InfoResponse{}, nil
	}
	downloadeTimes := make(map[string]int)
	for _, v := range infoResp {
		// 以天为单位
		date_str := v.DownloadTime.Format("2006-01-02")
		downloadeTimes[date_str] += 1
	}

	bookNum, err := is.BookRepo.GetNum(ctx)
	if err != nil {
		return response.InfoResponse{}, nil
	}
	userNum, err := is.UserRepo.GetNum(ctx)
	if err != nil {
		return response.InfoResponse{}, nil
	}
	paperNum, err := is.PaperRepo.GetNum(ctx)
	if err != nil {
		return response.InfoResponse{}, nil
	}
	borrowNum := make(map[string]int)
	borrowResp, err := is.BorrowRepo.GetBorrowInfo(ctx, 7)
	if err != nil {
		return response.InfoResponse{}, nil
	}
	for _, v := range borrowResp {
		date_str := v.BorrowDate.Format("2006-01-02")
		borrowNum[date_str] += 1
	}
	return response.InfoResponse{
		DownloadNum: downloadeTimes,
		BookNum:     bookNum,
		UserNum:     userNum,
		PaperNum:    paperNum,
		BorrowNum:   borrowNum,
	}, nil
}

func (is *InfoService) AddInfo(ctx *gin.Context, paper_id int) (err error) {
	dwinfo := model.Info{
		PaperID:      paper_id,
		DownloadTime: time.Now(),
	}
	err = is.InfoRepo.AddInfo(ctx, dwinfo)
	return
}
