package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/dictionary/graph/model"
	model1 "VehicleSupervision/internal/modules/dictionary/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteDataDictionary(ctx context.Context, where model.DataDictionaryBoolExp) (*model.DataDictionaryMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DataDictionary{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.DataDictionary
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
	return &model.DataDictionaryMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteDataDictionaryByPk(ctx context.Context, Id int64) (*model1.DataDictionary, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.DataDictionary
	tx := db.DB.Model(&model1.DataDictionary{})
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

func (r *mutationResolver) DeleteDataDictionaryByUnionPk(ctx context.Context, unionId string) (*model1.DataDictionary, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.DataDictionary
	tx := db.DB.Model(&model1.DataDictionary{})
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

func (r *mutationResolver) InsertDataDictionary(ctx context.Context, objects []*model.DataDictionaryInsertInput) (*model.DataDictionaryMutationResponse, error) {
	rs := make([]*model1.DataDictionary, 0)
	for _, object := range objects {
		v := &model1.DataDictionary{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.DataDictionary{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.DataDictionaryMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertDataDictionaryOne(ctx context.Context, object model.DataDictionaryInsertInput) (*model1.DataDictionary, error) {
	rs := &model1.DataDictionary{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.DataDictionary{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateDataDictionary(ctx context.Context, inc *model.DataDictionaryIncInput, set *model.DataDictionarySetInput, where model.DataDictionaryBoolExp) (*model.DataDictionaryMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DataDictionary{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.DataDictionaryMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.DataDictionary
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
	return &model.DataDictionaryMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateDataDictionaryByPk(ctx context.Context, inc *model.DataDictionaryIncInput, set *model.DataDictionarySetInput, Id int64) (*model1.DataDictionary, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.DataDictionary{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.DataDictionary
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *mutationResolver) UpdateDataDictionaryByUnionPk(ctx context.Context, inc *model.DataDictionaryIncInput, set *model.DataDictionarySetInput, unionId string) (*model1.DataDictionary, error) {
	var rs model1.DataDictionary
	tx := db.DB.Where(rs.UnionPrimaryColumnName()+" = ?", unionId)
	qt := util.NewQueryTranslator(tx, &model1.DataDictionary{})
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

func (r *queryResolver) DataDictionary(ctx context.Context, distinctOn []model.DataDictionarySelectColumn, limit *int, offset *int, orderBy []*model.DataDictionaryOrderBy, where *model.DataDictionaryBoolExp) ([]*model1.DataDictionary, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DataDictionary{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.DataDictionary
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) DataDictionaryAggregate(ctx context.Context, distinctOn []model.DataDictionarySelectColumn, limit *int, offset *int, orderBy []*model.DataDictionaryOrderBy, where *model.DataDictionaryBoolExp) (*model.DataDictionaryAggregate, error) {
	var rs model.DataDictionaryAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.DataDictionary{})
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

func (r *queryResolver) DataDictionaryByPk(ctx context.Context, Id int64) (*model1.DataDictionary, error) {
	var rs model1.DataDictionary
	tx := db.DB.Model(&model1.DataDictionary{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}

func (r *queryResolver) DataDictionaryByUnionPk(ctx context.Context, unionId string) (*model1.DataDictionary, error) {
	var rs model1.DataDictionary
	tx := db.DB.Model(&model1.DataDictionary{}).Select(util.GetTopPreloads(ctx)).Where(rs.UnionPrimaryColumnName()+" = ?", unionId).First(&rs)

	err := tx.Error
	return &rs, err
}
