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

func (r *mutationResolver) DeleteSellerRatingRecord(ctx context.Context, where model.SellerRatingRecordBoolExp) (*model.SellerRatingRecordMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.SellerRatingRecord{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.SellerRatingRecord
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
	return &model.SellerRatingRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteSellerRatingRecordByPk(ctx context.Context, Id int64) (*model1.SellerRatingRecord, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.SellerRatingRecord
	tx := db.DB.Model(&model1.SellerRatingRecord{})
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

func (r *mutationResolver) InsertSellerRatingRecord(ctx context.Context, objects []*model.SellerRatingRecordInsertInput) (*model.SellerRatingRecordMutationResponse, error) {
	rs := []*model1.SellerRatingRecord{}
	for _, object := range objects {
		v := &model1.SellerRatingRecord{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.SellerRatingRecord{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.SellerRatingRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertSellerRatingRecordOne(ctx context.Context, object model.SellerRatingRecordInsertInput) (*model1.SellerRatingRecord, error) {
	rs := &model1.SellerRatingRecord{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.SellerRatingRecord{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateSellerRatingRecord(ctx context.Context, inc *model.SellerRatingRecordIncInput, set *model.SellerRatingRecordSetInput, where model.SellerRatingRecordBoolExp) (*model.SellerRatingRecordMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.SellerRatingRecord{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.SellerRatingRecordMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.SellerRatingRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateSellerRatingRecordByPk(ctx context.Context, inc *model.SellerRatingRecordIncInput, set *model.SellerRatingRecordSetInput, Id int64) (*model1.SellerRatingRecord, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.SellerRatingRecord{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.SellerRatingRecord
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) SellerRatingRecord(ctx context.Context, distinctOn []model.SellerRatingRecordSelectColumn, limit *int, offset *int, orderBy []*model.SellerRatingRecordOrderBy, where *model.SellerRatingRecordBoolExp) ([]*model1.SellerRatingRecord, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.SellerRatingRecord{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.SellerRatingRecord
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) SellerRatingRecordAggregate(ctx context.Context, distinctOn []model.SellerRatingRecordSelectColumn, limit *int, offset *int, orderBy []*model.SellerRatingRecordOrderBy, where *model.SellerRatingRecordBoolExp) (*model.SellerRatingRecordAggregate, error) {
	var rs model.SellerRatingRecordAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.SellerRatingRecord{})
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

func (r *queryResolver) SellerRatingRecordByPk(ctx context.Context, Id int64) (*model1.SellerRatingRecord, error) {
	var rs model1.SellerRatingRecord
	tx := db.DB.Model(&model1.SellerRatingRecord{}).First(&rs, Id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
