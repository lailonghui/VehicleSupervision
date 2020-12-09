package resolver

import (
	model1 "VehicleSupervision/internal/modules/vehiclelocation/last/model"
	"VehicleSupervision/internal/modules/vehiclelocation/last/mutation/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

func (r Resolver) batchInsertParamConvert(objects []*model.VehicleLocationLastInsertInput) (rs []*model1.VehicleLocationLast) {
	for _, v := range objects {
		rs = append(rs, r.insertParamConvert(v))
	}
	return
}

func (r Resolver) insertParamConvert(v *model.VehicleLocationLastInsertInput) (rs *model1.VehicleLocationLast) {
	rs = &model1.VehicleLocationLast{
		Acceleration:        v.Acceleration,
		AlarmContent:        v.AlarmContent,
		Alititude:           v.Alititude,
		Coordinate:          v.Coordinate,
		CorrectCoordinate:   v.CorrectCoordinate,
		Direction:           v.Direction,
		DistrictID:          v.DistrictID,
		DriverID:            v.DriverID,
		EnterpriseID:        v.EnterpriseID,
		GpsSpeed:            v.GpsSpeed,
		Imei:                v.Imei,
		IsLocate:            v.IsLocate,
		LocateTime:          v.LocateTime,
		LocationDescription: v.LocationDescription,
		Mileage:             v.Mileage,
		RoadName:            v.RoadName,
		SimNumber:           v.SimNumber,
		SpeedLimitThreshold: v.SpeedLimitThreshold,
		StarCount:           v.StarCount,
		StarStatus:          v.StarStatus,
		SupervisionPhotoID:  v.SupervisionPhotoID,
		TachographSpeed:     v.TachographSpeed,
		VehicleID:           *v.VehicleID,
		VehicleStatus:       v.VehicleStatus,
	}
	return
}
