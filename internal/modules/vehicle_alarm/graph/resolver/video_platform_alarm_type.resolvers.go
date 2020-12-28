package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/vehicle_alarm/graph/model"
	model1 "VehicleSupervision/internal/modules/vehicle_alarm/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteVideoPlatformAlarmType(ctx context.Context, where model.VideoPlatformAlarmTypeBoolExp) (*model.VideoPlatformAlarmTypeMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VideoPlatformAlarmType{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.VideoPlatformAlarmType
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
	return &model.VideoPlatformAlarmTypeMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteVideoPlatformAlarmTypeByPk(ctx context.Context, Id int64) (*model1.VideoPlatformAlarmType, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.VideoPlatformAlarmType
	tx := db.DB.Model(&model1.VideoPlatformAlarmType{})
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

func (r *mutationResolver) InsertVideoPlatformAlarmType(ctx context.Context, objects []*model.VideoPlatformAlarmTypeInsertInput) (*model.VideoPlatformAlarmTypeMutationResponse, error) {
	rs := []*model1.VideoPlatformAlarmType{}
	for _, object := range objects {
		v := &model1.VideoPlatformAlarmType{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.VideoPlatformAlarmType{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.VideoPlatformAlarmTypeMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertVideoPlatformAlarmTypeOne(ctx context.Context, object model.VideoPlatformAlarmTypeInsertInput) (*model1.VideoPlatformAlarmType, error) {
	rs := &model1.VideoPlatformAlarmType{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.VideoPlatformAlarmType{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateVideoPlatformAlarmType(ctx context.Context, inc *model.VideoPlatformAlarmTypeIncInput, set *model.VideoPlatformAlarmTypeSetInput, where model.VideoPlatformAlarmTypeBoolExp) (*model.VideoPlatformAlarmTypeMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VideoPlatformAlarmType{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.VideoPlatformAlarmTypeMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.VideoPlatformAlarmTypeMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateVideoPlatformAlarmTypeByPk(ctx context.Context, inc *model.VideoPlatformAlarmTypeIncInput, set *model.VideoPlatformAlarmTypeSetInput, Id int64) (*model1.VideoPlatformAlarmType, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.VideoPlatformAlarmType{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.VideoPlatformAlarmType
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) VideoPlatformAlarmType(ctx context.Context, distinctOn []model.VideoPlatformAlarmTypeSelectColumn, limit *int, offset *int, orderBy []*model.VideoPlatformAlarmTypeOrderBy, where *model.VideoPlatformAlarmTypeBoolExp) ([]*model1.VideoPlatformAlarmType, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VideoPlatformAlarmType{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.VideoPlatformAlarmType
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) VideoPlatformAlarmTypeAggregate(ctx context.Context, distinctOn []model.VideoPlatformAlarmTypeSelectColumn, limit *int, offset *int, orderBy []*model.VideoPlatformAlarmTypeOrderBy, where *model.VideoPlatformAlarmTypeBoolExp) (*model.VideoPlatformAlarmTypeAggregate, error) {
	var rs model.VideoPlatformAlarmTypeAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.VideoPlatformAlarmType{})
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

func (r *queryResolver) VideoPlatformAlarmTypeByPk(ctx context.Context, Id int64) (*model1.VideoPlatformAlarmType, error) {
	var rs model1.VideoPlatformAlarmType
	tx := db.DB.Model(&model1.VideoPlatformAlarmType{}).First(&rs, Id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
