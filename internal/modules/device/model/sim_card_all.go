// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// sim卡管理（外部来源）
type SimCardAll struct {
	// ID
	ID int64 `json:"id"`
	// SIM卡管理（外部来源）ID
	SimcardAllID string `json:"simcard_all_id"`
	// 车牌号码
	PlateNumber *string `json:"plate_number"`
	// 车牌号码
	PlateColor *string `json:"plate_color"`
	// 所在部门ID
	DeptID *string `json:"dept_id"`
	// 代理商ID
	ProxyrgID *string `json:"proxyrg_id"`
	// 所在企业ID
	EnterpriseID *string `json:"enterprise_id"`
	// sim卡号
	SimNumber *string `json:"sim_number"`
	// 服务费到期时间
	ServiceEndTime *time.Time `json:"service_end_time"`
	// 终端ID
	TerminalID *string `json:"terminal_id"`
	// 应用于何种系统
	SystemName *string `json:"system_name"`
	// sim卡类型
	SimType *int `json:"sim_type"`
	// 用户ID
	UserID *string `json:"user_id"`
	// 运营商类型
	MobileType *int `json:"mobile_type"`
	// 解绑原因
	UpdateCause *string `json:"update_cause"`
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 创建人
	CreatedBy *string `json:"created_by"`
	// 更新时间
	UpdatedAt *time.Time `json:"updated_at"`
	// 更新人
	UpdatedBy *string `json:"updated_by"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人
	DeletedBy *string `json:"deleted_by"`
	// 是否删除
	IsDeleted bool `json:"is_deleted"`
	// 备注
	Remark *string `json:"remark"`
}
