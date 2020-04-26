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

type Observation struct {
	ObservationID              int64       `gorm:"column:observation_id;primary_key" json:"observation_id"`
	PersonID                   int64       `gorm:"column:person_id" json:"person_id"`
	ObservationConceptID       int         `gorm:"column:observation_concept_id" json:"observation_concept_id"`
	ObservationDate            null.Time   `gorm:"column:observation_date" json:"observation_date"`
	ObservationDatetime        time.Time   `gorm:"column:observation_datetime" json:"observation_datetime"`
	ObservationTypeConceptID   int         `gorm:"column:observation_type_concept_id" json:"observation_type_concept_id"`
	ValueAsNumber              null.Float  `gorm:"column:value_as_number" json:"value_as_number"`
	ValueAsString              null.String `gorm:"column:value_as_string" json:"value_as_string"`
	ValueAsConceptID           null.Int    `gorm:"column:value_as_concept_id" json:"value_as_concept_id"`
	QualifierConceptID         null.Int    `gorm:"column:qualifier_concept_id" json:"qualifier_concept_id"`
	UnitConceptID              null.Int    `gorm:"column:unit_concept_id" json:"unit_concept_id"`
	ProviderID                 null.Int    `gorm:"column:provider_id" json:"provider_id"`
	VisitOccurrenceID          null.Int    `gorm:"column:visit_occurrence_id" json:"visit_occurrence_id"`
	VisitDetailID              null.Int    `gorm:"column:visit_detail_id" json:"visit_detail_id"`
	ObservationSourceValue     null.String `gorm:"column:observation_source_value" json:"observation_source_value"`
	ObservationSourceConceptID int         `gorm:"column:observation_source_concept_id" json:"observation_source_concept_id"`
	UnitSourceValue            null.String `gorm:"column:unit_source_value" json:"unit_source_value"`
	QualifierSourceValue       null.String `gorm:"column:qualifier_source_value" json:"qualifier_source_value"`
	ObservationEventID         null.Int    `gorm:"column:observation_event_id" json:"observation_event_id"`
	ObsEventFieldConceptID     int         `gorm:"column:obs_event_field_concept_id" json:"obs_event_field_concept_id"`
	ValueAsDatetime            null.Time   `gorm:"column:value_as_datetime" json:"value_as_datetime"`
}

// TableName sets the insert table name for this struct type
func (o *Observation) TableName() string {
	return "observation"
}
