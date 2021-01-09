package resolver

import (
	"VehicleSupervision/internal/cache"
	"VehicleSupervision/internal/dataloader"
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/admin/graph/generated"
	"VehicleSupervision/internal/modules/admin/graph/model"
	model1 "VehicleSupervision/internal/modules/admin/model"
	"VehicleSupervision/internal/server/middle"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"strconv"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteDepartment(ctx context.Context, where model.DepartmentBoolExp) (*model.DepartmentMutationResponse, error) {
	var rs []*model1.Department
	var m = &model1.Department{}
	qt := util.NewQueryTranslator(db.DB, m)
	tx := qt.Where(where).Finish()
	// 获取ID,联系主键
	tx = tx.Select(m.PrimaryColumnName(), m.UnionPrimaryColumnName())
	tx = tx.Find(&rs)
	// 删除
	amount := len(rs)
	if amount == 0 {
		return &model.DepartmentMutationResponse{
			AffectedRows: 0,
			Returning:    nil,
		}, nil
	}
	var idList []int64 = make([]int64, 0, amount)
	var unionIdList []string = make([]string, 0, amount)
	for i := 0; i < amount; i++ {
		idList[i] = rs[i].GetPrimary()
		unionIdList[i] = rs[i].GetUnionPrimary()
	}
	tx = db.DB.Model(m).Delete(idList)
	if err := tx.Error; err != nil {
		return nil, err
	}
	cacheAspect, cacheErr := cache.GetGqlCacheAspect(m.TableName())
	if cacheErr == nil {
		rsLen := len(rs)
		var ids = make([]string, 0, rsLen)
		var unionIds = make([]string, 0, rsLen)
		for i := 0; i < rsLen; i++ {
			ids[i] = util2.ToStr(rs[i].GetPrimary())
			unionIds[i] = rs[i].GetUnionPrimary()
		}
		_ = cacheAspect.OnListRemove(ctx, ids, unionIds)
	}
	return &model.DepartmentMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    nil,
	}, nil
}

func (r *mutationResolver) DeleteDepartmentByPk(ctx context.Context, id int64) (*model1.Department, error) {
	var rs model1.Department
	m := &model1.Department{}
	tx := db.DB.Model(m)
	// 如果请求的字段不为空，则先查询一遍数据库
	tx = tx.Where(rs.PrimaryColumnName()+" = ?", id).First(&rs)
	// 如果查询结果含有错误，则返回错误
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	// 删除
	tx = tx.Delete(nil)
	if err := tx.Error; err != nil {
		return nil, err
	}
	cacheAspect, cacheErr := cache.GetGqlCacheAspect(m.TableName())
	if cacheErr == nil {
		_ = cacheAspect.OnPkRemove(ctx, util2.ToStr(rs.GetPrimary()))
		_ = cacheAspect.OnUnionPkRemove(ctx, rs.GetUnionPrimary())
	}
	return &rs, nil
}

func (r *mutationResolver) DeleteDepartmentByUnionPk(ctx context.Context, departmentID string) (*model1.Department, error) {
	var rs model1.Department
	m := &model1.Department{}
	tx := db.DB.Model(&model1.Department{})
	// 如果请求的字段不为空，则先查询一遍数据库
	tx = tx.Where(rs.UnionPrimaryColumnName()+" = ?", departmentID).First(&rs)
	// 如果查询结果含有错误，则返回错误
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	// 删除
	tx = tx.Delete(nil)
	if err := tx.Error; err != nil {
		return nil, err
	}
	cacheAspect, cacheErr := cache.GetGqlCacheAspect(m.TableName())
	if cacheErr == nil {
		_ = cacheAspect.OnPkRemove(ctx, util2.ToStr(rs.GetPrimary()))
		_ = cacheAspect.OnUnionPkRemove(ctx, rs.GetUnionPrimary())
	}
	return &rs, nil
}

func (r *mutationResolver) InsertDepartment(ctx context.Context, objects []*model.DepartmentInsertInput) (*model.DepartmentMutationResponse, error) {
	rs := make([]*model1.Department, 0)
	m := &model1.Department{}
	for _, object := range objects {
		v := &model1.Department{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(m).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	cacheAspect, cacheErr := cache.GetGqlCacheAspect(m.TableName())
	if cacheErr == nil {
		_ = cacheAspect.OnInsert(ctx)
	}
	return &model.DepartmentMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertDepartmentOne(ctx context.Context, objects model.DepartmentInsertInput) (*model1.Department, error) {
	rs := &model1.Department{}
	m := &model1.Department{}
	util2.StructAssign(rs, &objects)
	tx := db.DB.Model(m).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	cacheAspect, cacheErr := cache.GetGqlCacheAspect(m.TableName())
	if cacheErr == nil {
		_ = cacheAspect.OnInsert(ctx)
	}
	return rs, nil
}

func (r *mutationResolver) UpdateDepartment(ctx context.Context, inc *model.DepartmentIncInput, set *model.DepartmentSetInput, where model.DepartmentBoolExp) (*model.DepartmentMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.Department{})
	var m = &model1.Department{}
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
	var rs []*model1.Department
	// 如果请求的字段不为空，则先查询一遍数据库
	tx = tx.Find(&rs)
	// 如果查询结果含有错误，则返回错误
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	cacheAspect, cacheErr := cache.GetGqlCacheAspect(m.TableName())
	if cacheErr == nil {
		rsLen := len(rs)
		var ids = make([]string, 0, rsLen)
		var unionIds = make([]string, 0, rsLen)
		for i := 0; i < rsLen; i++ {
			ids[i] = util2.ToStr(rs[i].GetPrimary())
			unionIds[i] = rs[i].GetUnionPrimary()
		}
		_ = cacheAspect.OnListUpdate(ctx, ids, unionIds)
	}
	return &model.DepartmentMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateDepartmentByPk(ctx context.Context, inc *model.DepartmentIncInput, set *model.DepartmentSetInput, id int64) (*model1.Department, error) {
	var rs model1.Department
	var m = &model1.Department{}
	tx := db.DB.Where(rs.PrimaryColumnName()+" = ?", id)
	qt := util.NewQueryTranslator(tx, &model1.Department{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	cacheAspect, cacheErr := cache.GetGqlCacheAspect(m.TableName())
	if cacheErr == nil {
		_ = cacheAspect.OnPkRemove(ctx, util2.ToStr(rs.GetPrimary()))
		_ = cacheAspect.OnUnionPkRemove(ctx, rs.GetUnionPrimary())
	}
	return &rs, nil
}

func (r *mutationResolver) UpdateDepartmentByUnionPk(ctx context.Context, inc *model.DepartmentIncInput, set *model.DepartmentSetInput, departmentID string) (*model1.Department, error) {
	var rs model1.Department
	var m = &model1.Department{}
	tx := db.DB.Where(rs.UnionPrimaryColumnName()+" = ?", departmentID)
	qt := util.NewQueryTranslator(tx, &model1.Department{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}

	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	cacheAspect, cacheErr := cache.GetGqlCacheAspect(m.TableName())
	if cacheErr == nil {
		_ = cacheAspect.OnPkRemove(ctx, util2.ToStr(rs.GetPrimary()))
		_ = cacheAspect.OnUnionPkRemove(ctx, rs.GetUnionPrimary())
	}
	return &rs, nil
}

func (r *queryResolver) Department(ctx context.Context, distinctOn []model.DepartmentSelectColumn, limit *int, offset *int, orderBy []*model.DepartmentOrderBy, where *model.DepartmentBoolExp) ([]*model1.Department, error) {
	var rs []*model1.Department
	var m = &model1.Department{}
	if middle.IsEnableGqlCache(ctx) {
		cacheAspect, cacheErr := cache.GetGqlCacheAspect(m.TableName())
		cacheKey := graphql.GetOperationContext(ctx).RawQuery
		if cacheErr == nil {
			exist, err := cacheAspect.OnListQuery(ctx, cacheKey, &rs)
			if err != nil {
				return nil, err
			}
			if exist {
				return rs, err
			}
		}
		qt := util.NewQueryTranslator(db.DB, m)
		tx := qt.DistinctOn(distinctOn).
			Limit(limit).
			Offset(offset).
			OrderBy(orderBy).
			Where(where).
			Finish()
		tx = tx.Find(&rs)
		err := tx.Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return rs, nil
			}
			return rs, err
		}
		if cacheErr == nil {
			_ = cacheAspect.SetListQueryCache(ctx, cacheKey, rs)
		}
		return rs, err
	}
	qt := util.NewQueryTranslator(db.DB, m)
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	tx = tx.Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) DepartmentAggregate(ctx context.Context, distinctOn []model.DepartmentSelectColumn, limit *int, offset *int, orderBy []*model.DepartmentOrderBy, where *model.DepartmentBoolExp) (*model.DepartmentAggregate, error) {
	var rs model.DepartmentAggregate
	var m = &model1.Department{}
	if middle.IsEnableGqlCache(ctx) {
		cacheAspect, cacheErr := cache.GetGqlCacheAspect(m.TableName())
		cacheKey := graphql.GetOperationContext(ctx).RawQuery
		if cacheErr == nil {
			exist, err := cacheAspect.OnArrgegateQuery(ctx, cacheKey, &rs)
			if err != nil {
				return nil, err
			}
			if exist {
				return &rs, err
			}
		}
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
		if cacheErr == nil {
			_ = cacheAspect.SetArrgegateQueryCache(ctx, cacheKey, rs)
		}
		err = tx.Error
		return &rs, err
	}
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

func (r *queryResolver) DepartmentByPk(ctx context.Context, id int64) (*model1.Department, error) {

	var m = &model1.Department{}
	var rs model1.Department
	if middle.IsEnableGqlCache(ctx) {
		cacheAspect, cacheErr := cache.GetGqlCacheAspect(m.TableName())
		cacheKey := strconv.FormatInt(id, 10)
		if cacheErr == nil {
			exist, err := cacheAspect.OnPkQuery(ctx, cacheKey, &rs)
			if err != nil {
				return nil, err
			}
			if exist {
				return &rs, err
			}
		}
		tx := db.DB.Model(m).First(&rs, id)
		err := tx.Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				if cacheErr == nil {
					_ = cacheAspect.SetNotExistPkQueryCache(ctx, cacheKey, nil)
				}
				return &rs, nil
			}
			return &rs, err
		}
		if cacheErr == nil {
			_ = cacheAspect.SetPkQueryCache(ctx, cacheKey, rs)
		}
		return &rs, nil
	}
	tx := db.DB.Model(m).First(&rs, id)
	err := tx.Error
	if err != nil {
		return nil, err
	}
	return &rs, nil

}

func (r *queryResolver) DepartmentByUnionPk(ctx context.Context, departmentID string) (*model1.Department, error) {
	var rs model1.Department
	var m = &model1.Department{}
	if middle.IsEnableGqlCache(ctx) {
		cacheAspect, cacheErr := cache.GetGqlCacheAspect(rs.TableName())
		cacheKey := departmentID
		if cacheErr == nil {
			exist, err := cacheAspect.OnUnionPkQuery(ctx, cacheKey, &rs)
			if err != nil {
				return nil, err
			}
			if exist {
				return &rs, err
			}
		}
		tx := db.DB.Model(m).Where(rs.UnionPrimaryColumnName()+" = ?", departmentID).First(&rs)
		err := tx.Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				if cacheErr == nil {
					_ = cacheAspect.SetNotExistUnionPkQueryCache(ctx, cacheKey, model1.Department{})
				}
				return &rs, nil
			}
			return nil, err
		}
		if cacheErr == nil {
			_ = cacheAspect.SetUnionPkQueryCache(ctx, cacheKey, rs)
		}
		return &rs, nil
	}
	tx := db.DB.Model(m).Where(rs.UnionPrimaryColumnName()+" = ?", departmentID).First(&rs)
	err := tx.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &rs, nil
		}
		return nil, err
	}
	return &rs, nil
}

func (r *departmentResolver) Enterprise(ctx context.Context, obj *model1.Department) (*model1.Enterprise, error) {

	var loader model1.EnterpriseUnionPkLoader
	i := dataloader.GetLoaders(ctx).GetLoader(&loader)
	switch t := i.(type) {
	case *model1.EnterpriseUnionPkLoader:
		return t.Load(obj.EnterpriseID)
	default:
		panic("can not get dataloader")
	}
}

func (r *departmentResolver) SuperiorDepartment(ctx context.Context, obj *model1.Department) (*model1.Department, error) {
	if obj.SuperiorDepartmentID == nil {
		return nil, nil
	}
	var loader model1.DepartmentUnionPkLoader
	i := dataloader.GetLoaders(ctx).GetLoader(&loader)
	switch t := i.(type) {
	case *model1.DepartmentUnionPkLoader:
		return t.Load(*obj.SuperiorDepartmentID)
	default:
		panic("can not get dataloader")
	}
}

// Department returns generated.DepartmentResolver implementation.
func (r *Resolver) Department() generated.DepartmentResolver { return &departmentResolver{r} }

type departmentResolver struct{ *Resolver }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
