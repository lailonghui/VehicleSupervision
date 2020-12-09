package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/db"
	model1 "VehicleSupervision/internal/modules/blacklist/record/model"
	"VehicleSupervision/internal/modules/blacklist/record/mutation/graph/generated"
	"VehicleSupervision/internal/modules/blacklist/record/mutation/graph/model"
	"VehicleSupervision/pkg/graphql/util"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteBlacklistOperationRecord(ctx context.Context, where model.BlacklistOperationRecordBoolExp) (*model.BlacklistOperationRecordMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.BlacklistOperationRecord{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.BlacklistOperationRecord
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
	return &model.BlacklistOperationRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteBlacklistOperationRecordByPk(ctx context.Context, id int64) (*model1.BlacklistOperationRecord, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.BlacklistOperationRecord
	tx := db.DB.Model(&model1.BlacklistOperationRecord{})
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

func (r *mutationResolver) InsertBlacklistOperationRecord(ctx context.Context, objects []*model.BlacklistOperationRecordInsertInput, onConflict *model.BlacklistOperationRecordOnConflict) (*model.BlacklistOperationRecordMutationResponse, error) {
	rs := r.batchInsertParamConvert(objects)
	tx := db.DB.Model(&model1.BlacklistOperationRecord{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.BlacklistOperationRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertBlacklistOperationRecordOne(ctx context.Context, object model.BlacklistOperationRecordInsertInput, onConflict *model.BlacklistOperationRecordOnConflict) (*model1.BlacklistOperationRecord, error) {
	rs := r.insertParamConvert(&object)
	tx := db.DB.Model(&model1.BlacklistOperationRecord{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateBlacklistOperationRecord(ctx context.Context, inc *model.BlacklistOperationRecordIncInput, set *model.BlacklistOperationRecordSetInput, where model.BlacklistOperationRecordBoolExp) (*model.BlacklistOperationRecordMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.BlacklistOperationRecord{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.BlacklistOperationRecordMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.BlacklistOperationRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateBlacklistOperationRecordByPk(ctx context.Context, inc *model.BlacklistOperationRecordIncInput, set *model.BlacklistOperationRecordSetInput, pkColumns model.BlacklistOperationRecordPkColumnsInput) (*model1.BlacklistOperationRecord, error) {
	tx := db.DB.Where("id = ?", pkColumns.ID)
	qt := util.NewQueryTranslator(tx, &model1.BlacklistOperationRecord{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.BlacklistOperationRecord
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) T(ctx context.Context) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
