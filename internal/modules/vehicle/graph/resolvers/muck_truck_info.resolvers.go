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

func (r *mutationResolver) DeleteMuckTruckInfo(ctx context.Context, where model.MuckTruckInfoBoolExp) (*model.MuckTruckInfoMutationResponse, error) {
	var err error
	var rs []*model.MuckTruckInfo
	qt := util2.NewQueryTranslator(db.DB, &model.MuckTruckInfo{})
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
	return &model.MuckTruckInfoMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteMuckTruckInfoByPk(ctx context.Context, muckTruckID int64) (*model.MuckTruckInfo, error) {
	var rs model.MuckTruckInfo
	var err error
	tx := db.DB.Model(&model.MuckTruckInfo{})
	preloads := util2.GetPreloads(ctx)
	if len(preloads) > 0 {
		// 如果请求的字段不为空，则先查询一遍数据库
		tx = tx.Select(preloads).Where("muck_truck_id = ? ", muckTruckID).First(&rs)
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

func (r *mutationResolver) InsertMuckTruckInfo(ctx context.Context, objects []*model.MuckTruckInfoInsertInput, onConflict *model.MuckTruckInfoOnConflict) (*model.MuckTruckInfoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertMuckTruckInfoOne(ctx context.Context, object model.MuckTruckInfoInsertInput, onConflict *model.MuckTruckInfoOnConflict) (*model.MuckTruckInfo, error) {
	v := &model.MuckTruckInfo{}
	util.StructAssign(v, &object)
	err := db.DB.Create(v).Error
	if err != nil {
		fmt.Println(err)
	}
	return v, err
}

func (r *mutationResolver) UpdateMuckTruckInfo(ctx context.Context, inc *model.MuckTruckInfoIncInput, set *model.MuckTruckInfoSetInput, where model.MuckTruckInfoBoolExp) (*model.MuckTruckInfoMutationResponse, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.MuckTruckInfo{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return &model.MuckTruckInfoMutationResponse{
			AffectedRows: 0,
		}, err
	}
	var vehicleList []*model.MuckTruckInfo
	tx.Scan(&vehicleList)
	return &model.MuckTruckInfoMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    vehicleList,
	}, err
}

func (r *mutationResolver) UpdateMuckTruckInfoByPk(ctx context.Context, inc *model.MuckTruckInfoIncInput, set *model.MuckTruckInfoSetInput, pkColumns model.MuckTruckInfoPkColumnsInput) (*model.MuckTruckInfo, error) {
	var err error
	tx := db.DB.Where("muck_truck_id = ? ", pkColumns.MuckTruckID)
	qt := util2.NewQueryTranslator(tx, &model.MuckTruckInfo{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return nil, err
	}
	var rs model.MuckTruckInfo
	tx = tx.First(&rs)
	return &rs, err
}

func (r *queryResolver) MuckTruckInfo(ctx context.Context, distinctOn []model.MuckTruckInfoSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckInfoOrderBy, where *model.MuckTruckInfoBoolExp) ([]*model.MuckTruckInfo, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.MuckTruckInfo{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model.MuckTruckInfo
	tx = tx.Find(&rs)
	if err = tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, err
}

func (r *queryResolver) MuckTruckInfoAggregate(ctx context.Context, distinctOn []model.MuckTruckInfoSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckInfoOrderBy, where *model.MuckTruckInfoBoolExp) (*model.MuckTruckInfoAggregate, error) {
	var rs model.MuckTruckInfoAggregate
	qt := util2.NewQueryTranslator(db.DB, &model.MuckTruckInfo{})
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

func (r *queryResolver) MuckTruckInfoByPk(ctx context.Context, muckTruckID int64) (*model.MuckTruckInfo, error) {
	var rs model.MuckTruckInfo
	err := db.DB.Where("muck_truck_id = ?", muckTruckID).First(&rs).Error
	return &rs, err
}

func (r *subscriptionResolver) MuckTruckInfo(ctx context.Context, distinctOn []model.MuckTruckInfoSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckInfoOrderBy, where *model.MuckTruckInfoBoolExp) (<-chan []*model.MuckTruckInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) MuckTruckInfoAggregate(ctx context.Context, distinctOn []model.MuckTruckInfoSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckInfoOrderBy, where *model.MuckTruckInfoBoolExp) (<-chan *model.MuckTruckInfoAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) MuckTruckInfoByPk(ctx context.Context, muckTruckID int64) (<-chan *model.MuckTruckInfo, error) {
	panic(fmt.Errorf("not implemented"))
}
