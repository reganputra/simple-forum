package main

import (
	"simple-forum/internal/handler/membership"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
)

func main() {
	r := gin.Default()

	membershipHandler := membership.NewHandler(r)

	membershipHandler.PingRoutes()

	r.Run(":8080")
}
