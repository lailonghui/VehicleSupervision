//go:generate go run github.com/99designs/gqlgen
package resolver

import (
	"VehicleSupervision/internal/modules/area/mutation/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

func (r Resolver) batchCityInsertParamConvert(objects []*model.CityInsertInput) (rs []*model.City) {
	for _, v := range objects {
		rs = append(rs, r.insertCityParamConvert(v))
	}
	return
}

func (r Resolver) insertCityParamConvert(v *model.CityInsertInput) (rs *model.City) {
	rs = &model.City{
		CityID:     v.CityID,
		Code:       v.Code,
		CreateAt:   v.CreateAt,
		CreateBy:   v.CreateBy,
		DeleteAt:   v.DeleteAt,
		DeleteBy:   v.CreateBy,
		ID:         *v.ID,
		IsDelete:   v.IsDelete,
		Name:       v.Name,
		ProvinceID: v.ProvinceID,
		Remarks:    v.Remarks,
		UpdateAt:   v.UpdateAt,
		UpdateBy:   v.UpdateBy,
	}
	return
}

func (r Resolver) batchDistrictInsertParamConvert(objects []*model.DistrictInsertInput) (rs []*model.District) {
	for _, v := range objects {
		rs = append(rs, r.insertDistrictParamConvert(v))
	}
	return
}

func (r Resolver) insertDistrictParamConvert(v *model.DistrictInsertInput) (rs *model.District) {
	rs = &model.District{
		CityID:     v.CityID,
		Code:       v.Code,
		CreateAt:   v.CreateAt,
		CreateBy:   v.CreateBy,
		DeleteAt:   v.DeleteAt,
		DeleteBy:   v.CreateBy,
		ID:         *v.ID,
		IsDelete:   v.IsDelete,
		Name:       v.Name,
		DistrictID: v.DistrictID,
		Remarks:    v.Remarks,
		UpdateAt:   v.UpdateAt,
		UpdateBy:   v.UpdateBy,
	}
	return
}

func (r Resolver) batchProvinceInsertParamConvert(objects []*model.ProvinceInsertInput) (rs []*model.Province) {
	for _, v := range objects {
		rs = append(rs, r.insertProvinceParamConvert(v))
	}
	return
}

func (r Resolver) insertProvinceParamConvert(v *model.ProvinceInsertInput) (rs *model.Province) {
	rs = &model.Province{
		Code:       v.Code,
		CreateAt:   v.CreateAt,
		CreateBy:   v.CreateBy,
		DeleteAt:   v.DeleteAt,
		DeleteBy:   v.CreateBy,
		ID:         *v.ID,
		IsDelete:   v.IsDelete,
		Name:       v.Name,
		ProvinceID: v.ProvinceID,
		Remarks:    v.Remarks,
		UpdateAt:   v.UpdateAt,
		UpdateBy:   v.UpdateBy,
	}
	return
}
