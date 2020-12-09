package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/adas/query/graph/generated"
	"VehicleSupervision/internal/modules/adas/query/graph/model"
	"VehicleSupervision/pkg/graphql/util"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (r *queryResolver) AdasAttachment(ctx context.Context, distinctOn []model.AdasAttachmentSelectColumn, limit *int, offset *int, orderBy []*model.AdasAttachmentOrderBy, where *model.AdasAttachmentBoolExp) ([]*model.AdasAttachment, error) {
	qt := util.NewQueryTranslator(db.DB, &model.AdasAttachment{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model.AdasAttachment
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) AdasAttachmentAggregate(ctx context.Context, distinctOn []model.AdasAttachmentSelectColumn, limit *int, offset *int, orderBy []*model.AdasAttachmentOrderBy, where *model.AdasAttachmentBoolExp) (*model.AdasAttachmentAggregate, error) {
	var rs model.AdasAttachmentAggregate

	qt := util.NewQueryTranslator(db.DB, &model.AdasAttachment{})
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

func (r *queryResolver) AdasAttachmentByPk(ctx context.Context, id int64) (*model.AdasAttachment, error) {
	var rs model.AdasAttachment
	tx := db.DB.Model(&model.AdasAttachment{}).First(&rs, id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}

func (r *queryResolver) AdasData(ctx context.Context, distinctOn []model.AdasDataSelectColumn, limit *int, offset *int, orderBy []*model.AdasDataOrderBy, where *model.AdasDataBoolExp) ([]*model.AdasData, error) {
	qt := util.NewQueryTranslator(db.DB, &model.AdasData{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model.AdasData
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) AdasDataAggregate(ctx context.Context, distinctOn []model.AdasDataSelectColumn, limit *int, offset *int, orderBy []*model.AdasDataOrderBy, where *model.AdasDataBoolExp) (*model.AdasDataAggregate, error) {
	var rs model.AdasDataAggregate

	qt := util.NewQueryTranslator(db.DB, &model.AdasData{})
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

func (r *queryResolver) AdasDataByPk(ctx context.Context, id int64) (*model.AdasData, error) {
	var rs model.AdasData
	tx := db.DB.Model(&model.AdasData{}).First(&rs, id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
