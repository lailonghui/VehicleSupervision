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

func (r *mutationResolver) DeleteVehicleAlarmSupervision(ctx context.Context, where model.VehicleAlarmSupervisionBoolExp) (*model.VehicleAlarmSupervisionMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleAlarmSupervision{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.VehicleAlarmSupervision
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
	return &model.VehicleAlarmSupervisionMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteVehicleAlarmSupervisionByPk(ctx context.Context, Id int64) (*model1.VehicleAlarmSupervision, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.VehicleAlarmSupervision
	tx := db.DB.Model(&model1.VehicleAlarmSupervision{})
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

func (r *mutationResolver) InsertVehicleAlarmSupervision(ctx context.Context, objects []*model.VehicleAlarmSupervisionInsertInput) (*model.VehicleAlarmSupervisionMutationResponse, error) {
	rs := []*model1.VehicleAlarmSupervision{}
	for _, object := range objects {
		v := &model1.VehicleAlarmSupervision{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.VehicleAlarmSupervision{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.VehicleAlarmSupervisionMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertVehicleAlarmSupervisionOne(ctx context.Context, object model.VehicleAlarmSupervisionInsertInput) (*model1.VehicleAlarmSupervision, error) {
	rs := &model1.VehicleAlarmSupervision{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.VehicleAlarmSupervision{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateVehicleAlarmSupervision(ctx context.Context, inc *model.VehicleAlarmSupervisionIncInput, set *model.VehicleAlarmSupervisionSetInput, where model.VehicleAlarmSupervisionBoolExp) (*model.VehicleAlarmSupervisionMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleAlarmSupervision{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.VehicleAlarmSupervisionMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.VehicleAlarmSupervisionMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateVehicleAlarmSupervisionByPk(ctx context.Context, inc *model.VehicleAlarmSupervisionIncInput, set *model.VehicleAlarmSupervisionSetInput, Id int64) (*model1.VehicleAlarmSupervision, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.VehicleAlarmSupervision{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.VehicleAlarmSupervision
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) VehicleAlarmSupervision(ctx context.Context, distinctOn []model.VehicleAlarmSupervisionSelectColumn, limit *int, offset *int, orderBy []*model.VehicleAlarmSupervisionOrderBy, where *model.VehicleAlarmSupervisionBoolExp) ([]*model1.VehicleAlarmSupervision, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleAlarmSupervision{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.VehicleAlarmSupervision
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) VehicleAlarmSupervisionAggregate(ctx context.Context, distinctOn []model.VehicleAlarmSupervisionSelectColumn, limit *int, offset *int, orderBy []*model.VehicleAlarmSupervisionOrderBy, where *model.VehicleAlarmSupervisionBoolExp) (*model.VehicleAlarmSupervisionAggregate, error) {
	var rs model.VehicleAlarmSupervisionAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.VehicleAlarmSupervision{})
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

func (r *queryResolver) VehicleAlarmSupervisionByPk(ctx context.Context, Id int64) (*model1.VehicleAlarmSupervision, error) {
	var rs model1.VehicleAlarmSupervision
	tx := db.DB.Model(&model1.VehicleAlarmSupervision{}).First(&rs, Id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
