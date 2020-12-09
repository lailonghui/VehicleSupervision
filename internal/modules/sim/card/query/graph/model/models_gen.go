// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"VehicleSupervision/internal/modules/sim/card/model"
	model1 "VehicleSupervision/pkg/graphql/model"
	"fmt"
	"io"
	"strconv"
	"time"
)

// mutation root
type MutationRoot struct {
	// delete data from the table: "sim_card"
	DeleteSimCard *SimCardMutationResponse `json:"delete_sim_card"`
	// delete single row from the table: "sim_card"
	DeleteSimCardByPk *model.SimCard `json:"delete_sim_card_by_pk"`
	// insert data into the table: "sim_card"
	InsertSimCard *SimCardMutationResponse `json:"insert_sim_card"`
	// insert a single row into the table: "sim_card"
	InsertSimCardOne *model.SimCard `json:"insert_sim_card_one"`
	// update data of the table: "sim_card"
	UpdateSimCard *SimCardMutationResponse `json:"update_sim_card"`
	// update single row of the table: "sim_card"
	UpdateSimCardByPk *model.SimCard `json:"update_sim_card_by_pk"`
}

// aggregated selection of "sim_card"
type SimCardAggregate struct {
	Aggregate *SimCardAggregateFields `json:"aggregate"`
	Nodes     []*model.SimCard        `json:"nodes"`
}

// aggregate fields of "sim_card"
type SimCardAggregateFields struct {
	Avg        *SimCardAvgFields        `json:"avg"`
	Count      *int                     `json:"count"`
	Max        *SimCardMaxFields        `json:"max"`
	Min        *SimCardMinFields        `json:"min"`
	Stddev     *SimCardStddevFields     `json:"stddev"`
	StddevPop  *SimCardStddevPopFields  `json:"stddev_pop"`
	StddevSamp *SimCardStddevSampFields `json:"stddev_samp"`
	Sum        *SimCardSumFields        `json:"sum"`
	VarPop     *SimCardVarPopFields     `json:"var_pop"`
	VarSamp    *SimCardVarSampFields    `json:"var_samp"`
	Variance   *SimCardVarianceFields   `json:"variance"`
}

// order by aggregate values of table "sim_card"
type SimCardAggregateOrderBy struct {
	Avg        *SimCardAvgOrderBy        `json:"avg"`
	Count      *model1.OrderBy           `json:"count"`
	Max        *SimCardMaxOrderBy        `json:"max"`
	Min        *SimCardMinOrderBy        `json:"min"`
	Stddev     *SimCardStddevOrderBy     `json:"stddev"`
	StddevPop  *SimCardStddevPopOrderBy  `json:"stddev_pop"`
	StddevSamp *SimCardStddevSampOrderBy `json:"stddev_samp"`
	Sum        *SimCardSumOrderBy        `json:"sum"`
	VarPop     *SimCardVarPopOrderBy     `json:"var_pop"`
	VarSamp    *SimCardVarSampOrderBy    `json:"var_samp"`
	Variance   *SimCardVarianceOrderBy   `json:"variance"`
}

// input type for inserting array relation for remote table "sim_card"
type SimCardArrRelInsertInput struct {
	Data       []*SimCardInsertInput `json:"data"`
	OnConflict *SimCardOnConflict    `json:"on_conflict"`
}

// aggregate avg on columns
type SimCardAvgFields struct {
	ID         *float64 `json:"id"`
	MobileType *float64 `json:"mobile_type"`
}

// order by avg() on columns of table "sim_card"
type SimCardAvgOrderBy struct {
	ID         *model1.OrderBy `json:"id"`
	MobileType *model1.OrderBy `json:"mobile_type"`
}

// Boolean expression to filter rows from the table "sim_card". All fields are combined with a logical 'AND'.
type SimCardBoolExp struct {
	And         []*SimCardBoolExp                `json:"_and"`
	Not         *SimCardBoolExp                  `json:"_not"`
	Or          []*SimCardBoolExp                `json:"_or"`
	CreateAt    *model1.TimestamptzComparisonExp `json:"create_at"`
	CreateBy    *model1.StringComparisonExp      `json:"create_by"`
	DeleteAt    *model1.TimestamptzComparisonExp `json:"delete_at"`
	DeleteBy    *model1.StringComparisonExp      `json:"delete_by"`
	DeptID      *model1.StringComparisonExp      `json:"dept_id"`
	ID          *model1.BigintComparisonExp      `json:"id"`
	IsDelete    *model1.BooleanComparisonExp     `json:"is_delete"`
	MobileType  *model1.IntComparisonExp         `json:"mobile_type"`
	OperatorsID *model1.StringComparisonExp      `json:"operators_id"`
	ProxyrgID   *model1.StringComparisonExp      `json:"proxyrg_id"`
	Remark      *model1.StringComparisonExp      `json:"remark"`
	SimCardID   *model1.StringComparisonExp      `json:"sim_card_id"`
	SimNumber   *model1.StringComparisonExp      `json:"sim_number"`
	TerminalID  *model1.StringComparisonExp      `json:"terminal_id"`
	UpdateAt    *model1.TimestamptzComparisonExp `json:"update_at"`
	UpdateBy    *model1.StringComparisonExp      `json:"update_by"`
}

// input type for incrementing integer column in table "sim_card"
type SimCardIncInput struct {
	ID         *int64 `json:"id"`
	MobileType *int   `json:"mobile_type"`
}

// input type for inserting data into table "sim_card"
type SimCardInsertInput struct {
	CreateAt    *time.Time `json:"create_at"`
	CreateBy    *string    `json:"create_by"`
	DeleteAt    *time.Time `json:"delete_at"`
	DeleteBy    *string    `json:"delete_by"`
	DeptID      *string    `json:"dept_id"`
	ID          *int64     `json:"id"`
	IsDelete    *bool      `json:"is_delete"`
	MobileType  *int       `json:"mobile_type"`
	OperatorsID *string    `json:"operators_id"`
	ProxyrgID   *string    `json:"proxyrg_id"`
	Remark      *string    `json:"remark"`
	SimCardID   *string    `json:"sim_card_id"`
	SimNumber   *string    `json:"sim_number"`
	TerminalID  *string    `json:"terminal_id"`
	UpdateAt    *time.Time `json:"update_at"`
	UpdateBy    *string    `json:"update_by"`
}

// aggregate max on columns
type SimCardMaxFields struct {
	CreateAt    *time.Time `json:"create_at"`
	CreateBy    *string    `json:"create_by"`
	DeleteAt    *time.Time `json:"delete_at"`
	DeleteBy    *string    `json:"delete_by"`
	DeptID      *string    `json:"dept_id"`
	ID          *int64     `json:"id"`
	MobileType  *int       `json:"mobile_type"`
	OperatorsID *string    `json:"operators_id"`
	ProxyrgID   *string    `json:"proxyrg_id"`
	Remark      *string    `json:"remark"`
	SimCardID   *string    `json:"sim_card_id"`
	SimNumber   *string    `json:"sim_number"`
	TerminalID  *string    `json:"terminal_id"`
	UpdateAt    *time.Time `json:"update_at"`
	UpdateBy    *string    `json:"update_by"`
}

// order by max() on columns of table "sim_card"
type SimCardMaxOrderBy struct {
	CreateAt    *model1.OrderBy `json:"create_at"`
	CreateBy    *model1.OrderBy `json:"create_by"`
	DeleteAt    *model1.OrderBy `json:"delete_at"`
	DeleteBy    *model1.OrderBy `json:"delete_by"`
	DeptID      *model1.OrderBy `json:"dept_id"`
	ID          *model1.OrderBy `json:"id"`
	MobileType  *model1.OrderBy `json:"mobile_type"`
	OperatorsID *model1.OrderBy `json:"operators_id"`
	ProxyrgID   *model1.OrderBy `json:"proxyrg_id"`
	Remark      *model1.OrderBy `json:"remark"`
	SimCardID   *model1.OrderBy `json:"sim_card_id"`
	SimNumber   *model1.OrderBy `json:"sim_number"`
	TerminalID  *model1.OrderBy `json:"terminal_id"`
	UpdateAt    *model1.OrderBy `json:"update_at"`
	UpdateBy    *model1.OrderBy `json:"update_by"`
}

// aggregate min on columns
type SimCardMinFields struct {
	CreateAt    *time.Time `json:"create_at"`
	CreateBy    *string    `json:"create_by"`
	DeleteAt    *time.Time `json:"delete_at"`
	DeleteBy    *string    `json:"delete_by"`
	DeptID      *string    `json:"dept_id"`
	ID          *int64     `json:"id"`
	MobileType  *int       `json:"mobile_type"`
	OperatorsID *string    `json:"operators_id"`
	ProxyrgID   *string    `json:"proxyrg_id"`
	Remark      *string    `json:"remark"`
	SimCardID   *string    `json:"sim_card_id"`
	SimNumber   *string    `json:"sim_number"`
	TerminalID  *string    `json:"terminal_id"`
	UpdateAt    *time.Time `json:"update_at"`
	UpdateBy    *string    `json:"update_by"`
}

// order by min() on columns of table "sim_card"
type SimCardMinOrderBy struct {
	CreateAt    *model1.OrderBy `json:"create_at"`
	CreateBy    *model1.OrderBy `json:"create_by"`
	DeleteAt    *model1.OrderBy `json:"delete_at"`
	DeleteBy    *model1.OrderBy `json:"delete_by"`
	DeptID      *model1.OrderBy `json:"dept_id"`
	ID          *model1.OrderBy `json:"id"`
	MobileType  *model1.OrderBy `json:"mobile_type"`
	OperatorsID *model1.OrderBy `json:"operators_id"`
	ProxyrgID   *model1.OrderBy `json:"proxyrg_id"`
	Remark      *model1.OrderBy `json:"remark"`
	SimCardID   *model1.OrderBy `json:"sim_card_id"`
	SimNumber   *model1.OrderBy `json:"sim_number"`
	TerminalID  *model1.OrderBy `json:"terminal_id"`
	UpdateAt    *model1.OrderBy `json:"update_at"`
	UpdateBy    *model1.OrderBy `json:"update_by"`
}

// response of any mutation on the table "sim_card"
type SimCardMutationResponse struct {
	// number of affected rows by the mutation
	AffectedRows int `json:"affected_rows"`
	// data of the affected rows by the mutation
	Returning []*model.SimCard `json:"returning"`
}

// input type for inserting object relation for remote table "sim_card"
type SimCardObjRelInsertInput struct {
	Data       *SimCardInsertInput `json:"data"`
	OnConflict *SimCardOnConflict  `json:"on_conflict"`
}

// on conflict condition type for table "sim_card"
type SimCardOnConflict struct {
	Constraint    SimCardConstraint     `json:"constraint"`
	UpdateColumns []SimCardUpdateColumn `json:"update_columns"`
	Where         *SimCardBoolExp       `json:"where"`
}

// ordering options when selecting data from "sim_card"
type SimCardOrderBy struct {
	CreateAt    *model1.OrderBy `json:"create_at"`
	CreateBy    *model1.OrderBy `json:"create_by"`
	DeleteAt    *model1.OrderBy `json:"delete_at"`
	DeleteBy    *model1.OrderBy `json:"delete_by"`
	DeptID      *model1.OrderBy `json:"dept_id"`
	ID          *model1.OrderBy `json:"id"`
	IsDelete    *model1.OrderBy `json:"is_delete"`
	MobileType  *model1.OrderBy `json:"mobile_type"`
	OperatorsID *model1.OrderBy `json:"operators_id"`
	ProxyrgID   *model1.OrderBy `json:"proxyrg_id"`
	Remark      *model1.OrderBy `json:"remark"`
	SimCardID   *model1.OrderBy `json:"sim_card_id"`
	SimNumber   *model1.OrderBy `json:"sim_number"`
	TerminalID  *model1.OrderBy `json:"terminal_id"`
	UpdateAt    *model1.OrderBy `json:"update_at"`
	UpdateBy    *model1.OrderBy `json:"update_by"`
}

// primary key columns input for table: "sim_card"
type SimCardPkColumnsInput struct {
	// ID
	ID int64 `json:"id"`
}

// input type for updating data in table "sim_card"
type SimCardSetInput struct {
	CreateAt    *time.Time `json:"create_at"`
	CreateBy    *string    `json:"create_by"`
	DeleteAt    *time.Time `json:"delete_at"`
	DeleteBy    *string    `json:"delete_by"`
	DeptID      *string    `json:"dept_id"`
	ID          *int64     `json:"id"`
	IsDelete    *bool      `json:"is_delete"`
	MobileType  *int       `json:"mobile_type"`
	OperatorsID *string    `json:"operators_id"`
	ProxyrgID   *string    `json:"proxyrg_id"`
	Remark      *string    `json:"remark"`
	SimCardID   *string    `json:"sim_card_id"`
	SimNumber   *string    `json:"sim_number"`
	TerminalID  *string    `json:"terminal_id"`
	UpdateAt    *time.Time `json:"update_at"`
	UpdateBy    *string    `json:"update_by"`
}

// aggregate stddev on columns
type SimCardStddevFields struct {
	ID         *float64 `json:"id"`
	MobileType *float64 `json:"mobile_type"`
}

// order by stddev() on columns of table "sim_card"
type SimCardStddevOrderBy struct {
	ID         *model1.OrderBy `json:"id"`
	MobileType *model1.OrderBy `json:"mobile_type"`
}

// aggregate stddev_pop on columns
type SimCardStddevPopFields struct {
	ID         *float64 `json:"id"`
	MobileType *float64 `json:"mobile_type"`
}

// order by stddev_pop() on columns of table "sim_card"
type SimCardStddevPopOrderBy struct {
	ID         *model1.OrderBy `json:"id"`
	MobileType *model1.OrderBy `json:"mobile_type"`
}

// aggregate stddev_samp on columns
type SimCardStddevSampFields struct {
	ID         *float64 `json:"id"`
	MobileType *float64 `json:"mobile_type"`
}

// order by stddev_samp() on columns of table "sim_card"
type SimCardStddevSampOrderBy struct {
	ID         *model1.OrderBy `json:"id"`
	MobileType *model1.OrderBy `json:"mobile_type"`
}

// aggregate sum on columns
type SimCardSumFields struct {
	ID         *int64 `json:"id"`
	MobileType *int   `json:"mobile_type"`
}

// order by sum() on columns of table "sim_card"
type SimCardSumOrderBy struct {
	ID         *model1.OrderBy `json:"id"`
	MobileType *model1.OrderBy `json:"mobile_type"`
}

// aggregate var_pop on columns
type SimCardVarPopFields struct {
	ID         *float64 `json:"id"`
	MobileType *float64 `json:"mobile_type"`
}

// order by var_pop() on columns of table "sim_card"
type SimCardVarPopOrderBy struct {
	ID         *model1.OrderBy `json:"id"`
	MobileType *model1.OrderBy `json:"mobile_type"`
}

// aggregate var_samp on columns
type SimCardVarSampFields struct {
	ID         *float64 `json:"id"`
	MobileType *float64 `json:"mobile_type"`
}

// order by var_samp() on columns of table "sim_card"
type SimCardVarSampOrderBy struct {
	ID         *model1.OrderBy `json:"id"`
	MobileType *model1.OrderBy `json:"mobile_type"`
}

// aggregate variance on columns
type SimCardVarianceFields struct {
	ID         *float64 `json:"id"`
	MobileType *float64 `json:"mobile_type"`
}

// order by variance() on columns of table "sim_card"
type SimCardVarianceOrderBy struct {
	ID         *model1.OrderBy `json:"id"`
	MobileType *model1.OrderBy `json:"mobile_type"`
}

// subscription root
type SubscriptionRoot struct {
	// fetch data from the table: "sim_card"
	SimCard []*model.SimCard `json:"sim_card"`
	// fetch aggregated fields from the table: "sim_card"
	SimCardAggregate *SimCardAggregate `json:"sim_card_aggregate"`
	// fetch data from the table: "sim_card" using primary key columns
	SimCardByPk *model.SimCard `json:"sim_card_by_pk"`
}

// unique or primary key constraints on table "sim_card"
type SimCardConstraint string

const (
	// unique or primary key constraint
	SimCardConstraintSimCardPkey SimCardConstraint = "sim_card_pkey"
)

var AllSimCardConstraint = []SimCardConstraint{
	SimCardConstraintSimCardPkey,
}

func (e SimCardConstraint) IsValid() bool {
	switch e {
	case SimCardConstraintSimCardPkey:
		return true
	}
	return false
}

func (e SimCardConstraint) String() string {
	return string(e)
}

func (e *SimCardConstraint) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SimCardConstraint(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid sim_card_constraint", str)
	}
	return nil
}

func (e SimCardConstraint) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// select columns of table "sim_card"
type SimCardSelectColumn string

const (
	// column name
	SimCardSelectColumnCreateAt SimCardSelectColumn = "create_at"
	// column name
	SimCardSelectColumnCreateBy SimCardSelectColumn = "create_by"
	// column name
	SimCardSelectColumnDeleteAt SimCardSelectColumn = "delete_at"
	// column name
	SimCardSelectColumnDeleteBy SimCardSelectColumn = "delete_by"
	// column name
	SimCardSelectColumnDeptID SimCardSelectColumn = "dept_id"
	// column name
	SimCardSelectColumnID SimCardSelectColumn = "id"
	// column name
	SimCardSelectColumnIsDelete SimCardSelectColumn = "is_delete"
	// column name
	SimCardSelectColumnMobileType SimCardSelectColumn = "mobile_type"
	// column name
	SimCardSelectColumnOperatorsID SimCardSelectColumn = "operators_id"
	// column name
	SimCardSelectColumnProxyrgID SimCardSelectColumn = "proxyrg_id"
	// column name
	SimCardSelectColumnRemark SimCardSelectColumn = "remark"
	// column name
	SimCardSelectColumnSimCardID SimCardSelectColumn = "sim_card_id"
	// column name
	SimCardSelectColumnSimNumber SimCardSelectColumn = "sim_number"
	// column name
	SimCardSelectColumnTerminalID SimCardSelectColumn = "terminal_id"
	// column name
	SimCardSelectColumnUpdateAt SimCardSelectColumn = "update_at"
	// column name
	SimCardSelectColumnUpdateBy SimCardSelectColumn = "update_by"
)

var AllSimCardSelectColumn = []SimCardSelectColumn{
	SimCardSelectColumnCreateAt,
	SimCardSelectColumnCreateBy,
	SimCardSelectColumnDeleteAt,
	SimCardSelectColumnDeleteBy,
	SimCardSelectColumnDeptID,
	SimCardSelectColumnID,
	SimCardSelectColumnIsDelete,
	SimCardSelectColumnMobileType,
	SimCardSelectColumnOperatorsID,
	SimCardSelectColumnProxyrgID,
	SimCardSelectColumnRemark,
	SimCardSelectColumnSimCardID,
	SimCardSelectColumnSimNumber,
	SimCardSelectColumnTerminalID,
	SimCardSelectColumnUpdateAt,
	SimCardSelectColumnUpdateBy,
}

func (e SimCardSelectColumn) IsValid() bool {
	switch e {
	case SimCardSelectColumnCreateAt, SimCardSelectColumnCreateBy, SimCardSelectColumnDeleteAt, SimCardSelectColumnDeleteBy, SimCardSelectColumnDeptID, SimCardSelectColumnID, SimCardSelectColumnIsDelete, SimCardSelectColumnMobileType, SimCardSelectColumnOperatorsID, SimCardSelectColumnProxyrgID, SimCardSelectColumnRemark, SimCardSelectColumnSimCardID, SimCardSelectColumnSimNumber, SimCardSelectColumnTerminalID, SimCardSelectColumnUpdateAt, SimCardSelectColumnUpdateBy:
		return true
	}
	return false
}

func (e SimCardSelectColumn) String() string {
	return string(e)
}

func (e *SimCardSelectColumn) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SimCardSelectColumn(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid sim_card_select_column", str)
	}
	return nil
}

func (e SimCardSelectColumn) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// update columns of table "sim_card"
type SimCardUpdateColumn string

const (
	// column name
	SimCardUpdateColumnCreateAt SimCardUpdateColumn = "create_at"
	// column name
	SimCardUpdateColumnCreateBy SimCardUpdateColumn = "create_by"
	// column name
	SimCardUpdateColumnDeleteAt SimCardUpdateColumn = "delete_at"
	// column name
	SimCardUpdateColumnDeleteBy SimCardUpdateColumn = "delete_by"
	// column name
	SimCardUpdateColumnDeptID SimCardUpdateColumn = "dept_id"
	// column name
	SimCardUpdateColumnID SimCardUpdateColumn = "id"
	// column name
	SimCardUpdateColumnIsDelete SimCardUpdateColumn = "is_delete"
	// column name
	SimCardUpdateColumnMobileType SimCardUpdateColumn = "mobile_type"
	// column name
	SimCardUpdateColumnOperatorsID SimCardUpdateColumn = "operators_id"
	// column name
	SimCardUpdateColumnProxyrgID SimCardUpdateColumn = "proxyrg_id"
	// column name
	SimCardUpdateColumnRemark SimCardUpdateColumn = "remark"
	// column name
	SimCardUpdateColumnSimCardID SimCardUpdateColumn = "sim_card_id"
	// column name
	SimCardUpdateColumnSimNumber SimCardUpdateColumn = "sim_number"
	// column name
	SimCardUpdateColumnTerminalID SimCardUpdateColumn = "terminal_id"
	// column name
	SimCardUpdateColumnUpdateAt SimCardUpdateColumn = "update_at"
	// column name
	SimCardUpdateColumnUpdateBy SimCardUpdateColumn = "update_by"
)

var AllSimCardUpdateColumn = []SimCardUpdateColumn{
	SimCardUpdateColumnCreateAt,
	SimCardUpdateColumnCreateBy,
	SimCardUpdateColumnDeleteAt,
	SimCardUpdateColumnDeleteBy,
	SimCardUpdateColumnDeptID,
	SimCardUpdateColumnID,
	SimCardUpdateColumnIsDelete,
	SimCardUpdateColumnMobileType,
	SimCardUpdateColumnOperatorsID,
	SimCardUpdateColumnProxyrgID,
	SimCardUpdateColumnRemark,
	SimCardUpdateColumnSimCardID,
	SimCardUpdateColumnSimNumber,
	SimCardUpdateColumnTerminalID,
	SimCardUpdateColumnUpdateAt,
	SimCardUpdateColumnUpdateBy,
}

func (e SimCardUpdateColumn) IsValid() bool {
	switch e {
	case SimCardUpdateColumnCreateAt, SimCardUpdateColumnCreateBy, SimCardUpdateColumnDeleteAt, SimCardUpdateColumnDeleteBy, SimCardUpdateColumnDeptID, SimCardUpdateColumnID, SimCardUpdateColumnIsDelete, SimCardUpdateColumnMobileType, SimCardUpdateColumnOperatorsID, SimCardUpdateColumnProxyrgID, SimCardUpdateColumnRemark, SimCardUpdateColumnSimCardID, SimCardUpdateColumnSimNumber, SimCardUpdateColumnTerminalID, SimCardUpdateColumnUpdateAt, SimCardUpdateColumnUpdateBy:
		return true
	}
	return false
}

func (e SimCardUpdateColumn) String() string {
	return string(e)
}

func (e *SimCardUpdateColumn) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SimCardUpdateColumn(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid sim_card_update_column", str)
	}
	return nil
}

func (e SimCardUpdateColumn) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
