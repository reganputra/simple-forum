package membership

import (
	"simple-forum/internal/configs"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	engine *gin.Engine
	config *configs.Config
}

func NewHandler(api *gin.Engine, cfg *configs.Config) *Handler {
	return &Handler{
		engine: api,
		config: cfg,
	}
}

func (h *Handler) PingRoutes() {
	r := h.engine.Group("/membership")
	r.GET("/ping", h.Ping)
}
