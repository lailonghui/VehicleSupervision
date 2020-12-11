package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/modules/vehicle_alarm/graph/generated"
	"VehicleSupervision/internal/modules/vehicle_alarm/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) DeleteAlarmProcessingRecord(ctx context.Context, where model.AlarmProcessingRecordBoolExp) (*model.AlarmProcessingRecordMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertAlarmProcessingRecord(ctx context.Context, objects []*model.AlarmProcessingRecordInsertInput) (*model.AlarmProcessingRecordMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertAlarmProcessingRecordOne(ctx context.Context, object model.AlarmProcessingRecordInsertInput) (*model.AlarmProcessingRecord, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateAlarmProcessingRecord(ctx context.Context, inc *model.AlarmProcessingRecordIncInput, set *model.AlarmProcessingRecordSetInput, where model.AlarmProcessingRecordBoolExp) (*model.AlarmProcessingRecordMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AlarmProcessingRecord(ctx context.Context, distinctOn []model.AlarmProcessingRecordSelectColumn, limit *int, offset *int, orderBy []*model.AlarmProcessingRecordOrderBy, where *model.AlarmProcessingRecordBoolExp) ([]*model.AlarmProcessingRecord, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AlarmProcessingRecordAggregate(ctx context.Context, distinctOn []model.AlarmProcessingRecordSelectColumn, limit *int, offset *int, orderBy []*model.AlarmProcessingRecordOrderBy, where *model.AlarmProcessingRecordBoolExp) (*model.AlarmProcessingRecordAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) AlarmProcessingRecord(ctx context.Context, distinctOn []model.AlarmProcessingRecordSelectColumn, limit *int, offset *int, orderBy []*model.AlarmProcessingRecordOrderBy, where *model.AlarmProcessingRecordBoolExp) (<-chan []*model.AlarmProcessingRecord, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) AlarmProcessingRecordAggregate(ctx context.Context, distinctOn []model.AlarmProcessingRecordSelectColumn, limit *int, offset *int, orderBy []*model.AlarmProcessingRecordOrderBy, where *model.AlarmProcessingRecordBoolExp) (<-chan *model.AlarmProcessingRecordAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
