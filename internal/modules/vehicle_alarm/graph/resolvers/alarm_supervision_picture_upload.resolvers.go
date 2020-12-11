package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/modules/vehicle_alarm/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) DeleteAlarmSupervisionPictureUpload(ctx context.Context, where model.AlarmSupervisionPictureUploadBoolExp) (*model.AlarmSupervisionPictureUploadMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteAlarmSupervisionPictureUploadByPk(ctx context.Context, id int64) (*model.AlarmSupervisionPictureUpload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertAlarmSupervisionPictureUpload(ctx context.Context, objects []*model.AlarmSupervisionPictureUploadInsertInput, onConflict *model.AlarmSupervisionPictureUploadOnConflict) (*model.AlarmSupervisionPictureUploadMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertAlarmSupervisionPictureUploadOne(ctx context.Context, object model.AlarmSupervisionPictureUploadInsertInput, onConflict *model.AlarmSupervisionPictureUploadOnConflict) (*model.AlarmSupervisionPictureUpload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateAlarmSupervisionPictureUpload(ctx context.Context, inc *model.AlarmSupervisionPictureUploadIncInput, set *model.AlarmSupervisionPictureUploadSetInput, where model.AlarmSupervisionPictureUploadBoolExp) (*model.AlarmSupervisionPictureUploadMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateAlarmSupervisionPictureUploadByPk(ctx context.Context, inc *model.AlarmSupervisionPictureUploadIncInput, set *model.AlarmSupervisionPictureUploadSetInput, pkColumns model.AlarmSupervisionPictureUploadPkColumnsInput) (*model.AlarmSupervisionPictureUpload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AlarmSupervisionPictureUpload(ctx context.Context, distinctOn []model.AlarmSupervisionPictureUploadSelectColumn, limit *int, offset *int, orderBy []*model.AlarmSupervisionPictureUploadOrderBy, where *model.AlarmSupervisionPictureUploadBoolExp) ([]*model.AlarmSupervisionPictureUpload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AlarmSupervisionPictureUploadAggregate(ctx context.Context, distinctOn []model.AlarmSupervisionPictureUploadSelectColumn, limit *int, offset *int, orderBy []*model.AlarmSupervisionPictureUploadOrderBy, where *model.AlarmSupervisionPictureUploadBoolExp) (*model.AlarmSupervisionPictureUploadAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AlarmSupervisionPictureUploadByPk(ctx context.Context, id int64) (*model.AlarmSupervisionPictureUpload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) AlarmSupervisionPictureUpload(ctx context.Context, distinctOn []model.AlarmSupervisionPictureUploadSelectColumn, limit *int, offset *int, orderBy []*model.AlarmSupervisionPictureUploadOrderBy, where *model.AlarmSupervisionPictureUploadBoolExp) (<-chan []*model.AlarmSupervisionPictureUpload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) AlarmSupervisionPictureUploadAggregate(ctx context.Context, distinctOn []model.AlarmSupervisionPictureUploadSelectColumn, limit *int, offset *int, orderBy []*model.AlarmSupervisionPictureUploadOrderBy, where *model.AlarmSupervisionPictureUploadBoolExp) (<-chan *model.AlarmSupervisionPictureUploadAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) AlarmSupervisionPictureUploadByPk(ctx context.Context, id int64) (<-chan *model.AlarmSupervisionPictureUpload, error) {
	panic(fmt.Errorf("not implemented"))
}
