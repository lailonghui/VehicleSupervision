package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/modules/vehicle_violation/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) DeleteVehicleViolationScoringRecord(ctx context.Context, where model.VehicleViolationScoringRecordBoolExp) (*model.VehicleViolationScoringRecordMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteVehicleViolationScoringRecordByPk(ctx context.Context, id int64, violationScoringID string) (*model.VehicleViolationScoringRecord, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertVehicleViolationScoringRecord(ctx context.Context, objects []*model.VehicleViolationScoringRecordInsertInput, onConflict *model.VehicleViolationScoringRecordOnConflict) (*model.VehicleViolationScoringRecordMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertVehicleViolationScoringRecordOne(ctx context.Context, object model.VehicleViolationScoringRecordInsertInput, onConflict *model.VehicleViolationScoringRecordOnConflict) (*model.VehicleViolationScoringRecord, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateVehicleViolationScoringRecord(ctx context.Context, inc *model.VehicleViolationScoringRecordIncInput, set *model.VehicleViolationScoringRecordSetInput, where model.VehicleViolationScoringRecordBoolExp) (*model.VehicleViolationScoringRecordMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateVehicleViolationScoringRecordByPk(ctx context.Context, inc *model.VehicleViolationScoringRecordIncInput, set *model.VehicleViolationScoringRecordSetInput, pkColumns model.VehicleViolationScoringRecordPkColumnsInput) (*model.VehicleViolationScoringRecord, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) VehicleViolationScoringRecord(ctx context.Context, distinctOn []model.VehicleViolationScoringRecordSelectColumn, limit *int, offset *int, orderBy []*model.VehicleViolationScoringRecordOrderBy, where *model.VehicleViolationScoringRecordBoolExp) ([]*model.VehicleViolationScoringRecord, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) VehicleViolationScoringRecordAggregate(ctx context.Context, distinctOn []model.VehicleViolationScoringRecordSelectColumn, limit *int, offset *int, orderBy []*model.VehicleViolationScoringRecordOrderBy, where *model.VehicleViolationScoringRecordBoolExp) (*model.VehicleViolationScoringRecordAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) VehicleViolationScoringRecordByPk(ctx context.Context, id int64, violationScoringID string) (*model.VehicleViolationScoringRecord, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VehicleViolationScoringRecord(ctx context.Context, distinctOn []model.VehicleViolationScoringRecordSelectColumn, limit *int, offset *int, orderBy []*model.VehicleViolationScoringRecordOrderBy, where *model.VehicleViolationScoringRecordBoolExp) (<-chan []*model.VehicleViolationScoringRecord, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VehicleViolationScoringRecordAggregate(ctx context.Context, distinctOn []model.VehicleViolationScoringRecordSelectColumn, limit *int, offset *int, orderBy []*model.VehicleViolationScoringRecordOrderBy, where *model.VehicleViolationScoringRecordBoolExp) (<-chan *model.VehicleViolationScoringRecordAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VehicleViolationScoringRecordByPk(ctx context.Context, id int64, violationScoringID string) (<-chan *model.VehicleViolationScoringRecord, error) {
	panic(fmt.Errorf("not implemented"))
}
