package main

import (
	"log"
	"simple-forum/internal/configs"
	"simple-forum/internal/handler/membership"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
)

func main() {
	r := gin.Default()

	// Initialize config
	err := configs.Init(
		configs.WithConfigFolder([]string{"./internal/configs"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)
	if err != nil {
		log.Fatalf("Error initializing config: %v", err)
	}

	// Get the initialized config
	cfg := configs.GetConfig()
	log.Println("config:", cfg)

	membershipHandler := membership.NewHandler(r, cfg)

	membershipHandler.PingRoutes()

	r.Run(cfg.Service.Port)
}
