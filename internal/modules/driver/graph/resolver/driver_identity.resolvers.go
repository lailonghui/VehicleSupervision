package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/driver/graph/model"
	model1 "VehicleSupervision/internal/modules/driver/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteDriverIdentity(ctx context.Context, where model.DriverIdentityBoolExp) (*model.DriverIdentityMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DriverIdentity{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.DriverIdentity
	if len(preloads) > 0 {
		// 如果请求的字段不为空，则先查询一遍数据库
		tx := tx.Select(preloads)
		tx = tx.Find(&rs)
		// 如果查询结果含有错误，则返回错误
		if err := tx.Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, nil
			}
			return nil, err
		}
	}
	// 删除
	tx = tx.Delete(nil)
	if err := tx.Error; err != nil {
		return nil, err
	}
	return &model.DriverIdentityMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteDriverIdentityByPk(ctx context.Context, Id int64) (*model1.DriverIdentity, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.DriverIdentity
	tx := db.DB.Model(&model1.DriverIdentity{})
	if len(preloads) > 0 {
		// 如果请求的字段不为空，则先查询一遍数据库
		tx = tx.Select(preloads).Where("id = ?", Id).First(&rs)
		// 如果查询结果含有错误，则返回错误
		if err := tx.Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, nil
			}
			return nil, err
		}
	}
	// 删除
	tx = tx.Delete(nil)
	if err := tx.Error; err != nil {
		return nil, err
	}
	return &rs, nil
}

func (r *mutationResolver) InsertDriverIdentity(ctx context.Context, objects []*model.DriverIdentityInsertInput) (*model.DriverIdentityMutationResponse, error) {
	rs := make([]*model1.DriverIdentity, 0)
	for _, object := range objects {
		v := &model1.DriverIdentity{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.DriverIdentity{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.DriverIdentityMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertDriverIdentityOne(ctx context.Context, object model.DriverIdentityInsertInput) (*model1.DriverIdentity, error) {
	rs := &model1.DriverIdentity{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.DriverIdentity{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateDriverIdentity(ctx context.Context, inc *model.DriverIdentityIncInput, set *model.DriverIdentitySetInput, where model.DriverIdentityBoolExp) (*model.DriverIdentityMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DriverIdentity{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.DriverIdentityMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.DriverIdentityMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateDriverIdentityByPk(ctx context.Context, inc *model.DriverIdentityIncInput, set *model.DriverIdentitySetInput, Id int64) (*model1.DriverIdentity, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.DriverIdentity{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.DriverIdentity
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) DriverIdentity(ctx context.Context, distinctOn []model.DriverIdentitySelectColumn, limit *int, offset *int, orderBy []*model.DriverIdentityOrderBy, where *model.DriverIdentityBoolExp) ([]*model1.DriverIdentity, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DriverIdentity{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.DriverIdentity
	tx = tx.Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) DriverIdentityAggregate(ctx context.Context, distinctOn []model.DriverIdentitySelectColumn, limit *int, offset *int, orderBy []*model.DriverIdentityOrderBy, where *model.DriverIdentityBoolExp) (*model.DriverIdentityAggregate, error) {
	var rs model.DriverIdentityAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.DriverIdentity{})
	tx, err := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Aggregate(&rs, ctx)
	if err != nil {
		return nil, err
	}
	err = tx.Error
	return &rs, err
}

func (r *queryResolver) DriverIdentityByPk(ctx context.Context, Id int64) (*model1.DriverIdentity, error) {
	var rs model1.DriverIdentity
	tx := db.DB.Model(&model1.DriverIdentity{}).First(&rs, Id)
	err := tx.Error
	return &rs, err
}
