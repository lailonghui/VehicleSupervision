package resolver

import (
	model1 "VehicleSupervision/internal/modules/sim/flow/model"
	"VehicleSupervision/internal/modules/sim/flow/mutation/graph/model"
	"VehicleSupervision/pkg/xid"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

func (r Resolver) batchInsertParamConvert(objects []*model.SimCardFlowInsertInput) (rs []*model1.SimCardFlow) {
	for _, v := range objects {
		rs = append(rs, r.insertParamConvert(v))
	}
	return
}

func (r Resolver) insertParamConvert(v *model.SimCardFlowInsertInput) (rs *model1.SimCardFlow) {
	rs = &model1.SimCardFlow{
		CardAvgFlow:    v.CardAvgFlow,
		CardNoRemark:   v.CardNoRemark,
		CreateAt:       *v.CreateAt,
		CreateBy:       v.CreateBy,
		DeleteAt:       v.DeleteAt,
		DeleteBy:       v.DeleteBy,
		EnterpriseID:   v.EnterpriseID,
		Iccid:          v.Iccid,
		IotCardNo:      *v.IotCardNo,
		IsDelete:       false,
		IsSharePool:    *v.IsSharePool,
		PoolAvgFlow:    v.PoolAvgFlow,
		SimCardFlowID:  xid.GetXid(),
		SuitFlow:       v.SuitFlow,
		SuitLeftFlow:   v.SuitLeftFlow,
		SuitOverFlow:   v.SuitOverFlow,
		SuitSmsLeftNum: v.SuitSmsLeftNum,
		SuitSmsNum:     v.SuitSmsNum,
		SuitSmsOverNum: v.SuitSmsOverNum,
		SuitUseSmsNum:  v.SuitUseSmsNum,
		UpdateAt:       v.UpdateAt,
		UpdateBy:       v.UpdateBy,
		UseFlow:        v.UseFlow,
	}
	return
}
