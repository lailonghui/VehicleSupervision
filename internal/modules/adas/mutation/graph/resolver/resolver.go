//go:generate go run github.com/99designs/gqlgen
package resolver

import (
	"VehicleSupervision/internal/modules/adas/mutation/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

func (r Resolver) batchAdasDataInsertParamConvert(objects []*model.AdasDataInsertInput) (rs []*model.AdasData) {
	for _, v := range objects {
		rs = append(rs, r.insertAdasDataParamConvert(v))
	}
	return
}

func (r Resolver) insertAdasDataParamConvert(v *model.AdasDataInsertInput) (rs *model.AdasData) {
	rs = &model.AdasData{
		AdasAttachment:   v.AdasAttachment,
		AlarmID:          v.AlarmID,
		AlarmNo:          v.AlarmNo,
		AlarmSource:      v.AlarmSource,
		AlarmTime:        v.AlarmTime,
		AlarmType:        v.AlarmType,
		Altitude:         v.Altitude,
		CreatedDate:      v.CreatedDate,
		FatigueDegree:    v.FatigueDegree,
		FrontCarDistance: v.FrontCarDistance,
		ID:               *v.ID,
		Latitude:         v.Latitude,
		Levels:           v.Levels,
		Location:         v.Location,
		Longitude:        v.Longitude,
		PlateNo:          v.PlateNo,
		SimID:            v.SimID,
		Speed:            v.Speed,
		TerminalID:       v.TerminalID,
		VehicleID:        v.VehicleID,
		VehicleStatus:    v.VehicleStatus,
	}
	return
}

func (r Resolver) batchAdasAttachmentInsertParamConvert(objects []*model.AdasAttachmentInsertInput) (rs []*model.AdasAttachment) {
	for _, v := range objects {
		rs = append(rs, r.insertAdasAttachmentParamConvert(v))
	}
	return
}

func (r Resolver) insertAdasAttachmentParamConvert(v *model.AdasAttachmentInsertInput) (rs *model.AdasAttachment) {
	rs = &model.AdasAttachment{
		AlarmNo:      v.AlarmNo,
		AttachmentID: *v.AttachmentID,
		ChannelID:    v.ChannelID,
		CreateDate:   v.CreateDate,
		FileLength:   v.FileLength,
		FileName:     v.FileName,
		FilePath:     v.FilePath,
		FileType:     v.FileType,
		ID:           *v.ID,
		PlateNo:      v.PlateNo,
		SimNo:        v.SimNo,
		Status:       v.Status,
		VehicleID:    v.VehicleID,
	}
	return
}
