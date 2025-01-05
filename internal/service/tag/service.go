package tag

import (
	"blog-server/internal/common/constants"
	"blog-server/internal/common/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Service interface {
	// 获取tag
	GetTagList(ctx *gin.Context, data *ListData) (*response.PageListResponse, error)
	AddTag(ctx *gin.Context, tagName string) (int, error)
	GetTagDict(ctx *gin.Context) ([]*SimpleTagResp, error)
}

type service struct {
}

func NewService() Service {
	return &service{}
}
func GetDB(ctx *gin.Context) *gorm.DB {
	return ctx.MustGet(constants.DB).(*gorm.DB)
}
