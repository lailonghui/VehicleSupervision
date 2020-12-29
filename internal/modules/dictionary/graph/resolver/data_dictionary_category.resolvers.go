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

func (r *mutationResolver) DeleteDataDictionaryCategory(ctx context.Context, where model.DataDictionaryCategoryBoolExp) (*model.DataDictionaryCategoryMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DataDictionaryCategory{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.DataDictionaryCategory
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
	return &model.DataDictionaryCategoryMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteDataDictionaryCategoryByPk(ctx context.Context, Id int64) (*model1.DataDictionaryCategory, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.DataDictionaryCategory
	tx := db.DB.Model(&model1.DataDictionaryCategory{})
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

func (r *mutationResolver) InsertDataDictionaryCategory(ctx context.Context, objects []*model.DataDictionaryCategoryInsertInput) (*model.DataDictionaryCategoryMutationResponse, error) {
	rs := make([]*model1.DataDictionaryCategory, 0)
	for _, object := range objects {
		v := &model1.DataDictionaryCategory{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.DataDictionaryCategory{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.DataDictionaryCategoryMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertDataDictionaryCategoryOne(ctx context.Context, object model.DataDictionaryCategoryInsertInput) (*model1.DataDictionaryCategory, error) {
	rs := &model1.DataDictionaryCategory{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.DataDictionaryCategory{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateDataDictionaryCategory(ctx context.Context, inc *model.DataDictionaryCategoryIncInput, set *model.DataDictionaryCategorySetInput, where model.DataDictionaryCategoryBoolExp) (*model.DataDictionaryCategoryMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DataDictionaryCategory{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.DataDictionaryCategoryMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.DataDictionaryCategory
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
	return &model.DataDictionaryCategoryMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateDataDictionaryCategoryByPk(ctx context.Context, inc *model.DataDictionaryCategoryIncInput, set *model.DataDictionaryCategorySetInput, Id int64) (*model1.DataDictionaryCategory, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.DataDictionaryCategory{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.DataDictionaryCategory
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) DataDictionaryCategory(ctx context.Context, distinctOn []model.DataDictionaryCategorySelectColumn, limit *int, offset *int, orderBy []*model.DataDictionaryCategoryOrderBy, where *model.DataDictionaryCategoryBoolExp) ([]*model1.DataDictionaryCategory, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DataDictionaryCategory{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.DataDictionaryCategory
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) DataDictionaryCategoryAggregate(ctx context.Context, distinctOn []model.DataDictionaryCategorySelectColumn, limit *int, offset *int, orderBy []*model.DataDictionaryCategoryOrderBy, where *model.DataDictionaryCategoryBoolExp) (*model.DataDictionaryCategoryAggregate, error) {
	var rs model.DataDictionaryCategoryAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.DataDictionaryCategory{})
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

func (r *queryResolver) DataDictionaryCategoryByPk(ctx context.Context, Id int64) (*model1.DataDictionaryCategory, error) {
	var rs model1.DataDictionaryCategory
	tx := db.DB.Model(&model1.DataDictionaryCategory{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}
