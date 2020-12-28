package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/vehicle_alarm/graph/model"
	model1 "VehicleSupervision/internal/modules/vehicle_alarm/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteAlarmSupervisionPictureUpload(ctx context.Context, where model.AlarmSupervisionPictureUploadBoolExp) (*model.AlarmSupervisionPictureUploadMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.AlarmSupervisionPictureUpload{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.AlarmSupervisionPictureUpload
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
	return &model.AlarmSupervisionPictureUploadMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteAlarmSupervisionPictureUploadByPk(ctx context.Context, Id int64) (*model1.AlarmSupervisionPictureUpload, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.AlarmSupervisionPictureUpload
	tx := db.DB.Model(&model1.AlarmSupervisionPictureUpload{})
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

func (r *mutationResolver) InsertAlarmSupervisionPictureUpload(ctx context.Context, objects []*model.AlarmSupervisionPictureUploadInsertInput) (*model.AlarmSupervisionPictureUploadMutationResponse, error) {
	rs := []*model1.AlarmSupervisionPictureUpload{}
	for _, object := range objects {
		v := &model1.AlarmSupervisionPictureUpload{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.AlarmSupervisionPictureUpload{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.AlarmSupervisionPictureUploadMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertAlarmSupervisionPictureUploadOne(ctx context.Context, object model.AlarmSupervisionPictureUploadInsertInput) (*model1.AlarmSupervisionPictureUpload, error) {
	rs := &model1.AlarmSupervisionPictureUpload{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.AlarmSupervisionPictureUpload{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateAlarmSupervisionPictureUpload(ctx context.Context, inc *model.AlarmSupervisionPictureUploadIncInput, set *model.AlarmSupervisionPictureUploadSetInput, where model.AlarmSupervisionPictureUploadBoolExp) (*model.AlarmSupervisionPictureUploadMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.AlarmSupervisionPictureUpload{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.AlarmSupervisionPictureUploadMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.AlarmSupervisionPictureUploadMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateAlarmSupervisionPictureUploadByPk(ctx context.Context, inc *model.AlarmSupervisionPictureUploadIncInput, set *model.AlarmSupervisionPictureUploadSetInput, Id int64) (*model1.AlarmSupervisionPictureUpload, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.AlarmSupervisionPictureUpload{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.AlarmSupervisionPictureUpload
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) AlarmSupervisionPictureUpload(ctx context.Context, distinctOn []model.AlarmSupervisionPictureUploadSelectColumn, limit *int, offset *int, orderBy []*model.AlarmSupervisionPictureUploadOrderBy, where *model.AlarmSupervisionPictureUploadBoolExp) ([]*model1.AlarmSupervisionPictureUpload, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.AlarmSupervisionPictureUpload{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.AlarmSupervisionPictureUpload
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) AlarmSupervisionPictureUploadAggregate(ctx context.Context, distinctOn []model.AlarmSupervisionPictureUploadSelectColumn, limit *int, offset *int, orderBy []*model.AlarmSupervisionPictureUploadOrderBy, where *model.AlarmSupervisionPictureUploadBoolExp) (*model.AlarmSupervisionPictureUploadAggregate, error) {
	var rs model.AlarmSupervisionPictureUploadAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.AlarmSupervisionPictureUpload{})
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

func (r *queryResolver) AlarmSupervisionPictureUploadByPk(ctx context.Context, Id int64) (*model1.AlarmSupervisionPictureUpload, error) {
	var rs model1.AlarmSupervisionPictureUpload
	tx := db.DB.Model(&model1.AlarmSupervisionPictureUpload{}).First(&rs, Id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
