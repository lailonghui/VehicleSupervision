/*
@Time : 2020/12/11 14:41
@Author : lai
@Description :
@File : illegal_photo
*/
package model

import "time"

// 违法照片表
//
//
// columns and relationships of "illegal_photo"
type IllegalPhoto struct {
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 创建人
	CreatedBy string `json:"created_by"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人
	DeletedBy *string `json:"deleted_by"`
	// 主键
	ID int64 `json:"id"`
	// 联合主键
	IllegalPhotoID string `json:"illegal_photo_id"`
	// 是否删除
	IsDelete *bool `json:"is_delete"`
	// 是否同步
	IsSynchronized *bool `json:"is_synchronized"`
	// 违法照片地址
	PictureAddress *string `json:"picture_address"`
	// 违法照片名称
	PictureName *string `json:"picture_name"`
	// 修改时间
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人
	UpdatedBy *string `json:"updated_by"`
}

func (IllegalPhoto) TableName() string {
	return "illegal_photo"
}
