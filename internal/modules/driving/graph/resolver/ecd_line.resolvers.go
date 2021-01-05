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

func (r *mutationResolver) DeleteEcdLine(ctx context.Context, where model.EcdLineBoolExp) (*model.EcdLineMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EcdLine{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.EcdLine
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
	return &model.EcdLineMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteEcdLineByPk(ctx context.Context, Id int64) (*model1.EcdLine, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.EcdLine
	tx := db.DB.Model(&model1.EcdLine{})
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

func (r *mutationResolver) DeleteEcdLineByUnionPk(ctx context.Context, unionId string) (*model1.EcdLine, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.EcdLine
	tx := db.DB.Model(&model1.EcdLine{})
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

func (r *mutationResolver) InsertEcdLine(ctx context.Context, objects []*model.EcdLineInsertInput) (*model.EcdLineMutationResponse, error) {
	rs := make([]*model1.EcdLine, 0)
	for _, object := range objects {
		v := &model1.EcdLine{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.EcdLine{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.EcdLineMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertEcdLineOne(ctx context.Context, object model.EcdLineInsertInput) (*model1.EcdLine, error) {
	rs := &model1.EcdLine{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.EcdLine{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateEcdLine(ctx context.Context, inc *model.EcdLineIncInput, set *model.EcdLineSetInput, where model.EcdLineBoolExp) (*model.EcdLineMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EcdLine{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.EcdLineMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.EcdLine
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
	return &model.EcdLineMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateEcdLineByPk(ctx context.Context, inc *model.EcdLineIncInput, set *model.EcdLineSetInput, Id int64) (*model1.EcdLine, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.EcdLine{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.EcdLine
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *mutationResolver) UpdateEcdLineByUnionPk(ctx context.Context, inc *model.EcdLineIncInput, set *model.EcdLineSetInput, unionId string) (*model1.EcdLine, error) {
	var rs model1.EcdLine
	tx := db.DB.Where(rs.UnionPrimaryColumnName()+" = ?", unionId)
	qt := util.NewQueryTranslator(tx, &model1.EcdLine{})
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

func (r *queryResolver) EcdLine(ctx context.Context, distinctOn []model.EcdLineSelectColumn, limit *int, offset *int, orderBy []*model.EcdLineOrderBy, where *model.EcdLineBoolExp) ([]*model1.EcdLine, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EcdLine{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.EcdLine
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) EcdLineAggregate(ctx context.Context, distinctOn []model.EcdLineSelectColumn, limit *int, offset *int, orderBy []*model.EcdLineOrderBy, where *model.EcdLineBoolExp) (*model.EcdLineAggregate, error) {
	var rs model.EcdLineAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.EcdLine{})
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

func (r *queryResolver) EcdLineByPk(ctx context.Context, Id int64) (*model1.EcdLine, error) {
	var rs model1.EcdLine
	tx := db.DB.Model(&model1.EcdLine{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}

func (r *queryResolver) EcdLineByUnionPk(ctx context.Context, unionId string) (*model1.EcdLine, error) {
	var rs model1.EcdLine
	tx := db.DB.Model(&model1.EcdLine{}).Select(util.GetTopPreloads(ctx)).Where(rs.UnionPrimaryColumnName()+" = ?", unionId).First(&rs)

	err := tx.Error
	return &rs, err
}
