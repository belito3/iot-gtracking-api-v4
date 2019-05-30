package controller

import (
	"github.com/gin-gonic/gin"
	"iot-gtracking-api-v4/service"
	"net/http"
)

type SysDeviceController struct {
	sysDeviceService *service.SysDeviceService
}

type SysDeviceRequest struct {
	DeviceId string `json:"device_id"`
	CustomerId string `json:"customer_id"`
}

func NewSysDeviceController() *SysDeviceController {
	return &SysDeviceController{sysDeviceService: service.NewSysDeviceService()}
}

func(ctl *SysDeviceController) GetLastReport(c *gin.Context){
	var req SysDeviceRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		responseError(c, http.StatusBadRequest, err.Error())
		return
	}

	if req.DeviceId == "" || req.CustomerId == "" {
		responseError(c, http.StatusBadRequest, "Invalid Parameter")
		return
	}

	info, err := ctl.sysDeviceService.GetLastReport(req.DeviceId, req.CustomerId)
	if err != nil {
		responseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	responseSucess(c, http.StatusOK, info)
	return
}