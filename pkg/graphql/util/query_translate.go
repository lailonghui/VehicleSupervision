package util

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/pkg/graphql/model"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"reflect"
	"strings"
)

// graphql查询翻译器（翻译成sql）
type QueryTranslator struct {
	// 数据库model
	model interface{}
	// distinct
	distinctOn interface{}
	// limit
	limit *int
	// offset
	offset *int
	// order by
	orderBy interface{}
	// where
	where interface{}
	// tx
	tx *gorm.DB
}

func NewQueryTranslator(tx *gorm.DB, model interface{}) *QueryTranslator {
	tx = tx.Model(model)
	return &QueryTranslator{tx: tx, model: model}
}

// 排序
func (t *QueryTranslator) OrderBy(orderBy interface{}) (re *QueryTranslator) {
	re = t
	if isNil(orderBy) {
		return re
	}
	re.orderBy = orderBy

	orderByValue := getValue(orderBy)
	orderByKind := orderByValue.Kind()

	switch orderByKind {
	case reflect.Slice:
		sliceLength := orderByValue.Len()
		for i := 0; i < sliceLength; i++ {
			orderByItem := orderByValue.Index(i)
			if orderByItem.IsNil() {
				continue
			}
			if orderByItem.Kind() == reflect.Ptr {
				orderByItem = orderByItem.Elem()
			}
			orderByColumn, ok := buildOrderBy(orderByItem)
			if ok {
				re.tx = re.tx.Order(orderByColumn)
			}

		}
	default:
		panic(errors.New("unSupport type"))
	}
	return re
}

// 设置offset
func (t *QueryTranslator) Offset(offset *int) (re *QueryTranslator) {
	re = t
	if offset == nil {
		return re
	}
	re.offset = offset

	re.tx = re.tx.Offset(*offset)
	return re
}

// 设置limit
func (t *QueryTranslator) Limit(limit *int) (re *QueryTranslator) {
	re = t
	if limit == nil {
		return re
	}
	re.limit = limit
	re.tx = re.tx.Limit(*limit)
	return re
}

// 设置distinctOn
func (t *QueryTranslator) DistinctOn(distinctOn interface{}) (re *QueryTranslator) {
	re = t
	if isNil(distinctOn) {
		return re
	}
	re.distinctOn = distinctOn

	distinctOnValue := getValue(distinctOn)
	distinctOnKind := distinctOnValue.Kind()
	switch distinctOnKind {
	case reflect.Slice:
		sliceLength := distinctOnValue.Len()
		sArray := make([]string, sliceLength)
		for i := 0; i < sliceLength; i++ {
			sArray[i] = distinctOnValue.Index(i).String()
		}
		re.tx = re.tx.Distinct(sArray)
	default:
		panic(errors.New("unSupport type"))
	}

	return re
}

// 设置where
func (t *QueryTranslator) Where(where interface{}) (re *QueryTranslator) {
	re = t
	if isNil(where) {
		return re
	}
	re.where = where

	re.tx = buildWhere(re.tx, where)

	return re
}

// 结束，返回tx
func (t *QueryTranslator) Finish() *gorm.DB {

	return t.tx
}

// 聚合结果，返回tx
func (t *QueryTranslator) Aggregate(rs interface{}, ctx context.Context) *gorm.DB {
	// 获取聚合查询项
	queryStrings := GetPreloadsMustPrefix(ctx, "aggregate.")
	if queryStrings == nil || len(queryStrings) == 0 {
		return t.tx
	}

	// select语句buffer
	selects := getAggregateSelect(queryStrings)
	selectStr := strings.Join(selects, ",")
	var results []map[string]interface{}
	t.tx.Select(selectStr).Find(&results)
	fmt.Println(results)

	return t.tx
}

// 获取reflect.value
func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}

// 反射构建order
func buildOrderBy(value reflect.Value) (clause.OrderByColumn, bool) {
	valueKind := value.Kind()
	valueType := value.Type()
	switch valueKind {
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			fieldValue := value.Field(i)
			isFieldPtr := fieldValue.Kind() == reflect.Ptr
			if isFieldPtr && fieldValue.IsNil() {
				continue
			}
			if isFieldPtr {
				fieldValue = fieldValue.Elem()
			}
			columnName := valueType.Field(i).Tag.Get("json")
			switch fieldValue.String() {
			case string(model.OrderByAsc):
				return clause.OrderByColumn{Column: clause.Column{
					Name: columnName,
					Raw:  true,
				}, Desc: false}, true
			case string(model.OrderByDesc):
				return clause.OrderByColumn{Column: clause.Column{
					Name: columnName,
					Raw:  true,
				}, Desc: true}, true
			default:

			}

		}
	default:
		panic(errors.New("unSupport type"))
	}

	return clause.OrderByColumn{}, false
}

// 构建where
func buildWhere(tx *gorm.DB, where interface{}) *gorm.DB {
	value := getValue(where)
	valueKind := value.Kind()
	valueType := value.Type()
	switch valueKind {
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			fieldValue := value.Field(i)
			if fieldValue.IsNil() {
				continue
			}
			fieldKind := fieldValue.Kind()
			switch fieldKind {
			case reflect.Ptr:
				// 指针的情况，包括not子句和普通条件查询
				rFieldValue := fieldValue.Elem()
				rFieldType := rFieldValue.Type()
				columnName := valueType.Field(i).Tag.Get("json")
				if rFieldType == valueType {
					// not子句查询
					if columnName == "_not" {
						tx = tx.Not(buildWhere(db.DB, rFieldValue.Interface()))
					}
					continue
				}
				rValue := rFieldValue.Interface()
				// 条件查询
				switch exp := rValue.(type) {
				case model.BigintComparisonExp:
					tx = bigintCompare(tx, exp, columnName)
				case model.BooleanComparisonExp:
					tx = booleanCompare(tx, exp, columnName)
				case model.IntComparisonExp:
					tx = intCompare(tx, exp, columnName)
				case model.JsonbComparisonExp:
					tx = jsonbCompare(tx, exp, columnName)
				case model.StringComparisonExp:
					tx = stringCompare(tx, exp, columnName)
				case model.TimestamptzComparisonExp:
					tx = timestamptzCompare(tx, exp, columnName)
				default:
					panic(errors.New("unSupport type"))
				}
			case reflect.Slice:
				// 切片情况，包括and子句和or子句
				columnName := valueType.Field(i).Tag.Get("json")
				sliceLength := fieldValue.Len()
				for i := 0; i < sliceLength; i++ {
					sliceItem := fieldValue.Index(i).Interface()
					// and子句
					if columnName == "_and" {
						tx = tx.Where(buildWhere(db.DB, sliceItem))
					}
					// or子句
					if columnName == "_or" {
						tx = tx.Or(buildWhere(db.DB, sliceItem))
					}
				}

			default:
				panic(errors.New("unSupport type"))

			}

		}
	}

	return tx
}

// bigint 比较判断
func bigintCompare(tx *gorm.DB, exp model.BigintComparisonExp, columnName string) *gorm.DB {
	if exp.Eq != nil {
		tx = tx.Where(columnName+" = ? ", exp.Eq)
	}
	if exp.Gt != nil {
		tx = tx.Where(columnName+" > ? ", exp.Gt)
	}
	if exp.Gte != nil {
		tx = tx.Where(columnName+" >= ?", exp.Gte)
	}
	if exp.In != nil {
		tx = tx.Where(columnName+" in ? ", exp.In)
	}
	if exp.IsNull != nil {
		tx = tx.Where(columnName+" is null", exp.IsNull)
	}
	if exp.Lt != nil {
		tx = tx.Where(columnName+" < ? ", exp.Lt)
	}
	if exp.Lte != nil {
		tx = tx.Where(columnName+" <= ? ", exp.Lte)
	}
	if exp.Neq != nil {
		tx = tx.Not(columnName+" = ? ", exp.Neq)
	}
	if exp.Nin != nil {
		tx = tx.Not(columnName+" in ? ", exp.Nin)
	}

	return tx
}

// boolean 比较判断
func booleanCompare(tx *gorm.DB, exp model.BooleanComparisonExp, columnName string) *gorm.DB {
	if exp.Eq != nil {
		tx = tx.Where(columnName+" = ? ", exp.Eq)
	}
	if exp.Gt != nil {
		tx = tx.Where(columnName+" > ? ", exp.Gt)
	}
	if exp.Gte != nil {
		tx = tx.Where(columnName+" >= ?", exp.Gte)
	}
	if exp.In != nil {
		tx = tx.Where(columnName+" in ? ", exp.In)
	}
	if exp.IsNull != nil {
		tx = tx.Where(columnName+" is null", exp.IsNull)
	}
	if exp.Lt != nil {
		tx = tx.Where(columnName+" < ? ", exp.Lt)
	}
	if exp.Lte != nil {
		tx = tx.Where(columnName+" <= ? ", exp.Lte)
	}
	if exp.Neq != nil {
		tx = tx.Not(columnName+" = ? ", exp.Neq)
	}
	if exp.Nin != nil {
		tx = tx.Not(columnName+" in ? ", exp.Nin)
	}
	return tx
}

// int 比较判断
func intCompare(tx *gorm.DB, exp model.IntComparisonExp, columnName string) *gorm.DB {
	if exp.Eq != nil {
		tx = tx.Where(columnName+" = ? ", exp.Eq)
	}
	if exp.Gt != nil {
		tx = tx.Where(columnName+" > ? ", exp.Gt)
	}
	if exp.Gte != nil {
		tx = tx.Where(columnName+" >= ?", exp.Gte)
	}
	if exp.In != nil {
		tx = tx.Where(columnName+" in ? ", exp.In)
	}
	if exp.IsNull != nil {
		tx = tx.Where(columnName+" is null", exp.IsNull)
	}
	if exp.Lt != nil {
		tx = tx.Where(columnName+" < ? ", exp.Lt)
	}
	if exp.Lte != nil {
		tx = tx.Where(columnName+" <= ? ", exp.Lte)
	}
	if exp.Neq != nil {
		tx = tx.Not(columnName+" = ? ", exp.Neq)
	}
	if exp.Nin != nil {
		tx = tx.Not(columnName+" in ? ", exp.Nin)
	}
	return tx
}

// jsonb 比较判断
func jsonbCompare(tx *gorm.DB, exp model.JsonbComparisonExp, columnName string) *gorm.DB {
	if exp.Eq != nil {
		tx = tx.Where(columnName+" = ? ", exp.Eq)
	}
	if exp.Gt != nil {
		tx = tx.Where(columnName+" > ? ", exp.Gt)
	}
	if exp.Gte != nil {
		tx = tx.Where(columnName+" >= ?", exp.Gte)
	}
	if exp.In != nil {
		tx = tx.Where(columnName+" in ? ", exp.In)
	}
	if exp.IsNull != nil {
		tx = tx.Where(columnName+" is null", exp.IsNull)
	}
	if exp.Lt != nil {
		tx = tx.Where(columnName+" < ? ", exp.Lt)
	}
	if exp.Lte != nil {
		tx = tx.Where(columnName+" <= ? ", exp.Lte)
	}
	if exp.Neq != nil {
		tx = tx.Not(columnName+" = ? ", exp.Neq)
	}
	if exp.Nin != nil {
		tx = tx.Not(columnName+" in ? ", exp.Nin)
	}

	return tx
}

// string 比较判断
func stringCompare(tx *gorm.DB, exp model.StringComparisonExp, columnName string) *gorm.DB {
	if exp.Eq != nil {
		tx = tx.Where(columnName+" = ? ", exp.Eq)
	}
	if exp.Gt != nil {
		tx = tx.Where(columnName+" > ? ", exp.Gt)
	}
	if exp.Gte != nil {
		tx = tx.Where(columnName+" >= ?", exp.Gte)
	}
	if exp.In != nil {
		tx = tx.Where(columnName+" in ? ", exp.In)
	}
	if exp.IsNull != nil {
		tx = tx.Where(columnName+" is null", exp.IsNull)
	}
	if exp.Lt != nil {
		tx = tx.Where(columnName+" < ? ", exp.Lt)
	}
	if exp.Lte != nil {
		tx = tx.Where(columnName+" <= ? ", exp.Lte)
	}
	if exp.Neq != nil {
		tx = tx.Not(columnName+" = ? ", exp.Neq)
	}
	if exp.Nin != nil {
		tx = tx.Not(columnName+" in ? ", exp.Nin)
	}
	if exp.Like != nil {
		tx = tx.Where(columnName+" like ?% ", exp.Like)
	}
	if exp.Ilike != nil {
		tx = tx.Where(columnName+" like ?% ", exp.Ilike)
	}
	if exp.Similar != nil {
		tx = tx.Where(columnName+" like ?% ", exp.Similar)
	}
	if exp.Nlike != nil {
		tx = tx.Not(columnName+" like ?% ", exp.Nlike)
	}
	if exp.Nilike != nil {
		tx = tx.Not(columnName+" like ?% ", exp.Nilike)
	}
	if exp.Nsimilar != nil {
		tx = tx.Not(columnName+" like ?% ", exp.Nsimilar)
	}

	return tx
}

// timestamptz 比较判断
func timestamptzCompare(tx *gorm.DB, exp model.TimestamptzComparisonExp, columnName string) *gorm.DB {
	if exp.Eq != nil {
		tx = tx.Where(columnName+" = ? ", exp.Eq)
	}
	if exp.Gt != nil {
		tx = tx.Where(columnName+" > ? ", exp.Gt)
	}
	if exp.Gte != nil {
		tx = tx.Where(columnName+" >= ?", exp.Gte)
	}
	if exp.In != nil {
		tx = tx.Where(columnName+" in ? ", exp.In)
	}
	if exp.IsNull != nil {
		tx = tx.Where(columnName+" is null", exp.IsNull)
	}
	if exp.Lt != nil {
		tx = tx.Where(columnName+" < ? ", exp.Lt)
	}
	if exp.Lte != nil {
		tx = tx.Where(columnName+" <= ? ", exp.Lte)
	}
	if exp.Neq != nil {
		tx = tx.Not(columnName+" = ? ", exp.Neq)
	}
	if exp.Nin != nil {
		tx = tx.Not(columnName+" in ? ", exp.Nin)
	}

	return tx
}

// 判断interface{} 是否是nil
func isNil(i interface{}) bool {
	defer func() {
		recover()
	}()
	vi := reflect.ValueOf(i)
	return vi.IsNil()
}

// 获取聚合查询select内容数组
func getAggregateSelect(queryStrings []string) (selects []string) {
	for _, queryString := range queryStrings {
		queryInfos := strings.Split(queryString, ".")
		if len(queryInfos) < 2 {
			continue
		}
		aggregateType := queryInfos[1]
		switch aggregateType {
		case "avg":
			if len(queryInfos) >= 3 {
				column := queryInfos[2]
				selects = append(selects, "avg("+column+") as _avg_"+column)
			}
		case "count":
			selects = append(selects, "count(*) as _count")
		case "max":
			if len(queryInfos) >= 3 {
				column := queryInfos[2]
				selects = append(selects, "max("+column+") as _max_"+column)
			}
		case "min":
			if len(queryInfos) >= 3 {
				column := queryInfos[2]
				selects = append(selects, "min("+column+") as _min_"+column)
			}
		case "stddev":
			if len(queryInfos) >= 3 {
				column := queryInfos[2]
				selects = append(selects, "stddev("+column+") as _stddev_"+column)
			}
		case "stddev_pop":
			if len(queryInfos) >= 3 {
				column := queryInfos[2]
				selects = append(selects, "stddev_pop("+column+") as _stddev_pop_"+column)
			}
		case "stddev_samp":
			if len(queryInfos) >= 3 {
				column := queryInfos[2]
				selects = append(selects, "stddev_samp("+column+") as _stddev_samp_"+column)
			}
		case "sum":
			if len(queryInfos) >= 3 {
				column := queryInfos[2]
				selects = append(selects, "sum("+column+") as _sum_"+column)
			}
		case "var_pop":
			if len(queryInfos) >= 3 {
				column := queryInfos[2]
				selects = append(selects, "var_pop("+column+") as _var_pop_"+column)
			}
		case "var_samp":
			if len(queryInfos) >= 3 {
				column := queryInfos[2]
				selects = append(selects, "var_samp("+column+") as _var_samp_"+column)
			}
		case "variance":
			if len(queryInfos) >= 3 {
				column := queryInfos[2]
				selects = append(selects, "variance("+column+") as _variance_"+column)
			}
		default:
			panic("unsupport query")
		}
	}

	return
}
