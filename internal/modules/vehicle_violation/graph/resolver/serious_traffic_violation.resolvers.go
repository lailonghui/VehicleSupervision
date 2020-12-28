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

func (r *mutationResolver) DeleteSeriousTrafficViolation(ctx context.Context, where model.SeriousTrafficViolationBoolExp) (*model.SeriousTrafficViolationMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.SeriousTrafficViolation{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.SeriousTrafficViolation
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
	return &model.SeriousTrafficViolationMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteSeriousTrafficViolationByPk(ctx context.Context, Id int64) (*model1.SeriousTrafficViolation, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.SeriousTrafficViolation
	tx := db.DB.Model(&model1.SeriousTrafficViolation{})
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

func (r *mutationResolver) InsertSeriousTrafficViolation(ctx context.Context, objects []*model.SeriousTrafficViolationInsertInput) (*model.SeriousTrafficViolationMutationResponse, error) {
	rs := []*model1.SeriousTrafficViolation{}
	for _, object := range objects {
		v := &model1.SeriousTrafficViolation{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.SeriousTrafficViolation{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.SeriousTrafficViolationMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertSeriousTrafficViolationOne(ctx context.Context, object model.SeriousTrafficViolationInsertInput) (*model1.SeriousTrafficViolation, error) {
	rs := &model1.SeriousTrafficViolation{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.SeriousTrafficViolation{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateSeriousTrafficViolation(ctx context.Context, inc *model.SeriousTrafficViolationIncInput, set *model.SeriousTrafficViolationSetInput, where model.SeriousTrafficViolationBoolExp) (*model.SeriousTrafficViolationMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.SeriousTrafficViolation{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.SeriousTrafficViolationMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.SeriousTrafficViolationMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateSeriousTrafficViolationByPk(ctx context.Context, inc *model.SeriousTrafficViolationIncInput, set *model.SeriousTrafficViolationSetInput, Id int64) (*model1.SeriousTrafficViolation, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.SeriousTrafficViolation{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.SeriousTrafficViolation
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) SeriousTrafficViolation(ctx context.Context, distinctOn []model.SeriousTrafficViolationSelectColumn, limit *int, offset *int, orderBy []*model.SeriousTrafficViolationOrderBy, where *model.SeriousTrafficViolationBoolExp) ([]*model1.SeriousTrafficViolation, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.SeriousTrafficViolation{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.SeriousTrafficViolation
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) SeriousTrafficViolationAggregate(ctx context.Context, distinctOn []model.SeriousTrafficViolationSelectColumn, limit *int, offset *int, orderBy []*model.SeriousTrafficViolationOrderBy, where *model.SeriousTrafficViolationBoolExp) (*model.SeriousTrafficViolationAggregate, error) {
	var rs model.SeriousTrafficViolationAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.SeriousTrafficViolation{})
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

func (r *queryResolver) SeriousTrafficViolationByPk(ctx context.Context, Id int64) (*model1.SeriousTrafficViolation, error) {
	var rs model1.SeriousTrafficViolation
	tx := db.DB.Model(&model1.SeriousTrafficViolation{}).First(&rs, Id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
