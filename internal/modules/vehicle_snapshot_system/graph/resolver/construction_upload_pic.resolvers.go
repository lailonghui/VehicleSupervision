package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/vehicle_snapshot_system/graph/model"
	model1 "VehicleSupervision/internal/modules/vehicle_snapshot_system/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteConstructionUploadPic(ctx context.Context, where model.ConstructionUploadPicBoolExp) (*model.ConstructionUploadPicMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.ConstructionUploadPic{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.ConstructionUploadPic
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
	return &model.ConstructionUploadPicMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteConstructionUploadPicByPk(ctx context.Context, Id int64) (*model1.ConstructionUploadPic, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.ConstructionUploadPic
	tx := db.DB.Model(&model1.ConstructionUploadPic{})
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

func (r *mutationResolver) InsertConstructionUploadPic(ctx context.Context, objects []*model.ConstructionUploadPicInsertInput) (*model.ConstructionUploadPicMutationResponse, error) {
	rs := make([]*model1.ConstructionUploadPic, 0)
	for _, object := range objects {
		v := &model1.ConstructionUploadPic{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.ConstructionUploadPic{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.ConstructionUploadPicMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertConstructionUploadPicOne(ctx context.Context, object model.ConstructionUploadPicInsertInput) (*model1.ConstructionUploadPic, error) {
	rs := &model1.ConstructionUploadPic{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.ConstructionUploadPic{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateConstructionUploadPic(ctx context.Context, inc *model.ConstructionUploadPicIncInput, set *model.ConstructionUploadPicSetInput, where model.ConstructionUploadPicBoolExp) (*model.ConstructionUploadPicMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.ConstructionUploadPic{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.ConstructionUploadPicMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.ConstructionUploadPicMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateConstructionUploadPicByPk(ctx context.Context, inc *model.ConstructionUploadPicIncInput, set *model.ConstructionUploadPicSetInput, Id int64) (*model1.ConstructionUploadPic, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.ConstructionUploadPic{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.ConstructionUploadPic
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) ConstructionUploadPic(ctx context.Context, distinctOn []model.ConstructionUploadPicSelectColumn, limit *int, offset *int, orderBy []*model.ConstructionUploadPicOrderBy, where *model.ConstructionUploadPicBoolExp) ([]*model1.ConstructionUploadPic, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.ConstructionUploadPic{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.ConstructionUploadPic
	tx = tx.Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) ConstructionUploadPicAggregate(ctx context.Context, distinctOn []model.ConstructionUploadPicSelectColumn, limit *int, offset *int, orderBy []*model.ConstructionUploadPicOrderBy, where *model.ConstructionUploadPicBoolExp) (*model.ConstructionUploadPicAggregate, error) {
	var rs model.ConstructionUploadPicAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.ConstructionUploadPic{})
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

func (r *queryResolver) ConstructionUploadPicByPk(ctx context.Context, Id int64) (*model1.ConstructionUploadPic, error) {
	var rs model1.ConstructionUploadPic
	tx := db.DB.Model(&model1.ConstructionUploadPic{}).First(&rs, Id)
	err := tx.Error
	return &rs, err
}