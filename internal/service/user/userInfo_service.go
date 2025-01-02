package user

import (
	"blog-server/constants"
	"blog-server/internal/repository/userDao"
	bizErr "blog-server/pkg/errors"

	"errors"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserPasswordData struct {
	ID        uint64 `json:"id"`
	Password  string `json:"password"`
	Password1 string `json:"password1"`
	Password2 string `json:"password2"`
}

// 分页获取用户数据

type GetUserListData struct {
	Current  int    `json:"current"`
	Size     int    `json:"size"`
	NickName string `json:"nick_name"`
	Role     int    `json:"role"`
}

// 分页获取返回值

type GetUserListResp struct {
	List    []*userDao.BlogUser `json:"list"`
	Current int                 `json:"current"`
	Size    int                 `json:"size"`
	Total   int64               `json:"total"`
}

func (s *service) GetUserInfoById(ctx *gin.Context, id uint64) (*userDao.BlogUser, error) {
	// 获取db
	user, err := userDao.GetUserById(GetDB(ctx), id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (s *service) UpdateUserInfo(ctx *gin.Context, user *userDao.BlogUser) (bool, error) {
	_, err := userDao.UpdateUser(GetDB(ctx), user)
	if err != nil {
		return false, err
	}
	return true, nil
}

// 更改用户密码
func (s *service) UpdatePassword(ctx *gin.Context, data *UserPasswordData) (bool, error) {
	// 查询用户信息
	user, err := userDao.GetUserById(GetDB(ctx), data.ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return false, err
	}
	// 判断用户是否存在
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 用户不存在
		return false, bizErr.NewBizError(constants.USER_NO_REGISTERED, "用户不存在")
	}
	// 用户存在,判断旧密码是否相同
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)) != nil {
		return false, bizErr.NewBizError(constants.USER_PASS_FAILED, "用户密码错误")
	}
	// 判断两个新密码是否一致
	if data.Password1 != data.Password2 {
		return false, nil
	}
	password, err := bcrypt.GenerateFromPassword([]byte(data.Password1), bcrypt.DefaultCost)
	if err != nil {
		return false, err
	}
	_, err = userDao.UpdatePassword(GetDB(ctx), data.ID, string(password))
	if err != nil {
		return false, err
	}
	return true, nil
}

// 分页获取数据
func (s *service) GetUserList(ctx *gin.Context, data *GetUserListData) (*GetUserListResp, error) {
	// 获取数据
	list, err := userDao.GetUserList(GetDB(ctx), data.Current, data.Size, data.NickName, data.Role)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &GetUserListResp{}, nil
	}
	// 获取数量
	count, err := userDao.GetUserCount(GetDB(ctx), data.Current, data.Size, data.NickName, data.Role)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if count == 0 {
		return &GetUserListResp{
			List:    nil,
			Current: data.Current,
			Size:    data.Size,
			Total:   0,
		}, nil
	}
	return &GetUserListResp{
		List:    list,
		Current: data.Current,
		Size:    data.Size,
		Total:   count,
	}, nil
}
