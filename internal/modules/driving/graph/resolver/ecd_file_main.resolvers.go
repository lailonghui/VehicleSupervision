package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/driving/graph/generated"
	"VehicleSupervision/internal/modules/driving/graph/model"
	model1 "VehicleSupervision/internal/modules/driving/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteEcdFileMain(ctx context.Context, where model.EcdFileMainBoolExp) (*model.EcdFileMainMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EcdFileMain{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.EcdFileMain
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
	return &model.EcdFileMainMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteEcdFileMainByPk(ctx context.Context, Id int64) (*model1.EcdFileMain, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.EcdFileMain
	tx := db.DB.Model(&model1.EcdFileMain{})
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

func (r *mutationResolver) InsertEcdFileMain(ctx context.Context, objects []*model.EcdFileMainInsertInput) (*model.EcdFileMainMutationResponse, error) {
	rs := make([]*model1.EcdFileMain, 0)
	for _, object := range objects {
		v := &model1.EcdFileMain{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.EcdFileMain{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.EcdFileMainMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertEcdFileMainOne(ctx context.Context, object model.EcdFileMainInsertInput) (*model1.EcdFileMain, error) {
	rs := &model1.EcdFileMain{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.EcdFileMain{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateEcdFileMain(ctx context.Context, inc *model.EcdFileMainIncInput, set *model.EcdFileMainSetInput, where model.EcdFileMainBoolExp) (*model.EcdFileMainMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EcdFileMain{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.EcdFileMainMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.EcdFileMain
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
	return &model.EcdFileMainMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateEcdFileMainByPk(ctx context.Context, inc *model.EcdFileMainIncInput, set *model.EcdFileMainSetInput, Id int64) (*model1.EcdFileMain, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.EcdFileMain{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.EcdFileMain
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) EcdFileMain(ctx context.Context, distinctOn []model.EcdFileMainSelectColumn, limit *int, offset *int, orderBy []*model.EcdFileMainOrderBy, where *model.EcdFileMainBoolExp) ([]*model1.EcdFileMain, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EcdFileMain{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.EcdFileMain
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) EcdFileMainAggregate(ctx context.Context, distinctOn []model.EcdFileMainSelectColumn, limit *int, offset *int, orderBy []*model.EcdFileMainOrderBy, where *model.EcdFileMainBoolExp) (*model.EcdFileMainAggregate, error) {
	var rs model.EcdFileMainAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.EcdFileMain{})
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

func (r *queryResolver) EcdFileMainByPk(ctx context.Context, Id int64) (*model1.EcdFileMain, error) {
	var rs model1.EcdFileMain
	tx := db.DB.Model(&model1.EcdFileMain{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
