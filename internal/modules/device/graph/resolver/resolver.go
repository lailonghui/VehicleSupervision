package resolver

import (
	"VehicleSupervision/internal/modules/device/graph/model"
	model1 "VehicleSupervision/internal/modules/device/model"
	"VehicleSupervision/pkg/xid"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

func simCardInsertInputBatchConvert(objects []*model.SimCardInsertInput) (rs []*model1.SimCard) {
	for _, v := range objects {
		rs = append(rs, simCardInsertInputConvert(v))
	}
	return
}

func simCardInsertInputConvert(v *model.SimCardInsertInput) (rs *model1.SimCard) {
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

func simCardFlowInsertInputBatchConvert(objects []*model.SimCardFlowInsertInput) (rs []*model1.SimCardFlow) {
	for _, v := range objects {
		rs = append(rs, simCardFlowInsertInputConvert(v))
	}
	return
}

func simCardFlowInsertInputConvert(v *model.SimCardFlowInsertInput) (rs *model1.SimCardFlow) {
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
