package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/modules/admin3/graph/generated"
	"VehicleSupervision/internal/modules/admin3/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) DeleteEnterprise(ctx context.Context, where model.EnterpriseBoolExp) (*model.EnterpriseMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteEnterpriseByPk(ctx context.Context, id string) (*model.Enterprise, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertEnterprise(ctx context.Context, objects []*model.EnterpriseInsertInput, onConflict *model.EnterpriseOnConflict) (*model.EnterpriseMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertEnterpriseOne(ctx context.Context, object model.EnterpriseInsertInput, onConflict *model.EnterpriseOnConflict) (*model.Enterprise, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateEnterprise(ctx context.Context, inc *model.EnterpriseIncInput, set *model.EnterpriseSetInput, where model.EnterpriseBoolExp) (*model.EnterpriseMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateEnterpriseByPk(ctx context.Context, inc *model.EnterpriseIncInput, set *model.EnterpriseSetInput, pkColumns model.EnterprisePkColumnsInput) (*model.Enterprise, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Enterprise(ctx context.Context, distinctOn []model.EnterpriseSelectColumn, limit *int, offset *int, orderBy []*model.EnterpriseOrderBy, where *model.EnterpriseBoolExp) ([]*model.Enterprise, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) EnterpriseAggregate(ctx context.Context, distinctOn []model.EnterpriseSelectColumn, limit *int, offset *int, orderBy []*model.EnterpriseOrderBy, where *model.EnterpriseBoolExp) (*model.EnterpriseAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) EnterpriseByPk(ctx context.Context, id string) (*model.Enterprise, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
