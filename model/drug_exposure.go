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

type DrugExposure struct {
	DrugExposureID            int64       `gorm:"column:drug_exposure_id;primary_key" json:"drug_exposure_id"`
	PersonID                  int64       `gorm:"column:person_id" json:"person_id"`
	DrugConceptID             int         `gorm:"column:drug_concept_id" json:"drug_concept_id"`
	DrugExposureStartDate     null.Time   `gorm:"column:drug_exposure_start_date" json:"drug_exposure_start_date"`
	DrugExposureStartDatetime time.Time   `gorm:"column:drug_exposure_start_datetime" json:"drug_exposure_start_datetime"`
	DrugExposureEndDate       null.Time   `gorm:"column:drug_exposure_end_date" json:"drug_exposure_end_date"`
	DrugExposureEndDatetime   time.Time   `gorm:"column:drug_exposure_end_datetime" json:"drug_exposure_end_datetime"`
	VerbatimEndDate           null.Time   `gorm:"column:verbatim_end_date" json:"verbatim_end_date"`
	DrugTypeConceptID         int         `gorm:"column:drug_type_concept_id" json:"drug_type_concept_id"`
	StopReason                null.String `gorm:"column:stop_reason" json:"stop_reason"`
	Refills                   null.Int    `gorm:"column:refills" json:"refills"`
	Quantity                  null.Float  `gorm:"column:quantity" json:"quantity"`
	DaysSupply                null.Int    `gorm:"column:days_supply" json:"days_supply"`
	Sig                       null.String `gorm:"column:sig" json:"sig"`
	RouteConceptID            int         `gorm:"column:route_concept_id" json:"route_concept_id"`
	LotNumber                 null.String `gorm:"column:lot_number" json:"lot_number"`
	ProviderID                null.Int    `gorm:"column:provider_id" json:"provider_id"`
	VisitOccurrenceID         null.Int    `gorm:"column:visit_occurrence_id" json:"visit_occurrence_id"`
	VisitDetailID             null.Int    `gorm:"column:visit_detail_id" json:"visit_detail_id"`
	DrugSourceValue           null.String `gorm:"column:drug_source_value" json:"drug_source_value"`
	DrugSourceConceptID       int         `gorm:"column:drug_source_concept_id" json:"drug_source_concept_id"`
	RouteSourceValue          null.String `gorm:"column:route_source_value" json:"route_source_value"`
	DoseUnitSourceValue       null.String `gorm:"column:dose_unit_source_value" json:"dose_unit_source_value"`
}

// TableName sets the insert table name for this struct type
func (d *DrugExposure) TableName() string {
	return "drug_exposure"
}
