package util

import "gorm.io/gorm"

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

func (t *QueryTranslator) OrderBy(orderBy interface{}) (r *QueryTranslator) {
	r = t
	r.orderBy = orderBy
	return r
}

// 设置tx
func (t *QueryTranslator) Tx(tx *gorm.DB) (r *QueryTranslator) {
	r = t
	r.tx = tx
	return r
}

// 设置offset
func (t *QueryTranslator) Offset(offset *int) (r *QueryTranslator) {
	r = t
	r.offset = offset
	return r
}

// 设置limit
func (t *QueryTranslator) Limit(limit *int) (r *QueryTranslator) {
	r = t
	r.limit = limit
	return r
}

// 设置distinctOn
func (t *QueryTranslator) DistinctOn(distinctOn interface{}) (r *QueryTranslator) {
	r = t
	r.distinctOn = distinctOn
	return r
}

// 设置where
func (t *QueryTranslator) Where(where interface{}) (r *QueryTranslator) {

	r = t
	r.where = where
	return r
}

// 设置model
func (t *QueryTranslator) Model(m interface{}) (r *QueryTranslator) {
	r = t
	r.model = m
	return
}

// 执行翻译graphql到sql到操作
func (t QueryTranslator) DoTranslate() *gorm.DB {

	return t.tx
}
