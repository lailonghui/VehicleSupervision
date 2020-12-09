package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/db"
	model1 "VehicleSupervision/internal/modules/device/model"
	"VehicleSupervision/internal/modules/device/mutation/graph/generated"
	"VehicleSupervision/internal/modules/device/mutation/graph/model"
	"VehicleSupervision/pkg/graphql/util"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteSimCardFlow(ctx context.Context, where model.SimCardFlowBoolExp) (*model.SimCardFlowMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.SimCardFlow{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.SimCardFlow
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
	return &model.SimCardFlowMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteSimCardFlowByPk(ctx context.Context, id int64) (*model1.SimCardFlow, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.SimCardFlow
	tx := db.DB.Model(&model1.SimCardFlow{})
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

func (r *mutationResolver) InsertSimCardFlow(ctx context.Context, objects []*model.SimCardFlowInsertInput, onConflict *model.SimCardFlowOnConflict) (*model.SimCardFlowMutationResponse, error) {
	rs := simCardFlowInsertInputBatchConvert(objects)
	tx := db.DB.Model(&model1.SimCardFlow{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.SimCardFlowMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertSimCardFlowOne(ctx context.Context, object model.SimCardFlowInsertInput, onConflict *model.SimCardFlowOnConflict) (*model1.SimCardFlow, error) {
	rs := simCardFlowInsertInputConvert(&object)
	tx := db.DB.Model(&model1.SimCardFlow{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateSimCardFlow(ctx context.Context, inc *model.SimCardFlowIncInput, set *model.SimCardFlowSetInput, where model.SimCardFlowBoolExp) (*model.SimCardFlowMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.SimCardFlow{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.SimCardFlowMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.SimCardFlowMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateSimCardFlowByPk(ctx context.Context, inc *model.SimCardFlowIncInput, set *model.SimCardFlowSetInput, pkColumns model.SimCardFlowPkColumnsInput) (*model1.SimCardFlow, error) {
	tx := db.DB.Where("id = ?", pkColumns.ID)
	qt := util.NewQueryTranslator(tx, &model1.SimCardFlow{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.SimCardFlow
	tx = tx.First(&rs)
	return &rs, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
