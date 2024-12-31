package user

import (
	"blog-server/constants"
	"blog-server/internal/repository/userDao"
	"blog-server/pkg/authToken"
	bizErr "blog-server/pkg/errors"
	"gorm.io/gorm"

	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(ctx *gin.Context, username string, password string) (string, error) {
	// login的service
	// 判断是否存在该用户
	user, err := userDao.GetUserByUsername(ctx, username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", err
	}
	if user == nil {
		return "", bizErr.NewBizError(constants.USER_NO_REGISTERED, "用户未注册")
	}
	// 如果存在用户，判断密码是否正确
	passwordErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if passwordErr != nil {
		return "", err
	}
	// 如果正确的话，创建token
	token, err := authToken.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}
