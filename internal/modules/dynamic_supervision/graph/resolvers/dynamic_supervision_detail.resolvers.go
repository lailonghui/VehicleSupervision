package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/dynamic_supervision/graph/model"
	util2 "VehicleSupervision/pkg/graphql/util"
	"VehicleSupervision/pkg/util"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteDynamicSupervisionDetail(ctx context.Context, where model.DynamicSupervisionDetailBoolExp) (*model.DynamicSupervisionDetailMutationResponse, error) {
	var err error
	var rs []*model.DynamicSupervisionDetail
	qt := util2.NewQueryTranslator(db.DB, &model.DynamicSupervisionDetail{})
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
	return &model.DynamicSupervisionDetailMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, err
}

func (r *mutationResolver) DeleteDynamicSupervisionDetailByPk(ctx context.Context, id int64, supervisionDetailID string) (*model.DynamicSupervisionDetail, error) {
	var rs model.DynamicSupervisionDetail
	var err error
	tx := db.DB.Model(&model.DynamicSupervisionDetail{})
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

func (r *mutationResolver) InsertDynamicSupervisionDetail(ctx context.Context, objects []*model.DynamicSupervisionDetailInsertInput, onConflict *model.DynamicSupervisionDetailOnConflict) (*model.DynamicSupervisionDetailMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertDynamicSupervisionDetailOne(ctx context.Context, object model.DynamicSupervisionDetailInsertInput, onConflict *model.DynamicSupervisionDetailOnConflict) (*model.DynamicSupervisionDetail, error) {
	v := &model.DynamicSupervisionDetail{}
	util.StructAssign(v, &object)
	err := db.DB.Create(v).Error
	if err != nil {
		fmt.Println(err)
	}
	return v, err
}

func (r *mutationResolver) UpdateDynamicSupervisionDetail(ctx context.Context, inc *model.DynamicSupervisionDetailIncInput, set *model.DynamicSupervisionDetailSetInput, where model.DynamicSupervisionDetailBoolExp) (*model.DynamicSupervisionDetailMutationResponse, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.DynamicSupervisionDetail{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return &model.DynamicSupervisionDetailMutationResponse{
			AffectedRows: 0,
		}, err
	}
	var vehicleList []*model.DynamicSupervisionDetail
	tx.Scan(&vehicleList)
	return &model.DynamicSupervisionDetailMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    vehicleList,
	}, err
}

func (r *mutationResolver) UpdateDynamicSupervisionDetailByPk(ctx context.Context, inc *model.DynamicSupervisionDetailIncInput, set *model.DynamicSupervisionDetailSetInput, pkColumns model.DynamicSupervisionDetailPkColumnsInput) (*model.DynamicSupervisionDetail, error) {
	var err error
	tx := db.DB.Where("id = ? ", pkColumns.ID)
	qt := util2.NewQueryTranslator(tx, &model.DynamicSupervisionDetail{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return nil, err
	}
	var rs model.DynamicSupervisionDetail
	tx = tx.First(&rs)
	return &rs, err
}

func (r *queryResolver) DynamicSupervisionDetail(ctx context.Context, distinctOn []model.DynamicSupervisionDetailSelectColumn, limit *int, offset *int, orderBy []*model.DynamicSupervisionDetailOrderBy, where *model.DynamicSupervisionDetailBoolExp) ([]*model.DynamicSupervisionDetail, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.DynamicSupervisionDetail{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model.DynamicSupervisionDetail
	tx = tx.Find(&rs)
	if err = tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, err
}

func (r *queryResolver) DynamicSupervisionDetailAggregate(ctx context.Context, distinctOn []model.DynamicSupervisionDetailSelectColumn, limit *int, offset *int, orderBy []*model.DynamicSupervisionDetailOrderBy, where *model.DynamicSupervisionDetailBoolExp) (*model.DynamicSupervisionDetailAggregate, error) {
	var rs model.DynamicSupervisionDetailAggregate
	qt := util2.NewQueryTranslator(db.DB, &model.DynamicSupervisionDetail{})
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

func (r *queryResolver) DynamicSupervisionDetailByPk(ctx context.Context, id int64, supervisionDetailID string) (*model.DynamicSupervisionDetail, error) {
	var rs model.DynamicSupervisionDetail
	err := db.DB.Where("id = ?", id).First(&rs).Error
	return &rs, err
}

func (r *subscriptionResolver) DynamicSupervisionDetail(ctx context.Context, distinctOn []model.DynamicSupervisionDetailSelectColumn, limit *int, offset *int, orderBy []*model.DynamicSupervisionDetailOrderBy, where *model.DynamicSupervisionDetailBoolExp) (<-chan []*model.DynamicSupervisionDetail, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DynamicSupervisionDetailAggregate(ctx context.Context, distinctOn []model.DynamicSupervisionDetailSelectColumn, limit *int, offset *int, orderBy []*model.DynamicSupervisionDetailOrderBy, where *model.DynamicSupervisionDetailBoolExp) (<-chan *model.DynamicSupervisionDetailAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DynamicSupervisionDetailByPk(ctx context.Context, id int64, supervisionDetailID string) (<-chan *model.DynamicSupervisionDetail, error) {
	panic(fmt.Errorf("not implemented"))
}
