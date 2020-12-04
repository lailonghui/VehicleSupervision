package util

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"reflect"
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
	return &QueryTranslator{tx: tx, model: model}
}

// 排序
func (t *QueryTranslator) OrderBy(orderBy interface{}) (re *QueryTranslator) {
	re = t
	if orderBy == nil {
		return re
	}
	re.orderBy = orderBy

	orderByValue := getValue(orderBy)
	orderByKind := orderByValue.Kind()
	switch orderByKind {
	case reflect.Slice:
		sliceLength := orderByValue.Len()
		for i := 0; i < sliceLength; i++ {
			orderByItem := getValue(orderByValue.Index(i))
			re.tx = re.tx.Order(buildOrderBy(orderByItem))
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
	if distinctOn == nil {
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
	if where == nil {
		return re
	}
	re.where = where

	return re
}

// 执行翻译graphql到sql到操作
func (t *QueryTranslator) DoTranslate() *gorm.DB {

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
func buildOrderBy(value reflect.Value) clause.OrderByColumn {
	// 获取orderByItem第一个非nil的属性和这个属性的tag
	// 即为 order的方向和属性
	valueKind := value.Kind()
	switch valueKind {
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			fieldValue := value.Field(i)
			fmt.Println(fieldValue)
		}
	default:
		panic(errors.New("unSupport type"))
	}

	return clause.OrderByColumn{}
}
