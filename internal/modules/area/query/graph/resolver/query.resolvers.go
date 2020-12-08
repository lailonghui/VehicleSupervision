package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/modules/area/query/graph/generated"
	"VehicleSupervision/internal/modules/area/query/graph/model"
	"context"
	"fmt"
)

func (r *queryResolver) City(ctx context.Context, distinctOn []model.CitySelectColumn, limit *int, offset *int, orderBy []*model.CityOrderBy, where *model.CityBoolExp) ([]*model.City, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CityAggregate(ctx context.Context, distinctOn []model.CitySelectColumn, limit *int, offset *int, orderBy []*model.CityOrderBy, where *model.CityBoolExp) (*model.CityAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CityByPk(ctx context.Context, id int64) (*model.City, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) District(ctx context.Context, distinctOn []model.DistrictSelectColumn, limit *int, offset *int, orderBy []*model.DistrictOrderBy, where *model.DistrictBoolExp) ([]*model.District, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DistrictAggregate(ctx context.Context, distinctOn []model.DistrictSelectColumn, limit *int, offset *int, orderBy []*model.DistrictOrderBy, where *model.DistrictBoolExp) (*model.DistrictAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DistrictByPk(ctx context.Context, id int64) (*model.District, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
