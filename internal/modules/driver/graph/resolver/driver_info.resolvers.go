package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/driver/graph/generated"
	"VehicleSupervision/internal/modules/driver/graph/model"
	model1 "VehicleSupervision/internal/modules/driver/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteDriverInfo(ctx context.Context, where model.DriverInfoBoolExp) (*model.DriverInfoMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DriverInfo{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.DriverInfo
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
	return &model.DriverInfoMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteDriverInfoByPk(ctx context.Context, Id int64) (*model1.DriverInfo, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.DriverInfo
	tx := db.DB.Model(&model1.DriverInfo{})
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

func (r *mutationResolver) InsertDriverInfo(ctx context.Context, objects []*model.DriverInfoInsertInput) (*model.DriverInfoMutationResponse, error) {
	rs := make([]*model1.DriverInfo, 0)
	for _, object := range objects {
		v := &model1.DriverInfo{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.DriverInfo{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.DriverInfoMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertDriverInfoOne(ctx context.Context, object model.DriverInfoInsertInput) (*model1.DriverInfo, error) {
	rs := &model1.DriverInfo{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.DriverInfo{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateDriverInfo(ctx context.Context, inc *model.DriverInfoIncInput, set *model.DriverInfoSetInput, where model.DriverInfoBoolExp) (*model.DriverInfoMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DriverInfo{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.DriverInfoMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.DriverInfoMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateDriverInfoByPk(ctx context.Context, inc *model.DriverInfoIncInput, set *model.DriverInfoSetInput, Id int64) (*model1.DriverInfo, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.DriverInfo{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.DriverInfo
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) DriverInfo(ctx context.Context, distinctOn []model.DriverInfoSelectColumn, limit *int, offset *int, orderBy []*model.DriverInfoOrderBy, where *model.DriverInfoBoolExp) ([]*model1.DriverInfo, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DriverInfo{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.DriverInfo
	tx = tx.Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) DriverInfoAggregate(ctx context.Context, distinctOn []model.DriverInfoSelectColumn, limit *int, offset *int, orderBy []*model.DriverInfoOrderBy, where *model.DriverInfoBoolExp) (*model.DriverInfoAggregate, error) {
	var rs model.DriverInfoAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.DriverInfo{})
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

func (r *queryResolver) DriverInfoByPk(ctx context.Context, Id int64) (*model1.DriverInfo, error) {
	var rs model1.DriverInfo
	tx := db.DB.Model(&model1.DriverInfo{}).First(&rs, Id)
	err := tx.Error
	return &rs, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
