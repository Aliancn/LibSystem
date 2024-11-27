package service

import (
	"LibSystem/common"
	"LibSystem/common/utils"
	"LibSystem/global"
	"LibSystem/internal/api/request"
	"LibSystem/internal/api/response"
	"LibSystem/internal/model"
	"LibSystem/internal/repository"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IUserService interface {
	Login(ctx *gin.Context, login request.UserLogin) (response.UserLoginInfo, error)
	Logout(ctx *gin.Context) error
	AddUser(ctx *gin.Context, add request.UserDTO) error
	EditPassword(ctx *gin.Context, reqs request.UserEditPassword) error
	DeleteUser(ctx *gin.Context, id request.UserID) error
	UpdateUser(ctx *gin.Context, update request.UserDTO) error
	GetById(ctx *gin.Context, id request.UserID) (response.UserVO, error)
	GetByUsername(ctx *gin.Context, username request.Username) (response.UserVO, error)
	RegisterRole(ctx *gin.Context, register request.UserRegister) error
	RegisterAdmin(ctx *gin.Context, register request.UserRegister) error
	GetList(ctx *gin.Context, pageID, pageSize int) (response.UserList, error)
}

type UserService struct {
	repo repository.UserRepo
}

func NewUserService(repo repository.UserRepo) IUserService {
	return &UserService{repo: repo}
}

func (u UserService) GetById(ctx *gin.Context, id request.UserID) (response.UserVO, error) {
	resp, err := u.repo.GetById(ctx, uint(id.UserId))
	if err != nil {
		return response.UserVO{}, err
	}
	return response.UserVO{
		UserId:   int64(resp.ID),
		Username: resp.Username,
		Password: resp.Password,
		Role:     resp.Role,
		Name:     resp.Name,
		Phone:    resp.Phone,
		Email:    resp.Email,
	}, nil
}

func (u UserService) GetByUsername(ctx *gin.Context, username request.Username) (response.UserVO, error) {
	resp, err := u.repo.GetByUserName(ctx, username.Username)
	if err != nil {
		return response.UserVO{}, err
	}
	return response.UserVO{
		UserId:   int64(resp.ID),
		Username: resp.Username,
		Password: resp.Password,
		Role:     resp.Role,
		Name:     resp.Name,
		Phone:    resp.Phone,
		Email:    resp.Email,
	}, nil
}

func (u UserService) GetList(ctx *gin.Context, pageID, pageSize int) (response.UserList, error) {
	list, err := u.repo.GetAll(ctx, pageID, pageSize)
	if err != nil {
		return response.UserList{}, err
	}
	var respList []response.UserVO
	for _, item := range list {
		respList = append(respList, response.UserVO{
			UserId:   int64(item.ID),
			Username: item.Username,
			Password: item.Password,
			Role:     item.Role,
			Name:     item.Name,
			Phone:    item.Phone,
			Email:    item.Email,
		})
	}
	return response.UserList{
		List: respList,
	}, nil
}

func (u UserService) Login(ctx *gin.Context, login request.UserLogin) (response.UserLoginInfo, error) {
	user, err := u.repo.GetByUserName(ctx, login.Username)
	if err != nil {
		return response.UserLoginInfo{}, err
	}
	if user == nil {
		return response.UserLoginInfo{}, common.Error_ACCOUNT_NOT_FOUND
	}
	hashPassword := utils.MD5V(login.Password, "alia", 1)
	if user.Password != hashPassword {
		return response.UserLoginInfo{}, common.Error_PASSWORD_ERROR
	}

	token, err := utils.GetJwtToken(global.Config.Jwt.Secret, time.Now().Unix(), global.Config.Jwt.Expire, int64(user.ID), user.Role)
	if err != nil {
		return response.UserLoginInfo{}, err
	}
	return response.UserLoginInfo{
		UserId:   int64(user.ID),
		Username: user.Username,
		Role:     user.Role,
		Token:    token,
	}, nil
}

func (u UserService) Logout(ctx *gin.Context) error {
	return nil
}

func (u UserService) RegisterRole(ctx *gin.Context, register request.UserRegister) error {
	_, err := u.repo.GetByUserName(ctx, register.Username)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return common.Error_ALREADY_EXISTS
	}
	hashPassword := utils.MD5V(register.Password, "alia", 1)
	entity := model.User{
		Username: register.Username,
		Password: hashPassword,
		Role:     common.RoleUser,
	}
	err = u.repo.Insert(ctx, entity)
	return err
}

func (u UserService) RegisterAdmin(ctx *gin.Context, register request.UserRegister) error {
	_, err := u.repo.GetByUserName(ctx, register.Username)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return common.Error_ALREADY_EXISTS
	}
	hashPassword := utils.MD5V(register.Password, "alia", 1)
	entity := model.User{
		Username: register.Username,
		Password: hashPassword,
		Role:     common.RoleAdmin,
	}
	err = u.repo.Insert(ctx, entity)
	return err
}

func (u UserService) AddUser(ctx *gin.Context, add request.UserDTO) error {
	hashPassword := utils.MD5V(add.Password, "alia", 1)
	global.Log.Debug("hashPassword:", hashPassword)
	entity := model.User{
		Username: add.Username,
		Password: hashPassword,
		Role:     add.Role,
		Name:     add.Name,
		Phone:    add.Phone,
		Email:    add.Email,
	}
	err := u.repo.Insert(ctx, entity)
	return err
}

func (u UserService) EditPassword(ctx *gin.Context, reqs request.UserEditPassword) error {
	// 1.获取用户信息
	user, err := u.repo.GetById(ctx, uint(reqs.UserId))
	if err != nil {
		return err
	}
	// 校验用户老密码
	if user == nil {
		return common.Error_ACCOUNT_NOT_FOUND
	}
	oldHashPassword := utils.MD5V(reqs.OldPassword, "alia", 1)
	if user.Password != oldHashPassword {
		return common.Error_PASSWORD_ERROR
	}
	// 修改用户密码
	newHashPassword := utils.MD5V(reqs.NewPassword, "alia", 1)
	err = u.repo.Update(ctx, model.User{
		ID:       uint(reqs.UserId),
		Password: newHashPassword,
	})
	return err
}

func (u UserService) DeleteUser(ctx *gin.Context, id request.UserID) error {
	err := u.repo.Delete(ctx, uint(id.UserId))
	return err
}

func (u UserService) UpdateUser(ctx *gin.Context, update request.UserDTO) error {
	if update.UserId == 0 {
		return common.Error_ACCOUNT_NOT_FOUND
	}
	entity := model.User{
		ID:    uint(update.UserId),
		Role:  update.Role,
		Name:  update.Name,
		Phone: update.Phone,
		Email: update.Email,
	}
	err := u.repo.Update(ctx, entity)
	return err
}
