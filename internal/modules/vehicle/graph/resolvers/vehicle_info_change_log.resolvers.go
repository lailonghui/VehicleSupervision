package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/modules/vehicle/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) DeleteVehicleInfoChangeLog(ctx context.Context, where model.VehicleInfoChangeLogBoolExp) (*model.VehicleInfoChangeLogMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteVehicleInfoChangeLogByPk(ctx context.Context, id int64, vehicleInfoChangeID string) (*model.VehicleInfoChangeLog, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertVehicleInfoChangeLog(ctx context.Context, objects []*model.VehicleInfoChangeLogInsertInput, onConflict *model.VehicleInfoChangeLogOnConflict) (*model.VehicleInfoChangeLogMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertVehicleInfoChangeLogOne(ctx context.Context, object model.VehicleInfoChangeLogInsertInput, onConflict *model.VehicleInfoChangeLogOnConflict) (*model.VehicleInfoChangeLog, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateVehicleInfoChangeLog(ctx context.Context, inc *model.VehicleInfoChangeLogIncInput, set *model.VehicleInfoChangeLogSetInput, where model.VehicleInfoChangeLogBoolExp) (*model.VehicleInfoChangeLogMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateVehicleInfoChangeLogByPk(ctx context.Context, inc *model.VehicleInfoChangeLogIncInput, set *model.VehicleInfoChangeLogSetInput, pkColumns model.VehicleInfoChangeLogPkColumnsInput) (*model.VehicleInfoChangeLog, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) VehicleInfoChangeLog(ctx context.Context, distinctOn []model.VehicleInfoChangeLogSelectColumn, limit *int, offset *int, orderBy []*model.VehicleInfoChangeLogOrderBy, where *model.VehicleInfoChangeLogBoolExp) ([]*model.VehicleInfoChangeLog, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) VehicleInfoChangeLogAggregate(ctx context.Context, distinctOn []model.VehicleInfoChangeLogSelectColumn, limit *int, offset *int, orderBy []*model.VehicleInfoChangeLogOrderBy, where *model.VehicleInfoChangeLogBoolExp) (*model.VehicleInfoChangeLogAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) VehicleInfoChangeLogByPk(ctx context.Context, id int64, vehicleInfoChangeID string) (*model.VehicleInfoChangeLog, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VehicleInfoChangeLog(ctx context.Context, distinctOn []model.VehicleInfoChangeLogSelectColumn, limit *int, offset *int, orderBy []*model.VehicleInfoChangeLogOrderBy, where *model.VehicleInfoChangeLogBoolExp) (<-chan []*model.VehicleInfoChangeLog, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VehicleInfoChangeLogAggregate(ctx context.Context, distinctOn []model.VehicleInfoChangeLogSelectColumn, limit *int, offset *int, orderBy []*model.VehicleInfoChangeLogOrderBy, where *model.VehicleInfoChangeLogBoolExp) (<-chan *model.VehicleInfoChangeLogAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VehicleInfoChangeLogByPk(ctx context.Context, id int64, vehicleInfoChangeID string) (<-chan *model.VehicleInfoChangeLog, error) {
	panic(fmt.Errorf("not implemented"))
}
