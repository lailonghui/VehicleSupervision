// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// 车辆违章明细表
type VehicleViolationDetails struct {
	// 按指定方法生成                                              ( 主键                                                         )
	ID int64 `json:"id"`
	// 车辆违章明细外部编码，由golang程序生成的xid，暴露到外部使用 ( 联合主键                                                     )
	ViolationDetailID string `json:"violation_detail_id"`
	// 违章车辆id                                                  ( vehicle_info表的vehicle_id                               )
	VehicleID *string `json:"vehicle_id"`
	// 违章驾驶员id                                                ( driver_info表的driver_id                                 )
	DriverID *string `json:"driver_id"`
	// 所在企业id                                                  ( enterprise_info表的enterprise_id                         )
	EnterpriseID *string `json:"enterprise_id"`
	// 违法代码                                                    ( VIO_CODEWFDM 违法描述字典表                              )
	IllegalCode *string `json:"illegal_code"`
	// 违法时间
	IllegalTime *time.Time `json:"illegal_time"`
	// 违法处理状态                                                ( 车辆违法处理状态字典                                     )
	IllegalHandlingStatus *int `json:"illegal_handling_status"`
	// 违法地点
	IllegalLocation *string `json:"illegal_location"`
	// 标准值                                                      ( 路段的限速阈值或核载的人数，根据违法的种类不同而不同。       )
	StandardValue *string `json:"standard_value"`
	// 实测值                                                      ( 车辆实际行驶的车速或实际载的人数，根据违法的种类不同而不同。 )
	MeasuredValue *string `json:"measured_value"`
	// 发现机构
	DiscoveryAgency *string `json:"discovery_agency"`
	// 违法照片
	IllegalPhoto *string `json:"illegal_photo"`
	// 是否通知驾驶员
	IsNoticeDriver *bool `json:"is_notice_driver"`
	// 通知时间
	NoticeTime *time.Time `json:"notice_time"`
	// 决定书号
	DecisionNumber *string `json:"decision_number"`
	// 缴款标记                                                    ( 是否缴款字典                                             )
	PaymentMark *int `json:"payment_mark"`
	// 当事人姓名
	PartyName *string `json:"party_name"`
	// 信息来源：1，强制，2，非现场，0，简易                       ( 信息来源字典表                                           )
	InformationSource *int `json:"information_source"`
	// 驾驶人处理的交通违法记录对应的机动车信息
	VehicleInformation *string `json:"vehicle_information"`
	// 内网更新时间
	UpdateTimeIn *time.Time `json:"update_time_in"`
	// 是否处理
	IsHandle *bool `json:"is_handle"`
	// 处理人                                                      ( system_user表的user_id                                   )
	HandleBy *string `json:"handle_by"`
	// 处理时间
	HandleAt *time.Time `json:"handle_at"`
	// 是否发送短信
	IsSend *bool `json:"is_send"`
	// 是否删除
	IsDeleted *bool `json:"is_deleted"`
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 创建人                                                      ( system_user表的user_id                                   )
	CreatedBy string `json:"created_by"`
	// 修改时间
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人                                                      ( system_user表的user_id                                   )
	UpdatedBy *string `json:"updated_by"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人                                                      ( system_user表的user_id                                   )
	DeletedBy *string `json:"deleted_by"`
}
