package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/modules/dynamic_supervision/graph/generated"
	"VehicleSupervision/internal/modules/dynamic_supervision/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) DeleteDynamicSupervision(ctx context.Context, where model.DynamicSupervisionBoolExp) (*model.DynamicSupervisionMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteDynamicSupervisionByPk(ctx context.Context, id int64, supervisionID string) (*model.DynamicSupervision, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertDynamicSupervision(ctx context.Context, objects []*model.DynamicSupervisionInsertInput, onConflict *model.DynamicSupervisionOnConflict) (*model.DynamicSupervisionMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertDynamicSupervisionOne(ctx context.Context, object model.DynamicSupervisionInsertInput, onConflict *model.DynamicSupervisionOnConflict) (*model.DynamicSupervision, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateDynamicSupervision(ctx context.Context, inc *model.DynamicSupervisionIncInput, set *model.DynamicSupervisionSetInput, where model.DynamicSupervisionBoolExp) (*model.DynamicSupervisionMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateDynamicSupervisionByPk(ctx context.Context, inc *model.DynamicSupervisionIncInput, set *model.DynamicSupervisionSetInput, pkColumns model.DynamicSupervisionPkColumnsInput) (*model.DynamicSupervision, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DynamicSupervision(ctx context.Context, distinctOn []model.DynamicSupervisionSelectColumn, limit *int, offset *int, orderBy []*model.DynamicSupervisionOrderBy, where *model.DynamicSupervisionBoolExp) ([]*model.DynamicSupervision, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DynamicSupervisionAggregate(ctx context.Context, distinctOn []model.DynamicSupervisionSelectColumn, limit *int, offset *int, orderBy []*model.DynamicSupervisionOrderBy, where *model.DynamicSupervisionBoolExp) (*model.DynamicSupervisionAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DynamicSupervisionByPk(ctx context.Context, id int64, supervisionID string) (*model.DynamicSupervision, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DynamicSupervision(ctx context.Context, distinctOn []model.DynamicSupervisionSelectColumn, limit *int, offset *int, orderBy []*model.DynamicSupervisionOrderBy, where *model.DynamicSupervisionBoolExp) (<-chan []*model.DynamicSupervision, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DynamicSupervisionAggregate(ctx context.Context, distinctOn []model.DynamicSupervisionSelectColumn, limit *int, offset *int, orderBy []*model.DynamicSupervisionOrderBy, where *model.DynamicSupervisionBoolExp) (<-chan *model.DynamicSupervisionAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DynamicSupervisionByPk(ctx context.Context, id int64, supervisionID string) (<-chan *model.DynamicSupervision, error) {
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
