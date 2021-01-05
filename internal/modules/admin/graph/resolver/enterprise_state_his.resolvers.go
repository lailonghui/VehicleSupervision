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

func (r *mutationResolver) DeleteEnterpriseStateHis(ctx context.Context, where model.EnterpriseStateHisBoolExp) (*model.EnterpriseStateHisMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EnterpriseStateHis{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.EnterpriseStateHis
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
	return &model.EnterpriseStateHisMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteEnterpriseStateHisByPk(ctx context.Context, id int64) (*model1.EnterpriseStateHis, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.EnterpriseStateHis
	tx := db.DB.Model(&model1.EnterpriseStateHis{})
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

func (r *mutationResolver) DeleteEnterpriseStateHisByUnionPk(ctx context.Context, stateHisID string) (*model1.EnterpriseStateHis, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.EnterpriseStateHis
	tx := db.DB.Model(&model1.EnterpriseStateHis{})
	if len(preloads) > 0 {
		// 如果请求的字段不为空，则先查询一遍数据库
		tx = tx.Select(preloads).Where(rs.UnionPrimaryColumnName()+" = ?", stateHisID).First(&rs)
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

func (r *mutationResolver) InsertEnterpriseStateHis(ctx context.Context, objects []*model.EnterpriseStateHisInsertInput) (*model.EnterpriseStateHisMutationResponse, error) {
	rs := make([]*model1.EnterpriseStateHis, 0)
	for _, object := range objects {
		v := &model1.EnterpriseStateHis{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.EnterpriseStateHis{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.EnterpriseStateHisMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertEnterpriseStateHisOne(ctx context.Context, objects model.EnterpriseStateHisInsertInput) (*model1.EnterpriseStateHis, error) {
	rs := &model1.EnterpriseStateHis{}
	util2.StructAssign(rs, &objects)
	tx := db.DB.Model(&model1.EnterpriseStateHis{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateEnterpriseStateHis(ctx context.Context, inc *model.EnterpriseStateHisIncInput, set *model.EnterpriseStateHisSetInput, where model.EnterpriseStateHisBoolExp) (*model.EnterpriseStateHisMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EnterpriseStateHis{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.EnterpriseStateHisMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.EnterpriseStateHis
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
	return &model.EnterpriseStateHisMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateEnterpriseStateHisByPk(ctx context.Context, inc *model.EnterpriseStateHisIncInput, set *model.EnterpriseStateHisSetInput, id int64) (*model1.EnterpriseStateHis, error) {
	var rs model1.EnterpriseStateHis
	tx := db.DB.Where(rs.PrimaryColumnName()+" = ?", id)
	qt := util.NewQueryTranslator(tx, &model1.EnterpriseStateHis{})
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

func (r *mutationResolver) UpdateEnterpriseStateHisByUnionPk(ctx context.Context, inc *model.EnterpriseStateHisIncInput, set *model.EnterpriseStateHisSetInput, stateHisID string) (*model1.EnterpriseStateHis, error) {
	var rs model1.EnterpriseStateHis
	tx := db.DB.Where(rs.UnionPrimaryColumnName()+" = ?", stateHisID)
	qt := util.NewQueryTranslator(tx, &model1.EnterpriseStateHis{})
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

func (r *queryResolver) EnterpriseStateHis(ctx context.Context, distinctOn []model.EnterpriseStateHisSelectColumn, limit *int, offset *int, orderBy []*model.EnterpriseStateHisOrderBy, where *model.EnterpriseStateHisBoolExp) ([]*model1.EnterpriseStateHis, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EnterpriseStateHis{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.EnterpriseStateHis
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) EnterpriseStateHisAggregate(ctx context.Context, distinctOn []model.EnterpriseStateHisSelectColumn, limit *int, offset *int, orderBy []*model.EnterpriseStateHisOrderBy, where *model.EnterpriseStateHisBoolExp) (*model.EnterpriseStateHisAggregate, error) {
	var rs model.EnterpriseStateHisAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.EnterpriseStateHis{})
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

func (r *queryResolver) EnterpriseStateHisByPk(ctx context.Context, id int64) (*model1.EnterpriseStateHis, error) {
	var rs model1.EnterpriseStateHis
	tx := db.DB.Model(&model1.EnterpriseStateHis{}).Select(util.GetTopPreloads(ctx)).First(&rs, id)
	err := tx.Error
	return &rs, err
}

func (r *queryResolver) EnterpriseStateHisByUnionPk(ctx context.Context, stateHisID string) (*model1.EnterpriseStateHis, error) {
	var rs model1.EnterpriseStateHis
	tx := db.DB.Model(&model1.EnterpriseStateHis{}).Select(util.GetTopPreloads(ctx)).Where(rs.UnionPrimaryColumnName()+" = ?", stateHisID).First(&rs)

	err := tx.Error
	return &rs, err
}
