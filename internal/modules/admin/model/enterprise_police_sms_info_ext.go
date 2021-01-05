package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden EnterprisePoliceSmsInfoUnionPkLoader string *VehicleSupervision/internal/modules/admin/model.EnterprisePoliceSmsInfo

// 数据库表名
func (t EnterprisePoliceSmsInfo) TableName() string {
	return "enterprise_police_sms_info"
}

// 主键列名
func (t EnterprisePoliceSmsInfo) PrimaryColumnName() string {
	return "id"
}

// 新建主键dataloader
func (t *EnterprisePoliceSmsInfoPkLoader) NewLoader() *EnterprisePoliceSmsInfoPkLoader {
	return &EnterprisePoliceSmsInfoPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EnterprisePoliceSmsInfo, []error) {
			var rs []*EnterprisePoliceSmsInfo
			var m EnterprisePoliceSmsInfo
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t EnterprisePoliceSmsInfo) UnionPrimaryColumnName() string {
	return "enterprise_police_sms_info_id"
}

// 新建联合主键dataloader
func (t *EnterprisePoliceSmsInfoUnionPkLoader) NewLoader() *EnterprisePoliceSmsInfoUnionPkLoader {
	return &EnterprisePoliceSmsInfoUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EnterprisePoliceSmsInfo, []error) {
			var rs []*EnterprisePoliceSmsInfo
			var m EnterprisePoliceSmsInfo
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
