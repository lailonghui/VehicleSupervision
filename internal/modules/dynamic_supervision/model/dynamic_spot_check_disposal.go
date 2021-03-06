// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// 动态抽查处置表
type DynamicSpotCheckDisposal struct {
	// 按指定方法生成                           ( 主键                                                         )
	ID int64 `json:"id"`
	// vehicle_info 车辆信息表 的vehicle_id
	VehicleID string `json:"vehicle_id"`
	// 所在企业id                               ( enterprise_info表的enterprise_id                         )
	EnterpriseID *string `json:"enterprise_id"`
	// 手机号码
	Telephone *string `json:"telephone"`
	// 发送时间
	SendTime *time.Time `json:"send_time"`
	// 动态监管抽查明细表id                     ( dynamic_supervision_ detail 动态监管抽查明细表的supervision_detail_id )
	SupervisionDetailID *string `json:"supervision_detail_id"`
	// 图像异常处置措施
	ImageAbnormalHandingMeasure *string `json:"image_abnormal_handing_measure"`
	// 反馈时间
	FeedbackTime *time.Time `json:"feedback_time"`
	// 行车记录仪数据处置措施
	TachographDataDisposalMeasure *string `json:"tachograph_data_disposal_measure"`
	// 操作用户                                 ( system_user表的user_id                                   )
	OperationUser *string `json:"operation_user"`
	// 是否短信推送
	IsSmsPush *bool `json:"is_sms_push"`
	// 是否通报
	IsNotify *bool `json:"is_notify"`
	// 是否语音通知
	IsAnnounce *bool `json:"is_announce"`
	// 是否APP推送
	IsAppPush *bool `json:"is_app_push"`
	// 通报内容
	NotifyContent *string `json:"notify_content"`
	// 语音内容
	AnnounceContent *string `json:"announce_content"`
	// APP推送内容
	AppPushContent *string `json:"app_push_content"`
	// 处置内容
	DisposalContent *string `json:"disposal_content"`
	// 处置方式                                 ( 处置方式字典                                             )
	DisposalMethod *int `json:"disposal_method"`
	// 处置结果
	DisposalResult *string `json:"disposal_result"`
	// 是否删除                                 ( false                                                        )
	IsDeleted *bool `json:"is_deleted"`
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 创建人                                   ( system_user表的user_id                                   )
	CreatedBy string `json:"created_by"`
	// 修改时间
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人                                   ( system_user表的user_id                                   )
	UpdatedBy *string `json:"updated_by"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人                                   ( system_user表的user_id                                   )
	DeletedBy *string `json:"deleted_by"`
}
