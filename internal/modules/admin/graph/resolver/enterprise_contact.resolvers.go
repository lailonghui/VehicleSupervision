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

func (r *mutationResolver) DeleteEnterpriseContact(ctx context.Context, where model.EnterpriseContactBoolExp) (*model.EnterpriseContactMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EnterpriseContact{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.EnterpriseContact
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
	return &model.EnterpriseContactMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteEnterpriseContactByPk(ctx context.Context, Id int64) (*model1.EnterpriseContact, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.EnterpriseContact
	tx := db.DB.Model(&model1.EnterpriseContact{})
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

func (r *mutationResolver) InsertEnterpriseContact(ctx context.Context, objects []*model.EnterpriseContactInsertInput) (*model.EnterpriseContactMutationResponse, error) {
	rs := make([]*model1.EnterpriseContact, 0)
	for _, object := range objects {
		v := &model1.EnterpriseContact{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.EnterpriseContact{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.EnterpriseContactMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertEnterpriseContactOne(ctx context.Context, object model.EnterpriseContactInsertInput) (*model1.EnterpriseContact, error) {
	rs := &model1.EnterpriseContact{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.EnterpriseContact{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateEnterpriseContact(ctx context.Context, inc *model.EnterpriseContactIncInput, set *model.EnterpriseContactSetInput, where model.EnterpriseContactBoolExp) (*model.EnterpriseContactMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EnterpriseContact{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.EnterpriseContactMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.EnterpriseContact
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
	return &model.EnterpriseContactMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateEnterpriseContactByPk(ctx context.Context, inc *model.EnterpriseContactIncInput, set *model.EnterpriseContactSetInput, Id int64) (*model1.EnterpriseContact, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.EnterpriseContact{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.EnterpriseContact
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) EnterpriseContact(ctx context.Context, distinctOn []model.EnterpriseContactSelectColumn, limit *int, offset *int, orderBy []*model.EnterpriseContactOrderBy, where *model.EnterpriseContactBoolExp) ([]*model1.EnterpriseContact, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EnterpriseContact{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.EnterpriseContact
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) EnterpriseContactAggregate(ctx context.Context, distinctOn []model.EnterpriseContactSelectColumn, limit *int, offset *int, orderBy []*model.EnterpriseContactOrderBy, where *model.EnterpriseContactBoolExp) (*model.EnterpriseContactAggregate, error) {
	var rs model.EnterpriseContactAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.EnterpriseContact{})
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

func (r *queryResolver) EnterpriseContactByPk(ctx context.Context, Id int64) (*model1.EnterpriseContact, error) {
	var rs model1.EnterpriseContact
	tx := db.DB.Model(&model1.EnterpriseContact{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}
