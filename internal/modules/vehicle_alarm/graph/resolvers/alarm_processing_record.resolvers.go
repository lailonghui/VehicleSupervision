package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/vehicle_alarm/graph/generated"
	"VehicleSupervision/internal/modules/vehicle_alarm/graph/model"
	util2 "VehicleSupervision/pkg/graphql/util"
	"VehicleSupervision/pkg/util"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteAlarmProcessingRecord(ctx context.Context, where model.AlarmProcessingRecordBoolExp) (*model.AlarmProcessingRecordMutationResponse, error) {
	var err error
	var rs []*model.AlarmProcessingRecord
	qt := util2.NewQueryTranslator(db.DB, &model.AlarmProcessingRecord{})
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
	return &model.AlarmProcessingRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, err
}

func (r *mutationResolver) InsertAlarmProcessingRecord(ctx context.Context, objects []*model.AlarmProcessingRecordInsertInput) (*model.AlarmProcessingRecordMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertAlarmProcessingRecordOne(ctx context.Context, object model.AlarmProcessingRecordInsertInput) (*model.AlarmProcessingRecord, error) {
	v := &model.AlarmProcessingRecord{}
	util.StructAssign(v, &object)
	err := db.DB.Create(v).Error
	if err != nil {
		fmt.Println(err)
	}
	return v, err
}

func (r *mutationResolver) UpdateAlarmProcessingRecord(ctx context.Context, inc *model.AlarmProcessingRecordIncInput, set *model.AlarmProcessingRecordSetInput, where model.AlarmProcessingRecordBoolExp) (*model.AlarmProcessingRecordMutationResponse, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.AlarmProcessingRecord{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return &model.AlarmProcessingRecordMutationResponse{
			AffectedRows: 0,
		}, err
	}
	var vehicleList []*model.AlarmProcessingRecord
	tx.Scan(&vehicleList)
	return &model.AlarmProcessingRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    vehicleList,
	}, err
}

func (r *queryResolver) AlarmProcessingRecord(ctx context.Context, distinctOn []model.AlarmProcessingRecordSelectColumn, limit *int, offset *int, orderBy []*model.AlarmProcessingRecordOrderBy, where *model.AlarmProcessingRecordBoolExp) ([]*model.AlarmProcessingRecord, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.AlarmProcessingRecord{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model.AlarmProcessingRecord
	tx = tx.Find(&rs)
	if err = tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, err
}

func (r *queryResolver) AlarmProcessingRecordAggregate(ctx context.Context, distinctOn []model.AlarmProcessingRecordSelectColumn, limit *int, offset *int, orderBy []*model.AlarmProcessingRecordOrderBy, where *model.AlarmProcessingRecordBoolExp) (*model.AlarmProcessingRecordAggregate, error) {
	var rs model.AlarmProcessingRecordAggregate
	qt := util2.NewQueryTranslator(db.DB, &model.AlarmProcessingRecord{})
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

func (r *subscriptionResolver) AlarmProcessingRecord(ctx context.Context, distinctOn []model.AlarmProcessingRecordSelectColumn, limit *int, offset *int, orderBy []*model.AlarmProcessingRecordOrderBy, where *model.AlarmProcessingRecordBoolExp) (<-chan []*model.AlarmProcessingRecord, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) AlarmProcessingRecordAggregate(ctx context.Context, distinctOn []model.AlarmProcessingRecordSelectColumn, limit *int, offset *int, orderBy []*model.AlarmProcessingRecordOrderBy, where *model.AlarmProcessingRecordBoolExp) (<-chan *model.AlarmProcessingRecordAggregate, error) {
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
