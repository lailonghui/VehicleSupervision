// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// 渣土车购车意向
type MuckTruckPurchaseIntention struct {
	// 按指定方法生成                                  ( 主键                                 )
	ID int64 `json:"id"`
	// 外部编码，由golang程序生成的xid，暴露到外部使用 ( 联合主键                             )
	MuckTruckPurchaseIntentionID string `json:"muck_truck_purchase_intention_id"`
	// 供应商                                          ( enterprise_info表的enterprise_id )
	Supplier *string `json:"supplier"`
	// 购车用户姓名
	CustomerName *string `json:"customer_name"`
	// 购车用户电话
	CustomerPhone *string `json:"customer_phone"`
	// 所有人企业                                      ( enterprise_info表的enterprise_id )
	OwnerEnterprise *string `json:"owner_enterprise"`
	// 所有人所在省                                    ( 省份表province_id                )
	ProvinceID *string `json:"province_id"`
	// 所有人所在市                                    ( 城市表city_id                    )
	CityID *string `json:"city_id"`
	// 所有人所在县                                    ( 区域表district_id                )
	DistrictID *string `json:"district_id"`
	// 运力申请
	CapacigyApplication *string `json:"capacigy_application"`
	// 品牌型号
	BrandModel *string `json:"brand_model"`
	// 购车数量
	VehiclePurchase *int `json:"vehicle_purchase"`
	// 编码
	Code *string `json:"code"`
	// 登记日期
	RegistrationTime *time.Time `json:"registration_time"`
	// 登记用户                                        ( system_user表的user_id           )
	RegistrationUser *string `json:"registration_user"`
	// 审核
	Review *string `json:"review"`
	// 备注
	Remarks *string `json:"remarks"`
	// 是否删除
	IsDeleted *bool `json:"is_deleted"`
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 创建人                                          ( system_user表的user_id           )
	CreatedBy string `json:"created_by"`
	// 修改时间
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人                                          ( system_user表的user_id           )
	UpdatedBy *string `json:"updated_by"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人                                          ( system_user表的user_id           )
	DeletedBy *string `json:"deleted_by"`
}