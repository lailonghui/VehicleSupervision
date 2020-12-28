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

func (r *mutationResolver) DeleteDrivingLicenseRegistrationInspection(ctx context.Context, where model.DrivingLicenseRegistrationInspectionBoolExp) (*model.DrivingLicenseRegistrationInspectionMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DrivingLicenseRegistrationInspection{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.DrivingLicenseRegistrationInspection
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
	return &model.DrivingLicenseRegistrationInspectionMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteDrivingLicenseRegistrationInspectionByPk(ctx context.Context, Id int64) (*model1.DrivingLicenseRegistrationInspection, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.DrivingLicenseRegistrationInspection
	tx := db.DB.Model(&model1.DrivingLicenseRegistrationInspection{})
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

func (r *mutationResolver) InsertDrivingLicenseRegistrationInspection(ctx context.Context, objects []*model.DrivingLicenseRegistrationInspectionInsertInput) (*model.DrivingLicenseRegistrationInspectionMutationResponse, error) {
	rs := []*model1.DrivingLicenseRegistrationInspection{}
	for _, object := range objects {
		v := &model1.DrivingLicenseRegistrationInspection{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.DrivingLicenseRegistrationInspection{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.DrivingLicenseRegistrationInspectionMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertDrivingLicenseRegistrationInspectionOne(ctx context.Context, object model.DrivingLicenseRegistrationInspectionInsertInput) (*model1.DrivingLicenseRegistrationInspection, error) {
	rs := &model1.DrivingLicenseRegistrationInspection{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.DrivingLicenseRegistrationInspection{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateDrivingLicenseRegistrationInspection(ctx context.Context, inc *model.DrivingLicenseRegistrationInspectionIncInput, set *model.DrivingLicenseRegistrationInspectionSetInput, where model.DrivingLicenseRegistrationInspectionBoolExp) (*model.DrivingLicenseRegistrationInspectionMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DrivingLicenseRegistrationInspection{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.DrivingLicenseRegistrationInspectionMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.DrivingLicenseRegistrationInspectionMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateDrivingLicenseRegistrationInspectionByPk(ctx context.Context, inc *model.DrivingLicenseRegistrationInspectionIncInput, set *model.DrivingLicenseRegistrationInspectionSetInput, Id int64) (*model1.DrivingLicenseRegistrationInspection, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.DrivingLicenseRegistrationInspection{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.DrivingLicenseRegistrationInspection
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) DrivingLicenseRegistrationInspection(ctx context.Context, distinctOn []model.DrivingLicenseRegistrationInspectionSelectColumn, limit *int, offset *int, orderBy []*model.DrivingLicenseRegistrationInspectionOrderBy, where *model.DrivingLicenseRegistrationInspectionBoolExp) ([]*model1.DrivingLicenseRegistrationInspection, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DrivingLicenseRegistrationInspection{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.DrivingLicenseRegistrationInspection
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) DrivingLicenseRegistrationInspectionAggregate(ctx context.Context, distinctOn []model.DrivingLicenseRegistrationInspectionSelectColumn, limit *int, offset *int, orderBy []*model.DrivingLicenseRegistrationInspectionOrderBy, where *model.DrivingLicenseRegistrationInspectionBoolExp) (*model.DrivingLicenseRegistrationInspectionAggregate, error) {
	var rs model.DrivingLicenseRegistrationInspectionAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.DrivingLicenseRegistrationInspection{})
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

func (r *queryResolver) DrivingLicenseRegistrationInspectionByPk(ctx context.Context, Id int64) (*model1.DrivingLicenseRegistrationInspection, error) {
	var rs model1.DrivingLicenseRegistrationInspection
	tx := db.DB.Model(&model1.DrivingLicenseRegistrationInspection{}).First(&rs, Id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
