package model

import "time"

type SimCard struct {
	// 创建时间
	CreateAt time.Time `json:"create_at"`
	// 创建人
	CreateBy *string `json:"create_by"`
	// 删除时间
	DeleteAt *time.Time `json:"delete_at"`
	// 删除人
	DeleteBy *string `json:"delete_by"`
	// 所属部门ID
	DeptID *string `json:"dept_id"`
	// ID
	ID int64 `json:"id"`
	// 是否删除
	IsDelete bool `json:"is_delete"`
	// 运营商类型
	MobileType *int `json:"mobile_type"`
	// 运营商ID
	OperatorsID *string `json:"operators_id"`
	// 代理商ID
	ProxyrgID *string `json:"proxyrg_id"`
	// 备注
	Remark *string `json:"remark"`
	// SIM卡ID
	SimCardID string `json:"sim_card_id"`
	// SIM卡号
	SimNumber *string `json:"sim_number"`
	// 绑定终端ID
	TerminalID *string `json:"terminal_id"`
	// 修改时间
	UpdateAt *time.Time `json:"update_at"`
	// 修改人
	UpdateBy *string `json:"update_by"`
}

func (s SimCard) TableName() string {
	return "sim_card"
}
