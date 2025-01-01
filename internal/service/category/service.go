package category

import (
	"github.com/gin-gonic/gin"
)

type Service interface {
	GetCategoryDict(ctx *gin.Context) ([]*SimpleCategory, error)
}
type service struct {
}

func NewService() Service {
	return &service{}
}
