package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/vehicle_violation/graph/generated"
	"VehicleSupervision/internal/modules/vehicle_violation/graph/model"
	util2 "VehicleSupervision/pkg/graphql/util"
	"VehicleSupervision/pkg/util"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteAppEnforcement(ctx context.Context, where model.AppEnforcementBoolExp) (*model.AppEnforcementMutationResponse, error) {
	var err error
	var rs []*model.AppEnforcement
	qt := util2.NewQueryTranslator(db.DB, &model.AppEnforcement{})
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
	return &model.AppEnforcementMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, err
}

func (r *mutationResolver) DeleteAppEnforcementByPk(ctx context.Context, id int64, illegalPhotoID string) (*model.AppEnforcement, error) {
	var rs model.AppEnforcement
	var err error
	tx := db.DB.Model(&model.AppEnforcement{})
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

func (r *mutationResolver) InsertAppEnforcement(ctx context.Context, objects []*model.AppEnforcementInsertInput, onConflict *model.AppEnforcementOnConflict) (*model.AppEnforcementMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertAppEnforcementOne(ctx context.Context, object model.AppEnforcementInsertInput, onConflict *model.AppEnforcementOnConflict) (*model.AppEnforcement, error) {
	v := &model.AppEnforcement{}
	util.StructAssign(v, &object)
	err := db.DB.Create(v).Error
	if err != nil {
		fmt.Println(err)
	}
	return v, err
}

func (r *mutationResolver) UpdateAppEnforcement(ctx context.Context, inc *model.AppEnforcementIncInput, set *model.AppEnforcementSetInput, where model.AppEnforcementBoolExp) (*model.AppEnforcementMutationResponse, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.AppEnforcement{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return &model.AppEnforcementMutationResponse{
			AffectedRows: 0,
		}, err
	}
	var vehicleList []*model.AppEnforcement
	tx.Scan(&vehicleList)
	return &model.AppEnforcementMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    vehicleList,
	}, err
}

func (r *mutationResolver) UpdateAppEnforcementByPk(ctx context.Context, inc *model.AppEnforcementIncInput, set *model.AppEnforcementSetInput, pkColumns model.AppEnforcementPkColumnsInput) (*model.AppEnforcement, error) {
	var err error
	tx := db.DB.Where("id = ? ", pkColumns.ID)
	qt := util2.NewQueryTranslator(tx, &model.AppEnforcement{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return nil, err
	}
	var rs model.AppEnforcement
	tx = tx.First(&rs)
	return &rs, err
}

func (r *queryResolver) AppEnforcement(ctx context.Context, distinctOn []model.AppEnforcementSelectColumn, limit *int, offset *int, orderBy []*model.AppEnforcementOrderBy, where *model.AppEnforcementBoolExp) ([]*model.AppEnforcement, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.AppEnforcement{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model.AppEnforcement
	tx = tx.Find(&rs)
	if err = tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, err
}

func (r *queryResolver) AppEnforcementAggregate(ctx context.Context, distinctOn []model.AppEnforcementSelectColumn, limit *int, offset *int, orderBy []*model.AppEnforcementOrderBy, where *model.AppEnforcementBoolExp) (*model.AppEnforcementAggregate, error) {
	var rs model.AppEnforcementAggregate
	qt := util2.NewQueryTranslator(db.DB, &model.AppEnforcement{})
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

func (r *queryResolver) AppEnforcementByPk(ctx context.Context, id int64, illegalPhotoID string) (*model.AppEnforcement, error) {
	var rs model.AppEnforcement
	err := db.DB.Where("id = ?", id).First(&rs).Error
	return &rs, err
}

func (r *subscriptionResolver) AppEnforcement(ctx context.Context, distinctOn []model.AppEnforcementSelectColumn, limit *int, offset *int, orderBy []*model.AppEnforcementOrderBy, where *model.AppEnforcementBoolExp) (<-chan []*model.AppEnforcement, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) AppEnforcementAggregate(ctx context.Context, distinctOn []model.AppEnforcementSelectColumn, limit *int, offset *int, orderBy []*model.AppEnforcementOrderBy, where *model.AppEnforcementBoolExp) (<-chan *model.AppEnforcementAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) AppEnforcementByPk(ctx context.Context, id int64, illegalPhotoID string) (<-chan *model.AppEnforcement, error) {
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
