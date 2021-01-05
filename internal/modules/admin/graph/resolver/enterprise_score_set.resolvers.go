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

func (r *mutationResolver) DeleteEnterpriseScoreSet(ctx context.Context, where model.EnterpriseScoreSetBoolExp) (*model.EnterpriseScoreSetMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EnterpriseScoreSet{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.EnterpriseScoreSet
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
	return &model.EnterpriseScoreSetMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteEnterpriseScoreSetByPk(ctx context.Context, id int64) (*model1.EnterpriseScoreSet, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.EnterpriseScoreSet
	tx := db.DB.Model(&model1.EnterpriseScoreSet{})
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

func (r *mutationResolver) DeleteEnterpriseScoreSetByUnionPk(ctx context.Context, scoreSetID string) (*model1.EnterpriseScoreSet, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.EnterpriseScoreSet
	tx := db.DB.Model(&model1.EnterpriseScoreSet{})
	if len(preloads) > 0 {
		// 如果请求的字段不为空，则先查询一遍数据库
		tx = tx.Select(preloads).Where(rs.UnionPrimaryColumnName()+" = ?", scoreSetID).First(&rs)
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

func (r *mutationResolver) InsertEnterpriseScoreSet(ctx context.Context, objects []*model.EnterpriseScoreSetInsertInput) (*model.EnterpriseScoreSetMutationResponse, error) {
	rs := make([]*model1.EnterpriseScoreSet, 0)
	for _, object := range objects {
		v := &model1.EnterpriseScoreSet{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.EnterpriseScoreSet{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.EnterpriseScoreSetMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertEnterpriseScoreSetOne(ctx context.Context, objects model.EnterpriseScoreSetInsertInput) (*model1.EnterpriseScoreSet, error) {
	rs := &model1.EnterpriseScoreSet{}
	util2.StructAssign(rs, &objects)
	tx := db.DB.Model(&model1.EnterpriseScoreSet{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateEnterpriseScoreSet(ctx context.Context, inc *model.EnterpriseScoreSetIncInput, set *model.EnterpriseScoreSetSetInput, where model.EnterpriseScoreSetBoolExp) (*model.EnterpriseScoreSetMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EnterpriseScoreSet{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.EnterpriseScoreSetMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.EnterpriseScoreSet
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
	return &model.EnterpriseScoreSetMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateEnterpriseScoreSetByPk(ctx context.Context, inc *model.EnterpriseScoreSetIncInput, set *model.EnterpriseScoreSetSetInput, id int64) (*model1.EnterpriseScoreSet, error) {
	var rs model1.EnterpriseScoreSet
	tx := db.DB.Where(rs.PrimaryColumnName()+" = ?", id)
	qt := util.NewQueryTranslator(tx, &model1.EnterpriseScoreSet{})
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

func (r *mutationResolver) UpdateEnterpriseScoreSetByUnionPk(ctx context.Context, inc *model.EnterpriseScoreSetIncInput, set *model.EnterpriseScoreSetSetInput, scoreSetID string) (*model1.EnterpriseScoreSet, error) {
	var rs model1.EnterpriseScoreSet
	tx := db.DB.Where(rs.UnionPrimaryColumnName()+" = ?", scoreSetID)
	qt := util.NewQueryTranslator(tx, &model1.EnterpriseScoreSet{})
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

func (r *queryResolver) EnterpriseScoreSet(ctx context.Context, distinctOn []model.EnterpriseScoreSetSelectColumn, limit *int, offset *int, orderBy []*model.EnterpriseScoreSetOrderBy, where *model.EnterpriseScoreSetBoolExp) ([]*model1.EnterpriseScoreSet, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EnterpriseScoreSet{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.EnterpriseScoreSet
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) EnterpriseScoreSetAggregate(ctx context.Context, distinctOn []model.EnterpriseScoreSetSelectColumn, limit *int, offset *int, orderBy []*model.EnterpriseScoreSetOrderBy, where *model.EnterpriseScoreSetBoolExp) (*model.EnterpriseScoreSetAggregate, error) {
	var rs model.EnterpriseScoreSetAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.EnterpriseScoreSet{})
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

func (r *queryResolver) EnterpriseScoreSetByPk(ctx context.Context, id int64) (*model1.EnterpriseScoreSet, error) {
	var rs model1.EnterpriseScoreSet
	tx := db.DB.Model(&model1.EnterpriseScoreSet{}).Select(util.GetTopPreloads(ctx)).First(&rs, id)
	err := tx.Error
	return &rs, err
}

func (r *queryResolver) EnterpriseScoreSetByUnionPk(ctx context.Context, scoreSetID string) (*model1.EnterpriseScoreSet, error) {
	var rs model1.EnterpriseScoreSet
	tx := db.DB.Model(&model1.EnterpriseScoreSet{}).Select(util.GetTopPreloads(ctx)).Where(rs.UnionPrimaryColumnName()+" = ?", scoreSetID).First(&rs)

	err := tx.Error
	return &rs, err
}
