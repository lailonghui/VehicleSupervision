// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"VehicleSupervision/internal/modules/blacklist/model"
	model1 "VehicleSupervision/pkg/graphql/model"
	"fmt"
	"io"
	"strconv"
	"time"
)

// aggregated selection of "blacklist_operation_record"
type BlacklistOperationRecordAggregate struct {
	Aggregate *BlacklistOperationRecordAggregateFields `json:"aggregate"`
	Nodes     []*model.BlacklistOperationRecord        `json:"nodes"`
}

// aggregate fields of "blacklist_operation_record"
type BlacklistOperationRecordAggregateFields struct {
	Avg        *BlacklistOperationRecordAvgFields        `json:"avg"`
	Count      *int                                      `json:"count"`
	Max        *BlacklistOperationRecordMaxFields        `json:"max"`
	Min        *BlacklistOperationRecordMinFields        `json:"min"`
	Stddev     *BlacklistOperationRecordStddevFields     `json:"stddev"`
	StddevPop  *BlacklistOperationRecordStddevPopFields  `json:"stddev_pop"`
	StddevSamp *BlacklistOperationRecordStddevSampFields `json:"stddev_samp"`
	Sum        *BlacklistOperationRecordSumFields        `json:"sum"`
	VarPop     *BlacklistOperationRecordVarPopFields     `json:"var_pop"`
	VarSamp    *BlacklistOperationRecordVarSampFields    `json:"var_samp"`
	Variance   *BlacklistOperationRecordVarianceFields   `json:"variance"`
}

// order by aggregate values of table "blacklist_operation_record"
type BlacklistOperationRecordAggregateOrderBy struct {
	Avg        *BlacklistOperationRecordAvgOrderBy        `json:"avg"`
	Count      *model1.OrderBy                            `json:"count"`
	Max        *BlacklistOperationRecordMaxOrderBy        `json:"max"`
	Min        *BlacklistOperationRecordMinOrderBy        `json:"min"`
	Stddev     *BlacklistOperationRecordStddevOrderBy     `json:"stddev"`
	StddevPop  *BlacklistOperationRecordStddevPopOrderBy  `json:"stddev_pop"`
	StddevSamp *BlacklistOperationRecordStddevSampOrderBy `json:"stddev_samp"`
	Sum        *BlacklistOperationRecordSumOrderBy        `json:"sum"`
	VarPop     *BlacklistOperationRecordVarPopOrderBy     `json:"var_pop"`
	VarSamp    *BlacklistOperationRecordVarSampOrderBy    `json:"var_samp"`
	Variance   *BlacklistOperationRecordVarianceOrderBy   `json:"variance"`
}

// input type for inserting array relation for remote table "blacklist_operation_record"
type BlacklistOperationRecordArrRelInsertInput struct {
	Data       []*BlacklistOperationRecordInsertInput `json:"data"`
	OnConflict *BlacklistOperationRecordOnConflict    `json:"on_conflict"`
}

// aggregate avg on columns
type BlacklistOperationRecordAvgFields struct {
	BlacklistType *float64 `json:"blacklist_type"`
	ID            *float64 `json:"id"`
	Operate       *float64 `json:"operate"`
}

// order by avg() on columns of table "blacklist_operation_record"
type BlacklistOperationRecordAvgOrderBy struct {
	BlacklistType *model1.OrderBy `json:"blacklist_type"`
	ID            *model1.OrderBy `json:"id"`
	Operate       *model1.OrderBy `json:"operate"`
}

// Boolean expression to filter rows from the table "blacklist_operation_record". All fields are combined with a logical 'AND'.
type BlacklistOperationRecordBoolExp struct {
	And               []*BlacklistOperationRecordBoolExp `json:"_and"`
	Not               *BlacklistOperationRecordBoolExp   `json:"_not"`
	Or                []*BlacklistOperationRecordBoolExp `json:"_or"`
	BlacklistRecordID *model1.StringComparisonExp        `json:"blacklist_record_id"`
	BlacklistType     *model1.IntComparisonExp           `json:"blacklist_type"`
	CreateAt          *model1.TimestamptzComparisonExp   `json:"create_at"`
	CreateBy          *model1.StringComparisonExp        `json:"create_by"`
	DeleteAt          *model1.TimestamptzComparisonExp   `json:"delete_at"`
	DeleteBy          *model1.StringComparisonExp        `json:"delete_by"`
	ID                *model1.BigintComparisonExp        `json:"id"`
	IsDelete          *model1.BooleanComparisonExp       `json:"is_delete"`
	Operate           *model1.IntComparisonExp           `json:"operate"`
	Remarks           *model1.StringComparisonExp        `json:"remarks"`
	UpdateAt          *model1.TimestamptzComparisonExp   `json:"update_at"`
	UpdateBy          *model1.StringComparisonExp        `json:"update_by"`
	VSeqn             *model1.StringComparisonExp        `json:"v_seqn"`
}

// input type for incrementing integer column in table "blacklist_operation_record"
type BlacklistOperationRecordIncInput struct {
	BlacklistType *int   `json:"blacklist_type"`
	ID            *int64 `json:"id"`
	Operate       *int   `json:"operate"`
}

// input type for inserting data into table "blacklist_operation_record"
type BlacklistOperationRecordInsertInput struct {
	BlacklistRecordID *string    `json:"blacklist_record_id"`
	BlacklistType     *int       `json:"blacklist_type"`
	CreateAt          *time.Time `json:"create_at"`
	CreateBy          *string    `json:"create_by"`
	DeleteAt          *time.Time `json:"delete_at"`
	DeleteBy          *string    `json:"delete_by"`
	ID                *int64     `json:"id"`
	IsDelete          *bool      `json:"is_delete"`
	Operate           *int       `json:"operate"`
	Remarks           *string    `json:"remarks"`
	UpdateAt          *time.Time `json:"update_at"`
	UpdateBy          *string    `json:"update_by"`
	VSeqn             *string    `json:"v_seqn"`
}

// aggregate max on columns
type BlacklistOperationRecordMaxFields struct {
	BlacklistRecordID *string    `json:"blacklist_record_id"`
	BlacklistType     *int       `json:"blacklist_type"`
	CreateAt          *time.Time `json:"create_at"`
	CreateBy          *string    `json:"create_by"`
	DeleteAt          *time.Time `json:"delete_at"`
	DeleteBy          *string    `json:"delete_by"`
	ID                *int64     `json:"id"`
	Operate           *int       `json:"operate"`
	Remarks           *string    `json:"remarks"`
	UpdateAt          *time.Time `json:"update_at"`
	UpdateBy          *string    `json:"update_by"`
	VSeqn             *string    `json:"v_seqn"`
}

// order by max() on columns of table "blacklist_operation_record"
type BlacklistOperationRecordMaxOrderBy struct {
	BlacklistRecordID *model1.OrderBy `json:"blacklist_record_id"`
	BlacklistType     *model1.OrderBy `json:"blacklist_type"`
	CreateAt          *model1.OrderBy `json:"create_at"`
	CreateBy          *model1.OrderBy `json:"create_by"`
	DeleteAt          *model1.OrderBy `json:"delete_at"`
	DeleteBy          *model1.OrderBy `json:"delete_by"`
	ID                *model1.OrderBy `json:"id"`
	Operate           *model1.OrderBy `json:"operate"`
	Remarks           *model1.OrderBy `json:"remarks"`
	UpdateAt          *model1.OrderBy `json:"update_at"`
	UpdateBy          *model1.OrderBy `json:"update_by"`
	VSeqn             *model1.OrderBy `json:"v_seqn"`
}

// aggregate min on columns
type BlacklistOperationRecordMinFields struct {
	BlacklistRecordID *string    `json:"blacklist_record_id"`
	BlacklistType     *int       `json:"blacklist_type"`
	CreateAt          *time.Time `json:"create_at"`
	CreateBy          *string    `json:"create_by"`
	DeleteAt          *time.Time `json:"delete_at"`
	DeleteBy          *string    `json:"delete_by"`
	ID                *int64     `json:"id"`
	Operate           *int       `json:"operate"`
	Remarks           *string    `json:"remarks"`
	UpdateAt          *time.Time `json:"update_at"`
	UpdateBy          *string    `json:"update_by"`
	VSeqn             *string    `json:"v_seqn"`
}

// order by min() on columns of table "blacklist_operation_record"
type BlacklistOperationRecordMinOrderBy struct {
	BlacklistRecordID *model1.OrderBy `json:"blacklist_record_id"`
	BlacklistType     *model1.OrderBy `json:"blacklist_type"`
	CreateAt          *model1.OrderBy `json:"create_at"`
	CreateBy          *model1.OrderBy `json:"create_by"`
	DeleteAt          *model1.OrderBy `json:"delete_at"`
	DeleteBy          *model1.OrderBy `json:"delete_by"`
	ID                *model1.OrderBy `json:"id"`
	Operate           *model1.OrderBy `json:"operate"`
	Remarks           *model1.OrderBy `json:"remarks"`
	UpdateAt          *model1.OrderBy `json:"update_at"`
	UpdateBy          *model1.OrderBy `json:"update_by"`
	VSeqn             *model1.OrderBy `json:"v_seqn"`
}

// response of any mutation on the table "blacklist_operation_record"
type BlacklistOperationRecordMutationResponse struct {
	// number of affected rows by the mutation
	AffectedRows int `json:"affected_rows"`
	// data of the affected rows by the mutation
	Returning []*model.BlacklistOperationRecord `json:"returning"`
}

// input type for inserting object relation for remote table "blacklist_operation_record"
type BlacklistOperationRecordObjRelInsertInput struct {
	Data       *BlacklistOperationRecordInsertInput `json:"data"`
	OnConflict *BlacklistOperationRecordOnConflict  `json:"on_conflict"`
}

// on conflict condition type for table "blacklist_operation_record"
type BlacklistOperationRecordOnConflict struct {
	Constraint    BlacklistOperationRecordConstraint     `json:"constraint"`
	UpdateColumns []BlacklistOperationRecordUpdateColumn `json:"update_columns"`
	Where         *BlacklistOperationRecordBoolExp       `json:"where"`
}

// ordering options when selecting data from "blacklist_operation_record"
type BlacklistOperationRecordOrderBy struct {
	BlacklistRecordID *model1.OrderBy `json:"blacklist_record_id"`
	BlacklistType     *model1.OrderBy `json:"blacklist_type"`
	CreateAt          *model1.OrderBy `json:"create_at"`
	CreateBy          *model1.OrderBy `json:"create_by"`
	DeleteAt          *model1.OrderBy `json:"delete_at"`
	DeleteBy          *model1.OrderBy `json:"delete_by"`
	ID                *model1.OrderBy `json:"id"`
	IsDelete          *model1.OrderBy `json:"is_delete"`
	Operate           *model1.OrderBy `json:"operate"`
	Remarks           *model1.OrderBy `json:"remarks"`
	UpdateAt          *model1.OrderBy `json:"update_at"`
	UpdateBy          *model1.OrderBy `json:"update_by"`
	VSeqn             *model1.OrderBy `json:"v_seqn"`
}

// primary key columns input for table: "blacklist_operation_record"
type BlacklistOperationRecordPkColumnsInput struct {
	// ID
	ID int64 `json:"id"`
}

// input type for updating data in table "blacklist_operation_record"
type BlacklistOperationRecordSetInput struct {
	BlacklistRecordID *string    `json:"blacklist_record_id"`
	BlacklistType     *int       `json:"blacklist_type"`
	CreateAt          *time.Time `json:"create_at"`
	CreateBy          *string    `json:"create_by"`
	DeleteAt          *time.Time `json:"delete_at"`
	DeleteBy          *string    `json:"delete_by"`
	ID                *int64     `json:"id"`
	IsDelete          *bool      `json:"is_delete"`
	Operate           *int       `json:"operate"`
	Remarks           *string    `json:"remarks"`
	UpdateAt          *time.Time `json:"update_at"`
	UpdateBy          *string    `json:"update_by"`
	VSeqn             *string    `json:"v_seqn"`
}

// aggregate stddev on columns
type BlacklistOperationRecordStddevFields struct {
	BlacklistType *float64 `json:"blacklist_type"`
	ID            *float64 `json:"id"`
	Operate       *float64 `json:"operate"`
}

// order by stddev() on columns of table "blacklist_operation_record"
type BlacklistOperationRecordStddevOrderBy struct {
	BlacklistType *model1.OrderBy `json:"blacklist_type"`
	ID            *model1.OrderBy `json:"id"`
	Operate       *model1.OrderBy `json:"operate"`
}

// aggregate stddev_pop on columns
type BlacklistOperationRecordStddevPopFields struct {
	BlacklistType *float64 `json:"blacklist_type"`
	ID            *float64 `json:"id"`
	Operate       *float64 `json:"operate"`
}

// order by stddev_pop() on columns of table "blacklist_operation_record"
type BlacklistOperationRecordStddevPopOrderBy struct {
	BlacklistType *model1.OrderBy `json:"blacklist_type"`
	ID            *model1.OrderBy `json:"id"`
	Operate       *model1.OrderBy `json:"operate"`
}

// aggregate stddev_samp on columns
type BlacklistOperationRecordStddevSampFields struct {
	BlacklistType *float64 `json:"blacklist_type"`
	ID            *float64 `json:"id"`
	Operate       *float64 `json:"operate"`
}

// order by stddev_samp() on columns of table "blacklist_operation_record"
type BlacklistOperationRecordStddevSampOrderBy struct {
	BlacklistType *model1.OrderBy `json:"blacklist_type"`
	ID            *model1.OrderBy `json:"id"`
	Operate       *model1.OrderBy `json:"operate"`
}

// aggregate sum on columns
type BlacklistOperationRecordSumFields struct {
	BlacklistType *int   `json:"blacklist_type"`
	ID            *int64 `json:"id"`
	Operate       *int   `json:"operate"`
}

// order by sum() on columns of table "blacklist_operation_record"
type BlacklistOperationRecordSumOrderBy struct {
	BlacklistType *model1.OrderBy `json:"blacklist_type"`
	ID            *model1.OrderBy `json:"id"`
	Operate       *model1.OrderBy `json:"operate"`
}

// aggregate var_pop on columns
type BlacklistOperationRecordVarPopFields struct {
	BlacklistType *float64 `json:"blacklist_type"`
	ID            *float64 `json:"id"`
	Operate       *float64 `json:"operate"`
}

// order by var_pop() on columns of table "blacklist_operation_record"
type BlacklistOperationRecordVarPopOrderBy struct {
	BlacklistType *model1.OrderBy `json:"blacklist_type"`
	ID            *model1.OrderBy `json:"id"`
	Operate       *model1.OrderBy `json:"operate"`
}

// aggregate var_samp on columns
type BlacklistOperationRecordVarSampFields struct {
	BlacklistType *float64 `json:"blacklist_type"`
	ID            *float64 `json:"id"`
	Operate       *float64 `json:"operate"`
}

// order by var_samp() on columns of table "blacklist_operation_record"
type BlacklistOperationRecordVarSampOrderBy struct {
	BlacklistType *model1.OrderBy `json:"blacklist_type"`
	ID            *model1.OrderBy `json:"id"`
	Operate       *model1.OrderBy `json:"operate"`
}

// aggregate variance on columns
type BlacklistOperationRecordVarianceFields struct {
	BlacklistType *float64 `json:"blacklist_type"`
	ID            *float64 `json:"id"`
	Operate       *float64 `json:"operate"`
}

// order by variance() on columns of table "blacklist_operation_record"
type BlacklistOperationRecordVarianceOrderBy struct {
	BlacklistType *model1.OrderBy `json:"blacklist_type"`
	ID            *model1.OrderBy `json:"id"`
	Operate       *model1.OrderBy `json:"operate"`
}

// unique or primary key constraints on table "blacklist_operation_record"
type BlacklistOperationRecordConstraint string

const (
	// unique or primary key constraint
	BlacklistOperationRecordConstraintBlacklistOperationRecordPkey BlacklistOperationRecordConstraint = "blacklist_operation_record_pkey"
)

var AllBlacklistOperationRecordConstraint = []BlacklistOperationRecordConstraint{
	BlacklistOperationRecordConstraintBlacklistOperationRecordPkey,
}

func (e BlacklistOperationRecordConstraint) IsValid() bool {
	switch e {
	case BlacklistOperationRecordConstraintBlacklistOperationRecordPkey:
		return true
	}
	return false
}

func (e BlacklistOperationRecordConstraint) String() string {
	return string(e)
}

func (e *BlacklistOperationRecordConstraint) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = BlacklistOperationRecordConstraint(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid BlacklistOperationRecordConstraint", str)
	}
	return nil
}

func (e BlacklistOperationRecordConstraint) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// select columns of table "blacklist_operation_record"
type BlacklistOperationRecordSelectColumn string

const (
	// column name
	BlacklistOperationRecordSelectColumnBlacklistRecordID BlacklistOperationRecordSelectColumn = "blacklist_record_id"
	// column name
	BlacklistOperationRecordSelectColumnBlacklistType BlacklistOperationRecordSelectColumn = "blacklist_type"
	// column name
	BlacklistOperationRecordSelectColumnCreateAt BlacklistOperationRecordSelectColumn = "create_at"
	// column name
	BlacklistOperationRecordSelectColumnCreateBy BlacklistOperationRecordSelectColumn = "create_by"
	// column name
	BlacklistOperationRecordSelectColumnDeleteAt BlacklistOperationRecordSelectColumn = "delete_at"
	// column name
	BlacklistOperationRecordSelectColumnDeleteBy BlacklistOperationRecordSelectColumn = "delete_by"
	// column name
	BlacklistOperationRecordSelectColumnID BlacklistOperationRecordSelectColumn = "id"
	// column name
	BlacklistOperationRecordSelectColumnIsDelete BlacklistOperationRecordSelectColumn = "is_delete"
	// column name
	BlacklistOperationRecordSelectColumnOperate BlacklistOperationRecordSelectColumn = "operate"
	// column name
	BlacklistOperationRecordSelectColumnRemarks BlacklistOperationRecordSelectColumn = "remarks"
	// column name
	BlacklistOperationRecordSelectColumnUpdateAt BlacklistOperationRecordSelectColumn = "update_at"
	// column name
	BlacklistOperationRecordSelectColumnUpdateBy BlacklistOperationRecordSelectColumn = "update_by"
	// column name
	BlacklistOperationRecordSelectColumnVSeqn BlacklistOperationRecordSelectColumn = "v_seqn"
)

var AllBlacklistOperationRecordSelectColumn = []BlacklistOperationRecordSelectColumn{
	BlacklistOperationRecordSelectColumnBlacklistRecordID,
	BlacklistOperationRecordSelectColumnBlacklistType,
	BlacklistOperationRecordSelectColumnCreateAt,
	BlacklistOperationRecordSelectColumnCreateBy,
	BlacklistOperationRecordSelectColumnDeleteAt,
	BlacklistOperationRecordSelectColumnDeleteBy,
	BlacklistOperationRecordSelectColumnID,
	BlacklistOperationRecordSelectColumnIsDelete,
	BlacklistOperationRecordSelectColumnOperate,
	BlacklistOperationRecordSelectColumnRemarks,
	BlacklistOperationRecordSelectColumnUpdateAt,
	BlacklistOperationRecordSelectColumnUpdateBy,
	BlacklistOperationRecordSelectColumnVSeqn,
}

func (e BlacklistOperationRecordSelectColumn) IsValid() bool {
	switch e {
	case BlacklistOperationRecordSelectColumnBlacklistRecordID, BlacklistOperationRecordSelectColumnBlacklistType, BlacklistOperationRecordSelectColumnCreateAt, BlacklistOperationRecordSelectColumnCreateBy, BlacklistOperationRecordSelectColumnDeleteAt, BlacklistOperationRecordSelectColumnDeleteBy, BlacklistOperationRecordSelectColumnID, BlacklistOperationRecordSelectColumnIsDelete, BlacklistOperationRecordSelectColumnOperate, BlacklistOperationRecordSelectColumnRemarks, BlacklistOperationRecordSelectColumnUpdateAt, BlacklistOperationRecordSelectColumnUpdateBy, BlacklistOperationRecordSelectColumnVSeqn:
		return true
	}
	return false
}

func (e BlacklistOperationRecordSelectColumn) String() string {
	return string(e)
}

func (e *BlacklistOperationRecordSelectColumn) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = BlacklistOperationRecordSelectColumn(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid BlacklistOperationRecordSelectColumn", str)
	}
	return nil
}

func (e BlacklistOperationRecordSelectColumn) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// update columns of table "blacklist_operation_record"
type BlacklistOperationRecordUpdateColumn string

const (
	// column name
	BlacklistOperationRecordUpdateColumnBlacklistRecordID BlacklistOperationRecordUpdateColumn = "blacklist_record_id"
	// column name
	BlacklistOperationRecordUpdateColumnBlacklistType BlacklistOperationRecordUpdateColumn = "blacklist_type"
	// column name
	BlacklistOperationRecordUpdateColumnCreateAt BlacklistOperationRecordUpdateColumn = "create_at"
	// column name
	BlacklistOperationRecordUpdateColumnCreateBy BlacklistOperationRecordUpdateColumn = "create_by"
	// column name
	BlacklistOperationRecordUpdateColumnDeleteAt BlacklistOperationRecordUpdateColumn = "delete_at"
	// column name
	BlacklistOperationRecordUpdateColumnDeleteBy BlacklistOperationRecordUpdateColumn = "delete_by"
	// column name
	BlacklistOperationRecordUpdateColumnID BlacklistOperationRecordUpdateColumn = "id"
	// column name
	BlacklistOperationRecordUpdateColumnIsDelete BlacklistOperationRecordUpdateColumn = "is_delete"
	// column name
	BlacklistOperationRecordUpdateColumnOperate BlacklistOperationRecordUpdateColumn = "operate"
	// column name
	BlacklistOperationRecordUpdateColumnRemarks BlacklistOperationRecordUpdateColumn = "remarks"
	// column name
	BlacklistOperationRecordUpdateColumnUpdateAt BlacklistOperationRecordUpdateColumn = "update_at"
	// column name
	BlacklistOperationRecordUpdateColumnUpdateBy BlacklistOperationRecordUpdateColumn = "update_by"
	// column name
	BlacklistOperationRecordUpdateColumnVSeqn BlacklistOperationRecordUpdateColumn = "v_seqn"
)

var AllBlacklistOperationRecordUpdateColumn = []BlacklistOperationRecordUpdateColumn{
	BlacklistOperationRecordUpdateColumnBlacklistRecordID,
	BlacklistOperationRecordUpdateColumnBlacklistType,
	BlacklistOperationRecordUpdateColumnCreateAt,
	BlacklistOperationRecordUpdateColumnCreateBy,
	BlacklistOperationRecordUpdateColumnDeleteAt,
	BlacklistOperationRecordUpdateColumnDeleteBy,
	BlacklistOperationRecordUpdateColumnID,
	BlacklistOperationRecordUpdateColumnIsDelete,
	BlacklistOperationRecordUpdateColumnOperate,
	BlacklistOperationRecordUpdateColumnRemarks,
	BlacklistOperationRecordUpdateColumnUpdateAt,
	BlacklistOperationRecordUpdateColumnUpdateBy,
	BlacklistOperationRecordUpdateColumnVSeqn,
}

func (e BlacklistOperationRecordUpdateColumn) IsValid() bool {
	switch e {
	case BlacklistOperationRecordUpdateColumnBlacklistRecordID, BlacklistOperationRecordUpdateColumnBlacklistType, BlacklistOperationRecordUpdateColumnCreateAt, BlacklistOperationRecordUpdateColumnCreateBy, BlacklistOperationRecordUpdateColumnDeleteAt, BlacklistOperationRecordUpdateColumnDeleteBy, BlacklistOperationRecordUpdateColumnID, BlacklistOperationRecordUpdateColumnIsDelete, BlacklistOperationRecordUpdateColumnOperate, BlacklistOperationRecordUpdateColumnRemarks, BlacklistOperationRecordUpdateColumnUpdateAt, BlacklistOperationRecordUpdateColumnUpdateBy, BlacklistOperationRecordUpdateColumnVSeqn:
		return true
	}
	return false
}

func (e BlacklistOperationRecordUpdateColumn) String() string {
	return string(e)
}

func (e *BlacklistOperationRecordUpdateColumn) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = BlacklistOperationRecordUpdateColumn(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid BlacklistOperationRecordUpdateColumn", str)
	}
	return nil
}

func (e BlacklistOperationRecordUpdateColumn) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
