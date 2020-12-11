package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/driver/graph/generated"
	"VehicleSupervision/internal/modules/driver/graph/model"
	util2 "VehicleSupervision/pkg/graphql/util"
	"VehicleSupervision/pkg/util"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteVehicleDriverBinding(ctx context.Context, where model.VehicleDriverBindingBoolExp) (*model.VehicleDriverBindingMutationResponse, error) {
	var err error
	var rs []*model.VehicleDriverBinding
	qt := util2.NewQueryTranslator(db.DB, &model.VehicleDriverBinding{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util2.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	if len(preloads) > 0 {
		// 如果请求的字段不为空，则先查询一遍数据库
		tx := tx.Select(preloads)
		tx = tx.Find(&rs)
		// 如果查询结果含有错误，则返回错误
		if err := tx.Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, nil
			}
		}
	}
	// 删除
	tx = tx.Delete(nil)
	if err = tx.Error; err != nil {
		return nil, err
	}
	return &model.VehicleDriverBindingMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, err
}

func (r *mutationResolver) DeleteVehicleDriverBindingByPk(ctx context.Context, id int64) (*model.VehicleDriverBinding, error) {
	var rs model.VehicleDriverBinding
	var err error
	tx := db.DB.Model(&model.VehicleDriverBinding{})
	preloads := util2.GetPreloads(ctx)
	if len(preloads) > 0 {
		// 如果请求的字段不为空，则先查询一遍数据库
		tx = tx.Select(preloads).Where("id = ? ", id).First(&rs)
		// 如果查询结果含有错误，则返回错误
		if err := tx.Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				err = errors.New("0 rows affected or returned")
				return nil, err
			}
		}
	}
	// 删除
	tx = tx.Delete(nil)
	if err = tx.Error; err != nil {
		return nil, err
	}
	return &rs, nil
}

func (r *mutationResolver) InsertVehicleDriverBinding(ctx context.Context, objects []*model.VehicleDriverBindingInsertInput, onConflict *model.VehicleDriverBindingOnConflict) (*model.VehicleDriverBindingMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertVehicleDriverBindingOne(ctx context.Context, object model.VehicleDriverBindingInsertInput, onConflict *model.VehicleDriverBindingOnConflict) (*model.VehicleDriverBinding, error) {
	v := &model.VehicleDriverBinding{}
	util.StructAssign(v, &object)
	err := db.DB.Create(v).Error
	if err != nil {
		fmt.Println(err)
	}
	return v, err
}

func (r *mutationResolver) UpdateVehicleDriverBinding(ctx context.Context, inc *model.VehicleDriverBindingIncInput, set *model.VehicleDriverBindingSetInput, where model.VehicleDriverBindingBoolExp) (*model.VehicleDriverBindingMutationResponse, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.VehicleDriverBinding{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return &model.VehicleDriverBindingMutationResponse{
			AffectedRows: 0,
		}, err
	}
	var vehicleList []*model.VehicleDriverBinding
	tx.Scan(&vehicleList)
	return &model.VehicleDriverBindingMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    vehicleList,
	}, err
}

func (r *mutationResolver) UpdateVehicleDriverBindingByPk(ctx context.Context, inc *model.VehicleDriverBindingIncInput, set *model.VehicleDriverBindingSetInput, pkColumns model.VehicleDriverBindingPkColumnsInput) (*model.VehicleDriverBinding, error) {
	var err error
	tx := db.DB.Where("id = ? ", pkColumns.ID)
	qt := util2.NewQueryTranslator(tx, &model.VehicleDriverBinding{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return nil, err
	}
	var rs model.VehicleDriverBinding
	tx = tx.First(&rs)
	return &rs, err
}

func (r *queryResolver) VehicleDriverBinding(ctx context.Context, distinctOn []model.VehicleDriverBindingSelectColumn, limit *int, offset *int, orderBy []*model.VehicleDriverBindingOrderBy, where *model.VehicleDriverBindingBoolExp) ([]*model.VehicleDriverBinding, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.VehicleDriverBinding{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model.VehicleDriverBinding
	tx = tx.Find(&rs)
	if err = tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, err
}

func (r *queryResolver) VehicleDriverBindingAggregate(ctx context.Context, distinctOn []model.VehicleDriverBindingSelectColumn, limit *int, offset *int, orderBy []*model.VehicleDriverBindingOrderBy, where *model.VehicleDriverBindingBoolExp) (*model.VehicleDriverBindingAggregate, error) {
	var rs model.VehicleDriverBindingAggregate
	qt := util2.NewQueryTranslator(db.DB, &model.VehicleDriverBinding{})
	tx, err := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Aggregate(&rs, ctx)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if err = tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}
	if err != nil {
		fmt.Println(err)
	}
	return &rs, err
}

func (r *queryResolver) VehicleDriverBindingByPk(ctx context.Context, id int64) (*model.VehicleDriverBinding, error) {
	var rs model.VehicleDriverBinding
	err := db.DB.Where("id = ?", id).First(&rs).Error
	return &rs, err
}

func (r *subscriptionResolver) VehicleDriverBinding(ctx context.Context, distinctOn []model.VehicleDriverBindingSelectColumn, limit *int, offset *int, orderBy []*model.VehicleDriverBindingOrderBy, where *model.VehicleDriverBindingBoolExp) (<-chan []*model.VehicleDriverBinding, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VehicleDriverBindingAggregate(ctx context.Context, distinctOn []model.VehicleDriverBindingSelectColumn, limit *int, offset *int, orderBy []*model.VehicleDriverBindingOrderBy, where *model.VehicleDriverBindingBoolExp) (<-chan *model.VehicleDriverBindingAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VehicleDriverBindingByPk(ctx context.Context, id int64) (<-chan *model.VehicleDriverBinding, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
