// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// 企业扣分项
type EnterpriseScoreSet struct {
	// ID
	ID int64 `json:"id"`
	// 扣分项ID
	ScoreSetID string `json:"score_set_id"`
	// 扣分内容
	Content string `json:"content"`
	// 类型
	Type int `json:"type"`
	// 分值
	Score float64 `json:"score"`
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 创建人
	CreatedBy *string `json:"created_by"`
	// 修改时间
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人
	UpdatedBy *string `json:"updated_by"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人
	DeletedBy *string `json:"deleted_by"`
	// 是否删除
	IsDeleted bool `json:"is_deleted"`
}
