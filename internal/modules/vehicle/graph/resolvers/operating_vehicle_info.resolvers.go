package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/modules/vehicle/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) DeleteOperatingVehicleInfo(ctx context.Context, where model.OperatingVehicleInfoBoolExp) (*model.OperatingVehicleInfoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteOperatingVehicleInfoByPk(ctx context.Context, operatingVehicleID int64) (*model.OperatingVehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertOperatingVehicleInfo(ctx context.Context, objects []*model.OperatingVehicleInfoInsertInput, onConflict *model.OperatingVehicleInfoOnConflict) (*model.OperatingVehicleInfoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertOperatingVehicleInfoOne(ctx context.Context, object model.OperatingVehicleInfoInsertInput, onConflict *model.OperatingVehicleInfoOnConflict) (*model.OperatingVehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateOperatingVehicleInfo(ctx context.Context, inc *model.OperatingVehicleInfoIncInput, set *model.OperatingVehicleInfoSetInput, where model.OperatingVehicleInfoBoolExp) (*model.OperatingVehicleInfoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateOperatingVehicleInfoByPk(ctx context.Context, inc *model.OperatingVehicleInfoIncInput, set *model.OperatingVehicleInfoSetInput, pkColumns model.OperatingVehicleInfoPkColumnsInput) (*model.OperatingVehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) OperatingVehicleInfo(ctx context.Context, distinctOn []model.OperatingVehicleInfoSelectColumn, limit *int, offset *int, orderBy []*model.OperatingVehicleInfoOrderBy, where *model.OperatingVehicleInfoBoolExp) ([]*model.OperatingVehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) OperatingVehicleInfoAggregate(ctx context.Context, distinctOn []model.OperatingVehicleInfoSelectColumn, limit *int, offset *int, orderBy []*model.OperatingVehicleInfoOrderBy, where *model.OperatingVehicleInfoBoolExp) (*model.OperatingVehicleInfoAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) OperatingVehicleInfoByPk(ctx context.Context, operatingVehicleID int64) (*model.OperatingVehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) OperatingVehicleInfo(ctx context.Context, distinctOn []model.OperatingVehicleInfoSelectColumn, limit *int, offset *int, orderBy []*model.OperatingVehicleInfoOrderBy, where *model.OperatingVehicleInfoBoolExp) (<-chan []*model.OperatingVehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) OperatingVehicleInfoAggregate(ctx context.Context, distinctOn []model.OperatingVehicleInfoSelectColumn, limit *int, offset *int, orderBy []*model.OperatingVehicleInfoOrderBy, where *model.OperatingVehicleInfoBoolExp) (<-chan *model.OperatingVehicleInfoAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) OperatingVehicleInfoByPk(ctx context.Context, operatingVehicleID int64) (<-chan *model.OperatingVehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}
