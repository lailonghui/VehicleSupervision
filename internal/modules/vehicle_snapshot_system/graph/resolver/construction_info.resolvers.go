package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/vehicle_snapshot_system/graph/generated"
	"VehicleSupervision/internal/modules/vehicle_snapshot_system/graph/model"
	model1 "VehicleSupervision/internal/modules/vehicle_snapshot_system/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteConstructionInfo(ctx context.Context, where model.ConstructionInfoBoolExp) (*model.ConstructionInfoMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.ConstructionInfo{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.ConstructionInfo
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
	return &model.ConstructionInfoMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteConstructionInfoByPk(ctx context.Context, Id int64) (*model1.ConstructionInfo, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.ConstructionInfo
	tx := db.DB.Model(&model1.ConstructionInfo{})
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

func (r *mutationResolver) InsertConstructionInfo(ctx context.Context, objects []*model.ConstructionInfoInsertInput) (*model.ConstructionInfoMutationResponse, error) {
	rs := make([]*model1.ConstructionInfo, 0)
	for _, object := range objects {
		v := &model1.ConstructionInfo{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.ConstructionInfo{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.ConstructionInfoMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertConstructionInfoOne(ctx context.Context, object model.ConstructionInfoInsertInput) (*model1.ConstructionInfo, error) {
	rs := &model1.ConstructionInfo{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.ConstructionInfo{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateConstructionInfo(ctx context.Context, inc *model.ConstructionInfoIncInput, set *model.ConstructionInfoSetInput, where model.ConstructionInfoBoolExp) (*model.ConstructionInfoMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.ConstructionInfo{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.ConstructionInfoMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.ConstructionInfoMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateConstructionInfoByPk(ctx context.Context, inc *model.ConstructionInfoIncInput, set *model.ConstructionInfoSetInput, Id int64) (*model1.ConstructionInfo, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.ConstructionInfo{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.ConstructionInfo
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) ConstructionInfo(ctx context.Context, distinctOn []model.ConstructionInfoSelectColumn, limit *int, offset *int, orderBy []*model.ConstructionInfoOrderBy, where *model.ConstructionInfoBoolExp) ([]*model1.ConstructionInfo, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.ConstructionInfo{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.ConstructionInfo
	tx = tx.Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) ConstructionInfoAggregate(ctx context.Context, distinctOn []model.ConstructionInfoSelectColumn, limit *int, offset *int, orderBy []*model.ConstructionInfoOrderBy, where *model.ConstructionInfoBoolExp) (*model.ConstructionInfoAggregate, error) {
	var rs model.ConstructionInfoAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.ConstructionInfo{})
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

func (r *queryResolver) ConstructionInfoByPk(ctx context.Context, Id int64) (*model1.ConstructionInfo, error) {
	var rs model1.ConstructionInfo
	tx := db.DB.Model(&model1.ConstructionInfo{}).First(&rs, Id)
	err := tx.Error
	return &rs, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }