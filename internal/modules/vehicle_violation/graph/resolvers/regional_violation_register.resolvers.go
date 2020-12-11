package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/modules/vehicle_violation/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) DeleteRegionalViolationRegister(ctx context.Context, where model.RegionalViolationRegisterBoolExp) (*model.RegionalViolationRegisterMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteRegionalViolationRegisterByPk(ctx context.Context, id int64, regionalViolationRegisterID string) (*model.RegionalViolationRegister, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertRegionalViolationRegister(ctx context.Context, objects []*model.RegionalViolationRegisterInsertInput, onConflict *model.RegionalViolationRegisterOnConflict) (*model.RegionalViolationRegisterMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertRegionalViolationRegisterOne(ctx context.Context, object model.RegionalViolationRegisterInsertInput, onConflict *model.RegionalViolationRegisterOnConflict) (*model.RegionalViolationRegister, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateRegionalViolationRegister(ctx context.Context, inc *model.RegionalViolationRegisterIncInput, set *model.RegionalViolationRegisterSetInput, where model.RegionalViolationRegisterBoolExp) (*model.RegionalViolationRegisterMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateRegionalViolationRegisterByPk(ctx context.Context, inc *model.RegionalViolationRegisterIncInput, set *model.RegionalViolationRegisterSetInput, pkColumns model.RegionalViolationRegisterPkColumnsInput) (*model.RegionalViolationRegister, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) RegionalViolationRegister(ctx context.Context, distinctOn []model.RegionalViolationRegisterSelectColumn, limit *int, offset *int, orderBy []*model.RegionalViolationRegisterOrderBy, where *model.RegionalViolationRegisterBoolExp) ([]*model.RegionalViolationRegister, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) RegionalViolationRegisterAggregate(ctx context.Context, distinctOn []model.RegionalViolationRegisterSelectColumn, limit *int, offset *int, orderBy []*model.RegionalViolationRegisterOrderBy, where *model.RegionalViolationRegisterBoolExp) (*model.RegionalViolationRegisterAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) RegionalViolationRegisterByPk(ctx context.Context, id int64, regionalViolationRegisterID string) (*model.RegionalViolationRegister, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) RegionalViolationRegister(ctx context.Context, distinctOn []model.RegionalViolationRegisterSelectColumn, limit *int, offset *int, orderBy []*model.RegionalViolationRegisterOrderBy, where *model.RegionalViolationRegisterBoolExp) (<-chan []*model.RegionalViolationRegister, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) RegionalViolationRegisterAggregate(ctx context.Context, distinctOn []model.RegionalViolationRegisterSelectColumn, limit *int, offset *int, orderBy []*model.RegionalViolationRegisterOrderBy, where *model.RegionalViolationRegisterBoolExp) (<-chan *model.RegionalViolationRegisterAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) RegionalViolationRegisterByPk(ctx context.Context, id int64, regionalViolationRegisterID string) (<-chan *model.RegionalViolationRegister, error) {
	panic(fmt.Errorf("not implemented"))
}
