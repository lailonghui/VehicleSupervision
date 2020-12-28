package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/vehicle_snapshot_system/graph/model"
	model1 "VehicleSupervision/internal/modules/vehicle_snapshot_system/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteOfflineRegistrationRecord(ctx context.Context, where model.OfflineRegistrationRecordBoolExp) (*model.OfflineRegistrationRecordMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.OfflineRegistrationRecord{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.OfflineRegistrationRecord
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
	return &model.OfflineRegistrationRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteOfflineRegistrationRecordByPk(ctx context.Context, Id int64) (*model1.OfflineRegistrationRecord, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.OfflineRegistrationRecord
	tx := db.DB.Model(&model1.OfflineRegistrationRecord{})
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

func (r *mutationResolver) InsertOfflineRegistrationRecord(ctx context.Context, objects []*model.OfflineRegistrationRecordInsertInput) (*model.OfflineRegistrationRecordMutationResponse, error) {
	rs := []*model1.OfflineRegistrationRecord{}
	for _, object := range objects {
		v := &model1.OfflineRegistrationRecord{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.OfflineRegistrationRecord{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.OfflineRegistrationRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertOfflineRegistrationRecordOne(ctx context.Context, object model.OfflineRegistrationRecordInsertInput) (*model1.OfflineRegistrationRecord, error) {
	rs := &model1.OfflineRegistrationRecord{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.OfflineRegistrationRecord{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateOfflineRegistrationRecord(ctx context.Context, inc *model.OfflineRegistrationRecordIncInput, set *model.OfflineRegistrationRecordSetInput, where model.OfflineRegistrationRecordBoolExp) (*model.OfflineRegistrationRecordMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.OfflineRegistrationRecord{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.OfflineRegistrationRecordMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.OfflineRegistrationRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateOfflineRegistrationRecordByPk(ctx context.Context, inc *model.OfflineRegistrationRecordIncInput, set *model.OfflineRegistrationRecordSetInput, Id int64) (*model1.OfflineRegistrationRecord, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.OfflineRegistrationRecord{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.OfflineRegistrationRecord
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) OfflineRegistrationRecord(ctx context.Context, distinctOn []model.OfflineRegistrationRecordSelectColumn, limit *int, offset *int, orderBy []*model.OfflineRegistrationRecordOrderBy, where *model.OfflineRegistrationRecordBoolExp) ([]*model1.OfflineRegistrationRecord, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.OfflineRegistrationRecord{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.OfflineRegistrationRecord
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) OfflineRegistrationRecordAggregate(ctx context.Context, distinctOn []model.OfflineRegistrationRecordSelectColumn, limit *int, offset *int, orderBy []*model.OfflineRegistrationRecordOrderBy, where *model.OfflineRegistrationRecordBoolExp) (*model.OfflineRegistrationRecordAggregate, error) {
	var rs model.OfflineRegistrationRecordAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.OfflineRegistrationRecord{})
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

func (r *queryResolver) OfflineRegistrationRecordByPk(ctx context.Context, Id int64) (*model1.OfflineRegistrationRecord, error) {
	var rs model1.OfflineRegistrationRecord
	tx := db.DB.Model(&model1.OfflineRegistrationRecord{}).First(&rs, Id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
