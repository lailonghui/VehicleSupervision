package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/modules/vehicle_violation/graph/generated"
	"VehicleSupervision/internal/modules/vehicle_violation/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) DeleteAppEnforcement(ctx context.Context, where model.AppEnforcementBoolExp) (*model.AppEnforcementMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteAppEnforcementByPk(ctx context.Context, id int64, illegalPhotoID string) (*model.AppEnforcement, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertAppEnforcement(ctx context.Context, objects []*model.AppEnforcementInsertInput, onConflict *model.AppEnforcementOnConflict) (*model.AppEnforcementMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertAppEnforcementOne(ctx context.Context, object model.AppEnforcementInsertInput, onConflict *model.AppEnforcementOnConflict) (*model.AppEnforcement, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateAppEnforcement(ctx context.Context, inc *model.AppEnforcementIncInput, set *model.AppEnforcementSetInput, where model.AppEnforcementBoolExp) (*model.AppEnforcementMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateAppEnforcementByPk(ctx context.Context, inc *model.AppEnforcementIncInput, set *model.AppEnforcementSetInput, pkColumns model.AppEnforcementPkColumnsInput) (*model.AppEnforcement, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AppEnforcement(ctx context.Context, distinctOn []model.AppEnforcementSelectColumn, limit *int, offset *int, orderBy []*model.AppEnforcementOrderBy, where *model.AppEnforcementBoolExp) ([]*model.AppEnforcement, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AppEnforcementAggregate(ctx context.Context, distinctOn []model.AppEnforcementSelectColumn, limit *int, offset *int, orderBy []*model.AppEnforcementOrderBy, where *model.AppEnforcementBoolExp) (*model.AppEnforcementAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AppEnforcementByPk(ctx context.Context, id int64, illegalPhotoID string) (*model.AppEnforcement, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) AppEnforcement(ctx context.Context, distinctOn []model.AppEnforcementSelectColumn, limit *int, offset *int, orderBy []*model.AppEnforcementOrderBy, where *model.AppEnforcementBoolExp) (<-chan []*model.AppEnforcement, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) AppEnforcementAggregate(ctx context.Context, distinctOn []model.AppEnforcementSelectColumn, limit *int, offset *int, orderBy []*model.AppEnforcementOrderBy, where *model.AppEnforcementBoolExp) (<-chan *model.AppEnforcementAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) AppEnforcementByPk(ctx context.Context, id int64, illegalPhotoID string) (<-chan *model.AppEnforcement, error) {
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
