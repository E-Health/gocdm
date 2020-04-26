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

type VisitOccurrence struct {
	VisitOccurrenceID          int64       `gorm:"column:visit_occurrence_id;primary_key" json:"visit_occurrence_id"`
	PersonID                   int64       `gorm:"column:person_id" json:"person_id"`
	VisitConceptID             int         `gorm:"column:visit_concept_id" json:"visit_concept_id"`
	VisitStartDate             null.Time   `gorm:"column:visit_start_date" json:"visit_start_date"`
	VisitStartDatetime         time.Time   `gorm:"column:visit_start_datetime" json:"visit_start_datetime"`
	VisitEndDate               null.Time   `gorm:"column:visit_end_date" json:"visit_end_date"`
	VisitEndDatetime           time.Time   `gorm:"column:visit_end_datetime" json:"visit_end_datetime"`
	VisitTypeConceptID         int         `gorm:"column:visit_type_concept_id" json:"visit_type_concept_id"`
	ProviderID                 null.Int    `gorm:"column:provider_id" json:"provider_id"`
	CareSiteID                 null.Int    `gorm:"column:care_site_id" json:"care_site_id"`
	VisitSourceValue           null.String `gorm:"column:visit_source_value" json:"visit_source_value"`
	VisitSourceConceptID       int         `gorm:"column:visit_source_concept_id" json:"visit_source_concept_id"`
	AdmittedFromConceptID      int         `gorm:"column:admitted_from_concept_id" json:"admitted_from_concept_id"`
	AdmittedFromSourceValue    null.String `gorm:"column:admitted_from_source_value" json:"admitted_from_source_value"`
	DischargeToSourceValue     null.String `gorm:"column:discharge_to_source_value" json:"discharge_to_source_value"`
	DischargeToConceptID       int         `gorm:"column:discharge_to_concept_id" json:"discharge_to_concept_id"`
	PrecedingVisitOccurrenceID null.Int    `gorm:"column:preceding_visit_occurrence_id" json:"preceding_visit_occurrence_id"`
}

// TableName sets the insert table name for this struct type
func (v *VisitOccurrence) TableName() string {
	return "visit_occurrence"
}
