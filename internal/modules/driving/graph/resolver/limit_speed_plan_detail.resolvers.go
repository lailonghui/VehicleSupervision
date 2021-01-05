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

func (r *mutationResolver) DeleteLimitSpeedPlanDetail(ctx context.Context, where model.LimitSpeedPlanDetailBoolExp) (*model.LimitSpeedPlanDetailMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.LimitSpeedPlanDetail{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.LimitSpeedPlanDetail
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
	return &model.LimitSpeedPlanDetailMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteLimitSpeedPlanDetailByPk(ctx context.Context, Id int64) (*model1.LimitSpeedPlanDetail, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.LimitSpeedPlanDetail
	tx := db.DB.Model(&model1.LimitSpeedPlanDetail{})
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

func (r *mutationResolver) DeleteLimitSpeedPlanDetailByUnionPk(ctx context.Context, unionId string) (*model1.LimitSpeedPlanDetail, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.LimitSpeedPlanDetail
	tx := db.DB.Model(&model1.LimitSpeedPlanDetail{})
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

func (r *mutationResolver) InsertLimitSpeedPlanDetail(ctx context.Context, objects []*model.LimitSpeedPlanDetailInsertInput) (*model.LimitSpeedPlanDetailMutationResponse, error) {
	rs := make([]*model1.LimitSpeedPlanDetail, 0)
	for _, object := range objects {
		v := &model1.LimitSpeedPlanDetail{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.LimitSpeedPlanDetail{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.LimitSpeedPlanDetailMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertLimitSpeedPlanDetailOne(ctx context.Context, object model.LimitSpeedPlanDetailInsertInput) (*model1.LimitSpeedPlanDetail, error) {
	rs := &model1.LimitSpeedPlanDetail{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.LimitSpeedPlanDetail{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateLimitSpeedPlanDetail(ctx context.Context, inc *model.LimitSpeedPlanDetailIncInput, set *model.LimitSpeedPlanDetailSetInput, where model.LimitSpeedPlanDetailBoolExp) (*model.LimitSpeedPlanDetailMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.LimitSpeedPlanDetail{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.LimitSpeedPlanDetailMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.LimitSpeedPlanDetail
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
	return &model.LimitSpeedPlanDetailMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateLimitSpeedPlanDetailByPk(ctx context.Context, inc *model.LimitSpeedPlanDetailIncInput, set *model.LimitSpeedPlanDetailSetInput, Id int64) (*model1.LimitSpeedPlanDetail, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.LimitSpeedPlanDetail{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.LimitSpeedPlanDetail
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *mutationResolver) UpdateLimitSpeedPlanDetailByUnionPk(ctx context.Context, inc *model.LimitSpeedPlanDetailIncInput, set *model.LimitSpeedPlanDetailSetInput, unionId string) (*model1.LimitSpeedPlanDetail, error) {
	var rs model1.LimitSpeedPlanDetail
	tx := db.DB.Where(rs.UnionPrimaryColumnName()+" = ?", unionId)
	qt := util.NewQueryTranslator(tx, &model1.LimitSpeedPlanDetail{})
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

func (r *queryResolver) LimitSpeedPlanDetail(ctx context.Context, distinctOn []model.LimitSpeedPlanDetailSelectColumn, limit *int, offset *int, orderBy []*model.LimitSpeedPlanDetailOrderBy, where *model.LimitSpeedPlanDetailBoolExp) ([]*model1.LimitSpeedPlanDetail, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.LimitSpeedPlanDetail{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.LimitSpeedPlanDetail
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) LimitSpeedPlanDetailAggregate(ctx context.Context, distinctOn []model.LimitSpeedPlanDetailSelectColumn, limit *int, offset *int, orderBy []*model.LimitSpeedPlanDetailOrderBy, where *model.LimitSpeedPlanDetailBoolExp) (*model.LimitSpeedPlanDetailAggregate, error) {
	var rs model.LimitSpeedPlanDetailAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.LimitSpeedPlanDetail{})
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

func (r *queryResolver) LimitSpeedPlanDetailByPk(ctx context.Context, Id int64) (*model1.LimitSpeedPlanDetail, error) {
	var rs model1.LimitSpeedPlanDetail
	tx := db.DB.Model(&model1.LimitSpeedPlanDetail{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}

func (r *queryResolver) LimitSpeedPlanDetailByUnionPk(ctx context.Context, unionId string) (*model1.LimitSpeedPlanDetail, error) {
	var rs model1.LimitSpeedPlanDetail
	tx := db.DB.Model(&model1.LimitSpeedPlanDetail{}).Select(util.GetTopPreloads(ctx)).Where(rs.UnionPrimaryColumnName()+" = ?", unionId).First(&rs)

	err := tx.Error
	return &rs, err
}
