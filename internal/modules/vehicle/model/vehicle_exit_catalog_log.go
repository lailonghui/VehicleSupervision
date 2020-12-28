// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// 车辆退出目录库日志表
type VehicleExitCatalogLog struct {
	// 按指定方法生成                                               ( 主键                                                         )
	ID int64 `json:"id"`
	// 外部编码，由golang程序生成的xid，暴露到外部使用              ( 联合主键                                                     )
	VehicleExitCatalogLogID string `json:"vehicle_exit_catalog_log_id"`
	// 车辆退出目录库审核表id                                       ( vehicle_exit_catalog_review 车辆退出目录库审核表的vehicle_exit_catalog_review_id )
	VehicleExitCatalogReviewID *string `json:"vehicle_exit_catalog_review_id"`
	// 类别  3.企业车辆退出目录库  6.零散车辆退出目录库
	ExitType *int `json:"exit_type"`
	// 退出步骤  Type=3:{1.AA市A县a企业提交材料,2.A县交警同意,3.A县渣土办同意}  Type=6:{1.A县交警提交材料,2.A县渣土办同意}
	ExitStep *int `json:"exit_step"`
	// 操作人
	Operator *string `json:"operator"`
	// 审核状态  0.申请  1.审批  2.退回  3.撤销
	ReviewStatus *int `json:"review_status"`
	// 审批用户组  1.运输企业  2.管理部门
	ReviewUserGroup *int `json:"review_user_group"`
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 创建人                                                       ( system_user表的user_id                                   )
	CreatedBy string `json:"created_by"`
	// 修改时间
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人                                                       ( system_user表的user_id                                   )
	UpdatedBy *string `json:"updated_by"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人                                                       ( system_user表的user_id                                   )
	DeletedBy *string `json:"deleted_by"`
}
