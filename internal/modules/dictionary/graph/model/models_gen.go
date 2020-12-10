// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"VehicleSupervision/internal/modules/dictionary/model"
	model1 "VehicleSupervision/pkg/graphql/model"
	"fmt"
	"io"
	"strconv"
	"time"
)

// aggregated selection of "data_dictionary"
type DataDictionaryAggregate struct {
	Aggregate *DataDictionaryAggregateFields `json:"aggregate"`
	Nodes     []*model.DataDictionary        `json:"nodes"`
}

// aggregate fields of "data_dictionary"
type DataDictionaryAggregateFields struct {
	Avg        *DataDictionaryAvgFields        `json:"avg"`
	Count      *int                            `json:"count"`
	Max        *DataDictionaryMaxFields        `json:"max"`
	Min        *DataDictionaryMinFields        `json:"min"`
	Stddev     *DataDictionaryStddevFields     `json:"stddev"`
	StddevPop  *DataDictionaryStddevPopFields  `json:"stddev_pop"`
	StddevSamp *DataDictionaryStddevSampFields `json:"stddev_samp"`
	Sum        *DataDictionarySumFields        `json:"sum"`
	VarPop     *DataDictionaryVarPopFields     `json:"var_pop"`
	VarSamp    *DataDictionaryVarSampFields    `json:"var_samp"`
	Variance   *DataDictionaryVarianceFields   `json:"variance"`
}

// order by aggregate values of table "data_dictionary"
type DataDictionaryAggregateOrderBy struct {
	Avg        *DataDictionaryAvgOrderBy        `json:"avg"`
	Count      *model1.OrderBy                  `json:"count"`
	Max        *DataDictionaryMaxOrderBy        `json:"max"`
	Min        *DataDictionaryMinOrderBy        `json:"min"`
	Stddev     *DataDictionaryStddevOrderBy     `json:"stddev"`
	StddevPop  *DataDictionaryStddevPopOrderBy  `json:"stddev_pop"`
	StddevSamp *DataDictionaryStddevSampOrderBy `json:"stddev_samp"`
	Sum        *DataDictionarySumOrderBy        `json:"sum"`
	VarPop     *DataDictionaryVarPopOrderBy     `json:"var_pop"`
	VarSamp    *DataDictionaryVarSampOrderBy    `json:"var_samp"`
	Variance   *DataDictionaryVarianceOrderBy   `json:"variance"`
}

// input type for inserting array relation for remote table "data_dictionary"
type DataDictionaryArrRelInsertInput struct {
	Data       []*DataDictionaryInsertInput `json:"data"`
	OnConflict *DataDictionaryOnConflict    `json:"on_conflict"`
}

// aggregate avg on columns
type DataDictionaryAvgFields struct {
	ID    *float64 `json:"id"`
	Value *float64 `json:"value"`
}

// order by avg() on columns of table "data_dictionary"
type DataDictionaryAvgOrderBy struct {
	ID    *model1.OrderBy `json:"id"`
	Value *model1.OrderBy `json:"value"`
}

// Boolean expression to filter rows from the table "data_dictionary". All fields are combined with a logical 'AND'.
type DataDictionaryBoolExp struct {
	And                  []*DataDictionaryBoolExp         `json:"_and"`
	Not                  *DataDictionaryBoolExp           `json:"_not"`
	Or                   []*DataDictionaryBoolExp         `json:"_or"`
	CreateAt             *model1.TimestamptzComparisonExp `json:"create_at"`
	CreateBy             *model1.StringComparisonExp      `json:"create_by"`
	DeleteAt             *model1.TimestamptzComparisonExp `json:"delete_at"`
	DeleteBy             *model1.StringComparisonExp      `json:"delete_by"`
	DictionaryCategoryID *model1.StringComparisonExp      `json:"dictionary_category_id"`
	DictionaryID         *model1.StringComparisonExp      `json:"dictionary_id"`
	ID                   *model1.BigintComparisonExp      `json:"id"`
	IsDelete             *model1.BooleanComparisonExp     `json:"is_delete"`
	Name                 *model1.StringComparisonExp      `json:"name"`
	Remarks              *model1.StringComparisonExp      `json:"remarks"`
	UpdateAt             *model1.TimestamptzComparisonExp `json:"update_at"`
	UpdateBy             *model1.StringComparisonExp      `json:"update_by"`
	Value                *model1.IntComparisonExp         `json:"value"`
}

// aggregated selection of "data_dictionary_category"
type DataDictionaryCategoryAggregate struct {
	Aggregate *DataDictionaryCategoryAggregateFields `json:"aggregate"`
	Nodes     []*model.DataDictionaryCategory        `json:"nodes"`
}

// aggregate fields of "data_dictionary_category"
type DataDictionaryCategoryAggregateFields struct {
	Avg        *DataDictionaryCategoryAvgFields        `json:"avg"`
	Count      *int                                    `json:"count"`
	Max        *DataDictionaryCategoryMaxFields        `json:"max"`
	Min        *DataDictionaryCategoryMinFields        `json:"min"`
	Stddev     *DataDictionaryCategoryStddevFields     `json:"stddev"`
	StddevPop  *DataDictionaryCategoryStddevPopFields  `json:"stddev_pop"`
	StddevSamp *DataDictionaryCategoryStddevSampFields `json:"stddev_samp"`
	Sum        *DataDictionaryCategorySumFields        `json:"sum"`
	VarPop     *DataDictionaryCategoryVarPopFields     `json:"var_pop"`
	VarSamp    *DataDictionaryCategoryVarSampFields    `json:"var_samp"`
	Variance   *DataDictionaryCategoryVarianceFields   `json:"variance"`
}

// order by aggregate values of table "data_dictionary_category"
type DataDictionaryCategoryAggregateOrderBy struct {
	Avg        *DataDictionaryCategoryAvgOrderBy        `json:"avg"`
	Count      *model1.OrderBy                          `json:"count"`
	Max        *DataDictionaryCategoryMaxOrderBy        `json:"max"`
	Min        *DataDictionaryCategoryMinOrderBy        `json:"min"`
	Stddev     *DataDictionaryCategoryStddevOrderBy     `json:"stddev"`
	StddevPop  *DataDictionaryCategoryStddevPopOrderBy  `json:"stddev_pop"`
	StddevSamp *DataDictionaryCategoryStddevSampOrderBy `json:"stddev_samp"`
	Sum        *DataDictionaryCategorySumOrderBy        `json:"sum"`
	VarPop     *DataDictionaryCategoryVarPopOrderBy     `json:"var_pop"`
	VarSamp    *DataDictionaryCategoryVarSampOrderBy    `json:"var_samp"`
	Variance   *DataDictionaryCategoryVarianceOrderBy   `json:"variance"`
}

// input type for inserting array relation for remote table "data_dictionary_category"
type DataDictionaryCategoryArrRelInsertInput struct {
	Data       []*DataDictionaryCategoryInsertInput `json:"data"`
	OnConflict *DataDictionaryCategoryOnConflict    `json:"on_conflict"`
}

// aggregate avg on columns
type DataDictionaryCategoryAvgFields struct {
	ID *float64 `json:"id"`
}

// order by avg() on columns of table "data_dictionary_category"
type DataDictionaryCategoryAvgOrderBy struct {
	ID *model1.OrderBy `json:"id"`
}

// Boolean expression to filter rows from the table "data_dictionary_category". All fields are combined with a logical 'AND'.
type DataDictionaryCategoryBoolExp struct {
	And                  []*DataDictionaryCategoryBoolExp `json:"_and"`
	Not                  *DataDictionaryCategoryBoolExp   `json:"_not"`
	Or                   []*DataDictionaryCategoryBoolExp `json:"_or"`
	CategoryCode         *model1.StringComparisonExp      `json:"category_code"`
	CategoryName         *model1.StringComparisonExp      `json:"category_name"`
	CreateAt             *model1.TimestamptzComparisonExp `json:"create_at"`
	CreateBy             *model1.StringComparisonExp      `json:"create_by"`
	DeleteAt             *model1.TimestamptzComparisonExp `json:"delete_at"`
	DeleteBy             *model1.StringComparisonExp      `json:"delete_by"`
	DictionaryCategoryID *model1.StringComparisonExp      `json:"dictionary_category_id"`
	ID                   *model1.BigintComparisonExp      `json:"id"`
	IsDelete             *model1.BooleanComparisonExp     `json:"is_delete"`
	Remarks              *model1.StringComparisonExp      `json:"remarks"`
	UpdateAt             *model1.TimestamptzComparisonExp `json:"update_at"`
	UpdateBy             *model1.StringComparisonExp      `json:"update_by"`
}

// input type for incrementing integer column in table "data_dictionary_category"
type DataDictionaryCategoryIncInput struct {
	ID *int64 `json:"id"`
}

// input type for inserting data into table "data_dictionary_category"
type DataDictionaryCategoryInsertInput struct {
	CategoryCode         *string    `json:"category_code"`
	CategoryName         *string    `json:"category_name"`
	CreateAt             *time.Time `json:"create_at"`
	CreateBy             *string    `json:"create_by"`
	DeleteAt             *time.Time `json:"delete_at"`
	DeleteBy             *string    `json:"delete_by"`
	DictionaryCategoryID *string    `json:"dictionary_category_id"`
	ID                   *int64     `json:"id"`
	IsDelete             *bool      `json:"is_delete"`
	Remarks              *string    `json:"remarks"`
	UpdateAt             *time.Time `json:"update_at"`
	UpdateBy             *string    `json:"update_by"`
}

// aggregate max on columns
type DataDictionaryCategoryMaxFields struct {
	CategoryCode         *string    `json:"category_code"`
	CategoryName         *string    `json:"category_name"`
	CreateAt             *time.Time `json:"create_at"`
	CreateBy             *string    `json:"create_by"`
	DeleteAt             *time.Time `json:"delete_at"`
	DeleteBy             *string    `json:"delete_by"`
	DictionaryCategoryID *string    `json:"dictionary_category_id"`
	ID                   *int64     `json:"id"`
	Remarks              *string    `json:"remarks"`
	UpdateAt             *time.Time `json:"update_at"`
	UpdateBy             *string    `json:"update_by"`
}

// order by max() on columns of table "data_dictionary_category"
type DataDictionaryCategoryMaxOrderBy struct {
	CategoryCode         *model1.OrderBy `json:"category_code"`
	CategoryName         *model1.OrderBy `json:"category_name"`
	CreateAt             *model1.OrderBy `json:"create_at"`
	CreateBy             *model1.OrderBy `json:"create_by"`
	DeleteAt             *model1.OrderBy `json:"delete_at"`
	DeleteBy             *model1.OrderBy `json:"delete_by"`
	DictionaryCategoryID *model1.OrderBy `json:"dictionary_category_id"`
	ID                   *model1.OrderBy `json:"id"`
	Remarks              *model1.OrderBy `json:"remarks"`
	UpdateAt             *model1.OrderBy `json:"update_at"`
	UpdateBy             *model1.OrderBy `json:"update_by"`
}

// aggregate min on columns
type DataDictionaryCategoryMinFields struct {
	CategoryCode         *string    `json:"category_code"`
	CategoryName         *string    `json:"category_name"`
	CreateAt             *time.Time `json:"create_at"`
	CreateBy             *string    `json:"create_by"`
	DeleteAt             *time.Time `json:"delete_at"`
	DeleteBy             *string    `json:"delete_by"`
	DictionaryCategoryID *string    `json:"dictionary_category_id"`
	ID                   *int64     `json:"id"`
	Remarks              *string    `json:"remarks"`
	UpdateAt             *time.Time `json:"update_at"`
	UpdateBy             *string    `json:"update_by"`
}

// order by min() on columns of table "data_dictionary_category"
type DataDictionaryCategoryMinOrderBy struct {
	CategoryCode         *model1.OrderBy `json:"category_code"`
	CategoryName         *model1.OrderBy `json:"category_name"`
	CreateAt             *model1.OrderBy `json:"create_at"`
	CreateBy             *model1.OrderBy `json:"create_by"`
	DeleteAt             *model1.OrderBy `json:"delete_at"`
	DeleteBy             *model1.OrderBy `json:"delete_by"`
	DictionaryCategoryID *model1.OrderBy `json:"dictionary_category_id"`
	ID                   *model1.OrderBy `json:"id"`
	Remarks              *model1.OrderBy `json:"remarks"`
	UpdateAt             *model1.OrderBy `json:"update_at"`
	UpdateBy             *model1.OrderBy `json:"update_by"`
}

// response of any mutation on the table "data_dictionary_category"
type DataDictionaryCategoryMutationResponse struct {
	// number of affected rows by the mutation
	AffectedRows int `json:"affected_rows"`
	// data of the affected rows by the mutation
	Returning []*model.DataDictionaryCategory `json:"returning"`
}

// input type for inserting object relation for remote table "data_dictionary_category"
type DataDictionaryCategoryObjRelInsertInput struct {
	Data       *DataDictionaryCategoryInsertInput `json:"data"`
	OnConflict *DataDictionaryCategoryOnConflict  `json:"on_conflict"`
}

// on conflict condition type for table "data_dictionary_category"
type DataDictionaryCategoryOnConflict struct {
	Constraint    DataDictionaryCategoryConstraint     `json:"constraint"`
	UpdateColumns []DataDictionaryCategoryUpdateColumn `json:"update_columns"`
	Where         *DataDictionaryCategoryBoolExp       `json:"where"`
}

// ordering options when selecting data from "data_dictionary_category"
type DataDictionaryCategoryOrderBy struct {
	CategoryCode         *model1.OrderBy `json:"category_code"`
	CategoryName         *model1.OrderBy `json:"category_name"`
	CreateAt             *model1.OrderBy `json:"create_at"`
	CreateBy             *model1.OrderBy `json:"create_by"`
	DeleteAt             *model1.OrderBy `json:"delete_at"`
	DeleteBy             *model1.OrderBy `json:"delete_by"`
	DictionaryCategoryID *model1.OrderBy `json:"dictionary_category_id"`
	ID                   *model1.OrderBy `json:"id"`
	IsDelete             *model1.OrderBy `json:"is_delete"`
	Remarks              *model1.OrderBy `json:"remarks"`
	UpdateAt             *model1.OrderBy `json:"update_at"`
	UpdateBy             *model1.OrderBy `json:"update_by"`
}

// primary key columns input for table: "data_dictionary_category"
type DataDictionaryCategoryPkColumnsInput struct {
	// ID
	ID int64 `json:"id"`
}

// input type for updating data in table "data_dictionary_category"
type DataDictionaryCategorySetInput struct {
	CategoryCode         *string    `json:"category_code"`
	CategoryName         *string    `json:"category_name"`
	CreateAt             *time.Time `json:"create_at"`
	CreateBy             *string    `json:"create_by"`
	DeleteAt             *time.Time `json:"delete_at"`
	DeleteBy             *string    `json:"delete_by"`
	DictionaryCategoryID *string    `json:"dictionary_category_id"`
	ID                   *int64     `json:"id"`
	IsDelete             *bool      `json:"is_delete"`
	Remarks              *string    `json:"remarks"`
	UpdateAt             *time.Time `json:"update_at"`
	UpdateBy             *string    `json:"update_by"`
}

// aggregate stddev on columns
type DataDictionaryCategoryStddevFields struct {
	ID *float64 `json:"id"`
}

// order by stddev() on columns of table "data_dictionary_category"
type DataDictionaryCategoryStddevOrderBy struct {
	ID *model1.OrderBy `json:"id"`
}

// aggregate stddev_pop on columns
type DataDictionaryCategoryStddevPopFields struct {
	ID *float64 `json:"id"`
}

// order by stddev_pop() on columns of table "data_dictionary_category"
type DataDictionaryCategoryStddevPopOrderBy struct {
	ID *model1.OrderBy `json:"id"`
}

// aggregate stddev_samp on columns
type DataDictionaryCategoryStddevSampFields struct {
	ID *float64 `json:"id"`
}

// order by stddev_samp() on columns of table "data_dictionary_category"
type DataDictionaryCategoryStddevSampOrderBy struct {
	ID *model1.OrderBy `json:"id"`
}

// aggregate sum on columns
type DataDictionaryCategorySumFields struct {
	ID *int64 `json:"id"`
}

// order by sum() on columns of table "data_dictionary_category"
type DataDictionaryCategorySumOrderBy struct {
	ID *model1.OrderBy `json:"id"`
}

// aggregate var_pop on columns
type DataDictionaryCategoryVarPopFields struct {
	ID *float64 `json:"id"`
}

// order by var_pop() on columns of table "data_dictionary_category"
type DataDictionaryCategoryVarPopOrderBy struct {
	ID *model1.OrderBy `json:"id"`
}

// aggregate var_samp on columns
type DataDictionaryCategoryVarSampFields struct {
	ID *float64 `json:"id"`
}

// order by var_samp() on columns of table "data_dictionary_category"
type DataDictionaryCategoryVarSampOrderBy struct {
	ID *model1.OrderBy `json:"id"`
}

// aggregate variance on columns
type DataDictionaryCategoryVarianceFields struct {
	ID *float64 `json:"id"`
}

// order by variance() on columns of table "data_dictionary_category"
type DataDictionaryCategoryVarianceOrderBy struct {
	ID *model1.OrderBy `json:"id"`
}

// input type for incrementing integer column in table "data_dictionary"
type DataDictionaryIncInput struct {
	ID    *int64 `json:"id"`
	Value *int   `json:"value"`
}

// input type for inserting data into table "data_dictionary"
type DataDictionaryInsertInput struct {
	CreateAt             *time.Time `json:"create_at"`
	CreateBy             *string    `json:"create_by"`
	DeleteAt             *time.Time `json:"delete_at"`
	DeleteBy             *string    `json:"delete_by"`
	DictionaryCategoryID *string    `json:"dictionary_category_id"`
	DictionaryID         *string    `json:"dictionary_id"`
	ID                   *int64     `json:"id"`
	IsDelete             *bool      `json:"is_delete"`
	Name                 *string    `json:"name"`
	Remarks              *string    `json:"remarks"`
	UpdateAt             *time.Time `json:"update_at"`
	UpdateBy             *string    `json:"update_by"`
	Value                *int       `json:"value"`
}

// aggregate max on columns
type DataDictionaryMaxFields struct {
	CreateAt             *time.Time `json:"create_at"`
	CreateBy             *string    `json:"create_by"`
	DeleteAt             *time.Time `json:"delete_at"`
	DeleteBy             *string    `json:"delete_by"`
	DictionaryCategoryID *string    `json:"dictionary_category_id"`
	DictionaryID         *string    `json:"dictionary_id"`
	ID                   *int64     `json:"id"`
	Name                 *string    `json:"name"`
	Remarks              *string    `json:"remarks"`
	UpdateAt             *time.Time `json:"update_at"`
	UpdateBy             *string    `json:"update_by"`
	Value                *int       `json:"value"`
}

// order by max() on columns of table "data_dictionary"
type DataDictionaryMaxOrderBy struct {
	CreateAt             *model1.OrderBy `json:"create_at"`
	CreateBy             *model1.OrderBy `json:"create_by"`
	DeleteAt             *model1.OrderBy `json:"delete_at"`
	DeleteBy             *model1.OrderBy `json:"delete_by"`
	DictionaryCategoryID *model1.OrderBy `json:"dictionary_category_id"`
	DictionaryID         *model1.OrderBy `json:"dictionary_id"`
	ID                   *model1.OrderBy `json:"id"`
	Name                 *model1.OrderBy `json:"name"`
	Remarks              *model1.OrderBy `json:"remarks"`
	UpdateAt             *model1.OrderBy `json:"update_at"`
	UpdateBy             *model1.OrderBy `json:"update_by"`
	Value                *model1.OrderBy `json:"value"`
}

// aggregate min on columns
type DataDictionaryMinFields struct {
	CreateAt             *time.Time `json:"create_at"`
	CreateBy             *string    `json:"create_by"`
	DeleteAt             *time.Time `json:"delete_at"`
	DeleteBy             *string    `json:"delete_by"`
	DictionaryCategoryID *string    `json:"dictionary_category_id"`
	DictionaryID         *string    `json:"dictionary_id"`
	ID                   *int64     `json:"id"`
	Name                 *string    `json:"name"`
	Remarks              *string    `json:"remarks"`
	UpdateAt             *time.Time `json:"update_at"`
	UpdateBy             *string    `json:"update_by"`
	Value                *int       `json:"value"`
}

// order by min() on columns of table "data_dictionary"
type DataDictionaryMinOrderBy struct {
	CreateAt             *model1.OrderBy `json:"create_at"`
	CreateBy             *model1.OrderBy `json:"create_by"`
	DeleteAt             *model1.OrderBy `json:"delete_at"`
	DeleteBy             *model1.OrderBy `json:"delete_by"`
	DictionaryCategoryID *model1.OrderBy `json:"dictionary_category_id"`
	DictionaryID         *model1.OrderBy `json:"dictionary_id"`
	ID                   *model1.OrderBy `json:"id"`
	Name                 *model1.OrderBy `json:"name"`
	Remarks              *model1.OrderBy `json:"remarks"`
	UpdateAt             *model1.OrderBy `json:"update_at"`
	UpdateBy             *model1.OrderBy `json:"update_by"`
	Value                *model1.OrderBy `json:"value"`
}

// response of any mutation on the table "data_dictionary"
type DataDictionaryMutationResponse struct {
	// number of affected rows by the mutation
	AffectedRows int `json:"affected_rows"`
	// data of the affected rows by the mutation
	Returning []*model.DataDictionary `json:"returning"`
}

// input type for inserting object relation for remote table "data_dictionary"
type DataDictionaryObjRelInsertInput struct {
	Data       *DataDictionaryInsertInput `json:"data"`
	OnConflict *DataDictionaryOnConflict  `json:"on_conflict"`
}

// on conflict condition type for table "data_dictionary"
type DataDictionaryOnConflict struct {
	Constraint    DataDictionaryConstraint     `json:"constraint"`
	UpdateColumns []DataDictionaryUpdateColumn `json:"update_columns"`
	Where         *DataDictionaryBoolExp       `json:"where"`
}

// ordering options when selecting data from "data_dictionary"
type DataDictionaryOrderBy struct {
	CreateAt             *model1.OrderBy `json:"create_at"`
	CreateBy             *model1.OrderBy `json:"create_by"`
	DeleteAt             *model1.OrderBy `json:"delete_at"`
	DeleteBy             *model1.OrderBy `json:"delete_by"`
	DictionaryCategoryID *model1.OrderBy `json:"dictionary_category_id"`
	DictionaryID         *model1.OrderBy `json:"dictionary_id"`
	ID                   *model1.OrderBy `json:"id"`
	IsDelete             *model1.OrderBy `json:"is_delete"`
	Name                 *model1.OrderBy `json:"name"`
	Remarks              *model1.OrderBy `json:"remarks"`
	UpdateAt             *model1.OrderBy `json:"update_at"`
	UpdateBy             *model1.OrderBy `json:"update_by"`
	Value                *model1.OrderBy `json:"value"`
}

// primary key columns input for table: "data_dictionary"
type DataDictionaryPkColumnsInput struct {
	// ID
	ID int64 `json:"id"`
}

// input type for updating data in table "data_dictionary"
type DataDictionarySetInput struct {
	CreateAt             *time.Time `json:"create_at"`
	CreateBy             *string    `json:"create_by"`
	DeleteAt             *time.Time `json:"delete_at"`
	DeleteBy             *string    `json:"delete_by"`
	DictionaryCategoryID *string    `json:"dictionary_category_id"`
	DictionaryID         *string    `json:"dictionary_id"`
	ID                   *int64     `json:"id"`
	IsDelete             *bool      `json:"is_delete"`
	Name                 *string    `json:"name"`
	Remarks              *string    `json:"remarks"`
	UpdateAt             *time.Time `json:"update_at"`
	UpdateBy             *string    `json:"update_by"`
	Value                *int       `json:"value"`
}

// aggregate stddev on columns
type DataDictionaryStddevFields struct {
	ID    *float64 `json:"id"`
	Value *float64 `json:"value"`
}

// order by stddev() on columns of table "data_dictionary"
type DataDictionaryStddevOrderBy struct {
	ID    *model1.OrderBy `json:"id"`
	Value *model1.OrderBy `json:"value"`
}

// aggregate stddev_pop on columns
type DataDictionaryStddevPopFields struct {
	ID    *float64 `json:"id"`
	Value *float64 `json:"value"`
}

// order by stddev_pop() on columns of table "data_dictionary"
type DataDictionaryStddevPopOrderBy struct {
	ID    *model1.OrderBy `json:"id"`
	Value *model1.OrderBy `json:"value"`
}

// aggregate stddev_samp on columns
type DataDictionaryStddevSampFields struct {
	ID    *float64 `json:"id"`
	Value *float64 `json:"value"`
}

// order by stddev_samp() on columns of table "data_dictionary"
type DataDictionaryStddevSampOrderBy struct {
	ID    *model1.OrderBy `json:"id"`
	Value *model1.OrderBy `json:"value"`
}

// aggregate sum on columns
type DataDictionarySumFields struct {
	ID    *int64 `json:"id"`
	Value *int   `json:"value"`
}

// order by sum() on columns of table "data_dictionary"
type DataDictionarySumOrderBy struct {
	ID    *model1.OrderBy `json:"id"`
	Value *model1.OrderBy `json:"value"`
}

// aggregate var_pop on columns
type DataDictionaryVarPopFields struct {
	ID    *float64 `json:"id"`
	Value *float64 `json:"value"`
}

// order by var_pop() on columns of table "data_dictionary"
type DataDictionaryVarPopOrderBy struct {
	ID    *model1.OrderBy `json:"id"`
	Value *model1.OrderBy `json:"value"`
}

// aggregate var_samp on columns
type DataDictionaryVarSampFields struct {
	ID    *float64 `json:"id"`
	Value *float64 `json:"value"`
}

// order by var_samp() on columns of table "data_dictionary"
type DataDictionaryVarSampOrderBy struct {
	ID    *model1.OrderBy `json:"id"`
	Value *model1.OrderBy `json:"value"`
}

// aggregate variance on columns
type DataDictionaryVarianceFields struct {
	ID    *float64 `json:"id"`
	Value *float64 `json:"value"`
}

// order by variance() on columns of table "data_dictionary"
type DataDictionaryVarianceOrderBy struct {
	ID    *model1.OrderBy `json:"id"`
	Value *model1.OrderBy `json:"value"`
}

// expression to compare columns of type numeric. All fields are combined with logical 'AND'.
type NumericComparisonExp struct {
	Eq     *float64  `json:"_eq"`
	Gt     *float64  `json:"_gt"`
	Gte    *float64  `json:"_gte"`
	In     []float64 `json:"_in"`
	IsNull *bool     `json:"_is_null"`
	Lt     *float64  `json:"_lt"`
	Lte    *float64  `json:"_lte"`
	Neq    *float64  `json:"_neq"`
	Nin    []float64 `json:"_nin"`
}

// expression to compare columns of type point. All fields are combined with logical 'AND'.
type PointComparisonExp struct {
	Eq     *string  `json:"_eq"`
	Gt     *string  `json:"_gt"`
	Gte    *string  `json:"_gte"`
	In     []string `json:"_in"`
	IsNull *bool    `json:"_is_null"`
	Lt     *string  `json:"_lt"`
	Lte    *string  `json:"_lte"`
	Neq    *string  `json:"_neq"`
	Nin    []string `json:"_nin"`
}

// unique or primary key constraints on table "data_dictionary_category"
type DataDictionaryCategoryConstraint string

const (
	// unique or primary key constraint
	DataDictionaryCategoryConstraintDataDirtionaryCategoryPkey DataDictionaryCategoryConstraint = "data_dirtionary_category_pkey"
)

var AllDataDictionaryCategoryConstraint = []DataDictionaryCategoryConstraint{
	DataDictionaryCategoryConstraintDataDirtionaryCategoryPkey,
}

func (e DataDictionaryCategoryConstraint) IsValid() bool {
	switch e {
	case DataDictionaryCategoryConstraintDataDirtionaryCategoryPkey:
		return true
	}
	return false
}

func (e DataDictionaryCategoryConstraint) String() string {
	return string(e)
}

func (e *DataDictionaryCategoryConstraint) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DataDictionaryCategoryConstraint(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid data_dictionary_category_constraint", str)
	}
	return nil
}

func (e DataDictionaryCategoryConstraint) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// select columns of table "data_dictionary_category"
type DataDictionaryCategorySelectColumn string

const (
	// column name
	DataDictionaryCategorySelectColumnCategoryCode DataDictionaryCategorySelectColumn = "category_code"
	// column name
	DataDictionaryCategorySelectColumnCategoryName DataDictionaryCategorySelectColumn = "category_name"
	// column name
	DataDictionaryCategorySelectColumnCreateAt DataDictionaryCategorySelectColumn = "create_at"
	// column name
	DataDictionaryCategorySelectColumnCreateBy DataDictionaryCategorySelectColumn = "create_by"
	// column name
	DataDictionaryCategorySelectColumnDeleteAt DataDictionaryCategorySelectColumn = "delete_at"
	// column name
	DataDictionaryCategorySelectColumnDeleteBy DataDictionaryCategorySelectColumn = "delete_by"
	// column name
	DataDictionaryCategorySelectColumnDictionaryCategoryID DataDictionaryCategorySelectColumn = "dictionary_category_id"
	// column name
	DataDictionaryCategorySelectColumnID DataDictionaryCategorySelectColumn = "id"
	// column name
	DataDictionaryCategorySelectColumnIsDelete DataDictionaryCategorySelectColumn = "is_delete"
	// column name
	DataDictionaryCategorySelectColumnRemarks DataDictionaryCategorySelectColumn = "remarks"
	// column name
	DataDictionaryCategorySelectColumnUpdateAt DataDictionaryCategorySelectColumn = "update_at"
	// column name
	DataDictionaryCategorySelectColumnUpdateBy DataDictionaryCategorySelectColumn = "update_by"
)

var AllDataDictionaryCategorySelectColumn = []DataDictionaryCategorySelectColumn{
	DataDictionaryCategorySelectColumnCategoryCode,
	DataDictionaryCategorySelectColumnCategoryName,
	DataDictionaryCategorySelectColumnCreateAt,
	DataDictionaryCategorySelectColumnCreateBy,
	DataDictionaryCategorySelectColumnDeleteAt,
	DataDictionaryCategorySelectColumnDeleteBy,
	DataDictionaryCategorySelectColumnDictionaryCategoryID,
	DataDictionaryCategorySelectColumnID,
	DataDictionaryCategorySelectColumnIsDelete,
	DataDictionaryCategorySelectColumnRemarks,
	DataDictionaryCategorySelectColumnUpdateAt,
	DataDictionaryCategorySelectColumnUpdateBy,
}

func (e DataDictionaryCategorySelectColumn) IsValid() bool {
	switch e {
	case DataDictionaryCategorySelectColumnCategoryCode, DataDictionaryCategorySelectColumnCategoryName, DataDictionaryCategorySelectColumnCreateAt, DataDictionaryCategorySelectColumnCreateBy, DataDictionaryCategorySelectColumnDeleteAt, DataDictionaryCategorySelectColumnDeleteBy, DataDictionaryCategorySelectColumnDictionaryCategoryID, DataDictionaryCategorySelectColumnID, DataDictionaryCategorySelectColumnIsDelete, DataDictionaryCategorySelectColumnRemarks, DataDictionaryCategorySelectColumnUpdateAt, DataDictionaryCategorySelectColumnUpdateBy:
		return true
	}
	return false
}

func (e DataDictionaryCategorySelectColumn) String() string {
	return string(e)
}

func (e *DataDictionaryCategorySelectColumn) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DataDictionaryCategorySelectColumn(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid data_dictionary_category_select_column", str)
	}
	return nil
}

func (e DataDictionaryCategorySelectColumn) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// update columns of table "data_dictionary_category"
type DataDictionaryCategoryUpdateColumn string

const (
	// column name
	DataDictionaryCategoryUpdateColumnCategoryCode DataDictionaryCategoryUpdateColumn = "category_code"
	// column name
	DataDictionaryCategoryUpdateColumnCategoryName DataDictionaryCategoryUpdateColumn = "category_name"
	// column name
	DataDictionaryCategoryUpdateColumnCreateAt DataDictionaryCategoryUpdateColumn = "create_at"
	// column name
	DataDictionaryCategoryUpdateColumnCreateBy DataDictionaryCategoryUpdateColumn = "create_by"
	// column name
	DataDictionaryCategoryUpdateColumnDeleteAt DataDictionaryCategoryUpdateColumn = "delete_at"
	// column name
	DataDictionaryCategoryUpdateColumnDeleteBy DataDictionaryCategoryUpdateColumn = "delete_by"
	// column name
	DataDictionaryCategoryUpdateColumnDictionaryCategoryID DataDictionaryCategoryUpdateColumn = "dictionary_category_id"
	// column name
	DataDictionaryCategoryUpdateColumnID DataDictionaryCategoryUpdateColumn = "id"
	// column name
	DataDictionaryCategoryUpdateColumnIsDelete DataDictionaryCategoryUpdateColumn = "is_delete"
	// column name
	DataDictionaryCategoryUpdateColumnRemarks DataDictionaryCategoryUpdateColumn = "remarks"
	// column name
	DataDictionaryCategoryUpdateColumnUpdateAt DataDictionaryCategoryUpdateColumn = "update_at"
	// column name
	DataDictionaryCategoryUpdateColumnUpdateBy DataDictionaryCategoryUpdateColumn = "update_by"
)

var AllDataDictionaryCategoryUpdateColumn = []DataDictionaryCategoryUpdateColumn{
	DataDictionaryCategoryUpdateColumnCategoryCode,
	DataDictionaryCategoryUpdateColumnCategoryName,
	DataDictionaryCategoryUpdateColumnCreateAt,
	DataDictionaryCategoryUpdateColumnCreateBy,
	DataDictionaryCategoryUpdateColumnDeleteAt,
	DataDictionaryCategoryUpdateColumnDeleteBy,
	DataDictionaryCategoryUpdateColumnDictionaryCategoryID,
	DataDictionaryCategoryUpdateColumnID,
	DataDictionaryCategoryUpdateColumnIsDelete,
	DataDictionaryCategoryUpdateColumnRemarks,
	DataDictionaryCategoryUpdateColumnUpdateAt,
	DataDictionaryCategoryUpdateColumnUpdateBy,
}

func (e DataDictionaryCategoryUpdateColumn) IsValid() bool {
	switch e {
	case DataDictionaryCategoryUpdateColumnCategoryCode, DataDictionaryCategoryUpdateColumnCategoryName, DataDictionaryCategoryUpdateColumnCreateAt, DataDictionaryCategoryUpdateColumnCreateBy, DataDictionaryCategoryUpdateColumnDeleteAt, DataDictionaryCategoryUpdateColumnDeleteBy, DataDictionaryCategoryUpdateColumnDictionaryCategoryID, DataDictionaryCategoryUpdateColumnID, DataDictionaryCategoryUpdateColumnIsDelete, DataDictionaryCategoryUpdateColumnRemarks, DataDictionaryCategoryUpdateColumnUpdateAt, DataDictionaryCategoryUpdateColumnUpdateBy:
		return true
	}
	return false
}

func (e DataDictionaryCategoryUpdateColumn) String() string {
	return string(e)
}

func (e *DataDictionaryCategoryUpdateColumn) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DataDictionaryCategoryUpdateColumn(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid data_dictionary_category_update_column", str)
	}
	return nil
}

func (e DataDictionaryCategoryUpdateColumn) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// unique or primary key constraints on table "data_dictionary"
type DataDictionaryConstraint string

const (
	// unique or primary key constraint
	DataDictionaryConstraintDataDirtionaryPkey DataDictionaryConstraint = "data_dirtionary_pkey"
)

var AllDataDictionaryConstraint = []DataDictionaryConstraint{
	DataDictionaryConstraintDataDirtionaryPkey,
}

func (e DataDictionaryConstraint) IsValid() bool {
	switch e {
	case DataDictionaryConstraintDataDirtionaryPkey:
		return true
	}
	return false
}

func (e DataDictionaryConstraint) String() string {
	return string(e)
}

func (e *DataDictionaryConstraint) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DataDictionaryConstraint(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid data_dictionary_constraint", str)
	}
	return nil
}

func (e DataDictionaryConstraint) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// select columns of table "data_dictionary"
type DataDictionarySelectColumn string

const (
	// column name
	DataDictionarySelectColumnCreateAt DataDictionarySelectColumn = "create_at"
	// column name
	DataDictionarySelectColumnCreateBy DataDictionarySelectColumn = "create_by"
	// column name
	DataDictionarySelectColumnDeleteAt DataDictionarySelectColumn = "delete_at"
	// column name
	DataDictionarySelectColumnDeleteBy DataDictionarySelectColumn = "delete_by"
	// column name
	DataDictionarySelectColumnDictionaryCategoryID DataDictionarySelectColumn = "dictionary_category_id"
	// column name
	DataDictionarySelectColumnDictionaryID DataDictionarySelectColumn = "dictionary_id"
	// column name
	DataDictionarySelectColumnID DataDictionarySelectColumn = "id"
	// column name
	DataDictionarySelectColumnIsDelete DataDictionarySelectColumn = "is_delete"
	// column name
	DataDictionarySelectColumnName DataDictionarySelectColumn = "name"
	// column name
	DataDictionarySelectColumnRemarks DataDictionarySelectColumn = "remarks"
	// column name
	DataDictionarySelectColumnUpdateAt DataDictionarySelectColumn = "update_at"
	// column name
	DataDictionarySelectColumnUpdateBy DataDictionarySelectColumn = "update_by"
	// column name
	DataDictionarySelectColumnValue DataDictionarySelectColumn = "value"
)

var AllDataDictionarySelectColumn = []DataDictionarySelectColumn{
	DataDictionarySelectColumnCreateAt,
	DataDictionarySelectColumnCreateBy,
	DataDictionarySelectColumnDeleteAt,
	DataDictionarySelectColumnDeleteBy,
	DataDictionarySelectColumnDictionaryCategoryID,
	DataDictionarySelectColumnDictionaryID,
	DataDictionarySelectColumnID,
	DataDictionarySelectColumnIsDelete,
	DataDictionarySelectColumnName,
	DataDictionarySelectColumnRemarks,
	DataDictionarySelectColumnUpdateAt,
	DataDictionarySelectColumnUpdateBy,
	DataDictionarySelectColumnValue,
}

func (e DataDictionarySelectColumn) IsValid() bool {
	switch e {
	case DataDictionarySelectColumnCreateAt, DataDictionarySelectColumnCreateBy, DataDictionarySelectColumnDeleteAt, DataDictionarySelectColumnDeleteBy, DataDictionarySelectColumnDictionaryCategoryID, DataDictionarySelectColumnDictionaryID, DataDictionarySelectColumnID, DataDictionarySelectColumnIsDelete, DataDictionarySelectColumnName, DataDictionarySelectColumnRemarks, DataDictionarySelectColumnUpdateAt, DataDictionarySelectColumnUpdateBy, DataDictionarySelectColumnValue:
		return true
	}
	return false
}

func (e DataDictionarySelectColumn) String() string {
	return string(e)
}

func (e *DataDictionarySelectColumn) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DataDictionarySelectColumn(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid data_dictionary_select_column", str)
	}
	return nil
}

func (e DataDictionarySelectColumn) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// update columns of table "data_dictionary"
type DataDictionaryUpdateColumn string

const (
	// column name
	DataDictionaryUpdateColumnCreateAt DataDictionaryUpdateColumn = "create_at"
	// column name
	DataDictionaryUpdateColumnCreateBy DataDictionaryUpdateColumn = "create_by"
	// column name
	DataDictionaryUpdateColumnDeleteAt DataDictionaryUpdateColumn = "delete_at"
	// column name
	DataDictionaryUpdateColumnDeleteBy DataDictionaryUpdateColumn = "delete_by"
	// column name
	DataDictionaryUpdateColumnDictionaryCategoryID DataDictionaryUpdateColumn = "dictionary_category_id"
	// column name
	DataDictionaryUpdateColumnDictionaryID DataDictionaryUpdateColumn = "dictionary_id"
	// column name
	DataDictionaryUpdateColumnID DataDictionaryUpdateColumn = "id"
	// column name
	DataDictionaryUpdateColumnIsDelete DataDictionaryUpdateColumn = "is_delete"
	// column name
	DataDictionaryUpdateColumnName DataDictionaryUpdateColumn = "name"
	// column name
	DataDictionaryUpdateColumnRemarks DataDictionaryUpdateColumn = "remarks"
	// column name
	DataDictionaryUpdateColumnUpdateAt DataDictionaryUpdateColumn = "update_at"
	// column name
	DataDictionaryUpdateColumnUpdateBy DataDictionaryUpdateColumn = "update_by"
	// column name
	DataDictionaryUpdateColumnValue DataDictionaryUpdateColumn = "value"
)

var AllDataDictionaryUpdateColumn = []DataDictionaryUpdateColumn{
	DataDictionaryUpdateColumnCreateAt,
	DataDictionaryUpdateColumnCreateBy,
	DataDictionaryUpdateColumnDeleteAt,
	DataDictionaryUpdateColumnDeleteBy,
	DataDictionaryUpdateColumnDictionaryCategoryID,
	DataDictionaryUpdateColumnDictionaryID,
	DataDictionaryUpdateColumnID,
	DataDictionaryUpdateColumnIsDelete,
	DataDictionaryUpdateColumnName,
	DataDictionaryUpdateColumnRemarks,
	DataDictionaryUpdateColumnUpdateAt,
	DataDictionaryUpdateColumnUpdateBy,
	DataDictionaryUpdateColumnValue,
}

func (e DataDictionaryUpdateColumn) IsValid() bool {
	switch e {
	case DataDictionaryUpdateColumnCreateAt, DataDictionaryUpdateColumnCreateBy, DataDictionaryUpdateColumnDeleteAt, DataDictionaryUpdateColumnDeleteBy, DataDictionaryUpdateColumnDictionaryCategoryID, DataDictionaryUpdateColumnDictionaryID, DataDictionaryUpdateColumnID, DataDictionaryUpdateColumnIsDelete, DataDictionaryUpdateColumnName, DataDictionaryUpdateColumnRemarks, DataDictionaryUpdateColumnUpdateAt, DataDictionaryUpdateColumnUpdateBy, DataDictionaryUpdateColumnValue:
		return true
	}
	return false
}

func (e DataDictionaryUpdateColumn) String() string {
	return string(e)
}

func (e *DataDictionaryUpdateColumn) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DataDictionaryUpdateColumn(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid data_dictionary_update_column", str)
	}
	return nil
}

func (e DataDictionaryUpdateColumn) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
