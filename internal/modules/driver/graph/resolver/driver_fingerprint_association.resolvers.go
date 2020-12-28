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

func (r *mutationResolver) DeleteDriverFingerprintAssociation(ctx context.Context, where model.DriverFingerprintAssociationBoolExp) (*model.DriverFingerprintAssociationMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DriverFingerprintAssociation{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.DriverFingerprintAssociation
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
	return &model.DriverFingerprintAssociationMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteDriverFingerprintAssociationByPk(ctx context.Context, Id int64) (*model1.DriverFingerprintAssociation, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.DriverFingerprintAssociation
	tx := db.DB.Model(&model1.DriverFingerprintAssociation{})
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

func (r *mutationResolver) InsertDriverFingerprintAssociation(ctx context.Context, objects []*model.DriverFingerprintAssociationInsertInput) (*model.DriverFingerprintAssociationMutationResponse, error) {
	rs := []*model1.DriverFingerprintAssociation{}
	for _, object := range objects {
		v := &model1.DriverFingerprintAssociation{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.DriverFingerprintAssociation{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.DriverFingerprintAssociationMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertDriverFingerprintAssociationOne(ctx context.Context, object model.DriverFingerprintAssociationInsertInput) (*model1.DriverFingerprintAssociation, error) {
	rs := &model1.DriverFingerprintAssociation{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.DriverFingerprintAssociation{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateDriverFingerprintAssociation(ctx context.Context, inc *model.DriverFingerprintAssociationIncInput, set *model.DriverFingerprintAssociationSetInput, where model.DriverFingerprintAssociationBoolExp) (*model.DriverFingerprintAssociationMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DriverFingerprintAssociation{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.DriverFingerprintAssociationMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.DriverFingerprintAssociationMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateDriverFingerprintAssociationByPk(ctx context.Context, inc *model.DriverFingerprintAssociationIncInput, set *model.DriverFingerprintAssociationSetInput, Id int64) (*model1.DriverFingerprintAssociation, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.DriverFingerprintAssociation{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.DriverFingerprintAssociation
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) DriverFingerprintAssociation(ctx context.Context, distinctOn []model.DriverFingerprintAssociationSelectColumn, limit *int, offset *int, orderBy []*model.DriverFingerprintAssociationOrderBy, where *model.DriverFingerprintAssociationBoolExp) ([]*model1.DriverFingerprintAssociation, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DriverFingerprintAssociation{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.DriverFingerprintAssociation
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) DriverFingerprintAssociationAggregate(ctx context.Context, distinctOn []model.DriverFingerprintAssociationSelectColumn, limit *int, offset *int, orderBy []*model.DriverFingerprintAssociationOrderBy, where *model.DriverFingerprintAssociationBoolExp) (*model.DriverFingerprintAssociationAggregate, error) {
	var rs model.DriverFingerprintAssociationAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.DriverFingerprintAssociation{})
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

func (r *queryResolver) DriverFingerprintAssociationByPk(ctx context.Context, Id int64) (*model1.DriverFingerprintAssociation, error) {
	var rs model1.DriverFingerprintAssociation
	tx := db.DB.Model(&model1.DriverFingerprintAssociation{}).First(&rs, Id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
