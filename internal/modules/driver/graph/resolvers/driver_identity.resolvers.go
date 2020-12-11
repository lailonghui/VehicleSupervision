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

func (r *mutationResolver) DeleteDriverIdentity(ctx context.Context, where model.DriverIdentityBoolExp) (*model.DriverIdentityMutationResponse, error) {
	var err error
	var rs []*model.DriverIdentity
	qt := util2.NewQueryTranslator(db.DB, &model.DriverIdentity{})
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
	return &model.DriverIdentityMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, err
}

func (r *mutationResolver) DeleteDriverIdentityByPk(ctx context.Context, id int64, identityID string) (*model.DriverIdentity, error) {
	var rs model.DriverIdentity
	var err error
	tx := db.DB.Model(&model.DriverIdentity{})
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

func (r *mutationResolver) InsertDriverIdentity(ctx context.Context, objects []*model.DriverIdentityInsertInput, onConflict *model.DriverIdentityOnConflict) (*model.DriverIdentityMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertDriverIdentityOne(ctx context.Context, object model.DriverIdentityInsertInput, onConflict *model.DriverIdentityOnConflict) (*model.DriverIdentity, error) {
	v := &model.DriverIdentity{}
	util.StructAssign(v, &object)
	err := db.DB.Create(v).Error
	if err != nil {
		fmt.Println(err)
	}
	return v, err
}

func (r *mutationResolver) UpdateDriverIdentity(ctx context.Context, inc *model.DriverIdentityIncInput, set *model.DriverIdentitySetInput, where model.DriverIdentityBoolExp) (*model.DriverIdentityMutationResponse, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.DriverIdentity{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return &model.DriverIdentityMutationResponse{
			AffectedRows: 0,
		}, err
	}
	var vehicleList []*model.DriverIdentity
	tx.Scan(&vehicleList)
	return &model.DriverIdentityMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    vehicleList,
	}, err
}

func (r *mutationResolver) UpdateDriverIdentityByPk(ctx context.Context, inc *model.DriverIdentityIncInput, set *model.DriverIdentitySetInput, pkColumns model.DriverIdentityPkColumnsInput) (*model.DriverIdentity, error) {
	var err error
	tx := db.DB.Where("id = ? ", pkColumns.ID)
	qt := util2.NewQueryTranslator(tx, &model.DriverIdentity{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return nil, err
	}
	var rs model.DriverIdentity
	tx = tx.First(&rs)
	return &rs, err
}

func (r *queryResolver) DriverIdentity(ctx context.Context, distinctOn []model.DriverIdentitySelectColumn, limit *int, offset *int, orderBy []*model.DriverIdentityOrderBy, where *model.DriverIdentityBoolExp) ([]*model.DriverIdentity, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.DriverIdentity{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model.DriverIdentity
	tx = tx.Find(&rs)
	if err = tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, err
}

func (r *queryResolver) DriverIdentityAggregate(ctx context.Context, distinctOn []model.DriverIdentitySelectColumn, limit *int, offset *int, orderBy []*model.DriverIdentityOrderBy, where *model.DriverIdentityBoolExp) (*model.DriverIdentityAggregate, error) {
	var rs model.DriverIdentityAggregate
	qt := util2.NewQueryTranslator(db.DB, &model.DriverIdentity{})
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

func (r *queryResolver) DriverIdentityByPk(ctx context.Context, id int64, identityID string) (*model.DriverIdentity, error) {
	var rs model.DriverIdentity
	err := db.DB.Where("id = ?", id).First(&rs).Error
	return &rs, err
}

func (r *subscriptionResolver) DriverIdentity(ctx context.Context, distinctOn []model.DriverIdentitySelectColumn, limit *int, offset *int, orderBy []*model.DriverIdentityOrderBy, where *model.DriverIdentityBoolExp) (<-chan []*model.DriverIdentity, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DriverIdentityAggregate(ctx context.Context, distinctOn []model.DriverIdentitySelectColumn, limit *int, offset *int, orderBy []*model.DriverIdentityOrderBy, where *model.DriverIdentityBoolExp) (<-chan *model.DriverIdentityAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DriverIdentityByPk(ctx context.Context, id int64, identityID string) (<-chan *model.DriverIdentity, error) {
	panic(fmt.Errorf("not implemented"))
}
