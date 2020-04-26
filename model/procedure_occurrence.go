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

type ProcedureOccurrence struct {
	ProcedureOccurrenceID    int64       `gorm:"column:procedure_occurrence_id;primary_key" json:"procedure_occurrence_id"`
	PersonID                 int64       `gorm:"column:person_id" json:"person_id"`
	ProcedureConceptID       int         `gorm:"column:procedure_concept_id" json:"procedure_concept_id"`
	ProcedureDate            null.Time   `gorm:"column:procedure_date" json:"procedure_date"`
	ProcedureDatetime        time.Time   `gorm:"column:procedure_datetime" json:"procedure_datetime"`
	ProcedureTypeConceptID   int         `gorm:"column:procedure_type_concept_id" json:"procedure_type_concept_id"`
	ModifierConceptID        int         `gorm:"column:modifier_concept_id" json:"modifier_concept_id"`
	Quantity                 null.Int    `gorm:"column:quantity" json:"quantity"`
	ProviderID               null.Int    `gorm:"column:provider_id" json:"provider_id"`
	VisitOccurrenceID        null.Int    `gorm:"column:visit_occurrence_id" json:"visit_occurrence_id"`
	VisitDetailID            null.Int    `gorm:"column:visit_detail_id" json:"visit_detail_id"`
	ProcedureSourceValue     null.String `gorm:"column:procedure_source_value" json:"procedure_source_value"`
	ProcedureSourceConceptID int         `gorm:"column:procedure_source_concept_id" json:"procedure_source_concept_id"`
	ModifierSourceValue      null.String `gorm:"column:modifier_source_value" json:"modifier_source_value"`
}

// TableName sets the insert table name for this struct type
func (p *ProcedureOccurrence) TableName() string {
	return "procedure_occurrence"
}
