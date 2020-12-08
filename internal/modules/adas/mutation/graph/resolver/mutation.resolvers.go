package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/modules/adas/mutation/graph/generated"
	"VehicleSupervision/internal/modules/adas/mutation/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) DeleteAdasAttachment(ctx context.Context, where model.AdasAttachmentBoolExp) (*model.AdasAttachmentMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteAdasAttachmentByPk(ctx context.Context, id int64) (*model.AdasAttachment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteAdasData(ctx context.Context, where model.AdasDataBoolExp) (*model.AdasDataMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteAdasDataByPk(ctx context.Context, id int64) (*model.AdasData, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Bug(ctx context.Context) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
