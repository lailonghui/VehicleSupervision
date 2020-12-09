// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"VehicleSupervision/internal/modules/vehiclelocation/last/model"
	model1 "VehicleSupervision/pkg/graphql/model"
	"fmt"
	"io"
	"strconv"
	"time"
)

// subscription root
type SubscriptionRoot struct {
	// fetch data from the table: "vehicle_location_last"
	VehicleLocationLast []*model.VehicleLocationLast `json:"vehicle_location_last"`
	// fetch aggregated fields from the table: "vehicle_location_last"
	VehicleLocationLastAggregate *VehicleLocationLastAggregate `json:"vehicle_location_last_aggregate"`
	// fetch data from the table: "vehicle_location_last" using primary key columns
	VehicleLocationLastByPk *model.VehicleLocationLast `json:"vehicle_location_last_by_pk"`
}

// aggregated selection of "vehicle_location_last"
type VehicleLocationLastAggregate struct {
	Aggregate *VehicleLocationLastAggregateFields `json:"aggregate"`
	Nodes     []*model.VehicleLocationLast        `json:"nodes"`
}

// aggregate fields of "vehicle_location_last"
type VehicleLocationLastAggregateFields struct {
	Avg        *VehicleLocationLastAvgFields        `json:"avg"`
	Count      *int                                 `json:"count"`
	Max        *VehicleLocationLastMaxFields        `json:"max"`
	Min        *VehicleLocationLastMinFields        `json:"min"`
	Stddev     *VehicleLocationLastStddevFields     `json:"stddev"`
	StddevPop  *VehicleLocationLastStddevPopFields  `json:"stddev_pop"`
	StddevSamp *VehicleLocationLastStddevSampFields `json:"stddev_samp"`
	Sum        *VehicleLocationLastSumFields        `json:"sum"`
	VarPop     *VehicleLocationLastVarPopFields     `json:"var_pop"`
	VarSamp    *VehicleLocationLastVarSampFields    `json:"var_samp"`
	Variance   *VehicleLocationLastVarianceFields   `json:"variance"`
}

// order by aggregate values of table "vehicle_location_last"
type VehicleLocationLastAggregateOrderBy struct {
	Avg        *VehicleLocationLastAvgOrderBy        `json:"avg"`
	Count      *model1.OrderBy                       `json:"count"`
	Max        *VehicleLocationLastMaxOrderBy        `json:"max"`
	Min        *VehicleLocationLastMinOrderBy        `json:"min"`
	Stddev     *VehicleLocationLastStddevOrderBy     `json:"stddev"`
	StddevPop  *VehicleLocationLastStddevPopOrderBy  `json:"stddev_pop"`
	StddevSamp *VehicleLocationLastStddevSampOrderBy `json:"stddev_samp"`
	Sum        *VehicleLocationLastSumOrderBy        `json:"sum"`
	VarPop     *VehicleLocationLastVarPopOrderBy     `json:"var_pop"`
	VarSamp    *VehicleLocationLastVarSampOrderBy    `json:"var_samp"`
	Variance   *VehicleLocationLastVarianceOrderBy   `json:"variance"`
}

// input type for inserting array relation for remote table "vehicle_location_last"
type VehicleLocationLastArrRelInsertInput struct {
	Data       []*VehicleLocationLastInsertInput `json:"data"`
	OnConflict *VehicleLocationLastOnConflict    `json:"on_conflict"`
}

// aggregate avg on columns
type VehicleLocationLastAvgFields struct {
	GpsSpeed            *float64 `json:"gps_speed"`
	ID                  *float64 `json:"id"`
	SpeedLimitThreshold *float64 `json:"speed_limit_threshold"`
	TachographSpeed     *float64 `json:"tachograph_speed"`
}

// order by avg() on columns of table "vehicle_location_last"
type VehicleLocationLastAvgOrderBy struct {
	GpsSpeed            *model1.OrderBy `json:"gps_speed"`
	ID                  *model1.OrderBy `json:"id"`
	SpeedLimitThreshold *model1.OrderBy `json:"speed_limit_threshold"`
	TachographSpeed     *model1.OrderBy `json:"tachograph_speed"`
}

// Boolean expression to filter rows from the table "vehicle_location_last". All fields are combined with a logical 'AND'.
type VehicleLocationLastBoolExp struct {
	And                 []*VehicleLocationLastBoolExp    `json:"_and"`
	Not                 *VehicleLocationLastBoolExp      `json:"_not"`
	Or                  []*VehicleLocationLastBoolExp    `json:"_or"`
	Acceleration        *model1.StringComparisonExp      `json:"acceleration"`
	AlarmContent        *model1.StringComparisonExp      `json:"alarm_content"`
	Alititude           *model1.StringComparisonExp      `json:"alititude"`
	Coordinate          *model1.PointComparisonExp       `json:"coordinate"`
	CorrectCoordinate   *model1.PointComparisonExp       `json:"correct_coordinate"`
	Direction           *model1.StringComparisonExp      `json:"direction"`
	DistrictID          *model1.StringComparisonExp      `json:"district_id"`
	DriverID            *model1.StringComparisonExp      `json:"driver_id"`
	EnterpriseID        *model1.StringComparisonExp      `json:"enterprise_id"`
	GpsSpeed            *model1.NumericComparisonExp     `json:"gps_speed"`
	ID                  *model1.BigintComparisonExp      `json:"id"`
	Imei                *model1.StringComparisonExp      `json:"imei"`
	IsLocate            *model1.BooleanComparisonExp     `json:"is_locate"`
	LocateTime          *model1.TimestamptzComparisonExp `json:"locate_time"`
	LocationDescription *model1.StringComparisonExp      `json:"location_description"`
	Mileage             *model1.StringComparisonExp      `json:"mileage"`
	RoadName            *model1.StringComparisonExp      `json:"road_name"`
	SimNumber           *model1.StringComparisonExp      `json:"sim_number"`
	SpeedLimitThreshold *model1.NumericComparisonExp     `json:"speed_limit_threshold"`
	StarCount           *model1.StringComparisonExp      `json:"star_count"`
	StarStatus          *model1.StringComparisonExp      `json:"star_status"`
	SupervisionPhotoID  *model1.StringComparisonExp      `json:"supervision_photo_id"`
	TachographSpeed     *model1.NumericComparisonExp     `json:"tachograph_speed"`
	VehicleID           *model1.StringComparisonExp      `json:"vehicle_id"`
	VehicleStatus       *model1.StringComparisonExp      `json:"vehicle_status"`
}

// input type for incrementing integer column in table "vehicle_location_last"
type VehicleLocationLastIncInput struct {
	GpsSpeed            *float64 `json:"gps_speed"`
	ID                  *int64   `json:"id"`
	SpeedLimitThreshold *float64 `json:"speed_limit_threshold"`
	TachographSpeed     *float64 `json:"tachograph_speed"`
}

// input type for inserting data into table "vehicle_location_last"
type VehicleLocationLastInsertInput struct {
	Acceleration        *string    `json:"acceleration"`
	AlarmContent        *string    `json:"alarm_content"`
	Alititude           *string    `json:"alititude"`
	Coordinate          *string    `json:"coordinate"`
	CorrectCoordinate   *string    `json:"correct_coordinate"`
	Direction           *string    `json:"direction"`
	DistrictID          *string    `json:"district_id"`
	DriverID            *string    `json:"driver_id"`
	EnterpriseID        *string    `json:"enterprise_id"`
	GpsSpeed            *float64   `json:"gps_speed"`
	ID                  *int64     `json:"id"`
	Imei                *string    `json:"imei"`
	IsLocate            *bool      `json:"is_locate"`
	LocateTime          *time.Time `json:"locate_time"`
	LocationDescription *string    `json:"location_description"`
	Mileage             *string    `json:"mileage"`
	RoadName            *string    `json:"road_name"`
	SimNumber           *string    `json:"sim_number"`
	SpeedLimitThreshold *float64   `json:"speed_limit_threshold"`
	StarCount           *string    `json:"star_count"`
	StarStatus          *string    `json:"star_status"`
	SupervisionPhotoID  *string    `json:"supervision_photo_id"`
	TachographSpeed     *float64   `json:"tachograph_speed"`
	VehicleID           *string    `json:"vehicle_id"`
	VehicleStatus       *string    `json:"vehicle_status"`
}

// aggregate max on columns
type VehicleLocationLastMaxFields struct {
	Acceleration        *string    `json:"acceleration"`
	AlarmContent        *string    `json:"alarm_content"`
	Alititude           *string    `json:"alititude"`
	Direction           *string    `json:"direction"`
	DistrictID          *string    `json:"district_id"`
	DriverID            *string    `json:"driver_id"`
	EnterpriseID        *string    `json:"enterprise_id"`
	GpsSpeed            *float64   `json:"gps_speed"`
	ID                  *int64     `json:"id"`
	Imei                *string    `json:"imei"`
	LocateTime          *time.Time `json:"locate_time"`
	LocationDescription *string    `json:"location_description"`
	Mileage             *string    `json:"mileage"`
	RoadName            *string    `json:"road_name"`
	SimNumber           *string    `json:"sim_number"`
	SpeedLimitThreshold *float64   `json:"speed_limit_threshold"`
	StarCount           *string    `json:"star_count"`
	StarStatus          *string    `json:"star_status"`
	SupervisionPhotoID  *string    `json:"supervision_photo_id"`
	TachographSpeed     *float64   `json:"tachograph_speed"`
	VehicleID           *string    `json:"vehicle_id"`
	VehicleStatus       *string    `json:"vehicle_status"`
}

// order by max() on columns of table "vehicle_location_last"
type VehicleLocationLastMaxOrderBy struct {
	Acceleration        *model1.OrderBy `json:"acceleration"`
	AlarmContent        *model1.OrderBy `json:"alarm_content"`
	Alititude           *model1.OrderBy `json:"alititude"`
	Direction           *model1.OrderBy `json:"direction"`
	DistrictID          *model1.OrderBy `json:"district_id"`
	DriverID            *model1.OrderBy `json:"driver_id"`
	EnterpriseID        *model1.OrderBy `json:"enterprise_id"`
	GpsSpeed            *model1.OrderBy `json:"gps_speed"`
	ID                  *model1.OrderBy `json:"id"`
	Imei                *model1.OrderBy `json:"imei"`
	LocateTime          *model1.OrderBy `json:"locate_time"`
	LocationDescription *model1.OrderBy `json:"location_description"`
	Mileage             *model1.OrderBy `json:"mileage"`
	RoadName            *model1.OrderBy `json:"road_name"`
	SimNumber           *model1.OrderBy `json:"sim_number"`
	SpeedLimitThreshold *model1.OrderBy `json:"speed_limit_threshold"`
	StarCount           *model1.OrderBy `json:"star_count"`
	StarStatus          *model1.OrderBy `json:"star_status"`
	SupervisionPhotoID  *model1.OrderBy `json:"supervision_photo_id"`
	TachographSpeed     *model1.OrderBy `json:"tachograph_speed"`
	VehicleID           *model1.OrderBy `json:"vehicle_id"`
	VehicleStatus       *model1.OrderBy `json:"vehicle_status"`
}

// aggregate min on columns
type VehicleLocationLastMinFields struct {
	Acceleration        *string    `json:"acceleration"`
	AlarmContent        *string    `json:"alarm_content"`
	Alititude           *string    `json:"alititude"`
	Direction           *string    `json:"direction"`
	DistrictID          *string    `json:"district_id"`
	DriverID            *string    `json:"driver_id"`
	EnterpriseID        *string    `json:"enterprise_id"`
	GpsSpeed            *float64   `json:"gps_speed"`
	ID                  *int64     `json:"id"`
	Imei                *string    `json:"imei"`
	LocateTime          *time.Time `json:"locate_time"`
	LocationDescription *string    `json:"location_description"`
	Mileage             *string    `json:"mileage"`
	RoadName            *string    `json:"road_name"`
	SimNumber           *string    `json:"sim_number"`
	SpeedLimitThreshold *float64   `json:"speed_limit_threshold"`
	StarCount           *string    `json:"star_count"`
	StarStatus          *string    `json:"star_status"`
	SupervisionPhotoID  *string    `json:"supervision_photo_id"`
	TachographSpeed     *float64   `json:"tachograph_speed"`
	VehicleID           *string    `json:"vehicle_id"`
	VehicleStatus       *string    `json:"vehicle_status"`
}

// order by min() on columns of table "vehicle_location_last"
type VehicleLocationLastMinOrderBy struct {
	Acceleration        *model1.OrderBy `json:"acceleration"`
	AlarmContent        *model1.OrderBy `json:"alarm_content"`
	Alititude           *model1.OrderBy `json:"alititude"`
	Direction           *model1.OrderBy `json:"direction"`
	DistrictID          *model1.OrderBy `json:"district_id"`
	DriverID            *model1.OrderBy `json:"driver_id"`
	EnterpriseID        *model1.OrderBy `json:"enterprise_id"`
	GpsSpeed            *model1.OrderBy `json:"gps_speed"`
	ID                  *model1.OrderBy `json:"id"`
	Imei                *model1.OrderBy `json:"imei"`
	LocateTime          *model1.OrderBy `json:"locate_time"`
	LocationDescription *model1.OrderBy `json:"location_description"`
	Mileage             *model1.OrderBy `json:"mileage"`
	RoadName            *model1.OrderBy `json:"road_name"`
	SimNumber           *model1.OrderBy `json:"sim_number"`
	SpeedLimitThreshold *model1.OrderBy `json:"speed_limit_threshold"`
	StarCount           *model1.OrderBy `json:"star_count"`
	StarStatus          *model1.OrderBy `json:"star_status"`
	SupervisionPhotoID  *model1.OrderBy `json:"supervision_photo_id"`
	TachographSpeed     *model1.OrderBy `json:"tachograph_speed"`
	VehicleID           *model1.OrderBy `json:"vehicle_id"`
	VehicleStatus       *model1.OrderBy `json:"vehicle_status"`
}

// response of any mutation on the table "vehicle_location_last"
type VehicleLocationLastMutationResponse struct {
	// number of affected rows by the mutation
	AffectedRows int `json:"affected_rows"`
	// data of the affected rows by the mutation
	Returning []*model.VehicleLocationLast `json:"returning"`
}

// input type for inserting object relation for remote table "vehicle_location_last"
type VehicleLocationLastObjRelInsertInput struct {
	Data       *VehicleLocationLastInsertInput `json:"data"`
	OnConflict *VehicleLocationLastOnConflict  `json:"on_conflict"`
}

// on conflict condition type for table "vehicle_location_last"
type VehicleLocationLastOnConflict struct {
	Constraint    VehicleLocationLastConstraint     `json:"constraint"`
	UpdateColumns []VehicleLocationLastUpdateColumn `json:"update_columns"`
	Where         *VehicleLocationLastBoolExp       `json:"where"`
}

// ordering options when selecting data from "vehicle_location_last"
type VehicleLocationLastOrderBy struct {
	Acceleration        *model1.OrderBy `json:"acceleration"`
	AlarmContent        *model1.OrderBy `json:"alarm_content"`
	Alititude           *model1.OrderBy `json:"alititude"`
	Coordinate          *model1.OrderBy `json:"coordinate"`
	CorrectCoordinate   *model1.OrderBy `json:"correct_coordinate"`
	Direction           *model1.OrderBy `json:"direction"`
	DistrictID          *model1.OrderBy `json:"district_id"`
	DriverID            *model1.OrderBy `json:"driver_id"`
	EnterpriseID        *model1.OrderBy `json:"enterprise_id"`
	GpsSpeed            *model1.OrderBy `json:"gps_speed"`
	ID                  *model1.OrderBy `json:"id"`
	Imei                *model1.OrderBy `json:"imei"`
	IsLocate            *model1.OrderBy `json:"is_locate"`
	LocateTime          *model1.OrderBy `json:"locate_time"`
	LocationDescription *model1.OrderBy `json:"location_description"`
	Mileage             *model1.OrderBy `json:"mileage"`
	RoadName            *model1.OrderBy `json:"road_name"`
	SimNumber           *model1.OrderBy `json:"sim_number"`
	SpeedLimitThreshold *model1.OrderBy `json:"speed_limit_threshold"`
	StarCount           *model1.OrderBy `json:"star_count"`
	StarStatus          *model1.OrderBy `json:"star_status"`
	SupervisionPhotoID  *model1.OrderBy `json:"supervision_photo_id"`
	TachographSpeed     *model1.OrderBy `json:"tachograph_speed"`
	VehicleID           *model1.OrderBy `json:"vehicle_id"`
	VehicleStatus       *model1.OrderBy `json:"vehicle_status"`
}

// primary key columns input for table: "vehicle_location_last"
type VehicleLocationLastPkColumnsInput struct {
	// ID
	ID int64 `json:"id"`
}

// input type for updating data in table "vehicle_location_last"
type VehicleLocationLastSetInput struct {
	Acceleration        *string    `json:"acceleration"`
	AlarmContent        *string    `json:"alarm_content"`
	Alititude           *string    `json:"alititude"`
	Coordinate          *string    `json:"coordinate"`
	CorrectCoordinate   *string    `json:"correct_coordinate"`
	Direction           *string    `json:"direction"`
	DistrictID          *string    `json:"district_id"`
	DriverID            *string    `json:"driver_id"`
	EnterpriseID        *string    `json:"enterprise_id"`
	GpsSpeed            *float64   `json:"gps_speed"`
	ID                  *int64     `json:"id"`
	Imei                *string    `json:"imei"`
	IsLocate            *bool      `json:"is_locate"`
	LocateTime          *time.Time `json:"locate_time"`
	LocationDescription *string    `json:"location_description"`
	Mileage             *string    `json:"mileage"`
	RoadName            *string    `json:"road_name"`
	SimNumber           *string    `json:"sim_number"`
	SpeedLimitThreshold *float64   `json:"speed_limit_threshold"`
	StarCount           *string    `json:"star_count"`
	StarStatus          *string    `json:"star_status"`
	SupervisionPhotoID  *string    `json:"supervision_photo_id"`
	TachographSpeed     *float64   `json:"tachograph_speed"`
	VehicleID           *string    `json:"vehicle_id"`
	VehicleStatus       *string    `json:"vehicle_status"`
}

// aggregate stddev on columns
type VehicleLocationLastStddevFields struct {
	GpsSpeed            *float64 `json:"gps_speed"`
	ID                  *float64 `json:"id"`
	SpeedLimitThreshold *float64 `json:"speed_limit_threshold"`
	TachographSpeed     *float64 `json:"tachograph_speed"`
}

// order by stddev() on columns of table "vehicle_location_last"
type VehicleLocationLastStddevOrderBy struct {
	GpsSpeed            *model1.OrderBy `json:"gps_speed"`
	ID                  *model1.OrderBy `json:"id"`
	SpeedLimitThreshold *model1.OrderBy `json:"speed_limit_threshold"`
	TachographSpeed     *model1.OrderBy `json:"tachograph_speed"`
}

// aggregate stddev_pop on columns
type VehicleLocationLastStddevPopFields struct {
	GpsSpeed            *float64 `json:"gps_speed"`
	ID                  *float64 `json:"id"`
	SpeedLimitThreshold *float64 `json:"speed_limit_threshold"`
	TachographSpeed     *float64 `json:"tachograph_speed"`
}

// order by stddev_pop() on columns of table "vehicle_location_last"
type VehicleLocationLastStddevPopOrderBy struct {
	GpsSpeed            *model1.OrderBy `json:"gps_speed"`
	ID                  *model1.OrderBy `json:"id"`
	SpeedLimitThreshold *model1.OrderBy `json:"speed_limit_threshold"`
	TachographSpeed     *model1.OrderBy `json:"tachograph_speed"`
}

// aggregate stddev_samp on columns
type VehicleLocationLastStddevSampFields struct {
	GpsSpeed            *float64 `json:"gps_speed"`
	ID                  *float64 `json:"id"`
	SpeedLimitThreshold *float64 `json:"speed_limit_threshold"`
	TachographSpeed     *float64 `json:"tachograph_speed"`
}

// order by stddev_samp() on columns of table "vehicle_location_last"
type VehicleLocationLastStddevSampOrderBy struct {
	GpsSpeed            *model1.OrderBy `json:"gps_speed"`
	ID                  *model1.OrderBy `json:"id"`
	SpeedLimitThreshold *model1.OrderBy `json:"speed_limit_threshold"`
	TachographSpeed     *model1.OrderBy `json:"tachograph_speed"`
}

// aggregate sum on columns
type VehicleLocationLastSumFields struct {
	GpsSpeed            *float64 `json:"gps_speed"`
	ID                  *int64   `json:"id"`
	SpeedLimitThreshold *float64 `json:"speed_limit_threshold"`
	TachographSpeed     *float64 `json:"tachograph_speed"`
}

// order by sum() on columns of table "vehicle_location_last"
type VehicleLocationLastSumOrderBy struct {
	GpsSpeed            *model1.OrderBy `json:"gps_speed"`
	ID                  *model1.OrderBy `json:"id"`
	SpeedLimitThreshold *model1.OrderBy `json:"speed_limit_threshold"`
	TachographSpeed     *model1.OrderBy `json:"tachograph_speed"`
}

// aggregate var_pop on columns
type VehicleLocationLastVarPopFields struct {
	GpsSpeed            *float64 `json:"gps_speed"`
	ID                  *float64 `json:"id"`
	SpeedLimitThreshold *float64 `json:"speed_limit_threshold"`
	TachographSpeed     *float64 `json:"tachograph_speed"`
}

// order by var_pop() on columns of table "vehicle_location_last"
type VehicleLocationLastVarPopOrderBy struct {
	GpsSpeed            *model1.OrderBy `json:"gps_speed"`
	ID                  *model1.OrderBy `json:"id"`
	SpeedLimitThreshold *model1.OrderBy `json:"speed_limit_threshold"`
	TachographSpeed     *model1.OrderBy `json:"tachograph_speed"`
}

// aggregate var_samp on columns
type VehicleLocationLastVarSampFields struct {
	GpsSpeed            *float64 `json:"gps_speed"`
	ID                  *float64 `json:"id"`
	SpeedLimitThreshold *float64 `json:"speed_limit_threshold"`
	TachographSpeed     *float64 `json:"tachograph_speed"`
}

// order by var_samp() on columns of table "vehicle_location_last"
type VehicleLocationLastVarSampOrderBy struct {
	GpsSpeed            *model1.OrderBy `json:"gps_speed"`
	ID                  *model1.OrderBy `json:"id"`
	SpeedLimitThreshold *model1.OrderBy `json:"speed_limit_threshold"`
	TachographSpeed     *model1.OrderBy `json:"tachograph_speed"`
}

// aggregate variance on columns
type VehicleLocationLastVarianceFields struct {
	GpsSpeed            *float64 `json:"gps_speed"`
	ID                  *float64 `json:"id"`
	SpeedLimitThreshold *float64 `json:"speed_limit_threshold"`
	TachographSpeed     *float64 `json:"tachograph_speed"`
}

// order by variance() on columns of table "vehicle_location_last"
type VehicleLocationLastVarianceOrderBy struct {
	GpsSpeed            *model1.OrderBy `json:"gps_speed"`
	ID                  *model1.OrderBy `json:"id"`
	SpeedLimitThreshold *model1.OrderBy `json:"speed_limit_threshold"`
	TachographSpeed     *model1.OrderBy `json:"tachograph_speed"`
}

// unique or primary key constraints on table "vehicle_location_last"
type VehicleLocationLastConstraint string

const (
	// unique or primary key constraint
	VehicleLocationLastConstraintVehicleLocationLastPkey VehicleLocationLastConstraint = "vehicle_location_last_pkey"
)

var AllVehicleLocationLastConstraint = []VehicleLocationLastConstraint{
	VehicleLocationLastConstraintVehicleLocationLastPkey,
}

func (e VehicleLocationLastConstraint) IsValid() bool {
	switch e {
	case VehicleLocationLastConstraintVehicleLocationLastPkey:
		return true
	}
	return false
}

func (e VehicleLocationLastConstraint) String() string {
	return string(e)
}

func (e *VehicleLocationLastConstraint) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = VehicleLocationLastConstraint(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid vehicle_location_last_constraint", str)
	}
	return nil
}

func (e VehicleLocationLastConstraint) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// select columns of table "vehicle_location_last"
type VehicleLocationLastSelectColumn string

const (
	// column name
	VehicleLocationLastSelectColumnAcceleration VehicleLocationLastSelectColumn = "acceleration"
	// column name
	VehicleLocationLastSelectColumnAlarmContent VehicleLocationLastSelectColumn = "alarm_content"
	// column name
	VehicleLocationLastSelectColumnAlititude VehicleLocationLastSelectColumn = "alititude"
	// column name
	VehicleLocationLastSelectColumnCoordinate VehicleLocationLastSelectColumn = "coordinate"
	// column name
	VehicleLocationLastSelectColumnCorrectCoordinate VehicleLocationLastSelectColumn = "correct_coordinate"
	// column name
	VehicleLocationLastSelectColumnDirection VehicleLocationLastSelectColumn = "direction"
	// column name
	VehicleLocationLastSelectColumnDistrictID VehicleLocationLastSelectColumn = "district_id"
	// column name
	VehicleLocationLastSelectColumnDriverID VehicleLocationLastSelectColumn = "driver_id"
	// column name
	VehicleLocationLastSelectColumnEnterpriseID VehicleLocationLastSelectColumn = "enterprise_id"
	// column name
	VehicleLocationLastSelectColumnGpsSpeed VehicleLocationLastSelectColumn = "gps_speed"
	// column name
	VehicleLocationLastSelectColumnID VehicleLocationLastSelectColumn = "id"
	// column name
	VehicleLocationLastSelectColumnImei VehicleLocationLastSelectColumn = "imei"
	// column name
	VehicleLocationLastSelectColumnIsLocate VehicleLocationLastSelectColumn = "is_locate"
	// column name
	VehicleLocationLastSelectColumnLocateTime VehicleLocationLastSelectColumn = "locate_time"
	// column name
	VehicleLocationLastSelectColumnLocationDescription VehicleLocationLastSelectColumn = "location_description"
	// column name
	VehicleLocationLastSelectColumnMileage VehicleLocationLastSelectColumn = "mileage"
	// column name
	VehicleLocationLastSelectColumnRoadName VehicleLocationLastSelectColumn = "road_name"
	// column name
	VehicleLocationLastSelectColumnSimNumber VehicleLocationLastSelectColumn = "sim_number"
	// column name
	VehicleLocationLastSelectColumnSpeedLimitThreshold VehicleLocationLastSelectColumn = "speed_limit_threshold"
	// column name
	VehicleLocationLastSelectColumnStarCount VehicleLocationLastSelectColumn = "star_count"
	// column name
	VehicleLocationLastSelectColumnStarStatus VehicleLocationLastSelectColumn = "star_status"
	// column name
	VehicleLocationLastSelectColumnSupervisionPhotoID VehicleLocationLastSelectColumn = "supervision_photo_id"
	// column name
	VehicleLocationLastSelectColumnTachographSpeed VehicleLocationLastSelectColumn = "tachograph_speed"
	// column name
	VehicleLocationLastSelectColumnVehicleID VehicleLocationLastSelectColumn = "vehicle_id"
	// column name
	VehicleLocationLastSelectColumnVehicleStatus VehicleLocationLastSelectColumn = "vehicle_status"
)

var AllVehicleLocationLastSelectColumn = []VehicleLocationLastSelectColumn{
	VehicleLocationLastSelectColumnAcceleration,
	VehicleLocationLastSelectColumnAlarmContent,
	VehicleLocationLastSelectColumnAlititude,
	VehicleLocationLastSelectColumnCoordinate,
	VehicleLocationLastSelectColumnCorrectCoordinate,
	VehicleLocationLastSelectColumnDirection,
	VehicleLocationLastSelectColumnDistrictID,
	VehicleLocationLastSelectColumnDriverID,
	VehicleLocationLastSelectColumnEnterpriseID,
	VehicleLocationLastSelectColumnGpsSpeed,
	VehicleLocationLastSelectColumnID,
	VehicleLocationLastSelectColumnImei,
	VehicleLocationLastSelectColumnIsLocate,
	VehicleLocationLastSelectColumnLocateTime,
	VehicleLocationLastSelectColumnLocationDescription,
	VehicleLocationLastSelectColumnMileage,
	VehicleLocationLastSelectColumnRoadName,
	VehicleLocationLastSelectColumnSimNumber,
	VehicleLocationLastSelectColumnSpeedLimitThreshold,
	VehicleLocationLastSelectColumnStarCount,
	VehicleLocationLastSelectColumnStarStatus,
	VehicleLocationLastSelectColumnSupervisionPhotoID,
	VehicleLocationLastSelectColumnTachographSpeed,
	VehicleLocationLastSelectColumnVehicleID,
	VehicleLocationLastSelectColumnVehicleStatus,
}

func (e VehicleLocationLastSelectColumn) IsValid() bool {
	switch e {
	case VehicleLocationLastSelectColumnAcceleration, VehicleLocationLastSelectColumnAlarmContent, VehicleLocationLastSelectColumnAlititude, VehicleLocationLastSelectColumnCoordinate, VehicleLocationLastSelectColumnCorrectCoordinate, VehicleLocationLastSelectColumnDirection, VehicleLocationLastSelectColumnDistrictID, VehicleLocationLastSelectColumnDriverID, VehicleLocationLastSelectColumnEnterpriseID, VehicleLocationLastSelectColumnGpsSpeed, VehicleLocationLastSelectColumnID, VehicleLocationLastSelectColumnImei, VehicleLocationLastSelectColumnIsLocate, VehicleLocationLastSelectColumnLocateTime, VehicleLocationLastSelectColumnLocationDescription, VehicleLocationLastSelectColumnMileage, VehicleLocationLastSelectColumnRoadName, VehicleLocationLastSelectColumnSimNumber, VehicleLocationLastSelectColumnSpeedLimitThreshold, VehicleLocationLastSelectColumnStarCount, VehicleLocationLastSelectColumnStarStatus, VehicleLocationLastSelectColumnSupervisionPhotoID, VehicleLocationLastSelectColumnTachographSpeed, VehicleLocationLastSelectColumnVehicleID, VehicleLocationLastSelectColumnVehicleStatus:
		return true
	}
	return false
}

func (e VehicleLocationLastSelectColumn) String() string {
	return string(e)
}

func (e *VehicleLocationLastSelectColumn) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = VehicleLocationLastSelectColumn(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid vehicle_location_last_select_column", str)
	}
	return nil
}

func (e VehicleLocationLastSelectColumn) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// update columns of table "vehicle_location_last"
type VehicleLocationLastUpdateColumn string

const (
	// column name
	VehicleLocationLastUpdateColumnAcceleration VehicleLocationLastUpdateColumn = "acceleration"
	// column name
	VehicleLocationLastUpdateColumnAlarmContent VehicleLocationLastUpdateColumn = "alarm_content"
	// column name
	VehicleLocationLastUpdateColumnAlititude VehicleLocationLastUpdateColumn = "alititude"
	// column name
	VehicleLocationLastUpdateColumnCoordinate VehicleLocationLastUpdateColumn = "coordinate"
	// column name
	VehicleLocationLastUpdateColumnCorrectCoordinate VehicleLocationLastUpdateColumn = "correct_coordinate"
	// column name
	VehicleLocationLastUpdateColumnDirection VehicleLocationLastUpdateColumn = "direction"
	// column name
	VehicleLocationLastUpdateColumnDistrictID VehicleLocationLastUpdateColumn = "district_id"
	// column name
	VehicleLocationLastUpdateColumnDriverID VehicleLocationLastUpdateColumn = "driver_id"
	// column name
	VehicleLocationLastUpdateColumnEnterpriseID VehicleLocationLastUpdateColumn = "enterprise_id"
	// column name
	VehicleLocationLastUpdateColumnGpsSpeed VehicleLocationLastUpdateColumn = "gps_speed"
	// column name
	VehicleLocationLastUpdateColumnID VehicleLocationLastUpdateColumn = "id"
	// column name
	VehicleLocationLastUpdateColumnImei VehicleLocationLastUpdateColumn = "imei"
	// column name
	VehicleLocationLastUpdateColumnIsLocate VehicleLocationLastUpdateColumn = "is_locate"
	// column name
	VehicleLocationLastUpdateColumnLocateTime VehicleLocationLastUpdateColumn = "locate_time"
	// column name
	VehicleLocationLastUpdateColumnLocationDescription VehicleLocationLastUpdateColumn = "location_description"
	// column name
	VehicleLocationLastUpdateColumnMileage VehicleLocationLastUpdateColumn = "mileage"
	// column name
	VehicleLocationLastUpdateColumnRoadName VehicleLocationLastUpdateColumn = "road_name"
	// column name
	VehicleLocationLastUpdateColumnSimNumber VehicleLocationLastUpdateColumn = "sim_number"
	// column name
	VehicleLocationLastUpdateColumnSpeedLimitThreshold VehicleLocationLastUpdateColumn = "speed_limit_threshold"
	// column name
	VehicleLocationLastUpdateColumnStarCount VehicleLocationLastUpdateColumn = "star_count"
	// column name
	VehicleLocationLastUpdateColumnStarStatus VehicleLocationLastUpdateColumn = "star_status"
	// column name
	VehicleLocationLastUpdateColumnSupervisionPhotoID VehicleLocationLastUpdateColumn = "supervision_photo_id"
	// column name
	VehicleLocationLastUpdateColumnTachographSpeed VehicleLocationLastUpdateColumn = "tachograph_speed"
	// column name
	VehicleLocationLastUpdateColumnVehicleID VehicleLocationLastUpdateColumn = "vehicle_id"
	// column name
	VehicleLocationLastUpdateColumnVehicleStatus VehicleLocationLastUpdateColumn = "vehicle_status"
)

var AllVehicleLocationLastUpdateColumn = []VehicleLocationLastUpdateColumn{
	VehicleLocationLastUpdateColumnAcceleration,
	VehicleLocationLastUpdateColumnAlarmContent,
	VehicleLocationLastUpdateColumnAlititude,
	VehicleLocationLastUpdateColumnCoordinate,
	VehicleLocationLastUpdateColumnCorrectCoordinate,
	VehicleLocationLastUpdateColumnDirection,
	VehicleLocationLastUpdateColumnDistrictID,
	VehicleLocationLastUpdateColumnDriverID,
	VehicleLocationLastUpdateColumnEnterpriseID,
	VehicleLocationLastUpdateColumnGpsSpeed,
	VehicleLocationLastUpdateColumnID,
	VehicleLocationLastUpdateColumnImei,
	VehicleLocationLastUpdateColumnIsLocate,
	VehicleLocationLastUpdateColumnLocateTime,
	VehicleLocationLastUpdateColumnLocationDescription,
	VehicleLocationLastUpdateColumnMileage,
	VehicleLocationLastUpdateColumnRoadName,
	VehicleLocationLastUpdateColumnSimNumber,
	VehicleLocationLastUpdateColumnSpeedLimitThreshold,
	VehicleLocationLastUpdateColumnStarCount,
	VehicleLocationLastUpdateColumnStarStatus,
	VehicleLocationLastUpdateColumnSupervisionPhotoID,
	VehicleLocationLastUpdateColumnTachographSpeed,
	VehicleLocationLastUpdateColumnVehicleID,
	VehicleLocationLastUpdateColumnVehicleStatus,
}

func (e VehicleLocationLastUpdateColumn) IsValid() bool {
	switch e {
	case VehicleLocationLastUpdateColumnAcceleration, VehicleLocationLastUpdateColumnAlarmContent, VehicleLocationLastUpdateColumnAlititude, VehicleLocationLastUpdateColumnCoordinate, VehicleLocationLastUpdateColumnCorrectCoordinate, VehicleLocationLastUpdateColumnDirection, VehicleLocationLastUpdateColumnDistrictID, VehicleLocationLastUpdateColumnDriverID, VehicleLocationLastUpdateColumnEnterpriseID, VehicleLocationLastUpdateColumnGpsSpeed, VehicleLocationLastUpdateColumnID, VehicleLocationLastUpdateColumnImei, VehicleLocationLastUpdateColumnIsLocate, VehicleLocationLastUpdateColumnLocateTime, VehicleLocationLastUpdateColumnLocationDescription, VehicleLocationLastUpdateColumnMileage, VehicleLocationLastUpdateColumnRoadName, VehicleLocationLastUpdateColumnSimNumber, VehicleLocationLastUpdateColumnSpeedLimitThreshold, VehicleLocationLastUpdateColumnStarCount, VehicleLocationLastUpdateColumnStarStatus, VehicleLocationLastUpdateColumnSupervisionPhotoID, VehicleLocationLastUpdateColumnTachographSpeed, VehicleLocationLastUpdateColumnVehicleID, VehicleLocationLastUpdateColumnVehicleStatus:
		return true
	}
	return false
}

func (e VehicleLocationLastUpdateColumn) String() string {
	return string(e)
}

func (e *VehicleLocationLastUpdateColumn) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = VehicleLocationLastUpdateColumn(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid vehicle_location_last_update_column", str)
	}
	return nil
}

func (e VehicleLocationLastUpdateColumn) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
