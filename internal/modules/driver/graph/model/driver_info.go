/*
@Time : 2020/12/11 10:53
@Author : lai
@Description :
@File : driver_info
*/
package model

import "time"

// 驾驶员信息表
//
//
// columns and relationships of "driver_info"
type DriverInfo struct {
	// 代理商
	Agent *string `json:"agent"`
	// 检验时间
	CheckAt *time.Time `json:"check_at"`
	// 校验人
	CheckBy *string `json:"check_by"`
	// 联系地址
	ContactAddress *string `json:"contact_address"`
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 创建人
	CreatedBy string `json:"created_by"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人
	DeletedBy *string `json:"deleted_by"`
	// 所在部门id
	DepartmentID *string `json:"department_id"`
	// 联合主键
	DriverID string `json:"driver_id"`
	// 驾驶员身份验证信息ID
	DriverIdentityID *string `json:"driver_identity_id"`
	// 驾驶员姓名
	DriverName *string `json:"driver_name"`
	// 所在企业id
	EnterpriseID *string `json:"enterprise_id"`
	// 档案编号(后6位)
	FilesNumber *string `json:"files_number"`
	// 主键
	ID int64 `json:"id"`
	// 录入时间
	InputAt *time.Time `json:"input_at"`
	// 录入人
	InputBy *string `json:"input_by"`
	// 是否黑名单
	IsBlack *bool `json:"is_black"`
	// 是否校验数据,该字段代表是否用于校验驾驶员信息，未正式录入系统，但会同步到公安内容，用于查询驾驶员的违章数据。
	IsCheckData *bool `json:"is_check_data"`
	// 是否通过短信验证
	IsCheckSms *bool `json:"is_check_sms"`
	// 是否删除
	IsDelete *bool `json:"is_delete"`
	// 是否录入
	IsInput *bool `json:"is_input"`
	// 是否手动录入,驾驶员资料分为使用身份证读卡器读取身份证自动录入资料和手动填写资料
	IsManualInput *bool `json:"is_manual_input"`
	// 是否提交,用于标志驾驶员资料是否处于确定状态。未确定状态的驾驶员信息在系统上除驾驶员管理外的功能中都查不到。
	IsSubmit *bool `json:"is_submit"`
	// 邮寄地址
	MailingAddress *string `json:"mailing_address"`
	// 运营商
	Operator *string `json:"operator"`
	// 驾驶员信息同步内网反馈信息。驾驶员信息同步到公安内网后内网的反馈内容，如档案编号填写错误会反馈档案编号后六位不正确
	RemarkIn *string `json:"remark_in"`
	// 备注
	Remarks *string `json:"remarks"`
	// 性别字典
	Sex *int `json:"sex"`
	// 提交时间
	SubmitAt *time.Time `json:"submit_at"`
	// 提交人
	SubmitBy *string `json:"submit_by"`
	// 提交内容
	SubmitContent *string `json:"submit_content"`
	// 手机号码
	Telephone *string `json:"telephone"`
	// 内网更新时间
	UpdateTimeIn *time.Time `json:"update_time_in"`
	// 修改时间
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人
	UpdatedBy *string `json:"updated_by"`
}

func (DriverInfo) TableName() string {
	return "driver_info"
}
