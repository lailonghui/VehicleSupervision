package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/modules/vehicle/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) DeleteMuckTruckWorkerIDCardOrders(ctx context.Context, where model.MuckTruckWorkerIDCardOrdersBoolExp) (*model.MuckTruckWorkerIDCardOrdersMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteMuckTruckWorkerIDCardOrdersByPk(ctx context.Context, id int64) (*model.MuckTruckWorkerIDCardOrders, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertMuckTruckWorkerIDCardOrders(ctx context.Context, objects []*model.MuckTruckWorkerIDCardOrdersInsertInput, onConflict *model.MuckTruckWorkerIDCardOrdersOnConflict) (*model.MuckTruckWorkerIDCardOrdersMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertMuckTruckWorkerIDCardOrdersOne(ctx context.Context, object model.MuckTruckWorkerIDCardOrdersInsertInput, onConflict *model.MuckTruckWorkerIDCardOrdersOnConflict) (*model.MuckTruckWorkerIDCardOrders, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateMuckTruckWorkerIDCardOrders(ctx context.Context, inc *model.MuckTruckWorkerIDCardOrdersIncInput, set *model.MuckTruckWorkerIDCardOrdersSetInput, where model.MuckTruckWorkerIDCardOrdersBoolExp) (*model.MuckTruckWorkerIDCardOrdersMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateMuckTruckWorkerIDCardOrdersByPk(ctx context.Context, inc *model.MuckTruckWorkerIDCardOrdersIncInput, set *model.MuckTruckWorkerIDCardOrdersSetInput, pkColumns model.MuckTruckWorkerIDCardOrdersPkColumnsInput) (*model.MuckTruckWorkerIDCardOrders, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) MuckTruckWorkerIDCardOrders(ctx context.Context, distinctOn []model.MuckTruckWorkerIDCardOrdersSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckWorkerIDCardOrdersOrderBy, where *model.MuckTruckWorkerIDCardOrdersBoolExp) ([]*model.MuckTruckWorkerIDCardOrders, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) MuckTruckWorkerIDCardOrdersAggregate(ctx context.Context, distinctOn []model.MuckTruckWorkerIDCardOrdersSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckWorkerIDCardOrdersOrderBy, where *model.MuckTruckWorkerIDCardOrdersBoolExp) (*model.MuckTruckWorkerIDCardOrdersAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) MuckTruckWorkerIDCardOrdersByPk(ctx context.Context, id int64) (*model.MuckTruckWorkerIDCardOrders, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) MuckTruckWorkerIDCardOrders(ctx context.Context, distinctOn []model.MuckTruckWorkerIDCardOrdersSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckWorkerIDCardOrdersOrderBy, where *model.MuckTruckWorkerIDCardOrdersBoolExp) (<-chan []*model.MuckTruckWorkerIDCardOrders, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) MuckTruckWorkerIDCardOrdersAggregate(ctx context.Context, distinctOn []model.MuckTruckWorkerIDCardOrdersSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckWorkerIDCardOrdersOrderBy, where *model.MuckTruckWorkerIDCardOrdersBoolExp) (<-chan *model.MuckTruckWorkerIDCardOrdersAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) MuckTruckWorkerIDCardOrdersByPk(ctx context.Context, id int64) (<-chan *model.MuckTruckWorkerIDCardOrders, error) {
	panic(fmt.Errorf("not implemented"))
}
