package server

import (
	"libraryControlApp_backend/commonStructs"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	router.GET("/search", search)
	router.GET("/addAuthor", addAuthor)
	router.GET("/addType", addType)
	router.GET("/addPublisher", addPublisher)

	router.POST("/filter", filter)
	router.POST("/addBook", addBook)
}

func search(c *gin.Context) {
	searchTerm := c.DefaultQuery("searchTerm", "")
	data := Usecase.Search(searchTerm)

	c.JSON(200, gin.H{"data": data, "success": true, "errorCode": 0, "errorMessage": ""})

}

func addAuthor(c *gin.Context) {
	newAuthor := c.DefaultQuery("newAuthor", "")
	err := Usecase.AddAuthor(newAuthor)

	if err != nil {
		println("Filter error")
		c.JSON(500, gin.H{"data": nil, "success": false, "errorCode": 500, "errorMessage": "Author addition internal error"})
		return
	}
	c.JSON(200, gin.H{"data": "success", "success": true, "errorCode": 0, "errorMessage": ""})

}

func addType(c *gin.Context) {
	newType := c.DefaultQuery("newType", "")
	err := Usecase.AddAuthor(newType)

	if err != nil {
		println("Filter error")
		c.JSON(500, gin.H{"data": nil, "success": false, "errorCode": 500, "errorMessage": "Type addition internal error"})
		return
	}
	c.JSON(200, gin.H{"data": "success", "success": true, "errorCode": 0, "errorMessage": ""})
}

func addPublisher(c *gin.Context) {
	newPublisher := c.DefaultQuery("newPublisher", "")
	err := Usecase.AddAuthor(newPublisher)

	if err != nil {
		println("Filter error")
		c.JSON(500, gin.H{"data": nil, "success": false, "errorCode": 500, "errorMessage": "Publisher addition internal error"})
		return
	}
	c.JSON(200, gin.H{"data": "success", "success": true, "errorCode": 0, "errorMessage": ""})
}

func filter(c *gin.Context) {
	filters := &commonStructs.Filter{}
	err := c.BindJSON(filters)
	if err != nil {
		println("Filter error")
		c.JSON(500, gin.H{"data": nil, "success": false, "errorCode": 500, "errorMessage": "Filter data parse error"})
		return
	}
	data := Usecase.Filter(*filters)
	c.JSON(200, gin.H{"data": data, "success": true, "errorCode": 0, "errorMessage": ""})
}

func addBook(c *gin.Context) {
	newBook := &commonStructs.BookInput{}
	err := c.BindJSON(newBook)
	if err != nil {
		println("Filter error")
		c.JSON(500, gin.H{"data": nil, "success": false, "errorCode": 500, "errorMessage": "addBook data parse error"})
		return
	}
	err = Usecase.AddBook(*newBook)
	if err != nil {
		println("Filter error")
		c.JSON(500, gin.H{"data": nil, "success": false, "errorCode": 500, "errorMessage": "Book addition internal error"})
		return
	}
	c.JSON(200, gin.H{"data": "success", "success": true, "errorCode": 0, "errorMessage": ""})
}

// gin.H{"data": data, "success": true, "errorCode": 0, "errorMessage": ""}
// gin.H{"data": nil, "success": false, "errorCode": err.Code, "errorMessage": err.Message}
