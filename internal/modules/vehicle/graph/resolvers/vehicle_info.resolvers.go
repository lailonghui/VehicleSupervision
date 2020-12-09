package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/modules/vehicle/graph/generated"
	"VehicleSupervision/internal/modules/vehicle/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) DeleteVehicleInfo(ctx context.Context, where model.VehicleInfoBoolExp) (*model.VehicleInfoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteVehicleInfoByPk(ctx context.Context, id int64, vehicleID string) (*model.VehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertVehicleInfo(ctx context.Context, objects []*model.VehicleInfoInsertInput, onConflict *model.VehicleInfoOnConflict) (*model.VehicleInfoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertVehicleInfoOne(ctx context.Context, object model.VehicleInfoInsertInput, onConflict *model.VehicleInfoOnConflict) (*model.VehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateVehicleInfo(ctx context.Context, inc *model.VehicleInfoIncInput, set *model.VehicleInfoSetInput, where model.VehicleInfoBoolExp) (*model.VehicleInfoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateVehicleInfoByPk(ctx context.Context, inc *model.VehicleInfoIncInput, set *model.VehicleInfoSetInput, pkColumns model.VehicleInfoPkColumnsInput) (*model.VehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) VehicleInfo(ctx context.Context, distinctOn []model.VehicleInfoSelectColumn, limit *int, offset *int, orderBy []*model.VehicleInfoOrderBy, where *model.VehicleInfoBoolExp) ([]*model.VehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) VehicleInfoAggregate(ctx context.Context, distinctOn []model.VehicleInfoSelectColumn, limit *int, offset *int, orderBy []*model.VehicleInfoOrderBy, where *model.VehicleInfoBoolExp) (*model.VehicleInfoAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) VehicleInfoByPk(ctx context.Context, id int64, vehicleID string) (*model.VehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VehicleInfo(ctx context.Context, distinctOn []model.VehicleInfoSelectColumn, limit *int, offset *int, orderBy []*model.VehicleInfoOrderBy, where *model.VehicleInfoBoolExp) (<-chan []*model.VehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VehicleInfoAggregate(ctx context.Context, distinctOn []model.VehicleInfoSelectColumn, limit *int, offset *int, orderBy []*model.VehicleInfoOrderBy, where *model.VehicleInfoBoolExp) (<-chan *model.VehicleInfoAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VehicleInfoByPk(ctx context.Context, id int64, vehicleID string) (<-chan *model.VehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}
