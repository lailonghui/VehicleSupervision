/*
@Time : 2020/12/9 10:49
@Author : lai
@Description :
@File : muck_truck_worker_id_card_orders
*/
package model

import "time"

// 渣土车工号牌制作订单表
//
//
// columns and relationships of "muck_truck_worker_id_card_orders"
type MuckTruckWorkerIDCardOrders struct {
	// 创建时间
	CreatedAt time.Time `json:"create_at"`
	// 创建人
	CreatedBy string `json:"create_by"`
	// 删除时间
	DeletedAt *time.Time `json:"delete_at"`
	// 删除人
	DeletedBy *string `json:"delete_by"`
	// 主键
	ID int64 `json:"id"`
	// muck_truck_preview_number  渣土车车辆预编号表的id
	PreviewNumberID int64 `json:"preview_number_id"`
	// 修改时间
	UpdatedAt *time.Time `json:"update_at"`
	// 修改人
	UpdatedBy *string `json:"update_by"`
	// vehicle_info 车辆信息表 的vehicle_id
	VehicleID string `json:"vehicle_id"`
}
