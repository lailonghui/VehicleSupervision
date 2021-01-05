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

func (r *mutationResolver) DeleteTerminalParamType(ctx context.Context, where model.TerminalParamTypeBoolExp) (*model.TerminalParamTypeMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalParamType{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.TerminalParamType
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
	return &model.TerminalParamTypeMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteTerminalParamTypeByPk(ctx context.Context, Id int64) (*model1.TerminalParamType, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.TerminalParamType
	tx := db.DB.Model(&model1.TerminalParamType{})
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

func (r *mutationResolver) DeleteTerminalParamTypeByUnionPk(ctx context.Context, unionId string) (*model1.TerminalParamType, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.TerminalParamType
	tx := db.DB.Model(&model1.TerminalParamType{})
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

func (r *mutationResolver) InsertTerminalParamType(ctx context.Context, objects []*model.TerminalParamTypeInsertInput) (*model.TerminalParamTypeMutationResponse, error) {
	rs := make([]*model1.TerminalParamType, 0)
	for _, object := range objects {
		v := &model1.TerminalParamType{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.TerminalParamType{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.TerminalParamTypeMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertTerminalParamTypeOne(ctx context.Context, object model.TerminalParamTypeInsertInput) (*model1.TerminalParamType, error) {
	rs := &model1.TerminalParamType{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.TerminalParamType{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateTerminalParamType(ctx context.Context, inc *model.TerminalParamTypeIncInput, set *model.TerminalParamTypeSetInput, where model.TerminalParamTypeBoolExp) (*model.TerminalParamTypeMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalParamType{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.TerminalParamTypeMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.TerminalParamType
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
	return &model.TerminalParamTypeMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateTerminalParamTypeByPk(ctx context.Context, inc *model.TerminalParamTypeIncInput, set *model.TerminalParamTypeSetInput, Id int64) (*model1.TerminalParamType, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.TerminalParamType{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.TerminalParamType
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *mutationResolver) UpdateTerminalParamTypeByUnionPk(ctx context.Context, inc *model.TerminalParamTypeIncInput, set *model.TerminalParamTypeSetInput, unionId string) (*model1.TerminalParamType, error) {
	var rs model1.TerminalParamType
	tx := db.DB.Where(rs.UnionPrimaryColumnName()+" = ?", unionId)
	qt := util.NewQueryTranslator(tx, &model1.TerminalParamType{})
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

func (r *queryResolver) TerminalParamType(ctx context.Context, distinctOn []model.TerminalParamTypeSelectColumn, limit *int, offset *int, orderBy []*model.TerminalParamTypeOrderBy, where *model.TerminalParamTypeBoolExp) ([]*model1.TerminalParamType, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalParamType{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.TerminalParamType
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) TerminalParamTypeAggregate(ctx context.Context, distinctOn []model.TerminalParamTypeSelectColumn, limit *int, offset *int, orderBy []*model.TerminalParamTypeOrderBy, where *model.TerminalParamTypeBoolExp) (*model.TerminalParamTypeAggregate, error) {
	var rs model.TerminalParamTypeAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.TerminalParamType{})
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

func (r *queryResolver) TerminalParamTypeByPk(ctx context.Context, Id int64) (*model1.TerminalParamType, error) {
	var rs model1.TerminalParamType
	tx := db.DB.Model(&model1.TerminalParamType{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}

func (r *queryResolver) TerminalParamTypeByUnionPk(ctx context.Context, unionId string) (*model1.TerminalParamType, error) {
	var rs model1.TerminalParamType
	tx := db.DB.Model(&model1.TerminalParamType{}).Select(util.GetTopPreloads(ctx)).Where(rs.UnionPrimaryColumnName()+" = ?", unionId).First(&rs)

	err := tx.Error
	return &rs, err
}
