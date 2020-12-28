// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// 驾驶员指纹关联
type DriverFingerprintAssociation struct {
	// 按指定方法生成                                             ( 主键                       )
	ID int64 `json:"id"`
	// 外部编码，由golang程序生成的xid，暴露到外部使用            ( 联合主键                   )
	DriverFingerprintAssociationID string `json:"driver_fingerprint_association_id"`
	// driver_info驾驶员信息表的driver_id
	DriverID string `json:"driver_id"`
	// 指纹名称
	FingerprintName *string `json:"fingerprint_name"`
	// driver_fingerprint 驾驶员指纹表的driver_fingerprint_id
	DriverFingerprintID *string `json:"driver_fingerprint_id"`
	// 图片地址
	PictureAddress *string `json:"picture_address"`
	// 是否删除
	IsDeleted *bool `json:"is_deleted"`
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 创建人                                                     ( system_user表的user_id )
	CreatedBy string `json:"created_by"`
	// 修改时间
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人                                                     ( system_user表的user_id )
	UpdatedBy *string `json:"updated_by"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人                                                     ( system_user表的user_id )
	DeletedBy *string `json:"deleted_by"`
}
