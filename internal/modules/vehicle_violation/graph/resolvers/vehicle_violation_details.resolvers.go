package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/modules/vehicle_violation/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) DeleteVehicleViolationDetails(ctx context.Context, where model.VehicleViolationDetailsBoolExp) (*model.VehicleViolationDetailsMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteVehicleViolationDetailsByPk(ctx context.Context, id int64) (*model.VehicleViolationDetails, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertVehicleViolationDetails(ctx context.Context, objects []*model.VehicleViolationDetailsInsertInput, onConflict *model.VehicleViolationDetailsOnConflict) (*model.VehicleViolationDetailsMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertVehicleViolationDetailsOne(ctx context.Context, object model.VehicleViolationDetailsInsertInput, onConflict *model.VehicleViolationDetailsOnConflict) (*model.VehicleViolationDetails, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateVehicleViolationDetails(ctx context.Context, inc *model.VehicleViolationDetailsIncInput, set *model.VehicleViolationDetailsSetInput, where model.VehicleViolationDetailsBoolExp) (*model.VehicleViolationDetailsMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateVehicleViolationDetailsByPk(ctx context.Context, inc *model.VehicleViolationDetailsIncInput, set *model.VehicleViolationDetailsSetInput, pkColumns model.VehicleViolationDetailsPkColumnsInput) (*model.VehicleViolationDetails, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) VehicleViolationDetails(ctx context.Context, distinctOn []model.VehicleViolationDetailsSelectColumn, limit *int, offset *int, orderBy []*model.VehicleViolationDetailsOrderBy, where *model.VehicleViolationDetailsBoolExp) ([]*model.VehicleViolationDetails, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) VehicleViolationDetailsAggregate(ctx context.Context, distinctOn []model.VehicleViolationDetailsSelectColumn, limit *int, offset *int, orderBy []*model.VehicleViolationDetailsOrderBy, where *model.VehicleViolationDetailsBoolExp) (*model.VehicleViolationDetailsAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) VehicleViolationDetailsByPk(ctx context.Context, id int64) (*model.VehicleViolationDetails, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VehicleViolationDetails(ctx context.Context, distinctOn []model.VehicleViolationDetailsSelectColumn, limit *int, offset *int, orderBy []*model.VehicleViolationDetailsOrderBy, where *model.VehicleViolationDetailsBoolExp) (<-chan []*model.VehicleViolationDetails, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VehicleViolationDetailsAggregate(ctx context.Context, distinctOn []model.VehicleViolationDetailsSelectColumn, limit *int, offset *int, orderBy []*model.VehicleViolationDetailsOrderBy, where *model.VehicleViolationDetailsBoolExp) (<-chan *model.VehicleViolationDetailsAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VehicleViolationDetailsByPk(ctx context.Context, id int64) (<-chan *model.VehicleViolationDetails, error) {
	panic(fmt.Errorf("not implemented"))
}
