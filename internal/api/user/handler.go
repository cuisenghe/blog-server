package user

import "blog-server/internal/service/user"

type Handler struct {
	service user.Service
}

func NewHandler() *Handler {
	return &Handler{
		service: user.NewService(),
	}
}
