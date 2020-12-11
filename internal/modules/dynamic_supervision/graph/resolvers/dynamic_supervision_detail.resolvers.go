package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/modules/dynamic_supervision/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) DeleteDynamicSupervisionDetail(ctx context.Context, where model.DynamicSupervisionDetailBoolExp) (*model.DynamicSupervisionDetailMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteDynamicSupervisionDetailByPk(ctx context.Context, id int64, supervisionDetailID string) (*model.DynamicSupervisionDetail, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertDynamicSupervisionDetail(ctx context.Context, objects []*model.DynamicSupervisionDetailInsertInput, onConflict *model.DynamicSupervisionDetailOnConflict) (*model.DynamicSupervisionDetailMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertDynamicSupervisionDetailOne(ctx context.Context, object model.DynamicSupervisionDetailInsertInput, onConflict *model.DynamicSupervisionDetailOnConflict) (*model.DynamicSupervisionDetail, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateDynamicSupervisionDetail(ctx context.Context, inc *model.DynamicSupervisionDetailIncInput, set *model.DynamicSupervisionDetailSetInput, where model.DynamicSupervisionDetailBoolExp) (*model.DynamicSupervisionDetailMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateDynamicSupervisionDetailByPk(ctx context.Context, inc *model.DynamicSupervisionDetailIncInput, set *model.DynamicSupervisionDetailSetInput, pkColumns model.DynamicSupervisionDetailPkColumnsInput) (*model.DynamicSupervisionDetail, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DynamicSupervisionDetail(ctx context.Context, distinctOn []model.DynamicSupervisionDetailSelectColumn, limit *int, offset *int, orderBy []*model.DynamicSupervisionDetailOrderBy, where *model.DynamicSupervisionDetailBoolExp) ([]*model.DynamicSupervisionDetail, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DynamicSupervisionDetailAggregate(ctx context.Context, distinctOn []model.DynamicSupervisionDetailSelectColumn, limit *int, offset *int, orderBy []*model.DynamicSupervisionDetailOrderBy, where *model.DynamicSupervisionDetailBoolExp) (*model.DynamicSupervisionDetailAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DynamicSupervisionDetailByPk(ctx context.Context, id int64, supervisionDetailID string) (*model.DynamicSupervisionDetail, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DynamicSupervisionDetail(ctx context.Context, distinctOn []model.DynamicSupervisionDetailSelectColumn, limit *int, offset *int, orderBy []*model.DynamicSupervisionDetailOrderBy, where *model.DynamicSupervisionDetailBoolExp) (<-chan []*model.DynamicSupervisionDetail, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DynamicSupervisionDetailAggregate(ctx context.Context, distinctOn []model.DynamicSupervisionDetailSelectColumn, limit *int, offset *int, orderBy []*model.DynamicSupervisionDetailOrderBy, where *model.DynamicSupervisionDetailBoolExp) (<-chan *model.DynamicSupervisionDetailAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DynamicSupervisionDetailByPk(ctx context.Context, id int64, supervisionDetailID string) (<-chan *model.DynamicSupervisionDetail, error) {
	panic(fmt.Errorf("not implemented"))
}
