package model


// 企业
type Enterprise struct {
	// 企业ID
	EnterpriseID string `json:"enterprise_id"`
	// 上级企业
	SuperiorEnterprise *string `json:"superior_enterprise"`
	// 企业编号
	EnterpriseCode string `json:"enterprise_code"`
	// 企业名称
	EnterpriseName string `json:"enterprise_name"`
	// 企业级别
	EnterpriseLevel int `json:"enterprise_level"`
	// 显示序号
	DisplayNumber *int `json:"display_number"`
	// 联系人信息
	ContactPersons string `json:"contact_persons"`
	// 企业地址
	EnterpriseAddress *string `json:"enterprise_address"`
	// 经营范围
	BusinessScope *int `json:"business_scope"`
	// 传真号码
	FaxNumber *string `json:"fax_number"`
	// 省份
	Province *int `json:"province"`
	// 城市
	City *int `json:"city"`
	// 区
	District *int `json:"district"`
	// 机构类别
	InstitutionCategory *int `json:"institution_category"`
	// 经营许可证
	OperatingLicense *string `json:"operating_license"`
	// 营业执照
	BusinessLicense *string `json:"business_license"`
	// 营业执照发证日期
	BusinessLicenseIssuanceDate *string `json:"business_license_issuance_date"`
	// 营业执照到期日期
	BusinessLicenseExpiryDate *string `json:"business_license_expiry_date"`
	// 企业性质
	EnterpriseNature *int `json:"enterprise_nature"`
	// 企业法人代表
	LegalRepresentative *string `json:"legal_representative"`
	// 企业法人代表联系电话
	LegalRepresentativePhone *string `json:"legal_representative_phone"`
	// 企业法人代表身份证
	LegalRepresentativeIDCard *string `json:"legal_representative_id_card"`
	// 企业法人身份证号
	LegalPersonIDNumber *string `json:"legal_person_id_number"`
	// 委托代理人
	EntrustedAgent *string `json:"entrusted_agent"`
	// 委托代理人联系电话
	EntrustedAgentPhone *string `json:"entrusted_agent_phone"`
	// 委托代理人身份证
	EntrustedAgentIDCard *string `json:"entrusted_agent_id_card"`
	// 委托代理人身份证号
	AgentIDNumber *string `json:"agent_id_number"`
	// 组织机构代码(企业的营运证)
	OrganizationCode *string `json:"organization_code"`
	// 组织机构代码证
	OrganizationCodeCertificate *string `json:"organization_code_certificate"`
	// 内网更新时间
	UpdateTimeIn *string `json:"update_time_in"`
	// 业务办理扫描件
	BusinessPic *string `json:"business_pic"`
	// 是否黑名单
	IsBlack *bool `json:"is_black"`
	// 高速支队
	HighSpeedDetachment *string `json:"high_speed_detachment"`
	// 高速大队
	HighSpeedBrigade *string `json:"high_speed_brigade"`
	// 高速中队
	HighSpeedSquadron *string `json:"high_speed_squadron"`
	// 高速交警级别
	TrafficPoliceLevel *int `json:"traffic_police_level"`
	// 是否网约车驾校
	IsRideHailingDrivingSchool *bool `json:"is_ride_hailing_driving_school"`
	// 是否网约车企业
	IsRideHailingEnterprise *bool `json:"is_ride_hailing_enterprise"`
	// 是否试用期结束
	IsExpire *bool `json:"is_expire"`
	// 审核状态
	IsCheck *int `json:"is_check"`
	// 是否安装
	IsInstall *bool `json:"is_install"`
	// 是否录入完成
	IsInput *bool `json:"is_input"`
	// 所属大队
	Brigade *string `json:"brigade"`
	// 所属派出所
	PoliceStation *string `json:"police_station"`
	// 大队审核意见
	BrigadeReviewOpinion *string `json:"brigade_review_opinion"`
	// 大队审核时间
	BrigadeReviewTime *string `json:"brigade_review_time"`
	// 大队审核人
	BrigadeReviewBy *string `json:"brigade_review_by"`
	// 是否上传省厅
	IsUploadProvince *bool `json:"is_upload_province"`
	// 记分
	Score *int `json:"score"`
	// 渣土车供应商编号
	SlagCarSupplierCode *string `json:"slag_car_supplier_code"`
	// 渣土车供应商状态
	SlagCarSupplierState *int `json:"slag_car_supplier_state"`
	// 协会审核意见
	AssociationReviewOpinion *string `json:"association_review_opinion"`
	// 协会审核时间
	AssociationReviewTime *string `json:"association_review_time"`
	// 协会审核人员
	AssociationReviewBy *string `json:"association_review_by"`
	// 创建时间
	CreateAt *string `json:"create_at"`
	// 创建人
	CreateBy *string `json:"create_by"`
	// 修改时间
	UpdateAt *string `json:"update_at"`
	// 修改人
	UpdateBy *string `json:"update_by"`
	// 删除时间
	DeleteAt *string `json:"delete_at"`
	// 删除人
	DeleteBy *string `json:"delete_by"`
	// 登记时间
	RecordAt *string `json:"record_at"`
	// 登记人
	RecordBy *string `json:"record_by"`
	// 备注
	Remarks *string `json:"remarks"`
}
