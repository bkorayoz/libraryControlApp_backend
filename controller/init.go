package server

import (
	"MapApp/common/logger"
	"errors"
	"libraryControlApp_backend/interactors"
	"libraryControlApp_backend/interfaces"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
)

const ApiKey = "6ItOQZeFus1Rw3NODa3b5JZNIgMFzc4c"

var Usecase interfaces.LibraryUsecase

func StartServer() {
	router := gin.New()
	router.Use(gin.Recovery(), Logger(), Headers())
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	InitRoutes(router)

	Usecase = interactors.LibraryInteractor{}

	log.Fatal(http.ListenAndServe(":8090", router))
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

func Headers() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Run this on all requests
		// Should be moved to a proper middleware
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Token,Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,GET,HEAD,POST,PUT,OPTIONS,TRACE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		}

		c.Next()
	}
}

func ApiKeyCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.Request.Header.Get("Authorization")
		apiKey = strings.TrimLeft(apiKey, "Bearer ")
		if apiKey != ApiKey {
			// deny the request, show an error
			println("WrongKey")
			logger.LogError("Api Key does not match. Request denied")
			c.AbortWithError(401, errors.New("Api Key is wrong. Request denied."))
			return
		}

	}

}
