package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden SimCardFlowLoader string *VehicleSupervision/internal/modules/device/model.SimCardFlow

// sim卡流量信息
//
//
// columns and relationships of "sim_card_flow"
type SimCardFlow struct {
	// 已产生流量卡均流量(KB)
	CardAvgFlow *float64 `json:"card_avg_flow"`
	// 卡号备注
	CardNoRemark *string `json:"card_no_remark"`
	// 创建时间
	CreateAt time.Time `json:"create_at"`
	// 创建人
	CreateBy *string `json:"create_by"`
	// 删除时间
	DeleteAt *time.Time `json:"delete_at"`
	// 删除人
	DeleteBy *string `json:"delete_by"`
	// 企业ID
	EnterpriseID *string `json:"enterprise_id"`
	// ICCID
	Iccid *string `json:"iccid"`
	// ID
	ID int64 `json:"id"`
	// 物联卡号
	SimCardID string `json:"sim_card_id"`
	// 是否删除
	IsDelete bool `json:"is_delete"`
	// 是否共享池
	IsSharePool bool `json:"is_share_pool"`
	// 流量池卡均流量(KB)
	PoolAvgFlow *float64 `json:"pool_avg_flow"`
	// SIM卡流量ID
	SimCardFlowID string `json:"sim_card_flow_id"`
	// 套餐流量(MB)
	SuitFlow *float64 `json:"suit_flow"`
	// 套餐剩余流量(MB)
	SuitLeftFlow *float64 `json:"suit_left_flow"`
	// 已超套餐流量(MB)
	SuitOverFlow *float64 `json:"suit_over_flow"`
	// 套餐剩余短信条数
	SuitSmsLeftNum *int `json:"suit_sms_left_num"`
	// 套餐短信条数
	SuitSmsNum *int `json:"suit_sms_num"`
	// 已超套餐短信条数
	SuitSmsOverNum *int `json:"suit_sms_over_num"`
	// 上行短信已用条数
	SuitUseSmsNum *int `json:"suit_use_sms_num"`
	// 修改时间
	UpdateAt *time.Time `json:"update_at"`
	// 修改人
	UpdateBy *string `json:"update_by"`
	// 数据服务已用流量(KB)
	UseFlow *float64 `json:"use_flow"`
}

func (s SimCardFlow) TableName() string {
	return "sim_card_flow"
}

func (u *SimCardFlow) NewLoader() *SimCardFlowLoader {
	return &SimCardFlowLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*SimCardFlow, []error) {
			var rs []*SimCardFlow
			db.DB.Model(&SimCardFlow{}).Where("sim_card_flow_id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
