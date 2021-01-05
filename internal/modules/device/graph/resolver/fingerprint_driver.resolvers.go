package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/device/graph/generated"
	"VehicleSupervision/internal/modules/device/graph/model"
	model1 "VehicleSupervision/internal/modules/device/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteFingerprintDriver(ctx context.Context, where model.FingerprintDriverBoolExp) (*model.FingerprintDriverMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.FingerprintDriver{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.FingerprintDriver
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
	return &model.FingerprintDriverMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteFingerprintDriverByPk(ctx context.Context, Id int64) (*model1.FingerprintDriver, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.FingerprintDriver
	tx := db.DB.Model(&model1.FingerprintDriver{})
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

func (r *mutationResolver) DeleteFingerprintDriverByUnionPk(ctx context.Context, unionId string) (*model1.FingerprintDriver, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.FingerprintDriver
	tx := db.DB.Model(&model1.FingerprintDriver{})
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

func (r *mutationResolver) InsertFingerprintDriver(ctx context.Context, objects []*model.FingerprintDriverInsertInput) (*model.FingerprintDriverMutationResponse, error) {
	rs := make([]*model1.FingerprintDriver, 0)
	for _, object := range objects {
		v := &model1.FingerprintDriver{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.FingerprintDriver{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.FingerprintDriverMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertFingerprintDriverOne(ctx context.Context, object model.FingerprintDriverInsertInput) (*model1.FingerprintDriver, error) {
	rs := &model1.FingerprintDriver{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.FingerprintDriver{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateFingerprintDriver(ctx context.Context, inc *model.FingerprintDriverIncInput, set *model.FingerprintDriverSetInput, where model.FingerprintDriverBoolExp) (*model.FingerprintDriverMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.FingerprintDriver{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.FingerprintDriverMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.FingerprintDriver
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
	return &model.FingerprintDriverMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateFingerprintDriverByPk(ctx context.Context, inc *model.FingerprintDriverIncInput, set *model.FingerprintDriverSetInput, Id int64) (*model1.FingerprintDriver, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.FingerprintDriver{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.FingerprintDriver
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *mutationResolver) UpdateFingerprintDriverByUnionPk(ctx context.Context, inc *model.FingerprintDriverIncInput, set *model.FingerprintDriverSetInput, unionId string) (*model1.FingerprintDriver, error) {
	var rs model1.FingerprintDriver
	tx := db.DB.Where(rs.UnionPrimaryColumnName()+" = ?", unionId)
	qt := util.NewQueryTranslator(tx, &model1.FingerprintDriver{})
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

func (r *queryResolver) FingerprintDriver(ctx context.Context, distinctOn []model.FingerprintDriverSelectColumn, limit *int, offset *int, orderBy []*model.FingerprintDriverOrderBy, where *model.FingerprintDriverBoolExp) ([]*model1.FingerprintDriver, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.FingerprintDriver{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.FingerprintDriver
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) FingerprintDriverAggregate(ctx context.Context, distinctOn []model.FingerprintDriverSelectColumn, limit *int, offset *int, orderBy []*model.FingerprintDriverOrderBy, where *model.FingerprintDriverBoolExp) (*model.FingerprintDriverAggregate, error) {
	var rs model.FingerprintDriverAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.FingerprintDriver{})
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

func (r *queryResolver) FingerprintDriverByPk(ctx context.Context, Id int64) (*model1.FingerprintDriver, error) {
	var rs model1.FingerprintDriver
	tx := db.DB.Model(&model1.FingerprintDriver{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}

func (r *queryResolver) FingerprintDriverByUnionPk(ctx context.Context, unionId string) (*model1.FingerprintDriver, error) {
	var rs model1.FingerprintDriver
	tx := db.DB.Model(&model1.FingerprintDriver{}).Select(util.GetTopPreloads(ctx)).Where(rs.UnionPrimaryColumnName()+" = ?", unionId).First(&rs)

	err := tx.Error
	return &rs, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
