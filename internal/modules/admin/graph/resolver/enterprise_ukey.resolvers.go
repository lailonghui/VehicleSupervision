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

func (r *mutationResolver) DeleteEnterpriseUkey(ctx context.Context, where model.EnterpriseUkeyBoolExp) (*model.EnterpriseUkeyMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EnterpriseUkey{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.EnterpriseUkey
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
	return &model.EnterpriseUkeyMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteEnterpriseUkeyByPk(ctx context.Context, id int64) (*model1.EnterpriseUkey, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.EnterpriseUkey
	tx := db.DB.Model(&model1.EnterpriseUkey{})
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

func (r *mutationResolver) DeleteEnterpriseUkeyByUnionPk(ctx context.Context, ukeyID string) (*model1.EnterpriseUkey, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.EnterpriseUkey
	tx := db.DB.Model(&model1.EnterpriseUkey{})
	if len(preloads) > 0 {
		// 如果请求的字段不为空，则先查询一遍数据库
		tx = tx.Select(preloads).Where(rs.UnionPrimaryColumnName()+" = ?", ukeyID).First(&rs)
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

func (r *mutationResolver) InsertEnterpriseUkey(ctx context.Context, objects []*model.EnterpriseUkeyInsertInput) (*model.EnterpriseUkeyMutationResponse, error) {
	rs := make([]*model1.EnterpriseUkey, 0)
	for _, object := range objects {
		v := &model1.EnterpriseUkey{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.EnterpriseUkey{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.EnterpriseUkeyMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertEnterpriseUkeyOne(ctx context.Context, objects model.EnterpriseUkeyInsertInput) (*model1.EnterpriseUkey, error) {
	rs := &model1.EnterpriseUkey{}
	util2.StructAssign(rs, &objects)
	tx := db.DB.Model(&model1.EnterpriseUkey{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateEnterpriseUkey(ctx context.Context, inc *model.EnterpriseUkeyIncInput, set *model.EnterpriseUkeySetInput, where model.EnterpriseUkeyBoolExp) (*model.EnterpriseUkeyMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EnterpriseUkey{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.EnterpriseUkeyMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.EnterpriseUkey
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
	return &model.EnterpriseUkeyMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateEnterpriseUkeyByPk(ctx context.Context, inc *model.EnterpriseUkeyIncInput, set *model.EnterpriseUkeySetInput, id int64) (*model1.EnterpriseUkey, error) {
	var rs model1.EnterpriseUkey
	tx := db.DB.Where(rs.PrimaryColumnName()+" = ?", id)
	qt := util.NewQueryTranslator(tx, &model1.EnterpriseUkey{})
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

func (r *mutationResolver) UpdateEnterpriseUkeyByUnionPk(ctx context.Context, inc *model.EnterpriseUkeyIncInput, set *model.EnterpriseUkeySetInput, ukeyID string) (*model1.EnterpriseUkey, error) {
	var rs model1.EnterpriseUkey
	tx := db.DB.Where(rs.UnionPrimaryColumnName()+" = ?", ukeyID)
	qt := util.NewQueryTranslator(tx, &model1.EnterpriseUkey{})
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

func (r *queryResolver) EnterpriseUkey(ctx context.Context, distinctOn []model.EnterpriseUkeySelectColumn, limit *int, offset *int, orderBy []*model.EnterpriseUkeyOrderBy, where *model.EnterpriseUkeyBoolExp) ([]*model1.EnterpriseUkey, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EnterpriseUkey{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.EnterpriseUkey
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) EnterpriseUkeyAggregate(ctx context.Context, distinctOn []model.EnterpriseUkeySelectColumn, limit *int, offset *int, orderBy []*model.EnterpriseUkeyOrderBy, where *model.EnterpriseUkeyBoolExp) (*model.EnterpriseUkeyAggregate, error) {
	var rs model.EnterpriseUkeyAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.EnterpriseUkey{})
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

func (r *queryResolver) EnterpriseUkeyByPk(ctx context.Context, id int64) (*model1.EnterpriseUkey, error) {
	var rs model1.EnterpriseUkey
	tx := db.DB.Model(&model1.EnterpriseUkey{}).Select(util.GetTopPreloads(ctx)).First(&rs, id)
	err := tx.Error
	return &rs, err
}

func (r *queryResolver) EnterpriseUkeyByUnionPk(ctx context.Context, ukeyID string) (*model1.EnterpriseUkey, error) {
	var rs model1.EnterpriseUkey
	tx := db.DB.Model(&model1.EnterpriseUkey{}).Select(util.GetTopPreloads(ctx)).Where(rs.UnionPrimaryColumnName()+" = ?", ukeyID).First(&rs)

	err := tx.Error
	return &rs, err
}
