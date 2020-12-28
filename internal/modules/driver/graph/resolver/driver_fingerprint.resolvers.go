package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/driver/graph/model"
	model1 "VehicleSupervision/internal/modules/driver/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteDriverFingerprint(ctx context.Context, where model.DriverFingerprintBoolExp) (*model.DriverFingerprintMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DriverFingerprint{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.DriverFingerprint
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
	return &model.DriverFingerprintMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteDriverFingerprintByPk(ctx context.Context, Id int64) (*model1.DriverFingerprint, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.DriverFingerprint
	tx := db.DB.Model(&model1.DriverFingerprint{})
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

func (r *mutationResolver) InsertDriverFingerprint(ctx context.Context, objects []*model.DriverFingerprintInsertInput) (*model.DriverFingerprintMutationResponse, error) {
	rs := make([]*model1.DriverFingerprint, 0)
	for _, object := range objects {
		v := &model1.DriverFingerprint{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.DriverFingerprint{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.DriverFingerprintMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertDriverFingerprintOne(ctx context.Context, object model.DriverFingerprintInsertInput) (*model1.DriverFingerprint, error) {
	rs := &model1.DriverFingerprint{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.DriverFingerprint{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateDriverFingerprint(ctx context.Context, inc *model.DriverFingerprintIncInput, set *model.DriverFingerprintSetInput, where model.DriverFingerprintBoolExp) (*model.DriverFingerprintMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DriverFingerprint{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.DriverFingerprintMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.DriverFingerprintMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateDriverFingerprintByPk(ctx context.Context, inc *model.DriverFingerprintIncInput, set *model.DriverFingerprintSetInput, Id int64) (*model1.DriverFingerprint, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.DriverFingerprint{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.DriverFingerprint
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) DriverFingerprint(ctx context.Context, distinctOn []model.DriverFingerprintSelectColumn, limit *int, offset *int, orderBy []*model.DriverFingerprintOrderBy, where *model.DriverFingerprintBoolExp) ([]*model1.DriverFingerprint, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DriverFingerprint{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.DriverFingerprint
	tx = tx.Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) DriverFingerprintAggregate(ctx context.Context, distinctOn []model.DriverFingerprintSelectColumn, limit *int, offset *int, orderBy []*model.DriverFingerprintOrderBy, where *model.DriverFingerprintBoolExp) (*model.DriverFingerprintAggregate, error) {
	var rs model.DriverFingerprintAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.DriverFingerprint{})
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

func (r *queryResolver) DriverFingerprintByPk(ctx context.Context, Id int64) (*model1.DriverFingerprint, error) {
	var rs model1.DriverFingerprint
	tx := db.DB.Model(&model1.DriverFingerprint{}).First(&rs, Id)
	err := tx.Error
	return &rs, err
}
