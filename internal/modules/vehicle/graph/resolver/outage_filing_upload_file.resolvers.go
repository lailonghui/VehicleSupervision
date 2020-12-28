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

func (r *mutationResolver) DeleteOutageFilingUploadFile(ctx context.Context, where model.OutageFilingUploadFileBoolExp) (*model.OutageFilingUploadFileMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.OutageFilingUploadFile{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.OutageFilingUploadFile
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
	return &model.OutageFilingUploadFileMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteOutageFilingUploadFileByPk(ctx context.Context, Id int64) (*model1.OutageFilingUploadFile, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.OutageFilingUploadFile
	tx := db.DB.Model(&model1.OutageFilingUploadFile{})
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

func (r *mutationResolver) InsertOutageFilingUploadFile(ctx context.Context, objects []*model.OutageFilingUploadFileInsertInput) (*model.OutageFilingUploadFileMutationResponse, error) {
	rs := []*model1.OutageFilingUploadFile{}
	for _, object := range objects {
		v := &model1.OutageFilingUploadFile{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.OutageFilingUploadFile{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.OutageFilingUploadFileMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertOutageFilingUploadFileOne(ctx context.Context, object model.OutageFilingUploadFileInsertInput) (*model1.OutageFilingUploadFile, error) {
	rs := &model1.OutageFilingUploadFile{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.OutageFilingUploadFile{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateOutageFilingUploadFile(ctx context.Context, inc *model.OutageFilingUploadFileIncInput, set *model.OutageFilingUploadFileSetInput, where model.OutageFilingUploadFileBoolExp) (*model.OutageFilingUploadFileMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.OutageFilingUploadFile{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.OutageFilingUploadFileMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.OutageFilingUploadFileMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateOutageFilingUploadFileByPk(ctx context.Context, inc *model.OutageFilingUploadFileIncInput, set *model.OutageFilingUploadFileSetInput, Id int64) (*model1.OutageFilingUploadFile, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.OutageFilingUploadFile{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.OutageFilingUploadFile
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) OutageFilingUploadFile(ctx context.Context, distinctOn []model.OutageFilingUploadFileSelectColumn, limit *int, offset *int, orderBy []*model.OutageFilingUploadFileOrderBy, where *model.OutageFilingUploadFileBoolExp) ([]*model1.OutageFilingUploadFile, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.OutageFilingUploadFile{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.OutageFilingUploadFile
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) OutageFilingUploadFileAggregate(ctx context.Context, distinctOn []model.OutageFilingUploadFileSelectColumn, limit *int, offset *int, orderBy []*model.OutageFilingUploadFileOrderBy, where *model.OutageFilingUploadFileBoolExp) (*model.OutageFilingUploadFileAggregate, error) {
	var rs model.OutageFilingUploadFileAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.OutageFilingUploadFile{})
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

func (r *queryResolver) OutageFilingUploadFileByPk(ctx context.Context, Id int64) (*model1.OutageFilingUploadFile, error) {
	var rs model1.OutageFilingUploadFile
	tx := db.DB.Model(&model1.OutageFilingUploadFile{}).First(&rs, Id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
