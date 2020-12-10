/*
@Time : 2020/12/9 10:45
@Author : lai
@Description :
@File : vehicle_info
*/
package model

import "time"

// 车辆信息主表
//
//
// columns and relationships of "vehicle_info"
type VehicleInfo struct {
	// 主键
	ID int64 `json:"id" gorm:"column:id;default:id_generator();"`
	// 车辆外部编码，由golang程序生成的xid，暴露到外部使用
	VehicleID string `json:"vehicle_id" gorm:"column:vehicle_id"`
	// 经营范围字典
	BusinessScope *int `json:"business_scope" `
	// 租车标准价格
	CarRentalPrice *float64 `json:"car_rental_price"`
	// 校验状态
	CheckState *int `json:"check_state"`
	// 创建时间
	CreateAt time.Time `json:"create_at" gorm:"column:create_at"`
	// 创建人,
	CreateBy string `json:"create_by"`
	// 删除时间
	DeleteAt *time.Time `json:"delete_at"`
	// 删除人
	DeleteBy *string `json:"delete_by"`
	// 所在部门id,department 部门信息表
	DepartmentID *string `json:"department_id"`
	// 行驶证照片,云储存系统返回的路径
	DrivingLicenseePic *string `json:"driving_licensee_pic"`
	// 所在企业id,enterprise_info表的enterprise_id
	EnterpriseID *string `json:"enterprise_id"`
	// 吨位
	Heavy *float64 `json:"heavy"`
	// 行业类别字典
	IndustryCategory *int `json:"industry_category"`
	// 检验日期（六合一）
	InspectionDate *time.Time `json:"inspection_date"`
	// 投保公司
	InsuranceCompany *int `json:"insurance_company"`
	// 投保日期
	InsuranceDate *time.Time `json:"insurance_date"`
	// 是否激活
	IsActive *bool `json:"is_active"`
	// 是否申请安装智能终端
	IsApplyInstallTerminal *bool `json:"is_apply_install_terminal"`
	// 是否完成;用于标志车辆资料是否处于确定状态。未确定状态的车辆信息在系统上除车辆管理外的功能中都查不到
	IsComplete *bool `json:"is_complete"`
	// 是否删除
	IsDelete *bool `json:"is_delete"`
	// 是否导入
	IsImport *bool `json:"is_import"`
	// 是否录入完成
	IsInput *bool `json:"is_input"`
	// 是否上传省厅
	IsUploadProvince *bool `json:"is_upload_province"`
	// 车牌颜色字典
	LicensePlateColor *int `json:"license_plate_color"`
	// 车牌号
	LicensePlateNumber *string `json:"license_plate_number"`
	// 号牌种类字典
	LicensePlateType *int `json:"license_plate_type"`
	// muck_truck _info 渣土车信息表的id
	MuckTruckID *int64 `json:"muck_truck_id"`
	// 营运线路
	OperatingRoute *string `json:"operating_route"`
	// 营运状态字典
	OperatingState *int `json:"operating_state"`
	// 营运类型字典
	OperatingType *int `json:"operating_type"`
	// operating_vehi cle_ info 营运车信息表的id
	OperatingVehicleID *int64 `json:"operating_vehicle_id"`
	// 机动车所有人（六合一）
	Owner *string `json:"owner"`
	// 准驾车型
	QuasiDrivingModels *int `json:"quasi_driving_models"`
	// 登记时间
	RecordAt *time.Time `json:"record_at"`
	// 登记人,system_user表的user_id
	RecordBy *string `json:"record_by"`
	// 车辆信息同步内网反馈信息;车辆信息同步到公安内网后内网的反馈内容，如车牌号填写错误会反馈车辆号牌错误
	RemarkIn *string `json:"remark_in"`
	// 备注
	Remarks *string `json:"remarks"`
	// 报废日期（六合一）
	RetirementDate *time.Time `json:"retirement_date"`
	// 道路运输证号
	RoadTransportLicenseNumber *string `json:"road_transport_license_number"`
	// 座位
	Seats *int `json:"seats"`
	// 终端ID
	TerminalID *string `json:"terminal_id"`
	// 修改时间
	UpdateAt *time.Time `json:"update_at"`
	// 修改人
	UpdateBy *string `json:"update_by"`
	// 内网更新时间
	UpdateTimeIn *time.Time `json:"update_time_in"`
	// 使用性质（六合一）
	UseNature *string `json:"use_nature"`
	// 车辆品牌
	VehicleBrand *int `json:"vehicle_brand"`
	// 汽车排量
	VehicleDisplacement *string `json:"vehicle_displacement"`
	// 车架号(后6位),车辆识别代号vin,如D02133
	VehicleIdentificationNumber *string `json:"vehicle_identification_number"`
	// 维保数据数组，字段包括: 1.maintenance_ date 维保时间<br />2.maintenance_ kilometers 维保公里数
	VehicleMaintenances *string `json:"vehicle_maintenances"`
	// 机动车管理人
	VehicleManager *string `json:"vehicle_manager"`
	// 机动车管理人身份证
	VehicleManagerIDCard *string `json:"vehicle_manager_id_card"`
	// 机动车管理人联系电话
	VehicleManagerPhone *string `json:"vehicle_manager_phone"`
	// 机动车状态
	VehicleState *int `json:"vehicle_state"`
	// 车辆类型字典
	VehicleType *int `json:"vehicle_type"`
}

func (VehicleInfo) TableName() string {
	return "vehicle_info"
}
