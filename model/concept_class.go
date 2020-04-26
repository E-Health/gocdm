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

type ConceptClass struct {
	ConceptClassID        string `gorm:"column:concept_class_id;primary_key" json:"concept_class_id"`
	ConceptClassName      string `gorm:"column:concept_class_name" json:"concept_class_name"`
	ConceptClassConceptID int    `gorm:"column:concept_class_concept_id" json:"concept_class_concept_id"`
}

// TableName sets the insert table name for this struct type
func (c *ConceptClass) TableName() string {
	return "concept_class"
}
