package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/training/mutation/graph/generated"
	"VehicleSupervision/internal/modules/training/mutation/graph/model"
	"VehicleSupervision/pkg/graphql/util"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteAnswerLog(ctx context.Context, where model.AnswerLogBoolExp) (*model.AnswerLogMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model.AnswerLog{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model.AnswerLog
	if len(preloads) > 0 {
		// 如果请求的字段不为空，则先查询一遍数据库
		tx := tx.Select(preloads)
		tx = tx.Find(&rs)
		// 如果查询结果含有错误，则返回错误
		if err := tx.Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, nil
			}
			return nil, err
		}
	}
	// 删除
	tx = tx.Delete(nil)
	if err := tx.Error; err != nil {
		return nil, err
	}
	return &model.AnswerLogMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, tx.Error
}

func (r *mutationResolver) DeleteAnswerLogByPk(ctx context.Context, id int64) (*model.AnswerLog, error) {
	var rs = model.AnswerLog{}
	tx := db.DB.Model(&model.AnswerLog{}).Find(&rs, id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	tx = db.DB.Delete(id)
	return &rs, tx.Error
}

func (r *mutationResolver) DeleteEnterpriseTraining(ctx context.Context, where model.EnterpriseTrainingBoolExp) (*model.EnterpriseTrainingMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model.EnterpriseTraining{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model.EnterpriseTraining
	if len(preloads) > 0 {
		// 如果请求的字段不为空，则先查询一遍数据库
		tx := tx.Select(preloads)
		tx = tx.Find(&rs)
		// 如果查询结果含有错误，则返回错误
		if err := tx.Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, nil
			}
			return nil, err
		}
	}
	// 删除
	tx = tx.Delete(nil)
	if err := tx.Error; err != nil {
		return nil, err
	}
	return &model.EnterpriseTrainingMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, tx.Error
}

func (r *mutationResolver) DeleteEnterpriseTrainingByPk(ctx context.Context, id int64) (*model.EnterpriseTraining, error) {
	var rs = model.EnterpriseTraining{}
	tx := db.DB.Model(&model.EnterpriseTraining{}).Find(&rs, id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	tx = db.DB.Delete(id)
	return &rs, tx.Error
}

func (r *mutationResolver) DeleteGovernmentManager(ctx context.Context, where model.GovernmentManagerBoolExp) (*model.GovernmentManagerMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model.GovernmentManager{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model.GovernmentManager
	if len(preloads) > 0 {
		// 如果请求的字段不为空，则先查询一遍数据库
		tx := tx.Select(preloads)
		tx = tx.Find(&rs)
		// 如果查询结果含有错误，则返回错误
		if err := tx.Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, nil
			}
			return nil, err
		}
	}
	// 删除
	tx = tx.Delete(nil)
	if err := tx.Error; err != nil {
		return nil, err
	}
	return &model.GovernmentManagerMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, tx.Error
}

func (r *mutationResolver) DeleteGovernmentManagerByPk(ctx context.Context, id int64) (*model.GovernmentManager, error) {
	var rs = model.GovernmentManager{}
	tx := db.DB.Model(&model.GovernmentManager{}).Find(&rs, id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	tx = db.DB.Delete(id)
	return &rs, tx.Error
}

func (r *mutationResolver) DeleteScoreLog(ctx context.Context, where model.ScoreLogBoolExp) (*model.ScoreLogMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model.ScoreLog{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model.ScoreLog
	if len(preloads) > 0 {
		// 如果请求的字段不为空，则先查询一遍数据库
		tx := tx.Select(preloads)
		tx = tx.Find(&rs)
		// 如果查询结果含有错误，则返回错误
		if err := tx.Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, nil
			}
			return nil, err
		}
	}
	// 删除
	tx = tx.Delete(nil)
	if err := tx.Error; err != nil {
		return nil, err
	}
	return &model.ScoreLogMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, tx.Error
}

func (r *mutationResolver) DeleteScoreLogByPk(ctx context.Context, id int64) (*model.ScoreLog, error) {
	var rs = model.ScoreLog{}
	tx := db.DB.Model(&model.ScoreLog{}).Find(&rs, id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	tx = db.DB.Delete(id)
	return &rs, tx.Error
}

func (r *mutationResolver) DeleteTrainingMaterial(ctx context.Context, where model.TrainingMaterialBoolExp) (*model.TrainingMaterialMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model.TrainingMaterial{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model.TrainingMaterial
	if len(preloads) > 0 {
		// 如果请求的字段不为空，则先查询一遍数据库
		tx := tx.Select(preloads)
		tx = tx.Find(&rs)
		// 如果查询结果含有错误，则返回错误
		if err := tx.Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, nil
			}
			return nil, err
		}
	}
	// 删除
	tx = tx.Delete(nil)
	if err := tx.Error; err != nil {
		return nil, err
	}
	return &model.TrainingMaterialMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, tx.Error
}

func (r *mutationResolver) DeleteTrainingMaterialByPk(ctx context.Context, id int64) (*model.TrainingMaterial, error) {
	var rs = model.TrainingMaterial{}
	tx := db.DB.Model(&model.TrainingMaterial{}).Find(&rs, id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	tx = db.DB.Delete(id)
	return &rs, tx.Error
}

func (r *mutationResolver) InsertAnswerLog(ctx context.Context, objects []*model.AnswerLogInsertInput, onConflict *model.AnswerLogOnConflict) (*model.AnswerLogMutationResponse, error) {
	rs := r.batchAnswerLogInsertParamConvert(objects)
	tx := db.DB.Model(&model.AnswerLog{}).Save(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.AnswerLogMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertAnswerLogOne(ctx context.Context, object model.AnswerLogInsertInput, onConflict *model.AnswerLogOnConflict) (*model.AnswerLog, error) {
	rs := r.insertAnswerLogParamConvert(&object)
	tx := db.DB.Model(&model.AnswerLog{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) InsertDriverStudyDetails(ctx context.Context, objects []*model.DriverStudyDetailsInsertInput, onConflict *model.DriverStudyDetailsOnConflict) (*model.DriverStudyDetailsMutationResponse, error) {
	rs := r.batchDriverStudyInsertParamConvert(objects)
	tx := db.DB.Model(&model.DriverStudyDetails{}).Save(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.DriverStudyDetailsMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertDriverStudyDetailsOne(ctx context.Context, object model.DriverStudyDetailsInsertInput, onConflict *model.DriverStudyDetailsOnConflict) (*model.DriverStudyDetails, error) {
	rs := r.insertDriverStudyParamConvert(&object)
	tx := db.DB.Model(&model.DriverStudyDetails{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) InsertEnterpriseTraining(ctx context.Context, objects []*model.EnterpriseTrainingInsertInput, onConflict *model.EnterpriseTrainingOnConflict) (*model.EnterpriseTrainingMutationResponse, error) {
	rs := r.batchEnterpriseTrainingInsertParamConvert(objects)
	tx := db.DB.Model(&model.EnterpriseTraining{}).Save(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.EnterpriseTrainingMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertEnterpriseTrainingOne(ctx context.Context, object model.EnterpriseTrainingInsertInput, onConflict *model.EnterpriseTrainingOnConflict) (*model.EnterpriseTraining, error) {
	rs := r.insertEnterpriseTrainingParamConvert(&object)
	tx := db.DB.Model(&model.EnterpriseTraining{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) InsertGovernmentManager(ctx context.Context, objects []*model.GovernmentManagerInsertInput, onConflict *model.GovernmentManagerOnConflict) (*model.GovernmentManagerMutationResponse, error) {
	rs := r.batchGovernmentManagerInsertParamConvert(objects)
	tx := db.DB.Model(&model.GovernmentManager{}).Save(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.GovernmentManagerMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertGovernmentManagerOne(ctx context.Context, object model.GovernmentManagerInsertInput, onConflict *model.GovernmentManagerOnConflict) (*model.GovernmentManager, error) {
	rs := r.insertGovernmentManagerParamConvert(&object)
	tx := db.DB.Model(&model.GovernmentManager{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) InsertScoreLog(ctx context.Context, objects []*model.ScoreLogInsertInput, onConflict *model.ScoreLogOnConflict) (*model.ScoreLogMutationResponse, error) {
	rs := r.batchScoreLogInsertParamConvert(objects)
	tx := db.DB.Model(&model.ScoreLog{}).Save(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.ScoreLogMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertScoreLogOne(ctx context.Context, object model.ScoreLogInsertInput, onConflict *model.ScoreLogOnConflict) (*model.ScoreLog, error) {
	rs := r.insertScoreLogParamConvert(&object)
	tx := db.DB.Model(&model.ScoreLog{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) InsertTrainingMaterial(ctx context.Context, objects []*model.TrainingMaterialInsertInput, onConflict *model.TrainingMaterialOnConflict) (*model.TrainingMaterialMutationResponse, error) {
	rs := r.batchTrainingMaterialInsertParamConvert(objects)
	tx := db.DB.Model(&model.TrainingMaterial{}).Save(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.TrainingMaterialMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertTrainingMaterialOne(ctx context.Context, object model.TrainingMaterialInsertInput, onConflict *model.TrainingMaterialOnConflict) (*model.TrainingMaterial, error) {
	rs := r.insertTrainingMaterialParamConvert(&object)
	tx := db.DB.Model(&model.TrainingMaterial{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateAnswerLog(ctx context.Context, inc *model.AnswerLogIncInput, set *model.AnswerLogSetInput, where model.AnswerLogBoolExp) (*model.AnswerLogMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model.AnswerLog{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.AnswerLogMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.AnswerLogMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateAnswerLogByPk(ctx context.Context, inc *model.AnswerLogIncInput, set *model.AnswerLogSetInput, pkColumns model.AnswerLogPkColumnsInput) (*model.AnswerLog, error) {
	tx := db.DB.Where("id = ?", pkColumns.ID)
	qt := util.NewQueryTranslator(tx, &model.AnswerLog{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model.AnswerLog
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *mutationResolver) UpdateDriverStudyDetails(ctx context.Context, inc *model.DriverStudyDetailsIncInput, set *model.DriverStudyDetailsSetInput, where model.DriverStudyDetailsBoolExp) (*model.DriverStudyDetailsMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model.DriverStudyDetails{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.DriverStudyDetailsMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.DriverStudyDetailsMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateDriverStudyDetailsByPk(ctx context.Context, inc *model.DriverStudyDetailsIncInput, set *model.DriverStudyDetailsSetInput, pkColumns model.DriverStudyDetailsPkColumnsInput) (*model.DriverStudyDetails, error) {
	tx := db.DB.Where("id = ?", pkColumns.ID)
	qt := util.NewQueryTranslator(tx, &model.DriverStudyDetails{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model.DriverStudyDetails
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *mutationResolver) UpdateEnterpriseTraining(ctx context.Context, inc *model.EnterpriseTrainingIncInput, set *model.EnterpriseTrainingSetInput, where model.EnterpriseTrainingBoolExp) (*model.EnterpriseTrainingMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model.EnterpriseTraining{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.EnterpriseTrainingMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.EnterpriseTrainingMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateEnterpriseTrainingByPk(ctx context.Context, inc *model.EnterpriseTrainingIncInput, set *model.EnterpriseTrainingSetInput, pkColumns model.EnterpriseTrainingPkColumnsInput) (*model.EnterpriseTraining, error) {
	tx := db.DB.Where("id = ?", pkColumns.ID)
	qt := util.NewQueryTranslator(tx, &model.EnterpriseTraining{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model.EnterpriseTraining
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *mutationResolver) UpdateGovernmentManager(ctx context.Context, inc *model.GovernmentManagerIncInput, set *model.GovernmentManagerSetInput, where model.GovernmentManagerBoolExp) (*model.GovernmentManagerMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model.GovernmentManager{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.GovernmentManagerMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.GovernmentManagerMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateGovernmentManagerByPk(ctx context.Context, inc *model.GovernmentManagerIncInput, set *model.GovernmentManagerSetInput, pkColumns model.GovernmentManagerPkColumnsInput) (*model.GovernmentManager, error) {
	tx := db.DB.Where("id = ?", pkColumns.ID)
	qt := util.NewQueryTranslator(tx, &model.GovernmentManager{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model.GovernmentManager
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *mutationResolver) UpdateScoreLog(ctx context.Context, inc *model.ScoreLogIncInput, set *model.ScoreLogSetInput, where model.ScoreLogBoolExp) (*model.ScoreLogMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model.ScoreLog{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.ScoreLogMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.ScoreLogMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateScoreLogByPk(ctx context.Context, inc *model.ScoreLogIncInput, set *model.ScoreLogSetInput, pkColumns model.ScoreLogPkColumnsInput) (*model.ScoreLog, error) {
	tx := db.DB.Where("id = ?", pkColumns.ID)
	qt := util.NewQueryTranslator(tx, &model.ScoreLog{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model.ScoreLog
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *mutationResolver) UpdateTrainingMaterial(ctx context.Context, inc *model.TrainingMaterialIncInput, set *model.TrainingMaterialSetInput, where model.TrainingMaterialBoolExp) (*model.TrainingMaterialMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model.TrainingMaterial{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.TrainingMaterialMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.TrainingMaterialMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateTrainingMaterialByPk(ctx context.Context, inc *model.TrainingMaterialIncInput, set *model.TrainingMaterialSetInput, pkColumns model.TrainingMaterialPkColumnsInput) (*model.TrainingMaterial, error) {
	tx := db.DB.Where("id = ?", pkColumns.ID)
	qt := util.NewQueryTranslator(tx, &model.TrainingMaterial{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model.TrainingMaterial
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) Bug(ctx context.Context) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
