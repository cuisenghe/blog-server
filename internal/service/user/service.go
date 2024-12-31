package user

import (
	"blog-server/internal/repository/userDao"
	"github.com/gin-gonic/gin"
)

type Service interface {
	// login
	Login(ctx *gin.Context, username string, password string) (string, error)
	Register(ctx *gin.Context, request *RegData) (uint64, error)
	GetUserInfoById(ctx *gin.Context, id uint64) (*userDao.BlogUser, error)
	UpdateUserInfo(ctx *gin.Context, user *userDao.BlogUser) (bool, error)

	UpdatePassword(ctx *gin.Context, data *UserPasswordData) (bool, error)
	GetUserList(ctx *gin.Context, data *GetUserListData) (*GetUserListResp, error)

	AdminUpdateUserInfo(ctx *gin.Context, data *AdminUserInfoData) (bool, error)
}
type service struct {
}

func NewService() Service {
	return &service{}
}