package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/vehicle_driver_separation/graph/model"
	model1 "VehicleSupervision/internal/modules/vehicle_driver_separation/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteCaseApprovalReviewCall(ctx context.Context, where model.CaseApprovalReviewCallBoolExp) (*model.CaseApprovalReviewCallMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.CaseApprovalReviewCall{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.CaseApprovalReviewCall
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
	return &model.CaseApprovalReviewCallMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteCaseApprovalReviewCallByPk(ctx context.Context, Id int64) (*model1.CaseApprovalReviewCall, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.CaseApprovalReviewCall
	tx := db.DB.Model(&model1.CaseApprovalReviewCall{})
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

func (r *mutationResolver) InsertCaseApprovalReviewCall(ctx context.Context, objects []*model.CaseApprovalReviewCallInsertInput) (*model.CaseApprovalReviewCallMutationResponse, error) {
	rs := make([]*model1.CaseApprovalReviewCall, 0)
	for _, object := range objects {
		v := &model1.CaseApprovalReviewCall{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.CaseApprovalReviewCall{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.CaseApprovalReviewCallMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertCaseApprovalReviewCallOne(ctx context.Context, object model.CaseApprovalReviewCallInsertInput) (*model1.CaseApprovalReviewCall, error) {
	rs := &model1.CaseApprovalReviewCall{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.CaseApprovalReviewCall{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateCaseApprovalReviewCall(ctx context.Context, inc *model.CaseApprovalReviewCallIncInput, set *model.CaseApprovalReviewCallSetInput, where model.CaseApprovalReviewCallBoolExp) (*model.CaseApprovalReviewCallMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.CaseApprovalReviewCall{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.CaseApprovalReviewCallMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.CaseApprovalReviewCallMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateCaseApprovalReviewCallByPk(ctx context.Context, inc *model.CaseApprovalReviewCallIncInput, set *model.CaseApprovalReviewCallSetInput, Id int64) (*model1.CaseApprovalReviewCall, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.CaseApprovalReviewCall{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.CaseApprovalReviewCall
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) CaseApprovalReviewCall(ctx context.Context, distinctOn []model.CaseApprovalReviewCallSelectColumn, limit *int, offset *int, orderBy []*model.CaseApprovalReviewCallOrderBy, where *model.CaseApprovalReviewCallBoolExp) ([]*model1.CaseApprovalReviewCall, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.CaseApprovalReviewCall{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.CaseApprovalReviewCall
	tx = tx.Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) CaseApprovalReviewCallAggregate(ctx context.Context, distinctOn []model.CaseApprovalReviewCallSelectColumn, limit *int, offset *int, orderBy []*model.CaseApprovalReviewCallOrderBy, where *model.CaseApprovalReviewCallBoolExp) (*model.CaseApprovalReviewCallAggregate, error) {
	var rs model.CaseApprovalReviewCallAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.CaseApprovalReviewCall{})
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

func (r *queryResolver) CaseApprovalReviewCallByPk(ctx context.Context, Id int64) (*model1.CaseApprovalReviewCall, error) {
	var rs model1.CaseApprovalReviewCall
	tx := db.DB.Model(&model1.CaseApprovalReviewCall{}).First(&rs, Id)
	err := tx.Error
	return &rs, err
}
