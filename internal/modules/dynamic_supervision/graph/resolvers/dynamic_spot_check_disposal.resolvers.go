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

func (r *mutationResolver) DeleteDynamicSpotCheckDisposal(ctx context.Context, where model.DynamicSpotCheckDisposalBoolExp) (*model.DynamicSpotCheckDisposalMutationResponse, error) {
	var err error
	var rs []*model.DynamicSpotCheckDisposal
	qt := util2.NewQueryTranslator(db.DB, &model.DynamicSpotCheckDisposal{})
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
	return &model.DynamicSpotCheckDisposalMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, err
}

func (r *mutationResolver) DeleteDynamicSpotCheckDisposalByPk(ctx context.Context, id int64) (*model.DynamicSpotCheckDisposal, error) {
	var rs model.DynamicSpotCheckDisposal
	var err error
	tx := db.DB.Model(&model.DynamicSpotCheckDisposal{})
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

func (r *mutationResolver) InsertDynamicSpotCheckDisposal(ctx context.Context, objects []*model.DynamicSpotCheckDisposalInsertInput, onConflict *model.DynamicSpotCheckDisposalOnConflict) (*model.DynamicSpotCheckDisposalMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertDynamicSpotCheckDisposalOne(ctx context.Context, object model.DynamicSpotCheckDisposalInsertInput, onConflict *model.DynamicSpotCheckDisposalOnConflict) (*model.DynamicSpotCheckDisposal, error) {
	v := &model.DynamicSpotCheckDisposal{}
	util.StructAssign(v, &object)
	err := db.DB.Create(v).Error
	if err != nil {
		fmt.Println(err)
	}
	return v, err
}

func (r *mutationResolver) UpdateDynamicSpotCheckDisposal(ctx context.Context, inc *model.DynamicSpotCheckDisposalIncInput, set *model.DynamicSpotCheckDisposalSetInput, where model.DynamicSpotCheckDisposalBoolExp) (*model.DynamicSpotCheckDisposalMutationResponse, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.DynamicSpotCheckDisposal{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return &model.DynamicSpotCheckDisposalMutationResponse{
			AffectedRows: 0,
		}, err
	}
	var vehicleList []*model.DynamicSpotCheckDisposal
	tx.Scan(&vehicleList)
	return &model.DynamicSpotCheckDisposalMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    vehicleList,
	}, err
}

func (r *mutationResolver) UpdateDynamicSpotCheckDisposalByPk(ctx context.Context, inc *model.DynamicSpotCheckDisposalIncInput, set *model.DynamicSpotCheckDisposalSetInput, pkColumns model.DynamicSpotCheckDisposalPkColumnsInput) (*model.DynamicSpotCheckDisposal, error) {
	var err error
	tx := db.DB.Where("id = ? ", pkColumns.ID)
	qt := util2.NewQueryTranslator(tx, &model.DynamicSpotCheckDisposal{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return nil, err
	}
	var rs model.DynamicSpotCheckDisposal
	tx = tx.First(&rs)
	return &rs, err
}

func (r *queryResolver) DynamicSpotCheckDisposal(ctx context.Context, distinctOn []model.DynamicSpotCheckDisposalSelectColumn, limit *int, offset *int, orderBy []*model.DynamicSpotCheckDisposalOrderBy, where *model.DynamicSpotCheckDisposalBoolExp) ([]*model.DynamicSpotCheckDisposal, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.DynamicSpotCheckDisposal{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model.DynamicSpotCheckDisposal
	tx = tx.Find(&rs)
	if err = tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, err
}

func (r *queryResolver) DynamicSpotCheckDisposalAggregate(ctx context.Context, distinctOn []model.DynamicSpotCheckDisposalSelectColumn, limit *int, offset *int, orderBy []*model.DynamicSpotCheckDisposalOrderBy, where *model.DynamicSpotCheckDisposalBoolExp) (*model.DynamicSpotCheckDisposalAggregate, error) {
	var rs model.DynamicSpotCheckDisposalAggregate
	qt := util2.NewQueryTranslator(db.DB, &model.DynamicSpotCheckDisposal{})
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

func (r *queryResolver) DynamicSpotCheckDisposalByPk(ctx context.Context, id int64) (*model.DynamicSpotCheckDisposal, error) {
	var rs model.DynamicSpotCheckDisposal
	err := db.DB.Where("id = ?", id).First(&rs).Error
	return &rs, err
}

func (r *subscriptionResolver) DynamicSpotCheckDisposal(ctx context.Context, distinctOn []model.DynamicSpotCheckDisposalSelectColumn, limit *int, offset *int, orderBy []*model.DynamicSpotCheckDisposalOrderBy, where *model.DynamicSpotCheckDisposalBoolExp) (<-chan []*model.DynamicSpotCheckDisposal, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DynamicSpotCheckDisposalAggregate(ctx context.Context, distinctOn []model.DynamicSpotCheckDisposalSelectColumn, limit *int, offset *int, orderBy []*model.DynamicSpotCheckDisposalOrderBy, where *model.DynamicSpotCheckDisposalBoolExp) (<-chan *model.DynamicSpotCheckDisposalAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DynamicSpotCheckDisposalByPk(ctx context.Context, id int64) (<-chan *model.DynamicSpotCheckDisposal, error) {
	panic(fmt.Errorf("not implemented"))
}
