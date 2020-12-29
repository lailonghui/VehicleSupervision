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

func (r *mutationResolver) DeleteElectricFence(ctx context.Context, where model.ElectricFenceBoolExp) (*model.ElectricFenceMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.ElectricFence{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.ElectricFence
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
	return &model.ElectricFenceMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteElectricFenceByPk(ctx context.Context, Id int64) (*model1.ElectricFence, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.ElectricFence
	tx := db.DB.Model(&model1.ElectricFence{})
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

func (r *mutationResolver) InsertElectricFence(ctx context.Context, objects []*model.ElectricFenceInsertInput) (*model.ElectricFenceMutationResponse, error) {
	rs := make([]*model1.ElectricFence, 0)
	for _, object := range objects {
		v := &model1.ElectricFence{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.ElectricFence{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.ElectricFenceMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertElectricFenceOne(ctx context.Context, object model.ElectricFenceInsertInput) (*model1.ElectricFence, error) {
	rs := &model1.ElectricFence{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.ElectricFence{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateElectricFence(ctx context.Context, inc *model.ElectricFenceIncInput, set *model.ElectricFenceSetInput, where model.ElectricFenceBoolExp) (*model.ElectricFenceMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.ElectricFence{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.ElectricFenceMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.ElectricFence
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
	return &model.ElectricFenceMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateElectricFenceByPk(ctx context.Context, inc *model.ElectricFenceIncInput, set *model.ElectricFenceSetInput, Id int64) (*model1.ElectricFence, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.ElectricFence{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.ElectricFence
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) ElectricFence(ctx context.Context, distinctOn []model.ElectricFenceSelectColumn, limit *int, offset *int, orderBy []*model.ElectricFenceOrderBy, where *model.ElectricFenceBoolExp) ([]*model1.ElectricFence, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.ElectricFence{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.ElectricFence
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) ElectricFenceAggregate(ctx context.Context, distinctOn []model.ElectricFenceSelectColumn, limit *int, offset *int, orderBy []*model.ElectricFenceOrderBy, where *model.ElectricFenceBoolExp) (*model.ElectricFenceAggregate, error) {
	var rs model.ElectricFenceAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.ElectricFence{})
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

func (r *queryResolver) ElectricFenceByPk(ctx context.Context, Id int64) (*model1.ElectricFence, error) {
	var rs model1.ElectricFence
	tx := db.DB.Model(&model1.ElectricFence{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}
