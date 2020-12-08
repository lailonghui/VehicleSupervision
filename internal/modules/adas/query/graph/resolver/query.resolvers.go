package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/modules/adas/query/graph/generated"
	"VehicleSupervision/internal/modules/adas/query/graph/model"
	"context"
	"fmt"
)

func (r *queryResolver) AdasAttachment(ctx context.Context, distinctOn []model.AdasAttachmentSelectColumn, limit *int, offset *int, orderBy []*model.AdasAttachmentOrderBy, where *model.AdasAttachmentBoolExp) ([]*model.AdasAttachment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AdasAttachmentAggregate(ctx context.Context, distinctOn []model.AdasAttachmentSelectColumn, limit *int, offset *int, orderBy []*model.AdasAttachmentOrderBy, where *model.AdasAttachmentBoolExp) (*model.AdasAttachmentAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AdasAttachmentByPk(ctx context.Context, id int64) (*model.AdasAttachment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AdasData(ctx context.Context, distinctOn []model.AdasDataSelectColumn, limit *int, offset *int, orderBy []*model.AdasDataOrderBy, where *model.AdasDataBoolExp) ([]*model.AdasData, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AdasDataAggregate(ctx context.Context, distinctOn []model.AdasDataSelectColumn, limit *int, offset *int, orderBy []*model.AdasDataOrderBy, where *model.AdasDataBoolExp) (*model.AdasDataAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AdasDataByPk(ctx context.Context, id int64) (*model.AdasData, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
