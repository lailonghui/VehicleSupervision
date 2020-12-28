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

func (r *mutationResolver) DeleteMuckTruckExt(ctx context.Context, where model.MuckTruckExtBoolExp) (*model.MuckTruckExtMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.MuckTruckExt{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.MuckTruckExt
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
	return &model.MuckTruckExtMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteMuckTruckExtByPk(ctx context.Context, Id int64) (*model1.MuckTruckExt, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.MuckTruckExt
	tx := db.DB.Model(&model1.MuckTruckExt{})
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

func (r *mutationResolver) InsertMuckTruckExt(ctx context.Context, objects []*model.MuckTruckExtInsertInput) (*model.MuckTruckExtMutationResponse, error) {
	rs := []*model1.MuckTruckExt{}
	for _, object := range objects {
		v := &model1.MuckTruckExt{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.MuckTruckExt{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.MuckTruckExtMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertMuckTruckExtOne(ctx context.Context, object model.MuckTruckExtInsertInput) (*model1.MuckTruckExt, error) {
	rs := &model1.MuckTruckExt{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.MuckTruckExt{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateMuckTruckExt(ctx context.Context, inc *model.MuckTruckExtIncInput, set *model.MuckTruckExtSetInput, where model.MuckTruckExtBoolExp) (*model.MuckTruckExtMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.MuckTruckExt{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.MuckTruckExtMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.MuckTruckExtMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateMuckTruckExtByPk(ctx context.Context, inc *model.MuckTruckExtIncInput, set *model.MuckTruckExtSetInput, Id int64) (*model1.MuckTruckExt, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.MuckTruckExt{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.MuckTruckExt
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) MuckTruckExt(ctx context.Context, distinctOn []model.MuckTruckExtSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckExtOrderBy, where *model.MuckTruckExtBoolExp) ([]*model1.MuckTruckExt, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.MuckTruckExt{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.MuckTruckExt
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) MuckTruckExtAggregate(ctx context.Context, distinctOn []model.MuckTruckExtSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckExtOrderBy, where *model.MuckTruckExtBoolExp) (*model.MuckTruckExtAggregate, error) {
	var rs model.MuckTruckExtAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.MuckTruckExt{})
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

func (r *queryResolver) MuckTruckExtByPk(ctx context.Context, Id int64) (*model1.MuckTruckExt, error) {
	var rs model1.MuckTruckExt
	tx := db.DB.Model(&model1.MuckTruckExt{}).First(&rs, Id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
