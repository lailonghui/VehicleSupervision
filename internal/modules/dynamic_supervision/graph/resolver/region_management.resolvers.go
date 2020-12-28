package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/dynamic_supervision/graph/model"
	model1 "VehicleSupervision/internal/modules/dynamic_supervision/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteRegionManagement(ctx context.Context, where model.RegionManagementBoolExp) (*model.RegionManagementMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.RegionManagement{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.RegionManagement
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
	return &model.RegionManagementMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteRegionManagementByPk(ctx context.Context, Id int64) (*model1.RegionManagement, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.RegionManagement
	tx := db.DB.Model(&model1.RegionManagement{})
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

func (r *mutationResolver) InsertRegionManagement(ctx context.Context, objects []*model.RegionManagementInsertInput) (*model.RegionManagementMutationResponse, error) {
	rs := make([]*model1.RegionManagement, 0)
	for _, object := range objects {
		v := &model1.RegionManagement{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.RegionManagement{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.RegionManagementMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertRegionManagementOne(ctx context.Context, object model.RegionManagementInsertInput) (*model1.RegionManagement, error) {
	rs := &model1.RegionManagement{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.RegionManagement{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateRegionManagement(ctx context.Context, inc *model.RegionManagementIncInput, set *model.RegionManagementSetInput, where model.RegionManagementBoolExp) (*model.RegionManagementMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.RegionManagement{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.RegionManagementMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.RegionManagementMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateRegionManagementByPk(ctx context.Context, inc *model.RegionManagementIncInput, set *model.RegionManagementSetInput, Id int64) (*model1.RegionManagement, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.RegionManagement{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.RegionManagement
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) RegionManagement(ctx context.Context, distinctOn []model.RegionManagementSelectColumn, limit *int, offset *int, orderBy []*model.RegionManagementOrderBy, where *model.RegionManagementBoolExp) ([]*model1.RegionManagement, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.RegionManagement{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.RegionManagement
	tx = tx.Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) RegionManagementAggregate(ctx context.Context, distinctOn []model.RegionManagementSelectColumn, limit *int, offset *int, orderBy []*model.RegionManagementOrderBy, where *model.RegionManagementBoolExp) (*model.RegionManagementAggregate, error) {
	var rs model.RegionManagementAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.RegionManagement{})
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

func (r *queryResolver) RegionManagementByPk(ctx context.Context, Id int64) (*model1.RegionManagement, error) {
	var rs model1.RegionManagement
	tx := db.DB.Model(&model1.RegionManagement{}).First(&rs, Id)
	err := tx.Error
	return &rs, err
}
