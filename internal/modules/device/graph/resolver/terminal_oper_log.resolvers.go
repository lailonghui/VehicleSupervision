package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/device/graph/model"
	model1 "VehicleSupervision/internal/modules/device/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteTerminalOperLog(ctx context.Context, where model.TerminalOperLogBoolExp) (*model.TerminalOperLogMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalOperLog{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.TerminalOperLog
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
	return &model.TerminalOperLogMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteTerminalOperLogByPk(ctx context.Context, Id int64) (*model1.TerminalOperLog, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.TerminalOperLog
	tx := db.DB.Model(&model1.TerminalOperLog{})
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

func (r *mutationResolver) InsertTerminalOperLog(ctx context.Context, objects []*model.TerminalOperLogInsertInput) (*model.TerminalOperLogMutationResponse, error) {
	rs := make([]*model1.TerminalOperLog, 0)
	for _, object := range objects {
		v := &model1.TerminalOperLog{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.TerminalOperLog{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.TerminalOperLogMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertTerminalOperLogOne(ctx context.Context, object model.TerminalOperLogInsertInput) (*model1.TerminalOperLog, error) {
	rs := &model1.TerminalOperLog{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.TerminalOperLog{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateTerminalOperLog(ctx context.Context, inc *model.TerminalOperLogIncInput, set *model.TerminalOperLogSetInput, where model.TerminalOperLogBoolExp) (*model.TerminalOperLogMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalOperLog{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.TerminalOperLogMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.TerminalOperLog
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
	return &model.TerminalOperLogMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateTerminalOperLogByPk(ctx context.Context, inc *model.TerminalOperLogIncInput, set *model.TerminalOperLogSetInput, Id int64) (*model1.TerminalOperLog, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.TerminalOperLog{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.TerminalOperLog
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) TerminalOperLog(ctx context.Context, distinctOn []model.TerminalOperLogSelectColumn, limit *int, offset *int, orderBy []*model.TerminalOperLogOrderBy, where *model.TerminalOperLogBoolExp) ([]*model1.TerminalOperLog, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalOperLog{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.TerminalOperLog
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) TerminalOperLogAggregate(ctx context.Context, distinctOn []model.TerminalOperLogSelectColumn, limit *int, offset *int, orderBy []*model.TerminalOperLogOrderBy, where *model.TerminalOperLogBoolExp) (*model.TerminalOperLogAggregate, error) {
	var rs model.TerminalOperLogAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.TerminalOperLog{})
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

func (r *queryResolver) TerminalOperLogByPk(ctx context.Context, Id int64) (*model1.TerminalOperLog, error) {
	var rs model1.TerminalOperLog
	tx := db.DB.Model(&model1.TerminalOperLog{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}
