package model

import "gorm.io/gorm"

// 设备表
type Device struct {
	gorm.Model
	Mac    string // 设备mac地址
	Status bool   // 记录设备在线状态  true为在线
}

type DeviceData struct {
	gorm.Model
	MacID uint64 // 设备id  Device结构的主键
	Data  string // 设备记录的信息
}
