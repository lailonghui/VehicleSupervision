package model

import "time"

// expression to compare columns of type Boolean. All fields are combined with logical 'AND'.
type BooleanComparisonExp struct {
	Eq     *bool  `json:"_eq"`
	Gt     *bool  `json:"_gt"`
	Gte    *bool  `json:"_gte"`
	In     []bool `json:"_in"`
	IsNull *bool  `json:"_is_null"`
	Lt     *bool  `json:"_lt"`
	Lte    *bool  `json:"_lte"`
	Neq    *bool  `json:"_neq"`
	Nin    []bool `json:"_nin"`
}

// expression to compare columns of type Int. All fields are combined with logical 'AND'.
type IntComparisonExp struct {
	Eq     *int  `json:"_eq"`
	Gt     *int  `json:"_gt"`
	Gte    *int  `json:"_gte"`
	In     []int `json:"_in"`
	IsNull *bool `json:"_is_null"`
	Lt     *int  `json:"_lt"`
	Lte    *int  `json:"_lte"`
	Neq    *int  `json:"_neq"`
	Nin    []int `json:"_nin"`
}

// expression to compare columns of type String. All fields are combined with logical 'AND'.
type StringComparisonExp struct {
	Eq       *string  `json:"_eq"`
	Gt       *string  `json:"_gt"`
	Gte      *string  `json:"_gte"`
	Ilike    *string  `json:"_ilike"`
	In       []string `json:"_in"`
	IsNull   *bool    `json:"_is_null"`
	Like     *string  `json:"_like"`
	Lt       *string  `json:"_lt"`
	Lte      *string  `json:"_lte"`
	Neq      *string  `json:"_neq"`
	Nilike   *string  `json:"_nilike"`
	Nin      []string `json:"_nin"`
	Nlike    *string  `json:"_nlike"`
	Nsimilar *string  `json:"_nsimilar"`
	Similar  *string  `json:"_similar"`
}

// expression to compare columns of type _jsonb. All fields are combined with logical 'AND'.
type JsonbComparisonExp struct {
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

// expression to compare columns of type bigint. All fields are combined with logical 'AND'.
type BigintComparisonExp struct {
	Eq     *int64  `json:"_eq"`
	Gt     *int64  `json:"_gt"`
	Gte    *int64  `json:"_gte"`
	In     []int64 `json:"_in"`
	IsNull *bool   `json:"_is_null"`
	Lt     *int64  `json:"_lt"`
	Lte    *int64  `json:"_lte"`
	Neq    *int64  `json:"_neq"`
	Nin    []int64 `json:"_nin"`
}

// expression to compare columns of type timestamptz. All fields are combined with logical 'AND'.
type TimestamptzComparisonExp struct {
	Eq     *time.Time   `json:"_eq"`
	Gt     *time.Time   `json:"_gt"`
	Gte    *time.Time   `json:"_gte"`
	In     []*time.Time `json:"_in"`
	IsNull *bool        `json:"_is_null"`
	Lt     *time.Time   `json:"_lt"`
	Lte    *time.Time   `json:"_lte"`
	Neq    *time.Time   `json:"_neq"`
	Nin    []*time.Time `json:"_nin"`
}

// expression to compare columns of type date. All fields are combined with logical 'AND'.
type DateComparisonExp struct {
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
