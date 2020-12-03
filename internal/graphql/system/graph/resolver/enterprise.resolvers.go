package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"system-manage/internal/db"
	"system-manage/internal/graphql/system/graph/model"

	"github.com/dgryski/trifles/uuid"
)

func (r *mutationResolver) CreateEnterprise(ctx context.Context, input model.CreateEnterpriseParam) (*model.Enterprise, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateEnterprise(ctx context.Context, input model.UpdateEnterpriseParam) (*model.Enterprise, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RemoveEnterprise(ctx context.Context, input model.RemoveEnterpriseParam) (*model.Enterprise, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Enterprise(ctx context.Context, query model.EnterpriseQuery) (*model.Enterprise, error) {
	return &model.Enterprise{EnterpriseID: uuid.UUIDv4()}, nil
}

func (r *queryResolver) EnterpriseList(ctx context.Context, query *model.EnterpriseListQuery) (*model.EnterpriseConnection, error) {
	tx := db.DB.Model(&model.Enterprise{})
	if query.DistrictID != nil {
		if query.DistrictID.Eq != nil {
			tx.Where("district_id = ?", query.DistrictID.Eq)
		}
		if query.DistrictID.Neq != nil {
			tx.Where("district_id != ?", query.DistrictID.Neq)
		}
		if query.DistrictID.In != nil {
			tx.Where("district_id != ?", query.DistrictID.Neq)
		}
	}
	var rs []*model.Enterprise
	tx.Find(rs)

	return &model.EnterpriseConnection{
		PageInfo: &model.PageInfo{Total: 100},
		Nodes:    rs,
	}, nil
}

func (r *queryResolver) EnterpriseDistrictStatistics(ctx context.Context, query *model.EnterpriseDistrictStatisticsQuery) ([]*model.EnterpriseDistrictStatistics, error) {
	panic(fmt.Errorf("not implemented"))
}
