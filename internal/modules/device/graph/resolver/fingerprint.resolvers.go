package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/device/graph/model"
	model1 "VehicleSupervision/internal/modules/device/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteFingerprint(ctx context.Context, where model.FingerprintBoolExp) (*model.FingerprintMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.Fingerprint{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.Fingerprint
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
	return &model.FingerprintMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteFingerprintByPk(ctx context.Context, Id int64) (*model1.Fingerprint, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.Fingerprint
	tx := db.DB.Model(&model1.Fingerprint{})
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

func (r *mutationResolver) InsertFingerprint(ctx context.Context, objects []*model.FingerprintInsertInput) (*model.FingerprintMutationResponse, error) {
	rs := make([]*model1.Fingerprint, 0)
	for _, object := range objects {
		v := &model1.Fingerprint{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.Fingerprint{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.FingerprintMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertFingerprintOne(ctx context.Context, object model.FingerprintInsertInput) (*model1.Fingerprint, error) {
	rs := &model1.Fingerprint{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.Fingerprint{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateFingerprint(ctx context.Context, inc *model.FingerprintIncInput, set *model.FingerprintSetInput, where model.FingerprintBoolExp) (*model.FingerprintMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.Fingerprint{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.FingerprintMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.Fingerprint
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
	return &model.FingerprintMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateFingerprintByPk(ctx context.Context, inc *model.FingerprintIncInput, set *model.FingerprintSetInput, Id int64) (*model1.Fingerprint, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.Fingerprint{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.Fingerprint
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) Fingerprint(ctx context.Context, distinctOn []model.FingerprintSelectColumn, limit *int, offset *int, orderBy []*model.FingerprintOrderBy, where *model.FingerprintBoolExp) ([]*model1.Fingerprint, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.Fingerprint{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.Fingerprint
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) FingerprintAggregate(ctx context.Context, distinctOn []model.FingerprintSelectColumn, limit *int, offset *int, orderBy []*model.FingerprintOrderBy, where *model.FingerprintBoolExp) (*model.FingerprintAggregate, error) {
	var rs model.FingerprintAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.Fingerprint{})
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

func (r *queryResolver) FingerprintByPk(ctx context.Context, Id int64) (*model1.Fingerprint, error) {
	var rs model1.Fingerprint
	tx := db.DB.Model(&model1.Fingerprint{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}
