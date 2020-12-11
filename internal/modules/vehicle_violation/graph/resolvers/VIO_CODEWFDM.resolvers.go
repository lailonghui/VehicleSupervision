package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/vehicle_violation/graph/model"
	util2 "VehicleSupervision/pkg/graphql/util"
	"VehicleSupervision/pkg/util"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteVioCodewfdm(ctx context.Context, where model.VioCodewfdmBoolExp) (*model.VioCodewfdmMutationResponse, error) {
	var err error
	var rs []*model.VioCodewfdm
	qt := util2.NewQueryTranslator(db.DB, &model.VioCodewfdm{})
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
	return &model.VioCodewfdmMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, err
}

func (r *mutationResolver) InsertVioCodewfdm(ctx context.Context, objects []*model.VioCodewfdmInsertInput) (*model.VioCodewfdmMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertVioCodewfdmOne(ctx context.Context, object model.VioCodewfdmInsertInput) (*model.VioCodewfdm, error) {
	v := &model.VioCodewfdm{}
	util.StructAssign(v, &object)
	err := db.DB.Create(v).Error
	if err != nil {
		fmt.Println(err)
	}
	return v, err
}

func (r *mutationResolver) UpdateVioCodewfdm(ctx context.Context, inc *model.VioCodewfdmIncInput, set *model.VioCodewfdmSetInput, where model.VioCodewfdmBoolExp) (*model.VioCodewfdmMutationResponse, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.VioCodewfdm{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return &model.VioCodewfdmMutationResponse{
			AffectedRows: 0,
		}, err
	}
	var vehicleList []*model.VioCodewfdm
	tx.Scan(&vehicleList)
	return &model.VioCodewfdmMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    vehicleList,
	}, err
}

func (r *queryResolver) VioCodewfdm(ctx context.Context, distinctOn []model.VioCodewfdmSelectColumn, limit *int, offset *int, orderBy []*model.VioCodewfdmOrderBy, where *model.VioCodewfdmBoolExp) ([]*model.VioCodewfdm, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.VioCodewfdm{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model.VioCodewfdm
	tx = tx.Find(&rs)
	if err = tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, err
}

func (r *queryResolver) VioCodewfdmAggregate(ctx context.Context, distinctOn []model.VioCodewfdmSelectColumn, limit *int, offset *int, orderBy []*model.VioCodewfdmOrderBy, where *model.VioCodewfdmBoolExp) (*model.VioCodewfdmAggregate, error) {
	var rs model.VioCodewfdmAggregate
	qt := util2.NewQueryTranslator(db.DB, &model.VioCodewfdm{})
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

func (r *subscriptionResolver) VioCodewfdm(ctx context.Context, distinctOn []model.VioCodewfdmSelectColumn, limit *int, offset *int, orderBy []*model.VioCodewfdmOrderBy, where *model.VioCodewfdmBoolExp) (<-chan []*model.VioCodewfdm, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VioCodewfdmAggregate(ctx context.Context, distinctOn []model.VioCodewfdmSelectColumn, limit *int, offset *int, orderBy []*model.VioCodewfdmOrderBy, where *model.VioCodewfdmBoolExp) (<-chan *model.VioCodewfdmAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}
