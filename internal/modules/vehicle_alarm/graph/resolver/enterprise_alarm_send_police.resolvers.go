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

func (r *mutationResolver) DeleteEnterpriseAlarmSendPolice(ctx context.Context, where model.EnterpriseAlarmSendPoliceBoolExp) (*model.EnterpriseAlarmSendPoliceMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EnterpriseAlarmSendPolice{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.EnterpriseAlarmSendPolice
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
	return &model.EnterpriseAlarmSendPoliceMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteEnterpriseAlarmSendPoliceByPk(ctx context.Context, Id int64) (*model1.EnterpriseAlarmSendPolice, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.EnterpriseAlarmSendPolice
	tx := db.DB.Model(&model1.EnterpriseAlarmSendPolice{})
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

func (r *mutationResolver) InsertEnterpriseAlarmSendPolice(ctx context.Context, objects []*model.EnterpriseAlarmSendPoliceInsertInput) (*model.EnterpriseAlarmSendPoliceMutationResponse, error) {
	rs := make([]*model1.EnterpriseAlarmSendPolice, 0)
	for _, object := range objects {
		v := &model1.EnterpriseAlarmSendPolice{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.EnterpriseAlarmSendPolice{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.EnterpriseAlarmSendPoliceMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertEnterpriseAlarmSendPoliceOne(ctx context.Context, object model.EnterpriseAlarmSendPoliceInsertInput) (*model1.EnterpriseAlarmSendPolice, error) {
	rs := &model1.EnterpriseAlarmSendPolice{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.EnterpriseAlarmSendPolice{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateEnterpriseAlarmSendPolice(ctx context.Context, inc *model.EnterpriseAlarmSendPoliceIncInput, set *model.EnterpriseAlarmSendPoliceSetInput, where model.EnterpriseAlarmSendPoliceBoolExp) (*model.EnterpriseAlarmSendPoliceMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EnterpriseAlarmSendPolice{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.EnterpriseAlarmSendPoliceMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.EnterpriseAlarmSendPoliceMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateEnterpriseAlarmSendPoliceByPk(ctx context.Context, inc *model.EnterpriseAlarmSendPoliceIncInput, set *model.EnterpriseAlarmSendPoliceSetInput, Id int64) (*model1.EnterpriseAlarmSendPolice, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.EnterpriseAlarmSendPolice{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.EnterpriseAlarmSendPolice
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) EnterpriseAlarmSendPolice(ctx context.Context, distinctOn []model.EnterpriseAlarmSendPoliceSelectColumn, limit *int, offset *int, orderBy []*model.EnterpriseAlarmSendPoliceOrderBy, where *model.EnterpriseAlarmSendPoliceBoolExp) ([]*model1.EnterpriseAlarmSendPolice, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EnterpriseAlarmSendPolice{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.EnterpriseAlarmSendPolice
	tx = tx.Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) EnterpriseAlarmSendPoliceAggregate(ctx context.Context, distinctOn []model.EnterpriseAlarmSendPoliceSelectColumn, limit *int, offset *int, orderBy []*model.EnterpriseAlarmSendPoliceOrderBy, where *model.EnterpriseAlarmSendPoliceBoolExp) (*model.EnterpriseAlarmSendPoliceAggregate, error) {
	var rs model.EnterpriseAlarmSendPoliceAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.EnterpriseAlarmSendPolice{})
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

func (r *queryResolver) EnterpriseAlarmSendPoliceByPk(ctx context.Context, Id int64) (*model1.EnterpriseAlarmSendPolice, error) {
	var rs model1.EnterpriseAlarmSendPolice
	tx := db.DB.Model(&model1.EnterpriseAlarmSendPolice{}).First(&rs, Id)
	err := tx.Error
	return &rs, err
}
