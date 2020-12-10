package model

import "time"

// 终端工厂
//
//
// columns and relationships of "terminal_factory"
type TerminalFactory struct {
	// 厂家地址
	Address *string `json:"address"`
	// 联系人
	Contact *string `json:"contact"`
	// 联系电话
	ContactPhone *string `json:"contact_phone"`
	// 创建时间
	CreateAt *time.Time `json:"create_at"`
	// 创建人
	CreateBy *string `json:"create_by"`
	// 删除时间
	DeleteAt *time.Time `json:"delete_at"`
	// 删除人
	DeleteBy *string `json:"delete_by"`
	// 工厂ID
	FactoryID string `json:"factory_id"`
	// 厂家名称
	FactoryName string `json:"factory_name"`
	// ID
	ID int64 `json:"id"`
	// 是否删除
	IsDelete bool `json:"is_delete"`
	// 备注
	Remark *string `json:"remark"`
	// 技术支持人
	TechContact *string `json:"tech_contact"`
	// 技术支持电话
	TechContactPhone *string `json:"tech_contact_phone"`
	// 更新时间
	UpdateAt *time.Time `json:"update_at"`
	// 更新人
	UpdateBy *string `json:"update_by"`
}

func (t TerminalFactory) TableName() string {
	return "terminal_factory"
}
