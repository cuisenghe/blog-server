package userDao

import (
	"blog-server/constants"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

// 用户模块的dao
type BlogUser struct {
	ID        uint64
	Username  string
	Password  string
	Role      int
	NickName  string
	Avatar    string
	CreatedAt time.Time `gorm:"column:createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt"`
	QQ        string
	IP        string
}

func (user BlogUser) TableName() string {
	return "blog_user"
}
func GetUserByUsername(ctx *gin.Context, username string) (*BlogUser, error) {
	db := ctx.MustGet(constants.DB).(*gorm.DB)
	var user *BlogUser
	tx := db.Where("username = ?", username).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}
func CreateUser(ctx *gin.Context, user *BlogUser) (*BlogUser, error) {
	db := ctx.MustGet(constants.DB).(*gorm.DB)
	tx := db.Create(user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}

func GetUserById(ctx *gin.Context, id uint64) (*BlogUser, error) {
	db := ctx.MustGet(constants.DB).(*gorm.DB)
	var user *BlogUser
	tx := db.Where("id = ?", id).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}
func UpdateUser(ctx *gin.Context, user *BlogUser) (bool, error) {
	db := ctx.MustGet(constants.DB).(*gorm.DB)
	tx := db.Model(user).Updates(user)
	if tx.Error != nil {
		return false, tx.Error
	}
	return true, nil
}
func UpdatePassword(ctx *gin.Context, id uint64, password string) (bool, error) {
	db := ctx.MustGet(constants.DB).(*gorm.DB)
	tx := db.Model(&BlogUser{}).Where("id = ?", id).Update("password", password)
	if tx.Error != nil {
		return false, tx.Error
	}
	return true, nil
}

// getUserList
func GetUserList(ctx *gin.Context, current int, size int, nickname string, role int) ([]*BlogUser, error) {
	db := ctx.MustGet(constants.DB).(*gorm.DB)
	var users []*BlogUser
	offset := (current - 1) * size
	// 分页条件查询
	if len(nickname) != 0 {
		db.Where("nick_name LIKE ?", "%"+nickname+"%")
	}
	if role != 0 {
		db.Where("role = ?", role)
	}
	tx := db.Offset(offset).Limit(size).Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return users, nil
}
func GetUserCount(ctx *gin.Context, current int, size int, nickname string, role int) (int64, error) {
	db := ctx.MustGet(constants.DB).(*gorm.DB)
	var count int64
	// 分页条件查询
	if len(nickname) != 0 {
		db.Where("nick_name LIKE ?", "%"+nickname+"%")
	}
	if role != 0 {
		db.Where("role = ?", role)
	}
	// 求总数
	tx := db.Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return count, nil
}
