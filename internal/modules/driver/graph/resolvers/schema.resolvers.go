package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/modules/driver/graph/generated"
	"VehicleSupervision/internal/modules/driver/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) CreateDriver(ctx context.Context) (*model.Driver, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
