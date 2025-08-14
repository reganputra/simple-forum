package posts

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllPost(c *gin.Context) {
	pageIdxStr := c.Query("pageIndex")
	pageSizeStr := c.Query("pageSize")

	pageIndex, err := strconv.Atoi(pageIdxStr)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": errors.New("invalid Index").Error(),
		})
		return
	}
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": errors.New("invalid Size").Error(),
		})
	}

	resp, err := h.postSvc.GetAllPost(c.Request.Context(), pageIndex, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}
