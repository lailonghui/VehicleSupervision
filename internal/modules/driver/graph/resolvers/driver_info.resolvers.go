package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/modules/driver/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) DeleteDriverInfo(ctx context.Context, where model.DriverInfoBoolExp) (*model.DriverInfoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteDriverInfoByPk(ctx context.Context, driverID string, id int64) (*model.DriverInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertDriverInfo(ctx context.Context, objects []*model.DriverInfoInsertInput, onConflict *model.DriverInfoOnConflict) (*model.DriverInfoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertDriverInfoOne(ctx context.Context, object model.DriverInfoInsertInput, onConflict *model.DriverInfoOnConflict) (*model.DriverInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateDriverInfo(ctx context.Context, inc *model.DriverInfoIncInput, set *model.DriverInfoSetInput, where model.DriverInfoBoolExp) (*model.DriverInfoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateDriverInfoByPk(ctx context.Context, inc *model.DriverInfoIncInput, set *model.DriverInfoSetInput, pkColumns model.DriverInfoPkColumnsInput) (*model.DriverInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DriverInfo(ctx context.Context, distinctOn []model.DriverInfoSelectColumn, limit *int, offset *int, orderBy []*model.DriverInfoOrderBy, where *model.DriverInfoBoolExp) ([]*model.DriverInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DriverInfoAggregate(ctx context.Context, distinctOn []model.DriverInfoSelectColumn, limit *int, offset *int, orderBy []*model.DriverInfoOrderBy, where *model.DriverInfoBoolExp) (*model.DriverInfoAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DriverInfoByPk(ctx context.Context, driverID string, id int64) (*model.DriverInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DriverInfo(ctx context.Context, distinctOn []model.DriverInfoSelectColumn, limit *int, offset *int, orderBy []*model.DriverInfoOrderBy, where *model.DriverInfoBoolExp) (<-chan []*model.DriverInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DriverInfoAggregate(ctx context.Context, distinctOn []model.DriverInfoSelectColumn, limit *int, offset *int, orderBy []*model.DriverInfoOrderBy, where *model.DriverInfoBoolExp) (<-chan *model.DriverInfoAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DriverInfoByPk(ctx context.Context, driverID string, id int64) (<-chan *model.DriverInfo, error) {
	panic(fmt.Errorf("not implemented"))
}
