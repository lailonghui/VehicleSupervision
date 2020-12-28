package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/vehicle/graph/model"
	model1 "VehicleSupervision/internal/modules/vehicle/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteVehicleSupervisionPhoto(ctx context.Context, where model.VehicleSupervisionPhotoBoolExp) (*model.VehicleSupervisionPhotoMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleSupervisionPhoto{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.VehicleSupervisionPhoto
	if len(preloads) > 0 {
		// 如果请求的字段不为空，则先查询一遍数据库
		tx := tx.Select(preloads)
		tx = tx.Find(&rs)
		// 如果查询结果含有错误，则返回错误
		if err := tx.Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, nil
			}
			return nil, err
		}
	}
	// 删除
	tx = tx.Delete(nil)
	if err := tx.Error; err != nil {
		return nil, err
	}
	return &model.VehicleSupervisionPhotoMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteVehicleSupervisionPhotoByPk(ctx context.Context, Id int64) (*model1.VehicleSupervisionPhoto, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.VehicleSupervisionPhoto
	tx := db.DB.Model(&model1.VehicleSupervisionPhoto{})
	if len(preloads) > 0 {
		// 如果请求的字段不为空，则先查询一遍数据库
		tx = tx.Select(preloads).Where("id = ?", Id).First(&rs)
		// 如果查询结果含有错误，则返回错误
		if err := tx.Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, nil
			}
			return nil, err
		}
	}
	// 删除
	tx = tx.Delete(nil)
	if err := tx.Error; err != nil {
		return nil, err
	}
	return &rs, nil
}

func (r *mutationResolver) InsertVehicleSupervisionPhoto(ctx context.Context, objects []*model.VehicleSupervisionPhotoInsertInput) (*model.VehicleSupervisionPhotoMutationResponse, error) {
	rs := []*model1.VehicleSupervisionPhoto{}
	for _, object := range objects {
		v := &model1.VehicleSupervisionPhoto{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.VehicleSupervisionPhoto{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.VehicleSupervisionPhotoMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertVehicleSupervisionPhotoOne(ctx context.Context, object model.VehicleSupervisionPhotoInsertInput) (*model1.VehicleSupervisionPhoto, error) {
	rs := &model1.VehicleSupervisionPhoto{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.VehicleSupervisionPhoto{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateVehicleSupervisionPhoto(ctx context.Context, inc *model.VehicleSupervisionPhotoIncInput, set *model.VehicleSupervisionPhotoSetInput, where model.VehicleSupervisionPhotoBoolExp) (*model.VehicleSupervisionPhotoMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleSupervisionPhoto{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.VehicleSupervisionPhotoMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.VehicleSupervisionPhotoMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateVehicleSupervisionPhotoByPk(ctx context.Context, inc *model.VehicleSupervisionPhotoIncInput, set *model.VehicleSupervisionPhotoSetInput, Id int64) (*model1.VehicleSupervisionPhoto, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.VehicleSupervisionPhoto{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.VehicleSupervisionPhoto
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) VehicleSupervisionPhoto(ctx context.Context, distinctOn []model.VehicleSupervisionPhotoSelectColumn, limit *int, offset *int, orderBy []*model.VehicleSupervisionPhotoOrderBy, where *model.VehicleSupervisionPhotoBoolExp) ([]*model1.VehicleSupervisionPhoto, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleSupervisionPhoto{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.VehicleSupervisionPhoto
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) VehicleSupervisionPhotoAggregate(ctx context.Context, distinctOn []model.VehicleSupervisionPhotoSelectColumn, limit *int, offset *int, orderBy []*model.VehicleSupervisionPhotoOrderBy, where *model.VehicleSupervisionPhotoBoolExp) (*model.VehicleSupervisionPhotoAggregate, error) {
	var rs model.VehicleSupervisionPhotoAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.VehicleSupervisionPhoto{})
	tx, err := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Aggregate(&rs, ctx)
	if err != nil {
		return nil, err
	}
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &rs, nil
}

func (r *queryResolver) VehicleSupervisionPhotoByPk(ctx context.Context, Id int64) (*model1.VehicleSupervisionPhoto, error) {
	var rs model1.VehicleSupervisionPhoto
	tx := db.DB.Model(&model1.VehicleSupervisionPhoto{}).First(&rs, Id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
