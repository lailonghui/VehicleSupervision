package gqlcache

// gql缓存策略
type Strategy int

const (
	// 不使用缓存
	NONE Strategy = iota + 1
	// 实体缓存
	ENTITY
	// 列表缓存(包含实体缓存)
	LIST
)
