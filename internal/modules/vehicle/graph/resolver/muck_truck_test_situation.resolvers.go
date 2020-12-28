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

func (r *mutationResolver) DeleteMuckTruckTestSituation(ctx context.Context, where model.MuckTruckTestSituationBoolExp) (*model.MuckTruckTestSituationMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.MuckTruckTestSituation{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.MuckTruckTestSituation
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
	return &model.MuckTruckTestSituationMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteMuckTruckTestSituationByPk(ctx context.Context, Id int64) (*model1.MuckTruckTestSituation, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.MuckTruckTestSituation
	tx := db.DB.Model(&model1.MuckTruckTestSituation{})
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

func (r *mutationResolver) InsertMuckTruckTestSituation(ctx context.Context, objects []*model.MuckTruckTestSituationInsertInput) (*model.MuckTruckTestSituationMutationResponse, error) {
	rs := make([]*model1.MuckTruckTestSituation, 0)
	for _, object := range objects {
		v := &model1.MuckTruckTestSituation{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.MuckTruckTestSituation{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.MuckTruckTestSituationMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertMuckTruckTestSituationOne(ctx context.Context, object model.MuckTruckTestSituationInsertInput) (*model1.MuckTruckTestSituation, error) {
	rs := &model1.MuckTruckTestSituation{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.MuckTruckTestSituation{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateMuckTruckTestSituation(ctx context.Context, inc *model.MuckTruckTestSituationIncInput, set *model.MuckTruckTestSituationSetInput, where model.MuckTruckTestSituationBoolExp) (*model.MuckTruckTestSituationMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.MuckTruckTestSituation{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.MuckTruckTestSituationMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.MuckTruckTestSituationMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateMuckTruckTestSituationByPk(ctx context.Context, inc *model.MuckTruckTestSituationIncInput, set *model.MuckTruckTestSituationSetInput, Id int64) (*model1.MuckTruckTestSituation, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.MuckTruckTestSituation{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.MuckTruckTestSituation
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) MuckTruckTestSituation(ctx context.Context, distinctOn []model.MuckTruckTestSituationSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckTestSituationOrderBy, where *model.MuckTruckTestSituationBoolExp) ([]*model1.MuckTruckTestSituation, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.MuckTruckTestSituation{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.MuckTruckTestSituation
	tx = tx.Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) MuckTruckTestSituationAggregate(ctx context.Context, distinctOn []model.MuckTruckTestSituationSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckTestSituationOrderBy, where *model.MuckTruckTestSituationBoolExp) (*model.MuckTruckTestSituationAggregate, error) {
	var rs model.MuckTruckTestSituationAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.MuckTruckTestSituation{})
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

func (r *queryResolver) MuckTruckTestSituationByPk(ctx context.Context, Id int64) (*model1.MuckTruckTestSituation, error) {
	var rs model1.MuckTruckTestSituation
	tx := db.DB.Model(&model1.MuckTruckTestSituation{}).First(&rs, Id)
	err := tx.Error
	return &rs, err
}
