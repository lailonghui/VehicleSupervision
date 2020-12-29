package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/admin/graph/model"
	model1 "VehicleSupervision/internal/modules/admin/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteSystemUser(ctx context.Context, where model.SystemUserBoolExp) (*model.SystemUserMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.SystemUser{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.SystemUser
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
	return &model.SystemUserMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteSystemUserByPk(ctx context.Context, Id int64) (*model1.SystemUser, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.SystemUser
	tx := db.DB.Model(&model1.SystemUser{})
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

func (r *mutationResolver) InsertSystemUser(ctx context.Context, objects []*model.SystemUserInsertInput) (*model.SystemUserMutationResponse, error) {
	rs := make([]*model1.SystemUser, 0)
	for _, object := range objects {
		v := &model1.SystemUser{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.SystemUser{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.SystemUserMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertSystemUserOne(ctx context.Context, object model.SystemUserInsertInput) (*model1.SystemUser, error) {
	rs := &model1.SystemUser{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.SystemUser{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateSystemUser(ctx context.Context, inc *model.SystemUserIncInput, set *model.SystemUserSetInput, where model.SystemUserBoolExp) (*model.SystemUserMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.SystemUser{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.SystemUserMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.SystemUser
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
	return &model.SystemUserMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateSystemUserByPk(ctx context.Context, inc *model.SystemUserIncInput, set *model.SystemUserSetInput, Id int64) (*model1.SystemUser, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.SystemUser{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.SystemUser
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) SystemUser(ctx context.Context, distinctOn []model.SystemUserSelectColumn, limit *int, offset *int, orderBy []*model.SystemUserOrderBy, where *model.SystemUserBoolExp) ([]*model1.SystemUser, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.SystemUser{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.SystemUser
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) SystemUserAggregate(ctx context.Context, distinctOn []model.SystemUserSelectColumn, limit *int, offset *int, orderBy []*model.SystemUserOrderBy, where *model.SystemUserBoolExp) (*model.SystemUserAggregate, error) {
	var rs model.SystemUserAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.SystemUser{})
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

func (r *queryResolver) SystemUserByPk(ctx context.Context, Id int64) (*model1.SystemUser, error) {
	var rs model1.SystemUser
	tx := db.DB.Model(&model1.SystemUser{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}
