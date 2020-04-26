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

type ConditionOccurrence struct {
	ConditionOccurrenceID      int64       `gorm:"column:condition_occurrence_id;primary_key" json:"condition_occurrence_id"`
	PersonID                   int64       `gorm:"column:person_id" json:"person_id"`
	ConditionConceptID         int         `gorm:"column:condition_concept_id" json:"condition_concept_id"`
	ConditionStartDate         null.Time   `gorm:"column:condition_start_date" json:"condition_start_date"`
	ConditionStartDatetime     time.Time   `gorm:"column:condition_start_datetime" json:"condition_start_datetime"`
	ConditionEndDate           null.Time   `gorm:"column:condition_end_date" json:"condition_end_date"`
	ConditionEndDatetime       null.Time   `gorm:"column:condition_end_datetime" json:"condition_end_datetime"`
	ConditionTypeConceptID     int         `gorm:"column:condition_type_concept_id" json:"condition_type_concept_id"`
	ConditionStatusConceptID   int         `gorm:"column:condition_status_concept_id" json:"condition_status_concept_id"`
	StopReason                 null.String `gorm:"column:stop_reason" json:"stop_reason"`
	ProviderID                 null.Int    `gorm:"column:provider_id" json:"provider_id"`
	VisitOccurrenceID          null.Int    `gorm:"column:visit_occurrence_id" json:"visit_occurrence_id"`
	VisitDetailID              null.Int    `gorm:"column:visit_detail_id" json:"visit_detail_id"`
	ConditionSourceValue       null.String `gorm:"column:condition_source_value" json:"condition_source_value"`
	ConditionSourceConceptID   int         `gorm:"column:condition_source_concept_id" json:"condition_source_concept_id"`
	ConditionStatusSourceValue null.String `gorm:"column:condition_status_source_value" json:"condition_status_source_value"`
}

// TableName sets the insert table name for this struct type
func (c *ConditionOccurrence) TableName() string {
	return "condition_occurrence"
}
