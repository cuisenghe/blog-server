package user

import (
	"blog-server/constants"
	"blog-server/internal/api"
	"blog-server/internal/common/response"
	"blog-server/internal/service/user"
	"github.com/gin-gonic/gin"
	"log"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	NickName string `json:"nick_name" `
	QQ       string `json:"qq" `
	Avatar   string `json:"avatar" `
	Role     int    `json:"role" `
}
type RegisterResponse struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
}

func (h *Handler) Register(ctx *gin.Context) {
	// 参数绑定
	var req RegisterRequest
	api.Binding(ctx, &req)

	// 走service
	id, err := h.service.Register(ctx, &user.RegData{
		Username: req.Username,
		Password: req.Password,
		NickName: req.NickName,
		QQ:       req.QQ,
		Avatar:   req.Avatar,
		Role:     req.Role,
	})
	if err != nil {
		log.Println("register err:", err)
		response.FailWithData(ctx, constants.FAIL, err.Error(), &RegisterResponse{
			ID:       0,
			Username: req.Username,
		})
		return
	}
	response.Success(ctx, &RegisterResponse{
		ID:       id,
		Username: req.Username,
	})
}
