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

func (r *mutationResolver) DeleteCaseApprovalReviewOperation(ctx context.Context, where model.CaseApprovalReviewOperationBoolExp) (*model.CaseApprovalReviewOperationMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.CaseApprovalReviewOperation{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.CaseApprovalReviewOperation
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
	return &model.CaseApprovalReviewOperationMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteCaseApprovalReviewOperationByPk(ctx context.Context, Id int64) (*model1.CaseApprovalReviewOperation, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.CaseApprovalReviewOperation
	tx := db.DB.Model(&model1.CaseApprovalReviewOperation{})
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

func (r *mutationResolver) InsertCaseApprovalReviewOperation(ctx context.Context, objects []*model.CaseApprovalReviewOperationInsertInput) (*model.CaseApprovalReviewOperationMutationResponse, error) {
	rs := make([]*model1.CaseApprovalReviewOperation, 0)
	for _, object := range objects {
		v := &model1.CaseApprovalReviewOperation{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.CaseApprovalReviewOperation{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.CaseApprovalReviewOperationMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertCaseApprovalReviewOperationOne(ctx context.Context, object model.CaseApprovalReviewOperationInsertInput) (*model1.CaseApprovalReviewOperation, error) {
	rs := &model1.CaseApprovalReviewOperation{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.CaseApprovalReviewOperation{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateCaseApprovalReviewOperation(ctx context.Context, inc *model.CaseApprovalReviewOperationIncInput, set *model.CaseApprovalReviewOperationSetInput, where model.CaseApprovalReviewOperationBoolExp) (*model.CaseApprovalReviewOperationMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.CaseApprovalReviewOperation{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.CaseApprovalReviewOperationMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.CaseApprovalReviewOperationMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateCaseApprovalReviewOperationByPk(ctx context.Context, inc *model.CaseApprovalReviewOperationIncInput, set *model.CaseApprovalReviewOperationSetInput, Id int64) (*model1.CaseApprovalReviewOperation, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.CaseApprovalReviewOperation{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.CaseApprovalReviewOperation
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) CaseApprovalReviewOperation(ctx context.Context, distinctOn []model.CaseApprovalReviewOperationSelectColumn, limit *int, offset *int, orderBy []*model.CaseApprovalReviewOperationOrderBy, where *model.CaseApprovalReviewOperationBoolExp) ([]*model1.CaseApprovalReviewOperation, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.CaseApprovalReviewOperation{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.CaseApprovalReviewOperation
	tx = tx.Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) CaseApprovalReviewOperationAggregate(ctx context.Context, distinctOn []model.CaseApprovalReviewOperationSelectColumn, limit *int, offset *int, orderBy []*model.CaseApprovalReviewOperationOrderBy, where *model.CaseApprovalReviewOperationBoolExp) (*model.CaseApprovalReviewOperationAggregate, error) {
	var rs model.CaseApprovalReviewOperationAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.CaseApprovalReviewOperation{})
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

func (r *queryResolver) CaseApprovalReviewOperationByPk(ctx context.Context, Id int64) (*model1.CaseApprovalReviewOperation, error) {
	var rs model1.CaseApprovalReviewOperation
	tx := db.DB.Model(&model1.CaseApprovalReviewOperation{}).First(&rs, Id)
	err := tx.Error
	return &rs, err
}
