package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/driving/graph/model"
	model1 "VehicleSupervision/internal/modules/driving/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteEcdFileVehicleTime(ctx context.Context, where model.EcdFileVehicleTimeBoolExp) (*model.EcdFileVehicleTimeMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EcdFileVehicleTime{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.EcdFileVehicleTime
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
	return &model.EcdFileVehicleTimeMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteEcdFileVehicleTimeByPk(ctx context.Context, Id int64) (*model1.EcdFileVehicleTime, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.EcdFileVehicleTime
	tx := db.DB.Model(&model1.EcdFileVehicleTime{})
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

func (r *mutationResolver) DeleteEcdFileVehicleTimeByUnionPk(ctx context.Context, unionId string) (*model1.EcdFileVehicleTime, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.EcdFileVehicleTime
	tx := db.DB.Model(&model1.EcdFileVehicleTime{})
	if len(preloads) > 0 {
		// 如果请求的字段不为空，则先查询一遍数据库
		tx = tx.Select(preloads).Where(rs.UnionPrimaryColumnName()+" = ?", unionId).First(&rs)
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

func (r *mutationResolver) InsertEcdFileVehicleTime(ctx context.Context, objects []*model.EcdFileVehicleTimeInsertInput) (*model.EcdFileVehicleTimeMutationResponse, error) {
	rs := make([]*model1.EcdFileVehicleTime, 0)
	for _, object := range objects {
		v := &model1.EcdFileVehicleTime{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.EcdFileVehicleTime{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.EcdFileVehicleTimeMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertEcdFileVehicleTimeOne(ctx context.Context, object model.EcdFileVehicleTimeInsertInput) (*model1.EcdFileVehicleTime, error) {
	rs := &model1.EcdFileVehicleTime{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.EcdFileVehicleTime{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateEcdFileVehicleTime(ctx context.Context, inc *model.EcdFileVehicleTimeIncInput, set *model.EcdFileVehicleTimeSetInput, where model.EcdFileVehicleTimeBoolExp) (*model.EcdFileVehicleTimeMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EcdFileVehicleTime{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.EcdFileVehicleTimeMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.EcdFileVehicleTime
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
	return &model.EcdFileVehicleTimeMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateEcdFileVehicleTimeByPk(ctx context.Context, inc *model.EcdFileVehicleTimeIncInput, set *model.EcdFileVehicleTimeSetInput, Id int64) (*model1.EcdFileVehicleTime, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.EcdFileVehicleTime{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.EcdFileVehicleTime
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *mutationResolver) UpdateEcdFileVehicleTimeByUnionPk(ctx context.Context, inc *model.EcdFileVehicleTimeIncInput, set *model.EcdFileVehicleTimeSetInput, unionId string) (*model1.EcdFileVehicleTime, error) {
	var rs model1.EcdFileVehicleTime
	tx := db.DB.Where(rs.UnionPrimaryColumnName()+" = ?", unionId)
	qt := util.NewQueryTranslator(tx, &model1.EcdFileVehicleTime{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}

	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) EcdFileVehicleTime(ctx context.Context, distinctOn []model.EcdFileVehicleTimeSelectColumn, limit *int, offset *int, orderBy []*model.EcdFileVehicleTimeOrderBy, where *model.EcdFileVehicleTimeBoolExp) ([]*model1.EcdFileVehicleTime, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EcdFileVehicleTime{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.EcdFileVehicleTime
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) EcdFileVehicleTimeAggregate(ctx context.Context, distinctOn []model.EcdFileVehicleTimeSelectColumn, limit *int, offset *int, orderBy []*model.EcdFileVehicleTimeOrderBy, where *model.EcdFileVehicleTimeBoolExp) (*model.EcdFileVehicleTimeAggregate, error) {
	var rs model.EcdFileVehicleTimeAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.EcdFileVehicleTime{})
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

func (r *queryResolver) EcdFileVehicleTimeByPk(ctx context.Context, Id int64) (*model1.EcdFileVehicleTime, error) {
	var rs model1.EcdFileVehicleTime
	tx := db.DB.Model(&model1.EcdFileVehicleTime{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}

func (r *queryResolver) EcdFileVehicleTimeByUnionPk(ctx context.Context, unionId string) (*model1.EcdFileVehicleTime, error) {
	var rs model1.EcdFileVehicleTime
	tx := db.DB.Model(&model1.EcdFileVehicleTime{}).Select(util.GetTopPreloads(ctx)).Where(rs.UnionPrimaryColumnName()+" = ?", unionId).First(&rs)

	err := tx.Error
	return &rs, err
}
