package main

import (
	"log"
	"simple-forum/internal/configs"
	"simple-forum/internal/handler/membership"
	"simple-forum/internal/handler/posts"
	postRepo "simple-forum/internal/repository/posts"
	postSvc "simple-forum/internal/service/posts"
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
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Initialize Repository
	memberRepo := membershipRepo.NewRepository(db)
	postRepository := postRepo.NewRepository(db)
	// Initialize Service
	membershipSvc := membershipService.NewService(cfg, memberRepo)
	postService := postSvc.NewService(cfg, postRepository)
	// Initialize  Handler
	membershipHandler := membership.NewHandler(r, membershipSvc)
	postHandler := posts.NewHandler(r, postService)
	// Set up routes
	membershipHandler.RegisterRoutes()
	postHandler.PostRoutes()
	r.Run(cfg.Service.Port)
}
