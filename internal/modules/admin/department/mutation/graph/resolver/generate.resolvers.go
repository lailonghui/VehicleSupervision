package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/db"
	model1 "VehicleSupervision/internal/modules/admin/department/model"
	"VehicleSupervision/internal/modules/admin/department/mutation/graph/generated"
	"VehicleSupervision/internal/modules/admin/department/mutation/graph/model"
	"VehicleSupervision/pkg/graphql/util"
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

func (r *mutationResolver) DeleteDepartmentByPk(ctx context.Context, id int64) (*model1.Department, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.Department
	tx := db.DB.Model(&model1.Department{})
	if len(preloads) > 0 {
		// 如果请求的字段不为空，则先查询一遍数据库
		tx = tx.Select(preloads).Where("id = ?", id).First(&rs)
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

func (r *mutationResolver) InsertDepartment(ctx context.Context, objects []*model.DepartmentInsertInput, onConflict *model.DepartmentOnConflict) (*model.DepartmentMutationResponse, error) {
	rs := r.batchInsertParamConvert(objects)
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

func (r *mutationResolver) InsertDepartmentOne(ctx context.Context, object model.DepartmentInsertInput, onConflict *model.DepartmentOnConflict) (*model1.Department, error) {
	rs := r.insertParamConvert(&object)
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
	return &model.DepartmentMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateDepartmentByPk(ctx context.Context, inc *model.DepartmentIncInput, set *model.DepartmentSetInput, pkColumns model.DepartmentPkColumnsInput) (*model1.Department, error) {
	tx := db.DB.Where("id = ?", pkColumns.ID)
	qt := util.NewQueryTranslator(tx, &model1.Department{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.Department
	tx = tx.First(&rs)
	return &rs, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }