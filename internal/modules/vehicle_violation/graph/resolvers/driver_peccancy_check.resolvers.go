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

func (r *mutationResolver) DeleteDriverPeccancyCheck(ctx context.Context, where model.DriverPeccancyCheckBoolExp) (*model.DriverPeccancyCheckMutationResponse, error) {
	var err error
	var rs []*model.DriverPeccancyCheck
	qt := util2.NewQueryTranslator(db.DB, &model.DriverPeccancyCheck{})
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
	return &model.DriverPeccancyCheckMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, err
}

func (r *mutationResolver) DeleteDriverPeccancyCheckByPk(ctx context.Context, id int64) (*model.DriverPeccancyCheck, error) {
	var rs model.DriverPeccancyCheck
	var err error
	tx := db.DB.Model(&model.DriverPeccancyCheck{})
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

func (r *mutationResolver) InsertDriverPeccancyCheck(ctx context.Context, objects []*model.DriverPeccancyCheckInsertInput, onConflict *model.DriverPeccancyCheckOnConflict) (*model.DriverPeccancyCheckMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertDriverPeccancyCheckOne(ctx context.Context, object model.DriverPeccancyCheckInsertInput, onConflict *model.DriverPeccancyCheckOnConflict) (*model.DriverPeccancyCheck, error) {
	v := &model.DriverPeccancyCheck{}
	util.StructAssign(v, &object)
	err := db.DB.Create(v).Error
	if err != nil {
		fmt.Println(err)
	}
	return v, err
}

func (r *mutationResolver) UpdateDriverPeccancyCheck(ctx context.Context, inc *model.DriverPeccancyCheckIncInput, set *model.DriverPeccancyCheckSetInput, where model.DriverPeccancyCheckBoolExp) (*model.DriverPeccancyCheckMutationResponse, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.DriverPeccancyCheck{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return &model.DriverPeccancyCheckMutationResponse{
			AffectedRows: 0,
		}, err
	}
	var vehicleList []*model.DriverPeccancyCheck
	tx.Scan(&vehicleList)
	return &model.DriverPeccancyCheckMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    vehicleList,
	}, err
}

func (r *mutationResolver) UpdateDriverPeccancyCheckByPk(ctx context.Context, inc *model.DriverPeccancyCheckIncInput, set *model.DriverPeccancyCheckSetInput, pkColumns model.DriverPeccancyCheckPkColumnsInput) (*model.DriverPeccancyCheck, error) {
	var err error
	tx := db.DB.Where("id = ? ", pkColumns.ID)
	qt := util2.NewQueryTranslator(tx, &model.DriverPeccancyCheck{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return nil, err
	}
	var rs model.DriverPeccancyCheck
	tx = tx.First(&rs)
	return &rs, err
}

func (r *queryResolver) DriverPeccancyCheck(ctx context.Context, distinctOn []model.DriverPeccancyCheckSelectColumn, limit *int, offset *int, orderBy []*model.DriverPeccancyCheckOrderBy, where *model.DriverPeccancyCheckBoolExp) ([]*model.DriverPeccancyCheck, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.DriverPeccancyCheck{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model.DriverPeccancyCheck
	tx = tx.Find(&rs)
	if err = tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, err
}

func (r *queryResolver) DriverPeccancyCheckAggregate(ctx context.Context, distinctOn []model.DriverPeccancyCheckSelectColumn, limit *int, offset *int, orderBy []*model.DriverPeccancyCheckOrderBy, where *model.DriverPeccancyCheckBoolExp) (*model.DriverPeccancyCheckAggregate, error) {
	var rs model.DriverPeccancyCheckAggregate
	qt := util2.NewQueryTranslator(db.DB, &model.DriverPeccancyCheck{})
	tx, err := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Aggregate(&rs, ctx)
	return nil, err
	if err != nil {
		fmt.Println(err)
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

func (r *queryResolver) DriverPeccancyCheckByPk(ctx context.Context, id int64) (*model.DriverPeccancyCheck, error) {
	var rs model.DriverPeccancyCheck
	err := db.DB.Where("id = ?", id).First(&rs).Error
	return &rs, err
}

func (r *subscriptionResolver) DriverPeccancyCheck(ctx context.Context, distinctOn []model.DriverPeccancyCheckSelectColumn, limit *int, offset *int, orderBy []*model.DriverPeccancyCheckOrderBy, where *model.DriverPeccancyCheckBoolExp) (<-chan []*model.DriverPeccancyCheck, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DriverPeccancyCheckAggregate(ctx context.Context, distinctOn []model.DriverPeccancyCheckSelectColumn, limit *int, offset *int, orderBy []*model.DriverPeccancyCheckOrderBy, where *model.DriverPeccancyCheckBoolExp) (<-chan *model.DriverPeccancyCheckAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DriverPeccancyCheckByPk(ctx context.Context, id int64) (<-chan *model.DriverPeccancyCheck, error) {
	panic(fmt.Errorf("not implemented"))
}
