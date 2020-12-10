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

func terminalInsertInputBatchConvert(objects []*model.TerminalInsertInput) (rs []*model1.Terminal) {
	for _, v := range objects {
		rs = append(rs, terminalInsertInputConvert(v))
	}
	return
}

func terminalInsertInputConvert(v *model.TerminalInsertInput) (rs *model1.Terminal) {
	rs = &model1.Terminal{
		AdasModel:        v.AdasModel,
		AuthKey:          v.AuthKey,
		CameraNum:        v.CameraNum,
		ChannelNo:        v.ChannelNo,
		CityID:           v.CityID,
		CreateAt:         *v.CreateAt,
		CreateBy:         v.CreateBy,
		DeleteAt:         v.DeleteAt,
		DeleteBy:         v.DeleteBy,
		DeptID:           v.DeptID,
		FirstInstallTime: v.FirstInstallTime,
		GuaranteeDate:    v.GuaranteeDate,
		Imei:             *v.Imei,
		InstallManID:     v.InstallManID,
		IsDelete:         false,
		IsReg:            v.IsReg,
		IsSupportPhoto:   *v.IsSupportPhoto,
		MockAuthKey:      v.MockAuthKey,
		ProduceDate:      v.ProduceDate,
		ProveProxyrgID:   v.ProveProxyrgID,
		ProvinceID:       v.ProvinceID,
		ProxyrgID:        v.ProxyrgID,
		RecordDate:       v.RecordDate,
		RegID:            v.RegID,
		Remarks:          v.Remarks,
		RemoveReason:     v.RemoveReason,
		SimID:            v.SimID,
		SprgID:           v.SprgID,
		TerminalID:       xid.GetXid(),
		TypeID:           v.TypeID,
		UpdateAt:         v.UpdateAt,
		UpdateBy:         v.UpdateBy,
		VehicleID:        v.VehicleID,
		VersionNumber:    v.VersionNumber,
	}
	return
}

func terminalFactoryInsertInputBatchConvert(objects []*model.TerminalFactoryInsertInput) (rs []*model1.TerminalFactory) {
	for _, v := range objects {
		rs = append(rs, terminalFactoryInsertInputConvert(v))
	}
	return
}

func terminalFactoryInsertInputConvert(v *model.TerminalFactoryInsertInput) (rs *model1.TerminalFactory) {
	rs = &model1.TerminalFactory{
		Address:          v.Address,
		Contact:          v.Contact,
		ContactPhone:     v.ContactPhone,
		CreateAt:         v.CreateAt,
		CreateBy:         v.CreateBy,
		DeleteAt:         v.DeleteAt,
		DeleteBy:         v.DeleteBy,
		FactoryID:        xid.GetXid(),
		FactoryName:      *v.FactoryName,
		IsDelete:         false,
		Remark:           v.Remark,
		TechContact:      v.TechContact,
		TechContactPhone: v.TechContactPhone,
		UpdateAt:         v.UpdateAt,
		UpdateBy:         v.UpdateBy,
	}
	return
}

func terminalModalInsertInputBatchConvert(objects []*model.TerminalModalInsertInput) (rs []*model1.TerminalModal) {
	for _, v := range objects {
		rs = append(rs, terminalModalInsertInputConvert(v))
	}
	return
}

func terminalModalInsertInputConvert(v *model.TerminalModalInsertInput) (rs *model1.TerminalModal) {
	rs = &model1.TerminalModal{
		AdasModal:             v.AdasModal,
		CreateAt:              *v.CreateAt,
		CreateBy:              v.CreateBy,
		DeleteAt:              v.DeleteAt,
		DeleteBy:              v.DeleteBy,
		FactoryID:             v.FactoryID,
		IsDelete:              false,
		IsElectronicsPostCard: v.IsElectronicsPostCard,
		IsSlagCarTeminal:      v.IsSlagCarTeminal,
		IsTestingSituation:    v.IsTestingSituation,
		IsTransportDept4g:     v.IsTransportDept4g,
		ModalName:             v.ModalName,
		ProxyrgID:             v.ProxyrgID,
		RecordNo:              v.RecordNo,
		Remark:                v.Remark,
		TerminalModalID:       xid.GetXid(),
		TerminalTypeID:        v.TerminalTypeID,
		UpdateAt:              v.UpdateAt,
		UpdateBy:              v.UpdateBy,
	}
	return
}

func terminalTypeInsertInputBatchConvert(objects []*model.TerminalTypesInsertInput) (rs []*model1.TerminalType) {
	for _, v := range objects {
		rs = append(rs, terminalTypeInsertInputConvert(v))
	}
	return
}

func terminalTypeInsertInputConvert(v *model.TerminalTypesInsertInput) (rs *model1.TerminalType) {
	rs = &model1.TerminalType{
		CreateAt:     *v.CreateAt,
		CreateBy:     v.CreateBy,
		DeleteAt:     v.DeleteAt,
		DeleteBy:     v.DeleteBy,
		IsDelete:     false,
		ProtocolName: *v.ProtocolName,
		Remark:       v.Remark,
		TypeID:       xid.GetXid(),
		UpdateAt:     v.UpdateAt,
		UpdateBy:     v.UpdateBy,
	}
	return
}
