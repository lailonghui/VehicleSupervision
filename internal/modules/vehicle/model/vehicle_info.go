// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// 车辆信息表
type VehicleInfo struct {
	// 按指定方法生成，生成方法见下面说明                           ( 主键                                                         )
	ID int64 `json:"id"`
	// 车辆外部编码，由golang程序生成的xid，暴露到外部使用          ( 联合主键                                                     )
	VehicleID string `json:"vehicle_id"`
	// 所在企业id                                                   ( enterprise_info表的enterprise_id                         )
	EnterpriseID *string `json:"enterprise_id"`
	// 所在部门id                                                   ( department 部门信息表                                    )
	DepartmentID *string `json:"department_id"`
	// 行业类别                                                     ( 行业类别字典                                             )
	IndustryCategory *int `json:"industry_category"`
	// 经营范围                                                     ( 经营范围字典                                             )
	BusinessScope *int `json:"business_scope"`
	// 车辆类型                                                     ( 车辆类型字典                                             )
	VehicleType *int `json:"vehicle_type"`
	// 营运类型                                                     ( 营运类型字典                                             )
	OperatingType *int `json:"operating_type"`
	// 营运状态                                                     ( 营运状态字典                                             )
	OperatingState *int `json:"operating_state"`
	// 营运线路
	OperatingRoute *string `json:"operating_route"`
	// 终端ID
	TerminalID *string `json:"terminal_id"`
	// 是否申请安装智能终端
	IsApplyInstallTerminal *bool `json:"is_apply_install_terminal"`
	// 车牌号
	LicensePlateNumber *string `json:"license_plate_number"`
	// 车牌颜色                                                     ( 车牌颜色字典                                             )
	LicensePlateColor *int `json:"license_plate_color"`
	// 号牌种类                                                     ( 号牌种类字典                                             )
	LicensePlateType *int `json:"license_plate_type"`
	// 车架号(后6位)                                                ( 车辆识别代号vin,如D02133                                     )
	VehicleIdentificationNumber *string `json:"vehicle_identification_number"`
	// 道路运输证号
	RoadTransportLicenseNumber *string `json:"road_transport_license_number"`
	// 吨位
	Heavy *float64 `json:"heavy"`
	// 座位
	Seats *int `json:"seats"`
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
	// 报废日期（六合一）
	RetirementDate *time.Time `json:"retirement_date"`
	// 使用性质（六合一）
	UseNature *string `json:"use_nature"`
	// 机动车状态                                                   ( 车辆状态字典                                             )
	VehicleState *int `json:"vehicle_state"`
	// 内网更新时间
	UpdateTimeIn *time.Time `json:"update_time_in"`
	// 车辆信息同步内网反馈信息                                     ( 车辆信息同步到公安内网后内网的反馈内容，如车牌号填写错误会反馈车辆号牌错误 )
	RemarkIn *string `json:"remark_in"`
	// 是否完成                                                     ( 用于标志车辆资料是否处于确定状态。未确定状态的车辆信息在系统上除车辆管理外的功能中都查不到 )
	IsComplete *bool `json:"is_complete"`
	// 行驶证照片,云储存系统返回的路径
	DrivingLicenseePic *string `json:"driving_licensee_pic"`
	// 是否激活
	IsActive *bool `json:"is_active"`
	// 是否录入完成
	IsInput *bool `json:"is_input"`
	// 租车标准价格
	CarRentalPrice *float64 `json:"car_rental_price"`
	// 投保公司                                                     ( 投保公司字典                                             )
	InsuranceCompany *int `json:"insurance_company"`
	// 投保日期
	InsuranceDate *time.Time `json:"insurance_date"`
	// 维保数据数组，字段包括: 1.maintenance_ date 维保时间<br />2.maintenance_ kilometers 维保公里数
	VehicleMaintenances *string `json:"vehicle_maintenances"`
	// 汽车排量
	VehicleDisplacement *string `json:"vehicle_displacement"`
	// 车辆品牌                                                     ( 车辆品牌字典                                             )
	VehicleBrand *int `json:"vehicle_brand"`
	// 准驾车型                                                     ( 准驾车型字典                                             )
	QuasiDrivingModels *int `json:"quasi_driving_models"`
	// 是否上传省厅
	IsUploadProvince *bool `json:"is_upload_province"`
	// 校验状态                                                     ( 车辆校验状态字典                                         )
	CheckState *int `json:"check_state"`
	// 是否导入                                                     ( 是否通过外部导入的车辆信息                                   )
	IsImport *bool `json:"is_import"`
	// 是否工程运输车
	IsEngineeringVehicle *bool `json:"is_engineering_vehicle"`
	// 是否目录库
	IsCatalogLibrary *bool `json:"is_catalog_library"`
	// 备注
	Remarks *string `json:"remarks"`
	// 是否删除
	IsDeleted *bool `json:"is_deleted"`
	// 登记时间
	RecordAt *time.Time `json:"record_at"`
	// 登记人                                                       ( system_user表的user_id                                   )
	RecordBy *string `json:"record_by"`
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 创建人                                                       ( system_user表的user_id                                   )
	CreatedBy string `json:"created_by"`
	// 修改时间
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人                                                       ( system_user表的user_id                                   )
	UpdatedBy *string `json:"updated_by"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人                                                       ( system_user表的user_id                                   )
	DeletedBy *string `json:"deleted_by"`
}
