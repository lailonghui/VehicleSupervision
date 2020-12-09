package model

import "time"

// 终端型号
//
//
// columns and relationships of "terminal_modal"
type TerminalModal struct {
	// adas型号
	AdasModal *string `json:"adas_modal"`
	// 创建时间
	CreateAt time.Time `json:"create_at"`
	// 创建人
	CreateBy *string `json:"create_by"`
	// 删除时间
	DeleteAt *time.Time `json:"delete_at"`
	// 删除人
	DeleteBy *string `json:"delete_by"`
	// 生产厂家ID
	FactoryID *string `json:"factory_id"`
	// ID
	ID int64 `json:"id"`
	// 是否被删除
	IsDelete bool `json:"is_delete"`
	// 终端具备出租汽车电召、电子上岗证等功能
	IsElectronicsPostCard *bool `json:"is_electronics_post_card"`
	// 是否渣土车终端
	IsSlagCarTeminal *bool `json:"is_slag_car_teminal"`
	// 是否提交终端技术要求通过相关检测情况
	IsTestingSituation *bool `json:"is_testing_situation"`
	// 是否交通局4G终端
	IsTransportDept4g *bool `json:"is_transport_dept_4g"`
	// 型号
	ModalName *string `json:"modal_name"`
	// 代理商ID
	ProxyrgID *string `json:"proxyrg_id"`
	// 备案编号
	RecordNo *string `json:"record_no"`
	// 备注
	Remark *string `json:"remark"`
	// 终端型号ID
	TerminalModalID string `json:"terminal_modal_id"`
	// 终端类型ID
	TerminalTypeID *string `json:"terminal_type_id"`
	// 修改时间
	UpdateAt *time.Time `json:"update_at"`
	// 修改人
	UpdateBy *string `json:"update_by"`
}

func (t TerminalModal) TableName() string {
	return "terminal_modal"
}
