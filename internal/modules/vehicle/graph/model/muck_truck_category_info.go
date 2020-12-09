/*
@Time : 2020/12/9 10:48
@Author : lai
@Description :
@File : muck_truck_category_info
*/
package model

import "time"

// 渣土车目录库车辆信息表
//
//
// columns and relationships of "muck_truck_category_info"
type MuckTruckCategoryInfo struct {
	// 实际车主联系电话
	ActualOwnerContactPhone *string `json:"actual_owner_contact_phone"`
	// 实际车主身份证号
	ActualOwnerIDNumber *string `json:"actual_owner_id_number"`
	// 实际车主身份证照片
	ActualOwnerIDPhoto *string `json:"actual_owner_id_photo"`
	// 实际车主姓名
	ActualOwnerName *string `json:"actual_owner_name"`
	// 年检到期时间
	AnnualInspectionExpirationTime *time.Time `json:"annual_inspection_expiration_time"`
	// 申请验车时间
	ApplyInspectionTime *time.Time `json:"apply_inspection_time"`
	// 轴数
	AxesNumber *int `json:"axes_number"`
	// 黑名单截止日期
	BlacklistDeadline *time.Time `json:"blacklist_deadline"`
	// 交强险保单图片
	CompulsoryInsurancePolicyPicture *string `json:"compulsory_insurance_policy_picture"`
	// 创建时间
	CreateAt time.Time `json:"create_at"`
	// 创建人
	CreateBy string `json:"create_by"`
	// 删除时间
	DeleteAt *time.Time `json:"delete_at"`
	// 删除人
	DeleteBy *string `json:"delete_by"`
	// 行驶证登记的车主联系电话
	DrivingLicenseContactPhone *string `json:"driving_license_contact_phone"`
	// 行驶证登记的车主身份证号
	DrivingLicenseIDNumber *string `json:"driving_license_id_number"`
	// 行驶证登记的车主身份证照片
	DrivingLiscenseOwnerIDPhoto *string `json:"driving_liscense_owner_id_photo"`
	// 发动机号
	EngineNumber *string `json:"engine_number"`
	// 主键
	ID int64 `json:"id"`
	// 违法通知书编号签注
	IllegalNoticeNumberEndorsement *string `json:"illegal_notice_number_endorsement"`
	// 违法编号签注
	IllegalNumberEndorsement *string `json:"illegal_number_endorsement"`
	// 事故编号签注
	IncidentNumberEndorsement *string `json:"incident_number_endorsement"`
	// 初次登记日期
	InitialRegistrationDate *time.Time `json:"initial_registration_date"`
	// 保险到期时间
	InsuranceExpiryTime *time.Time `json:"insurance_expiry_time"`
	// 是否黑名单
	IsBlacklist *bool `json:"is_blacklist"`
	// 是否删除
	IsDelete *bool `json:"is_delete"`
	// 是否安检到期报警处理
	IsDueSecurityAlarmProcessing *bool `json:"is_due_security_alarm_processing"`
	// 是否首次注册
	IsFirstRegister *bool `json:"is_first_register"`
	// 是否渣土办审核
	IsMuckOfficeAudit *bool `json:"is_muck_office_audit"`
	// 是否发送短信
	IsSendSms *bool `json:"is_send_sms"`
	// 泉工号发放日期
	IssueDateOfQuangongNumber *time.Time `json:"issue_date_of_quangong_number"`
	// 车主身份证照片
	OwnerIDPhoto *string `json:"owner_id_photo"`
	// 处理备注
	ProcessingNotes *string `json:"processing_notes"`
	// 处理时间
	ProcessingTime *time.Time `json:"processing_time"`
	// 处理人
	Processor *string `json:"processor"`
	// 审核备注
	ReviewNotes *string `json:"review_notes"`
	// 审核状态
	ReviewStatus *string `json:"review_status"`
	// 审核时间
	ReviewTime *time.Time `json:"review_time"`
	// 审核人
	Reviewer *string `json:"reviewer"`
	// 二级维护检测到期时间
	SecondaryMaintenanceExpiryDate *time.Time `json:"secondary_maintenance_expiry_date"`
	// 二级维护地点
	SecondaryMaintenanceLocation *string `json:"secondary_maintenance_location"`
	// 签收验车申请时间
	SigningAcceptanceApplicationTime *time.Time `json:"signing_acceptance_application_time"`
	// 第三者保额
	ThirdPartyInsuranceCoverage *string `json:"third_party_insurance_coverage"`
	// 第三者保险到期时间
	ThirdPartyInsuranceExpiryTime *time.Time `json:"third_party_insurance_expiry_time"`
	// 第三者保单图片
	ThirdPartyInsurancePolicyPicture *string `json:"third_party_insurance_policy_picture"`
	// 总质量
	TotalMass *float64 `json:"total_mass"`
	// 修改时间
	UpdateAt *time.Time `json:"update_at"`
	// 修改人
	UpdateBy *string `json:"update_by"`
	// 车辆描述
	VehicleDescription *string `json:"vehicle_description"`
	// vehicle_info 车辆信息表 的vehicle_id
	VehicleID string `json:"vehicle_id"`
	// 车辆性质  1.本企业车辆  2.企业挂靠车辆  3.车队挂靠车辆
	VehicleNature *int `json:"vehicle_nature"`
	// 车辆营运证号
	VehicleOperatingCertificateNumber *string `json:"vehicle_operating_certificate_number"`
	// 车辆营运证照片
	VehicleOperatingCertificatePhoto *string `json:"vehicle_operating_certificate_photo"`
	// 车辆图片
	VehiclePicture *string `json:"vehicle_picture"`
}
