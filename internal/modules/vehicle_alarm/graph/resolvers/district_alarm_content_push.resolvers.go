package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/modules/vehicle_alarm/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) DeleteDistrictAlarmContentPush(ctx context.Context, where model.DistrictAlarmContentPushBoolExp) (*model.DistrictAlarmContentPushMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteDistrictAlarmContentPushByPk(ctx context.Context, id int64) (*model.DistrictAlarmContentPush, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertDistrictAlarmContentPush(ctx context.Context, objects []*model.DistrictAlarmContentPushInsertInput, onConflict *model.DistrictAlarmContentPushOnConflict) (*model.DistrictAlarmContentPushMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertDistrictAlarmContentPushOne(ctx context.Context, object model.DistrictAlarmContentPushInsertInput, onConflict *model.DistrictAlarmContentPushOnConflict) (*model.DistrictAlarmContentPush, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateDistrictAlarmContentPush(ctx context.Context, inc *model.DistrictAlarmContentPushIncInput, set *model.DistrictAlarmContentPushSetInput, where model.DistrictAlarmContentPushBoolExp) (*model.DistrictAlarmContentPushMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateDistrictAlarmContentPushByPk(ctx context.Context, inc *model.DistrictAlarmContentPushIncInput, set *model.DistrictAlarmContentPushSetInput, pkColumns model.DistrictAlarmContentPushPkColumnsInput) (*model.DistrictAlarmContentPush, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DistrictAlarmContentPush(ctx context.Context, distinctOn []model.DistrictAlarmContentPushSelectColumn, limit *int, offset *int, orderBy []*model.DistrictAlarmContentPushOrderBy, where *model.DistrictAlarmContentPushBoolExp) ([]*model.DistrictAlarmContentPush, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DistrictAlarmContentPushAggregate(ctx context.Context, distinctOn []model.DistrictAlarmContentPushSelectColumn, limit *int, offset *int, orderBy []*model.DistrictAlarmContentPushOrderBy, where *model.DistrictAlarmContentPushBoolExp) (*model.DistrictAlarmContentPushAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DistrictAlarmContentPushByPk(ctx context.Context, id int64) (*model.DistrictAlarmContentPush, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DistrictAlarmContentPush(ctx context.Context, distinctOn []model.DistrictAlarmContentPushSelectColumn, limit *int, offset *int, orderBy []*model.DistrictAlarmContentPushOrderBy, where *model.DistrictAlarmContentPushBoolExp) (<-chan []*model.DistrictAlarmContentPush, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DistrictAlarmContentPushAggregate(ctx context.Context, distinctOn []model.DistrictAlarmContentPushSelectColumn, limit *int, offset *int, orderBy []*model.DistrictAlarmContentPushOrderBy, where *model.DistrictAlarmContentPushBoolExp) (<-chan *model.DistrictAlarmContentPushAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DistrictAlarmContentPushByPk(ctx context.Context, id int64) (<-chan *model.DistrictAlarmContentPush, error) {
	panic(fmt.Errorf("not implemented"))
}
