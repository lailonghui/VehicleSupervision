package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	model1 "VehicleSupervision/internal/modules/admin/systemuser/model"
	"VehicleSupervision/internal/modules/admin/systemuser/query/graph/generated"
	"VehicleSupervision/internal/modules/admin/systemuser/query/graph/model"
	"context"
	"fmt"
)

func (r *queryResolver) SystemUser(ctx context.Context, distinctOn []model.SystemUserSelectColumn, limit *int, offset *int, orderBy []*model.SystemUserOrderBy, where *model.SystemUserBoolExp) ([]*model1.SystemUser, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) SystemUserAggregate(ctx context.Context, distinctOn []model.SystemUserSelectColumn, limit *int, offset *int, orderBy []*model.SystemUserOrderBy, where *model.SystemUserBoolExp) (*model.SystemUserAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) SystemUserByPk(ctx context.Context, id int64) (*model1.SystemUser, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
