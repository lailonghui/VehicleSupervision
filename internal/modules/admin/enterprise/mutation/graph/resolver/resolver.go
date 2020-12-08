package resolver

import (
	model1 "VehicleSupervision/internal/modules/admin/enterprise/model"
	"VehicleSupervision/internal/modules/admin/enterprise/mutation/graph/model"
	"VehicleSupervision/pkg/xid"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

func (r Resolver) batchInsertParamConvert(objects []*model.EnterpriseInsertInput) (rs []*model1.Enterprise) {
	for _, v := range objects {
		rs = append(rs, r.insertParamConvert(v))
	}
	return
}

func (r Resolver) insertParamConvert(v *model.EnterpriseInsertInput) (rs *model1.Enterprise) {
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
