package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/modules/vehicle/graph/generated"
	"VehicleSupervision/internal/modules/vehicle/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) DeleteJjVehicle(ctx context.Context, where model.JjVehicleBoolExp) (*model.JjVehicleMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteJjVehicleByPk(ctx context.Context, id int64) (*model.JjVehicle, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteMuckTruckCategoryInfo(ctx context.Context, where model.MuckTruckCategoryInfoBoolExp) (*model.MuckTruckCategoryInfoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteMuckTruckCategoryInfoByPk(ctx context.Context, id int64) (*model.MuckTruckCategoryInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteMuckTruckInfo(ctx context.Context, where model.MuckTruckInfoBoolExp) (*model.MuckTruckInfoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteMuckTruckInfoByPk(ctx context.Context, muckTruckID int64) (*model.MuckTruckInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteMuckTruckPreviewNumber(ctx context.Context, where model.MuckTruckPreviewNumberBoolExp) (*model.MuckTruckPreviewNumberMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteMuckTruckPreviewNumberByPk(ctx context.Context, id int64) (*model.MuckTruckPreviewNumber, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteMuckTruckWorkerIDCardOrders(ctx context.Context, where model.MuckTruckWorkerIDCardOrdersBoolExp) (*model.MuckTruckWorkerIDCardOrdersMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteMuckTruckWorkerIDCardOrdersByPk(ctx context.Context, id int64) (*model.MuckTruckWorkerIDCardOrders, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteOperatingVehicleInfo(ctx context.Context, where model.OperatingVehicleInfoBoolExp) (*model.OperatingVehicleInfoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteOperatingVehicleInfoByPk(ctx context.Context, operatingVehicleID int64) (*model.OperatingVehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteOwnerInfo(ctx context.Context, where model.OwnerInfoBoolExp) (*model.OwnerInfoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteOwnerInfoByPk(ctx context.Context, id int64) (*model.OwnerInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteVehicleInfo(ctx context.Context, where model.VehicleInfoBoolExp) (*model.VehicleInfoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteVehicleInfoByPk(ctx context.Context, id int64, vehicleID string) (*model.VehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteVehicleInfoChangeLog(ctx context.Context, where model.VehicleInfoChangeLogBoolExp) (*model.VehicleInfoChangeLogMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteVehicleInfoChangeLogByPk(ctx context.Context, id int64, vehicleInfoChangeID string) (*model.VehicleInfoChangeLog, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteVehicleSupervisionPhoto(ctx context.Context, where model.VehicleSupervisionPhotoBoolExp) (*model.VehicleSupervisionPhotoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteVehicleSupervisionPhotoByPk(ctx context.Context, id int64, supervisionPhotoID string) (*model.VehicleSupervisionPhoto, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertJjVehicle(ctx context.Context, objects []*model.JjVehicleInsertInput, onConflict *model.JjVehicleOnConflict) (*model.JjVehicleMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertJjVehicleOne(ctx context.Context, object model.JjVehicleInsertInput, onConflict *model.JjVehicleOnConflict) (*model.JjVehicle, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertMuckTruckCategoryInfo(ctx context.Context, objects []*model.MuckTruckCategoryInfoInsertInput, onConflict *model.MuckTruckCategoryInfoOnConflict) (*model.MuckTruckCategoryInfoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertMuckTruckCategoryInfoOne(ctx context.Context, object model.MuckTruckCategoryInfoInsertInput, onConflict *model.MuckTruckCategoryInfoOnConflict) (*model.MuckTruckCategoryInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertMuckTruckInfo(ctx context.Context, objects []*model.MuckTruckInfoInsertInput, onConflict *model.MuckTruckInfoOnConflict) (*model.MuckTruckInfoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertMuckTruckInfoOne(ctx context.Context, object model.MuckTruckInfoInsertInput, onConflict *model.MuckTruckInfoOnConflict) (*model.MuckTruckInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertMuckTruckPreviewNumber(ctx context.Context, objects []*model.MuckTruckPreviewNumberInsertInput, onConflict *model.MuckTruckPreviewNumberOnConflict) (*model.MuckTruckPreviewNumberMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertMuckTruckPreviewNumberOne(ctx context.Context, object model.MuckTruckPreviewNumberInsertInput, onConflict *model.MuckTruckPreviewNumberOnConflict) (*model.MuckTruckPreviewNumber, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertMuckTruckWorkerIDCardOrders(ctx context.Context, objects []*model.MuckTruckWorkerIDCardOrdersInsertInput, onConflict *model.MuckTruckWorkerIDCardOrdersOnConflict) (*model.MuckTruckWorkerIDCardOrdersMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertMuckTruckWorkerIDCardOrdersOne(ctx context.Context, object model.MuckTruckWorkerIDCardOrdersInsertInput, onConflict *model.MuckTruckWorkerIDCardOrdersOnConflict) (*model.MuckTruckWorkerIDCardOrders, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertOperatingVehicleInfo(ctx context.Context, objects []*model.OperatingVehicleInfoInsertInput, onConflict *model.OperatingVehicleInfoOnConflict) (*model.OperatingVehicleInfoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertOperatingVehicleInfoOne(ctx context.Context, object model.OperatingVehicleInfoInsertInput, onConflict *model.OperatingVehicleInfoOnConflict) (*model.OperatingVehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertOwnerInfo(ctx context.Context, objects []*model.OwnerInfoInsertInput, onConflict *model.OwnerInfoOnConflict) (*model.OwnerInfoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertOwnerInfoOne(ctx context.Context, object model.OwnerInfoInsertInput, onConflict *model.OwnerInfoOnConflict) (*model.OwnerInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertVehicleInfo(ctx context.Context, objects []*model.VehicleInfoInsertInput, onConflict *model.VehicleInfoOnConflict) (*model.VehicleInfoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertVehicleInfoChangeLog(ctx context.Context, objects []*model.VehicleInfoChangeLogInsertInput, onConflict *model.VehicleInfoChangeLogOnConflict) (*model.VehicleInfoChangeLogMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertVehicleInfoChangeLogOne(ctx context.Context, object model.VehicleInfoChangeLogInsertInput, onConflict *model.VehicleInfoChangeLogOnConflict) (*model.VehicleInfoChangeLog, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertVehicleInfoOne(ctx context.Context, object model.VehicleInfoInsertInput, onConflict *model.VehicleInfoOnConflict) (*model.VehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertVehicleSupervisionPhoto(ctx context.Context, objects []*model.VehicleSupervisionPhotoInsertInput, onConflict *model.VehicleSupervisionPhotoOnConflict) (*model.VehicleSupervisionPhotoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertVehicleSupervisionPhotoOne(ctx context.Context, object model.VehicleSupervisionPhotoInsertInput, onConflict *model.VehicleSupervisionPhotoOnConflict) (*model.VehicleSupervisionPhoto, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateJjVehicle(ctx context.Context, inc *model.JjVehicleIncInput, set *model.JjVehicleSetInput, where model.JjVehicleBoolExp) (*model.JjVehicleMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateJjVehicleByPk(ctx context.Context, inc *model.JjVehicleIncInput, set *model.JjVehicleSetInput, pkColumns model.JjVehiclePkColumnsInput) (*model.JjVehicle, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateMuckTruckCategoryInfo(ctx context.Context, inc *model.MuckTruckCategoryInfoIncInput, set *model.MuckTruckCategoryInfoSetInput, where model.MuckTruckCategoryInfoBoolExp) (*model.MuckTruckCategoryInfoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateMuckTruckCategoryInfoByPk(ctx context.Context, inc *model.MuckTruckCategoryInfoIncInput, set *model.MuckTruckCategoryInfoSetInput, pkColumns model.MuckTruckCategoryInfoPkColumnsInput) (*model.MuckTruckCategoryInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateMuckTruckInfo(ctx context.Context, inc *model.MuckTruckInfoIncInput, set *model.MuckTruckInfoSetInput, where model.MuckTruckInfoBoolExp) (*model.MuckTruckInfoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateMuckTruckInfoByPk(ctx context.Context, inc *model.MuckTruckInfoIncInput, set *model.MuckTruckInfoSetInput, pkColumns model.MuckTruckInfoPkColumnsInput) (*model.MuckTruckInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateMuckTruckPreviewNumber(ctx context.Context, inc *model.MuckTruckPreviewNumberIncInput, set *model.MuckTruckPreviewNumberSetInput, where model.MuckTruckPreviewNumberBoolExp) (*model.MuckTruckPreviewNumberMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateMuckTruckPreviewNumberByPk(ctx context.Context, inc *model.MuckTruckPreviewNumberIncInput, set *model.MuckTruckPreviewNumberSetInput, pkColumns model.MuckTruckPreviewNumberPkColumnsInput) (*model.MuckTruckPreviewNumber, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateMuckTruckWorkerIDCardOrders(ctx context.Context, inc *model.MuckTruckWorkerIDCardOrdersIncInput, set *model.MuckTruckWorkerIDCardOrdersSetInput, where model.MuckTruckWorkerIDCardOrdersBoolExp) (*model.MuckTruckWorkerIDCardOrdersMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateMuckTruckWorkerIDCardOrdersByPk(ctx context.Context, inc *model.MuckTruckWorkerIDCardOrdersIncInput, set *model.MuckTruckWorkerIDCardOrdersSetInput, pkColumns model.MuckTruckWorkerIDCardOrdersPkColumnsInput) (*model.MuckTruckWorkerIDCardOrders, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateOperatingVehicleInfo(ctx context.Context, inc *model.OperatingVehicleInfoIncInput, set *model.OperatingVehicleInfoSetInput, where model.OperatingVehicleInfoBoolExp) (*model.OperatingVehicleInfoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateOperatingVehicleInfoByPk(ctx context.Context, inc *model.OperatingVehicleInfoIncInput, set *model.OperatingVehicleInfoSetInput, pkColumns model.OperatingVehicleInfoPkColumnsInput) (*model.OperatingVehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateOwnerInfo(ctx context.Context, inc *model.OwnerInfoIncInput, set *model.OwnerInfoSetInput, where model.OwnerInfoBoolExp) (*model.OwnerInfoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateOwnerInfoByPk(ctx context.Context, inc *model.OwnerInfoIncInput, set *model.OwnerInfoSetInput, pkColumns model.OwnerInfoPkColumnsInput) (*model.OwnerInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateVehicleInfo(ctx context.Context, inc *model.VehicleInfoIncInput, set *model.VehicleInfoSetInput, where model.VehicleInfoBoolExp) (*model.VehicleInfoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateVehicleInfoByPk(ctx context.Context, inc *model.VehicleInfoIncInput, set *model.VehicleInfoSetInput, pkColumns model.VehicleInfoPkColumnsInput) (*model.VehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateVehicleInfoChangeLog(ctx context.Context, inc *model.VehicleInfoChangeLogIncInput, set *model.VehicleInfoChangeLogSetInput, where model.VehicleInfoChangeLogBoolExp) (*model.VehicleInfoChangeLogMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateVehicleInfoChangeLogByPk(ctx context.Context, inc *model.VehicleInfoChangeLogIncInput, set *model.VehicleInfoChangeLogSetInput, pkColumns model.VehicleInfoChangeLogPkColumnsInput) (*model.VehicleInfoChangeLog, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateVehicleSupervisionPhoto(ctx context.Context, inc *model.VehicleSupervisionPhotoIncInput, set *model.VehicleSupervisionPhotoSetInput, where model.VehicleSupervisionPhotoBoolExp) (*model.VehicleSupervisionPhotoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateVehicleSupervisionPhotoByPk(ctx context.Context, inc *model.VehicleSupervisionPhotoIncInput, set *model.VehicleSupervisionPhotoSetInput, pkColumns model.VehicleSupervisionPhotoPkColumnsInput) (*model.VehicleSupervisionPhoto, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) JjVehicle(ctx context.Context, distinctOn []model.JjVehicleSelectColumn, limit *int, offset *int, orderBy []*model.JjVehicleOrderBy, where *model.JjVehicleBoolExp) ([]*model.JjVehicle, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) JjVehicleAggregate(ctx context.Context, distinctOn []model.JjVehicleSelectColumn, limit *int, offset *int, orderBy []*model.JjVehicleOrderBy, where *model.JjVehicleBoolExp) (*model.JjVehicleAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) JjVehicleByPk(ctx context.Context, id int64) (*model.JjVehicle, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) MuckTruckCategoryInfo(ctx context.Context, distinctOn []model.MuckTruckCategoryInfoSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckCategoryInfoOrderBy, where *model.MuckTruckCategoryInfoBoolExp) ([]*model.MuckTruckCategoryInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) MuckTruckCategoryInfoAggregate(ctx context.Context, distinctOn []model.MuckTruckCategoryInfoSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckCategoryInfoOrderBy, where *model.MuckTruckCategoryInfoBoolExp) (*model.MuckTruckCategoryInfoAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) MuckTruckCategoryInfoByPk(ctx context.Context, id int64) (*model.MuckTruckCategoryInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) MuckTruckInfo(ctx context.Context, distinctOn []model.MuckTruckInfoSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckInfoOrderBy, where *model.MuckTruckInfoBoolExp) ([]*model.MuckTruckInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) MuckTruckInfoAggregate(ctx context.Context, distinctOn []model.MuckTruckInfoSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckInfoOrderBy, where *model.MuckTruckInfoBoolExp) (*model.MuckTruckInfoAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) MuckTruckInfoByPk(ctx context.Context, muckTruckID int64) (*model.MuckTruckInfo, error) {
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

func (r *queryResolver) MuckTruckWorkerIDCardOrders(ctx context.Context, distinctOn []model.MuckTruckWorkerIDCardOrdersSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckWorkerIDCardOrdersOrderBy, where *model.MuckTruckWorkerIDCardOrdersBoolExp) ([]*model.MuckTruckWorkerIDCardOrders, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) MuckTruckWorkerIDCardOrdersAggregate(ctx context.Context, distinctOn []model.MuckTruckWorkerIDCardOrdersSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckWorkerIDCardOrdersOrderBy, where *model.MuckTruckWorkerIDCardOrdersBoolExp) (*model.MuckTruckWorkerIDCardOrdersAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) MuckTruckWorkerIDCardOrdersByPk(ctx context.Context, id int64) (*model.MuckTruckWorkerIDCardOrders, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) OperatingVehicleInfo(ctx context.Context, distinctOn []model.OperatingVehicleInfoSelectColumn, limit *int, offset *int, orderBy []*model.OperatingVehicleInfoOrderBy, where *model.OperatingVehicleInfoBoolExp) ([]*model.OperatingVehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) OperatingVehicleInfoAggregate(ctx context.Context, distinctOn []model.OperatingVehicleInfoSelectColumn, limit *int, offset *int, orderBy []*model.OperatingVehicleInfoOrderBy, where *model.OperatingVehicleInfoBoolExp) (*model.OperatingVehicleInfoAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) OperatingVehicleInfoByPk(ctx context.Context, operatingVehicleID int64) (*model.OperatingVehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) OwnerInfo(ctx context.Context, distinctOn []model.OwnerInfoSelectColumn, limit *int, offset *int, orderBy []*model.OwnerInfoOrderBy, where *model.OwnerInfoBoolExp) ([]*model.OwnerInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) OwnerInfoAggregate(ctx context.Context, distinctOn []model.OwnerInfoSelectColumn, limit *int, offset *int, orderBy []*model.OwnerInfoOrderBy, where *model.OwnerInfoBoolExp) (*model.OwnerInfoAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) OwnerInfoByPk(ctx context.Context, id int64) (*model.OwnerInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) VehicleInfo(ctx context.Context, distinctOn []model.VehicleInfoSelectColumn, limit *int, offset *int, orderBy []*model.VehicleInfoOrderBy, where *model.VehicleInfoBoolExp) ([]*model.VehicleInfo, error) {
	var v []*model.VehicleInfo
	v = append(v, &model.VehicleInfo{
		ID:        12,
		VehicleID: "12",
	})
	return v, nil
}

func (r *queryResolver) VehicleInfoAggregate(ctx context.Context, distinctOn []model.VehicleInfoSelectColumn, limit *int, offset *int, orderBy []*model.VehicleInfoOrderBy, where *model.VehicleInfoBoolExp) (*model.VehicleInfoAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) VehicleInfoByPk(ctx context.Context, id int64, vehicleID string) (*model.VehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) VehicleInfoChangeLog(ctx context.Context, distinctOn []model.VehicleInfoChangeLogSelectColumn, limit *int, offset *int, orderBy []*model.VehicleInfoChangeLogOrderBy, where *model.VehicleInfoChangeLogBoolExp) ([]*model.VehicleInfoChangeLog, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) VehicleInfoChangeLogAggregate(ctx context.Context, distinctOn []model.VehicleInfoChangeLogSelectColumn, limit *int, offset *int, orderBy []*model.VehicleInfoChangeLogOrderBy, where *model.VehicleInfoChangeLogBoolExp) (*model.VehicleInfoChangeLogAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) VehicleInfoChangeLogByPk(ctx context.Context, id int64, vehicleInfoChangeID string) (*model.VehicleInfoChangeLog, error) {
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

func (r *subscriptionResolver) JjVehicle(ctx context.Context, distinctOn []model.JjVehicleSelectColumn, limit *int, offset *int, orderBy []*model.JjVehicleOrderBy, where *model.JjVehicleBoolExp) (<-chan []*model.JjVehicle, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) JjVehicleAggregate(ctx context.Context, distinctOn []model.JjVehicleSelectColumn, limit *int, offset *int, orderBy []*model.JjVehicleOrderBy, where *model.JjVehicleBoolExp) (<-chan *model.JjVehicleAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) JjVehicleByPk(ctx context.Context, id int64) (<-chan *model.JjVehicle, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) MuckTruckCategoryInfo(ctx context.Context, distinctOn []model.MuckTruckCategoryInfoSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckCategoryInfoOrderBy, where *model.MuckTruckCategoryInfoBoolExp) (<-chan []*model.MuckTruckCategoryInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) MuckTruckCategoryInfoAggregate(ctx context.Context, distinctOn []model.MuckTruckCategoryInfoSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckCategoryInfoOrderBy, where *model.MuckTruckCategoryInfoBoolExp) (<-chan *model.MuckTruckCategoryInfoAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) MuckTruckCategoryInfoByPk(ctx context.Context, id int64) (<-chan *model.MuckTruckCategoryInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) MuckTruckInfo(ctx context.Context, distinctOn []model.MuckTruckInfoSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckInfoOrderBy, where *model.MuckTruckInfoBoolExp) (<-chan []*model.MuckTruckInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) MuckTruckInfoAggregate(ctx context.Context, distinctOn []model.MuckTruckInfoSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckInfoOrderBy, where *model.MuckTruckInfoBoolExp) (<-chan *model.MuckTruckInfoAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) MuckTruckInfoByPk(ctx context.Context, muckTruckID int64) (<-chan *model.MuckTruckInfo, error) {
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

func (r *subscriptionResolver) MuckTruckWorkerIDCardOrders(ctx context.Context, distinctOn []model.MuckTruckWorkerIDCardOrdersSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckWorkerIDCardOrdersOrderBy, where *model.MuckTruckWorkerIDCardOrdersBoolExp) (<-chan []*model.MuckTruckWorkerIDCardOrders, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) MuckTruckWorkerIDCardOrdersAggregate(ctx context.Context, distinctOn []model.MuckTruckWorkerIDCardOrdersSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckWorkerIDCardOrdersOrderBy, where *model.MuckTruckWorkerIDCardOrdersBoolExp) (<-chan *model.MuckTruckWorkerIDCardOrdersAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) MuckTruckWorkerIDCardOrdersByPk(ctx context.Context, id int64) (<-chan *model.MuckTruckWorkerIDCardOrders, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) OperatingVehicleInfo(ctx context.Context, distinctOn []model.OperatingVehicleInfoSelectColumn, limit *int, offset *int, orderBy []*model.OperatingVehicleInfoOrderBy, where *model.OperatingVehicleInfoBoolExp) (<-chan []*model.OperatingVehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) OperatingVehicleInfoAggregate(ctx context.Context, distinctOn []model.OperatingVehicleInfoSelectColumn, limit *int, offset *int, orderBy []*model.OperatingVehicleInfoOrderBy, where *model.OperatingVehicleInfoBoolExp) (<-chan *model.OperatingVehicleInfoAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) OperatingVehicleInfoByPk(ctx context.Context, operatingVehicleID int64) (<-chan *model.OperatingVehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) OwnerInfo(ctx context.Context, distinctOn []model.OwnerInfoSelectColumn, limit *int, offset *int, orderBy []*model.OwnerInfoOrderBy, where *model.OwnerInfoBoolExp) (<-chan []*model.OwnerInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) OwnerInfoAggregate(ctx context.Context, distinctOn []model.OwnerInfoSelectColumn, limit *int, offset *int, orderBy []*model.OwnerInfoOrderBy, where *model.OwnerInfoBoolExp) (<-chan *model.OwnerInfoAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) OwnerInfoByPk(ctx context.Context, id int64) (<-chan *model.OwnerInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VehicleInfo(ctx context.Context, distinctOn []model.VehicleInfoSelectColumn, limit *int, offset *int, orderBy []*model.VehicleInfoOrderBy, where *model.VehicleInfoBoolExp) (<-chan []*model.VehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VehicleInfoAggregate(ctx context.Context, distinctOn []model.VehicleInfoSelectColumn, limit *int, offset *int, orderBy []*model.VehicleInfoOrderBy, where *model.VehicleInfoBoolExp) (<-chan *model.VehicleInfoAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VehicleInfoByPk(ctx context.Context, id int64, vehicleID string) (<-chan *model.VehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VehicleInfoChangeLog(ctx context.Context, distinctOn []model.VehicleInfoChangeLogSelectColumn, limit *int, offset *int, orderBy []*model.VehicleInfoChangeLogOrderBy, where *model.VehicleInfoChangeLogBoolExp) (<-chan []*model.VehicleInfoChangeLog, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VehicleInfoChangeLogAggregate(ctx context.Context, distinctOn []model.VehicleInfoChangeLogSelectColumn, limit *int, offset *int, orderBy []*model.VehicleInfoChangeLogOrderBy, where *model.VehicleInfoChangeLogBoolExp) (<-chan *model.VehicleInfoChangeLogAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VehicleInfoChangeLogByPk(ctx context.Context, id int64, vehicleInfoChangeID string) (<-chan *model.VehicleInfoChangeLog, error) {
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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
