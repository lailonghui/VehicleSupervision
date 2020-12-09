package resolver

import (
	model1 "VehicleSupervision/internal/modules/blacklist/record/model"
	"VehicleSupervision/internal/modules/blacklist/record/mutation/graph/model"
	"VehicleSupervision/pkg/xid"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

func (r Resolver) batchInsertParamConvert(objects []*model.BlacklistOperationRecordInsertInput) (rs []*model1.BlacklistOperationRecord) {
	for _, v := range objects {
		rs = append(rs, r.insertParamConvert(v))
	}
	return
}

func (r Resolver) insertParamConvert(v *model.BlacklistOperationRecordInsertInput) (rs *model1.BlacklistOperationRecord) {
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
