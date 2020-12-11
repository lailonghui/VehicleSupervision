/*
@Time : 2020/12/11 14:41
@Author : lai
@Description :
@File : user_operation_log
*/
package model

import "time"

// 交警大队窗口查询违章记录表
//
//
// columns and relationships of "user_operation_log"
type UserOperationLog struct {
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 创建人
	CreatedBy string `json:"created_by"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人
	DeletedBy *string `json:"deleted_by"`
	// 主键
	ID int64 `json:"id"`
	// 违法时间
	IllegalTime *time.Time `json:"illegal_time"`
	// 车牌号
	LicensePlateNumber *string `json:"license_plate_number"`
	// 修改时间
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人
	UpdatedBy *string `json:"updated_by"`
	// 用户IP
	UserIP *string `json:"user_ip"`
}

func (UserOperationLog) TableName() string {
	return "user_operation_log"
}
