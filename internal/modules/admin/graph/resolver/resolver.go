package resolver

import (
	"VehicleSupervision/internal/modules/admin/graph/model"
	model1 "VehicleSupervision/internal/modules/admin/model"
	"VehicleSupervision/internal/server/middle"
	"VehicleSupervision/pkg/xid"
	"context"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

// 获取dataloader实例
func (Resolver) loaders(ctx context.Context) *middle.Loaders {
	return ctx.Value(middle.DATA_LOADER_CONTEXT_KEY).(*middle.Loaders)
}

func (r Resolver) departmentInsertBatchConvert(objects []*model.DepartmentInsertInput) (rs []*model1.Department) {
	for _, v := range objects {
		rs = append(rs, r.departmentInsertConvert(v))
	}
	return
}

func (r Resolver) departmentInsertConvert(v *model.DepartmentInsertInput) (rs *model1.Department) {
	rs = &model1.Department{
		CreateAt:             v.CreateAt,
		CreateBy:             v.CreateBy,
		DeleteAt:             v.DeleteAt,
		DeleteBy:             v.DeleteBy,
		DepartmentCategory:   v.DepartmentCategory,
		DepartmentCode:       v.DepartmentCode,
		DepartmentID:         xid.GetXid(),
		DepartmentName:       v.DepartmentName,
		EnterpriseID:         *v.EnterpriseID,
		InternalNumber:       v.InternalNumber,
		IsDelete:             false,
		Remarks:              v.Remarks,
		SuperiorDepartmentID: v.SuperiorDepartmentID,
		UpdateAt:             v.UpdateAt,
		UpdateBy:             v.UpdateBy,
	}
	return
}

func (r Resolver) enterpriseInsertInputBatchConvert(objects []*model.EnterpriseInsertInput) (rs []*model1.Enterprise) {
	for _, v := range objects {
		rs = append(rs, r.enterpriseInsertInputConvert(v))
	}
	return
}

func (r Resolver) enterpriseInsertInputConvert(v *model.EnterpriseInsertInput) (rs *model1.Enterprise) {
	rs = &model1.Enterprise{
		AssociationReviewBy:              v.AssociationReviewBy,
		AssociationReviewOpinion:         v.AssociationReviewOpinion,
		AssociationReviewTime:            v.AssociationReviewTime,
		BrigadeID:                        v.BrigadeID,
		BrigadeReviewBy:                  v.BrigadeReviewBy,
		BrigadeReviewOpinion:             v.BrigadeReviewOpinion,
		BrigadeReviewTime:                v.BrigadeReviewTime,
		BusinessLicenseExpiryDate:        v.BusinessLicenseExpiryDate,
		BusinessLicenseIssuanceDate:      v.BusinessLicenseIssuanceDate,
		BusinessLicensePhoto:             v.BusinessLicensePhoto,
		BusinessPhoto:                    v.BusinessPhoto,
		BusinessScope:                    v.BusinessScope,
		CheckStatus:                      v.CheckStatus,
		CityID:                           v.CityID,
		ContactPersons:                   v.ContactPersons,
		CreateAt:                         v.CreateAt,
		CreateBy:                         v.CreateBy,
		DeleteAt:                         v.DeleteAt,
		DeleteBy:                         v.DeleteBy,
		DisplayNumber:                    v.DisplayNumber,
		DistrictID:                       v.DistrictID,
		EnterpriseAddress:                v.EnterpriseAddress,
		EnterpriseCode:                   v.EnterpriseCode,
		EnterpriseID:                     xid.GetXid(),
		EnterpriseLevel:                  v.EnterpriseLevel,
		EnterpriseName:                   v.EnterpriseName,
		EnterpriseNature:                 v.EnterpriseNature,
		EntrustedAgent:                   v.EntrustedAgent,
		EntrustedAgentIDCard:             v.EntrustedAgentIDCard,
		EntrustedAgentIDCardPhoto:        v.EntrustedAgentIDCardPhoto,
		EntrustedAgentPhone:              v.EntrustedAgentPhone,
		FaxNumber:                        v.FaxNumber,
		InstitutionCategory:              v.InstitutionCategory,
		IsBlack:                          v.IsBlack,
		IsDeleted:                        false,
		IsInput:                          v.IsInput,
		IsInstall:                        v.IsInstall,
		IsUploadProvince:                 v.IsUploadProvince,
		LegalRepresentative:              v.LegalRepresentative,
		LegalRepresentativeIDCard:        v.LegalRepresentativeIDCard,
		LegalRepresentativeIDCardPhoto:   v.LegalRepresentativeIDCardPhoto,
		LegalRepresentativePhone:         v.LegalRepresentativePhone,
		OperatingLicensePhoto:            v.OperatingLicensePhoto,
		OrganizationCode:                 v.OrganizationCode,
		OrganizationCodeCertificatePhoto: v.OrganizationCodeCertificatePhoto,
		PoliceStationID:                  v.PoliceStationID,
		ProvinceID:                       v.ProvinceID,
		RecordAt:                         v.RecordAt,
		RecordBy:                         v.RecordBy,
		Remarks:                          v.Remarks,
		Score:                            v.Score,
		SuperiorEnterpriseID:             v.SuperiorEnterpriseID,
		UpdateAt:                         v.UpdateAt,
		UpdateBy:                         v.UpdateBy,
		UpdateTimeIn:                     v.UpdateTimeIn,
	}
	return
}

func (r Resolver) systemUserInsertInputBatchConvert(objects []*model.SystemUserInsertInput) (rs []*model1.SystemUser) {
	for _, v := range objects {
		rs = append(rs, r.systemUserInsertInputConvert(v))
	}
	return
}

func (r Resolver) systemUserInsertInputConvert(v *model.SystemUserInsertInput) (rs *model1.SystemUser) {
	rs = &model1.SystemUser{
		IsBindIP:     v.IsBindIP,
		AppVersion:   v.AppVersion,
		AuditLevel:   v.AuditLevel,
		CreateBy:     v.CreateBy,
		CreatedAt:    v.CreatedAt,
		DeleteAt:     v.DeleteAt,
		DeleteBy:     v.DeleteBy,
		DepartmentID: v.DepartmentID,
		Email:        v.Email,
		EnterpriseID: v.EnterpriseID,
		Grade:        v.Grade,
		IPAddress:    v.IPAddress,
		IsDelete:     false,
		IsValid:      v.IsValid,
		UserID:       xid.GetXid(),
		Mkey:         v.Mkey,
		Mobile:       v.Mobile,
		Password:     *v.Password,
		ProxyUser:    v.ProxyUser,
		Remarks:      v.Remarks,
		Telephone:    v.Telephone,
		Username:     *v.Username,
		Ukey:         v.Ukey,
		UpdateAt:     v.UpdateAt,
		UpdateBy:     v.UpdateBy,
		UserState:    v.UserState,
		UserType:     v.UserType,
	}
	return
}
