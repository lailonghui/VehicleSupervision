package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden TerminalBuyBillUnionPkLoader string *VehicleSupervision/internal/modules/device/model.TerminalBuyBill

// 数据库表名
func (t *TerminalBuyBill) TableName() string {
	return "terminal_buy_bill"
}

// 主键列名
func (t *TerminalBuyBill) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *TerminalBuyBill) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *TerminalBuyBillPkLoader) NewLoader() *TerminalBuyBillPkLoader {
	return &TerminalBuyBillPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalBuyBill, []error) {
			var rs []*TerminalBuyBill
			var m TerminalBuyBill
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t *TerminalBuyBill) UnionPrimaryColumnName() string {
	return "bill_id"
}

// 获取联合主键
func (t *TerminalBuyBill) GetUnionPrimary() string {
	return t.BillID
}

// 新建联合主键dataloader
func (t *TerminalBuyBillUnionPkLoader) NewLoader() *TerminalBuyBillUnionPkLoader {
	return &TerminalBuyBillUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalBuyBill, []error) {
			var rs []*TerminalBuyBill
			var m TerminalBuyBill
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
