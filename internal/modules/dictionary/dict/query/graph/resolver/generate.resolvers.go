package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/modules/dictionary/dict/query/graph/generated"
	"VehicleSupervision/internal/modules/dictionary/dict/query/graph/model"
	"context"
	"fmt"
)

func (r *queryResolver) DataDictionaryCategory(ctx context.Context, distinctOn []model.DataDictionaryCategorySelectColumn, limit *int, offset *int, orderBy []*model.DataDictionaryCategoryOrderBy, where *model.DataDictionaryCategoryBoolExp) ([]*model.DataDictionaryCategory, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DataDictionaryCategoryAggregate(ctx context.Context, distinctOn []model.DataDictionaryCategorySelectColumn, limit *int, offset *int, orderBy []*model.DataDictionaryCategoryOrderBy, where *model.DataDictionaryCategoryBoolExp) (*model.DataDictionaryCategoryAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DataDictionaryCategoryByPk(ctx context.Context, id int64) (*model.DataDictionaryCategory, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
