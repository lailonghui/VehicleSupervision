package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/device/graph/model"
	model1 "VehicleSupervision/internal/modules/device/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteTerminalBuyBill(ctx context.Context, where model.TerminalBuyBillBoolExp) (*model.TerminalBuyBillMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalBuyBill{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.TerminalBuyBill
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
	return &model.TerminalBuyBillMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteTerminalBuyBillByPk(ctx context.Context, Id int64) (*model1.TerminalBuyBill, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.TerminalBuyBill
	tx := db.DB.Model(&model1.TerminalBuyBill{})
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

func (r *mutationResolver) DeleteTerminalBuyBillByUnionPk(ctx context.Context, unionId string) (*model1.TerminalBuyBill, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.TerminalBuyBill
	tx := db.DB.Model(&model1.TerminalBuyBill{})
	if len(preloads) > 0 {
		// 如果请求的字段不为空，则先查询一遍数据库
		tx = tx.Select(preloads).Where(rs.UnionPrimaryColumnName()+" = ?", unionId).First(&rs)
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

func (r *mutationResolver) InsertTerminalBuyBill(ctx context.Context, objects []*model.TerminalBuyBillInsertInput) (*model.TerminalBuyBillMutationResponse, error) {
	rs := make([]*model1.TerminalBuyBill, 0)
	for _, object := range objects {
		v := &model1.TerminalBuyBill{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.TerminalBuyBill{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.TerminalBuyBillMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertTerminalBuyBillOne(ctx context.Context, object model.TerminalBuyBillInsertInput) (*model1.TerminalBuyBill, error) {
	rs := &model1.TerminalBuyBill{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.TerminalBuyBill{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateTerminalBuyBill(ctx context.Context, inc *model.TerminalBuyBillIncInput, set *model.TerminalBuyBillSetInput, where model.TerminalBuyBillBoolExp) (*model.TerminalBuyBillMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalBuyBill{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.TerminalBuyBillMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.TerminalBuyBill
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
	return &model.TerminalBuyBillMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateTerminalBuyBillByPk(ctx context.Context, inc *model.TerminalBuyBillIncInput, set *model.TerminalBuyBillSetInput, Id int64) (*model1.TerminalBuyBill, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.TerminalBuyBill{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.TerminalBuyBill
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *mutationResolver) UpdateTerminalBuyBillByUnionPk(ctx context.Context, inc *model.TerminalBuyBillIncInput, set *model.TerminalBuyBillSetInput, unionId string) (*model1.TerminalBuyBill, error) {
	var rs model1.TerminalBuyBill
	tx := db.DB.Where(rs.UnionPrimaryColumnName()+" = ?", unionId)
	qt := util.NewQueryTranslator(tx, &model1.TerminalBuyBill{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}

	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) TerminalBuyBill(ctx context.Context, distinctOn []model.TerminalBuyBillSelectColumn, limit *int, offset *int, orderBy []*model.TerminalBuyBillOrderBy, where *model.TerminalBuyBillBoolExp) ([]*model1.TerminalBuyBill, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalBuyBill{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.TerminalBuyBill
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) TerminalBuyBillAggregate(ctx context.Context, distinctOn []model.TerminalBuyBillSelectColumn, limit *int, offset *int, orderBy []*model.TerminalBuyBillOrderBy, where *model.TerminalBuyBillBoolExp) (*model.TerminalBuyBillAggregate, error) {
	var rs model.TerminalBuyBillAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.TerminalBuyBill{})
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

func (r *queryResolver) TerminalBuyBillByPk(ctx context.Context, Id int64) (*model1.TerminalBuyBill, error) {
	var rs model1.TerminalBuyBill
	tx := db.DB.Model(&model1.TerminalBuyBill{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}

func (r *queryResolver) TerminalBuyBillByUnionPk(ctx context.Context, unionId string) (*model1.TerminalBuyBill, error) {
	var rs model1.TerminalBuyBill
	tx := db.DB.Model(&model1.TerminalBuyBill{}).Select(util.GetTopPreloads(ctx)).Where(rs.UnionPrimaryColumnName()+" = ?", unionId).First(&rs)

	err := tx.Error
	return &rs, err
}
