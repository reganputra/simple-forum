package membership

import (
	"context"
	"github.com/gin-gonic/gin"
	"simple-forum/internal/model/membership"
)

type membershipSevice interface {
	SignUp(ctx context.Context, req *membership.SignUpRequest) error
}

type Handler struct {
	engine        *gin.Engine
	membershipSvc membershipSevice
}

func NewHandler(api *gin.Engine, membershipSvc membershipSevice) *Handler {
	return &Handler{
		engine:        api,
		membershipSvc: membershipSvc,
	}
}

func (h *Handler) RegisterRoutes() {
	r := h.engine.Group("/membership")
	r.GET("/ping", h.Ping)
	r.POST("/register", h.SignUp)
}
