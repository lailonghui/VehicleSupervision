package model

import "time"

// 黑名单操作记录表
//
//
// columns and relationships of "blacklist_operation_record"
type BlacklistOperationRecord struct {
	// 黑名单记录ID
	BlacklistRecordID string `json:"blacklist_record_id"`
	// 黑名单类别
	BlacklistType int `json:"blacklist_type"`
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
	// 操作类别
	Operate int `json:"operate"`
	// 备注
	Remarks *string `json:"remarks"`
	// 修改时间
	UpdateAt *time.Time `json:"update_at"`
	// 修改人
	UpdateBy *string `json:"update_by"`
	// 驾驶员ID或者企业ID或者车辆ID
	VSeqn string `json:"v_seqn"`
}

func (b BlacklistOperationRecord) TableName() string {
	return "blacklist_operation_record"
}
