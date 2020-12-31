// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// 驾驶员指纹表
type DriverFingerprint struct {
	// 按指定方法生成                                  ( 主键                       )
	ID int64 `json:"id"`
	// 外部编码，由golang程序生成的xid，暴露到外部使用 ( 联合主键                   )
	DriverFingerprintID string `json:"driver_fingerprint_id"`
	// driver_info驾驶员信息表的driver_id
	DriverID string `json:"driver_id"`
	// 驾驶员姓名
	DriverName *string `json:"driver_name"`
	// 用户ID                                          ( system_user表的user_id )
	UserID *string `json:"user_id"`
	// 特征码
	Signature *string `json:"signature"`
	// 指纹名称
	FingerprintName *string `json:"fingerprint_name"`
	// SIM卡号
	SimNumber *string `json:"sim_number"`
	// 指令ID
	InstructionID *string `json:"instruction_id"`
	// 操作类型
	OperationType *int `json:"operation_type"`
	// 内容
	Content *string `json:"content"`
	// 操作时间
	OperationTime *time.Time `json:"operation_time"`
	// 上传时间
	UploadTime *time.Time `json:"upload_time"`
	// 时间戳
	Timestamp *string `json:"timestamp"`
	// 终端ID
	TernimalID *string `json:"ternimal_id"`
	// 是否成功
	IsSuccess *bool `json:"is_success"`
	// 是否删除
	IsDeleted *bool `json:"is_deleted"`
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 创建人                                          ( system_user表的user_id )
	CreatedBy string `json:"created_by"`
	// 修改时间
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人                                          ( system_user表的user_id )
	UpdatedBy *string `json:"updated_by"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人                                          ( system_user表的user_id )
	DeletedBy *string `json:"deleted_by"`
}