package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/dataloader"
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/device/graph/generated"
	"VehicleSupervision/internal/modules/device/graph/model"
	model1 "VehicleSupervision/internal/modules/device/model"
	"VehicleSupervision/pkg/graphql/util"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteSimCard(ctx context.Context, where model.SimCardBoolExp) (*model.SimCardMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.SimCard{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.SimCard
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
	return &model.SimCardMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteSimCardByPk(ctx context.Context, id int64) (*model1.SimCard, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.SimCard
	tx := db.DB.Model(&model1.SimCard{})
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

func (r *mutationResolver) InsertSimCard(ctx context.Context, objects []*model.SimCardInsertInput, onConflict *model.SimCardOnConflict) (*model.SimCardMutationResponse, error) {
	rs := simCardInsertInputBatchConvert(objects)
	tx := db.DB.Model(&model1.SimCard{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.SimCardMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertSimCardOne(ctx context.Context, object model.SimCardInsertInput, onConflict *model.SimCardOnConflict) (*model1.SimCard, error) {
	rs := simCardInsertInputConvert(&object)
	tx := db.DB.Model(&model1.SimCard{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateSimCard(ctx context.Context, inc *model.SimCardIncInput, set *model.SimCardSetInput, where model.SimCardBoolExp) (*model.SimCardMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.SimCard{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.SimCardMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.SimCardMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateSimCardByPk(ctx context.Context, inc *model.SimCardIncInput, set *model.SimCardSetInput, pkColumns model.SimCardPkColumnsInput) (*model1.SimCard, error) {
	tx := db.DB.Where("id = ?", pkColumns.ID)
	qt := util.NewQueryTranslator(tx, &model1.SimCard{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.SimCard
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) SimCard(ctx context.Context, distinctOn []model.SimCardSelectColumn, limit *int, offset *int, orderBy []*model.SimCardOrderBy, where *model.SimCardBoolExp) ([]*model1.SimCard, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.SimCard{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.SimCard
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) SimCardAggregate(ctx context.Context, distinctOn []model.SimCardSelectColumn, limit *int, offset *int, orderBy []*model.SimCardOrderBy, where *model.SimCardBoolExp) (*model.SimCardAggregate, error) {
	var rs model.SimCardAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.SimCard{})
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

func (r *queryResolver) SimCardByPk(ctx context.Context, id int64) (*model1.SimCard, error) {
	var rs model1.SimCard
	tx := db.DB.Model(&model1.SimCard{}).First(&rs, id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}

func (r *simCardResolver) TerminalID(ctx context.Context, obj *model1.SimCard) (*model1.Terminal, error) {
	if obj.TerminalID == nil || *obj.TerminalID == "" {
		return nil, nil
	}
	return dataloader.GetLoaders(ctx).TerminalLoader.Load(*obj.TerminalID)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// SimCard returns generated.SimCardResolver implementation.
func (r *Resolver) SimCard() generated.SimCardResolver { return &simCardResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type simCardResolver struct{ *Resolver }