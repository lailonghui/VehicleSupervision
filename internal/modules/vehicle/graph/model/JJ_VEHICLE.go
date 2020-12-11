/*
@Time : 2020/12/9 10:51
@Author : lai
@Description :
@File : JJ_VEHICLE
*/
package model

import "time"

// 公安内网六合一平台同步车辆表(不修改字段)
//
//
// columns and relationships of "JJ_VEHICLE"
type JjVehicle struct {
	// 初次登记日期
	Ccdjrq *time.Time `json:"CCDJRQ"`
	// 车辆类型
	Cllx *string `json:"CLLX"`
	// 车辆识别代号
	Clsbdh *string `json:"CLSBDH"`
	// 所在县
	County *string `json:"COUNTY"`
	// 登记日期
	Djrq *time.Time `json:"DJRQ"`
	// 发牌日期
	Fprq *time.Time `json:"FPRQ"`
	// 无
	Gxrq *time.Time `json:"GXRQ"`
	// 车牌号码
	Hphm *string `json:"HPHM"`
	// 号牌种类
	Hpzl *string `json:"HPZL"`
	// 是否删除
	IsDeteled *float64 `json:"IS_DETELED"`
	// 联系电话
	Lxdh *string `json:"LXDH"`
	// 联系地址
	Lxdz *string `json:"LXDZ"`
	// 固话
	Other *string `json:"OTHER"`
	// 强制报废期止
	Qzbfqz *time.Time `json:"QZBFQZ"`
	// 无
	Sjhm *string `json:"SJHM"`
	// 所有人
	Syr *string `json:"SYR"`
	// 使用性质
	Syxz *string `json:"SYXZ"`
	// 更新时间
	Updatetime *time.Time `json:"UPDATETIME"`
	// 经营范围
	Vehmontype *string `json:"VEHMONTYPE"`
	// 型号
	Xh *string `json:"XH"`
	// 有效期至
	Yxqz *time.Time `json:"YXQZ"`
	// 机动车状态
	Zt *int `json:"ZT"`
	// 总质量
	Zzl *float64 `json:"ZZL"`
	// 主键
	ID int64 `json:"id"`
	// vehicle_info 车辆信息表 的vehicle_id
	VehicleID string `json:"vehicle_id"`
}

func (JjVehicle) TableName() string {
	return "JJ_VEHICLE"
}
