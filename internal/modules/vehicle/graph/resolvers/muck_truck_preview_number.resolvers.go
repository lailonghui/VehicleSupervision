package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/vehicle/graph/model"
	util2 "VehicleSupervision/pkg/graphql/util"
	"VehicleSupervision/pkg/util"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteMuckTruckPreviewNumber(ctx context.Context, where model.MuckTruckPreviewNumberBoolExp) (*model.MuckTruckPreviewNumberMutationResponse, error) {
	var err error
	var rs []*model.MuckTruckPreviewNumber
	qt := util2.NewQueryTranslator(db.DB, &model.MuckTruckPreviewNumber{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util2.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	if len(preloads) > 0 {
		// 如果请求的字段不为空，则先查询一遍数据库
		tx := tx.Select(preloads)
		tx = tx.Find(&rs)
		// 如果查询结果含有错误，则返回错误
		if err := tx.Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, nil
			}
		}
	}
	// 删除
	tx = tx.Delete(nil)
	if err = tx.Error; err != nil {
		return nil, err
	}
	return &model.MuckTruckPreviewNumberMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteMuckTruckPreviewNumberByPk(ctx context.Context, id int64) (*model.MuckTruckPreviewNumber, error) {
	var rs model.MuckTruckPreviewNumber
	var err error
	tx := db.DB.Model(&model.MuckTruckPreviewNumber{})
	preloads := util2.GetPreloads(ctx)
	if len(preloads) > 0 {
		// 如果请求的字段不为空，则先查询一遍数据库
		tx = tx.Select(preloads).Where("id = ? ", id).First(&rs)
		// 如果查询结果含有错误，则返回错误
		if err := tx.Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				err = errors.New("0 rows affected or returned")
				return nil, err
			}
		}
	}
	// 删除
	tx = tx.Delete(nil)
	if err = tx.Error; err != nil {
		return nil, err
	}
	return &rs, nil
}

func (r *mutationResolver) InsertMuckTruckPreviewNumber(ctx context.Context, objects []*model.MuckTruckPreviewNumberInsertInput, onConflict *model.MuckTruckPreviewNumberOnConflict) (*model.MuckTruckPreviewNumberMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertMuckTruckPreviewNumberOne(ctx context.Context, object model.MuckTruckPreviewNumberInsertInput, onConflict *model.MuckTruckPreviewNumberOnConflict) (*model.MuckTruckPreviewNumber, error) {
	v := &model.MuckTruckPreviewNumber{}
	util.StructAssign(v, &object)
	err := db.DB.Create(v).Error
	if err != nil {
		fmt.Println(err)
	}
	return v, err
}

func (r *mutationResolver) UpdateMuckTruckPreviewNumber(ctx context.Context, inc *model.MuckTruckPreviewNumberIncInput, set *model.MuckTruckPreviewNumberSetInput, where model.MuckTruckPreviewNumberBoolExp) (*model.MuckTruckPreviewNumberMutationResponse, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.MuckTruckPreviewNumber{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return &model.MuckTruckPreviewNumberMutationResponse{
			AffectedRows: 0,
		}, err
	}
	var vehicleList []*model.MuckTruckPreviewNumber
	tx.Scan(&vehicleList)
	return &model.MuckTruckPreviewNumberMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    vehicleList,
	}, err
}

func (r *mutationResolver) UpdateMuckTruckPreviewNumberByPk(ctx context.Context, inc *model.MuckTruckPreviewNumberIncInput, set *model.MuckTruckPreviewNumberSetInput, pkColumns model.MuckTruckPreviewNumberPkColumnsInput) (*model.MuckTruckPreviewNumber, error) {
	var err error
	tx := db.DB.Where("id = ? ", pkColumns.ID)
	qt := util2.NewQueryTranslator(tx, &model.MuckTruckPreviewNumber{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return nil, err
	}
	var rs model.MuckTruckPreviewNumber
	tx = tx.First(&rs)
	return &rs, err
}

func (r *queryResolver) MuckTruckPreviewNumber(ctx context.Context, distinctOn []model.MuckTruckPreviewNumberSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckPreviewNumberOrderBy, where *model.MuckTruckPreviewNumberBoolExp) ([]*model.MuckTruckPreviewNumber, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.MuckTruckPreviewNumber{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model.MuckTruckPreviewNumber
	tx = tx.Find(&rs)
	if err = tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, err
}

func (r *queryResolver) MuckTruckPreviewNumberAggregate(ctx context.Context, distinctOn []model.MuckTruckPreviewNumberSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckPreviewNumberOrderBy, where *model.MuckTruckPreviewNumberBoolExp) (*model.MuckTruckPreviewNumberAggregate, error) {
	var rs model.MuckTruckPreviewNumberAggregate
	qt := util2.NewQueryTranslator(db.DB, &model.MuckTruckPreviewNumber{})
	tx, err := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Aggregate(&rs, ctx)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if err = tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}
	if err != nil {
		fmt.Println(err)
	}
	return &rs, err
}

func (r *queryResolver) MuckTruckPreviewNumberByPk(ctx context.Context, id int64) (*model.MuckTruckPreviewNumber, error) {
	var rs model.MuckTruckPreviewNumber
	err := db.DB.Where("id = ?", id).First(&rs).Error
	return &rs, err
}

func (r *subscriptionResolver) MuckTruckPreviewNumber(ctx context.Context, distinctOn []model.MuckTruckPreviewNumberSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckPreviewNumberOrderBy, where *model.MuckTruckPreviewNumberBoolExp) (<-chan []*model.MuckTruckPreviewNumber, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) MuckTruckPreviewNumberAggregate(ctx context.Context, distinctOn []model.MuckTruckPreviewNumberSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckPreviewNumberOrderBy, where *model.MuckTruckPreviewNumberBoolExp) (<-chan *model.MuckTruckPreviewNumberAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) MuckTruckPreviewNumberByPk(ctx context.Context, id int64) (<-chan *model.MuckTruckPreviewNumber, error) {
	panic(fmt.Errorf("not implemented"))
}
