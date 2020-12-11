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

func (r *mutationResolver) DeleteAlarmSupervisionPictureUpload(ctx context.Context, where model.AlarmSupervisionPictureUploadBoolExp) (*model.AlarmSupervisionPictureUploadMutationResponse, error) {
	var err error
	var rs []*model.AlarmSupervisionPictureUpload
	qt := util2.NewQueryTranslator(db.DB, &model.AlarmSupervisionPictureUpload{})
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
	return &model.AlarmSupervisionPictureUploadMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, err
}

func (r *mutationResolver) DeleteAlarmSupervisionPictureUploadByPk(ctx context.Context, id int64) (*model.AlarmSupervisionPictureUpload, error) {
	var rs model.AlarmSupervisionPictureUpload
	var err error
	tx := db.DB.Model(&model.AlarmSupervisionPictureUpload{})
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

func (r *mutationResolver) InsertAlarmSupervisionPictureUpload(ctx context.Context, objects []*model.AlarmSupervisionPictureUploadInsertInput, onConflict *model.AlarmSupervisionPictureUploadOnConflict) (*model.AlarmSupervisionPictureUploadMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertAlarmSupervisionPictureUploadOne(ctx context.Context, object model.AlarmSupervisionPictureUploadInsertInput, onConflict *model.AlarmSupervisionPictureUploadOnConflict) (*model.AlarmSupervisionPictureUpload, error) {
	v := &model.AlarmSupervisionPictureUpload{}
	util.StructAssign(v, &object)
	err := db.DB.Create(v).Error
	if err != nil {
		fmt.Println(err)
	}
	return v, err
}

func (r *mutationResolver) UpdateAlarmSupervisionPictureUpload(ctx context.Context, inc *model.AlarmSupervisionPictureUploadIncInput, set *model.AlarmSupervisionPictureUploadSetInput, where model.AlarmSupervisionPictureUploadBoolExp) (*model.AlarmSupervisionPictureUploadMutationResponse, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.AlarmSupervisionPictureUpload{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return &model.AlarmSupervisionPictureUploadMutationResponse{
			AffectedRows: 0,
		}, err
	}
	var vehicleList []*model.AlarmSupervisionPictureUpload
	tx.Scan(&vehicleList)
	return &model.AlarmSupervisionPictureUploadMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    vehicleList,
	}, err
}

func (r *mutationResolver) UpdateAlarmSupervisionPictureUploadByPk(ctx context.Context, inc *model.AlarmSupervisionPictureUploadIncInput, set *model.AlarmSupervisionPictureUploadSetInput, pkColumns model.AlarmSupervisionPictureUploadPkColumnsInput) (*model.AlarmSupervisionPictureUpload, error) {
	var err error
	tx := db.DB.Where("id = ? ", pkColumns.ID)
	qt := util2.NewQueryTranslator(tx, &model.AlarmSupervisionPictureUpload{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return nil, err
	}
	var rs model.AlarmSupervisionPictureUpload
	tx = tx.First(&rs)
	return &rs, err
}

func (r *queryResolver) AlarmSupervisionPictureUpload(ctx context.Context, distinctOn []model.AlarmSupervisionPictureUploadSelectColumn, limit *int, offset *int, orderBy []*model.AlarmSupervisionPictureUploadOrderBy, where *model.AlarmSupervisionPictureUploadBoolExp) ([]*model.AlarmSupervisionPictureUpload, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.AlarmSupervisionPictureUpload{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model.AlarmSupervisionPictureUpload
	tx = tx.Find(&rs)
	if err = tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, err
}

func (r *queryResolver) AlarmSupervisionPictureUploadAggregate(ctx context.Context, distinctOn []model.AlarmSupervisionPictureUploadSelectColumn, limit *int, offset *int, orderBy []*model.AlarmSupervisionPictureUploadOrderBy, where *model.AlarmSupervisionPictureUploadBoolExp) (*model.AlarmSupervisionPictureUploadAggregate, error) {
	var rs model.AlarmSupervisionPictureUploadAggregate
	qt := util2.NewQueryTranslator(db.DB, &model.AlarmSupervisionPictureUpload{})
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

func (r *queryResolver) AlarmSupervisionPictureUploadByPk(ctx context.Context, id int64) (*model.AlarmSupervisionPictureUpload, error) {
	var rs model.AlarmSupervisionPictureUpload
	err := db.DB.Where("id = ?", id).First(&rs).Error
	return &rs, err
}

func (r *subscriptionResolver) AlarmSupervisionPictureUpload(ctx context.Context, distinctOn []model.AlarmSupervisionPictureUploadSelectColumn, limit *int, offset *int, orderBy []*model.AlarmSupervisionPictureUploadOrderBy, where *model.AlarmSupervisionPictureUploadBoolExp) (<-chan []*model.AlarmSupervisionPictureUpload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) AlarmSupervisionPictureUploadAggregate(ctx context.Context, distinctOn []model.AlarmSupervisionPictureUploadSelectColumn, limit *int, offset *int, orderBy []*model.AlarmSupervisionPictureUploadOrderBy, where *model.AlarmSupervisionPictureUploadBoolExp) (<-chan *model.AlarmSupervisionPictureUploadAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) AlarmSupervisionPictureUploadByPk(ctx context.Context, id int64) (<-chan *model.AlarmSupervisionPictureUpload, error) {
	panic(fmt.Errorf("not implemented"))
}
