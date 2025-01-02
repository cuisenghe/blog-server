package user

import (
	"blog-server/internal/repository/userDao"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// admin
type AdminUserInfoData struct {
	ID       uint64 `json:"id"`
	NickName string `json:"nick_name"`
	Avatar   string `json:"avatar"`
}

func (s *service) AdminUpdateUserInfo(ctx *gin.Context, data *AdminUserInfoData) (bool, error) {
	// 查询
	user, err := userDao.GetUserById(GetDB(ctx), data.ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return false, err
	}
	if user == nil {
		return false, nil
	}
	// 更新
	user.Avatar = data.Avatar
	user.NickName = data.NickName
	isSuccess, err := userDao.UpdateUser(GetDB(ctx), user)
	if err != nil {
		return false, err
	}
	return isSuccess, nil

}
