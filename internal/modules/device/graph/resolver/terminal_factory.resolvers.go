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

func (r *mutationResolver) DeleteTerminalFactory(ctx context.Context, where model.TerminalFactoryBoolExp) (*model.TerminalFactoryMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalFactory{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.TerminalFactory
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
	return &model.TerminalFactoryMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteTerminalFactoryByPk(ctx context.Context, Id int64) (*model1.TerminalFactory, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.TerminalFactory
	tx := db.DB.Model(&model1.TerminalFactory{})
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

func (r *mutationResolver) DeleteTerminalFactoryByUnionPk(ctx context.Context, unionId string) (*model1.TerminalFactory, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.TerminalFactory
	tx := db.DB.Model(&model1.TerminalFactory{})
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

func (r *mutationResolver) InsertTerminalFactory(ctx context.Context, objects []*model.TerminalFactoryInsertInput) (*model.TerminalFactoryMutationResponse, error) {
	rs := make([]*model1.TerminalFactory, 0)
	for _, object := range objects {
		v := &model1.TerminalFactory{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.TerminalFactory{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.TerminalFactoryMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertTerminalFactoryOne(ctx context.Context, object model.TerminalFactoryInsertInput) (*model1.TerminalFactory, error) {
	rs := &model1.TerminalFactory{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.TerminalFactory{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateTerminalFactory(ctx context.Context, inc *model.TerminalFactoryIncInput, set *model.TerminalFactorySetInput, where model.TerminalFactoryBoolExp) (*model.TerminalFactoryMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalFactory{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.TerminalFactoryMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.TerminalFactory
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
	return &model.TerminalFactoryMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateTerminalFactoryByPk(ctx context.Context, inc *model.TerminalFactoryIncInput, set *model.TerminalFactorySetInput, Id int64) (*model1.TerminalFactory, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.TerminalFactory{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.TerminalFactory
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *mutationResolver) UpdateTerminalFactoryByUnionPk(ctx context.Context, inc *model.TerminalFactoryIncInput, set *model.TerminalFactorySetInput, unionId string) (*model1.TerminalFactory, error) {
	var rs model1.TerminalFactory
	tx := db.DB.Where(rs.UnionPrimaryColumnName()+" = ?", unionId)
	qt := util.NewQueryTranslator(tx, &model1.TerminalFactory{})
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

func (r *queryResolver) TerminalFactory(ctx context.Context, distinctOn []model.TerminalFactorySelectColumn, limit *int, offset *int, orderBy []*model.TerminalFactoryOrderBy, where *model.TerminalFactoryBoolExp) ([]*model1.TerminalFactory, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalFactory{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.TerminalFactory
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) TerminalFactoryAggregate(ctx context.Context, distinctOn []model.TerminalFactorySelectColumn, limit *int, offset *int, orderBy []*model.TerminalFactoryOrderBy, where *model.TerminalFactoryBoolExp) (*model.TerminalFactoryAggregate, error) {
	var rs model.TerminalFactoryAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.TerminalFactory{})
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

func (r *queryResolver) TerminalFactoryByPk(ctx context.Context, Id int64) (*model1.TerminalFactory, error) {
	var rs model1.TerminalFactory
	tx := db.DB.Model(&model1.TerminalFactory{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}

func (r *queryResolver) TerminalFactoryByUnionPk(ctx context.Context, unionId string) (*model1.TerminalFactory, error) {
	var rs model1.TerminalFactory
	tx := db.DB.Model(&model1.TerminalFactory{}).Select(util.GetTopPreloads(ctx)).Where(rs.UnionPrimaryColumnName()+" = ?", unionId).First(&rs)

	err := tx.Error
	return &rs, err
}
