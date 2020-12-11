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

func (r *mutationResolver) DeleteEnterpriseDeductionOperationRecord(ctx context.Context, where model.EnterpriseDeductionOperationRecordBoolExp) (*model.EnterpriseDeductionOperationRecordMutationResponse, error) {
	var err error
	var rs []*model.EnterpriseDeductionOperationRecord
	qt := util2.NewQueryTranslator(db.DB, &model.EnterpriseDeductionOperationRecord{})
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
	return &model.EnterpriseDeductionOperationRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, err
}

func (r *mutationResolver) DeleteEnterpriseDeductionOperationRecordByPk(ctx context.Context, enterpriseDuductionOperationID string, id int64) (*model.EnterpriseDeductionOperationRecord, error) {
	var rs model.EnterpriseDeductionOperationRecord
	var err error
	tx := db.DB.Model(&model.EnterpriseDeductionOperationRecord{})
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

func (r *mutationResolver) InsertEnterpriseDeductionOperationRecord(ctx context.Context, objects []*model.EnterpriseDeductionOperationRecordInsertInput, onConflict *model.EnterpriseDeductionOperationRecordOnConflict) (*model.EnterpriseDeductionOperationRecordMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertEnterpriseDeductionOperationRecordOne(ctx context.Context, object model.EnterpriseDeductionOperationRecordInsertInput, onConflict *model.EnterpriseDeductionOperationRecordOnConflict) (*model.EnterpriseDeductionOperationRecord, error) {
	v := &model.EnterpriseDeductionOperationRecord{}
	util.StructAssign(v, &object)
	err := db.DB.Create(v).Error
	if err != nil {
		fmt.Println(err)
	}
	return v, err
}

func (r *mutationResolver) UpdateEnterpriseDeductionOperationRecord(ctx context.Context, inc *model.EnterpriseDeductionOperationRecordIncInput, set *model.EnterpriseDeductionOperationRecordSetInput, where model.EnterpriseDeductionOperationRecordBoolExp) (*model.EnterpriseDeductionOperationRecordMutationResponse, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.EnterpriseDeductionOperationRecord{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return &model.EnterpriseDeductionOperationRecordMutationResponse{
			AffectedRows: 0,
		}, err
	}
	var vehicleList []*model.EnterpriseDeductionOperationRecord
	tx.Scan(&vehicleList)
	return &model.EnterpriseDeductionOperationRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    vehicleList,
	}, err
}

func (r *mutationResolver) UpdateEnterpriseDeductionOperationRecordByPk(ctx context.Context, inc *model.EnterpriseDeductionOperationRecordIncInput, set *model.EnterpriseDeductionOperationRecordSetInput, pkColumns model.EnterpriseDeductionOperationRecordPkColumnsInput) (*model.EnterpriseDeductionOperationRecord, error) {
	var err error
	tx := db.DB.Where("id = ? ", pkColumns.ID)
	qt := util2.NewQueryTranslator(tx, &model.EnterpriseDeductionOperationRecord{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return nil, err
	}
	var rs model.EnterpriseDeductionOperationRecord
	tx = tx.First(&rs)
	return &rs, err
}

func (r *queryResolver) EnterpriseDeductionOperationRecord(ctx context.Context, distinctOn []model.EnterpriseDeductionOperationRecordSelectColumn, limit *int, offset *int, orderBy []*model.EnterpriseDeductionOperationRecordOrderBy, where *model.EnterpriseDeductionOperationRecordBoolExp) ([]*model.EnterpriseDeductionOperationRecord, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.EnterpriseDeductionOperationRecord{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model.EnterpriseDeductionOperationRecord
	tx = tx.Find(&rs)
	if err = tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, err
}

func (r *queryResolver) EnterpriseDeductionOperationRecordAggregate(ctx context.Context, distinctOn []model.EnterpriseDeductionOperationRecordSelectColumn, limit *int, offset *int, orderBy []*model.EnterpriseDeductionOperationRecordOrderBy, where *model.EnterpriseDeductionOperationRecordBoolExp) (*model.EnterpriseDeductionOperationRecordAggregate, error) {
	var rs model.EnterpriseDeductionOperationRecordAggregate
	qt := util2.NewQueryTranslator(db.DB, &model.EnterpriseDeductionOperationRecord{})
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

func (r *queryResolver) EnterpriseDeductionOperationRecordByPk(ctx context.Context, enterpriseDuductionOperationID string, id int64) (*model.EnterpriseDeductionOperationRecord, error) {
	var rs model.EnterpriseDeductionOperationRecord
	err := db.DB.Where("id = ?", id).First(&rs).Error
	return &rs, err
}

func (r *subscriptionResolver) EnterpriseDeductionOperationRecord(ctx context.Context, distinctOn []model.EnterpriseDeductionOperationRecordSelectColumn, limit *int, offset *int, orderBy []*model.EnterpriseDeductionOperationRecordOrderBy, where *model.EnterpriseDeductionOperationRecordBoolExp) (<-chan []*model.EnterpriseDeductionOperationRecord, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) EnterpriseDeductionOperationRecordAggregate(ctx context.Context, distinctOn []model.EnterpriseDeductionOperationRecordSelectColumn, limit *int, offset *int, orderBy []*model.EnterpriseDeductionOperationRecordOrderBy, where *model.EnterpriseDeductionOperationRecordBoolExp) (<-chan *model.EnterpriseDeductionOperationRecordAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) EnterpriseDeductionOperationRecordByPk(ctx context.Context, enterpriseDuductionOperationID string, id int64) (<-chan *model.EnterpriseDeductionOperationRecord, error) {
	panic(fmt.Errorf("not implemented"))
}
