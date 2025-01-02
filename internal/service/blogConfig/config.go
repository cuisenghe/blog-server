package blogConfig

import (
	"blog-server/internal/repository/configDao"
	"github.com/gin-gonic/gin"
)

func (s *service) GetConfig(ctx *gin.Context) (*configDao.Config, error) {
	config, err := configDao.GetConfig(GetDB(ctx))
	if err != nil {
		return nil, err
	}
	return config, nil
}
