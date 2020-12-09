package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/modules/vehicle/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) DeleteOwnerInfo(ctx context.Context, where model.OwnerInfoBoolExp) (*model.OwnerInfoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteOwnerInfoByPk(ctx context.Context, id int64) (*model.OwnerInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertOwnerInfo(ctx context.Context, objects []*model.OwnerInfoInsertInput, onConflict *model.OwnerInfoOnConflict) (*model.OwnerInfoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertOwnerInfoOne(ctx context.Context, object model.OwnerInfoInsertInput, onConflict *model.OwnerInfoOnConflict) (*model.OwnerInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateOwnerInfo(ctx context.Context, inc *model.OwnerInfoIncInput, set *model.OwnerInfoSetInput, where model.OwnerInfoBoolExp) (*model.OwnerInfoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateOwnerInfoByPk(ctx context.Context, inc *model.OwnerInfoIncInput, set *model.OwnerInfoSetInput, pkColumns model.OwnerInfoPkColumnsInput) (*model.OwnerInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) OwnerInfo(ctx context.Context, distinctOn []model.OwnerInfoSelectColumn, limit *int, offset *int, orderBy []*model.OwnerInfoOrderBy, where *model.OwnerInfoBoolExp) ([]*model.OwnerInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) OwnerInfoAggregate(ctx context.Context, distinctOn []model.OwnerInfoSelectColumn, limit *int, offset *int, orderBy []*model.OwnerInfoOrderBy, where *model.OwnerInfoBoolExp) (*model.OwnerInfoAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) OwnerInfoByPk(ctx context.Context, id int64) (*model.OwnerInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) OwnerInfo(ctx context.Context, distinctOn []model.OwnerInfoSelectColumn, limit *int, offset *int, orderBy []*model.OwnerInfoOrderBy, where *model.OwnerInfoBoolExp) (<-chan []*model.OwnerInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) OwnerInfoAggregate(ctx context.Context, distinctOn []model.OwnerInfoSelectColumn, limit *int, offset *int, orderBy []*model.OwnerInfoOrderBy, where *model.OwnerInfoBoolExp) (<-chan *model.OwnerInfoAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) OwnerInfoByPk(ctx context.Context, id int64) (<-chan *model.OwnerInfo, error) {
	panic(fmt.Errorf("not implemented"))
}
