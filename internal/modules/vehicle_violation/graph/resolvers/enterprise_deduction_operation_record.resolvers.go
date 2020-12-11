package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/modules/vehicle_violation/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) DeleteEnterpriseDeductionOperationRecord(ctx context.Context, where model.EnterpriseDeductionOperationRecordBoolExp) (*model.EnterpriseDeductionOperationRecordMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteEnterpriseDeductionOperationRecordByPk(ctx context.Context, enterpriseDuductionOperationID string, id int64) (*model.EnterpriseDeductionOperationRecord, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertEnterpriseDeductionOperationRecord(ctx context.Context, objects []*model.EnterpriseDeductionOperationRecordInsertInput, onConflict *model.EnterpriseDeductionOperationRecordOnConflict) (*model.EnterpriseDeductionOperationRecordMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertEnterpriseDeductionOperationRecordOne(ctx context.Context, object model.EnterpriseDeductionOperationRecordInsertInput, onConflict *model.EnterpriseDeductionOperationRecordOnConflict) (*model.EnterpriseDeductionOperationRecord, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateEnterpriseDeductionOperationRecord(ctx context.Context, inc *model.EnterpriseDeductionOperationRecordIncInput, set *model.EnterpriseDeductionOperationRecordSetInput, where model.EnterpriseDeductionOperationRecordBoolExp) (*model.EnterpriseDeductionOperationRecordMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateEnterpriseDeductionOperationRecordByPk(ctx context.Context, inc *model.EnterpriseDeductionOperationRecordIncInput, set *model.EnterpriseDeductionOperationRecordSetInput, pkColumns model.EnterpriseDeductionOperationRecordPkColumnsInput) (*model.EnterpriseDeductionOperationRecord, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) EnterpriseDeductionOperationRecord(ctx context.Context, distinctOn []model.EnterpriseDeductionOperationRecordSelectColumn, limit *int, offset *int, orderBy []*model.EnterpriseDeductionOperationRecordOrderBy, where *model.EnterpriseDeductionOperationRecordBoolExp) ([]*model.EnterpriseDeductionOperationRecord, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) EnterpriseDeductionOperationRecordAggregate(ctx context.Context, distinctOn []model.EnterpriseDeductionOperationRecordSelectColumn, limit *int, offset *int, orderBy []*model.EnterpriseDeductionOperationRecordOrderBy, where *model.EnterpriseDeductionOperationRecordBoolExp) (*model.EnterpriseDeductionOperationRecordAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) EnterpriseDeductionOperationRecordByPk(ctx context.Context, enterpriseDuductionOperationID string, id int64) (*model.EnterpriseDeductionOperationRecord, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) EnterpriseDeductionOperationRecord(ctx context.Context, distinctOn []model.EnterpriseDeductionOperationRecordSelectColumn, limit *int, offset *int, orderBy []*model.EnterpriseDeductionOperationRecordOrderBy, where *model.EnterpriseDeductionOperationRecordBoolExp) (<-chan []*model.EnterpriseDeductionOperationRecord, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) EnterpriseDeductionOperationRecordAggregate(ctx context.Context, distinctOn []model.EnterpriseDeductionOperationRecordSelectColumn, limit *int, offset *int, orderBy []*model.EnterpriseDeductionOperationRecordOrderBy, where *model.EnterpriseDeductionOperationRecordBoolExp) (<-chan *model.EnterpriseDeductionOperationRecordAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) EnterpriseDeductionOperationRecordByPk(ctx context.Context, enterpriseDuductionOperationID string, id int64) (<-chan *model.EnterpriseDeductionOperationRecord, error) {
	panic(fmt.Errorf("not implemented"))
}
