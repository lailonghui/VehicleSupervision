package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/vehicle_violation/graph/model"
	model1 "VehicleSupervision/internal/modules/vehicle_violation/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteIllegalPhoto(ctx context.Context, where model.IllegalPhotoBoolExp) (*model.IllegalPhotoMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.IllegalPhoto{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.IllegalPhoto
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
	return &model.IllegalPhotoMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteIllegalPhotoByPk(ctx context.Context, Id int64) (*model1.IllegalPhoto, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.IllegalPhoto
	tx := db.DB.Model(&model1.IllegalPhoto{})
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

func (r *mutationResolver) InsertIllegalPhoto(ctx context.Context, objects []*model.IllegalPhotoInsertInput) (*model.IllegalPhotoMutationResponse, error) {
	rs := make([]*model1.IllegalPhoto, 0)
	for _, object := range objects {
		v := &model1.IllegalPhoto{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.IllegalPhoto{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.IllegalPhotoMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertIllegalPhotoOne(ctx context.Context, object model.IllegalPhotoInsertInput) (*model1.IllegalPhoto, error) {
	rs := &model1.IllegalPhoto{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.IllegalPhoto{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateIllegalPhoto(ctx context.Context, inc *model.IllegalPhotoIncInput, set *model.IllegalPhotoSetInput, where model.IllegalPhotoBoolExp) (*model.IllegalPhotoMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.IllegalPhoto{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.IllegalPhotoMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.IllegalPhotoMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateIllegalPhotoByPk(ctx context.Context, inc *model.IllegalPhotoIncInput, set *model.IllegalPhotoSetInput, Id int64) (*model1.IllegalPhoto, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.IllegalPhoto{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.IllegalPhoto
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) IllegalPhoto(ctx context.Context, distinctOn []model.IllegalPhotoSelectColumn, limit *int, offset *int, orderBy []*model.IllegalPhotoOrderBy, where *model.IllegalPhotoBoolExp) ([]*model1.IllegalPhoto, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.IllegalPhoto{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.IllegalPhoto
	tx = tx.Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) IllegalPhotoAggregate(ctx context.Context, distinctOn []model.IllegalPhotoSelectColumn, limit *int, offset *int, orderBy []*model.IllegalPhotoOrderBy, where *model.IllegalPhotoBoolExp) (*model.IllegalPhotoAggregate, error) {
	var rs model.IllegalPhotoAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.IllegalPhoto{})
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

func (r *queryResolver) IllegalPhotoByPk(ctx context.Context, Id int64) (*model1.IllegalPhoto, error) {
	var rs model1.IllegalPhoto
	tx := db.DB.Model(&model1.IllegalPhoto{}).First(&rs, Id)
	err := tx.Error
	return &rs, err
}
