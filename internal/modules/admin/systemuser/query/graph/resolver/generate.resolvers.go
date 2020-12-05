package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/db"
	model1 "VehicleSupervision/internal/modules/admin/systemuser/model"
	"VehicleSupervision/internal/modules/admin/systemuser/query/graph/generated"
	"VehicleSupervision/internal/modules/admin/systemuser/query/graph/model"
	"VehicleSupervision/pkg/graphql/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *queryResolver) SystemUser(ctx context.Context, distinctOn []model.SystemUserSelectColumn, limit *int, offset *int, orderBy []*model.SystemUserOrderBy, where *model.SystemUserBoolExp) ([]*model1.SystemUser, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.SystemUser{})
	/*qt.DistinctOn(distinctOn).
	Limit(limit).
	Offset(offset).
	OrderBy(orderBy).
	Where(where)*/
	// 执行翻译
	tx := qt.DoTranslate()
	var rs []*model1.SystemUser
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) SystemUserAggregate(ctx context.Context, distinctOn []model.SystemUserSelectColumn, limit *int, offset *int, orderBy []*model.SystemUserOrderBy, where *model.SystemUserBoolExp) (*model.SystemUserAggregate, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.SystemUser{})
	qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where)
	// 执行翻译
	//tx := qt.DoTranslate()

	return nil, nil
}

func (r *queryResolver) SystemUserByPk(ctx context.Context, id int64) (*model1.SystemUser, error) {
	var rs model1.SystemUser
	tx := db.DB.Model(&model1.SystemUser{}).First(&rs, id)
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
