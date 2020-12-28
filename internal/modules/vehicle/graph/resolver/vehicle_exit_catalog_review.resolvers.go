package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/vehicle/graph/model"
	model1 "VehicleSupervision/internal/modules/vehicle/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteVehicleExitCatalogReview(ctx context.Context, where model.VehicleExitCatalogReviewBoolExp) (*model.VehicleExitCatalogReviewMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleExitCatalogReview{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.VehicleExitCatalogReview
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
	return &model.VehicleExitCatalogReviewMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteVehicleExitCatalogReviewByPk(ctx context.Context, Id int64) (*model1.VehicleExitCatalogReview, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.VehicleExitCatalogReview
	tx := db.DB.Model(&model1.VehicleExitCatalogReview{})
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

func (r *mutationResolver) InsertVehicleExitCatalogReview(ctx context.Context, objects []*model.VehicleExitCatalogReviewInsertInput) (*model.VehicleExitCatalogReviewMutationResponse, error) {
	rs := []*model1.VehicleExitCatalogReview{}
	for _, object := range objects {
		v := &model1.VehicleExitCatalogReview{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.VehicleExitCatalogReview{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.VehicleExitCatalogReviewMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertVehicleExitCatalogReviewOne(ctx context.Context, object model.VehicleExitCatalogReviewInsertInput) (*model1.VehicleExitCatalogReview, error) {
	rs := &model1.VehicleExitCatalogReview{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.VehicleExitCatalogReview{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateVehicleExitCatalogReview(ctx context.Context, inc *model.VehicleExitCatalogReviewIncInput, set *model.VehicleExitCatalogReviewSetInput, where model.VehicleExitCatalogReviewBoolExp) (*model.VehicleExitCatalogReviewMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleExitCatalogReview{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.VehicleExitCatalogReviewMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.VehicleExitCatalogReviewMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateVehicleExitCatalogReviewByPk(ctx context.Context, inc *model.VehicleExitCatalogReviewIncInput, set *model.VehicleExitCatalogReviewSetInput, Id int64) (*model1.VehicleExitCatalogReview, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.VehicleExitCatalogReview{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.VehicleExitCatalogReview
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) VehicleExitCatalogReview(ctx context.Context, distinctOn []model.VehicleExitCatalogReviewSelectColumn, limit *int, offset *int, orderBy []*model.VehicleExitCatalogReviewOrderBy, where *model.VehicleExitCatalogReviewBoolExp) ([]*model1.VehicleExitCatalogReview, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleExitCatalogReview{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.VehicleExitCatalogReview
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) VehicleExitCatalogReviewAggregate(ctx context.Context, distinctOn []model.VehicleExitCatalogReviewSelectColumn, limit *int, offset *int, orderBy []*model.VehicleExitCatalogReviewOrderBy, where *model.VehicleExitCatalogReviewBoolExp) (*model.VehicleExitCatalogReviewAggregate, error) {
	var rs model.VehicleExitCatalogReviewAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.VehicleExitCatalogReview{})
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

func (r *queryResolver) VehicleExitCatalogReviewByPk(ctx context.Context, Id int64) (*model1.VehicleExitCatalogReview, error) {
	var rs model1.VehicleExitCatalogReview
	tx := db.DB.Model(&model1.VehicleExitCatalogReview{}).First(&rs, Id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
