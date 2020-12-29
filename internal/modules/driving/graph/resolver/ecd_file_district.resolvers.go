package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/driving/graph/model"
	model1 "VehicleSupervision/internal/modules/driving/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteEcdFileDistrict(ctx context.Context, where model.EcdFileDistrictBoolExp) (*model.EcdFileDistrictMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EcdFileDistrict{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.EcdFileDistrict
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
	return &model.EcdFileDistrictMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteEcdFileDistrictByPk(ctx context.Context, Id int64) (*model1.EcdFileDistrict, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.EcdFileDistrict
	tx := db.DB.Model(&model1.EcdFileDistrict{})
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

func (r *mutationResolver) InsertEcdFileDistrict(ctx context.Context, objects []*model.EcdFileDistrictInsertInput) (*model.EcdFileDistrictMutationResponse, error) {
	rs := make([]*model1.EcdFileDistrict, 0)
	for _, object := range objects {
		v := &model1.EcdFileDistrict{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.EcdFileDistrict{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.EcdFileDistrictMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertEcdFileDistrictOne(ctx context.Context, object model.EcdFileDistrictInsertInput) (*model1.EcdFileDistrict, error) {
	rs := &model1.EcdFileDistrict{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.EcdFileDistrict{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateEcdFileDistrict(ctx context.Context, inc *model.EcdFileDistrictIncInput, set *model.EcdFileDistrictSetInput, where model.EcdFileDistrictBoolExp) (*model.EcdFileDistrictMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EcdFileDistrict{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.EcdFileDistrictMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.EcdFileDistrict
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
	return &model.EcdFileDistrictMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateEcdFileDistrictByPk(ctx context.Context, inc *model.EcdFileDistrictIncInput, set *model.EcdFileDistrictSetInput, Id int64) (*model1.EcdFileDistrict, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.EcdFileDistrict{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.EcdFileDistrict
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) EcdFileDistrict(ctx context.Context, distinctOn []model.EcdFileDistrictSelectColumn, limit *int, offset *int, orderBy []*model.EcdFileDistrictOrderBy, where *model.EcdFileDistrictBoolExp) ([]*model1.EcdFileDistrict, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EcdFileDistrict{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.EcdFileDistrict
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) EcdFileDistrictAggregate(ctx context.Context, distinctOn []model.EcdFileDistrictSelectColumn, limit *int, offset *int, orderBy []*model.EcdFileDistrictOrderBy, where *model.EcdFileDistrictBoolExp) (*model.EcdFileDistrictAggregate, error) {
	var rs model.EcdFileDistrictAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.EcdFileDistrict{})
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

func (r *queryResolver) EcdFileDistrictByPk(ctx context.Context, Id int64) (*model1.EcdFileDistrict, error) {
	var rs model1.EcdFileDistrict
	tx := db.DB.Model(&model1.EcdFileDistrict{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}
