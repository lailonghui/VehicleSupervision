package resolver

import (
	model1 "VehicleSupervision/internal/modules/admin/systemuser/model"
	"VehicleSupervision/internal/modules/admin/systemuser/mutation/graph/model"
	"VehicleSupervision/pkg/xid"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

func (r Resolver) batchInsertParamConvert(objects []*model.SystemUserInsertInput) (rs []*model1.SystemUser) {
	for _, v := range objects {
		rs = append(rs, r.insertParamConvert(v))
	}
	return
}

func (r Resolver) insertParamConvert(v *model.SystemUserInsertInput) (rs *model1.SystemUser) {
	rs = &model1.SystemUser{
		IsBindIP:     v.IsBindIP,
		AppVersion:   v.AppVersion,
		AuditLevel:   v.AuditLevel,
		CreateBy:     v.CreateBy,
		CreatedAt:    v.CreatedAt,
		DeleteAt:     v.DeleteAt,
		DeleteBy:     v.DeleteBy,
		DepartmentID: v.DepartmentID,
		Email:        v.Email,
		EnterpriseID: v.EnterpriseID,
		Grade:        v.Grade,
		IPAddress:    v.IPAddress,
		IsDelete:     false,
		IsValid:      v.IsValid,
		UserID:       xid.GetXid(),
		Mkey:         v.Mkey,
		Mobile:       v.Mobile,
		Password:     *v.Password,
		ProxyUser:    v.ProxyUser,
		Remarks:      v.Remarks,
		Telephone:    v.Telephone,
		Username:     *v.Username,
		Ukey:         v.Ukey,
		UpdateAt:     v.UpdateAt,
		UpdateBy:     v.UpdateBy,
		UserState:    v.UserState,
		UserType:     v.UserType,
	}
	return
}
