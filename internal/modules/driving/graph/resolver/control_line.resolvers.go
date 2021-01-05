package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/driving/graph/model"
	model1 "VehicleSupervision/internal/modules/driving/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteControlLine(ctx context.Context, where model.ControlLineBoolExp) (*model.ControlLineMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.ControlLine{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.ControlLine
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
	return &model.ControlLineMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteControlLineByPk(ctx context.Context, Id int64) (*model1.ControlLine, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.ControlLine
	tx := db.DB.Model(&model1.ControlLine{})
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

func (r *mutationResolver) DeleteControlLineByUnionPk(ctx context.Context, unionId string) (*model1.ControlLine, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.ControlLine
	tx := db.DB.Model(&model1.ControlLine{})
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

func (r *mutationResolver) InsertControlLine(ctx context.Context, objects []*model.ControlLineInsertInput) (*model.ControlLineMutationResponse, error) {
	rs := make([]*model1.ControlLine, 0)
	for _, object := range objects {
		v := &model1.ControlLine{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.ControlLine{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.ControlLineMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertControlLineOne(ctx context.Context, object model.ControlLineInsertInput) (*model1.ControlLine, error) {
	rs := &model1.ControlLine{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.ControlLine{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateControlLine(ctx context.Context, inc *model.ControlLineIncInput, set *model.ControlLineSetInput, where model.ControlLineBoolExp) (*model.ControlLineMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.ControlLine{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.ControlLineMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.ControlLine
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
	return &model.ControlLineMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateControlLineByPk(ctx context.Context, inc *model.ControlLineIncInput, set *model.ControlLineSetInput, Id int64) (*model1.ControlLine, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.ControlLine{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.ControlLine
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *mutationResolver) UpdateControlLineByUnionPk(ctx context.Context, inc *model.ControlLineIncInput, set *model.ControlLineSetInput, unionId string) (*model1.ControlLine, error) {
	var rs model1.ControlLine
	tx := db.DB.Where(rs.UnionPrimaryColumnName()+" = ?", unionId)
	qt := util.NewQueryTranslator(tx, &model1.ControlLine{})
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

func (r *queryResolver) ControlLine(ctx context.Context, distinctOn []model.ControlLineSelectColumn, limit *int, offset *int, orderBy []*model.ControlLineOrderBy, where *model.ControlLineBoolExp) ([]*model1.ControlLine, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.ControlLine{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.ControlLine
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) ControlLineAggregate(ctx context.Context, distinctOn []model.ControlLineSelectColumn, limit *int, offset *int, orderBy []*model.ControlLineOrderBy, where *model.ControlLineBoolExp) (*model.ControlLineAggregate, error) {
	var rs model.ControlLineAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.ControlLine{})
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

func (r *queryResolver) ControlLineByPk(ctx context.Context, Id int64) (*model1.ControlLine, error) {
	var rs model1.ControlLine
	tx := db.DB.Model(&model1.ControlLine{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}

func (r *queryResolver) ControlLineByUnionPk(ctx context.Context, unionId string) (*model1.ControlLine, error) {
	var rs model1.ControlLine
	tx := db.DB.Model(&model1.ControlLine{}).Select(util.GetTopPreloads(ctx)).Where(rs.UnionPrimaryColumnName()+" = ?", unionId).First(&rs)

	err := tx.Error
	return &rs, err
}
