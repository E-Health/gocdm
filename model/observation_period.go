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

type ObservationPeriod struct {
	ObservationPeriodID        int64     `gorm:"column:observation_period_id;primary_key" json:"observation_period_id"`
	PersonID                   int64     `gorm:"column:person_id" json:"person_id"`
	ObservationPeriodStartDate time.Time `gorm:"column:observation_period_start_date" json:"observation_period_start_date"`
	ObservationPeriodEndDate   time.Time `gorm:"column:observation_period_end_date" json:"observation_period_end_date"`
	PeriodTypeConceptID        int       `gorm:"column:period_type_concept_id" json:"period_type_concept_id"`
}

// TableName sets the insert table name for this struct type
func (o *ObservationPeriod) TableName() string {
	return "observation_period"
}
