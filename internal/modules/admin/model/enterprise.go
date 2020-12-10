package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden EnterpriseLoader int64 *VehicleSupervision/internal/modules/admin/model.Enterprise

// 企业
type Enterprise struct {
	// 协会审核时间
	AssociationReviewBy *time.Time `json:"association_review_by"`
	// 协会审核意见
	AssociationReviewOpinion *string `json:"association_review_opinion"`
	// 协会审核时间
	AssociationReviewTime *time.Time `json:"association_review_time"`
	// 所属大队ID
	BrigadeID *string `json:"brigade_id"`
	// 大队审核人
	BrigadeReviewBy *string `json:"brigade_review_by"`
	// 大队审核意见
	BrigadeReviewOpinion *string `json:"brigade_review_opinion"`
	// 大队审核时间
	BrigadeReviewTime *time.Time `json:"brigade_review_time"`
	// 营业执照到期日期
	BusinessLicenseExpiryDate *time.Time `json:"business_license_expiry_date"`
	// 营业执照发证日期
	BusinessLicenseIssuanceDate *time.Time `json:"business_license_issuance_date"`
	// 营业执照图片
	BusinessLicensePhoto *string `json:"business_license_photo"`
	// 业务办理扫描件照片
	BusinessPhoto *string `json:"business_photo"`
	// 经营范围
	BusinessScope *int `json:"business_scope"`
	// 审核状态
	CheckStatus *int `json:"check_status"`
	// 城市ID
	CityID *int64 `json:"city_id"`
	// 联系人
	ContactPersons *string `json:"contact_persons"`
	// 创建时间
	CreateAt *time.Time `json:"create_at"`
	// 创建人
	CreateBy *string `json:"create_by"`
	// 删除时间
	DeleteAt *time.Time `json:"delete_at"`
	// 删除人
	DeleteBy *string `json:"delete_by"`
	// 显示顺序
	DisplayNumber *int `json:"display_number"`
	// 区域ID
	DistrictID *int64 `json:"district_id"`
	// 企业地址
	EnterpriseAddress *string `json:"enterprise_address"`
	// 企业码
	EnterpriseCode *string `json:"enterprise_code"`
	// 企业ID
	EnterpriseID string `json:"enterprise_id"`
	// 企业级别
	EnterpriseLevel *int `json:"enterprise_level"`
	// 企业名称
	EnterpriseName *string `json:"enterprise_name"`
	// 企业性质
	EnterpriseNature *int `json:"enterprise_nature"`
	// 委托代理人
	EntrustedAgent *string `json:"entrusted_agent"`
	// 委托代理人-身份证号码
	EntrustedAgentIDCard *string `json:"entrusted_agent_id_card"`
	// 委托代理人身份证图片
	EntrustedAgentIDCardPhoto *string `json:"entrusted_agent_id_card_photo"`
	// 委托代理人-电话号码
	EntrustedAgentPhone *string `json:"entrusted_agent_phone"`
	// 传真号码
	FaxNumber *string `json:"fax_number"`
	// ID
	ID int64 `gorm:"primarykey" json:"id"`
	// 机构类别
	InstitutionCategory *int64 `json:"institution_category"`
	// 是否黑名单
	IsBlack *bool `json:"is_black"`
	// 是否删除
	IsDeleted bool `json:"is_deleted"`
	// 是否录入完成
	IsInput *bool `json:"is_input"`
	// 是否安装
	IsInstall *bool `json:"is_install"`
	// 是否上次省厅
	IsUploadProvince *bool `json:"is_upload_province"`
	// 企业法人代表
	LegalRepresentative *string `json:"legal_representative"`
	// 企业法人-身份证号码
	LegalRepresentativeIDCard *string `json:"legal_representative_id_card"`
	// 企业法人代表身份证-图片
	LegalRepresentativeIDCardPhoto *string `json:"legal_representative_id_card_photo"`
	// 企业法人代表联系电话
	LegalRepresentativePhone *string `json:"legal_representative_phone"`
	// 经营许可证图片
	OperatingLicensePhoto *string `json:"operating_license_photo"`
	// 组织机构代码(企业的营运证)
	OrganizationCode *string `json:"organization_code"`
	// 组织机构代码证照片
	OrganizationCodeCertificatePhoto *string `json:"organization_code_certificate_photo"`
	// 所属派出所ID
	PoliceStationID *string `json:"police_station_id"`
	// 省份ID
	ProvinceID *int64 `json:"province_id"`
	// 登记时间
	RecordAt *time.Time `json:"record_at"`
	// 登记人
	RecordBy *string `json:"record_by"`
	// 备注
	Remarks *string `json:"remarks"`
	// 记分
	Score *int `json:"score"`
	// 上级企业ID
	SuperiorEnterpriseID *string `json:"superior_enterprise_id"`
	// 修改时间
	UpdateAt *time.Time `json:"update_at"`
	// 修改人
	UpdateBy *string `json:"update_by"`
	// 内网更新时间
	UpdateTimeIn *time.Time `json:"update_time_in"`
}

func (e Enterprise) TableName() string {
	return "enterprise"
}

func (u *Enterprise) NewLoader() *EnterpriseLoader {
	return &EnterpriseLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []int64) ([]*Enterprise, []error) {
			var rs []*Enterprise
			db.DB.Model(&Enterprise{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
