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

func (r *mutationResolver) DeleteEnterpriseScoreLog(ctx context.Context, where model.EnterpriseScoreLogBoolExp) (*model.EnterpriseScoreLogMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EnterpriseScoreLog{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.EnterpriseScoreLog
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
	return &model.EnterpriseScoreLogMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteEnterpriseScoreLogByPk(ctx context.Context, Id int64) (*model1.EnterpriseScoreLog, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.EnterpriseScoreLog
	tx := db.DB.Model(&model1.EnterpriseScoreLog{})
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

func (r *mutationResolver) InsertEnterpriseScoreLog(ctx context.Context, objects []*model.EnterpriseScoreLogInsertInput) (*model.EnterpriseScoreLogMutationResponse, error) {
	rs := make([]*model1.EnterpriseScoreLog, 0)
	for _, object := range objects {
		v := &model1.EnterpriseScoreLog{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.EnterpriseScoreLog{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.EnterpriseScoreLogMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertEnterpriseScoreLogOne(ctx context.Context, object model.EnterpriseScoreLogInsertInput) (*model1.EnterpriseScoreLog, error) {
	rs := &model1.EnterpriseScoreLog{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.EnterpriseScoreLog{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateEnterpriseScoreLog(ctx context.Context, inc *model.EnterpriseScoreLogIncInput, set *model.EnterpriseScoreLogSetInput, where model.EnterpriseScoreLogBoolExp) (*model.EnterpriseScoreLogMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EnterpriseScoreLog{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.EnterpriseScoreLogMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.EnterpriseScoreLog
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
	return &model.EnterpriseScoreLogMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateEnterpriseScoreLogByPk(ctx context.Context, inc *model.EnterpriseScoreLogIncInput, set *model.EnterpriseScoreLogSetInput, Id int64) (*model1.EnterpriseScoreLog, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.EnterpriseScoreLog{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.EnterpriseScoreLog
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) EnterpriseScoreLog(ctx context.Context, distinctOn []model.EnterpriseScoreLogSelectColumn, limit *int, offset *int, orderBy []*model.EnterpriseScoreLogOrderBy, where *model.EnterpriseScoreLogBoolExp) ([]*model1.EnterpriseScoreLog, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EnterpriseScoreLog{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.EnterpriseScoreLog
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) EnterpriseScoreLogAggregate(ctx context.Context, distinctOn []model.EnterpriseScoreLogSelectColumn, limit *int, offset *int, orderBy []*model.EnterpriseScoreLogOrderBy, where *model.EnterpriseScoreLogBoolExp) (*model.EnterpriseScoreLogAggregate, error) {
	var rs model.EnterpriseScoreLogAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.EnterpriseScoreLog{})
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

func (r *queryResolver) EnterpriseScoreLogByPk(ctx context.Context, Id int64) (*model1.EnterpriseScoreLog, error) {
	var rs model1.EnterpriseScoreLog
	tx := db.DB.Model(&model1.EnterpriseScoreLog{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}
