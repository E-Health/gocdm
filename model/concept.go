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

type Concept struct {
	ConceptID       int         `gorm:"column:concept_id;primary_key" json:"concept_id"`
	ConceptName     string      `gorm:"column:concept_name" json:"concept_name"`
	DomainID        string      `gorm:"column:domain_id" json:"domain_id"`
	VocabularyID    string      `gorm:"column:vocabulary_id" json:"vocabulary_id"`
	ConceptClassID  string      `gorm:"column:concept_class_id" json:"concept_class_id"`
	StandardConcept null.String `gorm:"column:standard_concept" json:"standard_concept"`
	ConceptCode     string      `gorm:"column:concept_code" json:"concept_code"`
	ValidStartDate  time.Time   `gorm:"column:valid_start_date" json:"valid_start_date"`
	ValidEndDate    time.Time   `gorm:"column:valid_end_date" json:"valid_end_date"`
	InvalidReason   null.String `gorm:"column:invalid_reason" json:"invalid_reason"`
}

// TableName sets the insert table name for this struct type
func (c *Concept) TableName() string {
	return "concept"
}
