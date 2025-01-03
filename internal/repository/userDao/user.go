package userDao

import (
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
func GetUserByUsername(db *gorm.DB, username string) (*BlogUser, error) {
	var user *BlogUser
	tx := db.Where("username = ?", username).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}
func CreateUser(db *gorm.DB, user *BlogUser) (*BlogUser, error) {
	tx := db.Create(user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}

func GetUserById(db *gorm.DB, id uint64) (*BlogUser, error) {
	var user *BlogUser
	tx := db.Where("id = ?", id).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}
func UpdateUser(db *gorm.DB, user *BlogUser) (bool, error) {
	tx := db.Model(user).Updates(user)
	if tx.Error != nil {
		return false, tx.Error
	}
	return true, nil
}
func UpdatePassword(db *gorm.DB, id uint64, password string) (bool, error) {
	tx := db.Model(&BlogUser{}).Where("id = ?", id).Update("password", password)
	if tx.Error != nil {
		return false, tx.Error
	}
	return true, nil
}

// getUserList
func GetUserList(db *gorm.DB, current int, size int, nickname string, role int) ([]*BlogUser, error) {
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
func GetUserCount(db *gorm.DB, current int, size int, nickname string, role int) (int, error) {
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
	return int(count), nil
}
