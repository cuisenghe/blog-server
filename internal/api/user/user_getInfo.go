package user

import (
	"blog-server/constants"
	"blog-server/internal/common/response"
	"blog-server/internal/repository/userDao"

	"github.com/gin-gonic/gin"

	"strconv"
)

type UserInfoResp struct {
	UserName string `json:"user_name"`
	Role     int    `json:"role"`
	NickName string `json:"nick_name"`
	Avatar   string `json:"avatar"`
	QQ       string `json:"qq"`
}

func (h *Handler) GetUserInfoById(ctx *gin.Context) {
	// 获取id
	id := ctx.Param("id")
	// string 转int64
	parseUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		response.Fail(ctx, constants.FAIL, err.Error())
		return
	}
	// 查询
	user, err := h.service.GetUserInfoById(ctx, parseUint)
	response.Success(ctx, map[string]interface{}{
		"user": convert(user),
	})
}
func convert(user *userDao.BlogUser) *UserInfoResp {
	return &UserInfoResp{
		UserName: user.Username,
		Role:     user.Role,
		NickName: user.NickName,
		Avatar:   user.Avatar,
		QQ:       user.QQ,
	}
}
