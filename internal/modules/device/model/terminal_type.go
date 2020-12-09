package model

import "time"

// 终端类型
//
//
// columns and relationships of "terminal_type"
type TerminalType struct {
	// 创建时间
	CreateAt time.Time `json:"create_at"`
	// 创建人
	CreateBy *string `json:"create_by"`
	// 删除时间
	DeleteAt *time.Time `json:"delete_at"`
	// 删除人
	DeleteBy *string `json:"delete_by"`
	// ID
	ID int64 `json:"id"`
	// 是否删除
	IsDelete bool `json:"is_delete"`
	// 终端类型
	ProtocolName string `json:"protocol_name"`
	// 备注
	Remark *string `json:"remark"`
	// 终端类型ID
	TypeID string `json:"type_id"`
	// 修改时间
	UpdateAt *time.Time `json:"update_at"`
	// 修改人
	UpdateBy *string `json:"update_by"`
}

func (t TerminalType) TableName() string {
	return "terminal_types"
}
