package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden TerminalBuyBillLoader string *VehicleSupervision/internal/modules/device/model.TerminalBuyBill

// 数据库表名
func (t TerminalBuyBill) TableName() string {
	return "terminal_buy_bill"
}

// 主键列名
func (t TerminalBuyBill) PrimaryColumnName() string {
	return "id"
}

// 联合主键列名
func (t TerminalBuyBill) UnionPrimaryColumnName() string {
	return "bill_id"
}

// 新建dataloader
func (t *TerminalBuyBill) NewLoader() *TerminalBuyBillLoader {
	return &TerminalBuyBillLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalBuyBill, []error) {
			var rs []*TerminalBuyBill

			db.DB.Model(&TerminalBuyBill{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)

			return rs, nil
		},
	}
}
