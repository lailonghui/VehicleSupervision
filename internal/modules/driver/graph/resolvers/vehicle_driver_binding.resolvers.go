package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/modules/driver/graph/generated"
	"VehicleSupervision/internal/modules/driver/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) DeleteVehicleDriverBinding(ctx context.Context, where model.VehicleDriverBindingBoolExp) (*model.VehicleDriverBindingMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteVehicleDriverBindingByPk(ctx context.Context, id int64) (*model.VehicleDriverBinding, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertVehicleDriverBinding(ctx context.Context, objects []*model.VehicleDriverBindingInsertInput, onConflict *model.VehicleDriverBindingOnConflict) (*model.VehicleDriverBindingMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertVehicleDriverBindingOne(ctx context.Context, object model.VehicleDriverBindingInsertInput, onConflict *model.VehicleDriverBindingOnConflict) (*model.VehicleDriverBinding, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateVehicleDriverBinding(ctx context.Context, inc *model.VehicleDriverBindingIncInput, set *model.VehicleDriverBindingSetInput, where model.VehicleDriverBindingBoolExp) (*model.VehicleDriverBindingMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateVehicleDriverBindingByPk(ctx context.Context, inc *model.VehicleDriverBindingIncInput, set *model.VehicleDriverBindingSetInput, pkColumns model.VehicleDriverBindingPkColumnsInput) (*model.VehicleDriverBinding, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) VehicleDriverBinding(ctx context.Context, distinctOn []model.VehicleDriverBindingSelectColumn, limit *int, offset *int, orderBy []*model.VehicleDriverBindingOrderBy, where *model.VehicleDriverBindingBoolExp) ([]*model.VehicleDriverBinding, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) VehicleDriverBindingAggregate(ctx context.Context, distinctOn []model.VehicleDriverBindingSelectColumn, limit *int, offset *int, orderBy []*model.VehicleDriverBindingOrderBy, where *model.VehicleDriverBindingBoolExp) (*model.VehicleDriverBindingAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) VehicleDriverBindingByPk(ctx context.Context, id int64) (*model.VehicleDriverBinding, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VehicleDriverBinding(ctx context.Context, distinctOn []model.VehicleDriverBindingSelectColumn, limit *int, offset *int, orderBy []*model.VehicleDriverBindingOrderBy, where *model.VehicleDriverBindingBoolExp) (<-chan []*model.VehicleDriverBinding, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VehicleDriverBindingAggregate(ctx context.Context, distinctOn []model.VehicleDriverBindingSelectColumn, limit *int, offset *int, orderBy []*model.VehicleDriverBindingOrderBy, where *model.VehicleDriverBindingBoolExp) (<-chan *model.VehicleDriverBindingAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VehicleDriverBindingByPk(ctx context.Context, id int64) (<-chan *model.VehicleDriverBinding, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
