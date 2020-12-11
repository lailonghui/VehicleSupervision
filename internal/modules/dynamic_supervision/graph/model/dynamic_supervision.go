/*
@Time : 2020/12/11 14:35
@Author : lai
@Description :
@File : dynamic_supervision
*/
package model

import "time"

// 动态监管抽查主表
//
//
// columns and relationships of "dynamic_supervision"
type DynamicSupervision struct {
	// 抽查人员
	CheckUserID *string `json:"check_user_id"`
	// 抽查人员位置的城市ID
	CityID *string `json:"city_id"`
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 创建人
	CreatedBy string `json:"created_by"`
	// 日
	Day *int `json:"day"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人
	DeletedBy *string `json:"deleted_by"`
	// 抽查人员位置的区域ID
	DistrictID *string `json:"district_id"`
	// 主键
	ID int64 `json:"id"`
	// 是否被删除
	IsDelete *bool `json:"is_delete"`
	// 月
	Month *int `json:"month"`
	// 抽查人员位置的省份ID
	ProvinceID *string `json:"province_id"`
	// 抽查日期
	SpotCheckDate *time.Time `json:"spot_check_date"`
	// 抽查数量
	SpotCheckNumber *int `json:"spot_check_number"`
	// 抽查比例
	SpotCheckRatio *float64 `json:"spot_check_ratio"`
	// 抽查总数
	SpotCheckTotalNumber *int `json:"spot_check_total_number"`
	// 联合主键
	SupervisionID string `json:"supervision_id"`
	// 总车辆数
	TotalNumberVehicle *int `json:"total_number_vehicle"`
	// 修改时间
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人
	UpdatedBy *string `json:"updated_by"`
	// 年
	Year *int `json:"year"`
}

func (DynamicSupervision) TableName() string {
	return "dynamic_supervision"
}
