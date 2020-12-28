package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/vehicle_alarm/graph/model"
	model1 "VehicleSupervision/internal/modules/vehicle_alarm/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteOfflineAlarmRegistration(ctx context.Context, where model.OfflineAlarmRegistrationBoolExp) (*model.OfflineAlarmRegistrationMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.OfflineAlarmRegistration{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.OfflineAlarmRegistration
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
	return &model.OfflineAlarmRegistrationMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteOfflineAlarmRegistrationByPk(ctx context.Context, Id int64) (*model1.OfflineAlarmRegistration, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.OfflineAlarmRegistration
	tx := db.DB.Model(&model1.OfflineAlarmRegistration{})
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

func (r *mutationResolver) InsertOfflineAlarmRegistration(ctx context.Context, objects []*model.OfflineAlarmRegistrationInsertInput) (*model.OfflineAlarmRegistrationMutationResponse, error) {
	rs := []*model1.OfflineAlarmRegistration{}
	for _, object := range objects {
		v := &model1.OfflineAlarmRegistration{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.OfflineAlarmRegistration{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.OfflineAlarmRegistrationMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertOfflineAlarmRegistrationOne(ctx context.Context, object model.OfflineAlarmRegistrationInsertInput) (*model1.OfflineAlarmRegistration, error) {
	rs := &model1.OfflineAlarmRegistration{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.OfflineAlarmRegistration{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateOfflineAlarmRegistration(ctx context.Context, inc *model.OfflineAlarmRegistrationIncInput, set *model.OfflineAlarmRegistrationSetInput, where model.OfflineAlarmRegistrationBoolExp) (*model.OfflineAlarmRegistrationMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.OfflineAlarmRegistration{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.OfflineAlarmRegistrationMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.OfflineAlarmRegistrationMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateOfflineAlarmRegistrationByPk(ctx context.Context, inc *model.OfflineAlarmRegistrationIncInput, set *model.OfflineAlarmRegistrationSetInput, Id int64) (*model1.OfflineAlarmRegistration, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.OfflineAlarmRegistration{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.OfflineAlarmRegistration
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) OfflineAlarmRegistration(ctx context.Context, distinctOn []model.OfflineAlarmRegistrationSelectColumn, limit *int, offset *int, orderBy []*model.OfflineAlarmRegistrationOrderBy, where *model.OfflineAlarmRegistrationBoolExp) ([]*model1.OfflineAlarmRegistration, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.OfflineAlarmRegistration{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.OfflineAlarmRegistration
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) OfflineAlarmRegistrationAggregate(ctx context.Context, distinctOn []model.OfflineAlarmRegistrationSelectColumn, limit *int, offset *int, orderBy []*model.OfflineAlarmRegistrationOrderBy, where *model.OfflineAlarmRegistrationBoolExp) (*model.OfflineAlarmRegistrationAggregate, error) {
	var rs model.OfflineAlarmRegistrationAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.OfflineAlarmRegistration{})
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

func (r *queryResolver) OfflineAlarmRegistrationByPk(ctx context.Context, Id int64) (*model1.OfflineAlarmRegistration, error) {
	var rs model1.OfflineAlarmRegistration
	tx := db.DB.Model(&model1.OfflineAlarmRegistration{}).First(&rs, Id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
