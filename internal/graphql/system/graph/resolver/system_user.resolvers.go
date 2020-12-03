package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"system-manage/internal/graphql/system/graph/model"
)

func (r *queryResolver) SystemUser(ctx context.Context, userID string) (*model.SystemUser, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) SystemUsers(ctx context.Context, query *model.SystemUsersQuery) (*model.SystemUserConnection, error) {
	panic(fmt.Errorf("not implemented"))
}
