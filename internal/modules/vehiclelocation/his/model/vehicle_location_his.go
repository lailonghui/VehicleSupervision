package model

import "time"

// 车辆位置历史表
//
//
// columns and relationships of "vehicle_location_his"
type VehicleLocationHis struct {
	// 加速度
	Acceleration *string `json:"acceleration"`
	// 报警内容
	AlarmContent *string `json:"alarm_content"`
	// 海拔
	Alititude *string `json:"alititude"`
	// 坐标
	Coordinate *string `json:"coordinate"`
	// 纠偏后坐标
	CorrectCoordinate *string `json:"correct_coordinate"`
	// 方向
	Direction *string `json:"direction"`
	// 区域ID
	DistrictID *string `json:"district_id"`
	// 驾驶员ID
	DriverID *string `json:"driver_id"`
	// 企业ID
	EnterpriseID *string `json:"enterprise_id"`
	// GPS速度
	GpsSpeed *float64 `json:"gps_speed"`
	// ID
	ID int64 `json:"id"`
	// 终端IMEI
	Imei *string `json:"imei"`
	// 是否定位
	IsLocate *bool `json:"is_locate"`
	// 定位时间
	LocateTime *time.Time `json:"locate_time"`
	// 位置描述
	LocationDescription *string `json:"location_description"`
	// 里程
	Mileage *string `json:"mileage"`
	// 道路名称
	RoadName *string `json:"road_name"`
	// SIM卡号
	SimNumber *string `json:"sim_number"`
	// 限速阀值
	SpeedLimitThreshold *float64 `json:"speed_limit_threshold"`
	// 星数
	StarCount *string `json:"star_count"`
	// 星况
	StarStatus *string `json:"star_status"`
	// 监控图片ID
	SupervisionPhotoID *string `json:"supervision_photo_id"`
	// 行驶记录仪速度
	TachographSpeed *float64 `json:"tachograph_speed"`
	// 车辆ID
	VehicleID string `json:"vehicle_id"`
	// 车辆状态
	VehicleStatus *string `json:"vehicle_status"`
}

func (v VehicleLocationHis) TableName() string {
	return "vehicle_location_his"
}
