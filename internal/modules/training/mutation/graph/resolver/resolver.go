//go:generate go run github.com/99designs/gqlgen
package resolver

import (
	"VehicleSupervision/internal/modules/training/mutation/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

func (r Resolver) batchDriverStudyInsertParamConvert(objects []*model.DriverStudyDetailsInsertInput) (rs []*model.DriverStudyDetails) {
	for _, v := range objects {
		rs = append(rs, r.insertDriverStudyParamConvert(v))
	}
	return
}

func (r Resolver) insertDriverStudyParamConvert(v *model.DriverStudyDetailsInsertInput) (rs *model.DriverStudyDetails) {
	rs = &model.DriverStudyDetails{
		CreateBy:            v.CreateBy,
		CreateTime:          v.CreateTime,
		DriverID:            v.DriverID,
		DriverTrainingID:    v.DriverTrainingID,
		EnterpriseID:        v.EnterpriseID,
		ID:                  *v.ID,
		IsTrainingCompleted: v.IsTrainingCompleted,
		MaterialID:          v.MaterialID,
		Score:               v.Score,
		TrainingTime:        v.TrainingTime,
	}
	return
}

func (r Resolver) batchAnswerLogInsertParamConvert(objects []*model.AnswerLogInsertInput) (rs []*model.AnswerLog) {
	for _, v := range objects {
		rs = append(rs, r.insertAnswerLogParamConvert(v))
	}
	return
}

func (r Resolver) insertAnswerLogParamConvert(v *model.AnswerLogInsertInput) (rs *model.AnswerLog) {
	rs = &model.AnswerLog{
		Answer:      v.Answer,
		AnswerLogID: v.AnswerLogID,
		CreateAt:    v.CreateAt,
		CreateBy:    v.CreateBy,
		DriverID:    v.DriverID,
		ID:          *v.ID,
		MaterialID:  v.MaterialID,
		UpdateAt:    v.UpdateAt,
		UpdateBy:    v.UpdateBy,
	}
	return
}

func (r Resolver) batchEnterpriseTrainingInsertParamConvert(objects []*model.EnterpriseTrainingInsertInput) (rs []*model.EnterpriseTraining) {
	for _, v := range objects {
		rs = append(rs, r.insertEnterpriseTrainingParamConvert(v))
	}
	return
}

func (r Resolver) insertEnterpriseTrainingParamConvert(v *model.EnterpriseTrainingInsertInput) (rs *model.EnterpriseTraining) {
	rs = &model.EnterpriseTraining{
		CreateAt:             v.CreateAt,
		CreateBy:             v.CreateBy,
		DeleteAt:             v.DeleteAt,
		DeleteBy:             v.DeleteBy,
		EnterpriseID:         v.EnterpriseID,
		EnterpriseTrainingID: v.EnterpriseTrainingID,
		ID:                   *v.ID,
		IsDelete:             v.IsDelete,
		IsReceived:           v.IsReceived,
		MaterialID:           v.MaterialID,
		UpdateAt:             v.UpdateAt,
		UpdateBy:             v.UpdateBy,
	}
	return
}

func (r Resolver) batchGovernmentManagerInsertParamConvert(objects []*model.GovernmentManagerInsertInput) (rs []*model.GovernmentManager) {
	for _, v := range objects {
		rs = append(rs, r.insertGovernmentManagerParamConvert(v))
	}
	return
}

func (r Resolver) insertGovernmentManagerParamConvert(v *model.GovernmentManagerInsertInput) (rs *model.GovernmentManager) {
	rs = &model.GovernmentManager{
		AllDriverCount: v.AllDriverCount,
		AreaID:         v.AreaID,
		CreateAt:       v.CreateAt,
		CreateBy:       v.CreateBy,
		Deadline:       v.Deadline,
		DeleteAt:       v.DeleteAt,
		DeleteBy:       v.DeleteBy,
		EnterpriseID:   v.EnterpriseID,
		GovID:          v.GovID,
		ID:             *v.ID,
		IsDelete:       v.IsDelete,
		IsReceived:     v.IsReceived,
		ManageID:       v.ManageID,
		MaterialID:     v.MaterialID,
		Note:           v.Note,
		StartTime:      v.StartTime,
		TrainedDrivers: v.TrainedDrivers,
		UpdateAt:       v.UpdateAt,
		UpdateBy:       v.UpdateBy,
	}
	return
}

func (r Resolver) batchScoreLogInsertParamConvert(objects []*model.ScoreLogInsertInput) (rs []*model.ScoreLog) {
	for _, v := range objects {
		rs = append(rs, r.insertScoreLogParamConvert(v))
	}
	return
}

func (r Resolver) insertScoreLogParamConvert(v *model.ScoreLogInsertInput) (rs *model.ScoreLog) {
	rs = &model.ScoreLog{
		CreateAt:         v.CreateAt,
		CreateBy:         v.CreateBy,
		ID:               *v.ID,
		IDNumber:         v.IDNumber,
		Name:             v.Name,
		ScoreLogID:       v.ScoreLogID,
		TrainingPrograms: v.TrainingPrograms,
		TrainingScore:    v.TrainingScore,
		UpdateAt:         v.UpdateAt,
		UpdateBy:         v.UpdateBy,
	}
	return
}

func (r Resolver) batchTrainingMaterialInsertParamConvert(objects []*model.TrainingMaterialInsertInput) (rs []*model.TrainingMaterial) {
	for _, v := range objects {
		rs = append(rs, r.insertTrainingMaterialParamConvert(v))
	}
	return
}

func (r Resolver) insertTrainingMaterialParamConvert(v *model.TrainingMaterialInsertInput) (rs *model.TrainingMaterial) {
	rs = &model.TrainingMaterial{
		Contents:        v.Contents,
		CreateAt:        v.CreateAt,
		CreateBy:        v.CreateBy,
		DeleteAt:        v.DeleteAt,
		DeleteBy:        v.DeleteBy,
		ID:              *v.ID,
		IsDelete:        v.IsDelete,
		MaterialAddress: v.MaterialAddress,
		MaterialID:      v.MaterialID,
		Title:           v.Title,
		Type:            v.Type,
		UpdateAt:        v.UpdateAt,
		UpdateBy:        v.UpdateBy,
	}
	return
}
