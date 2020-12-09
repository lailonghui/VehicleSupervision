// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"VehicleSupervision/internal/modules/driving/log/model"
	model1 "VehicleSupervision/pkg/graphql/model"
	"fmt"
	"io"
	"strconv"
	"time"
)

// aggregated selection of "driving_log"
type DrivingLogAggregate struct {
	Aggregate *DrivingLogAggregateFields `json:"aggregate"`
	Nodes     []*model.DrivingLog        `json:"nodes"`
}

// aggregate fields of "driving_log"
type DrivingLogAggregateFields struct {
	Avg        *DrivingLogAvgFields        `json:"avg"`
	Count      *int                        `json:"count"`
	Max        *DrivingLogMaxFields        `json:"max"`
	Min        *DrivingLogMinFields        `json:"min"`
	Stddev     *DrivingLogStddevFields     `json:"stddev"`
	StddevPop  *DrivingLogStddevPopFields  `json:"stddev_pop"`
	StddevSamp *DrivingLogStddevSampFields `json:"stddev_samp"`
	Sum        *DrivingLogSumFields        `json:"sum"`
	VarPop     *DrivingLogVarPopFields     `json:"var_pop"`
	VarSamp    *DrivingLogVarSampFields    `json:"var_samp"`
	Variance   *DrivingLogVarianceFields   `json:"variance"`
}

// order by aggregate values of table "driving_log"
type DrivingLogAggregateOrderBy struct {
	Avg        *DrivingLogAvgOrderBy        `json:"avg"`
	Count      *model1.OrderBy              `json:"count"`
	Max        *DrivingLogMaxOrderBy        `json:"max"`
	Min        *DrivingLogMinOrderBy        `json:"min"`
	Stddev     *DrivingLogStddevOrderBy     `json:"stddev"`
	StddevPop  *DrivingLogStddevPopOrderBy  `json:"stddev_pop"`
	StddevSamp *DrivingLogStddevSampOrderBy `json:"stddev_samp"`
	Sum        *DrivingLogSumOrderBy        `json:"sum"`
	VarPop     *DrivingLogVarPopOrderBy     `json:"var_pop"`
	VarSamp    *DrivingLogVarSampOrderBy    `json:"var_samp"`
	Variance   *DrivingLogVarianceOrderBy   `json:"variance"`
}

// input type for inserting array relation for remote table "driving_log"
type DrivingLogArrRelInsertInput struct {
	Data       []*DrivingLogInsertInput `json:"data"`
	OnConflict *DrivingLogOnConflict    `json:"on_conflict"`
}

// aggregate avg on columns
type DrivingLogAvgFields struct {
	CheckOrganizationLevel *float64 `json:"check_organization_level"`
	CheckState             *float64 `json:"check_state"`
	ID                     *float64 `json:"id"`
}

// order by avg() on columns of table "driving_log"
type DrivingLogAvgOrderBy struct {
	CheckOrganizationLevel *model1.OrderBy `json:"check_organization_level"`
	CheckState             *model1.OrderBy `json:"check_state"`
	ID                     *model1.OrderBy `json:"id"`
}

// Boolean expression to filter rows from the table "driving_log". All fields are combined with a logical 'AND'.
type DrivingLogBoolExp struct {
	And                    []*DrivingLogBoolExp             `json:"_and"`
	Not                    *DrivingLogBoolExp               `json:"_not"`
	Or                     []*DrivingLogBoolExp             `json:"_or"`
	Cause                  *model1.StringComparisonExp      `json:"cause"`
	CheckOrganizationLevel *model1.IntComparisonExp         `json:"check_organization_level"`
	CheckState             *model1.IntComparisonExp         `json:"check_state"`
	CreateAt               *model1.TimestamptzComparisonExp `json:"create_at"`
	CreateBy               *model1.StringComparisonExp      `json:"create_by"`
	DeleteAt               *model1.TimestamptzComparisonExp `json:"delete_at"`
	DeleteBy               *model1.StringComparisonExp      `json:"delete_by"`
	DriverID               *model1.StringComparisonExp      `json:"driver_id"`
	DrivingEndTime         *model1.DateComparisonExp        `json:"driving_end_time"`
	DrivingLogID           *model1.StringComparisonExp      `json:"driving_log_id"`
	DrivingStartTime       *model1.DateComparisonExp        `json:"driving_start_time"`
	EndTime                *model1.TimestamptzComparisonExp `json:"end_time"`
	ID                     *model1.BigintComparisonExp      `json:"id"`
	IsDelete               *model1.BooleanComparisonExp     `json:"is_delete"`
	IsFill                 *model1.BooleanComparisonExp     `json:"is_fill"`
	RegisterAt             *model1.TimestamptzComparisonExp `json:"register_at"`
	RegisterBy             *model1.StringComparisonExp      `json:"register_by"`
	Remarks                *model1.StringComparisonExp      `json:"remarks"`
	Route                  *model1.StringComparisonExp      `json:"route"`
	StartTime              *model1.TimestamptzComparisonExp `json:"start_time"`
	UpdateAt               *model1.TimestamptzComparisonExp `json:"update_at"`
	UpdateBy               *model1.StringComparisonExp      `json:"update_by"`
	VehicleID              *model1.StringComparisonExp      `json:"vehicle_id"`
}

// input type for incrementing integer column in table "driving_log"
type DrivingLogIncInput struct {
	CheckOrganizationLevel *int   `json:"check_organization_level"`
	CheckState             *int   `json:"check_state"`
	ID                     *int64 `json:"id"`
}

// input type for inserting data into table "driving_log"
type DrivingLogInsertInput struct {
	Cause                  *string    `json:"cause"`
	CheckOrganizationLevel *int       `json:"check_organization_level"`
	CheckState             *int       `json:"check_state"`
	CreateAt               *time.Time `json:"create_at"`
	CreateBy               *string    `json:"create_by"`
	DeleteAt               *time.Time `json:"delete_at"`
	DeleteBy               *string    `json:"delete_by"`
	DriverID               *string    `json:"driver_id"`
	DrivingEndTime         *string    `json:"driving_end_time"`
	DrivingLogID           *string    `json:"driving_log_id"`
	DrivingStartTime       *string    `json:"driving_start_time"`
	EndTime                *time.Time `json:"end_time"`
	ID                     *int64     `json:"id"`
	IsDelete               *bool      `json:"is_delete"`
	IsFill                 *bool      `json:"is_fill"`
	RegisterAt             *time.Time `json:"register_at"`
	RegisterBy             *string    `json:"register_by"`
	Remarks                *string    `json:"remarks"`
	Route                  *string    `json:"route"`
	StartTime              *time.Time `json:"start_time"`
	UpdateAt               *time.Time `json:"update_at"`
	UpdateBy               *string    `json:"update_by"`
	VehicleID              *string    `json:"vehicle_id"`
}

// aggregate max on columns
type DrivingLogMaxFields struct {
	Cause                  *string    `json:"cause"`
	CheckOrganizationLevel *int       `json:"check_organization_level"`
	CheckState             *int       `json:"check_state"`
	CreateAt               *time.Time `json:"create_at"`
	CreateBy               *string    `json:"create_by"`
	DeleteAt               *time.Time `json:"delete_at"`
	DeleteBy               *string    `json:"delete_by"`
	DriverID               *string    `json:"driver_id"`
	DrivingEndTime         *string    `json:"driving_end_time"`
	DrivingLogID           *string    `json:"driving_log_id"`
	DrivingStartTime       *string    `json:"driving_start_time"`
	EndTime                *time.Time `json:"end_time"`
	ID                     *int64     `json:"id"`
	RegisterAt             *time.Time `json:"register_at"`
	RegisterBy             *string    `json:"register_by"`
	Remarks                *string    `json:"remarks"`
	Route                  *string    `json:"route"`
	StartTime              *time.Time `json:"start_time"`
	UpdateAt               *time.Time `json:"update_at"`
	UpdateBy               *string    `json:"update_by"`
	VehicleID              *string    `json:"vehicle_id"`
}

// order by max() on columns of table "driving_log"
type DrivingLogMaxOrderBy struct {
	Cause                  *model1.OrderBy `json:"cause"`
	CheckOrganizationLevel *model1.OrderBy `json:"check_organization_level"`
	CheckState             *model1.OrderBy `json:"check_state"`
	CreateAt               *model1.OrderBy `json:"create_at"`
	CreateBy               *model1.OrderBy `json:"create_by"`
	DeleteAt               *model1.OrderBy `json:"delete_at"`
	DeleteBy               *model1.OrderBy `json:"delete_by"`
	DriverID               *model1.OrderBy `json:"driver_id"`
	DrivingEndTime         *model1.OrderBy `json:"driving_end_time"`
	DrivingLogID           *model1.OrderBy `json:"driving_log_id"`
	DrivingStartTime       *model1.OrderBy `json:"driving_start_time"`
	EndTime                *model1.OrderBy `json:"end_time"`
	ID                     *model1.OrderBy `json:"id"`
	RegisterAt             *model1.OrderBy `json:"register_at"`
	RegisterBy             *model1.OrderBy `json:"register_by"`
	Remarks                *model1.OrderBy `json:"remarks"`
	Route                  *model1.OrderBy `json:"route"`
	StartTime              *model1.OrderBy `json:"start_time"`
	UpdateAt               *model1.OrderBy `json:"update_at"`
	UpdateBy               *model1.OrderBy `json:"update_by"`
	VehicleID              *model1.OrderBy `json:"vehicle_id"`
}

// aggregate min on columns
type DrivingLogMinFields struct {
	Cause                  *string    `json:"cause"`
	CheckOrganizationLevel *int       `json:"check_organization_level"`
	CheckState             *int       `json:"check_state"`
	CreateAt               *time.Time `json:"create_at"`
	CreateBy               *string    `json:"create_by"`
	DeleteAt               *time.Time `json:"delete_at"`
	DeleteBy               *string    `json:"delete_by"`
	DriverID               *string    `json:"driver_id"`
	DrivingEndTime         *string    `json:"driving_end_time"`
	DrivingLogID           *string    `json:"driving_log_id"`
	DrivingStartTime       *string    `json:"driving_start_time"`
	EndTime                *time.Time `json:"end_time"`
	ID                     *int64     `json:"id"`
	RegisterAt             *time.Time `json:"register_at"`
	RegisterBy             *string    `json:"register_by"`
	Remarks                *string    `json:"remarks"`
	Route                  *string    `json:"route"`
	StartTime              *time.Time `json:"start_time"`
	UpdateAt               *time.Time `json:"update_at"`
	UpdateBy               *string    `json:"update_by"`
	VehicleID              *string    `json:"vehicle_id"`
}

// order by min() on columns of table "driving_log"
type DrivingLogMinOrderBy struct {
	Cause                  *model1.OrderBy `json:"cause"`
	CheckOrganizationLevel *model1.OrderBy `json:"check_organization_level"`
	CheckState             *model1.OrderBy `json:"check_state"`
	CreateAt               *model1.OrderBy `json:"create_at"`
	CreateBy               *model1.OrderBy `json:"create_by"`
	DeleteAt               *model1.OrderBy `json:"delete_at"`
	DeleteBy               *model1.OrderBy `json:"delete_by"`
	DriverID               *model1.OrderBy `json:"driver_id"`
	DrivingEndTime         *model1.OrderBy `json:"driving_end_time"`
	DrivingLogID           *model1.OrderBy `json:"driving_log_id"`
	DrivingStartTime       *model1.OrderBy `json:"driving_start_time"`
	EndTime                *model1.OrderBy `json:"end_time"`
	ID                     *model1.OrderBy `json:"id"`
	RegisterAt             *model1.OrderBy `json:"register_at"`
	RegisterBy             *model1.OrderBy `json:"register_by"`
	Remarks                *model1.OrderBy `json:"remarks"`
	Route                  *model1.OrderBy `json:"route"`
	StartTime              *model1.OrderBy `json:"start_time"`
	UpdateAt               *model1.OrderBy `json:"update_at"`
	UpdateBy               *model1.OrderBy `json:"update_by"`
	VehicleID              *model1.OrderBy `json:"vehicle_id"`
}

// response of any mutation on the table "driving_log"
type DrivingLogMutationResponse struct {
	// number of affected rows by the mutation
	AffectedRows int `json:"affected_rows"`
	// data of the affected rows by the mutation
	Returning []*model.DrivingLog `json:"returning"`
}

// input type for inserting object relation for remote table "driving_log"
type DrivingLogObjRelInsertInput struct {
	Data       *DrivingLogInsertInput `json:"data"`
	OnConflict *DrivingLogOnConflict  `json:"on_conflict"`
}

// on conflict condition type for table "driving_log"
type DrivingLogOnConflict struct {
	Constraint    DrivingLogConstraint     `json:"constraint"`
	UpdateColumns []DrivingLogUpdateColumn `json:"update_columns"`
	Where         *DrivingLogBoolExp       `json:"where"`
}

// ordering options when selecting data from "driving_log"
type DrivingLogOrderBy struct {
	Cause                  *model1.OrderBy `json:"cause"`
	CheckOrganizationLevel *model1.OrderBy `json:"check_organization_level"`
	CheckState             *model1.OrderBy `json:"check_state"`
	CreateAt               *model1.OrderBy `json:"create_at"`
	CreateBy               *model1.OrderBy `json:"create_by"`
	DeleteAt               *model1.OrderBy `json:"delete_at"`
	DeleteBy               *model1.OrderBy `json:"delete_by"`
	DriverID               *model1.OrderBy `json:"driver_id"`
	DrivingEndTime         *model1.OrderBy `json:"driving_end_time"`
	DrivingLogID           *model1.OrderBy `json:"driving_log_id"`
	DrivingStartTime       *model1.OrderBy `json:"driving_start_time"`
	EndTime                *model1.OrderBy `json:"end_time"`
	ID                     *model1.OrderBy `json:"id"`
	IsDelete               *model1.OrderBy `json:"is_delete"`
	IsFill                 *model1.OrderBy `json:"is_fill"`
	RegisterAt             *model1.OrderBy `json:"register_at"`
	RegisterBy             *model1.OrderBy `json:"register_by"`
	Remarks                *model1.OrderBy `json:"remarks"`
	Route                  *model1.OrderBy `json:"route"`
	StartTime              *model1.OrderBy `json:"start_time"`
	UpdateAt               *model1.OrderBy `json:"update_at"`
	UpdateBy               *model1.OrderBy `json:"update_by"`
	VehicleID              *model1.OrderBy `json:"vehicle_id"`
}

// primary key columns input for table: "driving_log"
type DrivingLogPkColumnsInput struct {
	// ID
	ID int64 `json:"id"`
}

// input type for updating data in table "driving_log"
type DrivingLogSetInput struct {
	Cause                  *string    `json:"cause"`
	CheckOrganizationLevel *int       `json:"check_organization_level"`
	CheckState             *int       `json:"check_state"`
	CreateAt               *time.Time `json:"create_at"`
	CreateBy               *string    `json:"create_by"`
	DeleteAt               *time.Time `json:"delete_at"`
	DeleteBy               *string    `json:"delete_by"`
	DriverID               *string    `json:"driver_id"`
	DrivingEndTime         *string    `json:"driving_end_time"`
	DrivingLogID           *string    `json:"driving_log_id"`
	DrivingStartTime       *string    `json:"driving_start_time"`
	EndTime                *time.Time `json:"end_time"`
	ID                     *int64     `json:"id"`
	IsDelete               *bool      `json:"is_delete"`
	IsFill                 *bool      `json:"is_fill"`
	RegisterAt             *time.Time `json:"register_at"`
	RegisterBy             *string    `json:"register_by"`
	Remarks                *string    `json:"remarks"`
	Route                  *string    `json:"route"`
	StartTime              *time.Time `json:"start_time"`
	UpdateAt               *time.Time `json:"update_at"`
	UpdateBy               *string    `json:"update_by"`
	VehicleID              *string    `json:"vehicle_id"`
}

// aggregate stddev on columns
type DrivingLogStddevFields struct {
	CheckOrganizationLevel *float64 `json:"check_organization_level"`
	CheckState             *float64 `json:"check_state"`
	ID                     *float64 `json:"id"`
}

// order by stddev() on columns of table "driving_log"
type DrivingLogStddevOrderBy struct {
	CheckOrganizationLevel *model1.OrderBy `json:"check_organization_level"`
	CheckState             *model1.OrderBy `json:"check_state"`
	ID                     *model1.OrderBy `json:"id"`
}

// aggregate stddev_pop on columns
type DrivingLogStddevPopFields struct {
	CheckOrganizationLevel *float64 `json:"check_organization_level"`
	CheckState             *float64 `json:"check_state"`
	ID                     *float64 `json:"id"`
}

// order by stddev_pop() on columns of table "driving_log"
type DrivingLogStddevPopOrderBy struct {
	CheckOrganizationLevel *model1.OrderBy `json:"check_organization_level"`
	CheckState             *model1.OrderBy `json:"check_state"`
	ID                     *model1.OrderBy `json:"id"`
}

// aggregate stddev_samp on columns
type DrivingLogStddevSampFields struct {
	CheckOrganizationLevel *float64 `json:"check_organization_level"`
	CheckState             *float64 `json:"check_state"`
	ID                     *float64 `json:"id"`
}

// order by stddev_samp() on columns of table "driving_log"
type DrivingLogStddevSampOrderBy struct {
	CheckOrganizationLevel *model1.OrderBy `json:"check_organization_level"`
	CheckState             *model1.OrderBy `json:"check_state"`
	ID                     *model1.OrderBy `json:"id"`
}

// aggregate sum on columns
type DrivingLogSumFields struct {
	CheckOrganizationLevel *int   `json:"check_organization_level"`
	CheckState             *int   `json:"check_state"`
	ID                     *int64 `json:"id"`
}

// order by sum() on columns of table "driving_log"
type DrivingLogSumOrderBy struct {
	CheckOrganizationLevel *model1.OrderBy `json:"check_organization_level"`
	CheckState             *model1.OrderBy `json:"check_state"`
	ID                     *model1.OrderBy `json:"id"`
}

// aggregate var_pop on columns
type DrivingLogVarPopFields struct {
	CheckOrganizationLevel *float64 `json:"check_organization_level"`
	CheckState             *float64 `json:"check_state"`
	ID                     *float64 `json:"id"`
}

// order by var_pop() on columns of table "driving_log"
type DrivingLogVarPopOrderBy struct {
	CheckOrganizationLevel *model1.OrderBy `json:"check_organization_level"`
	CheckState             *model1.OrderBy `json:"check_state"`
	ID                     *model1.OrderBy `json:"id"`
}

// aggregate var_samp on columns
type DrivingLogVarSampFields struct {
	CheckOrganizationLevel *float64 `json:"check_organization_level"`
	CheckState             *float64 `json:"check_state"`
	ID                     *float64 `json:"id"`
}

// order by var_samp() on columns of table "driving_log"
type DrivingLogVarSampOrderBy struct {
	CheckOrganizationLevel *model1.OrderBy `json:"check_organization_level"`
	CheckState             *model1.OrderBy `json:"check_state"`
	ID                     *model1.OrderBy `json:"id"`
}

// aggregate variance on columns
type DrivingLogVarianceFields struct {
	CheckOrganizationLevel *float64 `json:"check_organization_level"`
	CheckState             *float64 `json:"check_state"`
	ID                     *float64 `json:"id"`
}

// order by variance() on columns of table "driving_log"
type DrivingLogVarianceOrderBy struct {
	CheckOrganizationLevel *model1.OrderBy `json:"check_organization_level"`
	CheckState             *model1.OrderBy `json:"check_state"`
	ID                     *model1.OrderBy `json:"id"`
}

// subscription root
type SubscriptionRoot struct {
	// fetch data from the table: "driving_log"
	DrivingLog []*model.DrivingLog `json:"driving_log"`
	// fetch aggregated fields from the table: "driving_log"
	DrivingLogAggregate *DrivingLogAggregate `json:"driving_log_aggregate"`
	// fetch data from the table: "driving_log" using primary key columns
	DrivingLogByPk *model.DrivingLog `json:"driving_log_by_pk"`
}

// unique or primary key constraints on table "driving_log"
type DrivingLogConstraint string

const (
	// unique or primary key constraint
	DrivingLogConstraintDrivingLogPkey DrivingLogConstraint = "driving_log_pkey"
)

var AllDrivingLogConstraint = []DrivingLogConstraint{
	DrivingLogConstraintDrivingLogPkey,
}

func (e DrivingLogConstraint) IsValid() bool {
	switch e {
	case DrivingLogConstraintDrivingLogPkey:
		return true
	}
	return false
}

func (e DrivingLogConstraint) String() string {
	return string(e)
}

func (e *DrivingLogConstraint) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DrivingLogConstraint(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid driving_log_constraint", str)
	}
	return nil
}

func (e DrivingLogConstraint) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// select columns of table "driving_log"
type DrivingLogSelectColumn string

const (
	// column name
	DrivingLogSelectColumnCause DrivingLogSelectColumn = "cause"
	// column name
	DrivingLogSelectColumnCheckOrganizationLevel DrivingLogSelectColumn = "check_organization_level"
	// column name
	DrivingLogSelectColumnCheckState DrivingLogSelectColumn = "check_state"
	// column name
	DrivingLogSelectColumnCreateAt DrivingLogSelectColumn = "create_at"
	// column name
	DrivingLogSelectColumnCreateBy DrivingLogSelectColumn = "create_by"
	// column name
	DrivingLogSelectColumnDeleteAt DrivingLogSelectColumn = "delete_at"
	// column name
	DrivingLogSelectColumnDeleteBy DrivingLogSelectColumn = "delete_by"
	// column name
	DrivingLogSelectColumnDriverID DrivingLogSelectColumn = "driver_id"
	// column name
	DrivingLogSelectColumnDrivingEndTime DrivingLogSelectColumn = "driving_end_time"
	// column name
	DrivingLogSelectColumnDrivingLogID DrivingLogSelectColumn = "driving_log_id"
	// column name
	DrivingLogSelectColumnDrivingStartTime DrivingLogSelectColumn = "driving_start_time"
	// column name
	DrivingLogSelectColumnEndTime DrivingLogSelectColumn = "end_time"
	// column name
	DrivingLogSelectColumnID DrivingLogSelectColumn = "id"
	// column name
	DrivingLogSelectColumnIsDelete DrivingLogSelectColumn = "is_delete"
	// column name
	DrivingLogSelectColumnIsFill DrivingLogSelectColumn = "is_fill"
	// column name
	DrivingLogSelectColumnRegisterAt DrivingLogSelectColumn = "register_at"
	// column name
	DrivingLogSelectColumnRegisterBy DrivingLogSelectColumn = "register_by"
	// column name
	DrivingLogSelectColumnRemarks DrivingLogSelectColumn = "remarks"
	// column name
	DrivingLogSelectColumnRoute DrivingLogSelectColumn = "route"
	// column name
	DrivingLogSelectColumnStartTime DrivingLogSelectColumn = "start_time"
	// column name
	DrivingLogSelectColumnUpdateAt DrivingLogSelectColumn = "update_at"
	// column name
	DrivingLogSelectColumnUpdateBy DrivingLogSelectColumn = "update_by"
	// column name
	DrivingLogSelectColumnVehicleID DrivingLogSelectColumn = "vehicle_id"
)

var AllDrivingLogSelectColumn = []DrivingLogSelectColumn{
	DrivingLogSelectColumnCause,
	DrivingLogSelectColumnCheckOrganizationLevel,
	DrivingLogSelectColumnCheckState,
	DrivingLogSelectColumnCreateAt,
	DrivingLogSelectColumnCreateBy,
	DrivingLogSelectColumnDeleteAt,
	DrivingLogSelectColumnDeleteBy,
	DrivingLogSelectColumnDriverID,
	DrivingLogSelectColumnDrivingEndTime,
	DrivingLogSelectColumnDrivingLogID,
	DrivingLogSelectColumnDrivingStartTime,
	DrivingLogSelectColumnEndTime,
	DrivingLogSelectColumnID,
	DrivingLogSelectColumnIsDelete,
	DrivingLogSelectColumnIsFill,
	DrivingLogSelectColumnRegisterAt,
	DrivingLogSelectColumnRegisterBy,
	DrivingLogSelectColumnRemarks,
	DrivingLogSelectColumnRoute,
	DrivingLogSelectColumnStartTime,
	DrivingLogSelectColumnUpdateAt,
	DrivingLogSelectColumnUpdateBy,
	DrivingLogSelectColumnVehicleID,
}

func (e DrivingLogSelectColumn) IsValid() bool {
	switch e {
	case DrivingLogSelectColumnCause, DrivingLogSelectColumnCheckOrganizationLevel, DrivingLogSelectColumnCheckState, DrivingLogSelectColumnCreateAt, DrivingLogSelectColumnCreateBy, DrivingLogSelectColumnDeleteAt, DrivingLogSelectColumnDeleteBy, DrivingLogSelectColumnDriverID, DrivingLogSelectColumnDrivingEndTime, DrivingLogSelectColumnDrivingLogID, DrivingLogSelectColumnDrivingStartTime, DrivingLogSelectColumnEndTime, DrivingLogSelectColumnID, DrivingLogSelectColumnIsDelete, DrivingLogSelectColumnIsFill, DrivingLogSelectColumnRegisterAt, DrivingLogSelectColumnRegisterBy, DrivingLogSelectColumnRemarks, DrivingLogSelectColumnRoute, DrivingLogSelectColumnStartTime, DrivingLogSelectColumnUpdateAt, DrivingLogSelectColumnUpdateBy, DrivingLogSelectColumnVehicleID:
		return true
	}
	return false
}

func (e DrivingLogSelectColumn) String() string {
	return string(e)
}

func (e *DrivingLogSelectColumn) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DrivingLogSelectColumn(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid driving_log_select_column", str)
	}
	return nil
}

func (e DrivingLogSelectColumn) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// update columns of table "driving_log"
type DrivingLogUpdateColumn string

const (
	// column name
	DrivingLogUpdateColumnCause DrivingLogUpdateColumn = "cause"
	// column name
	DrivingLogUpdateColumnCheckOrganizationLevel DrivingLogUpdateColumn = "check_organization_level"
	// column name
	DrivingLogUpdateColumnCheckState DrivingLogUpdateColumn = "check_state"
	// column name
	DrivingLogUpdateColumnCreateAt DrivingLogUpdateColumn = "create_at"
	// column name
	DrivingLogUpdateColumnCreateBy DrivingLogUpdateColumn = "create_by"
	// column name
	DrivingLogUpdateColumnDeleteAt DrivingLogUpdateColumn = "delete_at"
	// column name
	DrivingLogUpdateColumnDeleteBy DrivingLogUpdateColumn = "delete_by"
	// column name
	DrivingLogUpdateColumnDriverID DrivingLogUpdateColumn = "driver_id"
	// column name
	DrivingLogUpdateColumnDrivingEndTime DrivingLogUpdateColumn = "driving_end_time"
	// column name
	DrivingLogUpdateColumnDrivingLogID DrivingLogUpdateColumn = "driving_log_id"
	// column name
	DrivingLogUpdateColumnDrivingStartTime DrivingLogUpdateColumn = "driving_start_time"
	// column name
	DrivingLogUpdateColumnEndTime DrivingLogUpdateColumn = "end_time"
	// column name
	DrivingLogUpdateColumnID DrivingLogUpdateColumn = "id"
	// column name
	DrivingLogUpdateColumnIsDelete DrivingLogUpdateColumn = "is_delete"
	// column name
	DrivingLogUpdateColumnIsFill DrivingLogUpdateColumn = "is_fill"
	// column name
	DrivingLogUpdateColumnRegisterAt DrivingLogUpdateColumn = "register_at"
	// column name
	DrivingLogUpdateColumnRegisterBy DrivingLogUpdateColumn = "register_by"
	// column name
	DrivingLogUpdateColumnRemarks DrivingLogUpdateColumn = "remarks"
	// column name
	DrivingLogUpdateColumnRoute DrivingLogUpdateColumn = "route"
	// column name
	DrivingLogUpdateColumnStartTime DrivingLogUpdateColumn = "start_time"
	// column name
	DrivingLogUpdateColumnUpdateAt DrivingLogUpdateColumn = "update_at"
	// column name
	DrivingLogUpdateColumnUpdateBy DrivingLogUpdateColumn = "update_by"
	// column name
	DrivingLogUpdateColumnVehicleID DrivingLogUpdateColumn = "vehicle_id"
)

var AllDrivingLogUpdateColumn = []DrivingLogUpdateColumn{
	DrivingLogUpdateColumnCause,
	DrivingLogUpdateColumnCheckOrganizationLevel,
	DrivingLogUpdateColumnCheckState,
	DrivingLogUpdateColumnCreateAt,
	DrivingLogUpdateColumnCreateBy,
	DrivingLogUpdateColumnDeleteAt,
	DrivingLogUpdateColumnDeleteBy,
	DrivingLogUpdateColumnDriverID,
	DrivingLogUpdateColumnDrivingEndTime,
	DrivingLogUpdateColumnDrivingLogID,
	DrivingLogUpdateColumnDrivingStartTime,
	DrivingLogUpdateColumnEndTime,
	DrivingLogUpdateColumnID,
	DrivingLogUpdateColumnIsDelete,
	DrivingLogUpdateColumnIsFill,
	DrivingLogUpdateColumnRegisterAt,
	DrivingLogUpdateColumnRegisterBy,
	DrivingLogUpdateColumnRemarks,
	DrivingLogUpdateColumnRoute,
	DrivingLogUpdateColumnStartTime,
	DrivingLogUpdateColumnUpdateAt,
	DrivingLogUpdateColumnUpdateBy,
	DrivingLogUpdateColumnVehicleID,
}

func (e DrivingLogUpdateColumn) IsValid() bool {
	switch e {
	case DrivingLogUpdateColumnCause, DrivingLogUpdateColumnCheckOrganizationLevel, DrivingLogUpdateColumnCheckState, DrivingLogUpdateColumnCreateAt, DrivingLogUpdateColumnCreateBy, DrivingLogUpdateColumnDeleteAt, DrivingLogUpdateColumnDeleteBy, DrivingLogUpdateColumnDriverID, DrivingLogUpdateColumnDrivingEndTime, DrivingLogUpdateColumnDrivingLogID, DrivingLogUpdateColumnDrivingStartTime, DrivingLogUpdateColumnEndTime, DrivingLogUpdateColumnID, DrivingLogUpdateColumnIsDelete, DrivingLogUpdateColumnIsFill, DrivingLogUpdateColumnRegisterAt, DrivingLogUpdateColumnRegisterBy, DrivingLogUpdateColumnRemarks, DrivingLogUpdateColumnRoute, DrivingLogUpdateColumnStartTime, DrivingLogUpdateColumnUpdateAt, DrivingLogUpdateColumnUpdateBy, DrivingLogUpdateColumnVehicleID:
		return true
	}
	return false
}

func (e DrivingLogUpdateColumn) String() string {
	return string(e)
}

func (e *DrivingLogUpdateColumn) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DrivingLogUpdateColumn(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid driving_log_update_column", str)
	}
	return nil
}

func (e DrivingLogUpdateColumn) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
