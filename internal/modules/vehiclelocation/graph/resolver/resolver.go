package resolver

import (
	"VehicleSupervision/internal/modules/vehiclelocation/graph/model"
	model1 "VehicleSupervision/internal/modules/vehiclelocation/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

func (r Resolver) vehicleLocationHisInsertInputBatchConvert(objects []*model.VehicleLocationHisInsertInput) (rs []*model1.VehicleLocationHis) {
	for _, v := range objects {
		rs = append(rs, r.vehicleLocationHisInsertInputConvert(v))
	}
	return
}

func (r Resolver) vehicleLocationHisInsertInputConvert(v *model.VehicleLocationHisInsertInput) (rs *model1.VehicleLocationHis) {
	rs = &model1.VehicleLocationHis{
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

func (r Resolver) vehicleLocationLastInsertInputBatchConvert(objects []*model.VehicleLocationLastInsertInput) (rs []*model1.VehicleLocationLast) {
	for _, v := range objects {
		rs = append(rs, r.vehicleLocationLastInsertInputConvert(v))
	}
	return
}

func (r Resolver) vehicleLocationLastInsertInputConvert(v *model.VehicleLocationLastInsertInput) (rs *model1.VehicleLocationLast) {
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
