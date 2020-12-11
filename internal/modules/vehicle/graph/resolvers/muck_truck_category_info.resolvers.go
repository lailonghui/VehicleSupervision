package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/vehicle/graph/model"
	util2 "VehicleSupervision/pkg/graphql/util"
	"VehicleSupervision/pkg/util"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteMuckTruckCategoryInfo(ctx context.Context, where model.MuckTruckCategoryInfoBoolExp) (*model.MuckTruckCategoryInfoMutationResponse, error) {
	var err error
	var rs []*model.MuckTruckCategoryInfo
	qt := util2.NewQueryTranslator(db.DB, &model.MuckTruckCategoryInfo{})
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
	return &model.MuckTruckCategoryInfoMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, err
}

func (r *mutationResolver) DeleteMuckTruckCategoryInfoByPk(ctx context.Context, id int64) (*model.MuckTruckCategoryInfo, error) {
	var rs model.MuckTruckCategoryInfo
	var err error
	tx := db.DB.Model(&model.MuckTruckCategoryInfo{})
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

func (r *mutationResolver) InsertMuckTruckCategoryInfo(ctx context.Context, objects []*model.MuckTruckCategoryInfoInsertInput, onConflict *model.MuckTruckCategoryInfoOnConflict) (*model.MuckTruckCategoryInfoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertMuckTruckCategoryInfoOne(ctx context.Context, object model.MuckTruckCategoryInfoInsertInput, onConflict *model.MuckTruckCategoryInfoOnConflict) (*model.MuckTruckCategoryInfo, error) {
	v := &model.MuckTruckCategoryInfo{}
	util.StructAssign(v, &object)
	err := db.DB.Create(v).Error
	if err != nil {
		fmt.Println(err)
	}
	return v, err
}

func (r *mutationResolver) UpdateMuckTruckCategoryInfo(ctx context.Context, inc *model.MuckTruckCategoryInfoIncInput, set *model.MuckTruckCategoryInfoSetInput, where model.MuckTruckCategoryInfoBoolExp) (*model.MuckTruckCategoryInfoMutationResponse, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.MuckTruckCategoryInfo{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return &model.MuckTruckCategoryInfoMutationResponse{
			AffectedRows: 0,
		}, err
	}
	var vehicleList []*model.MuckTruckCategoryInfo
	tx.Scan(&vehicleList)
	return &model.MuckTruckCategoryInfoMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    vehicleList,
	}, err
}

func (r *mutationResolver) UpdateMuckTruckCategoryInfoByPk(ctx context.Context, inc *model.MuckTruckCategoryInfoIncInput, set *model.MuckTruckCategoryInfoSetInput, pkColumns model.MuckTruckCategoryInfoPkColumnsInput) (*model.MuckTruckCategoryInfo, error) {
	var err error
	tx := db.DB.Where("id = ? ", pkColumns.ID)
	qt := util2.NewQueryTranslator(tx, &model.MuckTruckCategoryInfo{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return nil, err
	}
	var rs model.MuckTruckCategoryInfo
	tx = tx.First(&rs)
	return &rs, err
}

func (r *queryResolver) MuckTruckCategoryInfo(ctx context.Context, distinctOn []model.MuckTruckCategoryInfoSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckCategoryInfoOrderBy, where *model.MuckTruckCategoryInfoBoolExp) ([]*model.MuckTruckCategoryInfo, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.MuckTruckCategoryInfo{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model.MuckTruckCategoryInfo
	tx = tx.Find(&rs)
	if err = tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, err
}

func (r *queryResolver) MuckTruckCategoryInfoAggregate(ctx context.Context, distinctOn []model.MuckTruckCategoryInfoSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckCategoryInfoOrderBy, where *model.MuckTruckCategoryInfoBoolExp) (*model.MuckTruckCategoryInfoAggregate, error) {
	var rs model.MuckTruckCategoryInfoAggregate
	qt := util2.NewQueryTranslator(db.DB, &model.MuckTruckCategoryInfo{})
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

func (r *queryResolver) MuckTruckCategoryInfoByPk(ctx context.Context, id int64) (*model.MuckTruckCategoryInfo, error) {
	var rs model.MuckTruckCategoryInfo
	err := db.DB.Where("id = ?", id).First(&rs).Error
	return &rs, err
}

func (r *subscriptionResolver) MuckTruckCategoryInfo(ctx context.Context, distinctOn []model.MuckTruckCategoryInfoSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckCategoryInfoOrderBy, where *model.MuckTruckCategoryInfoBoolExp) (<-chan []*model.MuckTruckCategoryInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) MuckTruckCategoryInfoAggregate(ctx context.Context, distinctOn []model.MuckTruckCategoryInfoSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckCategoryInfoOrderBy, where *model.MuckTruckCategoryInfoBoolExp) (<-chan *model.MuckTruckCategoryInfoAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) MuckTruckCategoryInfoByPk(ctx context.Context, id int64) (<-chan *model.MuckTruckCategoryInfo, error) {
	panic(fmt.Errorf("not implemented"))
}
