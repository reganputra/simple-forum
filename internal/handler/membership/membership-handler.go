package membership

import "github.com/gin-gonic/gin"

type Handler struct {
	engine *gin.Engine
}

func NewHandler(api *gin.Engine) *Handler {
	return &Handler{
		engine: api,
	}
}

func (h *Handler) PingRoutes() {
	r := h.engine.Group("/membership")
	r.GET("/ping", h.Ping)
}
