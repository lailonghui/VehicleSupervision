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

func (r *mutationResolver) DeleteTempIdentityCardDownloadLog(ctx context.Context, where model.TempIdentityCardDownloadLogBoolExp) (*model.TempIdentityCardDownloadLogMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TempIdentityCardDownloadLog{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.TempIdentityCardDownloadLog
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
	return &model.TempIdentityCardDownloadLogMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteTempIdentityCardDownloadLogByPk(ctx context.Context, Id int64) (*model1.TempIdentityCardDownloadLog, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.TempIdentityCardDownloadLog
	tx := db.DB.Model(&model1.TempIdentityCardDownloadLog{})
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

func (r *mutationResolver) InsertTempIdentityCardDownloadLog(ctx context.Context, objects []*model.TempIdentityCardDownloadLogInsertInput) (*model.TempIdentityCardDownloadLogMutationResponse, error) {
	rs := make([]*model1.TempIdentityCardDownloadLog, 0)
	for _, object := range objects {
		v := &model1.TempIdentityCardDownloadLog{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.TempIdentityCardDownloadLog{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.TempIdentityCardDownloadLogMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertTempIdentityCardDownloadLogOne(ctx context.Context, object model.TempIdentityCardDownloadLogInsertInput) (*model1.TempIdentityCardDownloadLog, error) {
	rs := &model1.TempIdentityCardDownloadLog{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.TempIdentityCardDownloadLog{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateTempIdentityCardDownloadLog(ctx context.Context, inc *model.TempIdentityCardDownloadLogIncInput, set *model.TempIdentityCardDownloadLogSetInput, where model.TempIdentityCardDownloadLogBoolExp) (*model.TempIdentityCardDownloadLogMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TempIdentityCardDownloadLog{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.TempIdentityCardDownloadLogMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.TempIdentityCardDownloadLogMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateTempIdentityCardDownloadLogByPk(ctx context.Context, inc *model.TempIdentityCardDownloadLogIncInput, set *model.TempIdentityCardDownloadLogSetInput, Id int64) (*model1.TempIdentityCardDownloadLog, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.TempIdentityCardDownloadLog{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.TempIdentityCardDownloadLog
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) TempIdentityCardDownloadLog(ctx context.Context, distinctOn []model.TempIdentityCardDownloadLogSelectColumn, limit *int, offset *int, orderBy []*model.TempIdentityCardDownloadLogOrderBy, where *model.TempIdentityCardDownloadLogBoolExp) ([]*model1.TempIdentityCardDownloadLog, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TempIdentityCardDownloadLog{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.TempIdentityCardDownloadLog
	tx = tx.Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) TempIdentityCardDownloadLogAggregate(ctx context.Context, distinctOn []model.TempIdentityCardDownloadLogSelectColumn, limit *int, offset *int, orderBy []*model.TempIdentityCardDownloadLogOrderBy, where *model.TempIdentityCardDownloadLogBoolExp) (*model.TempIdentityCardDownloadLogAggregate, error) {
	var rs model.TempIdentityCardDownloadLogAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.TempIdentityCardDownloadLog{})
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

func (r *queryResolver) TempIdentityCardDownloadLogByPk(ctx context.Context, Id int64) (*model1.TempIdentityCardDownloadLog, error) {
	var rs model1.TempIdentityCardDownloadLog
	tx := db.DB.Model(&model1.TempIdentityCardDownloadLog{}).First(&rs, Id)
	err := tx.Error
	return &rs, err
}
