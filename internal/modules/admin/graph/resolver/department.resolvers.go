package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/admin/graph/model"
	model1 "VehicleSupervision/internal/modules/admin/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteDepartment(ctx context.Context, where model.DepartmentBoolExp) (*model.DepartmentMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.Department{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.Department
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
	return &model.DepartmentMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteDepartmentByPk(ctx context.Context, Id int64) (*model1.Department, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.Department
	tx := db.DB.Model(&model1.Department{})
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

func (r *mutationResolver) InsertDepartment(ctx context.Context, objects []*model.DepartmentInsertInput) (*model.DepartmentMutationResponse, error) {
	rs := make([]*model1.Department, 0)
	for _, object := range objects {
		v := &model1.Department{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.Department{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.DepartmentMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertDepartmentOne(ctx context.Context, object model.DepartmentInsertInput) (*model1.Department, error) {
	rs := &model1.Department{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.Department{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateDepartment(ctx context.Context, inc *model.DepartmentIncInput, set *model.DepartmentSetInput, where model.DepartmentBoolExp) (*model.DepartmentMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.Department{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.DepartmentMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.Department
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
	return &model.DepartmentMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateDepartmentByPk(ctx context.Context, inc *model.DepartmentIncInput, set *model.DepartmentSetInput, Id int64) (*model1.Department, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.Department{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.Department
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) Department(ctx context.Context, distinctOn []model.DepartmentSelectColumn, limit *int, offset *int, orderBy []*model.DepartmentOrderBy, where *model.DepartmentBoolExp) ([]*model1.Department, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.Department{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.Department
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) DepartmentAggregate(ctx context.Context, distinctOn []model.DepartmentSelectColumn, limit *int, offset *int, orderBy []*model.DepartmentOrderBy, where *model.DepartmentBoolExp) (*model.DepartmentAggregate, error) {
	var rs model.DepartmentAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.Department{})
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

func (r *queryResolver) DepartmentByPk(ctx context.Context, Id int64) (*model1.Department, error) {
	var rs model1.Department
	tx := db.DB.Model(&model1.Department{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}
