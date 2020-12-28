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

func (r *mutationResolver) DeleteDriverIdentityInfoReport(ctx context.Context, where model.DriverIdentityInfoReportBoolExp) (*model.DriverIdentityInfoReportMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DriverIdentityInfoReport{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.DriverIdentityInfoReport
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
	return &model.DriverIdentityInfoReportMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteDriverIdentityInfoReportByPk(ctx context.Context, Id int64) (*model1.DriverIdentityInfoReport, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.DriverIdentityInfoReport
	tx := db.DB.Model(&model1.DriverIdentityInfoReport{})
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

func (r *mutationResolver) InsertDriverIdentityInfoReport(ctx context.Context, objects []*model.DriverIdentityInfoReportInsertInput) (*model.DriverIdentityInfoReportMutationResponse, error) {
	rs := make([]*model1.DriverIdentityInfoReport, 0)
	for _, object := range objects {
		v := &model1.DriverIdentityInfoReport{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.DriverIdentityInfoReport{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.DriverIdentityInfoReportMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertDriverIdentityInfoReportOne(ctx context.Context, object model.DriverIdentityInfoReportInsertInput) (*model1.DriverIdentityInfoReport, error) {
	rs := &model1.DriverIdentityInfoReport{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.DriverIdentityInfoReport{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateDriverIdentityInfoReport(ctx context.Context, inc *model.DriverIdentityInfoReportIncInput, set *model.DriverIdentityInfoReportSetInput, where model.DriverIdentityInfoReportBoolExp) (*model.DriverIdentityInfoReportMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DriverIdentityInfoReport{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.DriverIdentityInfoReportMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.DriverIdentityInfoReportMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateDriverIdentityInfoReportByPk(ctx context.Context, inc *model.DriverIdentityInfoReportIncInput, set *model.DriverIdentityInfoReportSetInput, Id int64) (*model1.DriverIdentityInfoReport, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.DriverIdentityInfoReport{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.DriverIdentityInfoReport
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) DriverIdentityInfoReport(ctx context.Context, distinctOn []model.DriverIdentityInfoReportSelectColumn, limit *int, offset *int, orderBy []*model.DriverIdentityInfoReportOrderBy, where *model.DriverIdentityInfoReportBoolExp) ([]*model1.DriverIdentityInfoReport, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DriverIdentityInfoReport{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.DriverIdentityInfoReport
	tx = tx.Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) DriverIdentityInfoReportAggregate(ctx context.Context, distinctOn []model.DriverIdentityInfoReportSelectColumn, limit *int, offset *int, orderBy []*model.DriverIdentityInfoReportOrderBy, where *model.DriverIdentityInfoReportBoolExp) (*model.DriverIdentityInfoReportAggregate, error) {
	var rs model.DriverIdentityInfoReportAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.DriverIdentityInfoReport{})
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

func (r *queryResolver) DriverIdentityInfoReportByPk(ctx context.Context, Id int64) (*model1.DriverIdentityInfoReport, error) {
	var rs model1.DriverIdentityInfoReport
	tx := db.DB.Model(&model1.DriverIdentityInfoReport{}).First(&rs, Id)
	err := tx.Error
	return &rs, err
}
