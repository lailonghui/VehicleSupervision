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

func (r *mutationResolver) DeleteConstructionCamera(ctx context.Context, where model.ConstructionCameraBoolExp) (*model.ConstructionCameraMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.ConstructionCamera{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.ConstructionCamera
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
	return &model.ConstructionCameraMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteConstructionCameraByPk(ctx context.Context, Id int64) (*model1.ConstructionCamera, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.ConstructionCamera
	tx := db.DB.Model(&model1.ConstructionCamera{})
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

func (r *mutationResolver) InsertConstructionCamera(ctx context.Context, objects []*model.ConstructionCameraInsertInput) (*model.ConstructionCameraMutationResponse, error) {
	rs := []*model1.ConstructionCamera{}
	for _, object := range objects {
		v := &model1.ConstructionCamera{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.ConstructionCamera{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.ConstructionCameraMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertConstructionCameraOne(ctx context.Context, object model.ConstructionCameraInsertInput) (*model1.ConstructionCamera, error) {
	rs := &model1.ConstructionCamera{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.ConstructionCamera{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateConstructionCamera(ctx context.Context, inc *model.ConstructionCameraIncInput, set *model.ConstructionCameraSetInput, where model.ConstructionCameraBoolExp) (*model.ConstructionCameraMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.ConstructionCamera{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.ConstructionCameraMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.ConstructionCameraMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateConstructionCameraByPk(ctx context.Context, inc *model.ConstructionCameraIncInput, set *model.ConstructionCameraSetInput, Id int64) (*model1.ConstructionCamera, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.ConstructionCamera{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.ConstructionCamera
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) ConstructionCamera(ctx context.Context, distinctOn []model.ConstructionCameraSelectColumn, limit *int, offset *int, orderBy []*model.ConstructionCameraOrderBy, where *model.ConstructionCameraBoolExp) ([]*model1.ConstructionCamera, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.ConstructionCamera{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.ConstructionCamera
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) ConstructionCameraAggregate(ctx context.Context, distinctOn []model.ConstructionCameraSelectColumn, limit *int, offset *int, orderBy []*model.ConstructionCameraOrderBy, where *model.ConstructionCameraBoolExp) (*model.ConstructionCameraAggregate, error) {
	var rs model.ConstructionCameraAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.ConstructionCamera{})
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

func (r *queryResolver) ConstructionCameraByPk(ctx context.Context, Id int64) (*model1.ConstructionCamera, error) {
	var rs model1.ConstructionCamera
	tx := db.DB.Model(&model1.ConstructionCamera{}).First(&rs, Id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
