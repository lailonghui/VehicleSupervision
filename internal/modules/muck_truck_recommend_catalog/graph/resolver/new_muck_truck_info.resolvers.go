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

func (r *mutationResolver) DeleteNewMuckTruckInfo(ctx context.Context, where model.NewMuckTruckInfoBoolExp) (*model.NewMuckTruckInfoMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.NewMuckTruckInfo{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.NewMuckTruckInfo
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
	return &model.NewMuckTruckInfoMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteNewMuckTruckInfoByPk(ctx context.Context, Id int64) (*model1.NewMuckTruckInfo, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.NewMuckTruckInfo
	tx := db.DB.Model(&model1.NewMuckTruckInfo{})
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

func (r *mutationResolver) InsertNewMuckTruckInfo(ctx context.Context, objects []*model.NewMuckTruckInfoInsertInput) (*model.NewMuckTruckInfoMutationResponse, error) {
	rs := []*model1.NewMuckTruckInfo{}
	for _, object := range objects {
		v := &model1.NewMuckTruckInfo{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.NewMuckTruckInfo{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.NewMuckTruckInfoMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertNewMuckTruckInfoOne(ctx context.Context, object model.NewMuckTruckInfoInsertInput) (*model1.NewMuckTruckInfo, error) {
	rs := &model1.NewMuckTruckInfo{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.NewMuckTruckInfo{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateNewMuckTruckInfo(ctx context.Context, inc *model.NewMuckTruckInfoIncInput, set *model.NewMuckTruckInfoSetInput, where model.NewMuckTruckInfoBoolExp) (*model.NewMuckTruckInfoMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.NewMuckTruckInfo{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.NewMuckTruckInfoMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.NewMuckTruckInfoMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateNewMuckTruckInfoByPk(ctx context.Context, inc *model.NewMuckTruckInfoIncInput, set *model.NewMuckTruckInfoSetInput, Id int64) (*model1.NewMuckTruckInfo, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.NewMuckTruckInfo{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.NewMuckTruckInfo
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) NewMuckTruckInfo(ctx context.Context, distinctOn []model.NewMuckTruckInfoSelectColumn, limit *int, offset *int, orderBy []*model.NewMuckTruckInfoOrderBy, where *model.NewMuckTruckInfoBoolExp) ([]*model1.NewMuckTruckInfo, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.NewMuckTruckInfo{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.NewMuckTruckInfo
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) NewMuckTruckInfoAggregate(ctx context.Context, distinctOn []model.NewMuckTruckInfoSelectColumn, limit *int, offset *int, orderBy []*model.NewMuckTruckInfoOrderBy, where *model.NewMuckTruckInfoBoolExp) (*model.NewMuckTruckInfoAggregate, error) {
	var rs model.NewMuckTruckInfoAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.NewMuckTruckInfo{})
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

func (r *queryResolver) NewMuckTruckInfoByPk(ctx context.Context, Id int64) (*model1.NewMuckTruckInfo, error) {
	var rs model1.NewMuckTruckInfo
	tx := db.DB.Model(&model1.NewMuckTruckInfo{}).First(&rs, Id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
