// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// 车辆退出目录库审核表
type VehicleExitCatalogReview struct {
	// 按指定方法生成                                               ( 主键                                 )
	ID int64 `json:"id"`
	// 外部编码，由golang程序生成的xid，暴露到外部使用              ( 联合主键                             )
	VehicleExitCatalogReviewID string `json:"vehicle_exit_catalog_review_id"`
	// vehicle_info 车辆信息表 的vehicle_id
	VehicleID string `json:"vehicle_id"`
	// 所在企业id                                                   ( enterprise_info表的enterprise_id )
	EnterpriseID *string `json:"enterprise_id"`
	// 备注
	Remarks *string `json:"remarks"`
	// 审核状态  0.未完成 1.完成
	ReviewStatus *int `json:"review_status"`
	// 类别  3.企业车辆退出目录库  6.零散车辆退出目录库
	ExitType *int `json:"exit_type"`
	// 退出步骤  Type=3:{1.AA市A县a企业提交材料,2.A县交警同意,3.A县渣土办同意}  Type=6:{1.A县交警提交材料,2.A县渣土办同意}
	ExitStep *int `json:"exit_step"`
	// 地区ID
	AreaID *string `json:"area_id"`
	// 原单位ID                                                     ( enterprise_info表的enterprise_id )
	OriginalEnterpriseID *string `json:"original_enterprise_id"`
	// 原单位名称
	OriginalEnterpriseName *string `json:"original_enterprise_name"`
	// 原自编号
	OriginalSelfNumber *string `json:"original_self_number"`
	// 是否删除
	IsDeleted *bool `json:"is_deleted"`
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 创建人                                                       ( system_user表的user_id           )
	CreatedBy string `json:"created_by"`
	// 修改时间
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人                                                       ( system_user表的user_id           )
	UpdatedBy *string `json:"updated_by"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人                                                       ( system_user表的user_id           )
	DeletedBy *string `json:"deleted_by"`
}