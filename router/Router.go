package router

import "github.com/gin-gonic/gin"

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api/v1.0/")
	{
		SysDeviceRouter(api)
	}
}