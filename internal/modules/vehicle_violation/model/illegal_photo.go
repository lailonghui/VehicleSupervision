// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// 违法照片表
type IllegalPhoto struct {
	// 按指定方法生成                                            ( 主键                       )
	ID int64 `json:"id"`
	// 违法照片表外部编码，由golang程序生成的xid，暴露到外部使用 ( 联合主键                   )
	IllegalPhotoID string `json:"illegal_photo_id"`
	// 违法照片名称
	PictureName *string `json:"picture_name"`
	// 违法照片地址
	PictureAddress *string `json:"picture_address"`
	// 是否同步                                                  ( false                      )
	IsSynchronized *bool `json:"is_synchronized"`
	// 是否删除                                                  ( false                      )
	IsDeleted *bool `json:"is_deleted"`
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 创建人                                                    ( system_user表的user_id )
	CreatedBy string `json:"created_by"`
	// 修改时间
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人                                                    ( system_user表的user_id )
	UpdatedBy *string `json:"updated_by"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人                                                    ( system_user表的user_id )
	DeletedBy *string `json:"deleted_by"`
}
