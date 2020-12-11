package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/modules/dynamic_supervision/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) DeleteDynamicSpotCheckDisposal(ctx context.Context, where model.DynamicSpotCheckDisposalBoolExp) (*model.DynamicSpotCheckDisposalMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *mutationResolver) DeleteDynamicSpotCheckDisposalByPk(ctx context.Context, id int64) (*model.DynamicSpotCheckDisposal, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertDynamicSpotCheckDisposal(ctx context.Context, objects []*model.DynamicSpotCheckDisposalInsertInput, onConflict *model.DynamicSpotCheckDisposalOnConflict) (*model.DynamicSpotCheckDisposalMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertDynamicSpotCheckDisposalOne(ctx context.Context, object model.DynamicSpotCheckDisposalInsertInput, onConflict *model.DynamicSpotCheckDisposalOnConflict) (*model.DynamicSpotCheckDisposal, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateDynamicSpotCheckDisposal(ctx context.Context, inc *model.DynamicSpotCheckDisposalIncInput, set *model.DynamicSpotCheckDisposalSetInput, where model.DynamicSpotCheckDisposalBoolExp) (*model.DynamicSpotCheckDisposalMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateDynamicSpotCheckDisposalByPk(ctx context.Context, inc *model.DynamicSpotCheckDisposalIncInput, set *model.DynamicSpotCheckDisposalSetInput, pkColumns model.DynamicSpotCheckDisposalPkColumnsInput) (*model.DynamicSpotCheckDisposal, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DynamicSpotCheckDisposal(ctx context.Context, distinctOn []model.DynamicSpotCheckDisposalSelectColumn, limit *int, offset *int, orderBy []*model.DynamicSpotCheckDisposalOrderBy, where *model.DynamicSpotCheckDisposalBoolExp) ([]*model.DynamicSpotCheckDisposal, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DynamicSpotCheckDisposalAggregate(ctx context.Context, distinctOn []model.DynamicSpotCheckDisposalSelectColumn, limit *int, offset *int, orderBy []*model.DynamicSpotCheckDisposalOrderBy, where *model.DynamicSpotCheckDisposalBoolExp) (*model.DynamicSpotCheckDisposalAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DynamicSpotCheckDisposalByPk(ctx context.Context, id int64) (*model.DynamicSpotCheckDisposal, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DynamicSpotCheckDisposal(ctx context.Context, distinctOn []model.DynamicSpotCheckDisposalSelectColumn, limit *int, offset *int, orderBy []*model.DynamicSpotCheckDisposalOrderBy, where *model.DynamicSpotCheckDisposalBoolExp) (<-chan []*model.DynamicSpotCheckDisposal, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DynamicSpotCheckDisposalAggregate(ctx context.Context, distinctOn []model.DynamicSpotCheckDisposalSelectColumn, limit *int, offset *int, orderBy []*model.DynamicSpotCheckDisposalOrderBy, where *model.DynamicSpotCheckDisposalBoolExp) (<-chan *model.DynamicSpotCheckDisposalAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DynamicSpotCheckDisposalByPk(ctx context.Context, id int64) (<-chan *model.DynamicSpotCheckDisposal, error) {
	panic(fmt.Errorf("not implemented"))
}
