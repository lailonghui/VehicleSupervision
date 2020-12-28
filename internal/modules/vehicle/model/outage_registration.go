// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// 停运登记表
type OutageRegistration struct {
	// 按指定方法生成                                  ( 主键                                                         )
	ID int64 `json:"id"`
	// 外部编码，由golang程序生成的xid，暴露到外部使用 ( 联合主键                                                     )
	OutageRegistrationID string `json:"outage_registration_id"`
	// vehicle_info 车辆信息表 的vehicle_id
	VehicleID string `json:"vehicle_id"`
	// 用户                                            ( system_user表的user_id                                   )
	UserID *string `json:"user_id"`
	// 停运起始时间
	OutageStartTime *time.Time `json:"outage_start_time"`
	// 停运截止时间
	OutageEndTime *time.Time `json:"outage_end_time"`
	// 审核状态
	ReviewStatus *string `json:"review_status"`
	// 审核人                                          ( system_user表的user_id                                   )
	Reviewer *string `json:"reviewer"`
	// 审核时间
	ReviewTime *time.Time `json:"review_time"`
	// 停运报备上传文件表ID                            ( outage_filing_upload_file 的outage_filing_upload_file_id )
	OutageFilingUploadFileID *string `json:"outage_filing_upload_file_id"`
	// 停运起始经纬度
	OutageStartCoordinate *string `json:"outage_start_coordinate"`
	// 停运结束经纬度
	OutageEndCoordinate *string `json:"outage_end_coordinate"`
	// 上线时间
	OnlineTime *time.Time `json:"online_time"`
	// 停运起始位置
	OutageStartPosition *string `json:"outage_start_position"`
	// 是否管理部门审核
	IsManagementReview *bool `json:"is_management_review"`
	// 是否失效
	IsInvalid *bool `json:"is_invalid"`
	// 是否最新
	IsLatest *bool `json:"is_latest"`
	// 是否删除
	IsDeleted *bool `json:"is_deleted"`
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 创建人                                          ( system_user表的user_id                                   )
	CreatedBy string `json:"created_by"`
	// 修改时间
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人                                          ( system_user表的user_id                                   )
	UpdatedBy *string `json:"updated_by"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人                                          ( system_user表的user_id                                   )
	DeletedBy *string `json:"deleted_by"`
}
