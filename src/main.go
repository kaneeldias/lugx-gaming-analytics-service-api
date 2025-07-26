package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
)

func main() {
	_, err := GetDatabaseConnection()
	if err != nil {
		log.Fatal("Error initializing database: ", err)
	}

	router := gin.Default()
	router.Use(CORSMiddleware())
	router.GET("/", healthCheck)
	router.POST("/record-page-view", recordPageView)
	router.POST("/record-click", recordClick)

	err = router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		log.Println("Error starting server: ", err)
		return
	}
}

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": fmt.Sprintf("Server is running with tag %s", os.Getenv("TAG")),
	})
}

func recordPageView(c *gin.Context) {
	var request RecordPageViewRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	ipAddress := c.ClientIP()

	err := SavePageView(request.Path, ipAddress)
	if err != nil {
		log.Println("Error saving page view: ", err)
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(200, gin.H{"message": "Page view recorded successfully"})
}

func recordClick(c *gin.Context) {
	var request RecordClickRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	ipAddress := c.ClientIP()

	err := SaveClick(request.Path, request.Element, ipAddress)
	if err != nil {
		log.Println("Error saving click: ", err)
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(200, gin.H{"message": "Click recorded successfully"})
}
