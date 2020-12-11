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

func (r *mutationResolver) DeleteEnterpriseDeductionItems(ctx context.Context, where model.EnterpriseDeductionItemsBoolExp) (*model.EnterpriseDeductionItemsMutationResponse, error) {
	var err error
	var rs []*model.EnterpriseDeductionItems
	qt := util2.NewQueryTranslator(db.DB, &model.EnterpriseDeductionItems{})
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
	return &model.EnterpriseDeductionItemsMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, err
}

func (r *mutationResolver) DeleteEnterpriseDeductionItemsByPk(ctx context.Context, enterpriseDeductionItemID string, id int64) (*model.EnterpriseDeductionItems, error) {
	var rs model.EnterpriseDeductionItems
	var err error
	tx := db.DB.Model(&model.EnterpriseDeductionItems{})
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

func (r *mutationResolver) InsertEnterpriseDeductionItems(ctx context.Context, objects []*model.EnterpriseDeductionItemsInsertInput, onConflict *model.EnterpriseDeductionItemsOnConflict) (*model.EnterpriseDeductionItemsMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertEnterpriseDeductionItemsOne(ctx context.Context, object model.EnterpriseDeductionItemsInsertInput, onConflict *model.EnterpriseDeductionItemsOnConflict) (*model.EnterpriseDeductionItems, error) {
	v := &model.EnterpriseDeductionItems{}
	util.StructAssign(v, &object)
	err := db.DB.Create(v).Error
	if err != nil {
		fmt.Println(err)
	}
	return v, err
}

func (r *mutationResolver) UpdateEnterpriseDeductionItems(ctx context.Context, inc *model.EnterpriseDeductionItemsIncInput, set *model.EnterpriseDeductionItemsSetInput, where model.EnterpriseDeductionItemsBoolExp) (*model.EnterpriseDeductionItemsMutationResponse, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.EnterpriseDeductionItems{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return &model.EnterpriseDeductionItemsMutationResponse{
			AffectedRows: 0,
		}, err
	}
	var vehicleList []*model.EnterpriseDeductionItems
	tx.Scan(&vehicleList)
	return &model.EnterpriseDeductionItemsMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    vehicleList,
	}, err
}

func (r *mutationResolver) UpdateEnterpriseDeductionItemsByPk(ctx context.Context, inc *model.EnterpriseDeductionItemsIncInput, set *model.EnterpriseDeductionItemsSetInput, pkColumns model.EnterpriseDeductionItemsPkColumnsInput) (*model.EnterpriseDeductionItems, error) {
	var err error
	tx := db.DB.Where("id = ? ", pkColumns.ID)
	qt := util2.NewQueryTranslator(tx, &model.EnterpriseDeductionItems{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return nil, err
	}
	var rs model.EnterpriseDeductionItems
	tx = tx.First(&rs)
	return &rs, err
}

func (r *queryResolver) EnterpriseDeductionItems(ctx context.Context, distinctOn []model.EnterpriseDeductionItemsSelectColumn, limit *int, offset *int, orderBy []*model.EnterpriseDeductionItemsOrderBy, where *model.EnterpriseDeductionItemsBoolExp) ([]*model.EnterpriseDeductionItems, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.EnterpriseDeductionItems{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model.EnterpriseDeductionItems
	tx = tx.Find(&rs)
	if err = tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, err
}

func (r *queryResolver) EnterpriseDeductionItemsAggregate(ctx context.Context, distinctOn []model.EnterpriseDeductionItemsSelectColumn, limit *int, offset *int, orderBy []*model.EnterpriseDeductionItemsOrderBy, where *model.EnterpriseDeductionItemsBoolExp) (*model.EnterpriseDeductionItemsAggregate, error) {
	var rs model.EnterpriseDeductionItemsAggregate
	qt := util2.NewQueryTranslator(db.DB, &model.EnterpriseDeductionItems{})
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

func (r *queryResolver) EnterpriseDeductionItemsByPk(ctx context.Context, enterpriseDeductionItemID string, id int64) (*model.EnterpriseDeductionItems, error) {
	var rs model.EnterpriseDeductionItems
	err := db.DB.Where("id = ?", id).First(&rs).Error
	return &rs, err
}

func (r *subscriptionResolver) EnterpriseDeductionItems(ctx context.Context, distinctOn []model.EnterpriseDeductionItemsSelectColumn, limit *int, offset *int, orderBy []*model.EnterpriseDeductionItemsOrderBy, where *model.EnterpriseDeductionItemsBoolExp) (<-chan []*model.EnterpriseDeductionItems, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) EnterpriseDeductionItemsAggregate(ctx context.Context, distinctOn []model.EnterpriseDeductionItemsSelectColumn, limit *int, offset *int, orderBy []*model.EnterpriseDeductionItemsOrderBy, where *model.EnterpriseDeductionItemsBoolExp) (<-chan *model.EnterpriseDeductionItemsAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) EnterpriseDeductionItemsByPk(ctx context.Context, enterpriseDeductionItemID string, id int64) (<-chan *model.EnterpriseDeductionItems, error) {
	panic(fmt.Errorf("not implemented"))
}
