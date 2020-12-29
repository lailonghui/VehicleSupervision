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

func (r *mutationResolver) DeleteTerminalRegLog(ctx context.Context, where model.TerminalRegLogBoolExp) (*model.TerminalRegLogMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalRegLog{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.TerminalRegLog
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
	return &model.TerminalRegLogMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteTerminalRegLogByPk(ctx context.Context, Id int64) (*model1.TerminalRegLog, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.TerminalRegLog
	tx := db.DB.Model(&model1.TerminalRegLog{})
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

func (r *mutationResolver) InsertTerminalRegLog(ctx context.Context, objects []*model.TerminalRegLogInsertInput) (*model.TerminalRegLogMutationResponse, error) {
	rs := make([]*model1.TerminalRegLog, 0)
	for _, object := range objects {
		v := &model1.TerminalRegLog{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.TerminalRegLog{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.TerminalRegLogMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertTerminalRegLogOne(ctx context.Context, object model.TerminalRegLogInsertInput) (*model1.TerminalRegLog, error) {
	rs := &model1.TerminalRegLog{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.TerminalRegLog{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateTerminalRegLog(ctx context.Context, inc *model.TerminalRegLogIncInput, set *model.TerminalRegLogSetInput, where model.TerminalRegLogBoolExp) (*model.TerminalRegLogMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalRegLog{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.TerminalRegLogMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.TerminalRegLog
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
	return &model.TerminalRegLogMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateTerminalRegLogByPk(ctx context.Context, inc *model.TerminalRegLogIncInput, set *model.TerminalRegLogSetInput, Id int64) (*model1.TerminalRegLog, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.TerminalRegLog{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.TerminalRegLog
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) TerminalRegLog(ctx context.Context, distinctOn []model.TerminalRegLogSelectColumn, limit *int, offset *int, orderBy []*model.TerminalRegLogOrderBy, where *model.TerminalRegLogBoolExp) ([]*model1.TerminalRegLog, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalRegLog{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.TerminalRegLog
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) TerminalRegLogAggregate(ctx context.Context, distinctOn []model.TerminalRegLogSelectColumn, limit *int, offset *int, orderBy []*model.TerminalRegLogOrderBy, where *model.TerminalRegLogBoolExp) (*model.TerminalRegLogAggregate, error) {
	var rs model.TerminalRegLogAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.TerminalRegLog{})
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

func (r *queryResolver) TerminalRegLogByPk(ctx context.Context, Id int64) (*model1.TerminalRegLog, error) {
	var rs model1.TerminalRegLog
	tx := db.DB.Model(&model1.TerminalRegLog{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}
