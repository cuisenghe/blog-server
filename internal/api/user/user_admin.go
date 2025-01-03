package user

import (
	"blog-server/internal/api"
	"blog-server/internal/common/constants"
	"blog-server/internal/common/response"
	"blog-server/internal/service/user"
	"github.com/gin-gonic/gin"
)

// 管理员相关接口

func (h *Handler) UpdateRole(ctx *gin.Context) {
	//获取参数

}

type UpdateUserInfoReq struct {
	ID       uint64 `json:"id"`
	Nickname string `json:"nick_name"`
	Avatar   string `json:"avatar"`
}

// UpdateUserInfo

func (h *Handler) UpdateUserInfo(ctx *gin.Context) {
	//binding
	var req UpdateUserInfoReq
	api.Binding(ctx, &req)
	// 方法
	updateInfo, err := h.service.AdminUpdateUserInfo(ctx, &user.AdminUserInfoData{
		ID:       req.ID,
		NickName: req.Nickname,
		Avatar:   req.Avatar,
	})
	if err != nil {
		// 返回错误
		response.Fail(ctx, constants.FAIL, err.Error())
	}
	response.Success(ctx, updateInfo)
}

// 修改用户角色
func (h *Handler) UpdateUserRole(ctx *gin.Context) {
	// binding

}
