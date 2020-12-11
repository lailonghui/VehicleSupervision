package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/vehicle_alarm/graph/model"
	util2 "VehicleSupervision/pkg/graphql/util"
	"VehicleSupervision/pkg/util"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteDistrictAlarmContentPush(ctx context.Context, where model.DistrictAlarmContentPushBoolExp) (*model.DistrictAlarmContentPushMutationResponse, error) {
	var err error
	var rs []*model.DistrictAlarmContentPush
	qt := util2.NewQueryTranslator(db.DB, &model.DistrictAlarmContentPush{})
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
	return &model.DistrictAlarmContentPushMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, err
}

func (r *mutationResolver) DeleteDistrictAlarmContentPushByPk(ctx context.Context, id int64) (*model.DistrictAlarmContentPush, error) {
	var rs model.DistrictAlarmContentPush
	var err error
	tx := db.DB.Model(&model.DistrictAlarmContentPush{})
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

func (r *mutationResolver) InsertDistrictAlarmContentPush(ctx context.Context, objects []*model.DistrictAlarmContentPushInsertInput, onConflict *model.DistrictAlarmContentPushOnConflict) (*model.DistrictAlarmContentPushMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertDistrictAlarmContentPushOne(ctx context.Context, object model.DistrictAlarmContentPushInsertInput, onConflict *model.DistrictAlarmContentPushOnConflict) (*model.DistrictAlarmContentPush, error) {
	v := &model.DistrictAlarmContentPush{}
	util.StructAssign(v, &object)
	err := db.DB.Create(v).Error
	if err != nil {
		fmt.Println(err)
	}
	return v, err
}

func (r *mutationResolver) UpdateDistrictAlarmContentPush(ctx context.Context, inc *model.DistrictAlarmContentPushIncInput, set *model.DistrictAlarmContentPushSetInput, where model.DistrictAlarmContentPushBoolExp) (*model.DistrictAlarmContentPushMutationResponse, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.DistrictAlarmContentPush{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return &model.DistrictAlarmContentPushMutationResponse{
			AffectedRows: 0,
		}, err
	}
	var vehicleList []*model.DistrictAlarmContentPush
	tx.Scan(&vehicleList)
	return &model.DistrictAlarmContentPushMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    vehicleList,
	}, err
}

func (r *mutationResolver) UpdateDistrictAlarmContentPushByPk(ctx context.Context, inc *model.DistrictAlarmContentPushIncInput, set *model.DistrictAlarmContentPushSetInput, pkColumns model.DistrictAlarmContentPushPkColumnsInput) (*model.DistrictAlarmContentPush, error) {
	var err error
	tx := db.DB.Where("id = ? ", pkColumns.ID)
	qt := util2.NewQueryTranslator(tx, &model.DistrictAlarmContentPush{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return nil, err
	}
	var rs model.DistrictAlarmContentPush
	tx = tx.First(&rs)
	return &rs, err
}

func (r *queryResolver) DistrictAlarmContentPush(ctx context.Context, distinctOn []model.DistrictAlarmContentPushSelectColumn, limit *int, offset *int, orderBy []*model.DistrictAlarmContentPushOrderBy, where *model.DistrictAlarmContentPushBoolExp) ([]*model.DistrictAlarmContentPush, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.DistrictAlarmContentPush{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model.DistrictAlarmContentPush
	tx = tx.Find(&rs)
	if err = tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, err
}

func (r *queryResolver) DistrictAlarmContentPushAggregate(ctx context.Context, distinctOn []model.DistrictAlarmContentPushSelectColumn, limit *int, offset *int, orderBy []*model.DistrictAlarmContentPushOrderBy, where *model.DistrictAlarmContentPushBoolExp) (*model.DistrictAlarmContentPushAggregate, error) {
	var rs model.DistrictAlarmContentPushAggregate
	qt := util2.NewQueryTranslator(db.DB, &model.DistrictAlarmContentPush{})
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

func (r *queryResolver) DistrictAlarmContentPushByPk(ctx context.Context, id int64) (*model.DistrictAlarmContentPush, error) {
	var rs model.DistrictAlarmContentPush
	err := db.DB.Where("id = ?", id).First(&rs).Error
	return &rs, err
}

func (r *subscriptionResolver) DistrictAlarmContentPush(ctx context.Context, distinctOn []model.DistrictAlarmContentPushSelectColumn, limit *int, offset *int, orderBy []*model.DistrictAlarmContentPushOrderBy, where *model.DistrictAlarmContentPushBoolExp) (<-chan []*model.DistrictAlarmContentPush, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DistrictAlarmContentPushAggregate(ctx context.Context, distinctOn []model.DistrictAlarmContentPushSelectColumn, limit *int, offset *int, orderBy []*model.DistrictAlarmContentPushOrderBy, where *model.DistrictAlarmContentPushBoolExp) (<-chan *model.DistrictAlarmContentPushAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DistrictAlarmContentPushByPk(ctx context.Context, id int64) (<-chan *model.DistrictAlarmContentPush, error) {
	panic(fmt.Errorf("not implemented"))
}
