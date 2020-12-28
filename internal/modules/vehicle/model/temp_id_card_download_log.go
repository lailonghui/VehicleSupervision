// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// 临时工号牌下载记录表
type TempIDCardDownloadLog struct {
	// 按指定方法生成                                  ( 主键                       )
	ID int64 `json:"id"`
	// 外部编码，由golang程序生成的xid，暴露到外部使用 ( 联合主键                   )
	TempIDCardDownloadID string `json:"temp_id_card_download_id"`
	// vehicle_info 车辆信息表 的vehicle_id        (                            )
	VehicleID string `json:"vehicle_id"`
	// 有效期起始                                      (                            )
	ValidFrom *time.Time `json:"valid_from"`
	// 有效期截止                                      (                            )
	ValidUntil *time.Time `json:"valid_until"`
	// 操作人                                          ( system_user表的user_id )
	Operator *string `json:"operator"`
	// 是否删除                                        (                            )
	IsDeleted *bool `json:"is_deleted"`
	// 创建时间                                        (                            )
	CreatedAt time.Time `json:"created_at"`
	// 创建人                                          ( system_user表的user_id )
	CreatedBy string `json:"created_by"`
	// 修改时间                                        (                            )
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人                                          ( system_user表的user_id )
	UpdatedBy *string `json:"updated_by"`
	// 删除时间                                        (                            )
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人                                          ( system_user表的user_id )
	DeletedBy *string `json:"deleted_by"`
}
