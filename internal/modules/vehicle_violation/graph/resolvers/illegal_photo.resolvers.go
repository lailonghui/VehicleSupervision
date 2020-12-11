package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/vehicle_violation/graph/model"
	util2 "VehicleSupervision/pkg/graphql/util"
	"VehicleSupervision/pkg/util"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteIllegalPhoto(ctx context.Context, where model.IllegalPhotoBoolExp) (*model.IllegalPhotoMutationResponse, error) {
	var err error
	var rs []*model.IllegalPhoto
	qt := util2.NewQueryTranslator(db.DB, &model.IllegalPhoto{})
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
	return &model.IllegalPhotoMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, err
}

func (r *mutationResolver) DeleteIllegalPhotoByPk(ctx context.Context, id int64, illegalPhotoID string) (*model.IllegalPhoto, error) {
	var rs model.IllegalPhoto
	var err error
	tx := db.DB.Model(&model.IllegalPhoto{})
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

func (r *mutationResolver) InsertIllegalPhoto(ctx context.Context, objects []*model.IllegalPhotoInsertInput, onConflict *model.IllegalPhotoOnConflict) (*model.IllegalPhotoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertIllegalPhotoOne(ctx context.Context, object model.IllegalPhotoInsertInput, onConflict *model.IllegalPhotoOnConflict) (*model.IllegalPhoto, error) {
	v := &model.IllegalPhoto{}
	util.StructAssign(v, &object)
	err := db.DB.Create(v).Error
	if err != nil {
		fmt.Println(err)
	}
	return v, err
}

func (r *mutationResolver) UpdateIllegalPhoto(ctx context.Context, inc *model.IllegalPhotoIncInput, set *model.IllegalPhotoSetInput, where model.IllegalPhotoBoolExp) (*model.IllegalPhotoMutationResponse, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.IllegalPhoto{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return &model.IllegalPhotoMutationResponse{
			AffectedRows: 0,
		}, err
	}
	var vehicleList []*model.IllegalPhoto
	tx.Scan(&vehicleList)
	return &model.IllegalPhotoMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    vehicleList,
	}, err
}

func (r *mutationResolver) UpdateIllegalPhotoByPk(ctx context.Context, inc *model.IllegalPhotoIncInput, set *model.IllegalPhotoSetInput, pkColumns model.IllegalPhotoPkColumnsInput) (*model.IllegalPhoto, error) {
	var err error
	tx := db.DB.Where("id = ? ", pkColumns.ID)
	qt := util2.NewQueryTranslator(tx, &model.IllegalPhoto{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return nil, err
	}
	var rs model.IllegalPhoto
	tx = tx.First(&rs)
	return &rs, err
}

func (r *queryResolver) IllegalPhoto(ctx context.Context, distinctOn []model.IllegalPhotoSelectColumn, limit *int, offset *int, orderBy []*model.IllegalPhotoOrderBy, where *model.IllegalPhotoBoolExp) ([]*model.IllegalPhoto, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.IllegalPhoto{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model.IllegalPhoto
	tx = tx.Find(&rs)
	if err = tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, err
}

func (r *queryResolver) IllegalPhotoAggregate(ctx context.Context, distinctOn []model.IllegalPhotoSelectColumn, limit *int, offset *int, orderBy []*model.IllegalPhotoOrderBy, where *model.IllegalPhotoBoolExp) (*model.IllegalPhotoAggregate, error) {
	var rs model.IllegalPhotoAggregate
	qt := util2.NewQueryTranslator(db.DB, &model.IllegalPhoto{})
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

func (r *queryResolver) IllegalPhotoByPk(ctx context.Context, id int64, illegalPhotoID string) (*model.IllegalPhoto, error) {
	var rs model.IllegalPhoto
	err := db.DB.Where("id = ?", id).First(&rs).Error
	return &rs, err
}

func (r *subscriptionResolver) IllegalPhoto(ctx context.Context, distinctOn []model.IllegalPhotoSelectColumn, limit *int, offset *int, orderBy []*model.IllegalPhotoOrderBy, where *model.IllegalPhotoBoolExp) (<-chan []*model.IllegalPhoto, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) IllegalPhotoAggregate(ctx context.Context, distinctOn []model.IllegalPhotoSelectColumn, limit *int, offset *int, orderBy []*model.IllegalPhotoOrderBy, where *model.IllegalPhotoBoolExp) (<-chan *model.IllegalPhotoAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) IllegalPhotoByPk(ctx context.Context, id int64, illegalPhotoID string) (<-chan *model.IllegalPhoto, error) {
	panic(fmt.Errorf("not implemented"))
}
