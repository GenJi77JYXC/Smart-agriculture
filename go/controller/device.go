package controller

import (
	"demo/database"
	"demo/model"
	"demo/response"
	"github.com/gin-gonic/gin"
	"time"
)

func ListDevice(ctx *gin.Context) {
	DB := database.GetDB()
	var devices []model.Device
	DB.Table("devices").Find(&devices)
	response.Response(ctx, 200, 200, devices, "设备获取成功")
}

func GetDeviceData(ctx *gin.Context) {
	id := ctx.Query("id")
	DB := database.GetDB()
	var datas []model.DeviceData
	DB.Table("device_data").Where("mac_id = ?", id).Find(&datas)
	response.Response(ctx, 200, 200, datas, "设备信息列表获取成功")

}

func GetDeviceDataByTime(ctx *gin.Context) {
	id := ctx.Query("id")
	dateStart := ctx.Query("start")
	dateEnd := ctx.Query("end")
	start, _ := time.Parse("2006-01-02", dateStart)
	end, _ := time.Parse("2006-01-02", dateEnd)
	DB := database.GetDB()
	var datas []model.DeviceData
	DB.Table("device_data").Where("mac_id = ?", id).Where("created_at >= ?", start).Where("created_at <= ?", end).Find(&datas)
	response.Response(ctx, 200, 200, datas, "获取设备上线时间成功")
}
