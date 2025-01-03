package user

import (
	"blog-server/internal/api"
	"blog-server/internal/common/constants"
	"blog-server/internal/common/response"
	"blog-server/internal/service/user"
	"github.com/gin-gonic/gin"
)

// 分页获取用户列表
type GetListRequest struct {
	Current  int    `json:"current"`
	Size     int    `json:"size"`
	NickName string `json:"nick_name"`
	Role     int    `json:"role"`
}

func (h *Handler) GetUserList(ctx *gin.Context) {
	var req GetListRequest
	api.Binding(ctx, &req)
	// service
	resp, err := h.service.GetUserList(ctx, &user.GetUserListData{
		Current:  req.Current,
		Size:     req.Size,
		NickName: req.NickName,
		Role:     req.Role,
	})
	if err != nil {
		response.Fail(ctx, constants.FAIL, err.Error())
	}
	response.SuccessWithPage(ctx, resp)
}
