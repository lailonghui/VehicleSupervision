package resolver

import (
	"VehicleSupervision/internal/modules/blacklist/graph/model"
	model1 "VehicleSupervision/internal/modules/blacklist/model"
	"VehicleSupervision/pkg/xid"
)

//go:generate go run github.com/99designs/gqlgen
// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

func (r Resolver) blacklistOperationRecordInsertInputBatchConvert(objects []*model.BlacklistOperationRecordInsertInput) (rs []*model1.BlacklistOperationRecord) {
	for _, v := range objects {
		rs = append(rs, r.blacklistOperationRecordInsertInputConvert(v))
	}
	return
}

func (r Resolver) blacklistOperationRecordInsertInputConvert(v *model.BlacklistOperationRecordInsertInput) (rs *model1.BlacklistOperationRecord) {
	rs = &model1.BlacklistOperationRecord{
		BlacklistRecordID: xid.GetXid(),
		BlacklistType:     *v.BlacklistType,
		CreateAt:          *v.CreateAt,
		CreateBy:          v.CreateBy,
		DeleteAt:          v.DeleteAt,
		DeleteBy:          v.DeleteBy,
		IsDelete:          false,
		Operate:           *v.Operate,
		Remarks:           v.Remarks,
		UpdateAt:          v.UpdateAt,
		UpdateBy:          v.UpdateBy,
		VSeqn:             *v.VSeqn,
	}
	return
}
