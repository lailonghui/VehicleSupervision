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

func (r *mutationResolver) DeleteTerminalProve(ctx context.Context, where model.TerminalProveBoolExp) (*model.TerminalProveMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalProve{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.TerminalProve
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
	return &model.TerminalProveMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteTerminalProveByPk(ctx context.Context, Id int64) (*model1.TerminalProve, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.TerminalProve
	tx := db.DB.Model(&model1.TerminalProve{})
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

func (r *mutationResolver) DeleteTerminalProveByUnionPk(ctx context.Context, unionId string) (*model1.TerminalProve, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.TerminalProve
	tx := db.DB.Model(&model1.TerminalProve{})
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

func (r *mutationResolver) InsertTerminalProve(ctx context.Context, objects []*model.TerminalProveInsertInput) (*model.TerminalProveMutationResponse, error) {
	rs := make([]*model1.TerminalProve, 0)
	for _, object := range objects {
		v := &model1.TerminalProve{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.TerminalProve{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.TerminalProveMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertTerminalProveOne(ctx context.Context, object model.TerminalProveInsertInput) (*model1.TerminalProve, error) {
	rs := &model1.TerminalProve{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.TerminalProve{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateTerminalProve(ctx context.Context, inc *model.TerminalProveIncInput, set *model.TerminalProveSetInput, where model.TerminalProveBoolExp) (*model.TerminalProveMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalProve{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.TerminalProveMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.TerminalProve
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
	return &model.TerminalProveMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateTerminalProveByPk(ctx context.Context, inc *model.TerminalProveIncInput, set *model.TerminalProveSetInput, Id int64) (*model1.TerminalProve, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.TerminalProve{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.TerminalProve
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *mutationResolver) UpdateTerminalProveByUnionPk(ctx context.Context, inc *model.TerminalProveIncInput, set *model.TerminalProveSetInput, unionId string) (*model1.TerminalProve, error) {
	var rs model1.TerminalProve
	tx := db.DB.Where(rs.UnionPrimaryColumnName()+" = ?", unionId)
	qt := util.NewQueryTranslator(tx, &model1.TerminalProve{})
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

func (r *queryResolver) TerminalProve(ctx context.Context, distinctOn []model.TerminalProveSelectColumn, limit *int, offset *int, orderBy []*model.TerminalProveOrderBy, where *model.TerminalProveBoolExp) ([]*model1.TerminalProve, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalProve{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.TerminalProve
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) TerminalProveAggregate(ctx context.Context, distinctOn []model.TerminalProveSelectColumn, limit *int, offset *int, orderBy []*model.TerminalProveOrderBy, where *model.TerminalProveBoolExp) (*model.TerminalProveAggregate, error) {
	var rs model.TerminalProveAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.TerminalProve{})
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

func (r *queryResolver) TerminalProveByPk(ctx context.Context, Id int64) (*model1.TerminalProve, error) {
	var rs model1.TerminalProve
	tx := db.DB.Model(&model1.TerminalProve{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}

func (r *queryResolver) TerminalProveByUnionPk(ctx context.Context, unionId string) (*model1.TerminalProve, error) {
	var rs model1.TerminalProve
	tx := db.DB.Model(&model1.TerminalProve{}).Select(util.GetTopPreloads(ctx)).Where(rs.UnionPrimaryColumnName()+" = ?", unionId).First(&rs)

	err := tx.Error
	return &rs, err
}
