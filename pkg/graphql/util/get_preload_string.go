package util

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"strings"
)

// 获取preload
// preload是指graphql请求的值
func GetPreloads(ctx context.Context) []string {
	return GetNestedPreloads(
		graphql.GetOperationContext(ctx),
		graphql.CollectFieldsCtx(ctx, nil),
		"",
	)
}

// 获取嵌套preload
func GetNestedPreloads(ctx *graphql.OperationContext, fields []graphql.CollectedField, prefix string) (preloads []string) {
	for _, column := range fields {
		prefixColumn := GetPreloadString(prefix, column.Name)
		preloads = append(preloads, prefixColumn)
		preloads = append(preloads, GetNestedPreloads(ctx, graphql.CollectFields(ctx, column.Selections, nil), prefixColumn)...)
	}
	return
}

// 获取preload
func GetPreloadString(prefix, name string) string {
	if len(prefix) > 0 {
		return prefix + "." + name
	}
	return name
}

// 获取preload数组，数组内容必须以prefix为前缀，并且不等于prefix
func GetPreloadsMustPrefix(ctx context.Context, prefix string) (rs []string) {
	preloads := GetPreloads(ctx)
	if preloads == nil || len(preloads) == 0 {
		return
	}
	for _, preload := range preloads {
		if preload != prefix && strings.HasPrefix(preload, prefix) {
			rs = append(rs, preload)
		}
	}
	return
}

// 获取preload数组，数组内容必须以prefix为前缀，并且不等于prefix
// 并且移除prefix
func GetPreloadsMustPrefixAndRemovePrefix(ctx context.Context, prefix string) (rs []string) {
	preloads := GetPreloads(ctx)
	if preloads == nil || len(preloads) == 0 {
		return
	}
	for _, preload := range preloads {
		if preload != prefix && strings.HasPrefix(preload, prefix) {
			rs = append(rs, preload[len(prefix):])
		}
	}
	return
}

// 获取最顶层的preload
func GetTopPreloads(ctx context.Context) (rs []string) {

	preloads := GetPreloads(ctx)
	if preloads == nil || len(preloads) == 0 {
		return
	}
	relationTables := make(map[string]bool, 0)
	for _, preload := range preloads {
		dotIndex := strings.Index(preload, ".")
		if dotIndex != -1 {
			relationTables[preload[:dotIndex]] = true
			continue
		}
	}
	for _, preload := range preloads {
		dotIndex := strings.Index(preload, ".")
		if dotIndex != -1 {
			continue
		}
		ok, _ := relationTables[preload]
		if !ok {
			rs = append(rs, preload)
		}
	}
	return rs
}
