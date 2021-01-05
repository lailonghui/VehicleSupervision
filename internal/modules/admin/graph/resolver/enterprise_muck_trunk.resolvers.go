package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/admin/graph/model"
	model1 "VehicleSupervision/internal/modules/admin/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteEnterpriseMuckTrunk(ctx context.Context, where model.EnterpriseMuckTrunkBoolExp) (*model.EnterpriseMuckTrunkMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EnterpriseMuckTrunk{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.EnterpriseMuckTrunk
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
	return &model.EnterpriseMuckTrunkMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteEnterpriseMuckTrunkByPk(ctx context.Context, id int64) (*model1.EnterpriseMuckTrunk, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.EnterpriseMuckTrunk
	tx := db.DB.Model(&model1.EnterpriseMuckTrunk{})
	if len(preloads) > 0 {
		// 如果请求的字段不为空，则先查询一遍数据库
		tx = tx.Select(preloads).Where(rs.PrimaryColumnName()+" = ?", id).First(&rs)
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

func (r *mutationResolver) DeleteEnterpriseMuckTrunkByUnionPk(ctx context.Context, enterpriseMuckTrunkID string) (*model1.EnterpriseMuckTrunk, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.EnterpriseMuckTrunk
	tx := db.DB.Model(&model1.EnterpriseMuckTrunk{})
	if len(preloads) > 0 {
		// 如果请求的字段不为空，则先查询一遍数据库
		tx = tx.Select(preloads).Where(rs.UnionPrimaryColumnName()+" = ?", enterpriseMuckTrunkID).First(&rs)
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

func (r *mutationResolver) InsertEnterpriseMuckTrunk(ctx context.Context, objects []*model.EnterpriseMuckTrunkInsertInput) (*model.EnterpriseMuckTrunkMutationResponse, error) {
	rs := make([]*model1.EnterpriseMuckTrunk, 0)
	for _, object := range objects {
		v := &model1.EnterpriseMuckTrunk{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.EnterpriseMuckTrunk{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.EnterpriseMuckTrunkMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertEnterpriseMuckTrunkOne(ctx context.Context, objects model.EnterpriseMuckTrunkInsertInput) (*model1.EnterpriseMuckTrunk, error) {
	rs := &model1.EnterpriseMuckTrunk{}
	util2.StructAssign(rs, &objects)
	tx := db.DB.Model(&model1.EnterpriseMuckTrunk{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateEnterpriseMuckTrunk(ctx context.Context, inc *model.EnterpriseMuckTrunkIncInput, set *model.EnterpriseMuckTrunkSetInput, where model.EnterpriseMuckTrunkBoolExp) (*model.EnterpriseMuckTrunkMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EnterpriseMuckTrunk{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.EnterpriseMuckTrunkMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.EnterpriseMuckTrunk
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
	return &model.EnterpriseMuckTrunkMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateEnterpriseMuckTrunkByPk(ctx context.Context, inc *model.EnterpriseMuckTrunkIncInput, set *model.EnterpriseMuckTrunkSetInput, id int64) (*model1.EnterpriseMuckTrunk, error) {
	var rs model1.EnterpriseMuckTrunk
	tx := db.DB.Where(rs.PrimaryColumnName()+" = ?", id)
	qt := util.NewQueryTranslator(tx, &model1.EnterpriseMuckTrunk{})
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

func (r *mutationResolver) UpdateEnterpriseMuckTrunkByUnionPk(ctx context.Context, inc *model.EnterpriseMuckTrunkIncInput, set *model.EnterpriseMuckTrunkSetInput, enterpriseMuckTrunkID string) (*model1.EnterpriseMuckTrunk, error) {
	var rs model1.EnterpriseMuckTrunk
	tx := db.DB.Where(rs.UnionPrimaryColumnName()+" = ?", enterpriseMuckTrunkID)
	qt := util.NewQueryTranslator(tx, &model1.EnterpriseMuckTrunk{})
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

func (r *queryResolver) EnterpriseMuckTrunk(ctx context.Context, distinctOn []model.EnterpriseMuckTrunkSelectColumn, limit *int, offset *int, orderBy []*model.EnterpriseMuckTrunkOrderBy, where *model.EnterpriseMuckTrunkBoolExp) ([]*model1.EnterpriseMuckTrunk, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EnterpriseMuckTrunk{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.EnterpriseMuckTrunk
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) EnterpriseMuckTrunkAggregate(ctx context.Context, distinctOn []model.EnterpriseMuckTrunkSelectColumn, limit *int, offset *int, orderBy []*model.EnterpriseMuckTrunkOrderBy, where *model.EnterpriseMuckTrunkBoolExp) (*model.EnterpriseMuckTrunkAggregate, error) {
	var rs model.EnterpriseMuckTrunkAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.EnterpriseMuckTrunk{})
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

func (r *queryResolver) EnterpriseMuckTrunkByPk(ctx context.Context, id int64) (*model1.EnterpriseMuckTrunk, error) {
	var rs model1.EnterpriseMuckTrunk
	tx := db.DB.Model(&model1.EnterpriseMuckTrunk{}).Select(util.GetTopPreloads(ctx)).First(&rs, id)
	err := tx.Error
	return &rs, err
}

func (r *queryResolver) EnterpriseMuckTrunkByUnionPk(ctx context.Context, enterpriseMuckTrunkID string) (*model1.EnterpriseMuckTrunk, error) {
	var rs model1.EnterpriseMuckTrunk
	tx := db.DB.Model(&model1.EnterpriseMuckTrunk{}).Select(util.GetTopPreloads(ctx)).Where(rs.UnionPrimaryColumnName()+" = ?", enterpriseMuckTrunkID).First(&rs)

	err := tx.Error
	return &rs, err
}
