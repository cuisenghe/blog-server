package category

import (
	"blog-server/internal/common/constants"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Service interface {
	GetCategoryDict(ctx *gin.Context) ([]*SimpleCategory, error)
}
type service struct {
}

func NewService() Service {
	return &service{}
}

func GetDB(ctx *gin.Context) *gorm.DB {
	return ctx.MustGet(constants.DB).(*gorm.DB)
}
