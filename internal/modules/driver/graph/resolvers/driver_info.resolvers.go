package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/driver/graph/model"
	util2 "VehicleSupervision/pkg/graphql/util"
	"VehicleSupervision/pkg/util"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteDriverInfo(ctx context.Context, where model.DriverInfoBoolExp) (*model.DriverInfoMutationResponse, error) {
	var err error
	var rs []*model.DriverInfo
	qt := util2.NewQueryTranslator(db.DB, &model.DriverInfo{})
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
	return &model.DriverInfoMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, err
}

func (r *mutationResolver) DeleteDriverInfoByPk(ctx context.Context, driverID string, id int64) (*model.DriverInfo, error) {
	var rs model.DriverInfo
	var err error
	tx := db.DB.Model(&model.DriverInfo{})
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

func (r *mutationResolver) InsertDriverInfo(ctx context.Context, objects []*model.DriverInfoInsertInput, onConflict *model.DriverInfoOnConflict) (*model.DriverInfoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertDriverInfoOne(ctx context.Context, object model.DriverInfoInsertInput, onConflict *model.DriverInfoOnConflict) (*model.DriverInfo, error) {
	v := &model.DriverInfo{}
	util.StructAssign(v, &object)
	err := db.DB.Create(v).Error
	if err != nil {
		fmt.Println(err)
	}
	return v, err
}

func (r *mutationResolver) UpdateDriverInfo(ctx context.Context, inc *model.DriverInfoIncInput, set *model.DriverInfoSetInput, where model.DriverInfoBoolExp) (*model.DriverInfoMutationResponse, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.DriverInfo{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return &model.DriverInfoMutationResponse{
			AffectedRows: 0,
		}, err
	}
	var vehicleList []*model.DriverInfo
	tx.Scan(&vehicleList)
	return &model.DriverInfoMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    vehicleList,
	}, err
}

func (r *mutationResolver) UpdateDriverInfoByPk(ctx context.Context, inc *model.DriverInfoIncInput, set *model.DriverInfoSetInput, pkColumns model.DriverInfoPkColumnsInput) (*model.DriverInfo, error) {
	var err error
	tx := db.DB.Where("id = ? ", pkColumns.ID)
	qt := util2.NewQueryTranslator(tx, &model.DriverInfo{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return nil, err
	}
	var rs model.DriverInfo
	tx = tx.First(&rs)
	return &rs, err
}

func (r *queryResolver) DriverInfo(ctx context.Context, distinctOn []model.DriverInfoSelectColumn, limit *int, offset *int, orderBy []*model.DriverInfoOrderBy, where *model.DriverInfoBoolExp) ([]*model.DriverInfo, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.DriverInfo{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model.DriverInfo
	tx = tx.Find(&rs)
	if err = tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, err
}

func (r *queryResolver) DriverInfoAggregate(ctx context.Context, distinctOn []model.DriverInfoSelectColumn, limit *int, offset *int, orderBy []*model.DriverInfoOrderBy, where *model.DriverInfoBoolExp) (*model.DriverInfoAggregate, error) {
	var rs model.DriverInfoAggregate
	qt := util2.NewQueryTranslator(db.DB, &model.DriverInfo{})
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

func (r *queryResolver) DriverInfoByPk(ctx context.Context, driverID string, id int64) (*model.DriverInfo, error) {
	var rs model.DriverInfo
	err := db.DB.Where("id = ?", id).First(&rs).Error
	return &rs, err
}

func (r *subscriptionResolver) DriverInfo(ctx context.Context, distinctOn []model.DriverInfoSelectColumn, limit *int, offset *int, orderBy []*model.DriverInfoOrderBy, where *model.DriverInfoBoolExp) (<-chan []*model.DriverInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DriverInfoAggregate(ctx context.Context, distinctOn []model.DriverInfoSelectColumn, limit *int, offset *int, orderBy []*model.DriverInfoOrderBy, where *model.DriverInfoBoolExp) (<-chan *model.DriverInfoAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DriverInfoByPk(ctx context.Context, driverID string, id int64) (<-chan *model.DriverInfo, error) {
	panic(fmt.Errorf("not implemented"))
}
