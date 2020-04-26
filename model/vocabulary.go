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

type Vocabulary struct {
	VocabularyID        string      `gorm:"column:vocabulary_id;primary_key" json:"vocabulary_id"`
	VocabularyName      string      `gorm:"column:vocabulary_name" json:"vocabulary_name"`
	VocabularyReference string      `gorm:"column:vocabulary_reference" json:"vocabulary_reference"`
	VocabularyVersion   null.String `gorm:"column:vocabulary_version" json:"vocabulary_version"`
	VocabularyConceptID int         `gorm:"column:vocabulary_concept_id" json:"vocabulary_concept_id"`
}

// TableName sets the insert table name for this struct type
func (v *Vocabulary) TableName() string {
	return "vocabulary"
}
