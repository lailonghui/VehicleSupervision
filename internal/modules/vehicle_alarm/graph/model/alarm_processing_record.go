/*
@Time : 2020/12/11 14:38
@Author : lai
@Description :
@File : alarm_processing_record
*/
package model

import "time"

// 报警处理记录表
//
// columns and relationships of "alarm_processing_record"
type AlarmProcessingRecord struct {
	// vehicle_alarm_data报警数据表的alarm_data_id
	AlarmDataID string `json:"alarm_data_id"`
	// alarm_supervision_picture_ upload报警监管图片上传表的alarm_supervision_picture_id
	AlarmSupervisionPictureID string `json:"alarm_supervision_picture_id"`
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
	// 处置方式字典
	DisposalMethod *string `json:"disposal_method"`
	// 处置结果
	DisposalResult *string `json:"disposal_result"`
	// 主键
	ID int64 `json:"id"`
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
	// 处理内容
	ProcessingContent *string `json:"processing_content"`
	// 处理时间
	ProcessingTime *time.Time `json:"processing_time"`
	// 处理类型  1.超速报警  2.疲劳驾驶  3.工程报警  4.超三天断电报警  5.进出区域报警  6.进出区域报警  7.安检到期报警  11.进出工地报警
	ProcessingType *int `json:"processing_type"`
	// 修改时间
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人
	UpdatedBy *string `json:"updated_by"`
}

func (AlarmProcessingRecord) TableName() string {
	return "alarm_processing_record"
}
