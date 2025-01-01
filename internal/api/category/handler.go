package category

import "blog-server/internal/service/category"

type Handler struct {
	service category.Service
}

func NewHandler() *Handler {
	return &Handler{
		service: category.NewService(),
	}
}
