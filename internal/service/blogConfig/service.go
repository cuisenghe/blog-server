package blogConfig

import (
	"blog-server/internal/repository/configDao"
	"github.com/gin-gonic/gin"
)

type Service interface {
	GetConfig(ctx *gin.Context) (*configDao.Config, error)
}
type service struct {
}

func NewService() Service {
	return &service{}
}
