package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/modules/vehicle_violation/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) DeleteEnterpriseDeductionItems(ctx context.Context, where model.EnterpriseDeductionItemsBoolExp) (*model.EnterpriseDeductionItemsMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteEnterpriseDeductionItemsByPk(ctx context.Context, enterpriseDeductionItemID string, id int64) (*model.EnterpriseDeductionItems, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertEnterpriseDeductionItems(ctx context.Context, objects []*model.EnterpriseDeductionItemsInsertInput, onConflict *model.EnterpriseDeductionItemsOnConflict) (*model.EnterpriseDeductionItemsMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertEnterpriseDeductionItemsOne(ctx context.Context, object model.EnterpriseDeductionItemsInsertInput, onConflict *model.EnterpriseDeductionItemsOnConflict) (*model.EnterpriseDeductionItems, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateEnterpriseDeductionItems(ctx context.Context, inc *model.EnterpriseDeductionItemsIncInput, set *model.EnterpriseDeductionItemsSetInput, where model.EnterpriseDeductionItemsBoolExp) (*model.EnterpriseDeductionItemsMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateEnterpriseDeductionItemsByPk(ctx context.Context, inc *model.EnterpriseDeductionItemsIncInput, set *model.EnterpriseDeductionItemsSetInput, pkColumns model.EnterpriseDeductionItemsPkColumnsInput) (*model.EnterpriseDeductionItems, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) EnterpriseDeductionItems(ctx context.Context, distinctOn []model.EnterpriseDeductionItemsSelectColumn, limit *int, offset *int, orderBy []*model.EnterpriseDeductionItemsOrderBy, where *model.EnterpriseDeductionItemsBoolExp) ([]*model.EnterpriseDeductionItems, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) EnterpriseDeductionItemsAggregate(ctx context.Context, distinctOn []model.EnterpriseDeductionItemsSelectColumn, limit *int, offset *int, orderBy []*model.EnterpriseDeductionItemsOrderBy, where *model.EnterpriseDeductionItemsBoolExp) (*model.EnterpriseDeductionItemsAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) EnterpriseDeductionItemsByPk(ctx context.Context, enterpriseDeductionItemID string, id int64) (*model.EnterpriseDeductionItems, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) EnterpriseDeductionItems(ctx context.Context, distinctOn []model.EnterpriseDeductionItemsSelectColumn, limit *int, offset *int, orderBy []*model.EnterpriseDeductionItemsOrderBy, where *model.EnterpriseDeductionItemsBoolExp) (<-chan []*model.EnterpriseDeductionItems, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) EnterpriseDeductionItemsAggregate(ctx context.Context, distinctOn []model.EnterpriseDeductionItemsSelectColumn, limit *int, offset *int, orderBy []*model.EnterpriseDeductionItemsOrderBy, where *model.EnterpriseDeductionItemsBoolExp) (<-chan *model.EnterpriseDeductionItemsAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) EnterpriseDeductionItemsByPk(ctx context.Context, enterpriseDeductionItemID string, id int64) (<-chan *model.EnterpriseDeductionItems, error) {
	panic(fmt.Errorf("not implemented"))
}
