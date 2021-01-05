package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/blacklist/graph/model"
	model1 "VehicleSupervision/internal/modules/blacklist/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteVehicleBlacklistAlarm(ctx context.Context, where model.VehicleBlacklistAlarmBoolExp) (*model.VehicleBlacklistAlarmMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleBlacklistAlarm{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.VehicleBlacklistAlarm
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
	return &model.VehicleBlacklistAlarmMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteVehicleBlacklistAlarmByPk(ctx context.Context, Id int64) (*model1.VehicleBlacklistAlarm, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.VehicleBlacklistAlarm
	tx := db.DB.Model(&model1.VehicleBlacklistAlarm{})
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

func (r *mutationResolver) DeleteVehicleBlacklistAlarmByUnionPk(ctx context.Context, unionId string) (*model1.VehicleBlacklistAlarm, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.VehicleBlacklistAlarm
	tx := db.DB.Model(&model1.VehicleBlacklistAlarm{})
	if len(preloads) > 0 {
		// 如果请求的字段不为空，则先查询一遍数据库
		tx = tx.Select(preloads).Where(rs.UnionPrimaryColumnName()+" = ?", unionId).First(&rs)
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

func (r *mutationResolver) InsertVehicleBlacklistAlarm(ctx context.Context, objects []*model.VehicleBlacklistAlarmInsertInput) (*model.VehicleBlacklistAlarmMutationResponse, error) {
	rs := make([]*model1.VehicleBlacklistAlarm, 0)
	for _, object := range objects {
		v := &model1.VehicleBlacklistAlarm{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.VehicleBlacklistAlarm{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.VehicleBlacklistAlarmMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertVehicleBlacklistAlarmOne(ctx context.Context, object model.VehicleBlacklistAlarmInsertInput) (*model1.VehicleBlacklistAlarm, error) {
	rs := &model1.VehicleBlacklistAlarm{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.VehicleBlacklistAlarm{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateVehicleBlacklistAlarm(ctx context.Context, inc *model.VehicleBlacklistAlarmIncInput, set *model.VehicleBlacklistAlarmSetInput, where model.VehicleBlacklistAlarmBoolExp) (*model.VehicleBlacklistAlarmMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleBlacklistAlarm{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.VehicleBlacklistAlarmMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.VehicleBlacklistAlarm
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
	return &model.VehicleBlacklistAlarmMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateVehicleBlacklistAlarmByPk(ctx context.Context, inc *model.VehicleBlacklistAlarmIncInput, set *model.VehicleBlacklistAlarmSetInput, Id int64) (*model1.VehicleBlacklistAlarm, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.VehicleBlacklistAlarm{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.VehicleBlacklistAlarm
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *mutationResolver) UpdateVehicleBlacklistAlarmByUnionPk(ctx context.Context, inc *model.VehicleBlacklistAlarmIncInput, set *model.VehicleBlacklistAlarmSetInput, unionId string) (*model1.VehicleBlacklistAlarm, error) {
	var rs model1.VehicleBlacklistAlarm
	tx := db.DB.Where(rs.UnionPrimaryColumnName()+" = ?", unionId)
	qt := util.NewQueryTranslator(tx, &model1.VehicleBlacklistAlarm{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}

	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) VehicleBlacklistAlarm(ctx context.Context, distinctOn []model.VehicleBlacklistAlarmSelectColumn, limit *int, offset *int, orderBy []*model.VehicleBlacklistAlarmOrderBy, where *model.VehicleBlacklistAlarmBoolExp) ([]*model1.VehicleBlacklistAlarm, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleBlacklistAlarm{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.VehicleBlacklistAlarm
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) VehicleBlacklistAlarmAggregate(ctx context.Context, distinctOn []model.VehicleBlacklistAlarmSelectColumn, limit *int, offset *int, orderBy []*model.VehicleBlacklistAlarmOrderBy, where *model.VehicleBlacklistAlarmBoolExp) (*model.VehicleBlacklistAlarmAggregate, error) {
	var rs model.VehicleBlacklistAlarmAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.VehicleBlacklistAlarm{})
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

func (r *queryResolver) VehicleBlacklistAlarmByPk(ctx context.Context, Id int64) (*model1.VehicleBlacklistAlarm, error) {
	var rs model1.VehicleBlacklistAlarm
	tx := db.DB.Model(&model1.VehicleBlacklistAlarm{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}

func (r *queryResolver) VehicleBlacklistAlarmByUnionPk(ctx context.Context, unionId string) (*model1.VehicleBlacklistAlarm, error) {
	var rs model1.VehicleBlacklistAlarm
	tx := db.DB.Model(&model1.VehicleBlacklistAlarm{}).Select(util.GetTopPreloads(ctx)).Where(rs.UnionPrimaryColumnName()+" = ?", unionId).First(&rs)

	err := tx.Error
	return &rs, err
}
