package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/dynamic_supervision/graph/model"
	model1 "VehicleSupervision/internal/modules/dynamic_supervision/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteMuckTruckOnline(ctx context.Context, where model.MuckTruckOnlineBoolExp) (*model.MuckTruckOnlineMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.MuckTruckOnline{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.MuckTruckOnline
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
	return &model.MuckTruckOnlineMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteMuckTruckOnlineByPk(ctx context.Context, Id int64) (*model1.MuckTruckOnline, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.MuckTruckOnline
	tx := db.DB.Model(&model1.MuckTruckOnline{})
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

func (r *mutationResolver) InsertMuckTruckOnline(ctx context.Context, objects []*model.MuckTruckOnlineInsertInput) (*model.MuckTruckOnlineMutationResponse, error) {
	rs := []*model1.MuckTruckOnline{}
	for _, object := range objects {
		v := &model1.MuckTruckOnline{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.MuckTruckOnline{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.MuckTruckOnlineMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertMuckTruckOnlineOne(ctx context.Context, object model.MuckTruckOnlineInsertInput) (*model1.MuckTruckOnline, error) {
	rs := &model1.MuckTruckOnline{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.MuckTruckOnline{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateMuckTruckOnline(ctx context.Context, inc *model.MuckTruckOnlineIncInput, set *model.MuckTruckOnlineSetInput, where model.MuckTruckOnlineBoolExp) (*model.MuckTruckOnlineMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.MuckTruckOnline{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.MuckTruckOnlineMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.MuckTruckOnlineMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateMuckTruckOnlineByPk(ctx context.Context, inc *model.MuckTruckOnlineIncInput, set *model.MuckTruckOnlineSetInput, Id int64) (*model1.MuckTruckOnline, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.MuckTruckOnline{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.MuckTruckOnline
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) MuckTruckOnline(ctx context.Context, distinctOn []model.MuckTruckOnlineSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckOnlineOrderBy, where *model.MuckTruckOnlineBoolExp) ([]*model1.MuckTruckOnline, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.MuckTruckOnline{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.MuckTruckOnline
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) MuckTruckOnlineAggregate(ctx context.Context, distinctOn []model.MuckTruckOnlineSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckOnlineOrderBy, where *model.MuckTruckOnlineBoolExp) (*model.MuckTruckOnlineAggregate, error) {
	var rs model.MuckTruckOnlineAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.MuckTruckOnline{})
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

func (r *queryResolver) MuckTruckOnlineByPk(ctx context.Context, Id int64) (*model1.MuckTruckOnline, error) {
	var rs model1.MuckTruckOnline
	tx := db.DB.Model(&model1.MuckTruckOnline{}).First(&rs, Id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
