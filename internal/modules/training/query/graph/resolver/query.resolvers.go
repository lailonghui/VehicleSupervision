package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/modules/training/query/graph/generated"
	"VehicleSupervision/internal/modules/training/query/graph/model"
	"context"
	"fmt"
)

func (r *queryResolver) AnswerLog(ctx context.Context, distinctOn []model.AnswerLogSelectColumn, limit *int, offset *int, orderBy []*model.AnswerLogOrderBy, where *model.AnswerLogBoolExp) ([]*model.AnswerLog, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AnswerLogAggregate(ctx context.Context, distinctOn []model.AnswerLogSelectColumn, limit *int, offset *int, orderBy []*model.AnswerLogOrderBy, where *model.AnswerLogBoolExp) (*model.AnswerLogAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AnswerLogByPk(ctx context.Context, id int64) (*model.AnswerLog, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DriverStudyDetails(ctx context.Context, distinctOn []model.DriverStudyDetailsSelectColumn, limit *int, offset *int, orderBy []*model.DriverStudyDetailsOrderBy, where *model.DriverStudyDetailsBoolExp) ([]*model.DriverStudyDetails, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DriverStudyDetailsAggregate(ctx context.Context, distinctOn []model.DriverStudyDetailsSelectColumn, limit *int, offset *int, orderBy []*model.DriverStudyDetailsOrderBy, where *model.DriverStudyDetailsBoolExp) (*model.DriverStudyDetailsAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DriverStudyDetailsByPk(ctx context.Context, id int64) (*model.DriverStudyDetails, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) EnterpriseTraining(ctx context.Context, distinctOn []model.EnterpriseTrainingSelectColumn, limit *int, offset *int, orderBy []*model.EnterpriseTrainingOrderBy, where *model.EnterpriseTrainingBoolExp) ([]*model.EnterpriseTraining, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) EnterpriseTrainingAggregate(ctx context.Context, distinctOn []model.EnterpriseTrainingSelectColumn, limit *int, offset *int, orderBy []*model.EnterpriseTrainingOrderBy, where *model.EnterpriseTrainingBoolExp) (*model.EnterpriseTrainingAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) EnterpriseTrainingByPk(ctx context.Context, id int64) (*model.EnterpriseTraining, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GovernmentManager(ctx context.Context, distinctOn []model.GovernmentManagerSelectColumn, limit *int, offset *int, orderBy []*model.GovernmentManagerOrderBy, where *model.GovernmentManagerBoolExp) ([]*model.GovernmentManager, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GovernmentManagerAggregate(ctx context.Context, distinctOn []model.GovernmentManagerSelectColumn, limit *int, offset *int, orderBy []*model.GovernmentManagerOrderBy, where *model.GovernmentManagerBoolExp) (*model.GovernmentManagerAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GovernmentManagerByPk(ctx context.Context, id int64) (*model.GovernmentManager, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ScoreLog(ctx context.Context, distinctOn []model.ScoreLogSelectColumn, limit *int, offset *int, orderBy []*model.ScoreLogOrderBy, where *model.ScoreLogBoolExp) ([]*model.ScoreLog, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ScoreLogAggregate(ctx context.Context, distinctOn []model.ScoreLogSelectColumn, limit *int, offset *int, orderBy []*model.ScoreLogOrderBy, where *model.ScoreLogBoolExp) (*model.ScoreLogAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ScoreLogByPk(ctx context.Context, id int64) (*model.ScoreLog, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) TrainingMaterial(ctx context.Context, distinctOn []model.TrainingMaterialSelectColumn, limit *int, offset *int, orderBy []*model.TrainingMaterialOrderBy, where *model.TrainingMaterialBoolExp) ([]*model.TrainingMaterial, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) TrainingMaterialAggregate(ctx context.Context, distinctOn []model.TrainingMaterialSelectColumn, limit *int, offset *int, orderBy []*model.TrainingMaterialOrderBy, where *model.TrainingMaterialBoolExp) (*model.TrainingMaterialAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) TrainingMaterialByPk(ctx context.Context, id int64) (*model.TrainingMaterial, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
