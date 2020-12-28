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

func (r *mutationResolver) DeleteDriverPeccancyCheck(ctx context.Context, where model.DriverPeccancyCheckBoolExp) (*model.DriverPeccancyCheckMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DriverPeccancyCheck{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.DriverPeccancyCheck
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
	return &model.DriverPeccancyCheckMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteDriverPeccancyCheckByPk(ctx context.Context, Id int64) (*model1.DriverPeccancyCheck, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.DriverPeccancyCheck
	tx := db.DB.Model(&model1.DriverPeccancyCheck{})
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

func (r *mutationResolver) InsertDriverPeccancyCheck(ctx context.Context, objects []*model.DriverPeccancyCheckInsertInput) (*model.DriverPeccancyCheckMutationResponse, error) {
	rs := make([]*model1.DriverPeccancyCheck, 0)
	for _, object := range objects {
		v := &model1.DriverPeccancyCheck{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.DriverPeccancyCheck{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.DriverPeccancyCheckMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertDriverPeccancyCheckOne(ctx context.Context, object model.DriverPeccancyCheckInsertInput) (*model1.DriverPeccancyCheck, error) {
	rs := &model1.DriverPeccancyCheck{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.DriverPeccancyCheck{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateDriverPeccancyCheck(ctx context.Context, inc *model.DriverPeccancyCheckIncInput, set *model.DriverPeccancyCheckSetInput, where model.DriverPeccancyCheckBoolExp) (*model.DriverPeccancyCheckMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DriverPeccancyCheck{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.DriverPeccancyCheckMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.DriverPeccancyCheckMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateDriverPeccancyCheckByPk(ctx context.Context, inc *model.DriverPeccancyCheckIncInput, set *model.DriverPeccancyCheckSetInput, Id int64) (*model1.DriverPeccancyCheck, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.DriverPeccancyCheck{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.DriverPeccancyCheck
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) DriverPeccancyCheck(ctx context.Context, distinctOn []model.DriverPeccancyCheckSelectColumn, limit *int, offset *int, orderBy []*model.DriverPeccancyCheckOrderBy, where *model.DriverPeccancyCheckBoolExp) ([]*model1.DriverPeccancyCheck, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DriverPeccancyCheck{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.DriverPeccancyCheck
	tx = tx.Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) DriverPeccancyCheckAggregate(ctx context.Context, distinctOn []model.DriverPeccancyCheckSelectColumn, limit *int, offset *int, orderBy []*model.DriverPeccancyCheckOrderBy, where *model.DriverPeccancyCheckBoolExp) (*model.DriverPeccancyCheckAggregate, error) {
	var rs model.DriverPeccancyCheckAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.DriverPeccancyCheck{})
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

func (r *queryResolver) DriverPeccancyCheckByPk(ctx context.Context, Id int64) (*model1.DriverPeccancyCheck, error) {
	var rs model1.DriverPeccancyCheck
	tx := db.DB.Model(&model1.DriverPeccancyCheck{}).First(&rs, Id)
	err := tx.Error
	return &rs, err
}
