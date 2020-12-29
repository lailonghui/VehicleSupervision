package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden EnterprisePoliceSmsInfoLoader string *VehicleSupervision/internal/modules/admin/model.EnterprisePoliceSmsInfo

// 数据库表名
func (t EnterprisePoliceSmsInfo) TableName() string {
	return "enterprise_police_sms_info"
}

// 主键列名
func (t EnterprisePoliceSmsInfo) PrimaryColumnName() string {
	return "id"
}

// 联合主键列名
func (t EnterprisePoliceSmsInfo) UnionPrimaryColumnName() string {
	return "enterprise_police_sms_info_id"
}

// 新建dataloader
func (t *EnterprisePoliceSmsInfo) NewLoader() *EnterprisePoliceSmsInfoLoader {
	return &EnterprisePoliceSmsInfoLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EnterprisePoliceSmsInfo, []error) {
			var rs []*EnterprisePoliceSmsInfo

			db.DB.Model(&EnterprisePoliceSmsInfo{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)

			return rs, nil
		},
	}
}
