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

type VisitDetail struct {
	VisitDetailID              int64       `gorm:"column:visit_detail_id;primary_key" json:"visit_detail_id"`
	PersonID                   int64       `gorm:"column:person_id" json:"person_id"`
	VisitDetailConceptID       int         `gorm:"column:visit_detail_concept_id" json:"visit_detail_concept_id"`
	VisitDetailStartDate       null.Time   `gorm:"column:visit_detail_start_date" json:"visit_detail_start_date"`
	VisitDetailStartDatetime   time.Time   `gorm:"column:visit_detail_start_datetime" json:"visit_detail_start_datetime"`
	VisitDetailEndDate         null.Time   `gorm:"column:visit_detail_end_date" json:"visit_detail_end_date"`
	VisitDetailEndDatetime     time.Time   `gorm:"column:visit_detail_end_datetime" json:"visit_detail_end_datetime"`
	VisitDetailTypeConceptID   int         `gorm:"column:visit_detail_type_concept_id" json:"visit_detail_type_concept_id"`
	ProviderID                 null.Int    `gorm:"column:provider_id" json:"provider_id"`
	CareSiteID                 null.Int    `gorm:"column:care_site_id" json:"care_site_id"`
	DischargeToConceptID       int         `gorm:"column:discharge_to_concept_id" json:"discharge_to_concept_id"`
	AdmittedFromConceptID      int         `gorm:"column:admitted_from_concept_id" json:"admitted_from_concept_id"`
	AdmittedFromSourceValue    null.String `gorm:"column:admitted_from_source_value" json:"admitted_from_source_value"`
	VisitDetailSourceValue     null.String `gorm:"column:visit_detail_source_value" json:"visit_detail_source_value"`
	VisitDetailSourceConceptID int         `gorm:"column:visit_detail_source_concept_id" json:"visit_detail_source_concept_id"`
	DischargeToSourceValue     null.String `gorm:"column:discharge_to_source_value" json:"discharge_to_source_value"`
	PrecedingVisitDetailID     null.Int    `gorm:"column:preceding_visit_detail_id" json:"preceding_visit_detail_id"`
	VisitDetailParentID        null.Int    `gorm:"column:visit_detail_parent_id" json:"visit_detail_parent_id"`
	VisitOccurrenceID          int64       `gorm:"column:visit_occurrence_id" json:"visit_occurrence_id"`
}

// TableName sets the insert table name for this struct type
func (v *VisitDetail) TableName() string {
	return "visit_detail"
}
