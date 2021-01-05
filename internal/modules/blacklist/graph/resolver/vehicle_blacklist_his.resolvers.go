package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/blacklist/graph/model"
	model1 "VehicleSupervision/internal/modules/blacklist/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteVehicleBlacklistHis(ctx context.Context, where model.VehicleBlacklistHisBoolExp) (*model.VehicleBlacklistHisMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleBlacklistHis{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.VehicleBlacklistHis
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
	return &model.VehicleBlacklistHisMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteVehicleBlacklistHisByPk(ctx context.Context, Id int64) (*model1.VehicleBlacklistHis, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.VehicleBlacklistHis
	tx := db.DB.Model(&model1.VehicleBlacklistHis{})
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

func (r *mutationResolver) DeleteVehicleBlacklistHisByUnionPk(ctx context.Context, unionId string) (*model1.VehicleBlacklistHis, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.VehicleBlacklistHis
	tx := db.DB.Model(&model1.VehicleBlacklistHis{})
	if len(preloads) > 0 {
		// 如果请求的字段不为空，则先查询一遍数据库
		tx = tx.Select(preloads).Where(rs.UnionPrimaryColumnName()+" = ?", unionId).First(&rs)
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

func (r *mutationResolver) InsertVehicleBlacklistHis(ctx context.Context, objects []*model.VehicleBlacklistHisInsertInput) (*model.VehicleBlacklistHisMutationResponse, error) {
	rs := make([]*model1.VehicleBlacklistHis, 0)
	for _, object := range objects {
		v := &model1.VehicleBlacklistHis{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.VehicleBlacklistHis{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.VehicleBlacklistHisMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertVehicleBlacklistHisOne(ctx context.Context, object model.VehicleBlacklistHisInsertInput) (*model1.VehicleBlacklistHis, error) {
	rs := &model1.VehicleBlacklistHis{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.VehicleBlacklistHis{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateVehicleBlacklistHis(ctx context.Context, inc *model.VehicleBlacklistHisIncInput, set *model.VehicleBlacklistHisSetInput, where model.VehicleBlacklistHisBoolExp) (*model.VehicleBlacklistHisMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleBlacklistHis{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.VehicleBlacklistHisMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.VehicleBlacklistHis
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
	return &model.VehicleBlacklistHisMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateVehicleBlacklistHisByPk(ctx context.Context, inc *model.VehicleBlacklistHisIncInput, set *model.VehicleBlacklistHisSetInput, Id int64) (*model1.VehicleBlacklistHis, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.VehicleBlacklistHis{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.VehicleBlacklistHis
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *mutationResolver) UpdateVehicleBlacklistHisByUnionPk(ctx context.Context, inc *model.VehicleBlacklistHisIncInput, set *model.VehicleBlacklistHisSetInput, unionId string) (*model1.VehicleBlacklistHis, error) {
	var rs model1.VehicleBlacklistHis
	tx := db.DB.Where(rs.UnionPrimaryColumnName()+" = ?", unionId)
	qt := util.NewQueryTranslator(tx, &model1.VehicleBlacklistHis{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}

	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) VehicleBlacklistHis(ctx context.Context, distinctOn []model.VehicleBlacklistHisSelectColumn, limit *int, offset *int, orderBy []*model.VehicleBlacklistHisOrderBy, where *model.VehicleBlacklistHisBoolExp) ([]*model1.VehicleBlacklistHis, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleBlacklistHis{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.VehicleBlacklistHis
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) VehicleBlacklistHisAggregate(ctx context.Context, distinctOn []model.VehicleBlacklistHisSelectColumn, limit *int, offset *int, orderBy []*model.VehicleBlacklistHisOrderBy, where *model.VehicleBlacklistHisBoolExp) (*model.VehicleBlacklistHisAggregate, error) {
	var rs model.VehicleBlacklistHisAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.VehicleBlacklistHis{})
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

func (r *queryResolver) VehicleBlacklistHisByPk(ctx context.Context, Id int64) (*model1.VehicleBlacklistHis, error) {
	var rs model1.VehicleBlacklistHis
	tx := db.DB.Model(&model1.VehicleBlacklistHis{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}

func (r *queryResolver) VehicleBlacklistHisByUnionPk(ctx context.Context, unionId string) (*model1.VehicleBlacklistHis, error) {
	var rs model1.VehicleBlacklistHis
	tx := db.DB.Model(&model1.VehicleBlacklistHis{}).Select(util.GetTopPreloads(ctx)).Where(rs.UnionPrimaryColumnName()+" = ?", unionId).First(&rs)

	err := tx.Error
	return &rs, err
}
