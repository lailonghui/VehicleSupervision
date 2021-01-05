package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/device/graph/model"
	model1 "VehicleSupervision/internal/modules/device/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteTerminalBeidouValid(ctx context.Context, where model.TerminalBeidouValidBoolExp) (*model.TerminalBeidouValidMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalBeidouValid{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.TerminalBeidouValid
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
	return &model.TerminalBeidouValidMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteTerminalBeidouValidByPk(ctx context.Context, Id int64) (*model1.TerminalBeidouValid, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.TerminalBeidouValid
	tx := db.DB.Model(&model1.TerminalBeidouValid{})
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

func (r *mutationResolver) DeleteTerminalBeidouValidByUnionPk(ctx context.Context, unionId string) (*model1.TerminalBeidouValid, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.TerminalBeidouValid
	tx := db.DB.Model(&model1.TerminalBeidouValid{})
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

func (r *mutationResolver) InsertTerminalBeidouValid(ctx context.Context, objects []*model.TerminalBeidouValidInsertInput) (*model.TerminalBeidouValidMutationResponse, error) {
	rs := make([]*model1.TerminalBeidouValid, 0)
	for _, object := range objects {
		v := &model1.TerminalBeidouValid{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.TerminalBeidouValid{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.TerminalBeidouValidMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertTerminalBeidouValidOne(ctx context.Context, object model.TerminalBeidouValidInsertInput) (*model1.TerminalBeidouValid, error) {
	rs := &model1.TerminalBeidouValid{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.TerminalBeidouValid{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateTerminalBeidouValid(ctx context.Context, inc *model.TerminalBeidouValidIncInput, set *model.TerminalBeidouValidSetInput, where model.TerminalBeidouValidBoolExp) (*model.TerminalBeidouValidMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalBeidouValid{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.TerminalBeidouValidMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.TerminalBeidouValid
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
	return &model.TerminalBeidouValidMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateTerminalBeidouValidByPk(ctx context.Context, inc *model.TerminalBeidouValidIncInput, set *model.TerminalBeidouValidSetInput, Id int64) (*model1.TerminalBeidouValid, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.TerminalBeidouValid{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.TerminalBeidouValid
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *mutationResolver) UpdateTerminalBeidouValidByUnionPk(ctx context.Context, inc *model.TerminalBeidouValidIncInput, set *model.TerminalBeidouValidSetInput, unionId string) (*model1.TerminalBeidouValid, error) {
	var rs model1.TerminalBeidouValid
	tx := db.DB.Where(rs.UnionPrimaryColumnName()+" = ?", unionId)
	qt := util.NewQueryTranslator(tx, &model1.TerminalBeidouValid{})
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

func (r *queryResolver) TerminalBeidouValid(ctx context.Context, distinctOn []model.TerminalBeidouValidSelectColumn, limit *int, offset *int, orderBy []*model.TerminalBeidouValidOrderBy, where *model.TerminalBeidouValidBoolExp) ([]*model1.TerminalBeidouValid, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalBeidouValid{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.TerminalBeidouValid
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) TerminalBeidouValidAggregate(ctx context.Context, distinctOn []model.TerminalBeidouValidSelectColumn, limit *int, offset *int, orderBy []*model.TerminalBeidouValidOrderBy, where *model.TerminalBeidouValidBoolExp) (*model.TerminalBeidouValidAggregate, error) {
	var rs model.TerminalBeidouValidAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.TerminalBeidouValid{})
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

func (r *queryResolver) TerminalBeidouValidByPk(ctx context.Context, Id int64) (*model1.TerminalBeidouValid, error) {
	var rs model1.TerminalBeidouValid
	tx := db.DB.Model(&model1.TerminalBeidouValid{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}

func (r *queryResolver) TerminalBeidouValidByUnionPk(ctx context.Context, unionId string) (*model1.TerminalBeidouValid, error) {
	var rs model1.TerminalBeidouValid
	tx := db.DB.Model(&model1.TerminalBeidouValid{}).Select(util.GetTopPreloads(ctx)).Where(rs.UnionPrimaryColumnName()+" = ?", unionId).First(&rs)

	err := tx.Error
	return &rs, err
}
