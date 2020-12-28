package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/vehicle_driver_separation/graph/model"
	model1 "VehicleSupervision/internal/modules/vehicle_driver_separation/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteDrivingLogInfo(ctx context.Context, where model.DrivingLogInfoBoolExp) (*model.DrivingLogInfoMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DrivingLogInfo{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.DrivingLogInfo
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
	return &model.DrivingLogInfoMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteDrivingLogInfoByPk(ctx context.Context, Id int64) (*model1.DrivingLogInfo, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.DrivingLogInfo
	tx := db.DB.Model(&model1.DrivingLogInfo{})
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

func (r *mutationResolver) InsertDrivingLogInfo(ctx context.Context, objects []*model.DrivingLogInfoInsertInput) (*model.DrivingLogInfoMutationResponse, error) {
	rs := []*model1.DrivingLogInfo{}
	for _, object := range objects {
		v := &model1.DrivingLogInfo{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.DrivingLogInfo{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.DrivingLogInfoMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertDrivingLogInfoOne(ctx context.Context, object model.DrivingLogInfoInsertInput) (*model1.DrivingLogInfo, error) {
	rs := &model1.DrivingLogInfo{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.DrivingLogInfo{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateDrivingLogInfo(ctx context.Context, inc *model.DrivingLogInfoIncInput, set *model.DrivingLogInfoSetInput, where model.DrivingLogInfoBoolExp) (*model.DrivingLogInfoMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DrivingLogInfo{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.DrivingLogInfoMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.DrivingLogInfoMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateDrivingLogInfoByPk(ctx context.Context, inc *model.DrivingLogInfoIncInput, set *model.DrivingLogInfoSetInput, Id int64) (*model1.DrivingLogInfo, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.DrivingLogInfo{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.DrivingLogInfo
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) DrivingLogInfo(ctx context.Context, distinctOn []model.DrivingLogInfoSelectColumn, limit *int, offset *int, orderBy []*model.DrivingLogInfoOrderBy, where *model.DrivingLogInfoBoolExp) ([]*model1.DrivingLogInfo, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DrivingLogInfo{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.DrivingLogInfo
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) DrivingLogInfoAggregate(ctx context.Context, distinctOn []model.DrivingLogInfoSelectColumn, limit *int, offset *int, orderBy []*model.DrivingLogInfoOrderBy, where *model.DrivingLogInfoBoolExp) (*model.DrivingLogInfoAggregate, error) {
	var rs model.DrivingLogInfoAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.DrivingLogInfo{})
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

func (r *queryResolver) DrivingLogInfoByPk(ctx context.Context, Id int64) (*model1.DrivingLogInfo, error) {
	var rs model1.DrivingLogInfo
	tx := db.DB.Model(&model1.DrivingLogInfo{}).First(&rs, Id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
