/*
@Time : 2020/12/11 14:36
@Author : lai
@Description :
@File : dynamic_spot_check_disposal
*/
package model

import "time"

// 动态抽查处置表
//
//
// columns and relationships of "dynamic_spot_check_disposal"
type DynamicSpotCheckDisposal struct {
	// 语音内容
	AnnounceContent *string `json:"announce_content"`
	// APP推送内容
	AppPushContent *string `json:"app_push_content"`
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 创建人
	CreatedBy string `json:"created_by"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人
	DeletedBy *string `json:"deleted_by"`
	// 处置内容
	DisposalContent *string `json:"disposal_content"`
	// 处置方式字典
	DisposalMethod *int `json:"disposal_method"`
	// 处置结果
	DisposalResult *string `json:"disposal_result"`
	// 所在企业id
	EnterpriseID *string `json:"enterprise_id"`
	// 反馈时间
	FeedbackTime *time.Time `json:"feedback_time"`
	// 主键
	ID int64 `json:"id"`
	// 图像异常处置措施
	ImageAbnormalHandingMeasure *string `json:"image_abnormal_handing_measure"`
	// 是否语音通知
	IsAnnounce *bool `json:"is_announce"`
	// 是否APP推送
	IsAppPush *bool `json:"is_app_push"`
	// 是否删除
	IsDelete *bool `json:"is_delete"`
	// 是否通报
	IsNotify *bool `json:"is_notify"`
	// 是否短信推送
	IsSmsPush *bool `json:"is_sms_push"`
	// 通报内容
	NotifyContent *string `json:"notify_content"`
	// 操作用户
	OperationUser *string `json:"operation_user"`
	// 发送时间
	SendTime *time.Time `json:"send_time"`
	// dynamic_supervision_ detail 动态监管抽查明细表的supervision_detail_id
	SupervisionDetailID *string `json:"supervision_detail_id"`
	// 行车记录仪数据处置措施
	TachographDataDisposalMeasure *string `json:"tachograph_data_disposal_measure"`
	// 手机号码
	Telephone *string `json:"telephone"`
	// 修改时间
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人
	UpdatedBy *string `json:"updated_by"`
	// vehicle_info 车辆信息表 的vehicle_id
	VehicleID string `json:"vehicle_id"`
}

func (DynamicSpotCheckDisposal) TableName() string {
	return "dynamic_spot_check_disposal"
}
