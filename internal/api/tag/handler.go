package tag

import "blog-server/internal/service/tag"

type Handler struct {
	service tag.Service
}

func NewHandler() *Handler {
	return &Handler{
		service: tag.NewService(),
	}
}
