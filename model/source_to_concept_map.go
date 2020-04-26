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

type SourceToConceptMap struct {
	SourceCode            string      `gorm:"column:source_code;primary_key" json:"source_code"`
	SourceConceptID       int         `gorm:"column:source_concept_id" json:"source_concept_id"`
	SourceVocabularyID    string      `gorm:"column:source_vocabulary_id" json:"source_vocabulary_id"`
	SourceCodeDescription null.String `gorm:"column:source_code_description" json:"source_code_description"`
	TargetConceptID       int         `gorm:"column:target_concept_id" json:"target_concept_id"`
	TargetVocabularyID    string      `gorm:"column:target_vocabulary_id" json:"target_vocabulary_id"`
	ValidStartDate        time.Time   `gorm:"column:valid_start_date" json:"valid_start_date"`
	ValidEndDate          time.Time   `gorm:"column:valid_end_date" json:"valid_end_date"`
	InvalidReason         null.String `gorm:"column:invalid_reason" json:"invalid_reason"`
}

// TableName sets the insert table name for this struct type
func (s *SourceToConceptMap) TableName() string {
	return "source_to_concept_map"
}
