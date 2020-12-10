package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/vehiclelocation/graph/generated"
	"VehicleSupervision/internal/modules/vehiclelocation/graph/model"
	model1 "VehicleSupervision/internal/modules/vehiclelocation/model"
	"VehicleSupervision/pkg/graphql/util"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteVehicleLocationHis(ctx context.Context, where model.VehicleLocationHisBoolExp) (*model.VehicleLocationHisMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleLocationHis{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.VehicleLocationHis
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
	return &model.VehicleLocationHisMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteVehicleLocationHisByPk(ctx context.Context, id int64) (*model1.VehicleLocationHis, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.VehicleLocationHis
	tx := db.DB.Model(&model1.VehicleLocationHis{})
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

func (r *mutationResolver) InsertVehicleLocationHis(ctx context.Context, objects []*model.VehicleLocationHisInsertInput, onConflict *model.VehicleLocationHisOnConflict) (*model.VehicleLocationHisMutationResponse, error) {
	rs := r.vehicleLocationHisInsertInputBatchConvert(objects)
	tx := db.DB.Model(&model1.VehicleLocationHis{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.VehicleLocationHisMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertVehicleLocationHisOne(ctx context.Context, object model.VehicleLocationHisInsertInput, onConflict *model.VehicleLocationHisOnConflict) (*model1.VehicleLocationHis, error) {
	rs := r.vehicleLocationHisInsertInputConvert(&object)
	tx := db.DB.Model(&model1.VehicleLocationHis{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateVehicleLocationHis(ctx context.Context, inc *model.VehicleLocationHisIncInput, set *model.VehicleLocationHisSetInput, where model.VehicleLocationHisBoolExp) (*model.VehicleLocationHisMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleLocationHis{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.VehicleLocationHisMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.VehicleLocationHisMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateVehicleLocationHisByPk(ctx context.Context, inc *model.VehicleLocationHisIncInput, set *model.VehicleLocationHisSetInput, pkColumns model.VehicleLocationHisPkColumnsInput) (*model1.VehicleLocationHis, error) {
	tx := db.DB.Where("id = ?", pkColumns.ID)
	qt := util.NewQueryTranslator(tx, &model1.VehicleLocationHis{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.VehicleLocationHis
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) VehicleLocationHis(ctx context.Context, distinctOn []model.VehicleLocationHisSelectColumn, limit *int, offset *int, orderBy []*model.VehicleLocationHisOrderBy, where *model.VehicleLocationHisBoolExp) ([]*model1.VehicleLocationHis, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleLocationHis{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.VehicleLocationHis
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) VehicleLocationHisAggregate(ctx context.Context, distinctOn []model.VehicleLocationHisSelectColumn, limit *int, offset *int, orderBy []*model.VehicleLocationHisOrderBy, where *model.VehicleLocationHisBoolExp) (*model.VehicleLocationHisAggregate, error) {
	var rs model.VehicleLocationHisAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.VehicleLocationHis{})
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

func (r *queryResolver) VehicleLocationHisByPk(ctx context.Context, id int64) (*model1.VehicleLocationHis, error) {
	var rs model1.VehicleLocationHis
	tx := db.DB.Model(&model1.VehicleLocationHis{}).First(&rs, id)
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
