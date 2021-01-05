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

func (r *mutationResolver) DeleteDriverBlacklistHis(ctx context.Context, where model.DriverBlacklistHisBoolExp) (*model.DriverBlacklistHisMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DriverBlacklistHis{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.DriverBlacklistHis
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
	return &model.DriverBlacklistHisMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteDriverBlacklistHisByPk(ctx context.Context, Id int64) (*model1.DriverBlacklistHis, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.DriverBlacklistHis
	tx := db.DB.Model(&model1.DriverBlacklistHis{})
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

func (r *mutationResolver) DeleteDriverBlacklistHisByUnionPk(ctx context.Context, unionId string) (*model1.DriverBlacklistHis, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.DriverBlacklistHis
	tx := db.DB.Model(&model1.DriverBlacklistHis{})
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

func (r *mutationResolver) InsertDriverBlacklistHis(ctx context.Context, objects []*model.DriverBlacklistHisInsertInput) (*model.DriverBlacklistHisMutationResponse, error) {
	rs := make([]*model1.DriverBlacklistHis, 0)
	for _, object := range objects {
		v := &model1.DriverBlacklistHis{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.DriverBlacklistHis{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.DriverBlacklistHisMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertDriverBlacklistHisOne(ctx context.Context, object model.DriverBlacklistHisInsertInput) (*model1.DriverBlacklistHis, error) {
	rs := &model1.DriverBlacklistHis{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.DriverBlacklistHis{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateDriverBlacklistHis(ctx context.Context, inc *model.DriverBlacklistHisIncInput, set *model.DriverBlacklistHisSetInput, where model.DriverBlacklistHisBoolExp) (*model.DriverBlacklistHisMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DriverBlacklistHis{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.DriverBlacklistHisMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.DriverBlacklistHis
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
	return &model.DriverBlacklistHisMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateDriverBlacklistHisByPk(ctx context.Context, inc *model.DriverBlacklistHisIncInput, set *model.DriverBlacklistHisSetInput, Id int64) (*model1.DriverBlacklistHis, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.DriverBlacklistHis{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.DriverBlacklistHis
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *mutationResolver) UpdateDriverBlacklistHisByUnionPk(ctx context.Context, inc *model.DriverBlacklistHisIncInput, set *model.DriverBlacklistHisSetInput, unionId string) (*model1.DriverBlacklistHis, error) {
	var rs model1.DriverBlacklistHis
	tx := db.DB.Where(rs.UnionPrimaryColumnName()+" = ?", unionId)
	qt := util.NewQueryTranslator(tx, &model1.DriverBlacklistHis{})
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

func (r *queryResolver) DriverBlacklistHis(ctx context.Context, distinctOn []model.DriverBlacklistHisSelectColumn, limit *int, offset *int, orderBy []*model.DriverBlacklistHisOrderBy, where *model.DriverBlacklistHisBoolExp) ([]*model1.DriverBlacklistHis, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DriverBlacklistHis{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.DriverBlacklistHis
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) DriverBlacklistHisAggregate(ctx context.Context, distinctOn []model.DriverBlacklistHisSelectColumn, limit *int, offset *int, orderBy []*model.DriverBlacklistHisOrderBy, where *model.DriverBlacklistHisBoolExp) (*model.DriverBlacklistHisAggregate, error) {
	var rs model.DriverBlacklistHisAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.DriverBlacklistHis{})
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

func (r *queryResolver) DriverBlacklistHisByPk(ctx context.Context, Id int64) (*model1.DriverBlacklistHis, error) {
	var rs model1.DriverBlacklistHis
	tx := db.DB.Model(&model1.DriverBlacklistHis{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}

func (r *queryResolver) DriverBlacklistHisByUnionPk(ctx context.Context, unionId string) (*model1.DriverBlacklistHis, error) {
	var rs model1.DriverBlacklistHis
	tx := db.DB.Model(&model1.DriverBlacklistHis{}).Select(util.GetTopPreloads(ctx)).Where(rs.UnionPrimaryColumnName()+" = ?", unionId).First(&rs)

	err := tx.Error
	return &rs, err
}
