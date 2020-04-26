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

type DrugEra struct {
	DrugEraID            int64     `gorm:"column:drug_era_id;primary_key" json:"drug_era_id"`
	PersonID             int64     `gorm:"column:person_id" json:"person_id"`
	DrugConceptID        int       `gorm:"column:drug_concept_id" json:"drug_concept_id"`
	DrugEraStartDatetime time.Time `gorm:"column:drug_era_start_datetime" json:"drug_era_start_datetime"`
	DrugEraEndDatetime   time.Time `gorm:"column:drug_era_end_datetime" json:"drug_era_end_datetime"`
	DrugExposureCount    null.Int  `gorm:"column:drug_exposure_count" json:"drug_exposure_count"`
	GapDays              null.Int  `gorm:"column:gap_days" json:"gap_days"`
}

// TableName sets the insert table name for this struct type
func (d *DrugEra) TableName() string {
	return "drug_era"
}
