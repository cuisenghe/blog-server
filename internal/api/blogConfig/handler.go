package blogConfig

import "blog-server/internal/service/blogConfig"

type Handler struct {
	service blogConfig.Service
}

func NewHandler() *Handler {
	return &Handler{
		service: blogConfig.NewService(),
	}
}
