package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden SimCardMgrUnionPkLoader string *VehicleSupervision/internal/modules/device/model.SimCardMgr

// 数据库表名
func (t *SimCardMgr) TableName() string {
	return "sim_card_mgr"
}

// 主键列名
func (t *SimCardMgr) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *SimCardMgr) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *SimCardMgrPkLoader) NewLoader() *SimCardMgrPkLoader {
	return &SimCardMgrPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*SimCardMgr, []error) {
			var rs []*SimCardMgr
			var m SimCardMgr
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t *SimCardMgr) UnionPrimaryColumnName() string {
	return "mgr_id"
}

// 获取联合主键
func (t *SimCardMgr) GetUnionPrimary() string {
	return t.MgrID
}

// 新建联合主键dataloader
func (t *SimCardMgrUnionPkLoader) NewLoader() *SimCardMgrUnionPkLoader {
	return &SimCardMgrUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*SimCardMgr, []error) {
			var rs []*SimCardMgr
			var m SimCardMgr
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
