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

type DoseEra struct {
	DoseEraID            int64     `gorm:"column:dose_era_id;primary_key" json:"dose_era_id"`
	PersonID             int64     `gorm:"column:person_id" json:"person_id"`
	DrugConceptID        int       `gorm:"column:drug_concept_id" json:"drug_concept_id"`
	UnitConceptID        int       `gorm:"column:unit_concept_id" json:"unit_concept_id"`
	DoseValue            float64   `gorm:"column:dose_value" json:"dose_value"`
	DoseEraStartDatetime time.Time `gorm:"column:dose_era_start_datetime" json:"dose_era_start_datetime"`
	DoseEraEndDatetime   time.Time `gorm:"column:dose_era_end_datetime" json:"dose_era_end_datetime"`
}

// TableName sets the insert table name for this struct type
func (d *DoseEra) TableName() string {
	return "dose_era"
}
