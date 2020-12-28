package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/muck_truck_recommend_catalog/graph/generated"
	"VehicleSupervision/internal/modules/muck_truck_recommend_catalog/graph/model"
	model1 "VehicleSupervision/internal/modules/muck_truck_recommend_catalog/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteNewMuckTruckRecommendCatalog(ctx context.Context, where model.NewMuckTruckRecommendCatalogBoolExp) (*model.NewMuckTruckRecommendCatalogMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.NewMuckTruckRecommendCatalog{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.NewMuckTruckRecommendCatalog
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
	return &model.NewMuckTruckRecommendCatalogMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteNewMuckTruckRecommendCatalogByPk(ctx context.Context, Id int64) (*model1.NewMuckTruckRecommendCatalog, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.NewMuckTruckRecommendCatalog
	tx := db.DB.Model(&model1.NewMuckTruckRecommendCatalog{})
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

func (r *mutationResolver) InsertNewMuckTruckRecommendCatalog(ctx context.Context, objects []*model.NewMuckTruckRecommendCatalogInsertInput) (*model.NewMuckTruckRecommendCatalogMutationResponse, error) {
	rs := []*model1.NewMuckTruckRecommendCatalog{}
	for _, object := range objects {
		v := &model1.NewMuckTruckRecommendCatalog{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.NewMuckTruckRecommendCatalog{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.NewMuckTruckRecommendCatalogMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertNewMuckTruckRecommendCatalogOne(ctx context.Context, object model.NewMuckTruckRecommendCatalogInsertInput) (*model1.NewMuckTruckRecommendCatalog, error) {
	rs := &model1.NewMuckTruckRecommendCatalog{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.NewMuckTruckRecommendCatalog{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateNewMuckTruckRecommendCatalog(ctx context.Context, inc *model.NewMuckTruckRecommendCatalogIncInput, set *model.NewMuckTruckRecommendCatalogSetInput, where model.NewMuckTruckRecommendCatalogBoolExp) (*model.NewMuckTruckRecommendCatalogMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.NewMuckTruckRecommendCatalog{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.NewMuckTruckRecommendCatalogMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.NewMuckTruckRecommendCatalogMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateNewMuckTruckRecommendCatalogByPk(ctx context.Context, inc *model.NewMuckTruckRecommendCatalogIncInput, set *model.NewMuckTruckRecommendCatalogSetInput, Id int64) (*model1.NewMuckTruckRecommendCatalog, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.NewMuckTruckRecommendCatalog{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.NewMuckTruckRecommendCatalog
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) NewMuckTruckRecommendCatalog(ctx context.Context, distinctOn []model.NewMuckTruckRecommendCatalogSelectColumn, limit *int, offset *int, orderBy []*model.NewMuckTruckRecommendCatalogOrderBy, where *model.NewMuckTruckRecommendCatalogBoolExp) ([]*model1.NewMuckTruckRecommendCatalog, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.NewMuckTruckRecommendCatalog{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.NewMuckTruckRecommendCatalog
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) NewMuckTruckRecommendCatalogAggregate(ctx context.Context, distinctOn []model.NewMuckTruckRecommendCatalogSelectColumn, limit *int, offset *int, orderBy []*model.NewMuckTruckRecommendCatalogOrderBy, where *model.NewMuckTruckRecommendCatalogBoolExp) (*model.NewMuckTruckRecommendCatalogAggregate, error) {
	var rs model.NewMuckTruckRecommendCatalogAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.NewMuckTruckRecommendCatalog{})
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

func (r *queryResolver) NewMuckTruckRecommendCatalogByPk(ctx context.Context, Id int64) (*model1.NewMuckTruckRecommendCatalog, error) {
	var rs model1.NewMuckTruckRecommendCatalog
	tx := db.DB.Model(&model1.NewMuckTruckRecommendCatalog{}).First(&rs, Id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
