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

func (r *mutationResolver) DeleteVehicleOfflineDisposal(ctx context.Context, where model.VehicleOfflineDisposalBoolExp) (*model.VehicleOfflineDisposalMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleOfflineDisposal{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.VehicleOfflineDisposal
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
	return &model.VehicleOfflineDisposalMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteVehicleOfflineDisposalByPk(ctx context.Context, Id int64) (*model1.VehicleOfflineDisposal, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.VehicleOfflineDisposal
	tx := db.DB.Model(&model1.VehicleOfflineDisposal{})
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

func (r *mutationResolver) InsertVehicleOfflineDisposal(ctx context.Context, objects []*model.VehicleOfflineDisposalInsertInput) (*model.VehicleOfflineDisposalMutationResponse, error) {
	rs := []*model1.VehicleOfflineDisposal{}
	for _, object := range objects {
		v := &model1.VehicleOfflineDisposal{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.VehicleOfflineDisposal{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.VehicleOfflineDisposalMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertVehicleOfflineDisposalOne(ctx context.Context, object model.VehicleOfflineDisposalInsertInput) (*model1.VehicleOfflineDisposal, error) {
	rs := &model1.VehicleOfflineDisposal{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.VehicleOfflineDisposal{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateVehicleOfflineDisposal(ctx context.Context, inc *model.VehicleOfflineDisposalIncInput, set *model.VehicleOfflineDisposalSetInput, where model.VehicleOfflineDisposalBoolExp) (*model.VehicleOfflineDisposalMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleOfflineDisposal{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.VehicleOfflineDisposalMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.VehicleOfflineDisposalMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateVehicleOfflineDisposalByPk(ctx context.Context, inc *model.VehicleOfflineDisposalIncInput, set *model.VehicleOfflineDisposalSetInput, Id int64) (*model1.VehicleOfflineDisposal, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.VehicleOfflineDisposal{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.VehicleOfflineDisposal
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) VehicleOfflineDisposal(ctx context.Context, distinctOn []model.VehicleOfflineDisposalSelectColumn, limit *int, offset *int, orderBy []*model.VehicleOfflineDisposalOrderBy, where *model.VehicleOfflineDisposalBoolExp) ([]*model1.VehicleOfflineDisposal, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleOfflineDisposal{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.VehicleOfflineDisposal
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) VehicleOfflineDisposalAggregate(ctx context.Context, distinctOn []model.VehicleOfflineDisposalSelectColumn, limit *int, offset *int, orderBy []*model.VehicleOfflineDisposalOrderBy, where *model.VehicleOfflineDisposalBoolExp) (*model.VehicleOfflineDisposalAggregate, error) {
	var rs model.VehicleOfflineDisposalAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.VehicleOfflineDisposal{})
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

func (r *queryResolver) VehicleOfflineDisposalByPk(ctx context.Context, Id int64) (*model1.VehicleOfflineDisposal, error) {
	var rs model1.VehicleOfflineDisposal
	tx := db.DB.Model(&model1.VehicleOfflineDisposal{}).First(&rs, Id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
