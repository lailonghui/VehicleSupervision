package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"system-manage/internal/graphql/system/graph/model"
)

func (r *queryResolver) DictList(ctx context.Context, query model.DictListQuery) ([]*model.Dict, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ProvinceList(ctx context.Context, query *model.ProvinceListQuery) ([]*model.Province, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CityList(ctx context.Context, query *model.CityListQuery) ([]*model.City, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DistrictList(ctx context.Context, query *model.DistrictListQuery) ([]*model.District, error) {
	panic(fmt.Errorf("not implemented"))
}
