package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/muck_truck_recommend_catalog/graph/model"
	model1 "VehicleSupervision/internal/modules/muck_truck_recommend_catalog/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteNewMuckTruckPhoto(ctx context.Context, where model.NewMuckTruckPhotoBoolExp) (*model.NewMuckTruckPhotoMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.NewMuckTruckPhoto{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.NewMuckTruckPhoto
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
	return &model.NewMuckTruckPhotoMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteNewMuckTruckPhotoByPk(ctx context.Context, Id int64) (*model1.NewMuckTruckPhoto, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.NewMuckTruckPhoto
	tx := db.DB.Model(&model1.NewMuckTruckPhoto{})
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

func (r *mutationResolver) InsertNewMuckTruckPhoto(ctx context.Context, objects []*model.NewMuckTruckPhotoInsertInput) (*model.NewMuckTruckPhotoMutationResponse, error) {
	rs := make([]*model1.NewMuckTruckPhoto, 0)
	for _, object := range objects {
		v := &model1.NewMuckTruckPhoto{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.NewMuckTruckPhoto{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.NewMuckTruckPhotoMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertNewMuckTruckPhotoOne(ctx context.Context, object model.NewMuckTruckPhotoInsertInput) (*model1.NewMuckTruckPhoto, error) {
	rs := &model1.NewMuckTruckPhoto{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.NewMuckTruckPhoto{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateNewMuckTruckPhoto(ctx context.Context, inc *model.NewMuckTruckPhotoIncInput, set *model.NewMuckTruckPhotoSetInput, where model.NewMuckTruckPhotoBoolExp) (*model.NewMuckTruckPhotoMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.NewMuckTruckPhoto{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.NewMuckTruckPhotoMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.NewMuckTruckPhotoMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateNewMuckTruckPhotoByPk(ctx context.Context, inc *model.NewMuckTruckPhotoIncInput, set *model.NewMuckTruckPhotoSetInput, Id int64) (*model1.NewMuckTruckPhoto, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.NewMuckTruckPhoto{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.NewMuckTruckPhoto
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) NewMuckTruckPhoto(ctx context.Context, distinctOn []model.NewMuckTruckPhotoSelectColumn, limit *int, offset *int, orderBy []*model.NewMuckTruckPhotoOrderBy, where *model.NewMuckTruckPhotoBoolExp) ([]*model1.NewMuckTruckPhoto, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.NewMuckTruckPhoto{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.NewMuckTruckPhoto
	tx = tx.Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) NewMuckTruckPhotoAggregate(ctx context.Context, distinctOn []model.NewMuckTruckPhotoSelectColumn, limit *int, offset *int, orderBy []*model.NewMuckTruckPhotoOrderBy, where *model.NewMuckTruckPhotoBoolExp) (*model.NewMuckTruckPhotoAggregate, error) {
	var rs model.NewMuckTruckPhotoAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.NewMuckTruckPhoto{})
	tx, err := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Aggregate(&rs, ctx)
	if err != nil {
		return nil, err
	}
	err = tx.Error
	return &rs, err
}

func (r *queryResolver) NewMuckTruckPhotoByPk(ctx context.Context, Id int64) (*model1.NewMuckTruckPhoto, error) {
	var rs model1.NewMuckTruckPhoto
	tx := db.DB.Model(&model1.NewMuckTruckPhoto{}).First(&rs, Id)
	err := tx.Error
	return &rs, err
}
