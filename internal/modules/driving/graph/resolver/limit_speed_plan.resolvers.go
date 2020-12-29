package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/driving/graph/model"
	model1 "VehicleSupervision/internal/modules/driving/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteLimitSpeedPlan(ctx context.Context, where model.LimitSpeedPlanBoolExp) (*model.LimitSpeedPlanMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.LimitSpeedPlan{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.LimitSpeedPlan
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
	return &model.LimitSpeedPlanMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteLimitSpeedPlanByPk(ctx context.Context, Id int64) (*model1.LimitSpeedPlan, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.LimitSpeedPlan
	tx := db.DB.Model(&model1.LimitSpeedPlan{})
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

func (r *mutationResolver) InsertLimitSpeedPlan(ctx context.Context, objects []*model.LimitSpeedPlanInsertInput) (*model.LimitSpeedPlanMutationResponse, error) {
	rs := make([]*model1.LimitSpeedPlan, 0)
	for _, object := range objects {
		v := &model1.LimitSpeedPlan{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.LimitSpeedPlan{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.LimitSpeedPlanMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertLimitSpeedPlanOne(ctx context.Context, object model.LimitSpeedPlanInsertInput) (*model1.LimitSpeedPlan, error) {
	rs := &model1.LimitSpeedPlan{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.LimitSpeedPlan{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateLimitSpeedPlan(ctx context.Context, inc *model.LimitSpeedPlanIncInput, set *model.LimitSpeedPlanSetInput, where model.LimitSpeedPlanBoolExp) (*model.LimitSpeedPlanMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.LimitSpeedPlan{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.LimitSpeedPlanMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.LimitSpeedPlan
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
	return &model.LimitSpeedPlanMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateLimitSpeedPlanByPk(ctx context.Context, inc *model.LimitSpeedPlanIncInput, set *model.LimitSpeedPlanSetInput, Id int64) (*model1.LimitSpeedPlan, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.LimitSpeedPlan{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.LimitSpeedPlan
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) LimitSpeedPlan(ctx context.Context, distinctOn []model.LimitSpeedPlanSelectColumn, limit *int, offset *int, orderBy []*model.LimitSpeedPlanOrderBy, where *model.LimitSpeedPlanBoolExp) ([]*model1.LimitSpeedPlan, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.LimitSpeedPlan{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.LimitSpeedPlan
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) LimitSpeedPlanAggregate(ctx context.Context, distinctOn []model.LimitSpeedPlanSelectColumn, limit *int, offset *int, orderBy []*model.LimitSpeedPlanOrderBy, where *model.LimitSpeedPlanBoolExp) (*model.LimitSpeedPlanAggregate, error) {
	var rs model.LimitSpeedPlanAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.LimitSpeedPlan{})
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

func (r *queryResolver) LimitSpeedPlanByPk(ctx context.Context, Id int64) (*model1.LimitSpeedPlan, error) {
	var rs model1.LimitSpeedPlan
	tx := db.DB.Model(&model1.LimitSpeedPlan{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}
