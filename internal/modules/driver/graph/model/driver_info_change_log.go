/*
@Time : 2020/12/11 10:54
@Author : lai
@Description :
@File : driver_info_change_log
*/
package model

import "time"

// 驾驶员信息表
//
//
// columns and relationships of "driver_info_change_log"
type DriverInfoChangeLog struct {
	// 累计积分（六合一）
	AccumulativedPoints *float64 `json:"accumulatived_points"`
	// 代理商
	Agent *string `json:"agent"`
	// 年审日期（六合一）
	AnnualReviewDate *time.Time `json:"annual_review_date"`
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
	// 驾驶员手持身份证照片,云储存系统返回的路径
	DriverHoldingIDPhoto *string `json:"driver_holding_id_photo"`
	// 联合主键
	DriverID string `json:"driver_id"`
	// 驾驶员身份验证信息ID
	DriverInfoChangeLogID *string `json:"driver_info_change_log_id"`
	// 驾驶证发证所在地的城市ID
	DriverLicenseCityID *string `json:"driver_license_city_id"`
	// 驾驶证发证所在地的区域ID
	DriverLicenseDistrictID *string `json:"driver_license_district_id"`
	// 驾驶证初次领证日期
	DriverLicenseIssueDate *time.Time `json:"driver_license_issue_date"`
	// 驾驶证发证机关
	DriverLicenseIssuingAuthority *string `json:"driver_license_issuing_authority"`
	// 驾驶员驾驶证,云储存系统返回的路径
	DriverLicensePic *string `json:"driver_license_pic"`
	// 驾驶证发证所在地的省份ID
	DriverLicenseProvinceID *string `json:"driver_license_province_id"`
	// 驾驶证状态字典
	DriverLicenseStatus *int `json:"driver_license_status"`
	// 驾驶员姓名
	DriverName *string `json:"driver_name"`
	// 驾驶员的正面照,云储存系统返回的路径
	DriverPhoto *string `json:"driver_photo"`
	// 驾驶员签名,云储存系统返回的路径
	DriverSignature *string `json:"driver_signature"`
	// 所在企业id
	EnterpriseID *string `json:"enterprise_id"`
	// 档案编号(后6位)
	FilesNumber *string `json:"files_number"`
	// 主键
	ID int64 `json:"id"`
	// 身份证住址
	IDCardAddress *string `json:"id_card_address"`
	// 身份证背面照，云存储地址
	IDCardBackPic *string `json:"id_card_back_pic"`
	// 身份证出生日期
	IDCardBirthday *time.Time `json:"id_card_birthday"`
	// 身份证有效截止日期
	IDCardEndDate *time.Time `json:"id_card_end_date"`
	// 身份证正面照，云存储地址
	IDCardFrontPic *string `json:"id_card_front_pic"`
	// 身份证民族
	IDCardNation *string `json:"id_card_nation"`
	// 身份证号码
	IDCardNum *string `json:"id_card_num"`
	// 身份证签发机关
	IDCardSignGovernment *string `json:"id_card_sign_government"`
	// 身份证有效起始日期
	IDCardStartDate *time.Time `json:"id_card_start_date"`
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
	// 是否审核
	IsReview *bool `json:"is_review"`
	// 是否提交,用于标志驾驶员资料是否处于确定状态。未确定状态的驾驶员信息在系统上除驾驶员管理外的功能中都查不到。
	IsSubmit *bool `json:"is_submit"`
	// 劳动合同,云储存系统返回的完整劳动合同的图片路径
	LaborContract *string `json:"labor_contract"`
	// 邮寄地址
	MailingAddress *string `json:"mailing_address"`
	// 从业资格证有效期至
	OccupationalExpireDate *time.Time `json:"occupational_expire_date"`
	// 从业资格证发证机构
	OccupationalIssuingAuthority *string `json:"occupational_issuing_authority"`
	// 从业资格证号码
	OccupationalNumber *string `json:"occupational_number"`
	// 运营商
	Operator *string `json:"operator"`
	// 准驾车型（六合一）,准驾车型字典
	QuasiDrivingModels *int `json:"quasi_driving_models"`
	// 驾驶员信息同步内网反馈信息。驾驶员信息同步到公安内网后内网的反馈内容，如档案编号填写错误会反馈档案编号后六位不正确
	RemarkIn *string `json:"remark_in"`
	// 备注
	Remarks *string `json:"remarks"`
	// 换证日期（六合一）
	RenewalDate *time.Time `json:"renewal_date"`
	// 性别字典
	Sex *int `json:"sex"`
	// 清分日期（六合一）
	SortingDate *time.Time `json:"sorting_date"`
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
	// 从业时间
	WorkingTime *time.Time `json:"working_time"`
}

func (DriverInfoChangeLog) TableName() string {
	return "driver_info_change_log"
}
