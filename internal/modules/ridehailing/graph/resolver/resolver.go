package resolver

import (
	"VehicleSupervision/internal/modules/ridehailing/graph/model"
	model1 "VehicleSupervision/internal/modules/ridehailing/model"
	"VehicleSupervision/pkg/xid"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

func rideHailingDriverInsertInputBatchConvert(objects []*model.RideHailingDriverInsertInput) (rs []*model1.RideHailingDriver) {
	for _, v := range objects {
		rs = append(rs, rideHailingDriverInsertInputConvert(v))
	}
	return
}

func rideHailingDriverInsertInputConvert(v *model.RideHailingDriverInsertInput) (rs *model1.RideHailingDriver) {
	rs = &model1.RideHailingDriver{
		Birthday:                  v.Birthday,
		CheckStation:              v.CheckStation,
		CreateAt:                  v.CreateAt,
		CreateBy:                  v.CreateBy,
		CurrentAddress:            v.CurrentAddress,
		DriverName:                v.DriverName,
		DriverSchoolID:            v.DriverSchoolID,
		EndValidDate:              v.EndValidDate,
		FirstTimeReceivedDate:     v.FirstTimeReceivedDate,
		HandleIDPhoto:             v.HandleIDPhoto,
		IDAddress:                 v.IDAddress,
		IDNumber:                  v.IDNumber,
		IdcardPhoto:               v.IdcardPhoto,
		IsDelete:                  false,
		IsFormerDriver:            v.IsFormerDriver,
		Nation:                    v.Nation,
		OperatorID:                v.OperatorID,
		PhoneNumber:               v.PhoneNumber,
		QualificationNumber:       v.QualificationNumber,
		QuasiDrivingModels:        v.QuasiDrivingModels,
		Remarks:                   v.Remarks,
		RideHailingDriverID:       xid.GetXid(),
		RideHailingDriverVerifyID: v.RideHailingDriverVerifyID,
		Sex:                       v.Sex,
		SignGov:                   v.SignGov,
		SignnaturePhoto:           v.SignnaturePhoto,
		StartValidDate:            v.StartValidDate,
		UpdateAt:                  v.UpdateAt,
		UpdateBy:                  v.UpdateBy,
		UpdateTimeIn:              v.UpdateTimeIn,
	}
	return
}

func rideHailingDriverVerifyInsertInputBatchConvert(objects []*model.RideHailingDriverVerifyInsertInput) (rs []*model1.RideHailingDriverVerify) {
	for _, v := range objects {
		rs = append(rs, rideHailingDriverVerifyInsertInputConvert(v))
	}
	return
}

func rideHailingDriverVerifyInsertInputConvert(v *model.RideHailingDriverVerifyInsertInput) (rs *model1.RideHailingDriverVerify) {
	rs = &model1.RideHailingDriverVerify{
		CreateAt:                         v.CreateAt,
		CreateBy:                         v.CreateBy,
		DeleteAt:                         v.DeleteAt,
		DeleteBy:                         v.DeleteBy,
		DrivingExperienceExamineTime:     v.DrivingExperienceExamineTime,
		DrivingExperienceRemark:          v.DrivingExperienceRemark,
		DrivingExamTime:                  v.DrivingExamTime,
		DrugHistoryExamineTime:           v.DrugHistoryExamineTime,
		DrugHistoryRemark:                v.DrugHistoryRemark,
		DrunkDrugDrivingExamineTime:      v.DrunkDrugDrivingExamineTime,
		DrunkDrugDrivingRemark:           v.DrunkDrugDrivingRemark,
		IsDelete:                         false,
		IsDrugHistory:                    v.IsDrugHistory,
		IsDrunkDrugDriving:               v.IsDrunkDrugDriving,
		IsPassDrivingExam:                v.IsPassDrivingExam,
		IsThreeCycleTwelve:               v.IsThreeCycleTwelve,
		IsThreeYearsDrivingExperience:    v.IsThreeYearsDrivingExperience,
		IsTrafficAccidentEscapeRecord:    v.IsTrafficAccidentEscapeRecord,
		IsViolentCrime:                   v.IsViolentCrime,
		RideHailingDriverVerifyID:        xid.GetXid(),
		TrafficAccidentEscapeExamineTime: v.TrafficAccidentEscapeExamineTime,
		TrafficAccidentEscapeRemark:      v.TrafficAccidentEscapeRemark,
		UpdateAt:                         v.UpdateAt,
		UpdateBy:                         v.UpdateBy,
		ViolentCrimeExamineTime:          v.ViolentCrimeExamineTime,
		ViolentCrimeRemark:               v.ViolentCrimeRemark,
	}
	return
}
