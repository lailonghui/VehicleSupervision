package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/vehicle_violation/graph/generated"
	"VehicleSupervision/internal/modules/vehicle_violation/graph/model"
	model1 "VehicleSupervision/internal/modules/vehicle_violation/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteVehicleViolationDetails(ctx context.Context, where model.VehicleViolationDetailsBoolExp) (*model.VehicleViolationDetailsMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleViolationDetails{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.VehicleViolationDetails
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
	return &model.VehicleViolationDetailsMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteVehicleViolationDetailsByPk(ctx context.Context, Id int64) (*model1.VehicleViolationDetails, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.VehicleViolationDetails
	tx := db.DB.Model(&model1.VehicleViolationDetails{})
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

func (r *mutationResolver) InsertVehicleViolationDetails(ctx context.Context, objects []*model.VehicleViolationDetailsInsertInput) (*model.VehicleViolationDetailsMutationResponse, error) {
	rs := []*model1.VehicleViolationDetails{}
	for _, object := range objects {
		v := &model1.VehicleViolationDetails{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.VehicleViolationDetails{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.VehicleViolationDetailsMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertVehicleViolationDetailsOne(ctx context.Context, object model.VehicleViolationDetailsInsertInput) (*model1.VehicleViolationDetails, error) {
	rs := &model1.VehicleViolationDetails{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.VehicleViolationDetails{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateVehicleViolationDetails(ctx context.Context, inc *model.VehicleViolationDetailsIncInput, set *model.VehicleViolationDetailsSetInput, where model.VehicleViolationDetailsBoolExp) (*model.VehicleViolationDetailsMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleViolationDetails{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.VehicleViolationDetailsMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.VehicleViolationDetailsMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateVehicleViolationDetailsByPk(ctx context.Context, inc *model.VehicleViolationDetailsIncInput, set *model.VehicleViolationDetailsSetInput, Id int64) (*model1.VehicleViolationDetails, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.VehicleViolationDetails{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.VehicleViolationDetails
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) VehicleViolationDetails(ctx context.Context, distinctOn []model.VehicleViolationDetailsSelectColumn, limit *int, offset *int, orderBy []*model.VehicleViolationDetailsOrderBy, where *model.VehicleViolationDetailsBoolExp) ([]*model1.VehicleViolationDetails, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleViolationDetails{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.VehicleViolationDetails
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) VehicleViolationDetailsAggregate(ctx context.Context, distinctOn []model.VehicleViolationDetailsSelectColumn, limit *int, offset *int, orderBy []*model.VehicleViolationDetailsOrderBy, where *model.VehicleViolationDetailsBoolExp) (*model.VehicleViolationDetailsAggregate, error) {
	var rs model.VehicleViolationDetailsAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.VehicleViolationDetails{})
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

func (r *queryResolver) VehicleViolationDetailsByPk(ctx context.Context, Id int64) (*model1.VehicleViolationDetails, error) {
	var rs model1.VehicleViolationDetails
	tx := db.DB.Model(&model1.VehicleViolationDetails{}).First(&rs, Id)
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
