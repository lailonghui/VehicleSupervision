/*
@Time : 2020/12/3 16:17
@Author : lai
@Description :
@File : model_vehicle
*/
package model_vehicle

import "time"

// 车辆信息
type VehicleInfo struct {
	// 主键id
	ID string `json:"id"`
	// 车辆外部编码,由golang程序生成的xid，暴露到外部使用
	VehicleID string `json:"vehicle_id"`
	// 车牌号
	LicensePlateNumber *string `json:"license_plate_number"`
	// 车牌颜色
	LicensePlateColor *int `json:"license_plate_color"`
	// 号牌种类
	LicensePlateType *int `json:"license_plate_type"`
	// 所在企业id
	EnterpriseID *string `json:"enterprise_id"`
	// 所在部门id
	DepartmentID *string `json:"department_id"`
	// 车架号(后6位)
	VehicleIdentificationNumber *string `json:"vehicle_identification_number"`
	// 道路运输证号
	RoadTransportLicenseNumber *string `json:"road_transport_license_number"`
	// 车辆类型
	VehicleType *int `json:"vehicle_type"`
	// 行业类别
	IndustryCategory *int `json:"industry_category"`
	// 吨位
	Heavy *float64 `json:"heavy"`
	// 座位
	Seats *int `json:"seats"`
	// 营运类型
	OperatingType *int `json:"operating_type"`
	// 营运线路
	OperatingRoute *string `json:"operating_route"`
	// 经营范围
	BusinessScope *int `json:"business_scope"`
	// 机动车管理人
	VehicleManager *string `json:"vehicle_manager"`
	// 机动车管理人联系电话
	VehicleManagerPhone *string `json:"vehicle_manager_phone"`
	// 机动车管理人身份证
	VehicleManagerIDCard *string `json:"vehicle_manager_id_card"`
	// 机动车所有人（六合一）
	Owner *string `json:"owner"`
	// 检验日期（六合一）
	InspectionDate *time.Time `json:"inspection_date"`
	// 报废日期(六合一)
	RetirementDate *time.Time `json:"retirement_date"`
	// 使用性质（六合一）
	UseNature *string `json:"use_nature"`
	// 机动车状态
	VehicleState *int `json:"vehicle_state"`
	// 内网更新时间
	UpdateTimeIn *time.Time `json:"update_time_in"`
	// 车辆信息同步内网反馈信息
	RemarkIn *string `json:"remark_in"`
	// 是否完成
	IsComplete *bool `json:"is_complete"`
	// 行驶证照片,云储存系统返回的路径
	DrivingLicenseePic *string `json:"driving_licensee_pic"`
	// 是否激活
	IsActive *bool `json:"is_active"`
	// 是否录入完成
	IsInput *bool `json:"is_input"`
	// 租车标准价格
	CarRentalPrice *float64 `json:"car_rental_price"`
	// 投保公司
	InsuranceCompany *int `json:"insurance_company"`
	// 投保日期
	InsuranceDate *time.Time `json:"insurance_date"`
	// 汽车排量
	VehicleDisplacement *string `json:"vehicle_displacement"`
	// 车辆品牌
	VehicleBrand *int `json:"vehicle_brand"`
	// 准驾车型
	QuasiDrivingModels *int `json:"quasi_driving_models"`
	// 是否上传省厅
	IsUploadProvince *bool `json:"is_upload_province"`
	// 营运状态
	OperatingState *int `json:"operating_state"`
	// 终端ID
	TerminalID *string `json:"terminal_id"`
	// 是否申请安装智能终端
	IsApplyInstallTerminal *bool `json:"is_apply_install_terminal"`
	// 校验状态
	CheckState *int `json:"check_state"`
	// 是否导入
	IsImport *bool `json:"is_import"`
	// 登记时间
	RecordAt *time.Time `json:"record_at"`
	// 登记人
	RecordBy *string `json:"record_by"`
	// 是否删除
	IsDelete *bool `json:"is_delete"`
	// 备注
	Remarks *string `json:"remarks"`
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
}

func (VehicleInfo) TableName() string {
	return "vehicle_info"
}
