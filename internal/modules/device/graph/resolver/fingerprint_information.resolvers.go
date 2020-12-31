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

func (r *mutationResolver) DeleteFingerprintInformation(ctx context.Context, where model.FingerprintInformationBoolExp) (*model.FingerprintInformationMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.FingerprintInformation{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.FingerprintInformation
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
	return &model.FingerprintInformationMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteFingerprintInformationByPk(ctx context.Context, Id int64) (*model1.FingerprintInformation, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.FingerprintInformation
	tx := db.DB.Model(&model1.FingerprintInformation{})
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

func (r *mutationResolver) InsertFingerprintInformation(ctx context.Context, objects []*model.FingerprintInformationInsertInput) (*model.FingerprintInformationMutationResponse, error) {
	rs := make([]*model1.FingerprintInformation, 0)
	for _, object := range objects {
		v := &model1.FingerprintInformation{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.FingerprintInformation{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.FingerprintInformationMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertFingerprintInformationOne(ctx context.Context, object model.FingerprintInformationInsertInput) (*model1.FingerprintInformation, error) {
	rs := &model1.FingerprintInformation{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.FingerprintInformation{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateFingerprintInformation(ctx context.Context, inc *model.FingerprintInformationIncInput, set *model.FingerprintInformationSetInput, where model.FingerprintInformationBoolExp) (*model.FingerprintInformationMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.FingerprintInformation{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.FingerprintInformationMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.FingerprintInformation
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
	return &model.FingerprintInformationMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateFingerprintInformationByPk(ctx context.Context, inc *model.FingerprintInformationIncInput, set *model.FingerprintInformationSetInput, Id int64) (*model1.FingerprintInformation, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.FingerprintInformation{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.FingerprintInformation
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) FingerprintInformation(ctx context.Context, distinctOn []model.FingerprintInformationSelectColumn, limit *int, offset *int, orderBy []*model.FingerprintInformationOrderBy, where *model.FingerprintInformationBoolExp) ([]*model1.FingerprintInformation, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.FingerprintInformation{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.FingerprintInformation
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) FingerprintInformationAggregate(ctx context.Context, distinctOn []model.FingerprintInformationSelectColumn, limit *int, offset *int, orderBy []*model.FingerprintInformationOrderBy, where *model.FingerprintInformationBoolExp) (*model.FingerprintInformationAggregate, error) {
	var rs model.FingerprintInformationAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.FingerprintInformation{})
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

func (r *queryResolver) FingerprintInformationByPk(ctx context.Context, Id int64) (*model1.FingerprintInformation, error) {
	var rs model1.FingerprintInformation
	tx := db.DB.Model(&model1.FingerprintInformation{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}