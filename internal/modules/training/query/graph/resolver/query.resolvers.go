package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/training/query/graph/generated"
	"VehicleSupervision/internal/modules/training/query/graph/model"
	"VehicleSupervision/pkg/graphql/util"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (r *queryResolver) AnswerLog(ctx context.Context, distinctOn []model.AnswerLogSelectColumn, limit *int, offset *int, orderBy []*model.AnswerLogOrderBy, where *model.AnswerLogBoolExp) ([]*model.AnswerLog, error) {
	qt := util.NewQueryTranslator(db.DB, &model.AnswerLog{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model.AnswerLog
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) AnswerLogAggregate(ctx context.Context, distinctOn []model.AnswerLogSelectColumn, limit *int, offset *int, orderBy []*model.AnswerLogOrderBy, where *model.AnswerLogBoolExp) (*model.AnswerLogAggregate, error) {
	var rs model.AnswerLogAggregate

	qt := util.NewQueryTranslator(db.DB, &model.AnswerLog{})
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

func (r *queryResolver) AnswerLogByPk(ctx context.Context, id int64) (*model.AnswerLog, error) {
	var rs model.AnswerLog
	tx := db.DB.Model(&model.AnswerLog{}).First(&rs, id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}

func (r *queryResolver) DriverStudyDetails(ctx context.Context, distinctOn []model.DriverStudyDetailsSelectColumn, limit *int, offset *int, orderBy []*model.DriverStudyDetailsOrderBy, where *model.DriverStudyDetailsBoolExp) ([]*model.DriverStudyDetails, error) {
	qt := util.NewQueryTranslator(db.DB, &model.DriverStudyDetails{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model.DriverStudyDetails
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) DriverStudyDetailsAggregate(ctx context.Context, distinctOn []model.DriverStudyDetailsSelectColumn, limit *int, offset *int, orderBy []*model.DriverStudyDetailsOrderBy, where *model.DriverStudyDetailsBoolExp) (*model.DriverStudyDetailsAggregate, error) {
	var rs model.DriverStudyDetailsAggregate

	qt := util.NewQueryTranslator(db.DB, &model.DriverStudyDetails{})
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

func (r *queryResolver) DriverStudyDetailsByPk(ctx context.Context, id int64) (*model.DriverStudyDetails, error) {
	var rs model.DriverStudyDetails
	tx := db.DB.Model(&model.DriverStudyDetails{}).First(&rs, id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}

func (r *queryResolver) EnterpriseTraining(ctx context.Context, distinctOn []model.EnterpriseTrainingSelectColumn, limit *int, offset *int, orderBy []*model.EnterpriseTrainingOrderBy, where *model.EnterpriseTrainingBoolExp) ([]*model.EnterpriseTraining, error) {
	qt := util.NewQueryTranslator(db.DB, &model.EnterpriseTraining{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model.EnterpriseTraining
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) EnterpriseTrainingAggregate(ctx context.Context, distinctOn []model.EnterpriseTrainingSelectColumn, limit *int, offset *int, orderBy []*model.EnterpriseTrainingOrderBy, where *model.EnterpriseTrainingBoolExp) (*model.EnterpriseTrainingAggregate, error) {
	var rs model.EnterpriseTrainingAggregate

	qt := util.NewQueryTranslator(db.DB, &model.EnterpriseTraining{})
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

func (r *queryResolver) EnterpriseTrainingByPk(ctx context.Context, id int64) (*model.EnterpriseTraining, error) {
	var rs model.EnterpriseTraining
	tx := db.DB.Model(&model.EnterpriseTraining{}).First(&rs, id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}

func (r *queryResolver) GovernmentManager(ctx context.Context, distinctOn []model.GovernmentManagerSelectColumn, limit *int, offset *int, orderBy []*model.GovernmentManagerOrderBy, where *model.GovernmentManagerBoolExp) ([]*model.GovernmentManager, error) {
	qt := util.NewQueryTranslator(db.DB, &model.GovernmentManager{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model.GovernmentManager
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) GovernmentManagerAggregate(ctx context.Context, distinctOn []model.GovernmentManagerSelectColumn, limit *int, offset *int, orderBy []*model.GovernmentManagerOrderBy, where *model.GovernmentManagerBoolExp) (*model.GovernmentManagerAggregate, error) {
	var rs model.GovernmentManagerAggregate

	qt := util.NewQueryTranslator(db.DB, &model.GovernmentManager{})
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

func (r *queryResolver) GovernmentManagerByPk(ctx context.Context, id int64) (*model.GovernmentManager, error) {
	var rs model.GovernmentManager
	tx := db.DB.Model(&model.GovernmentManager{}).First(&rs, id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}

func (r *queryResolver) ScoreLog(ctx context.Context, distinctOn []model.ScoreLogSelectColumn, limit *int, offset *int, orderBy []*model.ScoreLogOrderBy, where *model.ScoreLogBoolExp) ([]*model.ScoreLog, error) {
	qt := util.NewQueryTranslator(db.DB, &model.ScoreLog{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model.ScoreLog
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) ScoreLogAggregate(ctx context.Context, distinctOn []model.ScoreLogSelectColumn, limit *int, offset *int, orderBy []*model.ScoreLogOrderBy, where *model.ScoreLogBoolExp) (*model.ScoreLogAggregate, error) {
	var rs model.ScoreLogAggregate

	qt := util.NewQueryTranslator(db.DB, &model.ScoreLog{})
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

func (r *queryResolver) ScoreLogByPk(ctx context.Context, id int64) (*model.ScoreLog, error) {
	var rs model.ScoreLog
	tx := db.DB.Model(&model.ScoreLog{}).First(&rs, id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}

func (r *queryResolver) TrainingMaterial(ctx context.Context, distinctOn []model.TrainingMaterialSelectColumn, limit *int, offset *int, orderBy []*model.TrainingMaterialOrderBy, where *model.TrainingMaterialBoolExp) ([]*model.TrainingMaterial, error) {
	qt := util.NewQueryTranslator(db.DB, &model.TrainingMaterial{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model.TrainingMaterial
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) TrainingMaterialAggregate(ctx context.Context, distinctOn []model.TrainingMaterialSelectColumn, limit *int, offset *int, orderBy []*model.TrainingMaterialOrderBy, where *model.TrainingMaterialBoolExp) (*model.TrainingMaterialAggregate, error) {
	var rs model.TrainingMaterialAggregate

	qt := util.NewQueryTranslator(db.DB, &model.TrainingMaterial{})
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

func (r *queryResolver) TrainingMaterialByPk(ctx context.Context, id int64) (*model.TrainingMaterial, error) {
	var rs model.TrainingMaterial
	tx := db.DB.Model(&model.TrainingMaterial{}).First(&rs, id)
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
