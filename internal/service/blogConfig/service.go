package blogConfig

import (
	"blog-server/internal/common/constants"
	"blog-server/internal/repository/configDao"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Service interface {
	GetConfig(ctx *gin.Context) (*configDao.Config, error)
}
type service struct {
}

func NewService() Service {
	return &service{}
}
func GetDB(ctx *gin.Context) *gorm.DB {
	return ctx.MustGet(constants.DB).(*gorm.DB)
}
