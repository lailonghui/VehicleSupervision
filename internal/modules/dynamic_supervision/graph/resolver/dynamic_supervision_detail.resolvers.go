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

func (r *mutationResolver) DeleteDynamicSupervisionDetail(ctx context.Context, where model.DynamicSupervisionDetailBoolExp) (*model.DynamicSupervisionDetailMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DynamicSupervisionDetail{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.DynamicSupervisionDetail
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
	return &model.DynamicSupervisionDetailMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteDynamicSupervisionDetailByPk(ctx context.Context, Id int64) (*model1.DynamicSupervisionDetail, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.DynamicSupervisionDetail
	tx := db.DB.Model(&model1.DynamicSupervisionDetail{})
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

func (r *mutationResolver) InsertDynamicSupervisionDetail(ctx context.Context, objects []*model.DynamicSupervisionDetailInsertInput) (*model.DynamicSupervisionDetailMutationResponse, error) {
	rs := []*model1.DynamicSupervisionDetail{}
	for _, object := range objects {
		v := &model1.DynamicSupervisionDetail{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.DynamicSupervisionDetail{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.DynamicSupervisionDetailMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertDynamicSupervisionDetailOne(ctx context.Context, object model.DynamicSupervisionDetailInsertInput) (*model1.DynamicSupervisionDetail, error) {
	rs := &model1.DynamicSupervisionDetail{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.DynamicSupervisionDetail{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateDynamicSupervisionDetail(ctx context.Context, inc *model.DynamicSupervisionDetailIncInput, set *model.DynamicSupervisionDetailSetInput, where model.DynamicSupervisionDetailBoolExp) (*model.DynamicSupervisionDetailMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DynamicSupervisionDetail{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.DynamicSupervisionDetailMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.DynamicSupervisionDetailMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateDynamicSupervisionDetailByPk(ctx context.Context, inc *model.DynamicSupervisionDetailIncInput, set *model.DynamicSupervisionDetailSetInput, Id int64) (*model1.DynamicSupervisionDetail, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.DynamicSupervisionDetail{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.DynamicSupervisionDetail
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) DynamicSupervisionDetail(ctx context.Context, distinctOn []model.DynamicSupervisionDetailSelectColumn, limit *int, offset *int, orderBy []*model.DynamicSupervisionDetailOrderBy, where *model.DynamicSupervisionDetailBoolExp) ([]*model1.DynamicSupervisionDetail, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DynamicSupervisionDetail{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.DynamicSupervisionDetail
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) DynamicSupervisionDetailAggregate(ctx context.Context, distinctOn []model.DynamicSupervisionDetailSelectColumn, limit *int, offset *int, orderBy []*model.DynamicSupervisionDetailOrderBy, where *model.DynamicSupervisionDetailBoolExp) (*model.DynamicSupervisionDetailAggregate, error) {
	var rs model.DynamicSupervisionDetailAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.DynamicSupervisionDetail{})
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

func (r *queryResolver) DynamicSupervisionDetailByPk(ctx context.Context, Id int64) (*model1.DynamicSupervisionDetail, error) {
	var rs model1.DynamicSupervisionDetail
	tx := db.DB.Model(&model1.DynamicSupervisionDetail{}).First(&rs, Id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
