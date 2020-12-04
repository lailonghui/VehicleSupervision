package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	model1 "VehicleSupervision/internal/modules/admin/enterprise/model"
	"VehicleSupervision/internal/modules/admin/enterprise/mutation/graph/generated"
	"VehicleSupervision/internal/modules/admin/enterprise/mutation/graph/model"
	"context"
	"fmt"
)

func (r *queryResolver) Enterprise(ctx context.Context, distinctOn []model.EnterpriseSelectColumn, limit *int, offset *int, orderBy []*model.EnterpriseOrderBy, where *model.EnterpriseBoolExp) ([]*model1.Enterprise, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) EnterpriseAggregate(ctx context.Context, distinctOn []model.EnterpriseSelectColumn, limit *int, offset *int, orderBy []*model.EnterpriseOrderBy, where *model.EnterpriseBoolExp) (*model.EnterpriseAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) EnterpriseByPk(ctx context.Context, id int64) (*model1.Enterprise, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
