// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// 汽车销售备案
type VehicleSaleRecord struct {
	// 按指定方法生成                                  ( 主键                                                    )
	ID int64 `json:"id"`
	// 外部编码，由golang程序生成的xid，暴露到外部使用 ( 联合主键                                                )
	VehicleSaleRecordID string `json:"vehicle_sale_record_id"`
	// 销售商                                          ( enterprise_info表的enterprise_id                    )
	Seller *string `json:"seller"`
	// 汽车型号                                        (                                                         )
	VehicleModel *string `json:"vehicle_model"`
	// 汽车相关图片的路径                              (                                                         )
	VehiclePicture *string `json:"vehicle_picture"`
	// 汽车参数                                        (                                                         )
	VehicleParameter *string `json:"vehicle_parameter"`
	// 参考报价                                        (                                                         )
	ReferencePrice *string `json:"reference_price"`
	// 其他相关材料                                    (                                                         )
	OtherMaterial *string `json:"other_material"`
	// 备注                                            (                                                         )
	Remarks *string `json:"remarks"`
	// 登记日期                                        (                                                         )
	RegistrationTime *time.Time `json:"registration_time"`
	// 登记用户                                        ( system_user表的user_id                              )
	RegistrationUser *string `json:"registration_user"`
	// 原因                                            (                                                         )
	Cause *string `json:"cause"`
	// 汽车品牌(简称)                                  ( new_muck_truck_recommend_catalog 新型渣土车推荐目录的id )
	CatalogID *string `json:"catalog_id"`
	// 轴数类型                                        ( 轴数类型字典                                        )
	AxisType *int `json:"axis_type"`
	// 运输方量                                        (                                                         )
	TransportVolume *string `json:"transport_volume"`
	// 是否审核                                        (                                                         )
	IsReview *bool `json:"is_review"`
	// 是否删除                                        (                                                         )
	IsDeleted *bool `json:"is_deleted"`
	// 创建时间                                        (                                                         )
	CreatedAt time.Time `json:"created_at"`
	// 创建人                                          ( system_user表的user_id                              )
	CreatedBy string `json:"created_by"`
	// 修改时间                                        (                                                         )
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人                                          ( system_user表的user_id                              )
	UpdatedBy *string `json:"updated_by"`
	// 删除时间                                        (                                                         )
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人                                          ( system_user表的user_id                              )
	DeletedBy *string `json:"deleted_by"`
}
