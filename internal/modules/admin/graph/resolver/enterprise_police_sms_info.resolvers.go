package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/admin/graph/model"
	model1 "VehicleSupervision/internal/modules/admin/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteEnterprisePoliceSmsInfo(ctx context.Context, where model.EnterprisePoliceSmsInfoBoolExp) (*model.EnterprisePoliceSmsInfoMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EnterprisePoliceSmsInfo{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.EnterprisePoliceSmsInfo
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
	return &model.EnterprisePoliceSmsInfoMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteEnterprisePoliceSmsInfoByPk(ctx context.Context, id int64) (*model1.EnterprisePoliceSmsInfo, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.EnterprisePoliceSmsInfo
	tx := db.DB.Model(&model1.EnterprisePoliceSmsInfo{})
	if len(preloads) > 0 {
		// 如果请求的字段不为空，则先查询一遍数据库
		tx = tx.Select(preloads).Where(rs.PrimaryColumnName()+" = ?", id).First(&rs)
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

func (r *mutationResolver) DeleteEnterprisePoliceSmsInfoByUnionPk(ctx context.Context, enterprisePoliceSmsInfoID string) (*model1.EnterprisePoliceSmsInfo, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.EnterprisePoliceSmsInfo
	tx := db.DB.Model(&model1.EnterprisePoliceSmsInfo{})
	if len(preloads) > 0 {
		// 如果请求的字段不为空，则先查询一遍数据库
		tx = tx.Select(preloads).Where(rs.UnionPrimaryColumnName()+" = ?", enterprisePoliceSmsInfoID).First(&rs)
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

func (r *mutationResolver) InsertEnterprisePoliceSmsInfo(ctx context.Context, objects []*model.EnterprisePoliceSmsInfoInsertInput) (*model.EnterprisePoliceSmsInfoMutationResponse, error) {
	rs := make([]*model1.EnterprisePoliceSmsInfo, 0)
	for _, object := range objects {
		v := &model1.EnterprisePoliceSmsInfo{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.EnterprisePoliceSmsInfo{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.EnterprisePoliceSmsInfoMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertEnterprisePoliceSmsInfoOne(ctx context.Context, objects model.EnterprisePoliceSmsInfoInsertInput) (*model1.EnterprisePoliceSmsInfo, error) {
	rs := &model1.EnterprisePoliceSmsInfo{}
	util2.StructAssign(rs, &objects)
	tx := db.DB.Model(&model1.EnterprisePoliceSmsInfo{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateEnterprisePoliceSmsInfo(ctx context.Context, inc *model.EnterprisePoliceSmsInfoIncInput, set *model.EnterprisePoliceSmsInfoSetInput, where model.EnterprisePoliceSmsInfoBoolExp) (*model.EnterprisePoliceSmsInfoMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EnterprisePoliceSmsInfo{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.EnterprisePoliceSmsInfoMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.EnterprisePoliceSmsInfo
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
	return &model.EnterprisePoliceSmsInfoMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateEnterprisePoliceSmsInfoByPk(ctx context.Context, inc *model.EnterprisePoliceSmsInfoIncInput, set *model.EnterprisePoliceSmsInfoSetInput, id int64) (*model1.EnterprisePoliceSmsInfo, error) {
	var rs model1.EnterprisePoliceSmsInfo
	tx := db.DB.Where(rs.PrimaryColumnName()+" = ?", id)
	qt := util.NewQueryTranslator(tx, &model1.EnterprisePoliceSmsInfo{})
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

func (r *mutationResolver) UpdateEnterprisePoliceSmsInfoByUnionPk(ctx context.Context, inc *model.EnterprisePoliceSmsInfoIncInput, set *model.EnterprisePoliceSmsInfoSetInput, enterprisePoliceSmsInfoID string) (*model1.EnterprisePoliceSmsInfo, error) {
	var rs model1.EnterprisePoliceSmsInfo
	tx := db.DB.Where(rs.UnionPrimaryColumnName()+" = ?", enterprisePoliceSmsInfoID)
	qt := util.NewQueryTranslator(tx, &model1.EnterprisePoliceSmsInfo{})
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

func (r *queryResolver) EnterprisePoliceSmsInfo(ctx context.Context, distinctOn []model.EnterprisePoliceSmsInfoSelectColumn, limit *int, offset *int, orderBy []*model.EnterprisePoliceSmsInfoOrderBy, where *model.EnterprisePoliceSmsInfoBoolExp) ([]*model1.EnterprisePoliceSmsInfo, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EnterprisePoliceSmsInfo{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.EnterprisePoliceSmsInfo
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) EnterprisePoliceSmsInfoAggregate(ctx context.Context, distinctOn []model.EnterprisePoliceSmsInfoSelectColumn, limit *int, offset *int, orderBy []*model.EnterprisePoliceSmsInfoOrderBy, where *model.EnterprisePoliceSmsInfoBoolExp) (*model.EnterprisePoliceSmsInfoAggregate, error) {
	var rs model.EnterprisePoliceSmsInfoAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.EnterprisePoliceSmsInfo{})
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

func (r *queryResolver) EnterprisePoliceSmsInfoByPk(ctx context.Context, id int64) (*model1.EnterprisePoliceSmsInfo, error) {
	var rs model1.EnterprisePoliceSmsInfo
	tx := db.DB.Model(&model1.EnterprisePoliceSmsInfo{}).Select(util.GetTopPreloads(ctx)).First(&rs, id)
	err := tx.Error
	return &rs, err
}

func (r *queryResolver) EnterprisePoliceSmsInfoByUnionPk(ctx context.Context, enterprisePoliceSmsInfoID string) (*model1.EnterprisePoliceSmsInfo, error) {
	var rs model1.EnterprisePoliceSmsInfo
	tx := db.DB.Model(&model1.EnterprisePoliceSmsInfo{}).Select(util.GetTopPreloads(ctx)).Where(rs.UnionPrimaryColumnName()+" = ?", enterprisePoliceSmsInfoID).First(&rs)

	err := tx.Error
	return &rs, err
}
