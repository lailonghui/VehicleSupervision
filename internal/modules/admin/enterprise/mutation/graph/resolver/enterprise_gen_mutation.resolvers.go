package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	model1 "VehicleSupervision/internal/modules/admin/enterprise/model"
	"VehicleSupervision/internal/modules/admin/enterprise/mutation/graph/generated"
	"VehicleSupervision/internal/modules/admin/enterprise/mutation/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) DeleteEnterprise(ctx context.Context, where model.EnterpriseBoolExp) (*model.EnterpriseMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteEnterpriseByPk(ctx context.Context, id int64) (*model1.Enterprise, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertEnterprise(ctx context.Context, objects []*model.EnterpriseInsertInput) (*model.EnterpriseMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertEnterpriseOne(ctx context.Context, object model.EnterpriseInsertInput) (*model1.Enterprise, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateEnterprise(ctx context.Context, inc *model.EnterpriseIncInput, set *model.EnterpriseSetInput, where model.EnterpriseBoolExp) (*model.EnterpriseMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateEnterpriseByPk(ctx context.Context, inc *model.EnterpriseIncInput, set *model.EnterpriseSetInput, pkColumns model.EnterprisePkColumnsInput) (*model1.Enterprise, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
