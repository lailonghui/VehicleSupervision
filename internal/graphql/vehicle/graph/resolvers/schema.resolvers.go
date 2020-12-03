package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"lai.com/gqlgen_study/keyVehicleSupervision03/graph/generated"
	"lai.com/gqlgen_study/keyVehicleSupervision03/graph/model"
)

func (r *mutationResolver) DeleteVehicleInfo(ctx context.Context, where model.VehicleInfoBoolExp) (*model.VehicleInfoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteVehicleInfoByPk(ctx context.Context, id int64, vehicleID string) (*model.VehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertVehicleInfo(ctx context.Context, objects []*model.VehicleInfoInsertInput, onConflict *model.VehicleInfoOnConflict) (*model.VehicleInfoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertVehicleInfoOne(ctx context.Context, object model.VehicleInfoInsertInput, onConflict *model.VehicleInfoOnConflict) (*model.VehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateVehicleInfo(ctx context.Context, inc *model.VehicleInfoIncInput, set *model.VehicleInfoSetInput, where model.VehicleInfoBoolExp) (*model.VehicleInfoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateVehicleInfoByPk(ctx context.Context, inc *model.VehicleInfoIncInput, set *model.VehicleInfoSetInput, pkColumns model.VehicleInfoPkColumnsInput) (*model.VehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) VehicleInfo(ctx context.Context, distinctOn []model.VehicleInfoSelectColumn, limit *int, offset *int, orderBy []*model.VehicleInfoOrderBy, where *model.VehicleInfoBoolExp) ([]*model.VehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) VehicleInfoAggregate(ctx context.Context, distinctOn []model.VehicleInfoSelectColumn, limit *int, offset *int, orderBy []*model.VehicleInfoOrderBy, where *model.VehicleInfoBoolExp) (*model.VehicleInfoAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) VehicleInfoByPk(ctx context.Context, id int64, vehicleID string) (*model.VehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VehicleInfo(ctx context.Context, distinctOn []model.VehicleInfoSelectColumn, limit *int, offset *int, orderBy []*model.VehicleInfoOrderBy, where *model.VehicleInfoBoolExp) (<-chan []*model.VehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VehicleInfoAggregate(ctx context.Context, distinctOn []model.VehicleInfoSelectColumn, limit *int, offset *int, orderBy []*model.VehicleInfoOrderBy, where *model.VehicleInfoBoolExp) (<-chan *model.VehicleInfoAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VehicleInfoByPk(ctx context.Context, id int64, vehicleID string) (<-chan *model.VehicleInfo, error) {
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
