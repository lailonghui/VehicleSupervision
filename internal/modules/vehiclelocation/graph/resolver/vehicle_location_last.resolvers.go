package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/vehiclelocation/graph/model"
	model1 "VehicleSupervision/internal/modules/vehiclelocation/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteVehicleLocationLast(ctx context.Context, where model.VehicleLocationLastBoolExp) (*model.VehicleLocationLastMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleLocationLast{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.VehicleLocationLast
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
	return &model.VehicleLocationLastMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteVehicleLocationLastByPk(ctx context.Context, Id int64) (*model1.VehicleLocationLast, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.VehicleLocationLast
	tx := db.DB.Model(&model1.VehicleLocationLast{})
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

func (r *mutationResolver) InsertVehicleLocationLast(ctx context.Context, objects []*model.VehicleLocationLastInsertInput) (*model.VehicleLocationLastMutationResponse, error) {
	rs := make([]*model1.VehicleLocationLast, 0)
	for _, object := range objects {
		v := &model1.VehicleLocationLast{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.VehicleLocationLast{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.VehicleLocationLastMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertVehicleLocationLastOne(ctx context.Context, object model.VehicleLocationLastInsertInput) (*model1.VehicleLocationLast, error) {
	rs := &model1.VehicleLocationLast{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.VehicleLocationLast{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateVehicleLocationLast(ctx context.Context, inc *model.VehicleLocationLastIncInput, set *model.VehicleLocationLastSetInput, where model.VehicleLocationLastBoolExp) (*model.VehicleLocationLastMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleLocationLast{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.VehicleLocationLastMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.VehicleLocationLast
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
	return &model.VehicleLocationLastMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateVehicleLocationLastByPk(ctx context.Context, inc *model.VehicleLocationLastIncInput, set *model.VehicleLocationLastSetInput, Id int64) (*model1.VehicleLocationLast, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.VehicleLocationLast{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.VehicleLocationLast
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) VehicleLocationLast(ctx context.Context, distinctOn []model.VehicleLocationLastSelectColumn, limit *int, offset *int, orderBy []*model.VehicleLocationLastOrderBy, where *model.VehicleLocationLastBoolExp) ([]*model1.VehicleLocationLast, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleLocationLast{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.VehicleLocationLast
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) VehicleLocationLastAggregate(ctx context.Context, distinctOn []model.VehicleLocationLastSelectColumn, limit *int, offset *int, orderBy []*model.VehicleLocationLastOrderBy, where *model.VehicleLocationLastBoolExp) (*model.VehicleLocationLastAggregate, error) {
	var rs model.VehicleLocationLastAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.VehicleLocationLast{})
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

func (r *queryResolver) VehicleLocationLastByPk(ctx context.Context, Id int64) (*model1.VehicleLocationLast, error) {
	var rs model1.VehicleLocationLast
	tx := db.DB.Model(&model1.VehicleLocationLast{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}
