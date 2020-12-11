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

func (r *mutationResolver) DeleteVehicleViolationScoringItems(ctx context.Context, where model.VehicleViolationScoringItemsBoolExp) (*model.VehicleViolationScoringItemsMutationResponse, error) {
	var err error
	var rs []*model.VehicleViolationScoringItems
	qt := util2.NewQueryTranslator(db.DB, &model.VehicleViolationScoringItems{})
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
	return &model.VehicleViolationScoringItemsMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, err
}

func (r *mutationResolver) DeleteVehicleViolationScoringItemsByPk(ctx context.Context, id int64, violationScoringItemID string) (*model.VehicleViolationScoringItems, error) {
	var rs model.VehicleViolationScoringItems
	var err error
	tx := db.DB.Model(&model.VehicleViolationScoringItems{})
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

func (r *mutationResolver) InsertVehicleViolationScoringItems(ctx context.Context, objects []*model.VehicleViolationScoringItemsInsertInput, onConflict *model.VehicleViolationScoringItemsOnConflict) (*model.VehicleViolationScoringItemsMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertVehicleViolationScoringItemsOne(ctx context.Context, object model.VehicleViolationScoringItemsInsertInput, onConflict *model.VehicleViolationScoringItemsOnConflict) (*model.VehicleViolationScoringItems, error) {
	v := &model.VehicleViolationScoringItems{}
	util.StructAssign(v, &object)
	err := db.DB.Create(v).Error
	if err != nil {
		fmt.Println(err)
	}
	return v, err
}

func (r *mutationResolver) UpdateVehicleViolationScoringItems(ctx context.Context, inc *model.VehicleViolationScoringItemsIncInput, set *model.VehicleViolationScoringItemsSetInput, where model.VehicleViolationScoringItemsBoolExp) (*model.VehicleViolationScoringItemsMutationResponse, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.VehicleViolationScoringItems{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return &model.VehicleViolationScoringItemsMutationResponse{
			AffectedRows: 0,
		}, err
	}
	var vehicleList []*model.VehicleViolationScoringItems
	tx.Scan(&vehicleList)
	return &model.VehicleViolationScoringItemsMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    vehicleList,
	}, err
}

func (r *mutationResolver) UpdateVehicleViolationScoringItemsByPk(ctx context.Context, inc *model.VehicleViolationScoringItemsIncInput, set *model.VehicleViolationScoringItemsSetInput, pkColumns model.VehicleViolationScoringItemsPkColumnsInput) (*model.VehicleViolationScoringItems, error) {
	var err error
	tx := db.DB.Where("id = ? ", pkColumns.ID)
	qt := util2.NewQueryTranslator(tx, &model.VehicleViolationScoringItems{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return nil, err
	}
	var rs model.VehicleViolationScoringItems
	tx = tx.First(&rs)
	return &rs, err
}

func (r *queryResolver) VehicleViolationScoringItems(ctx context.Context, distinctOn []model.VehicleViolationScoringItemsSelectColumn, limit *int, offset *int, orderBy []*model.VehicleViolationScoringItemsOrderBy, where *model.VehicleViolationScoringItemsBoolExp) ([]*model.VehicleViolationScoringItems, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.VehicleViolationScoringItems{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model.VehicleViolationScoringItems
	tx = tx.Find(&rs)
	if err = tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, err
}

func (r *queryResolver) VehicleViolationScoringItemsAggregate(ctx context.Context, distinctOn []model.VehicleViolationScoringItemsSelectColumn, limit *int, offset *int, orderBy []*model.VehicleViolationScoringItemsOrderBy, where *model.VehicleViolationScoringItemsBoolExp) (*model.VehicleViolationScoringItemsAggregate, error) {
	var rs model.VehicleViolationScoringItemsAggregate
	qt := util2.NewQueryTranslator(db.DB, &model.VehicleViolationScoringItems{})
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

func (r *queryResolver) VehicleViolationScoringItemsByPk(ctx context.Context, id int64, violationScoringItemID string) (*model.VehicleViolationScoringItems, error) {
	var rs model.VehicleViolationScoringItems
	err := db.DB.Where("id = ?", id).First(&rs).Error
	return &rs, err
}

func (r *subscriptionResolver) VehicleViolationScoringItems(ctx context.Context, distinctOn []model.VehicleViolationScoringItemsSelectColumn, limit *int, offset *int, orderBy []*model.VehicleViolationScoringItemsOrderBy, where *model.VehicleViolationScoringItemsBoolExp) (<-chan []*model.VehicleViolationScoringItems, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VehicleViolationScoringItemsAggregate(ctx context.Context, distinctOn []model.VehicleViolationScoringItemsSelectColumn, limit *int, offset *int, orderBy []*model.VehicleViolationScoringItemsOrderBy, where *model.VehicleViolationScoringItemsBoolExp) (<-chan *model.VehicleViolationScoringItemsAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VehicleViolationScoringItemsByPk(ctx context.Context, id int64, violationScoringItemID string) (<-chan *model.VehicleViolationScoringItems, error) {
	panic(fmt.Errorf("not implemented"))
}
