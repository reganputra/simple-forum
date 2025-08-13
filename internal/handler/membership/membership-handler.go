package membership

import (
	"context"
	"simple-forum/internal/model/membership"

	"github.com/gin-gonic/gin"
)

type membershipSevice interface {
	SignUp(ctx context.Context, req *membership.SignUpRequest) error
	Login(ctx context.Context, req membership.LoginRequest) (string, error)
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
	r.POST("/login", h.Login)
}
