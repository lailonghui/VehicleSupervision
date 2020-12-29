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

func (r *mutationResolver) DeleteElectricFenceEnteranceRecord(ctx context.Context, where model.ElectricFenceEnteranceRecordBoolExp) (*model.ElectricFenceEnteranceRecordMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.ElectricFenceEnteranceRecord{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.ElectricFenceEnteranceRecord
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
	return &model.ElectricFenceEnteranceRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteElectricFenceEnteranceRecordByPk(ctx context.Context, Id int64) (*model1.ElectricFenceEnteranceRecord, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.ElectricFenceEnteranceRecord
	tx := db.DB.Model(&model1.ElectricFenceEnteranceRecord{})
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

func (r *mutationResolver) InsertElectricFenceEnteranceRecord(ctx context.Context, objects []*model.ElectricFenceEnteranceRecordInsertInput) (*model.ElectricFenceEnteranceRecordMutationResponse, error) {
	rs := make([]*model1.ElectricFenceEnteranceRecord, 0)
	for _, object := range objects {
		v := &model1.ElectricFenceEnteranceRecord{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.ElectricFenceEnteranceRecord{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.ElectricFenceEnteranceRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertElectricFenceEnteranceRecordOne(ctx context.Context, object model.ElectricFenceEnteranceRecordInsertInput) (*model1.ElectricFenceEnteranceRecord, error) {
	rs := &model1.ElectricFenceEnteranceRecord{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.ElectricFenceEnteranceRecord{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateElectricFenceEnteranceRecord(ctx context.Context, inc *model.ElectricFenceEnteranceRecordIncInput, set *model.ElectricFenceEnteranceRecordSetInput, where model.ElectricFenceEnteranceRecordBoolExp) (*model.ElectricFenceEnteranceRecordMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.ElectricFenceEnteranceRecord{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.ElectricFenceEnteranceRecordMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.ElectricFenceEnteranceRecord
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
	return &model.ElectricFenceEnteranceRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateElectricFenceEnteranceRecordByPk(ctx context.Context, inc *model.ElectricFenceEnteranceRecordIncInput, set *model.ElectricFenceEnteranceRecordSetInput, Id int64) (*model1.ElectricFenceEnteranceRecord, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.ElectricFenceEnteranceRecord{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.ElectricFenceEnteranceRecord
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) ElectricFenceEnteranceRecord(ctx context.Context, distinctOn []model.ElectricFenceEnteranceRecordSelectColumn, limit *int, offset *int, orderBy []*model.ElectricFenceEnteranceRecordOrderBy, where *model.ElectricFenceEnteranceRecordBoolExp) ([]*model1.ElectricFenceEnteranceRecord, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.ElectricFenceEnteranceRecord{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.ElectricFenceEnteranceRecord
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) ElectricFenceEnteranceRecordAggregate(ctx context.Context, distinctOn []model.ElectricFenceEnteranceRecordSelectColumn, limit *int, offset *int, orderBy []*model.ElectricFenceEnteranceRecordOrderBy, where *model.ElectricFenceEnteranceRecordBoolExp) (*model.ElectricFenceEnteranceRecordAggregate, error) {
	var rs model.ElectricFenceEnteranceRecordAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.ElectricFenceEnteranceRecord{})
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

func (r *queryResolver) ElectricFenceEnteranceRecordByPk(ctx context.Context, Id int64) (*model1.ElectricFenceEnteranceRecord, error) {
	var rs model1.ElectricFenceEnteranceRecord
	tx := db.DB.Model(&model1.ElectricFenceEnteranceRecord{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}
