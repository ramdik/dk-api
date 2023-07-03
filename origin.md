package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	router := gin.Default()

	//membuat versioning
	v1 := router.Group("/v1")

	// route URL HTTP method GET
	v1.GET("/", rootHandler)

	v1.GET("/test", testHandler)
	v1.GET("/book/:id/:title", bookHandler)
	// menggunakan query?key=value
	v1.GET("/query", queryHandler)

	//POST method
	v1.POST("/book", postBookHandler)

	// go run main.go
	router.Run()
}

// tanda * adalah pointer
// http GET
func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Ramadhika",
		"age":  "26",
		"tes":  "asdasdas",
		"asd":  "asdasd",
	})
}

func testHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "afi",
		"age":  "20",
	})
}

func bookHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

func queryHandler(c *gin.Context) {
	title := c.Query("title")
	jumlah := c.Query("jumlah")

	c.JSON(http.StatusOK, gin.H{
		"title":  title,
		"jumlah": jumlah,
	})
}

type QcodeInput struct {
	Title string `json:"title" binding:"required"`
	Qcode string `json:"qcode" binding:"required"`
}

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

	// nrasdp masdn  nasd  nn vnasd  vnas  nqsio
}
