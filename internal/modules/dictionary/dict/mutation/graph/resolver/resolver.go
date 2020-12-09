package resolver

import (
	model1 "VehicleSupervision/internal/modules/dictionary/dict/model"
	"VehicleSupervision/internal/modules/dictionary/dict/mutation/graph/model"
	"VehicleSupervision/pkg/xid"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

func (r Resolver) batchInsertParamConvert(objects []*model.DataDictionaryInsertInput) (rs []*model1.DataDictionary) {
	for _, v := range objects {
		rs = append(rs, r.insertParamConvert(v))
	}
	return
}

func (r Resolver) insertParamConvert(v *model.DataDictionaryInsertInput) (rs *model1.DataDictionary) {
	rs = &model1.DataDictionary{
		CreateAt:             *v.CreateAt,
		CreateBy:             v.CreateBy,
		DeleteAt:             v.DeleteAt,
		DeleteBy:             v.DeleteBy,
		DictionaryCategoryID: *v.DictionaryCategoryID,
		DictionaryID:         xid.GetXid(),
		IsDelete:             false,
		Name:                 *v.Name,
		Remarks:              v.Remarks,
		UpdateAt:             v.UpdateAt,
		UpdateBy:             v.UpdateBy,
		Value:                *v.Value,
	}
	return
}
