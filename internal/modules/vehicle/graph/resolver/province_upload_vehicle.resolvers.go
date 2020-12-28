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

func (r *mutationResolver) DeleteProvinceUploadVehicle(ctx context.Context, where model.ProvinceUploadVehicleBoolExp) (*model.ProvinceUploadVehicleMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.ProvinceUploadVehicle{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.ProvinceUploadVehicle
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
	return &model.ProvinceUploadVehicleMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteProvinceUploadVehicleByPk(ctx context.Context, Id int64) (*model1.ProvinceUploadVehicle, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.ProvinceUploadVehicle
	tx := db.DB.Model(&model1.ProvinceUploadVehicle{})
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

func (r *mutationResolver) InsertProvinceUploadVehicle(ctx context.Context, objects []*model.ProvinceUploadVehicleInsertInput) (*model.ProvinceUploadVehicleMutationResponse, error) {
	rs := []*model1.ProvinceUploadVehicle{}
	for _, object := range objects {
		v := &model1.ProvinceUploadVehicle{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.ProvinceUploadVehicle{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.ProvinceUploadVehicleMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertProvinceUploadVehicleOne(ctx context.Context, object model.ProvinceUploadVehicleInsertInput) (*model1.ProvinceUploadVehicle, error) {
	rs := &model1.ProvinceUploadVehicle{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.ProvinceUploadVehicle{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateProvinceUploadVehicle(ctx context.Context, inc *model.ProvinceUploadVehicleIncInput, set *model.ProvinceUploadVehicleSetInput, where model.ProvinceUploadVehicleBoolExp) (*model.ProvinceUploadVehicleMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.ProvinceUploadVehicle{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.ProvinceUploadVehicleMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.ProvinceUploadVehicleMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateProvinceUploadVehicleByPk(ctx context.Context, inc *model.ProvinceUploadVehicleIncInput, set *model.ProvinceUploadVehicleSetInput, Id int64) (*model1.ProvinceUploadVehicle, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.ProvinceUploadVehicle{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.ProvinceUploadVehicle
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) ProvinceUploadVehicle(ctx context.Context, distinctOn []model.ProvinceUploadVehicleSelectColumn, limit *int, offset *int, orderBy []*model.ProvinceUploadVehicleOrderBy, where *model.ProvinceUploadVehicleBoolExp) ([]*model1.ProvinceUploadVehicle, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.ProvinceUploadVehicle{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.ProvinceUploadVehicle
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) ProvinceUploadVehicleAggregate(ctx context.Context, distinctOn []model.ProvinceUploadVehicleSelectColumn, limit *int, offset *int, orderBy []*model.ProvinceUploadVehicleOrderBy, where *model.ProvinceUploadVehicleBoolExp) (*model.ProvinceUploadVehicleAggregate, error) {
	var rs model.ProvinceUploadVehicleAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.ProvinceUploadVehicle{})
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

func (r *queryResolver) ProvinceUploadVehicleByPk(ctx context.Context, Id int64) (*model1.ProvinceUploadVehicle, error) {
	var rs model1.ProvinceUploadVehicle
	tx := db.DB.Model(&model1.ProvinceUploadVehicle{}).First(&rs, Id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
