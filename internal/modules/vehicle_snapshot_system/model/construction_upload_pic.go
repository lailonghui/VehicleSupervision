// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// 工地上传图片表
type ConstructionUploadPic struct {
	// 按指定方法生成                                  ( 主键                                                   )
	ID int64 `json:"id"`
	// 外部编码，由golang程序生成的xid，暴露到外部使用 ( 联合主键                                               )
	ConstructionUploadPicID string `json:"construction_upload_pic_id"`
	// 工地ID                                          ( construction_info 工地信息表的construction_info_id )
	ConstructionInfoID *string `json:"construction_info_id"`
	// 图片                                            (                                                        )
	PictureURL *string `json:"picture_url"`
	// 是否删除                                        (                                                        )
	IsDeleted *bool `json:"is_deleted"`
	// 创建时间                                        (                                                        )
	CreatedAt time.Time `json:"created_at"`
	// 创建人                                          ( system_user表的user_id                             )
	CreatedBy string `json:"created_by"`
	// 修改时间                                        (                                                        )
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人                                          ( system_user表的user_id                             )
	UpdatedBy *string `json:"updated_by"`
	// 删除时间                                        (                                                        )
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人                                          ( system_user表的user_id                             )
	DeletedBy *string `json:"deleted_by"`
}
