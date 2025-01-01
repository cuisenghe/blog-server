package user

import (
	"blog-server/internal/api"
	"github.com/gin-gonic/gin"
	"log"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Response struct {
	Token     string `json:"token"`
	Username  string `json:"username"`
	Role      int    `json:"role"`
	Id        int    `json:"id"`
	IpAddress string `json:"ipAddress"`
}

func (h *Handler) Login(ctx *gin.Context) {
	var request LoginRequest
	api.Binding(ctx, &request)
	resp, err := h.service.Login(ctx, request.Username, request.Password)
	if err != nil {
		log.Println("login error", err)
		api.ReturnBizErrorWithData(ctx, err, resp)
		return
	}
	api.ReturnSuccess(ctx, resp)
}
