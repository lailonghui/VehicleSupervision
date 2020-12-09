/*
@Time : 2020/12/9 10:50
@Author : lai
@Description :
@File : vehicle_supervision_photo
*/
package model

import "time"

// 车辆监控图片表
//
//
// columns and relationships of "vehicle_supervision_photo"
type VehicleSupervisionPhoto struct {
	// 终端IMEI
	Imei *string `json:"IMEI"`
	// 摄像头ID字典
	CameraID *int `json:"camera_id"`
	// 驾驶员id
	DriverID *string `json:"driver_id"`
	// 所在企业id
	EnterpriseID *string `json:"enterprise_id"`
	// 主键
	ID int64 `json:"id"`
	// 监控图片地址
	MonitoringPicAddress *string `json:"monitoring_pic_address"`
	// 监控图片名称
	MonitoringPicName *string `json:"monitoring_pic_name"`
	// 监控图片上传时间
	MonitoringPicUploadTime *time.Time `json:"monitoring_pic_upload_time"`
	// 拍照条件字典
	PhotoCondition *string `json:"photo_condition"`
	// SIM卡号
	SimNumber *string `json:"sim_number"`
	// 车辆监控图片外部编码，由golang程序生成的xid，暴露到外部使用
	SupervisionPhotoID string `json:"supervision_photo_id"`
	// 终端上报时间
	UpdateTime *time.Time `json:"update_time"`
	// vehicle_info表vehicle_id
	VehicleID string `json:"vehicle_id"`
}
