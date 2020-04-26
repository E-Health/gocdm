package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
)

type DeviceExposure struct {
	DeviceExposureID            int64       `gorm:"column:device_exposure_id;primary_key" json:"device_exposure_id"`
	PersonID                    int64       `gorm:"column:person_id" json:"person_id"`
	DeviceConceptID             int         `gorm:"column:device_concept_id" json:"device_concept_id"`
	DeviceExposureStartDate     null.Time   `gorm:"column:device_exposure_start_date" json:"device_exposure_start_date"`
	DeviceExposureStartDatetime time.Time   `gorm:"column:device_exposure_start_datetime" json:"device_exposure_start_datetime"`
	DeviceExposureEndDate       null.Time   `gorm:"column:device_exposure_end_date" json:"device_exposure_end_date"`
	DeviceExposureEndDatetime   null.Time   `gorm:"column:device_exposure_end_datetime" json:"device_exposure_end_datetime"`
	DeviceTypeConceptID         int         `gorm:"column:device_type_concept_id" json:"device_type_concept_id"`
	UniqueDeviceID              null.String `gorm:"column:unique_device_id" json:"unique_device_id"`
	Quantity                    null.Int    `gorm:"column:quantity" json:"quantity"`
	ProviderID                  null.Int    `gorm:"column:provider_id" json:"provider_id"`
	VisitOccurrenceID           null.Int    `gorm:"column:visit_occurrence_id" json:"visit_occurrence_id"`
	VisitDetailID               null.Int    `gorm:"column:visit_detail_id" json:"visit_detail_id"`
	DeviceSourceValue           null.String `gorm:"column:device_source_value" json:"device_source_value"`
	DeviceSourceConceptID       int         `gorm:"column:device_source_concept_id" json:"device_source_concept_id"`
}

// TableName sets the insert table name for this struct type
func (d *DeviceExposure) TableName() string {
	return "device_exposure"
}
