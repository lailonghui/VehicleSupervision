/*
@Time : 2020/12/11 10:53
@Author : lai
@Description :
@File : driver_identity
*/
package model

import "time"

// 驾驶员身份验证信息(各种证件信息，验证状态)
//
//
// columns and relationships of "driver_identity"
type DriverIdentity struct {
	// 累计积分（六合一）
	AccumulativedPoints *float64 `json:"accumulatived_points"`
	// 年审日期（六合一）
	AnnualReviewDate *time.Time `json:"annual_review_date"`
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 创建人
	CreatedBy string `json:"created_by"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人
	DeletedBy *string `json:"deleted_by"`
	// 驾驶员手持身份证照片,云储存系统返回的路径
	DriverHoldingIDPhoto *string `json:"driver_holding_id_photo"`
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
	// 驾驶员的正面照,云储存系统返回的路径
	DriverPhoto *string `json:"driver_photo"`
	// 驾驶员签名,云储存系统返回的路径
	DriverSignature *string `json:"driver_signature"`
	ID              int64   `json:"id"`
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
	// 联合主键
	IdentityID string `json:"identity_id"`
	// 是否删除
	IsDelete *bool `json:"is_delete"`
	// 是否审核
	IsReview *bool `json:"is_review"`
	// 劳动合同,云储存系统返回的完整劳动合同的图片路径
	LaborContract *string `json:"labor_contract"`
	// 从业资格证有效期至
	OccupationalExpireDate *time.Time `json:"occupational_expire_date"`
	// 从业资格证发证机构
	OccupationalIssuingAuthority *string `json:"occupational_issuing_authority"`
	// 从业资格证号码
	OccupationalNumber *string `json:"occupational_number"`
	// 准驾车型（六合一）,准驾车型字典
	QuasiDrivingModels *int `json:"quasi_driving_models"`
	// 换证日期（六合一）
	RenewalDate *time.Time `json:"renewal_date"`
	// 清分日期（六合一）
	SortingDate *time.Time `json:"sorting_date"`
	// 修改时间
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人
	UpdatedBy *string `json:"updated_by"`
	// 从业时间
	WorkingTime *time.Time `json:"working_time"`
}

func (DriverIdentity) TableName() string {
	return "driver_identity"
}
