/*
@Time : 2020/12/11 14:41
@Author : lai
@Description :
@File : dispute_violation_record
*/
package model

import "time"

// 违章争议记录表
//
//
// columns and relationships of "dispute_violation_record"
type DisputeViolationRecord struct {
	// 委托代理人身份证
	AgentIDNumber *string `json:"agent_id_number"`
	// 审批状态(车辆违法审批状态字典)
	ApproveState *int `json:"approve_state"`
	// 机动车所有人营业执照
	BusinessLicense *string `json:"business_license"`
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
	// 联合主键
	DisputeViolationID string `json:"dispute_violation_id"`
	// 行为人驾驶证
	DriverLicense *string `json:"driver_license"`
	// 机动车行驶证
	DrivingLicense *string `json:"driving_license"`
	// 行车日志
	DrivingLog *string `json:"driving_log"`
	// 主键
	ID int64 `json:"id"`
	// 行为人身份证
	IDCard *string `json:"id_card"`
	// 是否删除
	IsDelete *bool `json:"is_delete"`
	// 劳动合同或租赁合同
	LaborContract *string `json:"labor_contract"`
	// 法定代表人身份证
	LegalPersonIDNumber *string `json:"legal_person_id_number"`
	// 机动车所有人组织机构代码证
	OrganizationCode *string `json:"organization_code"`
	// 其他证据材料
	OtherEvidence *string `json:"other_evidence"`
	// 图像证据材料
	PicEvidence *string `json:"pic_evidence"`
	// 当事人陈述
	Statement *string `json:"statement"`
	// 内网更新时间
	UpdateTimeIn *time.Time `json:"update_time_in"`
	// 修改时间
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人
	UpdatedBy *string `json:"updated_by"`
	// 机动车管理人身份证
	VehicleManagerIDCard *string `json:"vehicle_manager_id_card"`
	// 违章明细表id(vehicle_violation _details 表的violation_detail_id)
	ViolationDetailID *string `json:"violation_detail_id"`
	// 证人证言
	Witness *string `json:"witness"`
	// 书面申请材料
	WrittenApplicationMaterials *string `json:"written_application_materials"`
}

func (DisputeViolationRecord) TableName() string {
	return "dispute_violation_record"
}
