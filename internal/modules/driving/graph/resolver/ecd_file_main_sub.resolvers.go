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

func (r *mutationResolver) DeleteEcdFileMainSub(ctx context.Context, where model.EcdFileMainSubBoolExp) (*model.EcdFileMainSubMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EcdFileMainSub{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.EcdFileMainSub
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
	return &model.EcdFileMainSubMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteEcdFileMainSubByPk(ctx context.Context, Id int64) (*model1.EcdFileMainSub, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.EcdFileMainSub
	tx := db.DB.Model(&model1.EcdFileMainSub{})
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

func (r *mutationResolver) InsertEcdFileMainSub(ctx context.Context, objects []*model.EcdFileMainSubInsertInput) (*model.EcdFileMainSubMutationResponse, error) {
	rs := make([]*model1.EcdFileMainSub, 0)
	for _, object := range objects {
		v := &model1.EcdFileMainSub{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.EcdFileMainSub{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.EcdFileMainSubMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertEcdFileMainSubOne(ctx context.Context, object model.EcdFileMainSubInsertInput) (*model1.EcdFileMainSub, error) {
	rs := &model1.EcdFileMainSub{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.EcdFileMainSub{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateEcdFileMainSub(ctx context.Context, inc *model.EcdFileMainSubIncInput, set *model.EcdFileMainSubSetInput, where model.EcdFileMainSubBoolExp) (*model.EcdFileMainSubMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EcdFileMainSub{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.EcdFileMainSubMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.EcdFileMainSub
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
	return &model.EcdFileMainSubMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateEcdFileMainSubByPk(ctx context.Context, inc *model.EcdFileMainSubIncInput, set *model.EcdFileMainSubSetInput, Id int64) (*model1.EcdFileMainSub, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.EcdFileMainSub{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.EcdFileMainSub
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) EcdFileMainSub(ctx context.Context, distinctOn []model.EcdFileMainSubSelectColumn, limit *int, offset *int, orderBy []*model.EcdFileMainSubOrderBy, where *model.EcdFileMainSubBoolExp) ([]*model1.EcdFileMainSub, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EcdFileMainSub{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.EcdFileMainSub
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) EcdFileMainSubAggregate(ctx context.Context, distinctOn []model.EcdFileMainSubSelectColumn, limit *int, offset *int, orderBy []*model.EcdFileMainSubOrderBy, where *model.EcdFileMainSubBoolExp) (*model.EcdFileMainSubAggregate, error) {
	var rs model.EcdFileMainSubAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.EcdFileMainSub{})
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

func (r *queryResolver) EcdFileMainSubByPk(ctx context.Context, Id int64) (*model1.EcdFileMainSub, error) {
	var rs model1.EcdFileMainSub
	tx := db.DB.Model(&model1.EcdFileMainSub{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}