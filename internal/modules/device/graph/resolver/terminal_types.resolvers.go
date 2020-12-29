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

func (r *mutationResolver) DeleteTerminalTypes(ctx context.Context, where model.TerminalTypesBoolExp) (*model.TerminalTypesMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalTypes{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.TerminalTypes
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
	return &model.TerminalTypesMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteTerminalTypesByPk(ctx context.Context, Id int64) (*model1.TerminalTypes, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.TerminalTypes
	tx := db.DB.Model(&model1.TerminalTypes{})
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

func (r *mutationResolver) InsertTerminalTypes(ctx context.Context, objects []*model.TerminalTypesInsertInput) (*model.TerminalTypesMutationResponse, error) {
	rs := make([]*model1.TerminalTypes, 0)
	for _, object := range objects {
		v := &model1.TerminalTypes{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.TerminalTypes{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.TerminalTypesMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertTerminalTypesOne(ctx context.Context, object model.TerminalTypesInsertInput) (*model1.TerminalTypes, error) {
	rs := &model1.TerminalTypes{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.TerminalTypes{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateTerminalTypes(ctx context.Context, inc *model.TerminalTypesIncInput, set *model.TerminalTypesSetInput, where model.TerminalTypesBoolExp) (*model.TerminalTypesMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalTypes{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.TerminalTypesMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.TerminalTypes
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
	return &model.TerminalTypesMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateTerminalTypesByPk(ctx context.Context, inc *model.TerminalTypesIncInput, set *model.TerminalTypesSetInput, Id int64) (*model1.TerminalTypes, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.TerminalTypes{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.TerminalTypes
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) TerminalTypes(ctx context.Context, distinctOn []model.TerminalTypesSelectColumn, limit *int, offset *int, orderBy []*model.TerminalTypesOrderBy, where *model.TerminalTypesBoolExp) ([]*model1.TerminalTypes, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalTypes{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.TerminalTypes
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) TerminalTypesAggregate(ctx context.Context, distinctOn []model.TerminalTypesSelectColumn, limit *int, offset *int, orderBy []*model.TerminalTypesOrderBy, where *model.TerminalTypesBoolExp) (*model.TerminalTypesAggregate, error) {
	var rs model.TerminalTypesAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.TerminalTypes{})
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

func (r *queryResolver) TerminalTypesByPk(ctx context.Context, Id int64) (*model1.TerminalTypes, error) {
	var rs model1.TerminalTypes
	tx := db.DB.Model(&model1.TerminalTypes{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}
