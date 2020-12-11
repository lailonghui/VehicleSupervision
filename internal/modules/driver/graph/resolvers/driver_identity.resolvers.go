package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/modules/driver/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) DeleteDriverIdentity(ctx context.Context, where model.DriverIdentityBoolExp) (*model.DriverIdentityMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteDriverIdentityByPk(ctx context.Context, id int64, identityID string) (*model.DriverIdentity, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertDriverIdentity(ctx context.Context, objects []*model.DriverIdentityInsertInput, onConflict *model.DriverIdentityOnConflict) (*model.DriverIdentityMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertDriverIdentityOne(ctx context.Context, object model.DriverIdentityInsertInput, onConflict *model.DriverIdentityOnConflict) (*model.DriverIdentity, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateDriverIdentity(ctx context.Context, inc *model.DriverIdentityIncInput, set *model.DriverIdentitySetInput, where model.DriverIdentityBoolExp) (*model.DriverIdentityMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateDriverIdentityByPk(ctx context.Context, inc *model.DriverIdentityIncInput, set *model.DriverIdentitySetInput, pkColumns model.DriverIdentityPkColumnsInput) (*model.DriverIdentity, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DriverIdentity(ctx context.Context, distinctOn []model.DriverIdentitySelectColumn, limit *int, offset *int, orderBy []*model.DriverIdentityOrderBy, where *model.DriverIdentityBoolExp) ([]*model.DriverIdentity, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DriverIdentityAggregate(ctx context.Context, distinctOn []model.DriverIdentitySelectColumn, limit *int, offset *int, orderBy []*model.DriverIdentityOrderBy, where *model.DriverIdentityBoolExp) (*model.DriverIdentityAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DriverIdentityByPk(ctx context.Context, id int64, identityID string) (*model.DriverIdentity, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DriverIdentity(ctx context.Context, distinctOn []model.DriverIdentitySelectColumn, limit *int, offset *int, orderBy []*model.DriverIdentityOrderBy, where *model.DriverIdentityBoolExp) (<-chan []*model.DriverIdentity, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DriverIdentityAggregate(ctx context.Context, distinctOn []model.DriverIdentitySelectColumn, limit *int, offset *int, orderBy []*model.DriverIdentityOrderBy, where *model.DriverIdentityBoolExp) (<-chan *model.DriverIdentityAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DriverIdentityByPk(ctx context.Context, id int64, identityID string) (<-chan *model.DriverIdentity, error) {
	panic(fmt.Errorf("not implemented"))
}
