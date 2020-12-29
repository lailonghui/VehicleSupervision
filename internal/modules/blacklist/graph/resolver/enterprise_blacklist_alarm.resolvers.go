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

func (r *mutationResolver) DeleteEnterpriseBlacklistAlarm(ctx context.Context, where model.EnterpriseBlacklistAlarmBoolExp) (*model.EnterpriseBlacklistAlarmMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EnterpriseBlacklistAlarm{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.EnterpriseBlacklistAlarm
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
	return &model.EnterpriseBlacklistAlarmMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteEnterpriseBlacklistAlarmByPk(ctx context.Context, Id int64) (*model1.EnterpriseBlacklistAlarm, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.EnterpriseBlacklistAlarm
	tx := db.DB.Model(&model1.EnterpriseBlacklistAlarm{})
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

func (r *mutationResolver) InsertEnterpriseBlacklistAlarm(ctx context.Context, objects []*model.EnterpriseBlacklistAlarmInsertInput) (*model.EnterpriseBlacklistAlarmMutationResponse, error) {
	rs := make([]*model1.EnterpriseBlacklistAlarm, 0)
	for _, object := range objects {
		v := &model1.EnterpriseBlacklistAlarm{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.EnterpriseBlacklistAlarm{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.EnterpriseBlacklistAlarmMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertEnterpriseBlacklistAlarmOne(ctx context.Context, object model.EnterpriseBlacklistAlarmInsertInput) (*model1.EnterpriseBlacklistAlarm, error) {
	rs := &model1.EnterpriseBlacklistAlarm{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.EnterpriseBlacklistAlarm{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateEnterpriseBlacklistAlarm(ctx context.Context, inc *model.EnterpriseBlacklistAlarmIncInput, set *model.EnterpriseBlacklistAlarmSetInput, where model.EnterpriseBlacklistAlarmBoolExp) (*model.EnterpriseBlacklistAlarmMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EnterpriseBlacklistAlarm{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.EnterpriseBlacklistAlarmMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.EnterpriseBlacklistAlarm
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
	return &model.EnterpriseBlacklistAlarmMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateEnterpriseBlacklistAlarmByPk(ctx context.Context, inc *model.EnterpriseBlacklistAlarmIncInput, set *model.EnterpriseBlacklistAlarmSetInput, Id int64) (*model1.EnterpriseBlacklistAlarm, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.EnterpriseBlacklistAlarm{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.EnterpriseBlacklistAlarm
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) EnterpriseBlacklistAlarm(ctx context.Context, distinctOn []model.EnterpriseBlacklistAlarmSelectColumn, limit *int, offset *int, orderBy []*model.EnterpriseBlacklistAlarmOrderBy, where *model.EnterpriseBlacklistAlarmBoolExp) ([]*model1.EnterpriseBlacklistAlarm, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EnterpriseBlacklistAlarm{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.EnterpriseBlacklistAlarm
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) EnterpriseBlacklistAlarmAggregate(ctx context.Context, distinctOn []model.EnterpriseBlacklistAlarmSelectColumn, limit *int, offset *int, orderBy []*model.EnterpriseBlacklistAlarmOrderBy, where *model.EnterpriseBlacklistAlarmBoolExp) (*model.EnterpriseBlacklistAlarmAggregate, error) {
	var rs model.EnterpriseBlacklistAlarmAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.EnterpriseBlacklistAlarm{})
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

func (r *queryResolver) EnterpriseBlacklistAlarmByPk(ctx context.Context, Id int64) (*model1.EnterpriseBlacklistAlarm, error) {
	var rs model1.EnterpriseBlacklistAlarm
	tx := db.DB.Model(&model1.EnterpriseBlacklistAlarm{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}
