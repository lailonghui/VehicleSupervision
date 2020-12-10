package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/ridehailing/graph/generated"
	"VehicleSupervision/internal/modules/ridehailing/graph/model"
	model1 "VehicleSupervision/internal/modules/ridehailing/model"
	"VehicleSupervision/pkg/graphql/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteRideHailingDriver(ctx context.Context, where model.RideHailingDriverBoolExp) (*model.RideHailingDriverMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.RideHailingDriver{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.RideHailingDriver
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
	return &model.RideHailingDriverMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteRideHailingDriverByPk(ctx context.Context, id int64) (*model1.RideHailingDriver, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.RideHailingDriver
	tx := db.DB.Model(&model1.RideHailingDriver{})
	if len(preloads) > 0 {
		// 如果请求的字段不为空，则先查询一遍数据库
		tx = tx.Select(preloads).Where("id = ?", id).First(&rs)
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

func (r *mutationResolver) InsertRideHailingDriver(ctx context.Context, objects []*model.RideHailingDriverInsertInput, onConflict *model.RideHailingDriverOnConflict) (*model.RideHailingDriverMutationResponse, error) {
	rs := rideHailingDriverInsertInputBatchConvert(objects)
	tx := db.DB.Model(&model1.RideHailingDriver{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.RideHailingDriverMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertRideHailingDriverOne(ctx context.Context, object model.RideHailingDriverInsertInput, onConflict *model.RideHailingDriverOnConflict) (*model1.RideHailingDriver, error) {
	rs := rideHailingDriverInsertInputConvert(&object)
	tx := db.DB.Model(&model1.RideHailingDriver{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateRideHailingDriver(ctx context.Context, inc *model.RideHailingDriverIncInput, set *model.RideHailingDriverSetInput, where model.RideHailingDriverBoolExp) (*model.RideHailingDriverMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.RideHailingDriver{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.RideHailingDriverMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.RideHailingDriverMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateRideHailingDriverByPk(ctx context.Context, inc *model.RideHailingDriverIncInput, set *model.RideHailingDriverSetInput, pkColumns model.RideHailingDriverPkColumnsInput) (*model1.RideHailingDriver, error) {
	tx := db.DB.Where("id = ?", pkColumns.ID)
	qt := util.NewQueryTranslator(tx, &model1.RideHailingDriver{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.RideHailingDriver
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) RideHailingDriver(ctx context.Context, distinctOn []model.RideHailingDriverSelectColumn, limit *int, offset *int, orderBy []*model.RideHailingDriverOrderBy, where *model.RideHailingDriverBoolExp) ([]*model1.RideHailingDriver, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.RideHailingDriver{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.RideHailingDriver
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) RideHailingDriverAggregate(ctx context.Context, distinctOn []model.RideHailingDriverSelectColumn, limit *int, offset *int, orderBy []*model.RideHailingDriverOrderBy, where *model.RideHailingDriverBoolExp) (*model.RideHailingDriverAggregate, error) {
	var rs model.RideHailingDriverAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.RideHailingDriver{})
	tx, err := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Aggregate(&rs, ctx)
	if err != nil {
		return nil, err
	}
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &rs, nil
}

func (r *queryResolver) RideHailingDriverByPk(ctx context.Context, id int64) (*model1.RideHailingDriver, error) {
	var rs model1.RideHailingDriver
	tx := db.DB.Model(&model1.RideHailingDriver{}).First(&rs, id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
