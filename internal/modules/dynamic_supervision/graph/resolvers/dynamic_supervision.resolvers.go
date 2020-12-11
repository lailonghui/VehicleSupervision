package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/dynamic_supervision/graph/generated"
	"VehicleSupervision/internal/modules/dynamic_supervision/graph/model"
	util2 "VehicleSupervision/pkg/graphql/util"
	"VehicleSupervision/pkg/util"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteDynamicSupervision(ctx context.Context, where model.DynamicSupervisionBoolExp) (*model.DynamicSupervisionMutationResponse, error) {
	var err error
	var rs []*model.DynamicSupervision
	qt := util2.NewQueryTranslator(db.DB, &model.DynamicSupervision{})
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
	return &model.DynamicSupervisionMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, err
}

func (r *mutationResolver) DeleteDynamicSupervisionByPk(ctx context.Context, id int64, supervisionID string) (*model.DynamicSupervision, error) {
	var rs model.DynamicSupervision
	var err error
	tx := db.DB.Model(&model.DynamicSupervision{})
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

func (r *mutationResolver) InsertDynamicSupervision(ctx context.Context, objects []*model.DynamicSupervisionInsertInput, onConflict *model.DynamicSupervisionOnConflict) (*model.DynamicSupervisionMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertDynamicSupervisionOne(ctx context.Context, object model.DynamicSupervisionInsertInput, onConflict *model.DynamicSupervisionOnConflict) (*model.DynamicSupervision, error) {
	v := &model.DynamicSupervision{}
	util.StructAssign(v, &object)
	err := db.DB.Create(v).Error
	if err != nil {
		fmt.Println(err)
	}
	return v, err
}

func (r *mutationResolver) UpdateDynamicSupervision(ctx context.Context, inc *model.DynamicSupervisionIncInput, set *model.DynamicSupervisionSetInput, where model.DynamicSupervisionBoolExp) (*model.DynamicSupervisionMutationResponse, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.DynamicSupervision{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return &model.DynamicSupervisionMutationResponse{
			AffectedRows: 0,
		}, err
	}
	var vehicleList []*model.DynamicSupervision
	tx.Scan(&vehicleList)
	return &model.DynamicSupervisionMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    vehicleList,
	}, err
}

func (r *mutationResolver) UpdateDynamicSupervisionByPk(ctx context.Context, inc *model.DynamicSupervisionIncInput, set *model.DynamicSupervisionSetInput, pkColumns model.DynamicSupervisionPkColumnsInput) (*model.DynamicSupervision, error) {
	var err error
	tx := db.DB.Where("id = ? ", pkColumns.ID)
	qt := util2.NewQueryTranslator(tx, &model.DynamicSupervision{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return nil, err
	}
	var rs model.DynamicSupervision
	tx = tx.First(&rs)
	return &rs, err
}

func (r *queryResolver) DynamicSupervision(ctx context.Context, distinctOn []model.DynamicSupervisionSelectColumn, limit *int, offset *int, orderBy []*model.DynamicSupervisionOrderBy, where *model.DynamicSupervisionBoolExp) ([]*model.DynamicSupervision, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.DynamicSupervision{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model.DynamicSupervision
	tx = tx.Find(&rs)
	if err = tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, err
}

func (r *queryResolver) DynamicSupervisionAggregate(ctx context.Context, distinctOn []model.DynamicSupervisionSelectColumn, limit *int, offset *int, orderBy []*model.DynamicSupervisionOrderBy, where *model.DynamicSupervisionBoolExp) (*model.DynamicSupervisionAggregate, error) {
	var rs model.DynamicSupervisionAggregate
	qt := util2.NewQueryTranslator(db.DB, &model.DynamicSupervision{})
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

func (r *queryResolver) DynamicSupervisionByPk(ctx context.Context, id int64, supervisionID string) (*model.DynamicSupervision, error) {
	var rs model.DynamicSupervision
	err := db.DB.Where("id = ?", id).First(&rs).Error
	return &rs, err
}

func (r *subscriptionResolver) DynamicSupervision(ctx context.Context, distinctOn []model.DynamicSupervisionSelectColumn, limit *int, offset *int, orderBy []*model.DynamicSupervisionOrderBy, where *model.DynamicSupervisionBoolExp) (<-chan []*model.DynamicSupervision, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DynamicSupervisionAggregate(ctx context.Context, distinctOn []model.DynamicSupervisionSelectColumn, limit *int, offset *int, orderBy []*model.DynamicSupervisionOrderBy, where *model.DynamicSupervisionBoolExp) (<-chan *model.DynamicSupervisionAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DynamicSupervisionByPk(ctx context.Context, id int64, supervisionID string) (<-chan *model.DynamicSupervision, error) {
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
