package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Hello World")

	// Connect to the database
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost", "5432", "vicccccw", "password", "sample-web3")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	// Obtain the underlying *sql.DB object
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get *sql.DB: %v", err)
	}

	// Ensure the database connection is closed when main function exits
	defer func() {
		if err := sqlDB.Close(); err != nil {
			log.Fatalf("failed to close the database: %v", err)
		}
	}()
	fmt.Println("Connected to database")

	r := gin.Default()

	// resolve cross origin
	r.Use(corsMiddleware())

	r.Use(dbMiddleware(db))

	// Initialize the routes
	initializeRoutes(r)

	r.Run() // listen and serve on 0.0.0.0:8080
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	}
}

func initializeRoutes(router *gin.Engine) {
	// Group routes under ""
	api := router.Group("")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
}

func dbMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
