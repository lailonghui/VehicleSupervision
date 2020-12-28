package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/vehicle_violation/graph/model"
	model1 "VehicleSupervision/internal/modules/vehicle_violation/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteDeductionReport(ctx context.Context, where model.DeductionReportBoolExp) (*model.DeductionReportMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DeductionReport{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.DeductionReport
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
	return &model.DeductionReportMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteDeductionReportByPk(ctx context.Context, Id int64) (*model1.DeductionReport, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.DeductionReport
	tx := db.DB.Model(&model1.DeductionReport{})
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

func (r *mutationResolver) InsertDeductionReport(ctx context.Context, objects []*model.DeductionReportInsertInput) (*model.DeductionReportMutationResponse, error) {
	rs := []*model1.DeductionReport{}
	for _, object := range objects {
		v := &model1.DeductionReport{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.DeductionReport{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.DeductionReportMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertDeductionReportOne(ctx context.Context, object model.DeductionReportInsertInput) (*model1.DeductionReport, error) {
	rs := &model1.DeductionReport{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.DeductionReport{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateDeductionReport(ctx context.Context, inc *model.DeductionReportIncInput, set *model.DeductionReportSetInput, where model.DeductionReportBoolExp) (*model.DeductionReportMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DeductionReport{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.DeductionReportMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.DeductionReportMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateDeductionReportByPk(ctx context.Context, inc *model.DeductionReportIncInput, set *model.DeductionReportSetInput, Id int64) (*model1.DeductionReport, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.DeductionReport{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.DeductionReport
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) DeductionReport(ctx context.Context, distinctOn []model.DeductionReportSelectColumn, limit *int, offset *int, orderBy []*model.DeductionReportOrderBy, where *model.DeductionReportBoolExp) ([]*model1.DeductionReport, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DeductionReport{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.DeductionReport
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) DeductionReportAggregate(ctx context.Context, distinctOn []model.DeductionReportSelectColumn, limit *int, offset *int, orderBy []*model.DeductionReportOrderBy, where *model.DeductionReportBoolExp) (*model.DeductionReportAggregate, error) {
	var rs model.DeductionReportAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.DeductionReport{})
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

func (r *queryResolver) DeductionReportByPk(ctx context.Context, Id int64) (*model1.DeductionReport, error) {
	var rs model1.DeductionReport
	tx := db.DB.Model(&model1.DeductionReport{}).First(&rs, Id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
