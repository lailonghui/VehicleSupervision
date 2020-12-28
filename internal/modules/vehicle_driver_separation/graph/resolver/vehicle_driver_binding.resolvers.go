package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/vehicle_driver_separation/graph/generated"
	"VehicleSupervision/internal/modules/vehicle_driver_separation/graph/model"
	model1 "VehicleSupervision/internal/modules/vehicle_driver_separation/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteVehicleDriverBinding(ctx context.Context, where model.VehicleDriverBindingBoolExp) (*model.VehicleDriverBindingMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleDriverBinding{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.VehicleDriverBinding
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
	return &model.VehicleDriverBindingMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteVehicleDriverBindingByPk(ctx context.Context, Id int64) (*model1.VehicleDriverBinding, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.VehicleDriverBinding
	tx := db.DB.Model(&model1.VehicleDriverBinding{})
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

func (r *mutationResolver) InsertVehicleDriverBinding(ctx context.Context, objects []*model.VehicleDriverBindingInsertInput) (*model.VehicleDriverBindingMutationResponse, error) {
	rs := []*model1.VehicleDriverBinding{}
	for _, object := range objects {
		v := &model1.VehicleDriverBinding{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.VehicleDriverBinding{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.VehicleDriverBindingMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertVehicleDriverBindingOne(ctx context.Context, object model.VehicleDriverBindingInsertInput) (*model1.VehicleDriverBinding, error) {
	rs := &model1.VehicleDriverBinding{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.VehicleDriverBinding{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateVehicleDriverBinding(ctx context.Context, inc *model.VehicleDriverBindingIncInput, set *model.VehicleDriverBindingSetInput, where model.VehicleDriverBindingBoolExp) (*model.VehicleDriverBindingMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleDriverBinding{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.VehicleDriverBindingMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.VehicleDriverBindingMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateVehicleDriverBindingByPk(ctx context.Context, inc *model.VehicleDriverBindingIncInput, set *model.VehicleDriverBindingSetInput, Id int64) (*model1.VehicleDriverBinding, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.VehicleDriverBinding{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.VehicleDriverBinding
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) VehicleDriverBinding(ctx context.Context, distinctOn []model.VehicleDriverBindingSelectColumn, limit *int, offset *int, orderBy []*model.VehicleDriverBindingOrderBy, where *model.VehicleDriverBindingBoolExp) ([]*model1.VehicleDriverBinding, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleDriverBinding{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.VehicleDriverBinding
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) VehicleDriverBindingAggregate(ctx context.Context, distinctOn []model.VehicleDriverBindingSelectColumn, limit *int, offset *int, orderBy []*model.VehicleDriverBindingOrderBy, where *model.VehicleDriverBindingBoolExp) (*model.VehicleDriverBindingAggregate, error) {
	var rs model.VehicleDriverBindingAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.VehicleDriverBinding{})
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

func (r *queryResolver) VehicleDriverBindingByPk(ctx context.Context, Id int64) (*model1.VehicleDriverBinding, error) {
	var rs model1.VehicleDriverBinding
	tx := db.DB.Model(&model1.VehicleDriverBinding{}).First(&rs, Id)
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
