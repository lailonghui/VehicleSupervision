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

func (r *mutationResolver) DeleteDriverInfoChangeLog(ctx context.Context, where model.DriverInfoChangeLogBoolExp) (*model.DriverInfoChangeLogMutationResponse, error) {
	var err error
	var rs []*model.DriverInfoChangeLog
	qt := util2.NewQueryTranslator(db.DB, &model.DriverInfoChangeLog{})
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
	return &model.DriverInfoChangeLogMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, err
}

func (r *mutationResolver) DeleteDriverInfoChangeLogByPk(ctx context.Context, driverID string, id int64) (*model.DriverInfoChangeLog, error) {
	var rs model.DriverInfoChangeLog
	var err error
	tx := db.DB.Model(&model.DriverInfoChangeLog{})
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

func (r *mutationResolver) InsertDriverInfoChangeLog(ctx context.Context, objects []*model.DriverInfoChangeLogInsertInput, onConflict *model.DriverInfoChangeLogOnConflict) (*model.DriverInfoChangeLogMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertDriverInfoChangeLogOne(ctx context.Context, object model.DriverInfoChangeLogInsertInput, onConflict *model.DriverInfoChangeLogOnConflict) (*model.DriverInfoChangeLog, error) {
	v := &model.DriverInfoChangeLog{}
	util.StructAssign(v, &object)
	err := db.DB.Create(v).Error
	if err != nil {
		fmt.Println(err)
	}
	return v, err
}

func (r *mutationResolver) UpdateDriverInfoChangeLog(ctx context.Context, inc *model.DriverInfoChangeLogIncInput, set *model.DriverInfoChangeLogSetInput, where model.DriverInfoChangeLogBoolExp) (*model.DriverInfoChangeLogMutationResponse, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.DriverInfoChangeLog{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return &model.DriverInfoChangeLogMutationResponse{
			AffectedRows: 0,
		}, err
	}
	var vehicleList []*model.DriverInfoChangeLog
	tx.Scan(&vehicleList)
	return &model.DriverInfoChangeLogMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    vehicleList,
	}, err
}

func (r *mutationResolver) UpdateDriverInfoChangeLogByPk(ctx context.Context, inc *model.DriverInfoChangeLogIncInput, set *model.DriverInfoChangeLogSetInput, pkColumns model.DriverInfoChangeLogPkColumnsInput) (*model.DriverInfoChangeLog, error) {
	var err error
	tx := db.DB.Where("id = ? ", pkColumns.ID)
	qt := util2.NewQueryTranslator(tx, &model.DriverInfoChangeLog{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return nil, err
	}
	var rs model.DriverInfoChangeLog
	tx = tx.First(&rs)
	return &rs, err
}

func (r *queryResolver) DriverInfoChangeLog(ctx context.Context, distinctOn []model.DriverInfoChangeLogSelectColumn, limit *int, offset *int, orderBy []*model.DriverInfoChangeLogOrderBy, where *model.DriverInfoChangeLogBoolExp) ([]*model.DriverInfoChangeLog, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.DriverInfoChangeLog{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model.DriverInfoChangeLog
	tx = tx.Find(&rs)
	if err = tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, err
}

func (r *queryResolver) DriverInfoChangeLogAggregate(ctx context.Context, distinctOn []model.DriverInfoChangeLogSelectColumn, limit *int, offset *int, orderBy []*model.DriverInfoChangeLogOrderBy, where *model.DriverInfoChangeLogBoolExp) (*model.DriverInfoChangeLogAggregate, error) {
	var rs model.DriverInfoChangeLogAggregate
	qt := util2.NewQueryTranslator(db.DB, &model.DriverInfoChangeLog{})
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

func (r *queryResolver) DriverInfoChangeLogByPk(ctx context.Context, driverID string, id int64) (*model.DriverInfoChangeLog, error) {
	var rs model.DriverInfoChangeLog
	err := db.DB.Where("id = ?", id).First(&rs).Error
	return &rs, err
}

func (r *subscriptionResolver) DriverInfoChangeLog(ctx context.Context, distinctOn []model.DriverInfoChangeLogSelectColumn, limit *int, offset *int, orderBy []*model.DriverInfoChangeLogOrderBy, where *model.DriverInfoChangeLogBoolExp) (<-chan []*model.DriverInfoChangeLog, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DriverInfoChangeLogAggregate(ctx context.Context, distinctOn []model.DriverInfoChangeLogSelectColumn, limit *int, offset *int, orderBy []*model.DriverInfoChangeLogOrderBy, where *model.DriverInfoChangeLogBoolExp) (<-chan *model.DriverInfoChangeLogAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DriverInfoChangeLogByPk(ctx context.Context, driverID string, id int64) (<-chan *model.DriverInfoChangeLog, error) {
	panic(fmt.Errorf("not implemented"))
}
