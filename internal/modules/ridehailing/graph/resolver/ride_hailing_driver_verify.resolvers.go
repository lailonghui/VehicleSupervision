package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/ridehailing/graph/model"
	model1 "VehicleSupervision/internal/modules/ridehailing/model"
	"VehicleSupervision/pkg/graphql/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteRideHailingDriverVerify(ctx context.Context, where model.RideHailingDriverVerifyBoolExp) (*model.RideHailingDriverVerifyMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.RideHailingDriverVerify{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.RideHailingDriverVerify
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
	return &model.RideHailingDriverVerifyMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteRideHailingDriverVerifyByPk(ctx context.Context, id int64) (*model1.RideHailingDriverVerify, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.RideHailingDriverVerify
	tx := db.DB.Model(&model1.RideHailingDriverVerify{})
	if len(preloads) > 0 {
		// 如果请求的字段不为空，则先查询一遍数据库
		tx = tx.Select(preloads).Where("id = ?", id).First(&rs)
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

func (r *mutationResolver) InsertRideHailingDriverVerify(ctx context.Context, objects []*model.RideHailingDriverVerifyInsertInput, onConflict *model.RideHailingDriverVerifyOnConflict) (*model.RideHailingDriverVerifyMutationResponse, error) {
	rs := rideHailingDriverVerifyInsertInputBatchConvert(objects)
	tx := db.DB.Model(&model1.RideHailingDriverVerify{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.RideHailingDriverVerifyMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertRideHailingDriverVerifyOne(ctx context.Context, object model.RideHailingDriverVerifyInsertInput, onConflict *model.RideHailingDriverVerifyOnConflict) (*model1.RideHailingDriverVerify, error) {
	rs := rideHailingDriverVerifyInsertInputConvert(&object)
	tx := db.DB.Model(&model1.RideHailingDriverVerify{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateRideHailingDriverVerify(ctx context.Context, inc *model.RideHailingDriverVerifyIncInput, set *model.RideHailingDriverVerifySetInput, where model.RideHailingDriverVerifyBoolExp) (*model.RideHailingDriverVerifyMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.RideHailingDriverVerify{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.RideHailingDriverVerifyMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.RideHailingDriverVerifyMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateRideHailingDriverVerifyByPk(ctx context.Context, inc *model.RideHailingDriverVerifyIncInput, set *model.RideHailingDriverVerifySetInput, pkColumns model.RideHailingDriverVerifyPkColumnsInput) (*model1.RideHailingDriverVerify, error) {
	tx := db.DB.Where("id = ?", pkColumns.ID)
	qt := util.NewQueryTranslator(tx, &model1.RideHailingDriverVerify{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.RideHailingDriverVerify
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) RideHailingDriverVerify(ctx context.Context, distinctOn []model.RideHailingDriverVerifySelectColumn, limit *int, offset *int, orderBy []*model.RideHailingDriverVerifyOrderBy, where *model.RideHailingDriverVerifyBoolExp) ([]*model1.RideHailingDriverVerify, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.RideHailingDriverVerify{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.RideHailingDriverVerify
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) RideHailingDriverVerifyAggregate(ctx context.Context, distinctOn []model.RideHailingDriverVerifySelectColumn, limit *int, offset *int, orderBy []*model.RideHailingDriverVerifyOrderBy, where *model.RideHailingDriverVerifyBoolExp) (*model.RideHailingDriverVerifyAggregate, error) {
	var rs model.RideHailingDriverVerifyAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.RideHailingDriverVerify{})
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

func (r *queryResolver) RideHailingDriverVerifyByPk(ctx context.Context, id int64) (*model1.RideHailingDriverVerify, error) {
	var rs model1.RideHailingDriverVerify
	tx := db.DB.Model(&model1.RideHailingDriverVerify{}).First(&rs, id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
