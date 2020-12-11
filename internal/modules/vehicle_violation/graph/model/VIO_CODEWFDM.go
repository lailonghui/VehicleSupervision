/*
@Time : 2020/12/11 14:45
@Author : lai
@Description :
@File : VIO_CODEWFDM
*/
package model

// 违法描述字典表(交警提供数据表，暂无修改)
//
//
// columns and relationships of "VIO_CODEWFDM"
type VioCodewfdm struct {
	// 最大罚款金额
	FkjeMax *float64 `json:"FKJE_MAX"`
	// 最小罚款金额
	FkjeMin *float64 `json:"FKJE_MIN"`
	// 违法计分数
	Wfjfs *float64 `json:"WFJFS"`
	// 违法描述
	Wfms *string `json:"WFMS"`
	// 违法行为
	Wfxw string `json:"WFXW"`
	// 序号
	Xh *string `json:"XH"`
}

func (VioCodewfdm) TableName() string {
	return "VIO_CODEWFDM"
}
