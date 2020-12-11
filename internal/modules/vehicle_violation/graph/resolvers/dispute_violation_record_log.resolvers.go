package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/modules/vehicle_violation/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) DeleteDisputeViolationRecordLog(ctx context.Context, where model.DisputeViolationRecordLogBoolExp) (*model.DisputeViolationRecordLogMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteDisputeViolationRecordLogByPk(ctx context.Context, disputeViolationLogID string, id int64) (*model.DisputeViolationRecordLog, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertDisputeViolationRecordLog(ctx context.Context, objects []*model.DisputeViolationRecordLogInsertInput, onConflict *model.DisputeViolationRecordLogOnConflict) (*model.DisputeViolationRecordLogMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertDisputeViolationRecordLogOne(ctx context.Context, object model.DisputeViolationRecordLogInsertInput, onConflict *model.DisputeViolationRecordLogOnConflict) (*model.DisputeViolationRecordLog, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateDisputeViolationRecordLog(ctx context.Context, inc *model.DisputeViolationRecordLogIncInput, set *model.DisputeViolationRecordLogSetInput, where model.DisputeViolationRecordLogBoolExp) (*model.DisputeViolationRecordLogMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateDisputeViolationRecordLogByPk(ctx context.Context, inc *model.DisputeViolationRecordLogIncInput, set *model.DisputeViolationRecordLogSetInput, pkColumns model.DisputeViolationRecordLogPkColumnsInput) (*model.DisputeViolationRecordLog, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DisputeViolationRecordLog(ctx context.Context, distinctOn []model.DisputeViolationRecordLogSelectColumn, limit *int, offset *int, orderBy []*model.DisputeViolationRecordLogOrderBy, where *model.DisputeViolationRecordLogBoolExp) ([]*model.DisputeViolationRecordLog, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DisputeViolationRecordLogAggregate(ctx context.Context, distinctOn []model.DisputeViolationRecordLogSelectColumn, limit *int, offset *int, orderBy []*model.DisputeViolationRecordLogOrderBy, where *model.DisputeViolationRecordLogBoolExp) (*model.DisputeViolationRecordLogAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DisputeViolationRecordLogByPk(ctx context.Context, disputeViolationLogID string, id int64) (*model.DisputeViolationRecordLog, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DisputeViolationRecordLog(ctx context.Context, distinctOn []model.DisputeViolationRecordLogSelectColumn, limit *int, offset *int, orderBy []*model.DisputeViolationRecordLogOrderBy, where *model.DisputeViolationRecordLogBoolExp) (<-chan []*model.DisputeViolationRecordLog, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DisputeViolationRecordLogAggregate(ctx context.Context, distinctOn []model.DisputeViolationRecordLogSelectColumn, limit *int, offset *int, orderBy []*model.DisputeViolationRecordLogOrderBy, where *model.DisputeViolationRecordLogBoolExp) (<-chan *model.DisputeViolationRecordLogAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DisputeViolationRecordLogByPk(ctx context.Context, disputeViolationLogID string, id int64) (<-chan *model.DisputeViolationRecordLog, error) {
	panic(fmt.Errorf("not implemented"))
}
