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

func (r *mutationResolver) DeleteEcdFileCheckHis(ctx context.Context, where model.EcdFileCheckHisBoolExp) (*model.EcdFileCheckHisMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EcdFileCheckHis{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.EcdFileCheckHis
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
	return &model.EcdFileCheckHisMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteEcdFileCheckHisByPk(ctx context.Context, Id int64) (*model1.EcdFileCheckHis, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.EcdFileCheckHis
	tx := db.DB.Model(&model1.EcdFileCheckHis{})
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

func (r *mutationResolver) InsertEcdFileCheckHis(ctx context.Context, objects []*model.EcdFileCheckHisInsertInput) (*model.EcdFileCheckHisMutationResponse, error) {
	rs := make([]*model1.EcdFileCheckHis, 0)
	for _, object := range objects {
		v := &model1.EcdFileCheckHis{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.EcdFileCheckHis{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.EcdFileCheckHisMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertEcdFileCheckHisOne(ctx context.Context, object model.EcdFileCheckHisInsertInput) (*model1.EcdFileCheckHis, error) {
	rs := &model1.EcdFileCheckHis{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.EcdFileCheckHis{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateEcdFileCheckHis(ctx context.Context, inc *model.EcdFileCheckHisIncInput, set *model.EcdFileCheckHisSetInput, where model.EcdFileCheckHisBoolExp) (*model.EcdFileCheckHisMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EcdFileCheckHis{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.EcdFileCheckHisMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.EcdFileCheckHis
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
	return &model.EcdFileCheckHisMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateEcdFileCheckHisByPk(ctx context.Context, inc *model.EcdFileCheckHisIncInput, set *model.EcdFileCheckHisSetInput, Id int64) (*model1.EcdFileCheckHis, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.EcdFileCheckHis{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.EcdFileCheckHis
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) EcdFileCheckHis(ctx context.Context, distinctOn []model.EcdFileCheckHisSelectColumn, limit *int, offset *int, orderBy []*model.EcdFileCheckHisOrderBy, where *model.EcdFileCheckHisBoolExp) ([]*model1.EcdFileCheckHis, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EcdFileCheckHis{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.EcdFileCheckHis
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) EcdFileCheckHisAggregate(ctx context.Context, distinctOn []model.EcdFileCheckHisSelectColumn, limit *int, offset *int, orderBy []*model.EcdFileCheckHisOrderBy, where *model.EcdFileCheckHisBoolExp) (*model.EcdFileCheckHisAggregate, error) {
	var rs model.EcdFileCheckHisAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.EcdFileCheckHis{})
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

func (r *queryResolver) EcdFileCheckHisByPk(ctx context.Context, Id int64) (*model1.EcdFileCheckHis, error) {
	var rs model1.EcdFileCheckHis
	tx := db.DB.Model(&model1.EcdFileCheckHis{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}
