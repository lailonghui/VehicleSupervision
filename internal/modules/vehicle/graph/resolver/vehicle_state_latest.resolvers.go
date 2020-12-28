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

func (r *mutationResolver) DeleteVehicleStateLatest(ctx context.Context, where model.VehicleStateLatestBoolExp) (*model.VehicleStateLatestMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleStateLatest{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.VehicleStateLatest
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
	return &model.VehicleStateLatestMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteVehicleStateLatestByPk(ctx context.Context, Id int64) (*model1.VehicleStateLatest, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.VehicleStateLatest
	tx := db.DB.Model(&model1.VehicleStateLatest{})
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

func (r *mutationResolver) InsertVehicleStateLatest(ctx context.Context, objects []*model.VehicleStateLatestInsertInput) (*model.VehicleStateLatestMutationResponse, error) {
	rs := make([]*model1.VehicleStateLatest, 0)
	for _, object := range objects {
		v := &model1.VehicleStateLatest{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.VehicleStateLatest{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.VehicleStateLatestMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertVehicleStateLatestOne(ctx context.Context, object model.VehicleStateLatestInsertInput) (*model1.VehicleStateLatest, error) {
	rs := &model1.VehicleStateLatest{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.VehicleStateLatest{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateVehicleStateLatest(ctx context.Context, inc *model.VehicleStateLatestIncInput, set *model.VehicleStateLatestSetInput, where model.VehicleStateLatestBoolExp) (*model.VehicleStateLatestMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleStateLatest{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.VehicleStateLatestMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.VehicleStateLatestMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateVehicleStateLatestByPk(ctx context.Context, inc *model.VehicleStateLatestIncInput, set *model.VehicleStateLatestSetInput, Id int64) (*model1.VehicleStateLatest, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.VehicleStateLatest{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.VehicleStateLatest
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) VehicleStateLatest(ctx context.Context, distinctOn []model.VehicleStateLatestSelectColumn, limit *int, offset *int, orderBy []*model.VehicleStateLatestOrderBy, where *model.VehicleStateLatestBoolExp) ([]*model1.VehicleStateLatest, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleStateLatest{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.VehicleStateLatest
	tx = tx.Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) VehicleStateLatestAggregate(ctx context.Context, distinctOn []model.VehicleStateLatestSelectColumn, limit *int, offset *int, orderBy []*model.VehicleStateLatestOrderBy, where *model.VehicleStateLatestBoolExp) (*model.VehicleStateLatestAggregate, error) {
	var rs model.VehicleStateLatestAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.VehicleStateLatest{})
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

func (r *queryResolver) VehicleStateLatestByPk(ctx context.Context, Id int64) (*model1.VehicleStateLatest, error) {
	var rs model1.VehicleStateLatest
	tx := db.DB.Model(&model1.VehicleStateLatest{}).First(&rs, Id)
	err := tx.Error
	return &rs, err
}
