// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// 最新终端检查
type TerminalCheckLast struct {
	// ID
	ID int64 `json:"id"`
	// 最新终端检查ID
	TerminalCheckLastID string `json:"terminal_check_last_id"`
	// 终端ID
	TerminalID string `json:"terminal_id"`
	// acc状态
	Acc *string `json:"acc"`
	// 制动信号（刹车）
	Brake *string `json:"brake"`
	// 左转向灯信号
	LeftLamp *string `json:"left_lamp"`
	// 右转向灯信号
	RightLamp *string `json:"right_lamp"`
	// 近光灯信号
	NearLamp *string `json:"near_lamp"`
	// 远光灯信号（大灯）
	FarLamp *string `json:"far_lamp"`
	// 喇叭信号
	LoudSpeaker *string `json:"loud_speaker"`
	// 定位状态
	Locate *string `json:"locate"`
	// 前门
	FrontDoor *string `json:"front_door"`
	// gps信号
	GpsOpen *string `json:"gps_open"`
	// 北斗信号
	BdOpen *string `json:"bd_open"`
	// 摄像头
	Camera *string `json:"camera"`
	// 车速
	VehicleSpeed *string `json:"vehicle_speed"`
	// 备注
	Remark *string `json:"remark"`
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 创建人
	CreatedBy *string `json:"created_by"`
	// 修改时间
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人
	UpdatedBy *string `json:"updated_by"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人
	DeletedBy *string `json:"deleted_by"`
	// 是否删除
	IsDeleted bool `json:"is_deleted"`
}
