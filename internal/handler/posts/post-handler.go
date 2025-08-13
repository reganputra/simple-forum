package posts

import (
	"context"
	"simple-forum/internal/middleware"
	"simple-forum/internal/model/posts"

	"github.com/gin-gonic/gin"
)

type postService interface {
	CreatePost(ctx context.Context, userId int64, req posts.CreatePostRequest) error
}

type Handler struct {
	engine  *gin.Engine
	postSvc postService
}

func NewHandler(api *gin.Engine, postSvc postService) *Handler {
	return &Handler{
		engine:  api,
		postSvc: postSvc,
	}
}

func (h *Handler) PostRoutes() {
	r := h.engine.Group("/posts")
	r.Use(middleware.AuthMiddleware())
	r.POST("/create-post", h.CreatePost)
}
