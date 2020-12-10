package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/modules/vehicle/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) DeleteMuckTruckInfo(ctx context.Context, where model.MuckTruckInfoBoolExp) (*model.MuckTruckInfoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteMuckTruckInfoByPk(ctx context.Context, muckTruckID int64) (*model.MuckTruckInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertMuckTruckInfo(ctx context.Context, objects []*model.MuckTruckInfoInsertInput, onConflict *model.MuckTruckInfoOnConflict) (*model.MuckTruckInfoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertMuckTruckInfoOne(ctx context.Context, object model.MuckTruckInfoInsertInput, onConflict *model.MuckTruckInfoOnConflict) (*model.MuckTruckInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateMuckTruckInfo(ctx context.Context, inc *model.MuckTruckInfoIncInput, set *model.MuckTruckInfoSetInput, where model.MuckTruckInfoBoolExp) (*model.MuckTruckInfoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateMuckTruckInfoByPk(ctx context.Context, inc *model.MuckTruckInfoIncInput, set *model.MuckTruckInfoSetInput, pkColumns model.MuckTruckInfoPkColumnsInput) (*model.MuckTruckInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) MuckTruckInfo(ctx context.Context, distinctOn []model.MuckTruckInfoSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckInfoOrderBy, where *model.MuckTruckInfoBoolExp) ([]*model.MuckTruckInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) MuckTruckInfoAggregate(ctx context.Context, distinctOn []model.MuckTruckInfoSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckInfoOrderBy, where *model.MuckTruckInfoBoolExp) (*model.MuckTruckInfoAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) MuckTruckInfoByPk(ctx context.Context, muckTruckID int64) (*model.MuckTruckInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) MuckTruckInfo(ctx context.Context, distinctOn []model.MuckTruckInfoSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckInfoOrderBy, where *model.MuckTruckInfoBoolExp) (<-chan []*model.MuckTruckInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) MuckTruckInfoAggregate(ctx context.Context, distinctOn []model.MuckTruckInfoSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckInfoOrderBy, where *model.MuckTruckInfoBoolExp) (<-chan *model.MuckTruckInfoAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) MuckTruckInfoByPk(ctx context.Context, muckTruckID int64) (<-chan *model.MuckTruckInfo, error) {
	panic(fmt.Errorf("not implemented"))
}
