package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden TerminalLoader string *VehicleSupervision/internal/modules/device/model.Terminal

// 终端
//
//
// columns and relationships of "terminal"
type Terminal struct {
	// ADAS型号
	AdasModel *string `json:"adas_model"`
	// 外接鉴权码
	AuthKey *string `json:"auth_key"`
	// 摄像头数量
	CameraNum *int `json:"camera_num"`
	// 摄像头安装位置
	ChannelNo *string `json:"channel_no"`
	// 城市ID
	CityID *string `json:"city_id"`
	// 新建时间
	CreateAt time.Time `json:"create_at"`
	// 新建人
	CreateBy *string `json:"create_by"`
	// 删除时间
	DeleteAt *time.Time `json:"delete_at"`
	// 删除人
	DeleteBy *string `json:"delete_by"`
	// 部门ID
	DeptID *string `json:"dept_id"`
	// 首次安装时间
	FirstInstallTime *time.Time `json:"first_install_time"`
	// 保修期到期时间
	GuaranteeDate *time.Time `json:"guarantee_date"`
	// ID
	ID int64 `json:"id"`
	// 终端IMEI
	Imei string `json:"imei"`
	// 安装人员ID
	InstallManID *string `json:"install_man_id"`
	// 是否删除
	IsDelete bool `json:"is_delete"`
	// 是否注册
	IsReg *bool `json:"is_reg"`
	// 是否支持拍照
	IsSupportPhoto bool `json:"is_support_photo"`
	// 模拟鉴权码
	MockAuthKey *string `json:"mock_auth_key"`
	// 生产日期
	ProduceDate *time.Time `json:"produce_date"`
	// 终端安装证明代理商
	ProveProxyrgID *string `json:"prove_proxyrg_id"`
	// 省份ID
	ProvinceID *string `json:"province_id"`
	// 代理商ID
	ProxyrgID *string `json:"proxyrg_id"`
	// 登记时间
	RecordDate *time.Time `json:"record_date"`
	// 终端注册ID
	RegID *string `json:"reg_id"`
	// 备注
	Remarks *string `json:"remarks"`
	// 解绑原因
	RemoveReason *string `json:"remove_reason"`
	// SIM卡ID
	SimID *string `json:"sim_id"`
	// 运营商
	SprgID *string `json:"sprg_id"`
	// 终端ID
	TerminalID string `json:"terminal_id"`
	// 终端类型ID
	TypeID *string `json:"type_id"`
	// 修改时间
	UpdateAt *time.Time `json:"update_at"`
	// 修改人
	UpdateBy *string `json:"update_by"`
	// 车辆ID
	VehicleID *string `json:"vehicle_id"`
	// 版本号
	VersionNumber *string `json:"version_number"`
}

func (t Terminal) TableName() string {
	return "terminal"
}

func (u *Terminal) NewLoader() *TerminalLoader {
	return &TerminalLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*Terminal, []error) {
			var rs []*Terminal
			db.DB.Model(&Terminal{}).Where("terminal_id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}