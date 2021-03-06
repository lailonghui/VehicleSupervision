// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// 网约车驾驶员审核表
type RideHailingDriverVerify struct {
	// ID
	ID int64 `json:"id"`
	// 审核ID
	RideHailingDriverVerifyID string `json:"ride_hailing_driver_verify_id"`
	// 是否有暴力犯罪记录（刑侦支队）
	IsViolentCrime *bool `json:"is_violent_crime"`
	// 刑侦支队审核备注
	ViolentCrimeRemark *string `json:"violent_crime_remark"`
	// 刑侦支队审核时间
	ViolentCrimeExamineTime *time.Time `json:"violent_crime_examine_time"`
	// 是否有吸毒记录（禁毒支队）
	IsDrugHistory *bool `json:"is_drug_history"`
	// 禁毒支队审核备注
	DrugHistoryRemark *string `json:"drug_history_remark"`
	// 禁毒支队审核时间
	DrugHistoryExamineTime *time.Time `json:"drug_history_examine_time"`
	// 是否满足三年驾龄（驾管大队）
	IsThreeYearsDrivingExperience *bool `json:"is_three_years_driving_experience"`
	// 驾管大队审核备注
	DrivingExperienceRemark *string `json:"driving_experience_remark"`
	// 驾管大队审核时间
	DrivingExperienceExamineTime *time.Time `json:"driving_experience_examine_time"`
	// 是否最近连续三个记分周期内没有记满12分（驾管大队）
	IsThreeCycleTwelve *bool `json:"is_three_cycle_twelve"`
	// 是否有交通肇事犯罪记录（事故大队）
	IsTrafficAccidentEscapeRecord *bool `json:"is_traffic_accident_escape_record"`
	// 事故大队审核备注
	TrafficAccidentEscapeRemark *string `json:"traffic_accident_escape_remark"`
	// 事故大队审核时间
	TrafficAccidentEscapeExamineTime *time.Time `json:"traffic_accident_escape_examine_time"`
	// 是否有酒驾毒驾记录(秩序大队)
	IsDrunkDrugDriving *bool `json:"is_drunk_drug_driving"`
	// 秩序大队审核备注
	DrunkDrugDrivingRemark *string `json:"drunk_drug_driving_remark"`
	// 秩序大队审核时间
	DrunkDrugDrivingExamineTime *time.Time `json:"drunk_drug_driving_examine_time"`
	// 是否通过考试(交通局审)
	IsPassDrivingExam *bool `json:"is_pass_driving_exam"`
	// 交通局考试认定时间
	DrivingExamTime *time.Time `json:"driving_exam_time"`
	// 创建时间
	CreatedAt *time.Time `json:"created_at"`
	// 创建人
	CreatedBy *string `json:"created_by"`
	// 修改时间
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人
	UpdateBy *string `json:"update_by"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人
	DeletedBy *string `json:"deleted_by"`
	// 是否删除
	IsDeleted string `json:"is_deleted"`
}
