package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/modules/vehicle_violation/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) DeleteDriverPeccancyCheck(ctx context.Context, where model.DriverPeccancyCheckBoolExp) (*model.DriverPeccancyCheckMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteDriverPeccancyCheckByPk(ctx context.Context, id int64) (*model.DriverPeccancyCheck, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertDriverPeccancyCheck(ctx context.Context, objects []*model.DriverPeccancyCheckInsertInput, onConflict *model.DriverPeccancyCheckOnConflict) (*model.DriverPeccancyCheckMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertDriverPeccancyCheckOne(ctx context.Context, object model.DriverPeccancyCheckInsertInput, onConflict *model.DriverPeccancyCheckOnConflict) (*model.DriverPeccancyCheck, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateDriverPeccancyCheck(ctx context.Context, inc *model.DriverPeccancyCheckIncInput, set *model.DriverPeccancyCheckSetInput, where model.DriverPeccancyCheckBoolExp) (*model.DriverPeccancyCheckMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateDriverPeccancyCheckByPk(ctx context.Context, inc *model.DriverPeccancyCheckIncInput, set *model.DriverPeccancyCheckSetInput, pkColumns model.DriverPeccancyCheckPkColumnsInput) (*model.DriverPeccancyCheck, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DriverPeccancyCheck(ctx context.Context, distinctOn []model.DriverPeccancyCheckSelectColumn, limit *int, offset *int, orderBy []*model.DriverPeccancyCheckOrderBy, where *model.DriverPeccancyCheckBoolExp) ([]*model.DriverPeccancyCheck, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DriverPeccancyCheckAggregate(ctx context.Context, distinctOn []model.DriverPeccancyCheckSelectColumn, limit *int, offset *int, orderBy []*model.DriverPeccancyCheckOrderBy, where *model.DriverPeccancyCheckBoolExp) (*model.DriverPeccancyCheckAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DriverPeccancyCheckByPk(ctx context.Context, id int64) (*model.DriverPeccancyCheck, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DriverPeccancyCheck(ctx context.Context, distinctOn []model.DriverPeccancyCheckSelectColumn, limit *int, offset *int, orderBy []*model.DriverPeccancyCheckOrderBy, where *model.DriverPeccancyCheckBoolExp) (<-chan []*model.DriverPeccancyCheck, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DriverPeccancyCheckAggregate(ctx context.Context, distinctOn []model.DriverPeccancyCheckSelectColumn, limit *int, offset *int, orderBy []*model.DriverPeccancyCheckOrderBy, where *model.DriverPeccancyCheckBoolExp) (<-chan *model.DriverPeccancyCheckAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DriverPeccancyCheckByPk(ctx context.Context, id int64) (<-chan *model.DriverPeccancyCheck, error) {
	panic(fmt.Errorf("not implemented"))
}
