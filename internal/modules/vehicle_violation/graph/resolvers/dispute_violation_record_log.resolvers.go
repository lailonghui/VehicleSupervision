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

func (r *mutationResolver) DeleteDisputeViolationRecordLog(ctx context.Context, where model.DisputeViolationRecordLogBoolExp) (*model.DisputeViolationRecordLogMutationResponse, error) {
	var err error
	var rs []*model.DisputeViolationRecordLog
	qt := util2.NewQueryTranslator(db.DB, &model.DisputeViolationRecordLog{})
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
	return &model.DisputeViolationRecordLogMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, err
}

func (r *mutationResolver) DeleteDisputeViolationRecordLogByPk(ctx context.Context, disputeViolationLogID string, id int64) (*model.DisputeViolationRecordLog, error) {
	var rs model.DisputeViolationRecordLog
	var err error
	tx := db.DB.Model(&model.DisputeViolationRecordLog{})
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

func (r *mutationResolver) InsertDisputeViolationRecordLog(ctx context.Context, objects []*model.DisputeViolationRecordLogInsertInput, onConflict *model.DisputeViolationRecordLogOnConflict) (*model.DisputeViolationRecordLogMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertDisputeViolationRecordLogOne(ctx context.Context, object model.DisputeViolationRecordLogInsertInput, onConflict *model.DisputeViolationRecordLogOnConflict) (*model.DisputeViolationRecordLog, error) {
	v := &model.DisputeViolationRecordLog{}
	util.StructAssign(v, &object)
	err := db.DB.Create(v).Error
	if err != nil {
		fmt.Println(err)
	}
	return v, err
}

func (r *mutationResolver) UpdateDisputeViolationRecordLog(ctx context.Context, inc *model.DisputeViolationRecordLogIncInput, set *model.DisputeViolationRecordLogSetInput, where model.DisputeViolationRecordLogBoolExp) (*model.DisputeViolationRecordLogMutationResponse, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.DisputeViolationRecordLog{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return &model.DisputeViolationRecordLogMutationResponse{
			AffectedRows: 0,
		}, err
	}
	var vehicleList []*model.DisputeViolationRecordLog
	tx.Scan(&vehicleList)
	return &model.DisputeViolationRecordLogMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    vehicleList,
	}, err
}

func (r *mutationResolver) UpdateDisputeViolationRecordLogByPk(ctx context.Context, inc *model.DisputeViolationRecordLogIncInput, set *model.DisputeViolationRecordLogSetInput, pkColumns model.DisputeViolationRecordLogPkColumnsInput) (*model.DisputeViolationRecordLog, error) {
	var err error
	tx := db.DB.Where("id = ? ", pkColumns.ID)
	qt := util2.NewQueryTranslator(tx, &model.DisputeViolationRecordLog{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return nil, err
	}
	var rs model.DisputeViolationRecordLog
	tx = tx.First(&rs)
	return &rs, err
}

func (r *queryResolver) DisputeViolationRecordLog(ctx context.Context, distinctOn []model.DisputeViolationRecordLogSelectColumn, limit *int, offset *int, orderBy []*model.DisputeViolationRecordLogOrderBy, where *model.DisputeViolationRecordLogBoolExp) ([]*model.DisputeViolationRecordLog, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.DisputeViolationRecordLog{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model.DisputeViolationRecordLog
	tx = tx.Find(&rs)
	if err = tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, err
}

func (r *queryResolver) DisputeViolationRecordLogAggregate(ctx context.Context, distinctOn []model.DisputeViolationRecordLogSelectColumn, limit *int, offset *int, orderBy []*model.DisputeViolationRecordLogOrderBy, where *model.DisputeViolationRecordLogBoolExp) (*model.DisputeViolationRecordLogAggregate, error) {
	var rs model.DisputeViolationRecordLogAggregate
	qt := util2.NewQueryTranslator(db.DB, &model.DisputeViolationRecordLog{})
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

func (r *queryResolver) DisputeViolationRecordLogByPk(ctx context.Context, disputeViolationLogID string, id int64) (*model.DisputeViolationRecordLog, error) {
	var rs model.DisputeViolationRecordLog
	err := db.DB.Where("id = ?", id).First(&rs).Error
	return &rs, err
}

func (r *subscriptionResolver) DisputeViolationRecordLog(ctx context.Context, distinctOn []model.DisputeViolationRecordLogSelectColumn, limit *int, offset *int, orderBy []*model.DisputeViolationRecordLogOrderBy, where *model.DisputeViolationRecordLogBoolExp) (<-chan []*model.DisputeViolationRecordLog, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DisputeViolationRecordLogAggregate(ctx context.Context, distinctOn []model.DisputeViolationRecordLogSelectColumn, limit *int, offset *int, orderBy []*model.DisputeViolationRecordLogOrderBy, where *model.DisputeViolationRecordLogBoolExp) (<-chan *model.DisputeViolationRecordLogAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DisputeViolationRecordLogByPk(ctx context.Context, disputeViolationLogID string, id int64) (<-chan *model.DisputeViolationRecordLog, error) {
	panic(fmt.Errorf("not implemented"))
}
