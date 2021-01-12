// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// sim卡使用情况管理
type SimCardMgr struct {
	// ID
	ID int64 `json:"id"`
	// sim卡使用情况管理ID
	MgrID string `json:"mgr_id"`
	// 状态
	Status *int `json:"status"`
	// 终端ID
	TerminalID *string `json:"terminal_id"`
	// 运营商类型
	MobileType *int `json:"mobile_type"`
	// 代理商ID
	ProxyrgID *string `json:"proxyrg_id"`
	// 部门ID
	DeptID *string `json:"dept_id"`
	// sim卡类型
	SimType *int `json:"sim_type"`
	// 领用用途
	UseType *int `json:"use_type"`
	// 领用人员
	UseUserID *string `json:"use_user_id"`
	// 应用系统名称
	SystemName *string `json:"system_name"`
	// sim卡号
	SimNumber *string `json:"sim_number"`
	// 是否外部来源
	IsOutside bool `json:"is_outside"`
	// 换卡原因
	UpdateCause *string `json:"update_cause"`
	// sim卡来自哪个区域
	SimArea *string `json:"sim_area"`
	// 去掉1和2位的sim卡号
	SimNumber12 *string `json:"sim_number_12"`
	// 去掉2和3位的sim卡号
	SimNumber23 *string `json:"sim_number_23"`
	// 外部sim卡号相匹配的sim卡
	SimBak *string `json:"sim_bak"`
	// 10位sim卡号
	SimNumber10 *string `json:"sim_number10"`
	// 注销时间
	CancelTime *time.Time `json:"cancel_time"`
	// 注销原因
	CancelCause *string `json:"cancel_cause"`
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
