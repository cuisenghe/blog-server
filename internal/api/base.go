package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"net/http"
)

// Binding 进行数据绑定
func Binding(ctx *gin.Context, data interface{}) {
	if err := ctx.ShouldBindJSON(data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	return
}
func GetDB(ctx *gin.Context) *gorm.DB {
	return ctx.MustGet("db").(*gorm.DB)
}
