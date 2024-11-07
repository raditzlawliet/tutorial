package main

import (
	"go-3-db-digi/handler"
	"go-3-db-digi/middleware"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	// Database
	db := NewDatabase()
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("failed to get DB from GORM:", err)
	}
	defer sqlDB.Close()

	// secret-key
	signingKey := os.Getenv("SIGNING_KEY")

	// import "github.com/gin-contrib/cors"
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AddAllowHeaders("Authorization") // Allow Header: Authorization
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	// grouping route with /auth
	authHandler := handler.NewAuth(db, []byte(signingKey))
	authRoute := r.Group("/auth")
	authRoute.POST("/login", authHandler.Login)
	authRoute.POST("/upsert", authHandler.Upsert)

	// grouping route with /account
	accountHandler := handler.NewAccount(db)
	accountRoutes := r.Group("/account")
	accountRoutes.POST("/create", accountHandler.Create)
	accountRoutes.GET("/read/:id", accountHandler.Read)
	accountRoutes.PATCH("/update/:id", accountHandler.Update)
	accountRoutes.DELETE("/delete/:id", accountHandler.Delete)
	accountRoutes.GET("/list", accountHandler.List)

	accountRoutes.GET("/my", middleware.AuthMiddleware(signingKey), accountHandler.My)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func NewDatabase() *gorm.DB {
	// dsn := "host=localhost port=5432 user=postgres dbname=digi sslmode=disable TimeZone=Asia/Jakarta"
	dsn := os.Getenv("DATABASE")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get DB object: %v", err)
	}

	var currentDB string
	err = sqlDB.QueryRow("SELECT current_database()").Scan(&currentDB)
	if err != nil {
		log.Fatalf("failed to query current database: %v", err)
	}

	log.Printf("Current Database: %s\n", currentDB)

	return db
}
