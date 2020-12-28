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

func (r *mutationResolver) DeleteVehicleSecurityCheckRecord(ctx context.Context, where model.VehicleSecurityCheckRecordBoolExp) (*model.VehicleSecurityCheckRecordMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleSecurityCheckRecord{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.VehicleSecurityCheckRecord
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
	return &model.VehicleSecurityCheckRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteVehicleSecurityCheckRecordByPk(ctx context.Context, Id int64) (*model1.VehicleSecurityCheckRecord, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.VehicleSecurityCheckRecord
	tx := db.DB.Model(&model1.VehicleSecurityCheckRecord{})
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

func (r *mutationResolver) InsertVehicleSecurityCheckRecord(ctx context.Context, objects []*model.VehicleSecurityCheckRecordInsertInput) (*model.VehicleSecurityCheckRecordMutationResponse, error) {
	rs := []*model1.VehicleSecurityCheckRecord{}
	for _, object := range objects {
		v := &model1.VehicleSecurityCheckRecord{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.VehicleSecurityCheckRecord{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.VehicleSecurityCheckRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertVehicleSecurityCheckRecordOne(ctx context.Context, object model.VehicleSecurityCheckRecordInsertInput) (*model1.VehicleSecurityCheckRecord, error) {
	rs := &model1.VehicleSecurityCheckRecord{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.VehicleSecurityCheckRecord{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateVehicleSecurityCheckRecord(ctx context.Context, inc *model.VehicleSecurityCheckRecordIncInput, set *model.VehicleSecurityCheckRecordSetInput, where model.VehicleSecurityCheckRecordBoolExp) (*model.VehicleSecurityCheckRecordMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleSecurityCheckRecord{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.VehicleSecurityCheckRecordMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.VehicleSecurityCheckRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateVehicleSecurityCheckRecordByPk(ctx context.Context, inc *model.VehicleSecurityCheckRecordIncInput, set *model.VehicleSecurityCheckRecordSetInput, Id int64) (*model1.VehicleSecurityCheckRecord, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.VehicleSecurityCheckRecord{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.VehicleSecurityCheckRecord
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) VehicleSecurityCheckRecord(ctx context.Context, distinctOn []model.VehicleSecurityCheckRecordSelectColumn, limit *int, offset *int, orderBy []*model.VehicleSecurityCheckRecordOrderBy, where *model.VehicleSecurityCheckRecordBoolExp) ([]*model1.VehicleSecurityCheckRecord, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleSecurityCheckRecord{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.VehicleSecurityCheckRecord
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) VehicleSecurityCheckRecordAggregate(ctx context.Context, distinctOn []model.VehicleSecurityCheckRecordSelectColumn, limit *int, offset *int, orderBy []*model.VehicleSecurityCheckRecordOrderBy, where *model.VehicleSecurityCheckRecordBoolExp) (*model.VehicleSecurityCheckRecordAggregate, error) {
	var rs model.VehicleSecurityCheckRecordAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.VehicleSecurityCheckRecord{})
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

func (r *queryResolver) VehicleSecurityCheckRecordByPk(ctx context.Context, Id int64) (*model1.VehicleSecurityCheckRecord, error) {
	var rs model1.VehicleSecurityCheckRecord
	tx := db.DB.Model(&model1.VehicleSecurityCheckRecord{}).First(&rs, Id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
