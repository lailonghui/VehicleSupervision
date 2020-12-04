package model

import (
	"fmt"
	"io"
	"strconv"
)

// column ordering options
type OrderBy string

const (
	// in the ascending order, nulls last
	OrderByAsc OrderBy = "asc"
	// in the ascending order, nulls first
	OrderByAscNullsFirst OrderBy = "asc_nulls_first"
	// in the ascending order, nulls last
	OrderByAscNullsLast OrderBy = "asc_nulls_last"
	// in the descending order, nulls first
	OrderByDesc OrderBy = "desc"
	// in the descending order, nulls first
	OrderByDescNullsFirst OrderBy = "desc_nulls_first"
	// in the descending order, nulls last
	OrderByDescNullsLast OrderBy = "desc_nulls_last"
)

var AllOrderBy = []OrderBy{
	OrderByAsc,
	OrderByAscNullsFirst,
	OrderByAscNullsLast,
	OrderByDesc,
	OrderByDescNullsFirst,
	OrderByDescNullsLast,
}

func (e OrderBy) IsValid() bool {
	switch e {
	case OrderByAsc, OrderByAscNullsFirst, OrderByAscNullsLast, OrderByDesc, OrderByDescNullsFirst, OrderByDescNullsLast:
		return true
	}
	return false
}

func (e OrderBy) String() string {
	return string(e)
}

func (e *OrderBy) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderBy(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid order_by", str)
	}
	return nil
}

func (e OrderBy) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

