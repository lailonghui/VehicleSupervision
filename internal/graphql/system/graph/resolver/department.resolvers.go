package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"system-manage/internal/graphql/system/graph/generated"
	"system-manage/internal/graphql/system/graph/model"
)

func (r *mutationResolver) CreateDepartment(ctx context.Context, input model.CreateDepartmentParam) (*model.Department, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateDepartment(ctx context.Context, input model.UpdateDepartmentParam) (*model.Department, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RemoveDepartment(ctx context.Context, input model.RemoveDepartmentParam) (*model.Department, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Department(ctx context.Context, query model.DepartmentQuery) (*model.Department, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DepartmentList(ctx context.Context, query model.DepartmentListQuery) (*model.DepartmentConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
