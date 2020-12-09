package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/modules/vehicle/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) DeleteMuckTruckPreviewNumber(ctx context.Context, where model.MuckTruckPreviewNumberBoolExp) (*model.MuckTruckPreviewNumberMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteMuckTruckPreviewNumberByPk(ctx context.Context, id int64) (*model.MuckTruckPreviewNumber, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertMuckTruckPreviewNumber(ctx context.Context, objects []*model.MuckTruckPreviewNumberInsertInput, onConflict *model.MuckTruckPreviewNumberOnConflict) (*model.MuckTruckPreviewNumberMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertMuckTruckPreviewNumberOne(ctx context.Context, object model.MuckTruckPreviewNumberInsertInput, onConflict *model.MuckTruckPreviewNumberOnConflict) (*model.MuckTruckPreviewNumber, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateMuckTruckPreviewNumber(ctx context.Context, inc *model.MuckTruckPreviewNumberIncInput, set *model.MuckTruckPreviewNumberSetInput, where model.MuckTruckPreviewNumberBoolExp) (*model.MuckTruckPreviewNumberMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateMuckTruckPreviewNumberByPk(ctx context.Context, inc *model.MuckTruckPreviewNumberIncInput, set *model.MuckTruckPreviewNumberSetInput, pkColumns model.MuckTruckPreviewNumberPkColumnsInput) (*model.MuckTruckPreviewNumber, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) MuckTruckPreviewNumber(ctx context.Context, distinctOn []model.MuckTruckPreviewNumberSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckPreviewNumberOrderBy, where *model.MuckTruckPreviewNumberBoolExp) ([]*model.MuckTruckPreviewNumber, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) MuckTruckPreviewNumberAggregate(ctx context.Context, distinctOn []model.MuckTruckPreviewNumberSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckPreviewNumberOrderBy, where *model.MuckTruckPreviewNumberBoolExp) (*model.MuckTruckPreviewNumberAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) MuckTruckPreviewNumberByPk(ctx context.Context, id int64) (*model.MuckTruckPreviewNumber, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) MuckTruckPreviewNumber(ctx context.Context, distinctOn []model.MuckTruckPreviewNumberSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckPreviewNumberOrderBy, where *model.MuckTruckPreviewNumberBoolExp) (<-chan []*model.MuckTruckPreviewNumber, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) MuckTruckPreviewNumberAggregate(ctx context.Context, distinctOn []model.MuckTruckPreviewNumberSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckPreviewNumberOrderBy, where *model.MuckTruckPreviewNumberBoolExp) (<-chan *model.MuckTruckPreviewNumberAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) MuckTruckPreviewNumberByPk(ctx context.Context, id int64) (<-chan *model.MuckTruckPreviewNumber, error) {
	panic(fmt.Errorf("not implemented"))
}
