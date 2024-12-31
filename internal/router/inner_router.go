package router

import "github.com/gin-gonic/gin"

// 内部接口
func InitInnerRouter(router *gin.Engine) {
	innerGroup := router.Group("/inner")
	{
		articleGroup := innerGroup.Group("/article")
		{
			articleGroup.GET("/")
		}
	}
}
