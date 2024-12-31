package user

import (
	"blog-server/internal/api"
	"blog-server/internal/repository/userDao"
	"blog-server/internal/service/user"
	"github.com/gin-gonic/gin"
)

type UpdateUserReq struct {
	ID       uint64 `json:"id"`
	Avatar   string `json:"avatar"`
	NickName string `json:"nick_name"`
	QQ       string `json:"qq"`
}
type UpdatePasswordReq struct {
	ID        uint64 `json:"id"`
	Password  string `json:"password"`
	Password1 string `json:"password1"`
	Password2 string `json:"password2"`
}

func (h *Handler) UpdateOwnUserInfo(ctx *gin.Context) {
	//获取参数
	var req *UpdateUserReq
	api.Binding(ctx, req)
	// 更新
	info, err := h.service.UpdateUserInfo(ctx, &userDao.BlogUser{
		ID:       req.ID,
		Avatar:   req.Avatar,
		NickName: req.NickName,
		QQ:       req.QQ,
	})
	if err != nil {
		api.ReturnBizError(ctx, err)
	}
	api.ReturnSuccess(ctx, info)
}

// 更改用户密码
func (h *Handler) UpdatePassword(ctx *gin.Context) {
	// 绑定
	var req *UpdatePasswordReq
	api.Binding(ctx, req)
	// 服务
	isSuccess, err := h.service.UpdatePassword(ctx, &user.UserPasswordData{
		ID:        req.ID,
		Password:  req.Password,
		Password1: req.Password1,
		Password2: req.Password2,
	})
	if err != nil {
		api.ReturnBizError(ctx, err)
		return
	}
	api.ReturnSuccess(ctx, isSuccess)
}
