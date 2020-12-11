package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/modules/vehicle_violation/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) DeleteUserOperationLog(ctx context.Context, where model.UserOperationLogBoolExp) (*model.UserOperationLogMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteUserOperationLogByPk(ctx context.Context, id int64) (*model.UserOperationLog, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertUserOperationLog(ctx context.Context, objects []*model.UserOperationLogInsertInput, onConflict *model.UserOperationLogOnConflict) (*model.UserOperationLogMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertUserOperationLogOne(ctx context.Context, object model.UserOperationLogInsertInput, onConflict *model.UserOperationLogOnConflict) (*model.UserOperationLog, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateUserOperationLog(ctx context.Context, inc *model.UserOperationLogIncInput, set *model.UserOperationLogSetInput, where model.UserOperationLogBoolExp) (*model.UserOperationLogMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateUserOperationLogByPk(ctx context.Context, inc *model.UserOperationLogIncInput, set *model.UserOperationLogSetInput, pkColumns model.UserOperationLogPkColumnsInput) (*model.UserOperationLog, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) UserOperationLog(ctx context.Context, distinctOn []model.UserOperationLogSelectColumn, limit *int, offset *int, orderBy []*model.UserOperationLogOrderBy, where *model.UserOperationLogBoolExp) ([]*model.UserOperationLog, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) UserOperationLogAggregate(ctx context.Context, distinctOn []model.UserOperationLogSelectColumn, limit *int, offset *int, orderBy []*model.UserOperationLogOrderBy, where *model.UserOperationLogBoolExp) (*model.UserOperationLogAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) UserOperationLogByPk(ctx context.Context, id int64) (*model.UserOperationLog, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) UserOperationLog(ctx context.Context, distinctOn []model.UserOperationLogSelectColumn, limit *int, offset *int, orderBy []*model.UserOperationLogOrderBy, where *model.UserOperationLogBoolExp) (<-chan []*model.UserOperationLog, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) UserOperationLogAggregate(ctx context.Context, distinctOn []model.UserOperationLogSelectColumn, limit *int, offset *int, orderBy []*model.UserOperationLogOrderBy, where *model.UserOperationLogBoolExp) (<-chan *model.UserOperationLogAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) UserOperationLogByPk(ctx context.Context, id int64) (<-chan *model.UserOperationLog, error) {
	panic(fmt.Errorf("not implemented"))
}
