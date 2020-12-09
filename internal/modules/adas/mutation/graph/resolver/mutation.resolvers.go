package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/adas/mutation/graph/generated"
	"VehicleSupervision/internal/modules/adas/mutation/graph/model"
	"VehicleSupervision/pkg/graphql/util"
	"VehicleSupervision/pkg/xid"
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteAdasAttachment(ctx context.Context, where model.AdasAttachmentBoolExp) (*model.AdasAttachmentMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model.AdasAttachment{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model.AdasAttachment
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
	return &model.AdasAttachmentMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, tx.Error
}

func (r *mutationResolver) DeleteAdasAttachmentByPk(ctx context.Context, id int64) (*model.AdasAttachment, error) {
	var rs = model.AdasAttachment{}
	tx := db.DB.Model(&model.AdasAttachment{}).Find(&rs, id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	tx = db.DB.Delete(id)
	return &rs, tx.Error
}

func (r *mutationResolver) DeleteAdasData(ctx context.Context, where model.AdasDataBoolExp) (*model.AdasDataMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model.AdasData{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model.AdasData
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
	return &model.AdasDataMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, tx.Error
}

func (r *mutationResolver) DeleteAdasDataByPk(ctx context.Context, id int64) (*model.AdasData, error) {
	var rs = model.AdasData{}
	tx := db.DB.Model(&model.AdasData{}).Find(&rs, id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	tx = db.DB.Delete(id)
	return &rs, tx.Error
}

func (r *mutationResolver) InsertAdasAttachment(ctx context.Context, objects []*model.AdasAttachmentInsertInput, onConflict *model.AdasAttachmentOnConflict) (*model.AdasAttachmentMutationResponse, error) {
	for _, input := range objects {
		xidStr := xid.GetXid()
		input.AttachmentID = &xidStr
		input.ID = nil
	}
	tx := db.DB.Model(&model.AdasAttachment{}).Save(objects)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.AdasAttachmentMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) InsertAdasAttachmentOne(ctx context.Context, object model.AdasAttachmentInsertInput, onConflict *model.AdasAttachmentOnConflict) (*model.AdasAttachment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertAdasData(ctx context.Context, objects []*model.AdasDataInsertInput, onConflict *model.AdasDataOnConflict) (*model.AdasDataMutationResponse, error) {
	for _, input := range objects {
		xidStr := xid.GetXid()
		input.AlarmNo = &xidStr
		input.ID = nil
	}
	tx := db.DB.Model(&model.AdasAttachment{}).Save(objects)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.AdasDataMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) InsertAdasDataOne(ctx context.Context, object model.AdasDataInsertInput, onConflict *model.AdasDataOnConflict) (*model.AdasData, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateAdasAttachment(ctx context.Context, inc *model.AdasAttachmentIncInput, set *model.AdasAttachmentSetInput, where model.AdasAttachmentBoolExp) (*model.AdasAttachmentMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model.AdasAttachment{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.AdasAttachmentMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.AdasAttachmentMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateAdasAttachmentByPk(ctx context.Context, inc *model.AdasAttachmentIncInput, set *model.AdasAttachmentSetInput, pkColumns model.AdasAttachmentPkColumnsInput) (*model.AdasAttachment, error) {
	tx := db.DB.Where("id = ?", pkColumns.ID)
	qt := util.NewQueryTranslator(tx, &model.AdasAttachment{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model.AdasAttachment
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *mutationResolver) UpdateAdasData(ctx context.Context, inc *model.AdasDataIncInput, set *model.AdasDataSetInput, where model.AdasDataBoolExp) (*model.AdasDataMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model.AdasAttachment{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.AdasDataMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.AdasDataMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateAdasDataByPk(ctx context.Context, inc *model.AdasDataIncInput, set *model.AdasDataSetInput, pkColumns model.AdasDataPkColumnsInput) (*model.AdasData, error) {
	tx := db.DB.Where("id = ?", pkColumns.ID)
	qt := util.NewQueryTranslator(tx, &model.AdasData{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model.AdasData
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
