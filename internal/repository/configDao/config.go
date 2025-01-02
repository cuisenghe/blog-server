package configDao

import (
	"errors"
	"gorm.io/gorm"
)

type Config struct {
	ID           string `json:"id"`
	AliPay       string `json:"ali_pay"`
	QQGroup      string `json:"qq_group"`
	WeChatPay    string `json:"we_chat_pay"`
	WeChatGroup  string `json:"we_chat_group"`
	UpdatedAt    string `json:"updated_at"`
	CreatedAt    string `json:"created_at"`
	ViewTime     string `json:"view_time"`
	BilibiliLink string `json:"bilibili_link"`
	GitEELink    string `json:"git_ee_link"`
	GithubLink   string `json:"github_link"`
	WeChatLink   string `json:"we_chat_link"`
	QQLink       string `json:"qq_link"`
	BlogNotice   string `json:"blog_notice"`
	PersonalSay  string `json:"personal_say"`
	AvatarBg     string `json:"avatar_bg"`
	BlogAvatar   string `json:"blog_avatar"`
	BlogName     string `json:"blog_name"`
}

func (Config) TableName() string {
	return "blog_config"
}
func GetConfig(db *gorm.DB) (*Config, error) {
	var config Config
	// 查找最新的配置
	tx := db.Order("createdAt desc").First(&config)
	if tx.Error != nil && !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, tx.Error
	}
	return &config, nil
}
