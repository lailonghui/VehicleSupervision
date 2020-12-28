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

func (r *mutationResolver) DeleteVehicleOnlineTime(ctx context.Context, where model.VehicleOnlineTimeBoolExp) (*model.VehicleOnlineTimeMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleOnlineTime{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.VehicleOnlineTime
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
	return &model.VehicleOnlineTimeMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteVehicleOnlineTimeByPk(ctx context.Context, Id int64) (*model1.VehicleOnlineTime, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.VehicleOnlineTime
	tx := db.DB.Model(&model1.VehicleOnlineTime{})
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

func (r *mutationResolver) InsertVehicleOnlineTime(ctx context.Context, objects []*model.VehicleOnlineTimeInsertInput) (*model.VehicleOnlineTimeMutationResponse, error) {
	rs := make([]*model1.VehicleOnlineTime, 0)
	for _, object := range objects {
		v := &model1.VehicleOnlineTime{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.VehicleOnlineTime{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.VehicleOnlineTimeMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertVehicleOnlineTimeOne(ctx context.Context, object model.VehicleOnlineTimeInsertInput) (*model1.VehicleOnlineTime, error) {
	rs := &model1.VehicleOnlineTime{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.VehicleOnlineTime{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateVehicleOnlineTime(ctx context.Context, inc *model.VehicleOnlineTimeIncInput, set *model.VehicleOnlineTimeSetInput, where model.VehicleOnlineTimeBoolExp) (*model.VehicleOnlineTimeMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleOnlineTime{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.VehicleOnlineTimeMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.VehicleOnlineTimeMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateVehicleOnlineTimeByPk(ctx context.Context, inc *model.VehicleOnlineTimeIncInput, set *model.VehicleOnlineTimeSetInput, Id int64) (*model1.VehicleOnlineTime, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.VehicleOnlineTime{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.VehicleOnlineTime
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) VehicleOnlineTime(ctx context.Context, distinctOn []model.VehicleOnlineTimeSelectColumn, limit *int, offset *int, orderBy []*model.VehicleOnlineTimeOrderBy, where *model.VehicleOnlineTimeBoolExp) ([]*model1.VehicleOnlineTime, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleOnlineTime{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.VehicleOnlineTime
	tx = tx.Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) VehicleOnlineTimeAggregate(ctx context.Context, distinctOn []model.VehicleOnlineTimeSelectColumn, limit *int, offset *int, orderBy []*model.VehicleOnlineTimeOrderBy, where *model.VehicleOnlineTimeBoolExp) (*model.VehicleOnlineTimeAggregate, error) {
	var rs model.VehicleOnlineTimeAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.VehicleOnlineTime{})
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

func (r *queryResolver) VehicleOnlineTimeByPk(ctx context.Context, Id int64) (*model1.VehicleOnlineTime, error) {
	var rs model1.VehicleOnlineTime
	tx := db.DB.Model(&model1.VehicleOnlineTime{}).First(&rs, Id)
	err := tx.Error
	return &rs, err
}
