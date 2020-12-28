package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/muck_truck_recommend_catalog/graph/model"
	model1 "VehicleSupervision/internal/modules/muck_truck_recommend_catalog/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteMuckTruckPurchaseIntention(ctx context.Context, where model.MuckTruckPurchaseIntentionBoolExp) (*model.MuckTruckPurchaseIntentionMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.MuckTruckPurchaseIntention{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.MuckTruckPurchaseIntention
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
	return &model.MuckTruckPurchaseIntentionMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteMuckTruckPurchaseIntentionByPk(ctx context.Context, Id int64) (*model1.MuckTruckPurchaseIntention, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.MuckTruckPurchaseIntention
	tx := db.DB.Model(&model1.MuckTruckPurchaseIntention{})
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

func (r *mutationResolver) InsertMuckTruckPurchaseIntention(ctx context.Context, objects []*model.MuckTruckPurchaseIntentionInsertInput) (*model.MuckTruckPurchaseIntentionMutationResponse, error) {
	rs := make([]*model1.MuckTruckPurchaseIntention, 0)
	for _, object := range objects {
		v := &model1.MuckTruckPurchaseIntention{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.MuckTruckPurchaseIntention{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.MuckTruckPurchaseIntentionMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertMuckTruckPurchaseIntentionOne(ctx context.Context, object model.MuckTruckPurchaseIntentionInsertInput) (*model1.MuckTruckPurchaseIntention, error) {
	rs := &model1.MuckTruckPurchaseIntention{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.MuckTruckPurchaseIntention{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateMuckTruckPurchaseIntention(ctx context.Context, inc *model.MuckTruckPurchaseIntentionIncInput, set *model.MuckTruckPurchaseIntentionSetInput, where model.MuckTruckPurchaseIntentionBoolExp) (*model.MuckTruckPurchaseIntentionMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.MuckTruckPurchaseIntention{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.MuckTruckPurchaseIntentionMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.MuckTruckPurchaseIntentionMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateMuckTruckPurchaseIntentionByPk(ctx context.Context, inc *model.MuckTruckPurchaseIntentionIncInput, set *model.MuckTruckPurchaseIntentionSetInput, Id int64) (*model1.MuckTruckPurchaseIntention, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.MuckTruckPurchaseIntention{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.MuckTruckPurchaseIntention
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) MuckTruckPurchaseIntention(ctx context.Context, distinctOn []model.MuckTruckPurchaseIntentionSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckPurchaseIntentionOrderBy, where *model.MuckTruckPurchaseIntentionBoolExp) ([]*model1.MuckTruckPurchaseIntention, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.MuckTruckPurchaseIntention{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.MuckTruckPurchaseIntention
	tx = tx.Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) MuckTruckPurchaseIntentionAggregate(ctx context.Context, distinctOn []model.MuckTruckPurchaseIntentionSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckPurchaseIntentionOrderBy, where *model.MuckTruckPurchaseIntentionBoolExp) (*model.MuckTruckPurchaseIntentionAggregate, error) {
	var rs model.MuckTruckPurchaseIntentionAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.MuckTruckPurchaseIntention{})
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

func (r *queryResolver) MuckTruckPurchaseIntentionByPk(ctx context.Context, Id int64) (*model1.MuckTruckPurchaseIntention, error) {
	var rs model1.MuckTruckPurchaseIntention
	tx := db.DB.Model(&model1.MuckTruckPurchaseIntention{}).First(&rs, Id)
	err := tx.Error
	return &rs, err
}
