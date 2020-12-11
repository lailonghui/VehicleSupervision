package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/modules/vehicle_alarm/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) DeleteVehicleAlarmData(ctx context.Context, where model.VehicleAlarmDataBoolExp) (*model.VehicleAlarmDataMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteVehicleAlarmDataByPk(ctx context.Context, id int64, vehicleAlarmDataID string) (*model.VehicleAlarmData, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertVehicleAlarmData(ctx context.Context, objects []*model.VehicleAlarmDataInsertInput, onConflict *model.VehicleAlarmDataOnConflict) (*model.VehicleAlarmDataMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertVehicleAlarmDataOne(ctx context.Context, object model.VehicleAlarmDataInsertInput, onConflict *model.VehicleAlarmDataOnConflict) (*model.VehicleAlarmData, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateVehicleAlarmData(ctx context.Context, inc *model.VehicleAlarmDataIncInput, set *model.VehicleAlarmDataSetInput, where model.VehicleAlarmDataBoolExp) (*model.VehicleAlarmDataMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateVehicleAlarmDataByPk(ctx context.Context, inc *model.VehicleAlarmDataIncInput, set *model.VehicleAlarmDataSetInput, pkColumns model.VehicleAlarmDataPkColumnsInput) (*model.VehicleAlarmData, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) VehicleAlarmData(ctx context.Context, distinctOn []model.VehicleAlarmDataSelectColumn, limit *int, offset *int, orderBy []*model.VehicleAlarmDataOrderBy, where *model.VehicleAlarmDataBoolExp) ([]*model.VehicleAlarmData, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) VehicleAlarmDataAggregate(ctx context.Context, distinctOn []model.VehicleAlarmDataSelectColumn, limit *int, offset *int, orderBy []*model.VehicleAlarmDataOrderBy, where *model.VehicleAlarmDataBoolExp) (*model.VehicleAlarmDataAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) VehicleAlarmDataByPk(ctx context.Context, id int64, vehicleAlarmDataID string) (*model.VehicleAlarmData, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VehicleAlarmData(ctx context.Context, distinctOn []model.VehicleAlarmDataSelectColumn, limit *int, offset *int, orderBy []*model.VehicleAlarmDataOrderBy, where *model.VehicleAlarmDataBoolExp) (<-chan []*model.VehicleAlarmData, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VehicleAlarmDataAggregate(ctx context.Context, distinctOn []model.VehicleAlarmDataSelectColumn, limit *int, offset *int, orderBy []*model.VehicleAlarmDataOrderBy, where *model.VehicleAlarmDataBoolExp) (<-chan *model.VehicleAlarmDataAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VehicleAlarmDataByPk(ctx context.Context, id int64, vehicleAlarmDataID string) (<-chan *model.VehicleAlarmData, error) {
	panic(fmt.Errorf("not implemented"))
}
