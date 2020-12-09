package resolver

import (
	model1 "VehicleSupervision/internal/modules/sim/card/model"
	"VehicleSupervision/internal/modules/sim/card/mutation/graph/model"
	"VehicleSupervision/pkg/xid"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

func (r Resolver) batchInsertParamConvert(objects []*model.SimCardInsertInput) (rs []*model1.SimCard) {
	for _, v := range objects {
		rs = append(rs, r.insertParamConvert(v))
	}
	return
}

func (r Resolver) insertParamConvert(v *model.SimCardInsertInput) (rs *model1.SimCard) {
	rs = &model1.SimCard{
		CreateAt:    *v.CreateAt,
		CreateBy:    v.CreateBy,
		DeleteAt:    v.DeleteAt,
		DeleteBy:    v.DeleteBy,
		DeptID:      v.DeptID,
		IsDelete:    false,
		MobileType:  v.MobileType,
		OperatorsID: v.OperatorsID,
		ProxyrgID:   v.ProxyrgID,
		Remark:      v.Remark,
		SimCardID:   xid.GetXid(),
		SimNumber:   v.SimNumber,
		TerminalID:  v.TerminalID,
		UpdateAt:    v.UpdateAt,
		UpdateBy:    v.UpdateBy,
	}
	return
}
