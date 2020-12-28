package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/vehicle_alarm/graph/generated"
	"VehicleSupervision/internal/modules/vehicle_alarm/graph/model"
	model1 "VehicleSupervision/internal/modules/vehicle_alarm/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteVehicleAlarmData(ctx context.Context, where model.VehicleAlarmDataBoolExp) (*model.VehicleAlarmDataMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleAlarmData{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.VehicleAlarmData
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
	return &model.VehicleAlarmDataMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteVehicleAlarmDataByPk(ctx context.Context, Id int64) (*model1.VehicleAlarmData, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.VehicleAlarmData
	tx := db.DB.Model(&model1.VehicleAlarmData{})
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

func (r *mutationResolver) InsertVehicleAlarmData(ctx context.Context, objects []*model.VehicleAlarmDataInsertInput) (*model.VehicleAlarmDataMutationResponse, error) {
	rs := make([]*model1.VehicleAlarmData, 0)
	for _, object := range objects {
		v := &model1.VehicleAlarmData{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.VehicleAlarmData{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.VehicleAlarmDataMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertVehicleAlarmDataOne(ctx context.Context, object model.VehicleAlarmDataInsertInput) (*model1.VehicleAlarmData, error) {
	rs := &model1.VehicleAlarmData{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.VehicleAlarmData{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateVehicleAlarmData(ctx context.Context, inc *model.VehicleAlarmDataIncInput, set *model.VehicleAlarmDataSetInput, where model.VehicleAlarmDataBoolExp) (*model.VehicleAlarmDataMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleAlarmData{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.VehicleAlarmDataMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.VehicleAlarmDataMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateVehicleAlarmDataByPk(ctx context.Context, inc *model.VehicleAlarmDataIncInput, set *model.VehicleAlarmDataSetInput, Id int64) (*model1.VehicleAlarmData, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.VehicleAlarmData{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.VehicleAlarmData
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) VehicleAlarmData(ctx context.Context, distinctOn []model.VehicleAlarmDataSelectColumn, limit *int, offset *int, orderBy []*model.VehicleAlarmDataOrderBy, where *model.VehicleAlarmDataBoolExp) ([]*model1.VehicleAlarmData, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleAlarmData{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.VehicleAlarmData
	tx = tx.Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) VehicleAlarmDataAggregate(ctx context.Context, distinctOn []model.VehicleAlarmDataSelectColumn, limit *int, offset *int, orderBy []*model.VehicleAlarmDataOrderBy, where *model.VehicleAlarmDataBoolExp) (*model.VehicleAlarmDataAggregate, error) {
	var rs model.VehicleAlarmDataAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.VehicleAlarmData{})
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

func (r *queryResolver) VehicleAlarmDataByPk(ctx context.Context, Id int64) (*model1.VehicleAlarmData, error) {
	var rs model1.VehicleAlarmData
	tx := db.DB.Model(&model1.VehicleAlarmData{}).First(&rs, Id)
	err := tx.Error
	return &rs, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
