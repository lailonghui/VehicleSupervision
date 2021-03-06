// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// 终端变更
type TerminalChange struct {
	// ID
	ID int64 `json:"id"`
	// 终端变更ID
	ChangeID string `json:"change_id"`
	// 新车牌
	NewPlateNumber string `json:"new_plate_number"`
	// 终端ID
	TerminalID string `json:"terminal_id"`
	// 终端型号ID
	TerminalTypeID *string `json:"terminal_type_id"`
	// 旧车牌
	OldPlateNumber string `json:"old_plate_number"`
	// 所在部门ID
	DeptID *string `json:"dept_id"`
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 创建人
	CreatedBy *string `json:"created_by"`
	// 备注
	Remark *string `json:"remark"`
}
