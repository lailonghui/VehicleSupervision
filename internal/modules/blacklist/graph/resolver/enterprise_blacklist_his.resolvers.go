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

func (r *mutationResolver) DeleteEnterpriseBlacklistHis(ctx context.Context, where model.EnterpriseBlacklistHisBoolExp) (*model.EnterpriseBlacklistHisMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EnterpriseBlacklistHis{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.EnterpriseBlacklistHis
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
	return &model.EnterpriseBlacklistHisMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteEnterpriseBlacklistHisByPk(ctx context.Context, Id int64) (*model1.EnterpriseBlacklistHis, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.EnterpriseBlacklistHis
	tx := db.DB.Model(&model1.EnterpriseBlacklistHis{})
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

func (r *mutationResolver) DeleteEnterpriseBlacklistHisByUnionPk(ctx context.Context, unionId string) (*model1.EnterpriseBlacklistHis, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.EnterpriseBlacklistHis
	tx := db.DB.Model(&model1.EnterpriseBlacklistHis{})
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

func (r *mutationResolver) InsertEnterpriseBlacklistHis(ctx context.Context, objects []*model.EnterpriseBlacklistHisInsertInput) (*model.EnterpriseBlacklistHisMutationResponse, error) {
	rs := make([]*model1.EnterpriseBlacklistHis, 0)
	for _, object := range objects {
		v := &model1.EnterpriseBlacklistHis{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.EnterpriseBlacklistHis{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.EnterpriseBlacklistHisMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertEnterpriseBlacklistHisOne(ctx context.Context, object model.EnterpriseBlacklistHisInsertInput) (*model1.EnterpriseBlacklistHis, error) {
	rs := &model1.EnterpriseBlacklistHis{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.EnterpriseBlacklistHis{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateEnterpriseBlacklistHis(ctx context.Context, inc *model.EnterpriseBlacklistHisIncInput, set *model.EnterpriseBlacklistHisSetInput, where model.EnterpriseBlacklistHisBoolExp) (*model.EnterpriseBlacklistHisMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EnterpriseBlacklistHis{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.EnterpriseBlacklistHisMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.EnterpriseBlacklistHis
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
	return &model.EnterpriseBlacklistHisMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateEnterpriseBlacklistHisByPk(ctx context.Context, inc *model.EnterpriseBlacklistHisIncInput, set *model.EnterpriseBlacklistHisSetInput, Id int64) (*model1.EnterpriseBlacklistHis, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.EnterpriseBlacklistHis{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.EnterpriseBlacklistHis
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *mutationResolver) UpdateEnterpriseBlacklistHisByUnionPk(ctx context.Context, inc *model.EnterpriseBlacklistHisIncInput, set *model.EnterpriseBlacklistHisSetInput, unionId string) (*model1.EnterpriseBlacklistHis, error) {
	var rs model1.EnterpriseBlacklistHis
	tx := db.DB.Where(rs.UnionPrimaryColumnName()+" = ?", unionId)
	qt := util.NewQueryTranslator(tx, &model1.EnterpriseBlacklistHis{})
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

func (r *queryResolver) EnterpriseBlacklistHis(ctx context.Context, distinctOn []model.EnterpriseBlacklistHisSelectColumn, limit *int, offset *int, orderBy []*model.EnterpriseBlacklistHisOrderBy, where *model.EnterpriseBlacklistHisBoolExp) ([]*model1.EnterpriseBlacklistHis, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EnterpriseBlacklistHis{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.EnterpriseBlacklistHis
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) EnterpriseBlacklistHisAggregate(ctx context.Context, distinctOn []model.EnterpriseBlacklistHisSelectColumn, limit *int, offset *int, orderBy []*model.EnterpriseBlacklistHisOrderBy, where *model.EnterpriseBlacklistHisBoolExp) (*model.EnterpriseBlacklistHisAggregate, error) {
	var rs model.EnterpriseBlacklistHisAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.EnterpriseBlacklistHis{})
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

func (r *queryResolver) EnterpriseBlacklistHisByPk(ctx context.Context, Id int64) (*model1.EnterpriseBlacklistHis, error) {
	var rs model1.EnterpriseBlacklistHis
	tx := db.DB.Model(&model1.EnterpriseBlacklistHis{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}

func (r *queryResolver) EnterpriseBlacklistHisByUnionPk(ctx context.Context, unionId string) (*model1.EnterpriseBlacklistHis, error) {
	var rs model1.EnterpriseBlacklistHis
	tx := db.DB.Model(&model1.EnterpriseBlacklistHis{}).Select(util.GetTopPreloads(ctx)).Where(rs.UnionPrimaryColumnName()+" = ?", unionId).First(&rs)

	err := tx.Error
	return &rs, err
}
