package service

import (
	"LibSystem/internal/api/request"
	"LibSystem/internal/api/response"
	"LibSystem/internal/model"
	"LibSystem/internal/repository"
	"github.com/gin-gonic/gin"
)

type IPaperService interface {
	GetPaperList(ctx *gin.Context) (response.PaperList, error)
	GetPaperById(ctx *gin.Context, id int) (response.PaperVO, error)
	GetPaperByTitle(ctx *gin.Context, title string) (response.PaperList, error)
	AddPaper(ctx *gin.Context, add request.PaperDTO) error
	UpdatePaper(ctx *gin.Context, update request.PaperDTO) error
	DeletePaper(ctx *gin.Context, id int) error
	GetPaperFilePath(ctx *gin.Context, id int) (string, error)
}

type PaperService struct {
	repo repository.PaperRepo
}

func (p PaperService) GetPaperList(ctx *gin.Context) (response.PaperList, error) {
	list, err := p.repo.GetAll(ctx)
	if err != nil {
		return response.PaperList{}, err
	}
	if len(list) == 0 {
		return response.PaperList{}, nil
	}
	var paperList []response.PaperVO
	for _, paper := range list {
		paperList = append(paperList, response.PaperVO{
			PaperId:       int64(paper.ID),
			Title:         paper.Title,
			Author:        paper.Author,
			Department:    paper.Department,
			Year:          paper.Year,
			Status:        paper.Status,
			FilePath:      paper.FilePath,
			DownloadTimes: paper.DownloadTimes,
			UploadTime:    paper.UpdatedAt,
		})
	}
	return response.PaperList{PaperList: paperList}, nil
}

func (p PaperService) GetPaperById(ctx *gin.Context, id int) (response.PaperVO, error) {
	paper, err := p.repo.GetById(ctx, uint(id))
	if err != nil {
		return response.PaperVO{}, err
	}
	return response.PaperVO{
		PaperId:       int64(paper.ID),
		Title:         paper.Title,
		Author:        paper.Author,
		Department:    paper.Department,
		Year:          paper.Year,
		Status:        paper.Status,
		FilePath:      paper.FilePath,
		DownloadTimes: paper.DownloadTimes,
		UploadTime:    paper.UpdatedAt,
	}, nil
}

func (p PaperService) GetPaperByTitle(ctx *gin.Context, title string) (response.PaperList, error) {
	// 模糊查询
	list, err := p.repo.GetByPaperName(ctx, title)
	if err != nil {
		return response.PaperList{}, err
	}
	if len(list) == 0 {
		return response.PaperList{}, nil
	}
	var paperList []response.PaperVO
	for _, paper := range list {
		paperList = append(paperList, response.PaperVO{
			PaperId:       int64(paper.ID),
			Title:         paper.Title,
			Author:        paper.Author,
			Department:    paper.Department,
			Year:          paper.Year,
			Status:        paper.Status,
			FilePath:      paper.FilePath,
			DownloadTimes: paper.DownloadTimes,
			UploadTime:    paper.UpdatedAt,
		})
	}
	return response.PaperList{PaperList: paperList}, nil
}

func (p PaperService) AddPaper(ctx *gin.Context, add request.PaperDTO) error {
	var paper = model.Paper{
		Title:      add.Title,
		Author:     add.Author,
		Department: add.Department,
		Year:       add.Year,
		Status:     add.Status,
		FilePath:   add.FilePath,
	}
	err := p.repo.Insert(ctx, paper)
	if err != nil {
		return err
	}
	return nil
}

func (p PaperService) UpdatePaper(ctx *gin.Context, update request.PaperDTO) error {
	var paper = model.Paper{
		ID:         update.PaperID,
		Title:      update.Title,
		Author:     update.Author,
		Department: update.Department,
		Year:       update.Year,
		Status:     update.Status,
	}
	err := p.repo.Update(ctx, paper)
	if err != nil {
		return err
	}
	return nil
}

func (p PaperService) DeletePaper(ctx *gin.Context, id int) error {
	err := p.repo.Delete(ctx, uint(id))
	if err != nil {
		return err
	}
	return nil
}

func (p PaperService) GetPaperFilePath(ctx *gin.Context, id int) (string, error) {
	filePath, err := p.repo.GetFilePath(ctx, uint(id))
	if err != nil {
		return "", err
	}
	return filePath, nil
}

func NewPaperService(repo repository.PaperRepo) IPaperService {
	return &PaperService{repo: repo}
}
