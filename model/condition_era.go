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

type ConditionEra struct {
	ConditionEraID            int64     `gorm:"column:condition_era_id;primary_key" json:"condition_era_id"`
	PersonID                  int64     `gorm:"column:person_id" json:"person_id"`
	ConditionConceptID        int       `gorm:"column:condition_concept_id" json:"condition_concept_id"`
	ConditionEraStartDatetime time.Time `gorm:"column:condition_era_start_datetime" json:"condition_era_start_datetime"`
	ConditionEraEndDatetime   time.Time `gorm:"column:condition_era_end_datetime" json:"condition_era_end_datetime"`
	ConditionOccurrenceCount  null.Int  `gorm:"column:condition_occurrence_count" json:"condition_occurrence_count"`
}

// TableName sets the insert table name for this struct type
func (c *ConditionEra) TableName() string {
	return "condition_era"
}
