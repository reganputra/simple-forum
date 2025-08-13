package membership

import (
	"fmt"
	"net/http"
	"simple-forum/internal/model/membership"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUp(c *gin.Context) {
	ctx := c.Request.Context()

	var req membership.SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Printf("Received signup request: Email='%s', Username='%s'\n", req.Email, req.Username)

	err := h.membershipSvc.SignUp(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user": gin.H{
			"email":    req.Email,
			"username": req.Username,
		},
	})
}
