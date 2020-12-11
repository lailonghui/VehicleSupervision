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

func (r *mutationResolver) DeleteDisputeViolationRecord(ctx context.Context, where model.DisputeViolationRecordBoolExp) (*model.DisputeViolationRecordMutationResponse, error) {
	var err error
	var rs []*model.DisputeViolationRecord
	qt := util2.NewQueryTranslator(db.DB, &model.DisputeViolationRecord{})
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
	return &model.DisputeViolationRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, err
}

func (r *mutationResolver) DeleteDisputeViolationRecordByPk(ctx context.Context, disputeViolationID string, id int64) (*model.DisputeViolationRecord, error) {
	var rs model.DisputeViolationRecord
	var err error
	tx := db.DB.Model(&model.DisputeViolationRecord{})
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

func (r *mutationResolver) InsertDisputeViolationRecord(ctx context.Context, objects []*model.DisputeViolationRecordInsertInput, onConflict *model.DisputeViolationRecordOnConflict) (*model.DisputeViolationRecordMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertDisputeViolationRecordOne(ctx context.Context, object model.DisputeViolationRecordInsertInput, onConflict *model.DisputeViolationRecordOnConflict) (*model.DisputeViolationRecord, error) {
	v := &model.DisputeViolationRecord{}
	util.StructAssign(v, &object)
	err := db.DB.Create(v).Error
	if err != nil {
		fmt.Println(err)
	}
	return v, err
}

func (r *mutationResolver) UpdateDisputeViolationRecord(ctx context.Context, inc *model.DisputeViolationRecordIncInput, set *model.DisputeViolationRecordSetInput, where model.DisputeViolationRecordBoolExp) (*model.DisputeViolationRecordMutationResponse, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.DisputeViolationRecord{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return &model.DisputeViolationRecordMutationResponse{
			AffectedRows: 0,
		}, err
	}
	var vehicleList []*model.DisputeViolationRecord
	tx.Scan(&vehicleList)
	return &model.DisputeViolationRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    vehicleList,
	}, err
}

func (r *mutationResolver) UpdateDisputeViolationRecordByPk(ctx context.Context, inc *model.DisputeViolationRecordIncInput, set *model.DisputeViolationRecordSetInput, pkColumns model.DisputeViolationRecordPkColumnsInput) (*model.DisputeViolationRecord, error) {
	var err error
	tx := db.DB.Where("id = ? ", pkColumns.ID)
	qt := util2.NewQueryTranslator(tx, &model.DisputeViolationRecord{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return nil, err
	}
	var rs model.DisputeViolationRecord
	tx = tx.First(&rs)
	return &rs, err
}

func (r *queryResolver) DisputeViolationRecord(ctx context.Context, distinctOn []model.DisputeViolationRecordSelectColumn, limit *int, offset *int, orderBy []*model.DisputeViolationRecordOrderBy, where *model.DisputeViolationRecordBoolExp) ([]*model.DisputeViolationRecord, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.DisputeViolationRecord{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model.DisputeViolationRecord
	tx = tx.Find(&rs)
	if err = tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, err
}

func (r *queryResolver) DisputeViolationRecordAggregate(ctx context.Context, distinctOn []model.DisputeViolationRecordSelectColumn, limit *int, offset *int, orderBy []*model.DisputeViolationRecordOrderBy, where *model.DisputeViolationRecordBoolExp) (*model.DisputeViolationRecordAggregate, error) {
	var rs model.DisputeViolationRecordAggregate
	qt := util2.NewQueryTranslator(db.DB, &model.DisputeViolationRecord{})
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

func (r *queryResolver) DisputeViolationRecordByPk(ctx context.Context, disputeViolationID string, id int64) (*model.DisputeViolationRecord, error) {
	var rs model.DisputeViolationRecord
	err := db.DB.Where("id = ?", id).First(&rs).Error
	return &rs, err
}

func (r *subscriptionResolver) DisputeViolationRecord(ctx context.Context, distinctOn []model.DisputeViolationRecordSelectColumn, limit *int, offset *int, orderBy []*model.DisputeViolationRecordOrderBy, where *model.DisputeViolationRecordBoolExp) (<-chan []*model.DisputeViolationRecord, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DisputeViolationRecordAggregate(ctx context.Context, distinctOn []model.DisputeViolationRecordSelectColumn, limit *int, offset *int, orderBy []*model.DisputeViolationRecordOrderBy, where *model.DisputeViolationRecordBoolExp) (<-chan *model.DisputeViolationRecordAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DisputeViolationRecordByPk(ctx context.Context, disputeViolationID string, id int64) (<-chan *model.DisputeViolationRecord, error) {
	panic(fmt.Errorf("not implemented"))
}
