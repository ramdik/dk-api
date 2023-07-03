package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	/* //membuat versioning
	v1 := router.Group("/v1")

	// http GET route (root, search by parameter :id, search by query)
	// menggunakan query?key=value
	v1.GET("/", getHandler)
	v1.GET("/book/:id/:title", getParamHandler)
	v1.GET("/query", getQueryHandler)

	// http POST route
	v1.POST("/book", postBookHandler) */

	router.GET("/", geBooktHandler)

	// go run main.go
	router.Run()
}

// http GET all
/* func getHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "afi",
		"age":  "20",
	})
}

// http GET with query
func getQueryHandler(c *gin.Context) {
	title := c.Query("title")
	jumlah := c.Query("jumlah")

	c.JSON(http.StatusOK, gin.H{
		"title":  title,
		"jumlah": jumlah,
	})
}

// http GET with id, and tittle
func getParamHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

type QcodeInput struct {
	Title string `json:"title" binding:"required"`
	Qcode string `json:"qcode" binding:"required"`
}

// http POST
func postBookHandler(c *gin.Context) {
	var qcodeInput QcodeInput

	err := c.ShouldBindJSON(&qcodeInput)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on failed %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
		//log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"title": qcodeInput.Title,
		"qcode": qcodeInput.Qcode,
	})

} */

//new code api

// Object map data model
type Book struct {
	Id    string  `json:"id"`
	Title string  `json:"title"`
	Write string  `json:"write"`
	Price float64 `json:"price"`
}

// data dummy slice (obj any)
var books = []Book{
	{Id: "1", Title: "Kisah shope", Write: "jarot", Price: 10000.0},
	{Id: "2", Title: "Negeri berdarah putih", Write: "mamat", Price: 144000.0},
	{Id: "3", Title: "Trik membeli barang murah shopee", Write: "karin", Price: 690000.99},
	{Id: "4", Title: "Cara jadi influencer", Write: "Malih", Price: 99999.9999},
}

// http GET All data
func geBooktHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}
