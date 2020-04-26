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

type Measurement struct {
	MeasurementID              int64       `gorm:"column:measurement_id;primary_key" json:"measurement_id"`
	PersonID                   int64       `gorm:"column:person_id" json:"person_id"`
	MeasurementConceptID       int         `gorm:"column:measurement_concept_id" json:"measurement_concept_id"`
	MeasurementDate            null.Time   `gorm:"column:measurement_date" json:"measurement_date"`
	MeasurementDatetime        time.Time   `gorm:"column:measurement_datetime" json:"measurement_datetime"`
	MeasurementTime            null.String `gorm:"column:measurement_time" json:"measurement_time"`
	MeasurementTypeConceptID   int         `gorm:"column:measurement_type_concept_id" json:"measurement_type_concept_id"`
	OperatorConceptID          null.Int    `gorm:"column:operator_concept_id" json:"operator_concept_id"`
	ValueAsNumber              null.Float  `gorm:"column:value_as_number" json:"value_as_number"`
	ValueAsConceptID           null.Int    `gorm:"column:value_as_concept_id" json:"value_as_concept_id"`
	UnitConceptID              null.Int    `gorm:"column:unit_concept_id" json:"unit_concept_id"`
	RangeLow                   null.Float  `gorm:"column:range_low" json:"range_low"`
	RangeHigh                  null.Float  `gorm:"column:range_high" json:"range_high"`
	ProviderID                 null.Int    `gorm:"column:provider_id" json:"provider_id"`
	VisitOccurrenceID          null.Int    `gorm:"column:visit_occurrence_id" json:"visit_occurrence_id"`
	VisitDetailID              null.Int    `gorm:"column:visit_detail_id" json:"visit_detail_id"`
	MeasurementSourceValue     null.String `gorm:"column:measurement_source_value" json:"measurement_source_value"`
	MeasurementSourceConceptID int         `gorm:"column:measurement_source_concept_id" json:"measurement_source_concept_id"`
	UnitSourceValue            null.String `gorm:"column:unit_source_value" json:"unit_source_value"`
	ValueSourceValue           null.String `gorm:"column:value_source_value" json:"value_source_value"`
}

// TableName sets the insert table name for this struct type
func (m *Measurement) TableName() string {
	return "measurement"
}
