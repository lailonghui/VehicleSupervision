package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/dynamic_supervision/graph/model"
	model1 "VehicleSupervision/internal/modules/dynamic_supervision/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteDynamicSpotCheckDisposal(ctx context.Context, where model.DynamicSpotCheckDisposalBoolExp) (*model.DynamicSpotCheckDisposalMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DynamicSpotCheckDisposal{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.DynamicSpotCheckDisposal
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
	return &model.DynamicSpotCheckDisposalMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteDynamicSpotCheckDisposalByPk(ctx context.Context, Id int64) (*model1.DynamicSpotCheckDisposal, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.DynamicSpotCheckDisposal
	tx := db.DB.Model(&model1.DynamicSpotCheckDisposal{})
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

func (r *mutationResolver) InsertDynamicSpotCheckDisposal(ctx context.Context, objects []*model.DynamicSpotCheckDisposalInsertInput) (*model.DynamicSpotCheckDisposalMutationResponse, error) {
	rs := []*model1.DynamicSpotCheckDisposal{}
	for _, object := range objects {
		v := &model1.DynamicSpotCheckDisposal{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.DynamicSpotCheckDisposal{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.DynamicSpotCheckDisposalMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertDynamicSpotCheckDisposalOne(ctx context.Context, object model.DynamicSpotCheckDisposalInsertInput) (*model1.DynamicSpotCheckDisposal, error) {
	rs := &model1.DynamicSpotCheckDisposal{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.DynamicSpotCheckDisposal{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateDynamicSpotCheckDisposal(ctx context.Context, inc *model.DynamicSpotCheckDisposalIncInput, set *model.DynamicSpotCheckDisposalSetInput, where model.DynamicSpotCheckDisposalBoolExp) (*model.DynamicSpotCheckDisposalMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DynamicSpotCheckDisposal{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.DynamicSpotCheckDisposalMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.DynamicSpotCheckDisposalMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateDynamicSpotCheckDisposalByPk(ctx context.Context, inc *model.DynamicSpotCheckDisposalIncInput, set *model.DynamicSpotCheckDisposalSetInput, Id int64) (*model1.DynamicSpotCheckDisposal, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.DynamicSpotCheckDisposal{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.DynamicSpotCheckDisposal
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) DynamicSpotCheckDisposal(ctx context.Context, distinctOn []model.DynamicSpotCheckDisposalSelectColumn, limit *int, offset *int, orderBy []*model.DynamicSpotCheckDisposalOrderBy, where *model.DynamicSpotCheckDisposalBoolExp) ([]*model1.DynamicSpotCheckDisposal, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DynamicSpotCheckDisposal{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.DynamicSpotCheckDisposal
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) DynamicSpotCheckDisposalAggregate(ctx context.Context, distinctOn []model.DynamicSpotCheckDisposalSelectColumn, limit *int, offset *int, orderBy []*model.DynamicSpotCheckDisposalOrderBy, where *model.DynamicSpotCheckDisposalBoolExp) (*model.DynamicSpotCheckDisposalAggregate, error) {
	var rs model.DynamicSpotCheckDisposalAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.DynamicSpotCheckDisposal{})
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

func (r *queryResolver) DynamicSpotCheckDisposalByPk(ctx context.Context, Id int64) (*model1.DynamicSpotCheckDisposal, error) {
	var rs model1.DynamicSpotCheckDisposal
	tx := db.DB.Model(&model1.DynamicSpotCheckDisposal{}).First(&rs, Id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
