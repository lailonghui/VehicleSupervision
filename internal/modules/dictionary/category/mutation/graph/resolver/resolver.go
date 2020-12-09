package resolver

import (
	model1 "VehicleSupervision/internal/modules/dictionary/category/model"
	"VehicleSupervision/internal/modules/dictionary/category/mutation/graph/model"
	"VehicleSupervision/pkg/xid"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

func (r Resolver) batchInsertParamConvert(objects []*model.DataDictionaryCategoryInsertInput) (rs []*model1.DataDictionaryCategory) {
	for _, v := range objects {
		rs = append(rs, r.insertParamConvert(v))
	}
	return
}

func (r Resolver) insertParamConvert(v *model.DataDictionaryCategoryInsertInput) (rs *model1.DataDictionaryCategory) {
	rs = &model1.DataDictionaryCategory{
		CategoryCode:         *v.CategoryCode,
		CategoryName:         *v.CategoryName,
		CreateAt:             *v.CreateAt,
		CreateBy:             v.CreateBy,
		DeleteAt:             v.DeleteAt,
		DeleteBy:             v.DeleteBy,
		DictionaryCategoryID: xid.GetXid(),
		IsDelete:             false,
		Remarks:              v.Remarks,
		UpdateAt:             v.UpdateAt,
		UpdateBy:             v.UpdateBy,
	}
	return
}
