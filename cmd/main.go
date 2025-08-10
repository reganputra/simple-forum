package main

import (
	"log"
	"simple-forum/internal/configs"
	"simple-forum/internal/handler/membership"
	"simple-forum/pkg/internalsql"

	membershipRepo "simple-forum/internal/repository/membership"
	membershipService "simple-forum/internal/service/membership"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
)

func main() {
	r := gin.Default()

	var (
		cfg *configs.Config
	)

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
	cfg = configs.GetConfig()
	log.Println("config:", cfg)

	// Connect to Database
	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Initialize Repository
	memberRepo := membershipRepo.NewRepository(db)
	// Initialize Service
	membershipSvc := membershipService.NewService(memberRepo)
	// Initialize  Handler
	membershipHandler := membership.NewHandler(r, membershipSvc)
	// Set up routes
	membershipHandler.RegisterRoutes()
	r.Run(cfg.Service.Port)
}
