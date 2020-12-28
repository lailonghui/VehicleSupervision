package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden OutageFilingUploadFileLoader string *VehicleSupervision/internal/modules/vehicle/model.OutageFilingUploadFile

func (t OutageFilingUploadFile) TableName() string {
	return "outage_filing_upload_file"
}

func (t *OutageFilingUploadFile) NewLoader() *OutageFilingUploadFileLoader {
	return &OutageFilingUploadFileLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*OutageFilingUploadFile, []error) {
			var rs []*OutageFilingUploadFile
			// TODO 按实际需要实现
			db.DB.Model(&OutageFilingUploadFile{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
