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

func (r *mutationResolver) DeleteVehicleEnterpriseChangeReview(ctx context.Context, where model.VehicleEnterpriseChangeReviewBoolExp) (*model.VehicleEnterpriseChangeReviewMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleEnterpriseChangeReview{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.VehicleEnterpriseChangeReview
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
	return &model.VehicleEnterpriseChangeReviewMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteVehicleEnterpriseChangeReviewByPk(ctx context.Context, Id int64) (*model1.VehicleEnterpriseChangeReview, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.VehicleEnterpriseChangeReview
	tx := db.DB.Model(&model1.VehicleEnterpriseChangeReview{})
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

func (r *mutationResolver) InsertVehicleEnterpriseChangeReview(ctx context.Context, objects []*model.VehicleEnterpriseChangeReviewInsertInput) (*model.VehicleEnterpriseChangeReviewMutationResponse, error) {
	rs := []*model1.VehicleEnterpriseChangeReview{}
	for _, object := range objects {
		v := &model1.VehicleEnterpriseChangeReview{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.VehicleEnterpriseChangeReview{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.VehicleEnterpriseChangeReviewMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertVehicleEnterpriseChangeReviewOne(ctx context.Context, object model.VehicleEnterpriseChangeReviewInsertInput) (*model1.VehicleEnterpriseChangeReview, error) {
	rs := &model1.VehicleEnterpriseChangeReview{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.VehicleEnterpriseChangeReview{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateVehicleEnterpriseChangeReview(ctx context.Context, inc *model.VehicleEnterpriseChangeReviewIncInput, set *model.VehicleEnterpriseChangeReviewSetInput, where model.VehicleEnterpriseChangeReviewBoolExp) (*model.VehicleEnterpriseChangeReviewMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleEnterpriseChangeReview{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.VehicleEnterpriseChangeReviewMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.VehicleEnterpriseChangeReviewMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateVehicleEnterpriseChangeReviewByPk(ctx context.Context, inc *model.VehicleEnterpriseChangeReviewIncInput, set *model.VehicleEnterpriseChangeReviewSetInput, Id int64) (*model1.VehicleEnterpriseChangeReview, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.VehicleEnterpriseChangeReview{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.VehicleEnterpriseChangeReview
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) VehicleEnterpriseChangeReview(ctx context.Context, distinctOn []model.VehicleEnterpriseChangeReviewSelectColumn, limit *int, offset *int, orderBy []*model.VehicleEnterpriseChangeReviewOrderBy, where *model.VehicleEnterpriseChangeReviewBoolExp) ([]*model1.VehicleEnterpriseChangeReview, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleEnterpriseChangeReview{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.VehicleEnterpriseChangeReview
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) VehicleEnterpriseChangeReviewAggregate(ctx context.Context, distinctOn []model.VehicleEnterpriseChangeReviewSelectColumn, limit *int, offset *int, orderBy []*model.VehicleEnterpriseChangeReviewOrderBy, where *model.VehicleEnterpriseChangeReviewBoolExp) (*model.VehicleEnterpriseChangeReviewAggregate, error) {
	var rs model.VehicleEnterpriseChangeReviewAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.VehicleEnterpriseChangeReview{})
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

func (r *queryResolver) VehicleEnterpriseChangeReviewByPk(ctx context.Context, Id int64) (*model1.VehicleEnterpriseChangeReview, error) {
	var rs model1.VehicleEnterpriseChangeReview
	tx := db.DB.Model(&model1.VehicleEnterpriseChangeReview{}).First(&rs, Id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
