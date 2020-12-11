package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/modules/vehicle_violation/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) DeleteIllegalPhoto(ctx context.Context, where model.IllegalPhotoBoolExp) (*model.IllegalPhotoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteIllegalPhotoByPk(ctx context.Context, id int64, illegalPhotoID string) (*model.IllegalPhoto, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertIllegalPhoto(ctx context.Context, objects []*model.IllegalPhotoInsertInput, onConflict *model.IllegalPhotoOnConflict) (*model.IllegalPhotoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertIllegalPhotoOne(ctx context.Context, object model.IllegalPhotoInsertInput, onConflict *model.IllegalPhotoOnConflict) (*model.IllegalPhoto, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateIllegalPhoto(ctx context.Context, inc *model.IllegalPhotoIncInput, set *model.IllegalPhotoSetInput, where model.IllegalPhotoBoolExp) (*model.IllegalPhotoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateIllegalPhotoByPk(ctx context.Context, inc *model.IllegalPhotoIncInput, set *model.IllegalPhotoSetInput, pkColumns model.IllegalPhotoPkColumnsInput) (*model.IllegalPhoto, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) IllegalPhoto(ctx context.Context, distinctOn []model.IllegalPhotoSelectColumn, limit *int, offset *int, orderBy []*model.IllegalPhotoOrderBy, where *model.IllegalPhotoBoolExp) ([]*model.IllegalPhoto, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) IllegalPhotoAggregate(ctx context.Context, distinctOn []model.IllegalPhotoSelectColumn, limit *int, offset *int, orderBy []*model.IllegalPhotoOrderBy, where *model.IllegalPhotoBoolExp) (*model.IllegalPhotoAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) IllegalPhotoByPk(ctx context.Context, id int64, illegalPhotoID string) (*model.IllegalPhoto, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) IllegalPhoto(ctx context.Context, distinctOn []model.IllegalPhotoSelectColumn, limit *int, offset *int, orderBy []*model.IllegalPhotoOrderBy, where *model.IllegalPhotoBoolExp) (<-chan []*model.IllegalPhoto, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) IllegalPhotoAggregate(ctx context.Context, distinctOn []model.IllegalPhotoSelectColumn, limit *int, offset *int, orderBy []*model.IllegalPhotoOrderBy, where *model.IllegalPhotoBoolExp) (<-chan *model.IllegalPhotoAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) IllegalPhotoByPk(ctx context.Context, id int64, illegalPhotoID string) (<-chan *model.IllegalPhoto, error) {
	panic(fmt.Errorf("not implemented"))
}
