package resolver

import (
	model1 "VehicleSupervision/internal/modules/admin/department/model"
	"VehicleSupervision/internal/modules/admin/department/mutation/graph/model"
	"VehicleSupervision/pkg/xid"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

func (r Resolver) batchInsertParamConvert(objects []*model.DepartmentInsertInput) (rs []*model1.Department) {
	for _, v := range objects {
		rs = append(rs, r.insertParamConvert(v))
	}
	return
}

func (r Resolver) insertParamConvert(v *model.DepartmentInsertInput) (rs *model1.Department) {
	rs = &model1.Department{
		CreateAt:             v.CreateAt,
		CreateBy:             v.CreateBy,
		DeleteAt:             v.DeleteAt,
		DeleteBy:             v.DeleteBy,
		DepartmentCategory:   v.DepartmentCategory,
		DepartmentCode:       v.DepartmentCode,
		DepartmentID:         xid.GetXid(),
		DepartmentName:       v.DepartmentName,
		EnterpriseID:         *v.EnterpriseID,
		InternalNumber:       v.InternalNumber,
		IsDelete:             false,
		Remarks:              v.Remarks,
		SuperiorDepartmentID: v.SuperiorDepartmentID,
		UpdateAt:             v.UpdateAt,
		UpdateBy:             v.UpdateBy,
	}
	return
}
