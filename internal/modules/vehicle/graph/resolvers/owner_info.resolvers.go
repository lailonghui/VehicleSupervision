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

func (r *mutationResolver) DeleteOwnerInfo(ctx context.Context, where model.OwnerInfoBoolExp) (*model.OwnerInfoMutationResponse, error) {
	var err error
	var rs []*model.OwnerInfo
	qt := util2.NewQueryTranslator(db.DB, &model.OwnerInfo{})
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
	return &model.OwnerInfoMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, err
}

func (r *mutationResolver) DeleteOwnerInfoByPk(ctx context.Context, id int64) (*model.OwnerInfo, error) {
	var rs model.OwnerInfo
	var err error
	tx := db.DB.Model(&model.OwnerInfo{})
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

func (r *mutationResolver) InsertOwnerInfo(ctx context.Context, objects []*model.OwnerInfoInsertInput, onConflict *model.OwnerInfoOnConflict) (*model.OwnerInfoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertOwnerInfoOne(ctx context.Context, object model.OwnerInfoInsertInput, onConflict *model.OwnerInfoOnConflict) (*model.OwnerInfo, error) {
	v := &model.OwnerInfo{}
	util.StructAssign(v, &object)
	err := db.DB.Create(v).Error
	if err != nil {
		fmt.Println(err)
	}
	return v, err
}

func (r *mutationResolver) UpdateOwnerInfo(ctx context.Context, inc *model.OwnerInfoIncInput, set *model.OwnerInfoSetInput, where model.OwnerInfoBoolExp) (*model.OwnerInfoMutationResponse, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.OwnerInfo{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return &model.OwnerInfoMutationResponse{
			AffectedRows: 0,
		}, err
	}
	var vehicleList []*model.OwnerInfo
	tx.Scan(&vehicleList)
	return &model.OwnerInfoMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    vehicleList,
	}, err
}

func (r *mutationResolver) UpdateOwnerInfoByPk(ctx context.Context, inc *model.OwnerInfoIncInput, set *model.OwnerInfoSetInput, pkColumns model.OwnerInfoPkColumnsInput) (*model.OwnerInfo, error) {
	var err error
	tx := db.DB.Where("id = ? ", pkColumns.ID)
	qt := util2.NewQueryTranslator(tx, &model.OwnerInfo{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return nil, err
	}
	var rs model.OwnerInfo
	tx = tx.First(&rs)
	return &rs, err
}

func (r *queryResolver) OwnerInfo(ctx context.Context, distinctOn []model.OwnerInfoSelectColumn, limit *int, offset *int, orderBy []*model.OwnerInfoOrderBy, where *model.OwnerInfoBoolExp) ([]*model.OwnerInfo, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.OwnerInfo{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model.OwnerInfo
	tx = tx.Find(&rs)
	if err = tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, err
}

func (r *queryResolver) OwnerInfoAggregate(ctx context.Context, distinctOn []model.OwnerInfoSelectColumn, limit *int, offset *int, orderBy []*model.OwnerInfoOrderBy, where *model.OwnerInfoBoolExp) (*model.OwnerInfoAggregate, error) {
	var rs model.OwnerInfoAggregate
	qt := util2.NewQueryTranslator(db.DB, &model.OwnerInfo{})
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

func (r *queryResolver) OwnerInfoByPk(ctx context.Context, id int64) (*model.OwnerInfo, error) {
	var rs model.OwnerInfo
	err := db.DB.Where("id = ?", id).First(&rs).Error
	return &rs, err
}

func (r *subscriptionResolver) OwnerInfo(ctx context.Context, distinctOn []model.OwnerInfoSelectColumn, limit *int, offset *int, orderBy []*model.OwnerInfoOrderBy, where *model.OwnerInfoBoolExp) (<-chan []*model.OwnerInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) OwnerInfoAggregate(ctx context.Context, distinctOn []model.OwnerInfoSelectColumn, limit *int, offset *int, orderBy []*model.OwnerInfoOrderBy, where *model.OwnerInfoBoolExp) (<-chan *model.OwnerInfoAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) OwnerInfoByPk(ctx context.Context, id int64) (<-chan *model.OwnerInfo, error) {
	panic(fmt.Errorf("not implemented"))
}
