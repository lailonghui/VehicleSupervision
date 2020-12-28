package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/driver/graph/model"
	model1 "VehicleSupervision/internal/modules/driver/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteDriverIdInfoReport(ctx context.Context, where model.DriverIdInfoReportBoolExp) (*model.DriverIdInfoReportMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DriverIdInfoReport{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.DriverIdInfoReport
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
	return &model.DriverIdInfoReportMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteDriverIdInfoReportByPk(ctx context.Context, Id int64) (*model1.DriverIdInfoReport, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.DriverIdInfoReport
	tx := db.DB.Model(&model1.DriverIdInfoReport{})
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

func (r *mutationResolver) InsertDriverIdInfoReport(ctx context.Context, objects []*model.DriverIdInfoReportInsertInput) (*model.DriverIdInfoReportMutationResponse, error) {
	rs := []*model1.DriverIdInfoReport{}
	for _, object := range objects {
		v := &model1.DriverIdInfoReport{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.DriverIdInfoReport{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.DriverIdInfoReportMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertDriverIdInfoReportOne(ctx context.Context, object model.DriverIdInfoReportInsertInput) (*model1.DriverIdInfoReport, error) {
	rs := &model1.DriverIdInfoReport{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.DriverIdInfoReport{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateDriverIdInfoReport(ctx context.Context, inc *model.DriverIdInfoReportIncInput, set *model.DriverIdInfoReportSetInput, where model.DriverIdInfoReportBoolExp) (*model.DriverIdInfoReportMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DriverIdInfoReport{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.DriverIdInfoReportMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.DriverIdInfoReportMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateDriverIdInfoReportByPk(ctx context.Context, inc *model.DriverIdInfoReportIncInput, set *model.DriverIdInfoReportSetInput, Id int64) (*model1.DriverIdInfoReport, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.DriverIdInfoReport{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.DriverIdInfoReport
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) DriverIdInfoReport(ctx context.Context, distinctOn []model.DriverIdInfoReportSelectColumn, limit *int, offset *int, orderBy []*model.DriverIdInfoReportOrderBy, where *model.DriverIdInfoReportBoolExp) ([]*model1.DriverIdInfoReport, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DriverIdInfoReport{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.DriverIdInfoReport
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) DriverIdInfoReportAggregate(ctx context.Context, distinctOn []model.DriverIdInfoReportSelectColumn, limit *int, offset *int, orderBy []*model.DriverIdInfoReportOrderBy, where *model.DriverIdInfoReportBoolExp) (*model.DriverIdInfoReportAggregate, error) {
	var rs model.DriverIdInfoReportAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.DriverIdInfoReport{})
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

func (r *queryResolver) DriverIdInfoReportByPk(ctx context.Context, Id int64) (*model1.DriverIdInfoReport, error) {
	var rs model1.DriverIdInfoReport
	tx := db.DB.Model(&model1.DriverIdInfoReport{}).First(&rs, Id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
