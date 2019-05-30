package controller

import "github.com/gin-gonic/gin"

type Request struct {

}

// responseError
func responseError(c *gin.Context, statusCode int, message string){
	c.JSON(statusCode, gin.H{
		"StatusCode": statusCode,
		"Message": message,
	})
}

// responseSucess
func responseSucess(c *gin.Context, statusCode int, data interface{}){
	c.JSON(statusCode, gin.H{
		"StatusCode": statusCode,
		"Message": "Success",
		"Data": data,
	})
}