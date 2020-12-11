package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/modules/vehicle_violation/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) DeleteVioCodewfdm(ctx context.Context, where model.VioCodewfdmBoolExp) (*model.VioCodewfdmMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertVioCodewfdm(ctx context.Context, objects []*model.VioCodewfdmInsertInput) (*model.VioCodewfdmMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertVioCodewfdmOne(ctx context.Context, object model.VioCodewfdmInsertInput) (*model.VioCodewfdm, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateVioCodewfdm(ctx context.Context, inc *model.VioCodewfdmIncInput, set *model.VioCodewfdmSetInput, where model.VioCodewfdmBoolExp) (*model.VioCodewfdmMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) VioCodewfdm(ctx context.Context, distinctOn []model.VioCodewfdmSelectColumn, limit *int, offset *int, orderBy []*model.VioCodewfdmOrderBy, where *model.VioCodewfdmBoolExp) ([]*model.VioCodewfdm, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) VioCodewfdmAggregate(ctx context.Context, distinctOn []model.VioCodewfdmSelectColumn, limit *int, offset *int, orderBy []*model.VioCodewfdmOrderBy, where *model.VioCodewfdmBoolExp) (*model.VioCodewfdmAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VioCodewfdm(ctx context.Context, distinctOn []model.VioCodewfdmSelectColumn, limit *int, offset *int, orderBy []*model.VioCodewfdmOrderBy, where *model.VioCodewfdmBoolExp) (<-chan []*model.VioCodewfdm, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VioCodewfdmAggregate(ctx context.Context, distinctOn []model.VioCodewfdmSelectColumn, limit *int, offset *int, orderBy []*model.VioCodewfdmOrderBy, where *model.VioCodewfdmBoolExp) (<-chan *model.VioCodewfdmAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}
