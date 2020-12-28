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

func (r *mutationResolver) DeleteDistrictAlarmContentPush(ctx context.Context, where model.DistrictAlarmContentPushBoolExp) (*model.DistrictAlarmContentPushMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DistrictAlarmContentPush{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.DistrictAlarmContentPush
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
	return &model.DistrictAlarmContentPushMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteDistrictAlarmContentPushByPk(ctx context.Context, Id int64) (*model1.DistrictAlarmContentPush, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.DistrictAlarmContentPush
	tx := db.DB.Model(&model1.DistrictAlarmContentPush{})
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

func (r *mutationResolver) InsertDistrictAlarmContentPush(ctx context.Context, objects []*model.DistrictAlarmContentPushInsertInput) (*model.DistrictAlarmContentPushMutationResponse, error) {
	rs := []*model1.DistrictAlarmContentPush{}
	for _, object := range objects {
		v := &model1.DistrictAlarmContentPush{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.DistrictAlarmContentPush{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.DistrictAlarmContentPushMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertDistrictAlarmContentPushOne(ctx context.Context, object model.DistrictAlarmContentPushInsertInput) (*model1.DistrictAlarmContentPush, error) {
	rs := &model1.DistrictAlarmContentPush{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.DistrictAlarmContentPush{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateDistrictAlarmContentPush(ctx context.Context, inc *model.DistrictAlarmContentPushIncInput, set *model.DistrictAlarmContentPushSetInput, where model.DistrictAlarmContentPushBoolExp) (*model.DistrictAlarmContentPushMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DistrictAlarmContentPush{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.DistrictAlarmContentPushMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.DistrictAlarmContentPushMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateDistrictAlarmContentPushByPk(ctx context.Context, inc *model.DistrictAlarmContentPushIncInput, set *model.DistrictAlarmContentPushSetInput, Id int64) (*model1.DistrictAlarmContentPush, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.DistrictAlarmContentPush{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.DistrictAlarmContentPush
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) DistrictAlarmContentPush(ctx context.Context, distinctOn []model.DistrictAlarmContentPushSelectColumn, limit *int, offset *int, orderBy []*model.DistrictAlarmContentPushOrderBy, where *model.DistrictAlarmContentPushBoolExp) ([]*model1.DistrictAlarmContentPush, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DistrictAlarmContentPush{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.DistrictAlarmContentPush
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) DistrictAlarmContentPushAggregate(ctx context.Context, distinctOn []model.DistrictAlarmContentPushSelectColumn, limit *int, offset *int, orderBy []*model.DistrictAlarmContentPushOrderBy, where *model.DistrictAlarmContentPushBoolExp) (*model.DistrictAlarmContentPushAggregate, error) {
	var rs model.DistrictAlarmContentPushAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.DistrictAlarmContentPush{})
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

func (r *queryResolver) DistrictAlarmContentPushByPk(ctx context.Context, Id int64) (*model1.DistrictAlarmContentPush, error) {
	var rs model1.DistrictAlarmContentPush
	tx := db.DB.Model(&model1.DistrictAlarmContentPush{}).First(&rs, Id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
