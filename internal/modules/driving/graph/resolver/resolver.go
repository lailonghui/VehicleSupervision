package resolver

import (
	"VehicleSupervision/internal/modules/driving/graph/model"
	model1 "VehicleSupervision/internal/modules/driving/model"
	"VehicleSupervision/pkg/xid"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

func (r Resolver) drivingLogInsertInputBatchConvert(objects []*model.DrivingLogInsertInput) (rs []*model1.DrivingLog) {
	for _, v := range objects {
		rs = append(rs, r.drivingLogInsertInputConvert(v))
	}
	return
}

func (r Resolver) drivingLogInsertInputConvert(v *model.DrivingLogInsertInput) (rs *model1.DrivingLog) {
	rs = &model1.DrivingLog{
		Cause:                  v.Cause,
		CheckOrganizationLevel: *v.CheckOrganizationLevel,
		CheckState:             *v.CheckState,
		CreateAt:               *v.CreateAt,
		CreateBy:               v.CreateBy,
		DeleteAt:               v.DeleteAt,
		DeleteBy:               v.DeleteBy,
		DriverID:               *v.DriverID,
		DrivingEndTime:         *v.DrivingEndTime,
		DrivingLogID:           xid.GetXid(),
		DrivingStartTime:       *v.DrivingStartTime,
		EndTime:                *v.EndTime,
		IsDelete:               false,
		IsFill:                 *v.IsFill,
		RegisterAt:             v.RegisterAt,
		RegisterBy:             v.RegisterBy,
		Remarks:                v.Remarks,
		Route:                  v.Route,
		StartTime:              *v.StartTime,
		UpdateAt:               v.UpdateAt,
		UpdateBy:               v.UpdateBy,
		VehicleID:              *v.VehicleID,
	}
	return
}
