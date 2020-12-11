/*
@Time : 2020/12/11 14:41
@Author : lai
@Description :
@File : vehicle_detain_details
*/
package model

import "time"

// 扣车明细表
//
//
// columns and relationships of "vehicle_detain_details"
type VehicleDetainDetails struct {
	// 卡口图片
	BayonetPicture *string `json:"bayonet_picture"`
	// 主键
	ID int64 `json:"id"`
	// 卡口判断(是否通过)
	IsBayonet *bool `json:"is_bayonet"`
	// 是否目录库
	IsCategory *bool `json:"is_category"`
	// 是否删除
	IsDelete *bool `json:"is_delete"`
	// 违法判断(是否通过)
	IsIllegal *bool `json:"is_illegal"`
	// 卫星定位判断(是否通过)
	IsSatelliteJudgment *bool `json:"is_satellite_judgment"`
	// 车牌颜色字典
	LicensePlateColor *string `json:"license_plate_color"`
	// 车牌号码
	LicensePlateNumber *string `json:"license_plate_number"`
	// 记录时间
	RecordTime *time.Time `json:"record_time"`
	// 登记所在机构
	RegisterRegion *string `json:"register_region"`
	// 登记用户
	RegisterUser *string `json:"register_user"`
	// 备注
	Remarks *string `json:"remarks"`
	// 更新时间
	UpdateTime *time.Time `json:"update_time"`
	// 扣车图片
	VehicleDetainPicture *string `json:"vehicle_detain_picture"`
	// 状态(车辆扣车状态字典)
	VehicleDetainStatus *int `json:"vehicle_detain_status"`
	// 扣车时间
	VehicleDetainTime *time.Time `json:"vehicle_detain_time"`
	// 放车图片
	VehicleReleasePicture *string `json:"vehicle_release_picture"`
	// 放车时间
	VehicleReleaseTime *time.Time `json:"vehicle_release_time"`
}

func (VehicleDetainDetails) TableName() string {
	return "vehicle_detain_details"
}
