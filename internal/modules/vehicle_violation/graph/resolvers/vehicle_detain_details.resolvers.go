package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/modules/vehicle_violation/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) DeleteVehicleDetainDetails(ctx context.Context, where model.VehicleDetainDetailsBoolExp) (*model.VehicleDetainDetailsMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteVehicleDetainDetailsByPk(ctx context.Context, id int64) (*model.VehicleDetainDetails, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertVehicleDetainDetails(ctx context.Context, objects []*model.VehicleDetainDetailsInsertInput, onConflict *model.VehicleDetainDetailsOnConflict) (*model.VehicleDetainDetailsMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertVehicleDetainDetailsOne(ctx context.Context, object model.VehicleDetainDetailsInsertInput, onConflict *model.VehicleDetainDetailsOnConflict) (*model.VehicleDetainDetails, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateVehicleDetainDetails(ctx context.Context, inc *model.VehicleDetainDetailsIncInput, set *model.VehicleDetainDetailsSetInput, where model.VehicleDetainDetailsBoolExp) (*model.VehicleDetainDetailsMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateVehicleDetainDetailsByPk(ctx context.Context, inc *model.VehicleDetainDetailsIncInput, set *model.VehicleDetainDetailsSetInput, pkColumns model.VehicleDetainDetailsPkColumnsInput) (*model.VehicleDetainDetails, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) VehicleDetainDetails(ctx context.Context, distinctOn []model.VehicleDetainDetailsSelectColumn, limit *int, offset *int, orderBy []*model.VehicleDetainDetailsOrderBy, where *model.VehicleDetainDetailsBoolExp) ([]*model.VehicleDetainDetails, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) VehicleDetainDetailsAggregate(ctx context.Context, distinctOn []model.VehicleDetainDetailsSelectColumn, limit *int, offset *int, orderBy []*model.VehicleDetainDetailsOrderBy, where *model.VehicleDetainDetailsBoolExp) (*model.VehicleDetainDetailsAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) VehicleDetainDetailsByPk(ctx context.Context, id int64) (*model.VehicleDetainDetails, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VehicleDetainDetails(ctx context.Context, distinctOn []model.VehicleDetainDetailsSelectColumn, limit *int, offset *int, orderBy []*model.VehicleDetainDetailsOrderBy, where *model.VehicleDetainDetailsBoolExp) (<-chan []*model.VehicleDetainDetails, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VehicleDetainDetailsAggregate(ctx context.Context, distinctOn []model.VehicleDetainDetailsSelectColumn, limit *int, offset *int, orderBy []*model.VehicleDetainDetailsOrderBy, where *model.VehicleDetainDetailsBoolExp) (<-chan *model.VehicleDetainDetailsAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VehicleDetainDetailsByPk(ctx context.Context, id int64) (<-chan *model.VehicleDetainDetails, error) {
	panic(fmt.Errorf("not implemented"))
}
