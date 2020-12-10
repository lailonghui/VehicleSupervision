package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/modules/vehicle/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) DeleteVehicleSupervisionPhoto(ctx context.Context, where model.VehicleSupervisionPhotoBoolExp) (*model.VehicleSupervisionPhotoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteVehicleSupervisionPhotoByPk(ctx context.Context, id int64, supervisionPhotoID string) (*model.VehicleSupervisionPhoto, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertVehicleSupervisionPhoto(ctx context.Context, objects []*model.VehicleSupervisionPhotoInsertInput, onConflict *model.VehicleSupervisionPhotoOnConflict) (*model.VehicleSupervisionPhotoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertVehicleSupervisionPhotoOne(ctx context.Context, object model.VehicleSupervisionPhotoInsertInput, onConflict *model.VehicleSupervisionPhotoOnConflict) (*model.VehicleSupervisionPhoto, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateVehicleSupervisionPhoto(ctx context.Context, inc *model.VehicleSupervisionPhotoIncInput, set *model.VehicleSupervisionPhotoSetInput, where model.VehicleSupervisionPhotoBoolExp) (*model.VehicleSupervisionPhotoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateVehicleSupervisionPhotoByPk(ctx context.Context, inc *model.VehicleSupervisionPhotoIncInput, set *model.VehicleSupervisionPhotoSetInput, pkColumns model.VehicleSupervisionPhotoPkColumnsInput) (*model.VehicleSupervisionPhoto, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) VehicleSupervisionPhoto(ctx context.Context, distinctOn []model.VehicleSupervisionPhotoSelectColumn, limit *int, offset *int, orderBy []*model.VehicleSupervisionPhotoOrderBy, where *model.VehicleSupervisionPhotoBoolExp) ([]*model.VehicleSupervisionPhoto, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) VehicleSupervisionPhotoAggregate(ctx context.Context, distinctOn []model.VehicleSupervisionPhotoSelectColumn, limit *int, offset *int, orderBy []*model.VehicleSupervisionPhotoOrderBy, where *model.VehicleSupervisionPhotoBoolExp) (*model.VehicleSupervisionPhotoAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) VehicleSupervisionPhotoByPk(ctx context.Context, id int64, supervisionPhotoID string) (*model.VehicleSupervisionPhoto, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VehicleSupervisionPhoto(ctx context.Context, distinctOn []model.VehicleSupervisionPhotoSelectColumn, limit *int, offset *int, orderBy []*model.VehicleSupervisionPhotoOrderBy, where *model.VehicleSupervisionPhotoBoolExp) (<-chan []*model.VehicleSupervisionPhoto, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VehicleSupervisionPhotoAggregate(ctx context.Context, distinctOn []model.VehicleSupervisionPhotoSelectColumn, limit *int, offset *int, orderBy []*model.VehicleSupervisionPhotoOrderBy, where *model.VehicleSupervisionPhotoBoolExp) (<-chan *model.VehicleSupervisionPhotoAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VehicleSupervisionPhotoByPk(ctx context.Context, id int64, supervisionPhotoID string) (<-chan *model.VehicleSupervisionPhoto, error) {
	panic(fmt.Errorf("not implemented"))
}
