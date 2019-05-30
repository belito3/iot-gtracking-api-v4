package router

import (
	"github.com/gin-gonic/gin"
	"iot-gtracking-api-v4/controller"
)

func SysDeviceRouter(r *gin.RouterGroup){
	sysDeviceController := controller.NewSysDeviceController()

	client := r.Group("/gps")
	{
		client.POST("/last-report", sysDeviceController.GetLastReport)
	}
}