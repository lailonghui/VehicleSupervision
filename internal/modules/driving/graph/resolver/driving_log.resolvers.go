package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/driving/graph/generated"
	"VehicleSupervision/internal/modules/driving/graph/model"
	model1 "VehicleSupervision/internal/modules/driving/model"
	"VehicleSupervision/pkg/graphql/util"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteDrivingLog(ctx context.Context, where model.DrivingLogBoolExp) (*model.DrivingLogMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DrivingLog{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.DrivingLog
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
	return &model.DrivingLogMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteDrivingLogByPk(ctx context.Context, id int64) (*model1.DrivingLog, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.DrivingLog
	tx := db.DB.Model(&model1.DrivingLog{})
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

func (r *mutationResolver) InsertDrivingLog(ctx context.Context, objects []*model.DrivingLogInsertInput, onConflict *model.DrivingLogOnConflict) (*model.DrivingLogMutationResponse, error) {
	rs := r.drivingLogInsertInputBatchConvert(objects)
	tx := db.DB.Model(&model1.DrivingLog{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.DrivingLogMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertDrivingLogOne(ctx context.Context, object model.DrivingLogInsertInput, onConflict *model.DrivingLogOnConflict) (*model1.DrivingLog, error) {
	rs := r.drivingLogInsertInputConvert(&object)
	tx := db.DB.Model(&model1.DrivingLog{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateDrivingLog(ctx context.Context, inc *model.DrivingLogIncInput, set *model.DrivingLogSetInput, where model.DrivingLogBoolExp) (*model.DrivingLogMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DrivingLog{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.DrivingLogMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.DrivingLogMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateDrivingLogByPk(ctx context.Context, inc *model.DrivingLogIncInput, set *model.DrivingLogSetInput, pkColumns model.DrivingLogPkColumnsInput) (*model1.DrivingLog, error) {
	tx := db.DB.Where("id = ?", pkColumns.ID)
	qt := util.NewQueryTranslator(tx, &model1.DrivingLog{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.DrivingLog
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) DrivingLog(ctx context.Context, distinctOn []model.DrivingLogSelectColumn, limit *int, offset *int, orderBy []*model.DrivingLogOrderBy, where *model.DrivingLogBoolExp) ([]*model1.DrivingLog, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DrivingLog{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.DrivingLog
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) DrivingLogAggregate(ctx context.Context, distinctOn []model.DrivingLogSelectColumn, limit *int, offset *int, orderBy []*model.DrivingLogOrderBy, where *model.DrivingLogBoolExp) (*model.DrivingLogAggregate, error) {
	var rs model.DrivingLogAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.DrivingLog{})
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

func (r *queryResolver) DrivingLogByPk(ctx context.Context, id int64) (*model1.DrivingLog, error) {
	var rs model1.DrivingLog
	tx := db.DB.Model(&model1.DrivingLog{}).First(&rs, id)
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
