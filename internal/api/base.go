package api

import (
	"blog-server/constants"
	BizErr "blog-server/pkg/errors"
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"net/http"
)

// base 工具类，用于结果的反回

// ReturnSuccess 返回成功
func ReturnSuccess(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    constants.SUCCESS,
		"result":  data,
		"message": constants.SUCCESS_MSG,
	})
}

// ReturnError 返回失败
func ReturnError(ctx *gin.Context, code int, msg string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": msg,
		"result":  nil,
	})
}
func ReturnBizError(ctx *gin.Context, err error) {
	var bizError *BizErr.BizError
	ok := errors.As(err, &bizError)
	if !ok {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    bizError.Code,
		"message": bizError.Message,
		"result":  nil,
	})
}
func ReturnBizErrorWithData(ctx *gin.Context, err error, data interface{}) {
	var bizError *BizErr.BizError
	ok := errors.As(err, &bizError)
	if !ok {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    bizError.Code,
		"message": bizError.Message,
		"result":  data,
	})
}

// ReturnErrWithData 返回失败伴随数据
func ReturnErrWithData(ctx *gin.Context, code int, msg string, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": msg,
		"result":  data,
	})
}

// 进行数据绑定
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

// 分页返回
type PageInfo struct {
	Current int         `json:"current"`
	Size    int         `json:"size"`
	List    interface{} `json:"list"`
	Total   int64       `json:"total"`
}

func ReturnSuccessWithPage(ctx *gin.Context, size int, current int, data interface{}, total int64) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": constants.SUCCESS,
		"result": &PageInfo{
			Current: current,
			Size:    size,
			List:    data,
			Total:   total,
		},
		"message": constants.SUCCESS_MSG,
	})
}
func ReturnFailWithPage(ctx *gin.Context, code int, msg string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"result": &PageInfo{
			Current: 0,
			Size:    0,
			List:    nil,
			Total:   0,
		},
		"message": msg,
	})
}
