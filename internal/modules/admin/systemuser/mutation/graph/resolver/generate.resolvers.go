package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/db"
	model1 "VehicleSupervision/internal/modules/admin/systemuser/model"
	"VehicleSupervision/internal/modules/admin/systemuser/mutation/graph/generated"
	"VehicleSupervision/internal/modules/admin/systemuser/mutation/graph/model"
	"VehicleSupervision/pkg/graphql/util"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteSystemUser(ctx context.Context, where model.SystemUserBoolExp) (*model.SystemUserMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.SystemUser{})
	tx := qt.Where(where).Finish()
	tx.Delete(model1.SystemUser{})
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.SystemUserMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) DeleteSystemUserByPk(ctx context.Context, id int64) (*model1.SystemUser, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertSystemUser(ctx context.Context, objects []*model.SystemUserInsertInput, onConflict *model.SystemUserOnConflict) (*model.SystemUserMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertSystemUserOne(ctx context.Context, object model.SystemUserInsertInput, onConflict *model.SystemUserOnConflict) (*model1.SystemUser, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateSystemUser(ctx context.Context, inc *model.SystemUserIncInput, set *model.SystemUserSetInput, where model.SystemUserBoolExp) (*model.SystemUserMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateSystemUserByPk(ctx context.Context, inc *model.SystemUserIncInput, set *model.SystemUserSetInput, pkColumns model.SystemUserPkColumnsInput) (*model1.SystemUser, error) {
	panic(fmt.Errorf("not implemented"))
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
