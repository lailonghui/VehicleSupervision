package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/dynamic_supervision/graph/generated"
	"VehicleSupervision/internal/modules/dynamic_supervision/graph/model"
	model1 "VehicleSupervision/internal/modules/dynamic_supervision/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteDynamicSupervision(ctx context.Context, where model.DynamicSupervisionBoolExp) (*model.DynamicSupervisionMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DynamicSupervision{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.DynamicSupervision
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
	return &model.DynamicSupervisionMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteDynamicSupervisionByPk(ctx context.Context, Id int64) (*model1.DynamicSupervision, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.DynamicSupervision
	tx := db.DB.Model(&model1.DynamicSupervision{})
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

func (r *mutationResolver) InsertDynamicSupervision(ctx context.Context, objects []*model.DynamicSupervisionInsertInput) (*model.DynamicSupervisionMutationResponse, error) {
	rs := make([]*model1.DynamicSupervision, 0)
	for _, object := range objects {
		v := &model1.DynamicSupervision{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.DynamicSupervision{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.DynamicSupervisionMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertDynamicSupervisionOne(ctx context.Context, object model.DynamicSupervisionInsertInput) (*model1.DynamicSupervision, error) {
	rs := &model1.DynamicSupervision{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.DynamicSupervision{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateDynamicSupervision(ctx context.Context, inc *model.DynamicSupervisionIncInput, set *model.DynamicSupervisionSetInput, where model.DynamicSupervisionBoolExp) (*model.DynamicSupervisionMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DynamicSupervision{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.DynamicSupervisionMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.DynamicSupervisionMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateDynamicSupervisionByPk(ctx context.Context, inc *model.DynamicSupervisionIncInput, set *model.DynamicSupervisionSetInput, Id int64) (*model1.DynamicSupervision, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.DynamicSupervision{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.DynamicSupervision
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) DynamicSupervision(ctx context.Context, distinctOn []model.DynamicSupervisionSelectColumn, limit *int, offset *int, orderBy []*model.DynamicSupervisionOrderBy, where *model.DynamicSupervisionBoolExp) ([]*model1.DynamicSupervision, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DynamicSupervision{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.DynamicSupervision
	tx = tx.Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) DynamicSupervisionAggregate(ctx context.Context, distinctOn []model.DynamicSupervisionSelectColumn, limit *int, offset *int, orderBy []*model.DynamicSupervisionOrderBy, where *model.DynamicSupervisionBoolExp) (*model.DynamicSupervisionAggregate, error) {
	var rs model.DynamicSupervisionAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.DynamicSupervision{})
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

func (r *queryResolver) DynamicSupervisionByPk(ctx context.Context, Id int64) (*model1.DynamicSupervision, error) {
	var rs model1.DynamicSupervision
	tx := db.DB.Model(&model1.DynamicSupervision{}).First(&rs, Id)
	err := tx.Error
	return &rs, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
