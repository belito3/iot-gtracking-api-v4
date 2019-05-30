package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"iot-gtracking-api-v4/config"
	"iot-gtracking-api-v4/router"
)

func main(){

	port := config.AppConf.ServerPort
	app := gin.Default() // create gin app
	app.Use(CORSMiddleware())
	router.ApplyRoutes(app) // apply api router
	app.Run(port)  // listen to given port
}

// CORSMiddleware set header CORS
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}