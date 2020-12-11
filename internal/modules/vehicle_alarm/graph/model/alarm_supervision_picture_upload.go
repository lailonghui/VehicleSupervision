/*
@Time : 2020/12/11 14:39
@Author : lai
@Description :
@File : alarm_supervision_picture_upload
*/
package model

import "time"

// 报警监管图片上传表
//
//
// columns and relationships of "alarm_supervision_picture_upload"
type AlarmSupervisionPictureUpload struct {
	// 终端IMEI
	Imei *string `json:"IMEI"`
	// 联合主键
	AlarmSupervisionPictureID string `json:"alarm_supervision_picture_id"`
	// 摄像头ID字典
	CameraID *int `json:"camera_id"`
	// 驾驶员id
	DriverID *string `json:"driver_id"`
	// 所在企业id
	EnterpriseID *string `json:"enterprise_id"`
	// 主键
	ID int64 `json:"id"`
	// 报警监控图片地址
	MonitoringPicAddress *string `json:"monitoring_pic_address"`
	// 报警监控图片名称
	MonitoringPicName *string `json:"monitoring_pic_name"`
	// 报警监控图片上传时间
	MonitoringPicUploadTime *time.Time `json:"monitoring_pic_upload_time"`
	// 拍照条件字典
	PhotoCondition *string `json:"photo_condition"`
	// SIM卡号
	SimNumber *string `json:"sim_number"`
	// 终端上报时间
	UpdateTime *time.Time `json:"update_time"`
	// 车辆ID
	VehicleID *string `json:"vehicle_id"`
}

func (AlarmSupervisionPictureUpload) TableName() string {
	return "alarm_supervision_picture_upload"
}
