package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/modules/area/mutation/graph/generated"
	"VehicleSupervision/internal/modules/area/mutation/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) DeleteCity(ctx context.Context, where model.CityBoolExp) (*model.CityMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteCityByPk(ctx context.Context, id int64) (*model.City, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteDistrict(ctx context.Context, where model.DistrictBoolExp) (*model.DistrictMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteDistrictByPk(ctx context.Context, id int64) (*model.District, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteProvince(ctx context.Context, where model.ProvinceBoolExp) (*model.ProvinceMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteProvinceByPk(ctx context.Context, id int64) (*model.Province, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertCity(ctx context.Context, objects []*model.CityInsertInput, onConflict *model.CityOnConflict) (*model.CityMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertCityOne(ctx context.Context, object model.CityInsertInput, onConflict *model.CityOnConflict) (*model.City, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertDistrict(ctx context.Context, objects []*model.DistrictInsertInput, onConflict *model.DistrictOnConflict) (*model.DistrictMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertDistrictOne(ctx context.Context, object model.DistrictInsertInput, onConflict *model.DistrictOnConflict) (*model.District, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertProvince(ctx context.Context, objects []*model.ProvinceInsertInput, onConflict *model.ProvinceOnConflict) (*model.ProvinceMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertProvinceOne(ctx context.Context, object model.ProvinceInsertInput, onConflict *model.ProvinceOnConflict) (*model.Province, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateCity(ctx context.Context, inc *model.CityIncInput, set *model.CitySetInput, where model.CityBoolExp) (*model.CityMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateCityByPk(ctx context.Context, inc *model.CityIncInput, set *model.CitySetInput, pkColumns model.CityPkColumnsInput) (*model.City, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateDistrict(ctx context.Context, inc *model.DistrictIncInput, set *model.DistrictSetInput, where model.DistrictBoolExp) (*model.DistrictMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateDistrictByPk(ctx context.Context, inc *model.DistrictIncInput, set *model.DistrictSetInput, pkColumns model.DistrictPkColumnsInput) (*model.District, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateProvince(ctx context.Context, inc *model.ProvinceIncInput, set *model.ProvinceSetInput, where model.ProvinceBoolExp) (*model.ProvinceMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateProvinceByPk(ctx context.Context, inc *model.ProvinceIncInput, set *model.ProvinceSetInput, pkColumns model.ProvincePkColumnsInput) (*model.Province, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Bug(ctx context.Context) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
