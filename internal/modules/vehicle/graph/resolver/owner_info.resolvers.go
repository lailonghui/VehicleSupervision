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

func (r *mutationResolver) DeleteOwnerInfo(ctx context.Context, where model.OwnerInfoBoolExp) (*model.OwnerInfoMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.OwnerInfo{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.OwnerInfo
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
	return &model.OwnerInfoMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteOwnerInfoByPk(ctx context.Context, Id int64) (*model1.OwnerInfo, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.OwnerInfo
	tx := db.DB.Model(&model1.OwnerInfo{})
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

func (r *mutationResolver) InsertOwnerInfo(ctx context.Context, objects []*model.OwnerInfoInsertInput) (*model.OwnerInfoMutationResponse, error) {
	rs := make([]*model1.OwnerInfo, 0)
	for _, object := range objects {
		v := &model1.OwnerInfo{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.OwnerInfo{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.OwnerInfoMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertOwnerInfoOne(ctx context.Context, object model.OwnerInfoInsertInput) (*model1.OwnerInfo, error) {
	rs := &model1.OwnerInfo{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.OwnerInfo{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateOwnerInfo(ctx context.Context, inc *model.OwnerInfoIncInput, set *model.OwnerInfoSetInput, where model.OwnerInfoBoolExp) (*model.OwnerInfoMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.OwnerInfo{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.OwnerInfoMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.OwnerInfoMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateOwnerInfoByPk(ctx context.Context, inc *model.OwnerInfoIncInput, set *model.OwnerInfoSetInput, Id int64) (*model1.OwnerInfo, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.OwnerInfo{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.OwnerInfo
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) OwnerInfo(ctx context.Context, distinctOn []model.OwnerInfoSelectColumn, limit *int, offset *int, orderBy []*model.OwnerInfoOrderBy, where *model.OwnerInfoBoolExp) ([]*model1.OwnerInfo, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.OwnerInfo{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.OwnerInfo
	tx = tx.Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) OwnerInfoAggregate(ctx context.Context, distinctOn []model.OwnerInfoSelectColumn, limit *int, offset *int, orderBy []*model.OwnerInfoOrderBy, where *model.OwnerInfoBoolExp) (*model.OwnerInfoAggregate, error) {
	var rs model.OwnerInfoAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.OwnerInfo{})
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

func (r *queryResolver) OwnerInfoByPk(ctx context.Context, Id int64) (*model1.OwnerInfo, error) {
	var rs model1.OwnerInfo
	tx := db.DB.Model(&model1.OwnerInfo{}).First(&rs, Id)
	err := tx.Error
	return &rs, err
}
