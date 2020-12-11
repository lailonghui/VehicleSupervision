package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/vehicle/graph/model"
	util2 "VehicleSupervision/pkg/graphql/util"
	"VehicleSupervision/pkg/util"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteVehicleSupervisionPhoto(ctx context.Context, where model.VehicleSupervisionPhotoBoolExp) (*model.VehicleSupervisionPhotoMutationResponse, error) {
	var err error
	var rs []*model.VehicleSupervisionPhoto
	qt := util2.NewQueryTranslator(db.DB, &model.VehicleSupervisionPhoto{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util2.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	if len(preloads) > 0 {
		// 如果请求的字段不为空，则先查询一遍数据库
		tx := tx.Select(preloads)
		tx = tx.Find(&rs)
		// 如果查询结果含有错误，则返回错误
		if err := tx.Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, nil
			}
		}
	}
	// 删除
	tx = tx.Delete(nil)
	if err = tx.Error; err != nil {
		return nil, err
	}
	return &model.VehicleSupervisionPhotoMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, err
}

func (r *mutationResolver) DeleteVehicleSupervisionPhotoByPk(ctx context.Context, id int64, supervisionPhotoID string) (*model.VehicleSupervisionPhoto, error) {
	var rs model.VehicleSupervisionPhoto
	var err error
	tx := db.DB.Model(&model.VehicleSupervisionPhoto{})
	preloads := util2.GetPreloads(ctx)
	if len(preloads) > 0 {
		// 如果请求的字段不为空，则先查询一遍数据库
		tx = tx.Select(preloads).Where("id = ? ", id).First(&rs)
		// 如果查询结果含有错误，则返回错误
		if err := tx.Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				err = errors.New("0 rows affected or returned")
				return nil, err
			}
		}
	}
	// 删除
	tx = tx.Delete(nil)
	if err = tx.Error; err != nil {
		return nil, err
	}
	return &rs, nil
}

func (r *mutationResolver) InsertVehicleSupervisionPhoto(ctx context.Context, objects []*model.VehicleSupervisionPhotoInsertInput, onConflict *model.VehicleSupervisionPhotoOnConflict) (*model.VehicleSupervisionPhotoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertVehicleSupervisionPhotoOne(ctx context.Context, object model.VehicleSupervisionPhotoInsertInput, onConflict *model.VehicleSupervisionPhotoOnConflict) (*model.VehicleSupervisionPhoto, error) {
	v := &model.VehicleSupervisionPhoto{}
	util.StructAssign(v, &object)
	err := db.DB.Create(v).Error
	if err != nil {
		fmt.Println(err)
	}
	return v, err
}

func (r *mutationResolver) UpdateVehicleSupervisionPhoto(ctx context.Context, inc *model.VehicleSupervisionPhotoIncInput, set *model.VehicleSupervisionPhotoSetInput, where model.VehicleSupervisionPhotoBoolExp) (*model.VehicleSupervisionPhotoMutationResponse, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.VehicleSupervisionPhoto{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return &model.VehicleSupervisionPhotoMutationResponse{
			AffectedRows: 0,
		}, err
	}
	var vehicleList []*model.VehicleSupervisionPhoto
	tx.Scan(&vehicleList)
	return &model.VehicleSupervisionPhotoMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    vehicleList,
	}, err
}

func (r *mutationResolver) UpdateVehicleSupervisionPhotoByPk(ctx context.Context, inc *model.VehicleSupervisionPhotoIncInput, set *model.VehicleSupervisionPhotoSetInput, pkColumns model.VehicleSupervisionPhotoPkColumnsInput) (*model.VehicleSupervisionPhoto, error) {
	var err error
	tx := db.DB.Where("id = ? ", pkColumns.ID)
	qt := util2.NewQueryTranslator(tx, &model.VehicleSupervisionPhoto{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return nil, err
	}
	var rs model.VehicleSupervisionPhoto
	tx = tx.First(&rs)
	return &rs, err
}

func (r *queryResolver) VehicleSupervisionPhoto(ctx context.Context, distinctOn []model.VehicleSupervisionPhotoSelectColumn, limit *int, offset *int, orderBy []*model.VehicleSupervisionPhotoOrderBy, where *model.VehicleSupervisionPhotoBoolExp) ([]*model.VehicleSupervisionPhoto, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.VehicleSupervisionPhoto{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model.VehicleSupervisionPhoto
	tx = tx.Find(&rs)
	if err = tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, err
}

func (r *queryResolver) VehicleSupervisionPhotoAggregate(ctx context.Context, distinctOn []model.VehicleSupervisionPhotoSelectColumn, limit *int, offset *int, orderBy []*model.VehicleSupervisionPhotoOrderBy, where *model.VehicleSupervisionPhotoBoolExp) (*model.VehicleSupervisionPhotoAggregate, error) {
	var rs model.VehicleSupervisionPhotoAggregate
	qt := util2.NewQueryTranslator(db.DB, &model.VehicleSupervisionPhoto{})
	tx, err := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Aggregate(&rs, ctx)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if err = tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}
	if err != nil {
		fmt.Println(err)
	}
	return &rs, err
}

func (r *queryResolver) VehicleSupervisionPhotoByPk(ctx context.Context, id int64, supervisionPhotoID string) (*model.VehicleSupervisionPhoto, error) {
	var rs model.VehicleSupervisionPhoto
	err := db.DB.Where("id = ?", id).First(&rs).Error
	return &rs, err
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
