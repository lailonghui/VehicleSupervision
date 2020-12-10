package resolver

import (
	"VehicleSupervision/internal/modules/dictionary/graph/model"
	model1 "VehicleSupervision/internal/modules/dictionary/model"
	"VehicleSupervision/pkg/xid"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

func (r Resolver) dataDictionaryCategoryInsertInputBatchConvert(objects []*model.DataDictionaryCategoryInsertInput) (rs []*model1.DataDictionaryCategory) {
	for _, v := range objects {
		rs = append(rs, r.dataDictionaryCategoryInsertInputConvert(v))
	}
	return
}

func (r Resolver) dataDictionaryCategoryInsertInputConvert(v *model.DataDictionaryCategoryInsertInput) (rs *model1.DataDictionaryCategory) {
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

func (r Resolver) dataDictionaryInsertInputBatchConvert(objects []*model.DataDictionaryInsertInput) (rs []*model1.DataDictionary) {
	for _, v := range objects {
		rs = append(rs, r.dataDictionaryInsertInputConvert(v))
	}
	return
}

func (r Resolver) dataDictionaryInsertInputConvert(v *model.DataDictionaryInsertInput) (rs *model1.DataDictionary) {
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
