package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/area/mutation/graph/generated"
	"VehicleSupervision/internal/modules/area/mutation/graph/model"
	"VehicleSupervision/pkg/graphql/util"
	"VehicleSupervision/pkg/xid"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteCity(ctx context.Context, where model.CityBoolExp) (*model.CityMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model.City{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model.City
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
	return &model.CityMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, tx.Error
}

func (r *mutationResolver) DeleteCityByPk(ctx context.Context, id int64) (*model.City, error) {
	var rs = model.City{}
	tx := db.DB.Model(&model.City{}).Find(&rs, id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	tx = db.DB.Delete(id)
	return &rs, tx.Error
}

func (r *mutationResolver) DeleteDistrict(ctx context.Context, where model.DistrictBoolExp) (*model.DistrictMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model.District{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model.District
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
	return &model.DistrictMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, tx.Error
}

func (r *mutationResolver) DeleteDistrictByPk(ctx context.Context, id int64) (*model.District, error) {
	var rs = model.District{}
	tx := db.DB.Model(&model.District{}).Find(&rs, id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	tx = db.DB.Delete(id)
	return &rs, tx.Error
}

func (r *mutationResolver) DeleteProvince(ctx context.Context, where model.ProvinceBoolExp) (*model.ProvinceMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model.Province{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model.Province
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
	return &model.ProvinceMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, tx.Error
}

func (r *mutationResolver) DeleteProvinceByPk(ctx context.Context, id int64) (*model.Province, error) {
	var rs = model.Province{}
	tx := db.DB.Model(&model.Province{}).Find(&rs, id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	tx = db.DB.Delete(id)
	return &rs, tx.Error
}

func (r *mutationResolver) InsertCity(ctx context.Context, objects []*model.CityInsertInput, onConflict *model.CityOnConflict) (*model.CityMutationResponse, error) {
	for _, input := range objects {
		xidStr := xid.GetXid()
		input.CityID = &xidStr
		input.ID = nil
	}
	tx := db.DB.Model(&model.City{}).Save(objects)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.CityMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) InsertCityOne(ctx context.Context, object model.CityInsertInput, onConflict *model.CityOnConflict) (*model.City, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertDistrict(ctx context.Context, objects []*model.DistrictInsertInput, onConflict *model.DistrictOnConflict) (*model.DistrictMutationResponse, error) {
	for _, input := range objects {
		xidStr := xid.GetXid()
		input.DistrictID = &xidStr
		input.ID = nil
	}
	tx := db.DB.Model(&model.City{}).Save(objects)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.DistrictMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) InsertDistrictOne(ctx context.Context, object model.DistrictInsertInput, onConflict *model.DistrictOnConflict) (*model.District, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertProvince(ctx context.Context, objects []*model.ProvinceInsertInput, onConflict *model.ProvinceOnConflict) (*model.ProvinceMutationResponse, error) {
	for _, input := range objects {
		xidStr := xid.GetXid()
		input.ProvinceID = &xidStr
		input.ID = nil
	}
	tx := db.DB.Model(&model.City{}).Save(objects)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.ProvinceMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) InsertProvinceOne(ctx context.Context, object model.ProvinceInsertInput, onConflict *model.ProvinceOnConflict) (*model.Province, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateCity(ctx context.Context, inc *model.CityIncInput, set *model.CitySetInput, where model.CityBoolExp) (*model.CityMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model.City{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.CityMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.CityMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateCityByPk(ctx context.Context, inc *model.CityIncInput, set *model.CitySetInput, pkColumns model.CityPkColumnsInput) (*model.City, error) {
	tx := db.DB.Where("id = ?", pkColumns.ID)
	qt := util.NewQueryTranslator(tx, &model.City{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model.City
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *mutationResolver) UpdateDistrict(ctx context.Context, inc *model.DistrictIncInput, set *model.DistrictSetInput, where model.DistrictBoolExp) (*model.DistrictMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model.District{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.DistrictMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.DistrictMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateDistrictByPk(ctx context.Context, inc *model.DistrictIncInput, set *model.DistrictSetInput, pkColumns model.DistrictPkColumnsInput) (*model.District, error) {
	tx := db.DB.Where("id = ?", pkColumns.ID)
	qt := util.NewQueryTranslator(tx, &model.District{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model.District
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *mutationResolver) UpdateProvince(ctx context.Context, inc *model.ProvinceIncInput, set *model.ProvinceSetInput, where model.ProvinceBoolExp) (*model.ProvinceMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model.Province{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.ProvinceMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.ProvinceMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateProvinceByPk(ctx context.Context, inc *model.ProvinceIncInput, set *model.ProvinceSetInput, pkColumns model.ProvincePkColumnsInput) (*model.Province, error) {
	tx := db.DB.Where("id = ?", pkColumns.ID)
	qt := util.NewQueryTranslator(tx, &model.Province{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model.Province
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) Bug(ctx context.Context) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
