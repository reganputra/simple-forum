package membership

import (
	"net/http"
	"simple-forum/internal/model/membership"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(c *gin.Context) {
	ctx := c.Request.Context()

	var req membership.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	accessToken, err := h.membershipSvc.Login(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp := membership.LoginResponse{
		AccessToken: accessToken,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   resp,
	})
}
