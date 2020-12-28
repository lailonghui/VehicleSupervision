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

func (r *mutationResolver) DeleteMuckTruckPreviewNumber(ctx context.Context, where model.MuckTruckPreviewNumberBoolExp) (*model.MuckTruckPreviewNumberMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.MuckTruckPreviewNumber{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.MuckTruckPreviewNumber
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
	return &model.MuckTruckPreviewNumberMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteMuckTruckPreviewNumberByPk(ctx context.Context, Id int64) (*model1.MuckTruckPreviewNumber, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.MuckTruckPreviewNumber
	tx := db.DB.Model(&model1.MuckTruckPreviewNumber{})
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

func (r *mutationResolver) InsertMuckTruckPreviewNumber(ctx context.Context, objects []*model.MuckTruckPreviewNumberInsertInput) (*model.MuckTruckPreviewNumberMutationResponse, error) {
	rs := []*model1.MuckTruckPreviewNumber{}
	for _, object := range objects {
		v := &model1.MuckTruckPreviewNumber{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.MuckTruckPreviewNumber{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.MuckTruckPreviewNumberMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertMuckTruckPreviewNumberOne(ctx context.Context, object model.MuckTruckPreviewNumberInsertInput) (*model1.MuckTruckPreviewNumber, error) {
	rs := &model1.MuckTruckPreviewNumber{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.MuckTruckPreviewNumber{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateMuckTruckPreviewNumber(ctx context.Context, inc *model.MuckTruckPreviewNumberIncInput, set *model.MuckTruckPreviewNumberSetInput, where model.MuckTruckPreviewNumberBoolExp) (*model.MuckTruckPreviewNumberMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.MuckTruckPreviewNumber{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.MuckTruckPreviewNumberMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.MuckTruckPreviewNumberMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateMuckTruckPreviewNumberByPk(ctx context.Context, inc *model.MuckTruckPreviewNumberIncInput, set *model.MuckTruckPreviewNumberSetInput, Id int64) (*model1.MuckTruckPreviewNumber, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.MuckTruckPreviewNumber{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.MuckTruckPreviewNumber
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) MuckTruckPreviewNumber(ctx context.Context, distinctOn []model.MuckTruckPreviewNumberSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckPreviewNumberOrderBy, where *model.MuckTruckPreviewNumberBoolExp) ([]*model1.MuckTruckPreviewNumber, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.MuckTruckPreviewNumber{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.MuckTruckPreviewNumber
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) MuckTruckPreviewNumberAggregate(ctx context.Context, distinctOn []model.MuckTruckPreviewNumberSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckPreviewNumberOrderBy, where *model.MuckTruckPreviewNumberBoolExp) (*model.MuckTruckPreviewNumberAggregate, error) {
	var rs model.MuckTruckPreviewNumberAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.MuckTruckPreviewNumber{})
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

func (r *queryResolver) MuckTruckPreviewNumberByPk(ctx context.Context, Id int64) (*model1.MuckTruckPreviewNumber, error) {
	var rs model1.MuckTruckPreviewNumber
	tx := db.DB.Model(&model1.MuckTruckPreviewNumber{}).First(&rs, Id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
