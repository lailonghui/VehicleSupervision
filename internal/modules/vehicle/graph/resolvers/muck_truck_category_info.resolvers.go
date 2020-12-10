package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/modules/vehicle/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) DeleteMuckTruckCategoryInfo(ctx context.Context, where model.MuckTruckCategoryInfoBoolExp) (*model.MuckTruckCategoryInfoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteMuckTruckCategoryInfoByPk(ctx context.Context, id int64) (*model.MuckTruckCategoryInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertMuckTruckCategoryInfo(ctx context.Context, objects []*model.MuckTruckCategoryInfoInsertInput, onConflict *model.MuckTruckCategoryInfoOnConflict) (*model.MuckTruckCategoryInfoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertMuckTruckCategoryInfoOne(ctx context.Context, object model.MuckTruckCategoryInfoInsertInput, onConflict *model.MuckTruckCategoryInfoOnConflict) (*model.MuckTruckCategoryInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateMuckTruckCategoryInfo(ctx context.Context, inc *model.MuckTruckCategoryInfoIncInput, set *model.MuckTruckCategoryInfoSetInput, where model.MuckTruckCategoryInfoBoolExp) (*model.MuckTruckCategoryInfoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateMuckTruckCategoryInfoByPk(ctx context.Context, inc *model.MuckTruckCategoryInfoIncInput, set *model.MuckTruckCategoryInfoSetInput, pkColumns model.MuckTruckCategoryInfoPkColumnsInput) (*model.MuckTruckCategoryInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) MuckTruckCategoryInfo(ctx context.Context, distinctOn []model.MuckTruckCategoryInfoSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckCategoryInfoOrderBy, where *model.MuckTruckCategoryInfoBoolExp) ([]*model.MuckTruckCategoryInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) MuckTruckCategoryInfoAggregate(ctx context.Context, distinctOn []model.MuckTruckCategoryInfoSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckCategoryInfoOrderBy, where *model.MuckTruckCategoryInfoBoolExp) (*model.MuckTruckCategoryInfoAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) MuckTruckCategoryInfoByPk(ctx context.Context, id int64) (*model.MuckTruckCategoryInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) MuckTruckCategoryInfo(ctx context.Context, distinctOn []model.MuckTruckCategoryInfoSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckCategoryInfoOrderBy, where *model.MuckTruckCategoryInfoBoolExp) (<-chan []*model.MuckTruckCategoryInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) MuckTruckCategoryInfoAggregate(ctx context.Context, distinctOn []model.MuckTruckCategoryInfoSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckCategoryInfoOrderBy, where *model.MuckTruckCategoryInfoBoolExp) (<-chan *model.MuckTruckCategoryInfoAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) MuckTruckCategoryInfoByPk(ctx context.Context, id int64) (<-chan *model.MuckTruckCategoryInfo, error) {
	panic(fmt.Errorf("not implemented"))
}
