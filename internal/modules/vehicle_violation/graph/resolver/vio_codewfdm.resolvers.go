package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/vehicle_violation/graph/model"
	model1 "VehicleSupervision/internal/modules/vehicle_violation/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteVioCodewfdm(ctx context.Context, where model.VioCodewfdmBoolExp) (*model.VioCodewfdmMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VioCodewfdm{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.VioCodewfdm
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
	return &model.VioCodewfdmMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteVioCodewfdmByPk(ctx context.Context, Wfxw string) (*model1.VioCodewfdm, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.VioCodewfdm
	tx := db.DB.Model(&model1.VioCodewfdm{})
	if len(preloads) > 0 {
		// 如果请求的字段不为空，则先查询一遍数据库
		tx = tx.Select(preloads).Where("wfxw = ?", Wfxw).First(&rs)
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

func (r *mutationResolver) InsertVioCodewfdm(ctx context.Context, objects []*model.VioCodewfdmInsertInput) (*model.VioCodewfdmMutationResponse, error) {
	rs := []*model1.VioCodewfdm{}
	for _, object := range objects {
		v := &model1.VioCodewfdm{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.VioCodewfdm{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.VioCodewfdmMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertVioCodewfdmOne(ctx context.Context, object model.VioCodewfdmInsertInput) (*model1.VioCodewfdm, error) {
	rs := &model1.VioCodewfdm{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.VioCodewfdm{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateVioCodewfdm(ctx context.Context, inc *model.VioCodewfdmIncInput, set *model.VioCodewfdmSetInput, where model.VioCodewfdmBoolExp) (*model.VioCodewfdmMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VioCodewfdm{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.VioCodewfdmMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.VioCodewfdmMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateVioCodewfdmByPk(ctx context.Context, inc *model.VioCodewfdmIncInput, set *model.VioCodewfdmSetInput, Wfxw string) (*model1.VioCodewfdm, error) {
	tx := db.DB.Where("wfxw = ?", Wfxw)
	qt := util.NewQueryTranslator(tx, &model1.VioCodewfdm{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.VioCodewfdm
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) VioCodewfdm(ctx context.Context, distinctOn []model.VioCodewfdmSelectColumn, limit *int, offset *int, orderBy []*model.VioCodewfdmOrderBy, where *model.VioCodewfdmBoolExp) ([]*model1.VioCodewfdm, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VioCodewfdm{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.VioCodewfdm
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) VioCodewfdmAggregate(ctx context.Context, distinctOn []model.VioCodewfdmSelectColumn, limit *int, offset *int, orderBy []*model.VioCodewfdmOrderBy, where *model.VioCodewfdmBoolExp) (*model.VioCodewfdmAggregate, error) {
	var rs model.VioCodewfdmAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.VioCodewfdm{})
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

func (r *queryResolver) VioCodewfdmByPk(ctx context.Context, Wfxw string) (*model1.VioCodewfdm, error) {
	var rs model1.VioCodewfdm
	tx := db.DB.Model(&model1.VioCodewfdm{}).First(&rs, Wfxw)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
