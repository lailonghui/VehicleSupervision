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

func (r *mutationResolver) DeleteAppEnforcement(ctx context.Context, where model.AppEnforcementBoolExp) (*model.AppEnforcementMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.AppEnforcement{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.AppEnforcement
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
	return &model.AppEnforcementMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteAppEnforcementByPk(ctx context.Context, Id int64) (*model1.AppEnforcement, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.AppEnforcement
	tx := db.DB.Model(&model1.AppEnforcement{})
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

func (r *mutationResolver) InsertAppEnforcement(ctx context.Context, objects []*model.AppEnforcementInsertInput) (*model.AppEnforcementMutationResponse, error) {
	rs := []*model1.AppEnforcement{}
	for _, object := range objects {
		v := &model1.AppEnforcement{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.AppEnforcement{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.AppEnforcementMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertAppEnforcementOne(ctx context.Context, object model.AppEnforcementInsertInput) (*model1.AppEnforcement, error) {
	rs := &model1.AppEnforcement{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.AppEnforcement{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateAppEnforcement(ctx context.Context, inc *model.AppEnforcementIncInput, set *model.AppEnforcementSetInput, where model.AppEnforcementBoolExp) (*model.AppEnforcementMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.AppEnforcement{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.AppEnforcementMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.AppEnforcementMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateAppEnforcementByPk(ctx context.Context, inc *model.AppEnforcementIncInput, set *model.AppEnforcementSetInput, Id int64) (*model1.AppEnforcement, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.AppEnforcement{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.AppEnforcement
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) AppEnforcement(ctx context.Context, distinctOn []model.AppEnforcementSelectColumn, limit *int, offset *int, orderBy []*model.AppEnforcementOrderBy, where *model.AppEnforcementBoolExp) ([]*model1.AppEnforcement, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.AppEnforcement{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.AppEnforcement
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) AppEnforcementAggregate(ctx context.Context, distinctOn []model.AppEnforcementSelectColumn, limit *int, offset *int, orderBy []*model.AppEnforcementOrderBy, where *model.AppEnforcementBoolExp) (*model.AppEnforcementAggregate, error) {
	var rs model.AppEnforcementAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.AppEnforcement{})
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

func (r *queryResolver) AppEnforcementByPk(ctx context.Context, Id int64) (*model1.AppEnforcement, error) {
	var rs model1.AppEnforcement
	tx := db.DB.Model(&model1.AppEnforcement{}).First(&rs, Id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
