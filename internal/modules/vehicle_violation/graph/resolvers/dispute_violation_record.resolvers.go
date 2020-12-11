package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/modules/vehicle_violation/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) DeleteDisputeViolationRecord(ctx context.Context, where model.DisputeViolationRecordBoolExp) (*model.DisputeViolationRecordMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteDisputeViolationRecordByPk(ctx context.Context, disputeViolationID string, id int64) (*model.DisputeViolationRecord, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertDisputeViolationRecord(ctx context.Context, objects []*model.DisputeViolationRecordInsertInput, onConflict *model.DisputeViolationRecordOnConflict) (*model.DisputeViolationRecordMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertDisputeViolationRecordOne(ctx context.Context, object model.DisputeViolationRecordInsertInput, onConflict *model.DisputeViolationRecordOnConflict) (*model.DisputeViolationRecord, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateDisputeViolationRecord(ctx context.Context, inc *model.DisputeViolationRecordIncInput, set *model.DisputeViolationRecordSetInput, where model.DisputeViolationRecordBoolExp) (*model.DisputeViolationRecordMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateDisputeViolationRecordByPk(ctx context.Context, inc *model.DisputeViolationRecordIncInput, set *model.DisputeViolationRecordSetInput, pkColumns model.DisputeViolationRecordPkColumnsInput) (*model.DisputeViolationRecord, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DisputeViolationRecord(ctx context.Context, distinctOn []model.DisputeViolationRecordSelectColumn, limit *int, offset *int, orderBy []*model.DisputeViolationRecordOrderBy, where *model.DisputeViolationRecordBoolExp) ([]*model.DisputeViolationRecord, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DisputeViolationRecordAggregate(ctx context.Context, distinctOn []model.DisputeViolationRecordSelectColumn, limit *int, offset *int, orderBy []*model.DisputeViolationRecordOrderBy, where *model.DisputeViolationRecordBoolExp) (*model.DisputeViolationRecordAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DisputeViolationRecordByPk(ctx context.Context, disputeViolationID string, id int64) (*model.DisputeViolationRecord, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DisputeViolationRecord(ctx context.Context, distinctOn []model.DisputeViolationRecordSelectColumn, limit *int, offset *int, orderBy []*model.DisputeViolationRecordOrderBy, where *model.DisputeViolationRecordBoolExp) (<-chan []*model.DisputeViolationRecord, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DisputeViolationRecordAggregate(ctx context.Context, distinctOn []model.DisputeViolationRecordSelectColumn, limit *int, offset *int, orderBy []*model.DisputeViolationRecordOrderBy, where *model.DisputeViolationRecordBoolExp) (<-chan *model.DisputeViolationRecordAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DisputeViolationRecordByPk(ctx context.Context, disputeViolationID string, id int64) (<-chan *model.DisputeViolationRecord, error) {
	panic(fmt.Errorf("not implemented"))
}
