package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/blacklist/graph/generated"
	"VehicleSupervision/internal/modules/blacklist/graph/model"
	model1 "VehicleSupervision/internal/modules/blacklist/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteDriverBlacklistApply(ctx context.Context, where model.DriverBlacklistApplyBoolExp) (*model.DriverBlacklistApplyMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DriverBlacklistApply{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.DriverBlacklistApply
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
	return &model.DriverBlacklistApplyMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteDriverBlacklistApplyByPk(ctx context.Context, Id int64) (*model1.DriverBlacklistApply, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.DriverBlacklistApply
	tx := db.DB.Model(&model1.DriverBlacklistApply{})
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

func (r *mutationResolver) InsertDriverBlacklistApply(ctx context.Context, objects []*model.DriverBlacklistApplyInsertInput) (*model.DriverBlacklistApplyMutationResponse, error) {
	rs := make([]*model1.DriverBlacklistApply, 0)
	for _, object := range objects {
		v := &model1.DriverBlacklistApply{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.DriverBlacklistApply{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.DriverBlacklistApplyMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertDriverBlacklistApplyOne(ctx context.Context, object model.DriverBlacklistApplyInsertInput) (*model1.DriverBlacklistApply, error) {
	rs := &model1.DriverBlacklistApply{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.DriverBlacklistApply{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateDriverBlacklistApply(ctx context.Context, inc *model.DriverBlacklistApplyIncInput, set *model.DriverBlacklistApplySetInput, where model.DriverBlacklistApplyBoolExp) (*model.DriverBlacklistApplyMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DriverBlacklistApply{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.DriverBlacklistApplyMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.DriverBlacklistApply
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
	return &model.DriverBlacklistApplyMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateDriverBlacklistApplyByPk(ctx context.Context, inc *model.DriverBlacklistApplyIncInput, set *model.DriverBlacklistApplySetInput, Id int64) (*model1.DriverBlacklistApply, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.DriverBlacklistApply{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.DriverBlacklistApply
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) DriverBlacklistApply(ctx context.Context, distinctOn []model.DriverBlacklistApplySelectColumn, limit *int, offset *int, orderBy []*model.DriverBlacklistApplyOrderBy, where *model.DriverBlacklistApplyBoolExp) ([]*model1.DriverBlacklistApply, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DriverBlacklistApply{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.DriverBlacklistApply
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) DriverBlacklistApplyAggregate(ctx context.Context, distinctOn []model.DriverBlacklistApplySelectColumn, limit *int, offset *int, orderBy []*model.DriverBlacklistApplyOrderBy, where *model.DriverBlacklistApplyBoolExp) (*model.DriverBlacklistApplyAggregate, error) {
	var rs model.DriverBlacklistApplyAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.DriverBlacklistApply{})
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

func (r *queryResolver) DriverBlacklistApplyByPk(ctx context.Context, Id int64) (*model1.DriverBlacklistApply, error) {
	var rs model1.DriverBlacklistApply
	tx := db.DB.Model(&model1.DriverBlacklistApply{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
