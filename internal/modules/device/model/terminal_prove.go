// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// 终端证明表
type TerminalProve struct {
	// ID
	ID int64 `json:"id"`
	// 终端证明表ID
	ProveID string `json:"prove_id"`
	// 车牌号码
	PlateNumber string `json:"plate_number"`
	// 车牌颜色
	PlateColor *string `json:"plate_color"`
	// 车辆类型
	VehicleType *int `json:"vehicle_type"`
	// 合同编号
	ContractNum *string `json:"contract_num"`
	// 签定时间
	SignDate *time.Time `json:"sign_date"`
	// 合同有效期限至
	ContractEndTime *time.Time `json:"contract_end_time"`
	// 终端型号ID
	TerminalTypeID *string `json:"terminal_type_id"`
	// 终端ID
	TerminalID *string `json:"terminal_id"`
	// 终端imei
	TerminalImei *string `json:"terminal_imei"`
	// 安装时间
	InstallDate *time.Time `json:"install_date"`
	// 终端备案编号
	TerminalRecordNum *string `json:"terminal_record_num"`
	// 安装人签名照片
	InstallManSignPhoto *string `json:"install_man_sign_photo"`
	// 代理商ID
	ProxyrgID *string `json:"proxyrg_id"`
	// 是否具备电子上岗
	ElectronicInduction bool `json:"electronic_induction"`
	// 检测报告编号
	TestReportNum *string `json:"test_report_num"`
	// 批次
	Batch *string `json:"batch"`
	// 平台编号
	PlateformNum *string `json:"plateform_num"`
	// 平台备案编号
	PlatformRecordNum *string `json:"platform_record_num"`
	// 经营许可证编号
	BusinessLicenseNum *string `json:"business_license_num"`
	// 经营许可证截止日期
	BusinessLicenseEndDate *time.Time `json:"business_license_end_date"`
	// 登记保护备案编号
	LevelProtectRecordNum *string `json:"level_protect_record_num"`
	// 登记保护签发日期
	LevelProtectSignDate *time.Time `json:"level_protect_sign_date"`
	// 车属单位ID
	EnterpriseID *string `json:"enterprise_id"`
	// 办理人ID
	FillManID *string `json:"fill_man_id"`
	// 审查人ID
	VerifyManID *string `json:"verify_man_id"`
	// 提交人ID
	SealManID *string `json:"seal_man_id"`
	// 办理人姓名
	FillManName *string `json:"fill_man_name"`
	// 审查人姓名
	VerifyManName *string `json:"verify_man_name"`
	// 提交人姓名
	SealManName *string `json:"seal_man_name"`
	// 审查阶段
	CheckStatus *int `json:"check_status"`
	// 是否审核完成
	IsCheckEnd bool `json:"is_check_end"`
	// 审核时间
	CheckTime *time.Time `json:"check_time"`
	// 办理时间
	FillTime *time.Time `json:"fill_time"`
	// 办理单号
	FillID *string `json:"fill_id"`
	// 车辆ID
	VehicleID *string `json:"vehicle_id"`
	// 审核备注
	CheckRemark *string `json:"check_remark"`
	// 办理备注
	FillRemark *string `json:"fill_remark"`
	// 删除备注
	DelRemark *string `json:"del_remark"`
	// 工程车自编号
	SvNum *string `json:"sv_num"`
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 创建人
	CreatedBy *string `json:"created_by"`
	// 更新时间
	UpdatedAt *time.Time `json:"updated_at"`
	// 更新人
	UpdatedBy *string `json:"updated_by"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除时间
	DeletedBy *string `json:"deleted_by"`
	// 是否删除
	IsDeleted bool `json:"is_deleted"`
}
