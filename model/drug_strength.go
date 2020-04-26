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

type DrugStrength struct {
	DrugConceptID            int         `gorm:"column:drug_concept_id;primary_key" json:"drug_concept_id"`
	IngredientConceptID      int         `gorm:"column:ingredient_concept_id" json:"ingredient_concept_id"`
	AmountValue              null.Float  `gorm:"column:amount_value" json:"amount_value"`
	AmountUnitConceptID      null.Int    `gorm:"column:amount_unit_concept_id" json:"amount_unit_concept_id"`
	NumeratorValue           null.Float  `gorm:"column:numerator_value" json:"numerator_value"`
	NumeratorUnitConceptID   null.Int    `gorm:"column:numerator_unit_concept_id" json:"numerator_unit_concept_id"`
	DenominatorValue         null.Float  `gorm:"column:denominator_value" json:"denominator_value"`
	DenominatorUnitConceptID null.Int    `gorm:"column:denominator_unit_concept_id" json:"denominator_unit_concept_id"`
	BoxSize                  null.Int    `gorm:"column:box_size" json:"box_size"`
	ValidStartDate           time.Time   `gorm:"column:valid_start_date" json:"valid_start_date"`
	ValidEndDate             time.Time   `gorm:"column:valid_end_date" json:"valid_end_date"`
	InvalidReason            null.String `gorm:"column:invalid_reason" json:"invalid_reason"`
}

// TableName sets the insert table name for this struct type
func (d *DrugStrength) TableName() string {
	return "drug_strength"
}
