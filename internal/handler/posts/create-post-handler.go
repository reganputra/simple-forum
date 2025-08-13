package posts

import (
	"net/http"
	"simple-forum/internal/model/posts"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreatePost(c *gin.Context) {
	ctx := c.Request.Context()

	var req posts.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userId := c.GetInt64("userId")
	err := h.postSvc.CreatePost(ctx, userId, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Post created successfully",
	})
}
