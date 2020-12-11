/*
@Time : 2020/12/11 14:41
@Author : lai
@Description :
@File : vehicle_violation_details
*/
package model

import "time"

// 车辆违章明细表
//
//
// columns and relationships of "vehicle_violation_details"
type VehicleViolationDetails struct {
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 创建人
	CreatedBy string `json:"created_by"`
	// 决定书号
	DecisionNumber *string `json:"decision_number"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人
	DeletedBy *string `json:"deleted_by"`
	// 发现机构
	DiscoveryAgency *string `json:"discovery_agency"`
	// 违章驾驶员id
	DriverID *string `json:"driver_id"`
	// 所在企业id
	EnterpriseID *string `json:"enterprise_id"`
	// 处理时间
	HandleAt *time.Time `json:"handle_at"`
	// 处理人
	HandleBy *string `json:"handle_by"`
	// 主键
	ID int64 `json:"id"`
	// 违法代码,VIO_CODEWFDM 违法描述字典表
	IllegalCode *string `json:"illegal_code"`
	// 车辆违法处理状态字典
	IllegalHandlingStatus *int `json:"illegal_handling_status"`
	// 违法地点
	IllegalLocation *string `json:"illegal_location"`
	// 违法照片
	IllegalPhoto *string `json:"illegal_photo"`
	// 违法时间
	IllegalTime *time.Time `json:"illegal_time"`
	// 信息来源：1，强制，2，非现场，0，简易（信息来源字典表）
	InformationSource *int `json:"information_source"`
	// 是否删除
	IsDelete *bool `json:"is_delete"`
	// 是否处理
	IsHandle *bool `json:"is_handle"`
	// 是否通知驾驶员
	IsNoticeDriver *bool `json:"is_notice_driver"`
	// 是否发送短信
	IsSend *bool `json:"is_send"`
	// 实测值,车辆实际行驶的车速或实际载的人数，根据违法的种类不同而不同。
	MeasuredValue *string `json:"measured_value"`
	// 通知时间
	NoticeTime *time.Time `json:"notice_time"`
	// 当事人姓名
	PartyName *string `json:"party_name"`
	// 缴款标记,是否缴款字典
	PaymentMark *int `json:"payment_mark"`
	// 标准值,路段的限速阈值或核载的人数，根据违法的种类不同而不同。
	StandardValue *string `json:"standard_value"`
	// 内网更新时间
	UpdateTimeIn *time.Time `json:"update_time_in"`
	// 修改时间
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人
	UpdatedBy *string `json:"updated_by"`
	// 违章车辆id
	VehicleID *string `json:"vehicle_id"`
	// 驾驶人处理的交通违法记录对应的机动车信息
	VehicleInformation *string `json:"vehicle_information"`
	// 联合主键
	ViolationDetailID string `json:"violation_detail_id"`
}

func (VehicleViolationDetails) TableName() string {
	return "vehicle_violation_details"
}
