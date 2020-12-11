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

func (r *mutationResolver) DeleteRegionalViolationRegister(ctx context.Context, where model.RegionalViolationRegisterBoolExp) (*model.RegionalViolationRegisterMutationResponse, error) {
	var err error
	var rs []*model.RegionalViolationRegister
	qt := util2.NewQueryTranslator(db.DB, &model.RegionalViolationRegister{})
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
	return &model.RegionalViolationRegisterMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, err
}

func (r *mutationResolver) DeleteRegionalViolationRegisterByPk(ctx context.Context, id int64, regionalViolationRegisterID string) (*model.RegionalViolationRegister, error) {
	var rs model.RegionalViolationRegister
	var err error
	tx := db.DB.Model(&model.RegionalViolationRegister{})
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

func (r *mutationResolver) InsertRegionalViolationRegister(ctx context.Context, objects []*model.RegionalViolationRegisterInsertInput, onConflict *model.RegionalViolationRegisterOnConflict) (*model.RegionalViolationRegisterMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertRegionalViolationRegisterOne(ctx context.Context, object model.RegionalViolationRegisterInsertInput, onConflict *model.RegionalViolationRegisterOnConflict) (*model.RegionalViolationRegister, error) {
	v := &model.RegionalViolationRegister{}
	util.StructAssign(v, &object)
	err := db.DB.Create(v).Error
	if err != nil {
		fmt.Println(err)
	}
	return v, err
}

func (r *mutationResolver) UpdateRegionalViolationRegister(ctx context.Context, inc *model.RegionalViolationRegisterIncInput, set *model.RegionalViolationRegisterSetInput, where model.RegionalViolationRegisterBoolExp) (*model.RegionalViolationRegisterMutationResponse, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.RegionalViolationRegister{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return &model.RegionalViolationRegisterMutationResponse{
			AffectedRows: 0,
		}, err
	}
	var vehicleList []*model.RegionalViolationRegister
	tx.Scan(&vehicleList)
	return &model.RegionalViolationRegisterMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    vehicleList,
	}, err
}

func (r *mutationResolver) UpdateRegionalViolationRegisterByPk(ctx context.Context, inc *model.RegionalViolationRegisterIncInput, set *model.RegionalViolationRegisterSetInput, pkColumns model.RegionalViolationRegisterPkColumnsInput) (*model.RegionalViolationRegister, error) {
	var err error
	tx := db.DB.Where("id = ? ", pkColumns.ID)
	qt := util2.NewQueryTranslator(tx, &model.RegionalViolationRegister{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return nil, err
	}
	var rs model.RegionalViolationRegister
	tx = tx.First(&rs)
	return &rs, err
}

func (r *queryResolver) RegionalViolationRegister(ctx context.Context, distinctOn []model.RegionalViolationRegisterSelectColumn, limit *int, offset *int, orderBy []*model.RegionalViolationRegisterOrderBy, where *model.RegionalViolationRegisterBoolExp) ([]*model.RegionalViolationRegister, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.RegionalViolationRegister{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model.RegionalViolationRegister
	tx = tx.Find(&rs)
	if err = tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, err
}

func (r *queryResolver) RegionalViolationRegisterAggregate(ctx context.Context, distinctOn []model.RegionalViolationRegisterSelectColumn, limit *int, offset *int, orderBy []*model.RegionalViolationRegisterOrderBy, where *model.RegionalViolationRegisterBoolExp) (*model.RegionalViolationRegisterAggregate, error) {
	var rs model.RegionalViolationRegisterAggregate
	qt := util2.NewQueryTranslator(db.DB, &model.RegionalViolationRegister{})
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

func (r *queryResolver) RegionalViolationRegisterByPk(ctx context.Context, id int64, regionalViolationRegisterID string) (*model.RegionalViolationRegister, error) {
	var rs model.RegionalViolationRegister
	err := db.DB.Where("id = ?", id).First(&rs).Error
	return &rs, err
}

func (r *subscriptionResolver) RegionalViolationRegister(ctx context.Context, distinctOn []model.RegionalViolationRegisterSelectColumn, limit *int, offset *int, orderBy []*model.RegionalViolationRegisterOrderBy, where *model.RegionalViolationRegisterBoolExp) (<-chan []*model.RegionalViolationRegister, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) RegionalViolationRegisterAggregate(ctx context.Context, distinctOn []model.RegionalViolationRegisterSelectColumn, limit *int, offset *int, orderBy []*model.RegionalViolationRegisterOrderBy, where *model.RegionalViolationRegisterBoolExp) (<-chan *model.RegionalViolationRegisterAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) RegionalViolationRegisterByPk(ctx context.Context, id int64, regionalViolationRegisterID string) (<-chan *model.RegionalViolationRegister, error) {
	panic(fmt.Errorf("not implemented"))
}
