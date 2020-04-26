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

type ConceptSynonym struct {
	ConceptID          int    `gorm:"column:concept_id;primary_key" json:"concept_id"`
	ConceptSynonymName string `gorm:"column:concept_synonym_name" json:"concept_synonym_name"`
	LanguageConceptID  int    `gorm:"column:language_concept_id" json:"language_concept_id"`
}

// TableName sets the insert table name for this struct type
func (c *ConceptSynonym) TableName() string {
	return "concept_synonym"
}
