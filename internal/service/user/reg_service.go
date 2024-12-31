package user

import (
	"blog-server/constants"
	"blog-server/internal/repository/userDao"
	bizErr "blog-server/pkg/errors"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type RegData struct {
	Username string `json:"username"`
	Password string `json:"password"`
	NickName string `json:"nick_name"`
	QQ       string `json:"qq"`
	Avatar   string `json:"avatar"`
	Role     int    `json:"role"`
}

func (s *service) Register(ctx *gin.Context, req *RegData) (uint64, error) {
	// 先判断是否存在该用户
	userInfo, err := userDao.GetUserByUsername(ctx, req.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, err
	}

	if userInfo != nil {
		// 存在该用户，返回用户已经注册
		return 0, bizErr.NewBizError(constants.USER_REGISTERED, "用户已经注册")
	}
	// 如果不存在该用户，则进行注册
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	userInfo, err = userDao.CreateUser(ctx, &userDao.BlogUser{
		Username: req.Username,
		Password: string(password),
		Role:     req.Role,
		NickName: req.NickName,
		Avatar:   req.Avatar,
		QQ:       req.QQ,
	})
	if err != nil {
		return 0, err
	}
	// 创建成功
	return userInfo.ID, nil
}
