package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/modules/vehicle/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) DeleteJjVehicle(ctx context.Context, where model.JjVehicleBoolExp) (*model.JjVehicleMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteJjVehicleByPk(ctx context.Context, id int64) (*model.JjVehicle, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertJjVehicle(ctx context.Context, objects []*model.JjVehicleInsertInput, onConflict *model.JjVehicleOnConflict) (*model.JjVehicleMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertJjVehicleOne(ctx context.Context, object model.JjVehicleInsertInput, onConflict *model.JjVehicleOnConflict) (*model.JjVehicle, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateJjVehicle(ctx context.Context, inc *model.JjVehicleIncInput, set *model.JjVehicleSetInput, where model.JjVehicleBoolExp) (*model.JjVehicleMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateJjVehicleByPk(ctx context.Context, inc *model.JjVehicleIncInput, set *model.JjVehicleSetInput, pkColumns model.JjVehiclePkColumnsInput) (*model.JjVehicle, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) JjVehicle(ctx context.Context, distinctOn []model.JjVehicleSelectColumn, limit *int, offset *int, orderBy []*model.JjVehicleOrderBy, where *model.JjVehicleBoolExp) ([]*model.JjVehicle, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) JjVehicleAggregate(ctx context.Context, distinctOn []model.JjVehicleSelectColumn, limit *int, offset *int, orderBy []*model.JjVehicleOrderBy, where *model.JjVehicleBoolExp) (*model.JjVehicleAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) JjVehicleByPk(ctx context.Context, id int64) (*model.JjVehicle, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) JjVehicle(ctx context.Context, distinctOn []model.JjVehicleSelectColumn, limit *int, offset *int, orderBy []*model.JjVehicleOrderBy, where *model.JjVehicleBoolExp) (<-chan []*model.JjVehicle, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) JjVehicleAggregate(ctx context.Context, distinctOn []model.JjVehicleSelectColumn, limit *int, offset *int, orderBy []*model.JjVehicleOrderBy, where *model.JjVehicleBoolExp) (<-chan *model.JjVehicleAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) JjVehicleByPk(ctx context.Context, id int64) (<-chan *model.JjVehicle, error) {
	panic(fmt.Errorf("not implemented"))
}
