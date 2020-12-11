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

func (r *mutationResolver) DeleteUserOperationLog(ctx context.Context, where model.UserOperationLogBoolExp) (*model.UserOperationLogMutationResponse, error) {
	var err error
	var rs []*model.UserOperationLog
	qt := util2.NewQueryTranslator(db.DB, &model.UserOperationLog{})
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
	return &model.UserOperationLogMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, err
}

func (r *mutationResolver) DeleteUserOperationLogByPk(ctx context.Context, id int64) (*model.UserOperationLog, error) {
	var rs model.UserOperationLog
	var err error
	tx := db.DB.Model(&model.UserOperationLog{})
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

func (r *mutationResolver) InsertUserOperationLog(ctx context.Context, objects []*model.UserOperationLogInsertInput, onConflict *model.UserOperationLogOnConflict) (*model.UserOperationLogMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertUserOperationLogOne(ctx context.Context, object model.UserOperationLogInsertInput, onConflict *model.UserOperationLogOnConflict) (*model.UserOperationLog, error) {
	v := &model.UserOperationLog{}
	util.StructAssign(v, &object)
	err := db.DB.Create(v).Error
	if err != nil {
		fmt.Println(err)
	}
	return v, err
}

func (r *mutationResolver) UpdateUserOperationLog(ctx context.Context, inc *model.UserOperationLogIncInput, set *model.UserOperationLogSetInput, where model.UserOperationLogBoolExp) (*model.UserOperationLogMutationResponse, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.UserOperationLog{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return &model.UserOperationLogMutationResponse{
			AffectedRows: 0,
		}, err
	}
	var vehicleList []*model.UserOperationLog
	tx.Scan(&vehicleList)
	return &model.UserOperationLogMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    vehicleList,
	}, err
}

func (r *mutationResolver) UpdateUserOperationLogByPk(ctx context.Context, inc *model.UserOperationLogIncInput, set *model.UserOperationLogSetInput, pkColumns model.UserOperationLogPkColumnsInput) (*model.UserOperationLog, error) {
	var err error
	tx := db.DB.Where("id = ? ", pkColumns.ID)
	qt := util2.NewQueryTranslator(tx, &model.UserOperationLog{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return nil, err
	}
	var rs model.UserOperationLog
	tx = tx.First(&rs)
	return &rs, err
}

func (r *queryResolver) UserOperationLog(ctx context.Context, distinctOn []model.UserOperationLogSelectColumn, limit *int, offset *int, orderBy []*model.UserOperationLogOrderBy, where *model.UserOperationLogBoolExp) ([]*model.UserOperationLog, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.UserOperationLog{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model.UserOperationLog
	tx = tx.Find(&rs)
	if err = tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, err
}

func (r *queryResolver) UserOperationLogAggregate(ctx context.Context, distinctOn []model.UserOperationLogSelectColumn, limit *int, offset *int, orderBy []*model.UserOperationLogOrderBy, where *model.UserOperationLogBoolExp) (*model.UserOperationLogAggregate, error) {
	var rs model.UserOperationLogAggregate
	qt := util2.NewQueryTranslator(db.DB, &model.UserOperationLog{})
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

func (r *queryResolver) UserOperationLogByPk(ctx context.Context, id int64) (*model.UserOperationLog, error) {
	var rs model.UserOperationLog
	err := db.DB.Where("id = ?", id).First(&rs).Error
	return &rs, err
}

func (r *subscriptionResolver) UserOperationLog(ctx context.Context, distinctOn []model.UserOperationLogSelectColumn, limit *int, offset *int, orderBy []*model.UserOperationLogOrderBy, where *model.UserOperationLogBoolExp) (<-chan []*model.UserOperationLog, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) UserOperationLogAggregate(ctx context.Context, distinctOn []model.UserOperationLogSelectColumn, limit *int, offset *int, orderBy []*model.UserOperationLogOrderBy, where *model.UserOperationLogBoolExp) (<-chan *model.UserOperationLogAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) UserOperationLogByPk(ctx context.Context, id int64) (<-chan *model.UserOperationLog, error) {
	panic(fmt.Errorf("not implemented"))
}
