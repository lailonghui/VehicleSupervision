package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/modules/vehicle_violation/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) DeleteVehicleViolationScoringItems(ctx context.Context, where model.VehicleViolationScoringItemsBoolExp) (*model.VehicleViolationScoringItemsMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteVehicleViolationScoringItemsByPk(ctx context.Context, id int64, violationScoringItemID string) (*model.VehicleViolationScoringItems, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertVehicleViolationScoringItems(ctx context.Context, objects []*model.VehicleViolationScoringItemsInsertInput, onConflict *model.VehicleViolationScoringItemsOnConflict) (*model.VehicleViolationScoringItemsMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertVehicleViolationScoringItemsOne(ctx context.Context, object model.VehicleViolationScoringItemsInsertInput, onConflict *model.VehicleViolationScoringItemsOnConflict) (*model.VehicleViolationScoringItems, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateVehicleViolationScoringItems(ctx context.Context, inc *model.VehicleViolationScoringItemsIncInput, set *model.VehicleViolationScoringItemsSetInput, where model.VehicleViolationScoringItemsBoolExp) (*model.VehicleViolationScoringItemsMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateVehicleViolationScoringItemsByPk(ctx context.Context, inc *model.VehicleViolationScoringItemsIncInput, set *model.VehicleViolationScoringItemsSetInput, pkColumns model.VehicleViolationScoringItemsPkColumnsInput) (*model.VehicleViolationScoringItems, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) VehicleViolationScoringItems(ctx context.Context, distinctOn []model.VehicleViolationScoringItemsSelectColumn, limit *int, offset *int, orderBy []*model.VehicleViolationScoringItemsOrderBy, where *model.VehicleViolationScoringItemsBoolExp) ([]*model.VehicleViolationScoringItems, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) VehicleViolationScoringItemsAggregate(ctx context.Context, distinctOn []model.VehicleViolationScoringItemsSelectColumn, limit *int, offset *int, orderBy []*model.VehicleViolationScoringItemsOrderBy, where *model.VehicleViolationScoringItemsBoolExp) (*model.VehicleViolationScoringItemsAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) VehicleViolationScoringItemsByPk(ctx context.Context, id int64, violationScoringItemID string) (*model.VehicleViolationScoringItems, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VehicleViolationScoringItems(ctx context.Context, distinctOn []model.VehicleViolationScoringItemsSelectColumn, limit *int, offset *int, orderBy []*model.VehicleViolationScoringItemsOrderBy, where *model.VehicleViolationScoringItemsBoolExp) (<-chan []*model.VehicleViolationScoringItems, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VehicleViolationScoringItemsAggregate(ctx context.Context, distinctOn []model.VehicleViolationScoringItemsSelectColumn, limit *int, offset *int, orderBy []*model.VehicleViolationScoringItemsOrderBy, where *model.VehicleViolationScoringItemsBoolExp) (<-chan *model.VehicleViolationScoringItemsAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VehicleViolationScoringItemsByPk(ctx context.Context, id int64, violationScoringItemID string) (<-chan *model.VehicleViolationScoringItems, error) {
	panic(fmt.Errorf("not implemented"))
}
