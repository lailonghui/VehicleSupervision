package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/vehicle/graph/model"
	model1 "VehicleSupervision/internal/modules/vehicle/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteJjVehicle(ctx context.Context, where model.JjVehicleBoolExp) (*model.JjVehicleMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.JjVehicle{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.JjVehicle
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
	return &model.JjVehicleMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteJjVehicleByPk(ctx context.Context, Id int64) (*model1.JjVehicle, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.JjVehicle
	tx := db.DB.Model(&model1.JjVehicle{})
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

func (r *mutationResolver) InsertJjVehicle(ctx context.Context, objects []*model.JjVehicleInsertInput) (*model.JjVehicleMutationResponse, error) {
	rs := []*model1.JjVehicle{}
	for _, object := range objects {
		v := &model1.JjVehicle{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.JjVehicle{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.JjVehicleMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertJjVehicleOne(ctx context.Context, object model.JjVehicleInsertInput) (*model1.JjVehicle, error) {
	rs := &model1.JjVehicle{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.JjVehicle{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateJjVehicle(ctx context.Context, inc *model.JjVehicleIncInput, set *model.JjVehicleSetInput, where model.JjVehicleBoolExp) (*model.JjVehicleMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.JjVehicle{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.JjVehicleMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.JjVehicleMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateJjVehicleByPk(ctx context.Context, inc *model.JjVehicleIncInput, set *model.JjVehicleSetInput, Id int64) (*model1.JjVehicle, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.JjVehicle{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.JjVehicle
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) JjVehicle(ctx context.Context, distinctOn []model.JjVehicleSelectColumn, limit *int, offset *int, orderBy []*model.JjVehicleOrderBy, where *model.JjVehicleBoolExp) ([]*model1.JjVehicle, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.JjVehicle{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.JjVehicle
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) JjVehicleAggregate(ctx context.Context, distinctOn []model.JjVehicleSelectColumn, limit *int, offset *int, orderBy []*model.JjVehicleOrderBy, where *model.JjVehicleBoolExp) (*model.JjVehicleAggregate, error) {
	var rs model.JjVehicleAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.JjVehicle{})
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

func (r *queryResolver) JjVehicleByPk(ctx context.Context, Id int64) (*model1.JjVehicle, error) {
	var rs model1.JjVehicle
	tx := db.DB.Model(&model1.JjVehicle{}).First(&rs, Id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
