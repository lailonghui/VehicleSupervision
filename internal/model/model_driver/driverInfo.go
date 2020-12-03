/*
@Time : 2020/12/3 16:21
@Author : lai
@Description :
@File : driverInfo
*/
package model_driver

import "time"

// 驾驶员信息
type DriverInfo struct {
	// 主键id
	ID string `json:"id"`
	// 驾驶员外部编码,由golang程序生成的xid，暴露到外部使用
	DriverID string `json:"driver_id"`
	// 所在企业id
	EnterpriseID *string `json:"enterprise_id"`
	// 所在部门id
	DepartmentID *string `json:"department_id"`
	// 驾驶员身份验证信息ID
	DriverIdentityID *string `json:"driver_identity_id"`
	// 驾驶员姓名
	DriverName *string `json:"driver_name"`
	// 手机号码
	Telephone *string `json:"telephone"`
	// 性别
	Sex int `json:"sex"`
	// 档案编号(后6位)
	FilesNumber *string `json:"files_number"`
	// 联系地址
	ContactAddress *string `json:"contact_address"`
	// 邮寄地址
	MailingAddress *string `json:"mailing_address"`
	// 是否提交
	IsSubmit *bool `json:"is_submit"`
	// 提交内容
	SubmitContent *string `json:"submit_content"`
	// 提交时间
	SubmitAt *time.Time `json:"submit_at"`
	// 提交人
	SubmitBy *string `json:"submit_by"`
	// 是否手动录入
	IsManualInput *bool `json:"is_manual_input"`
	// 是否录入
	IsInput *bool `json:"is_input"`
	// 录入时间
	InputAt *time.Time `json:"input_at"`
	// 录入人
	InputBy *string `json:"input_by"`
	// 是否校验数据
	IsCheckData *bool `json:"is_check_data"`
	// 检验时间
	CheckAt *time.Time `json:"check_at"`
	// 校验人
	CheckBy *string `json:"check_by"`
	// 驾驶员信息同步内网反馈信息
	RemarkIn *string `json:"remark_in"`
	// 内网更新时间
	UpdateTimeIn *time.Time `json:"update_time_in"`
	// 是否通过短信验证
	IsCheckSms *bool `json:"is_check_sms"`
	// 是否黑名单
	IsBlack *bool `json:"is_black"`
	// 创建时间
	CreateAt time.Time `json:"create_at"`
	// 创建人
	CreateBy string `json:"create_by"`
	// 修改时间
	UpdateAt time.Time `json:"update_at"`
	// 修改人
	UpdateBy string `json:"update_by"`
	// 删除时间
	DeleteAt *time.Time `json:"delete_at"`
	// 删除人
	DeleteBy *string `json:"delete_by"`
	// 是否删除
	IsDelete *bool `json:"is_delete"`
	// 备注
	Remarks *string `json:"remarks"`
}

func (DriverInfo) TableName() string {
	return "vehicle_info"
}
