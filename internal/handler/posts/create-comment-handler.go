package posts

import (
	"net/http"
	"simple-forum/internal/model/posts"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateComment(c *gin.Context) {
	ctx := c.Request.Context()

	var req posts.CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	postIdStr := c.Param("postId")
	postId, err := strconv.ParseInt(postIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "postId is not valid",
		})
		return
	}
	userId := c.GetInt64("userId")

	err = h.postSvc.CreateComment(ctx, postId, userId, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Comment created successfully",
		"postId":  postId,
		"comment": req.CommentContent,
	})
}
