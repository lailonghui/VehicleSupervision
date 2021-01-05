package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/ridehailing/graph/model"
	model1 "VehicleSupervision/internal/modules/ridehailing/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteRideHailingDriverVerify(ctx context.Context, where model.RideHailingDriverVerifyBoolExp) (*model.RideHailingDriverVerifyMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.RideHailingDriverVerify{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.RideHailingDriverVerify
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
	return &model.RideHailingDriverVerifyMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteRideHailingDriverVerifyByPk(ctx context.Context, Id int64) (*model1.RideHailingDriverVerify, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.RideHailingDriverVerify
	tx := db.DB.Model(&model1.RideHailingDriverVerify{})
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

func (r *mutationResolver) DeleteRideHailingDriverVerifyByUnionPk(ctx context.Context, unionId string) (*model1.RideHailingDriverVerify, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.RideHailingDriverVerify
	tx := db.DB.Model(&model1.RideHailingDriverVerify{})
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

func (r *mutationResolver) InsertRideHailingDriverVerify(ctx context.Context, objects []*model.RideHailingDriverVerifyInsertInput) (*model.RideHailingDriverVerifyMutationResponse, error) {
	rs := make([]*model1.RideHailingDriverVerify, 0)
	for _, object := range objects {
		v := &model1.RideHailingDriverVerify{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.RideHailingDriverVerify{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.RideHailingDriverVerifyMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertRideHailingDriverVerifyOne(ctx context.Context, object model.RideHailingDriverVerifyInsertInput) (*model1.RideHailingDriverVerify, error) {
	rs := &model1.RideHailingDriverVerify{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.RideHailingDriverVerify{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateRideHailingDriverVerify(ctx context.Context, inc *model.RideHailingDriverVerifyIncInput, set *model.RideHailingDriverVerifySetInput, where model.RideHailingDriverVerifyBoolExp) (*model.RideHailingDriverVerifyMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.RideHailingDriverVerify{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.RideHailingDriverVerifyMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.RideHailingDriverVerify
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
	return &model.RideHailingDriverVerifyMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateRideHailingDriverVerifyByPk(ctx context.Context, inc *model.RideHailingDriverVerifyIncInput, set *model.RideHailingDriverVerifySetInput, Id int64) (*model1.RideHailingDriverVerify, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.RideHailingDriverVerify{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.RideHailingDriverVerify
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *mutationResolver) UpdateRideHailingDriverVerifyByUnionPk(ctx context.Context, inc *model.RideHailingDriverVerifyIncInput, set *model.RideHailingDriverVerifySetInput, unionId string) (*model1.RideHailingDriverVerify, error) {
	var rs model1.RideHailingDriverVerify
	tx := db.DB.Where(rs.UnionPrimaryColumnName()+" = ?", unionId)
	qt := util.NewQueryTranslator(tx, &model1.RideHailingDriverVerify{})
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

func (r *queryResolver) RideHailingDriverVerify(ctx context.Context, distinctOn []model.RideHailingDriverVerifySelectColumn, limit *int, offset *int, orderBy []*model.RideHailingDriverVerifyOrderBy, where *model.RideHailingDriverVerifyBoolExp) ([]*model1.RideHailingDriverVerify, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.RideHailingDriverVerify{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.RideHailingDriverVerify
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) RideHailingDriverVerifyAggregate(ctx context.Context, distinctOn []model.RideHailingDriverVerifySelectColumn, limit *int, offset *int, orderBy []*model.RideHailingDriverVerifyOrderBy, where *model.RideHailingDriverVerifyBoolExp) (*model.RideHailingDriverVerifyAggregate, error) {
	var rs model.RideHailingDriverVerifyAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.RideHailingDriverVerify{})
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

func (r *queryResolver) RideHailingDriverVerifyByPk(ctx context.Context, Id int64) (*model1.RideHailingDriverVerify, error) {
	var rs model1.RideHailingDriverVerify
	tx := db.DB.Model(&model1.RideHailingDriverVerify{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}

func (r *queryResolver) RideHailingDriverVerifyByUnionPk(ctx context.Context, unionId string) (*model1.RideHailingDriverVerify, error) {
	var rs model1.RideHailingDriverVerify
	tx := db.DB.Model(&model1.RideHailingDriverVerify{}).Select(util.GetTopPreloads(ctx)).Where(rs.UnionPrimaryColumnName()+" = ?", unionId).First(&rs)

	err := tx.Error
	return &rs, err
}
