package main

import (
	"fmt"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"net/http"
	"skingenius/database"
	"skingenius/handlers"
	"strings"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
)

func main() {
	_, err := database.NewClient(host, port, user, password)
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to establich db connection, error: %v", err))
	}

	router := gin.Default()
	router.Use(gin.Recovery(), gin.Logger(), CORSMiddleware())

	router.POST("/match", handlers.FindMatch)

	router.Use(SPAMiddleware("/", "./skingenius"))

	router.Run(":8080")
	fmt.Println(fmt.Sprintf("Start server on port %d", 8080))
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

func SPAMiddleware(urlPrefix, spaDirectory string) gin.HandlerFunc {

	directory := static.LocalFile(spaDirectory, true)
	fileserver := http.FileServer(directory)
	if urlPrefix != "" {
		fileserver = http.StripPrefix(urlPrefix, fileserver)
	}
	return func(c *gin.Context) {

		if strings.Contains(c.Request.URL.Path, "api") {
			c.Next()
		}

		if directory.Exists(urlPrefix, c.Request.URL.Path) {
			fileserver.ServeHTTP(c.Writer, c.Request)
			c.Abort()
		} else {
			c.Request.URL.Path = "/"
			fileserver.ServeHTTP(c.Writer, c.Request)
			c.Abort()
		}
	}
}
