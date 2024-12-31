package article

import (
	"blog-server/internal/service/article"
)

type Handler struct {
	service article.Service
}

func NewHandler() *Handler {
	return &Handler{
		service: article.NewService(),
	}
}
