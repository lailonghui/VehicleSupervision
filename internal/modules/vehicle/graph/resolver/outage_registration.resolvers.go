package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/vehicle/graph/model"
	model1 "VehicleSupervision/internal/modules/vehicle/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteOutageRegistration(ctx context.Context, where model.OutageRegistrationBoolExp) (*model.OutageRegistrationMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.OutageRegistration{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.OutageRegistration
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
	return &model.OutageRegistrationMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteOutageRegistrationByPk(ctx context.Context, Id int64) (*model1.OutageRegistration, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.OutageRegistration
	tx := db.DB.Model(&model1.OutageRegistration{})
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

func (r *mutationResolver) InsertOutageRegistration(ctx context.Context, objects []*model.OutageRegistrationInsertInput) (*model.OutageRegistrationMutationResponse, error) {
	rs := []*model1.OutageRegistration{}
	for _, object := range objects {
		v := &model1.OutageRegistration{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.OutageRegistration{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.OutageRegistrationMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertOutageRegistrationOne(ctx context.Context, object model.OutageRegistrationInsertInput) (*model1.OutageRegistration, error) {
	rs := &model1.OutageRegistration{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.OutageRegistration{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateOutageRegistration(ctx context.Context, inc *model.OutageRegistrationIncInput, set *model.OutageRegistrationSetInput, where model.OutageRegistrationBoolExp) (*model.OutageRegistrationMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.OutageRegistration{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.OutageRegistrationMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.OutageRegistrationMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateOutageRegistrationByPk(ctx context.Context, inc *model.OutageRegistrationIncInput, set *model.OutageRegistrationSetInput, Id int64) (*model1.OutageRegistration, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.OutageRegistration{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.OutageRegistration
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) OutageRegistration(ctx context.Context, distinctOn []model.OutageRegistrationSelectColumn, limit *int, offset *int, orderBy []*model.OutageRegistrationOrderBy, where *model.OutageRegistrationBoolExp) ([]*model1.OutageRegistration, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.OutageRegistration{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.OutageRegistration
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) OutageRegistrationAggregate(ctx context.Context, distinctOn []model.OutageRegistrationSelectColumn, limit *int, offset *int, orderBy []*model.OutageRegistrationOrderBy, where *model.OutageRegistrationBoolExp) (*model.OutageRegistrationAggregate, error) {
	var rs model.OutageRegistrationAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.OutageRegistration{})
	tx, err := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Aggregate(&rs, ctx)
	if err != nil {
		return nil, err
	}
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &rs, nil
}

func (r *queryResolver) OutageRegistrationByPk(ctx context.Context, Id int64) (*model1.OutageRegistration, error) {
	var rs model1.OutageRegistration
	tx := db.DB.Model(&model1.OutageRegistration{}).First(&rs, Id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
