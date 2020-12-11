package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/modules/driver/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) DeleteDriverInfoChangeLog(ctx context.Context, where model.DriverInfoChangeLogBoolExp) (*model.DriverInfoChangeLogMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteDriverInfoChangeLogByPk(ctx context.Context, driverID string, id int64) (*model.DriverInfoChangeLog, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertDriverInfoChangeLog(ctx context.Context, objects []*model.DriverInfoChangeLogInsertInput, onConflict *model.DriverInfoChangeLogOnConflict) (*model.DriverInfoChangeLogMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertDriverInfoChangeLogOne(ctx context.Context, object model.DriverInfoChangeLogInsertInput, onConflict *model.DriverInfoChangeLogOnConflict) (*model.DriverInfoChangeLog, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateDriverInfoChangeLog(ctx context.Context, inc *model.DriverInfoChangeLogIncInput, set *model.DriverInfoChangeLogSetInput, where model.DriverInfoChangeLogBoolExp) (*model.DriverInfoChangeLogMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateDriverInfoChangeLogByPk(ctx context.Context, inc *model.DriverInfoChangeLogIncInput, set *model.DriverInfoChangeLogSetInput, pkColumns model.DriverInfoChangeLogPkColumnsInput) (*model.DriverInfoChangeLog, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DriverInfoChangeLog(ctx context.Context, distinctOn []model.DriverInfoChangeLogSelectColumn, limit *int, offset *int, orderBy []*model.DriverInfoChangeLogOrderBy, where *model.DriverInfoChangeLogBoolExp) ([]*model.DriverInfoChangeLog, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DriverInfoChangeLogAggregate(ctx context.Context, distinctOn []model.DriverInfoChangeLogSelectColumn, limit *int, offset *int, orderBy []*model.DriverInfoChangeLogOrderBy, where *model.DriverInfoChangeLogBoolExp) (*model.DriverInfoChangeLogAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DriverInfoChangeLogByPk(ctx context.Context, driverID string, id int64) (*model.DriverInfoChangeLog, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DriverInfoChangeLog(ctx context.Context, distinctOn []model.DriverInfoChangeLogSelectColumn, limit *int, offset *int, orderBy []*model.DriverInfoChangeLogOrderBy, where *model.DriverInfoChangeLogBoolExp) (<-chan []*model.DriverInfoChangeLog, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DriverInfoChangeLogAggregate(ctx context.Context, distinctOn []model.DriverInfoChangeLogSelectColumn, limit *int, offset *int, orderBy []*model.DriverInfoChangeLogOrderBy, where *model.DriverInfoChangeLogBoolExp) (<-chan *model.DriverInfoChangeLogAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DriverInfoChangeLogByPk(ctx context.Context, driverID string, id int64) (<-chan *model.DriverInfoChangeLog, error) {
	panic(fmt.Errorf("not implemented"))
}
