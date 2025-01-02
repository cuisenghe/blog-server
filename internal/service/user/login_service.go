package user

import (
	"blog-server/constants"
	"blog-server/internal/repository/userDao"
	"blog-server/pkg/authToken"
	bizErr "blog-server/pkg/errors"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type Response struct {
	Token    string `json:"token"`
	Username string `json:"username"`
	Role     int    `json:"role"`
	Id       int    `json:"id"`
}

func (s *service) Login(ctx *gin.Context, username string, password string) (*Response, error) {
	// login的service
	// 判断是否存在该用户
	user, err := userDao.GetUserByUsername(GetDB(ctx), username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return &Response{
			Token:    "",
			Username: "",
			Role:     0,
			Id:       0,
		}, err
	}
	if user == nil {
		return &Response{
			Token:    "",
			Username: "",
			Role:     0,
			Id:       0,
		}, bizErr.NewBizError(constants.USER_NO_REGISTERED, "用户未注册")
	}
	// 如果存在用户，判断密码是否正确
	passwordErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if passwordErr != nil {
		return &Response{
			Token:    "",
			Username: user.Username,
			Role:     user.Role,
			Id:       int(user.ID),
		}, bizErr.NewBizError(constants.USER_PASS_FAILED, "用户名密码错误")
	}
	// 如果正确的话，创建token
	token, err := authToken.GenerateToken(user.Username)
	if err != nil {
		log.Println(err)
		return &Response{
			Token:    "",
			Username: user.Username,
			Role:     user.Role,
			Id:       int(user.ID),
		}, err
	}
	return &Response{
		Token:    token,
		Username: user.Username,
		Role:     user.Role,
		Id:       int(user.ID),
	}, nil
}
