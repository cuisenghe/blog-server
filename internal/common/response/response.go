package response

import (
	bizErr "blog-server/pkg/errors"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PageListResponse struct {
	Current int         `json:"current"`
	Size    int         `json:"size"`
	Total   int         `json:"total"`
	List    interface{} `json:"list"`
}

func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"result":  data,
		"message": "success",
	})
}
func Fail(ctx *gin.Context, code int, message string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    code,
		"result":  nil,
		"message": message,
	})
}
func FailWithBizError(ctx *gin.Context, err error) {
	var bizErr *bizErr.BizError
	if !errors.As(err, &bizErr) {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    1,
			"result":  nil,
			"message": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    bizErr.Code,
		"result":  nil,
		"message": bizErr.Message,
	})
}
func SuccessWithPage(ctx *gin.Context, response *PageListResponse) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"result":  response,
		"message": "success",
	})
}
func FailWithPage(ctx *gin.Context, code int, message string, response *PageListResponse) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    code,
		"result":  response,
		"message": message,
	})
}
func FailWithData(ctx *gin.Context, code int, message string, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    code,
		"result":  data,
		"message": message,
	})
}
