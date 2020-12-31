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

func (r *mutationResolver) DeleteEcdFileVehicle(ctx context.Context, where model.EcdFileVehicleBoolExp) (*model.EcdFileVehicleMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EcdFileVehicle{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.EcdFileVehicle
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
	return &model.EcdFileVehicleMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteEcdFileVehicleByPk(ctx context.Context, Id int64) (*model1.EcdFileVehicle, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.EcdFileVehicle
	tx := db.DB.Model(&model1.EcdFileVehicle{})
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

func (r *mutationResolver) InsertEcdFileVehicle(ctx context.Context, objects []*model.EcdFileVehicleInsertInput) (*model.EcdFileVehicleMutationResponse, error) {
	rs := make([]*model1.EcdFileVehicle, 0)
	for _, object := range objects {
		v := &model1.EcdFileVehicle{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.EcdFileVehicle{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.EcdFileVehicleMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertEcdFileVehicleOne(ctx context.Context, object model.EcdFileVehicleInsertInput) (*model1.EcdFileVehicle, error) {
	rs := &model1.EcdFileVehicle{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.EcdFileVehicle{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateEcdFileVehicle(ctx context.Context, inc *model.EcdFileVehicleIncInput, set *model.EcdFileVehicleSetInput, where model.EcdFileVehicleBoolExp) (*model.EcdFileVehicleMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EcdFileVehicle{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.EcdFileVehicleMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.EcdFileVehicle
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
	return &model.EcdFileVehicleMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateEcdFileVehicleByPk(ctx context.Context, inc *model.EcdFileVehicleIncInput, set *model.EcdFileVehicleSetInput, Id int64) (*model1.EcdFileVehicle, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.EcdFileVehicle{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.EcdFileVehicle
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) EcdFileVehicle(ctx context.Context, distinctOn []model.EcdFileVehicleSelectColumn, limit *int, offset *int, orderBy []*model.EcdFileVehicleOrderBy, where *model.EcdFileVehicleBoolExp) ([]*model1.EcdFileVehicle, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EcdFileVehicle{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.EcdFileVehicle
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) EcdFileVehicleAggregate(ctx context.Context, distinctOn []model.EcdFileVehicleSelectColumn, limit *int, offset *int, orderBy []*model.EcdFileVehicleOrderBy, where *model.EcdFileVehicleBoolExp) (*model.EcdFileVehicleAggregate, error) {
	var rs model.EcdFileVehicleAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.EcdFileVehicle{})
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

func (r *queryResolver) EcdFileVehicleByPk(ctx context.Context, Id int64) (*model1.EcdFileVehicle, error) {
	var rs model1.EcdFileVehicle
	tx := db.DB.Model(&model1.EcdFileVehicle{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}