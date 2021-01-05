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

func (r *mutationResolver) DeleteTerminalCheck(ctx context.Context, where model.TerminalCheckBoolExp) (*model.TerminalCheckMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalCheck{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.TerminalCheck
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
	return &model.TerminalCheckMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteTerminalCheckByPk(ctx context.Context, Id int64) (*model1.TerminalCheck, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.TerminalCheck
	tx := db.DB.Model(&model1.TerminalCheck{})
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

func (r *mutationResolver) DeleteTerminalCheckByUnionPk(ctx context.Context, unionId string) (*model1.TerminalCheck, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.TerminalCheck
	tx := db.DB.Model(&model1.TerminalCheck{})
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

func (r *mutationResolver) InsertTerminalCheck(ctx context.Context, objects []*model.TerminalCheckInsertInput) (*model.TerminalCheckMutationResponse, error) {
	rs := make([]*model1.TerminalCheck, 0)
	for _, object := range objects {
		v := &model1.TerminalCheck{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.TerminalCheck{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.TerminalCheckMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertTerminalCheckOne(ctx context.Context, object model.TerminalCheckInsertInput) (*model1.TerminalCheck, error) {
	rs := &model1.TerminalCheck{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.TerminalCheck{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateTerminalCheck(ctx context.Context, inc *model.TerminalCheckIncInput, set *model.TerminalCheckSetInput, where model.TerminalCheckBoolExp) (*model.TerminalCheckMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalCheck{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.TerminalCheckMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.TerminalCheck
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
	return &model.TerminalCheckMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateTerminalCheckByPk(ctx context.Context, inc *model.TerminalCheckIncInput, set *model.TerminalCheckSetInput, Id int64) (*model1.TerminalCheck, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.TerminalCheck{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.TerminalCheck
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *mutationResolver) UpdateTerminalCheckByUnionPk(ctx context.Context, inc *model.TerminalCheckIncInput, set *model.TerminalCheckSetInput, unionId string) (*model1.TerminalCheck, error) {
	var rs model1.TerminalCheck
	tx := db.DB.Where(rs.UnionPrimaryColumnName()+" = ?", unionId)
	qt := util.NewQueryTranslator(tx, &model1.TerminalCheck{})
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

func (r *queryResolver) TerminalCheck(ctx context.Context, distinctOn []model.TerminalCheckSelectColumn, limit *int, offset *int, orderBy []*model.TerminalCheckOrderBy, where *model.TerminalCheckBoolExp) ([]*model1.TerminalCheck, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalCheck{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.TerminalCheck
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) TerminalCheckAggregate(ctx context.Context, distinctOn []model.TerminalCheckSelectColumn, limit *int, offset *int, orderBy []*model.TerminalCheckOrderBy, where *model.TerminalCheckBoolExp) (*model.TerminalCheckAggregate, error) {
	var rs model.TerminalCheckAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.TerminalCheck{})
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

func (r *queryResolver) TerminalCheckByPk(ctx context.Context, Id int64) (*model1.TerminalCheck, error) {
	var rs model1.TerminalCheck
	tx := db.DB.Model(&model1.TerminalCheck{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}

func (r *queryResolver) TerminalCheckByUnionPk(ctx context.Context, unionId string) (*model1.TerminalCheck, error) {
	var rs model1.TerminalCheck
	tx := db.DB.Model(&model1.TerminalCheck{}).Select(util.GetTopPreloads(ctx)).Where(rs.UnionPrimaryColumnName()+" = ?", unionId).First(&rs)

	err := tx.Error
	return &rs, err
}
