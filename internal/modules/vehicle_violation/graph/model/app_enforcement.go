/*
@Time : 2020/12/11 14:41
@Author : lai
@Description :
@File : app_enforcement
*/
package model

import "time"

// APP现场执法表
//
//
// columns and relationships of "app_enforcement"
type AppEnforcement struct {
	// 空间数据类型point表示经度(longitude)和纬度(latitude)
	Coordinate *string `json:"coordinate"`
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 创建人
	CreatedBy string `json:"created_by"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人
	DeletedBy *string `json:"deleted_by"`
	// 编辑文本
	EditText *string `json:"edit_text"`
	// 企业类型字典
	EnterpriseType *string `json:"enterprise_type"`
	// 主键
	ID int64 `json:"id"`
	// 联合主键
	IllegalPhotoID string `json:"illegal_photo_id"`
	// 位置描述
	LocationDescription *string `json:"location_description"`
	// 操作用户
	OperationUser *string `json:"operation_user"`
	// 纠察状态
	PicketStatus *int `json:"picket_status"`
	// 修改时间
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人
	UpdatedBy *string `json:"updated_by"`
	// vehicle_info 车辆信息表 的vehicle_id
	VehicleID *string `json:"vehicle_id"`
	// 违法明细表ID(vehicle_violation_details的violation_detail_id)
	ViolationDetailID *string `json:"violation_detail_id"`
}

func (AppEnforcement) TableName() string {
	return "app_enforcement"
}
