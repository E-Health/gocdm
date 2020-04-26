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

type NoteNlp struct {
	NoteNlpID              int64       `gorm:"column:note_nlp_id;primary_key" json:"note_nlp_id"`
	NoteID                 int64       `gorm:"column:note_id" json:"note_id"`
	SectionConceptID       int         `gorm:"column:section_concept_id" json:"section_concept_id"`
	Snippet                null.String `gorm:"column:snippet" json:"snippet"`
	Offset                 null.String `gorm:"column:offset" json:"offset"`
	LexicalVariant         string      `gorm:"column:lexical_variant" json:"lexical_variant"`
	NoteNlpConceptID       int         `gorm:"column:note_nlp_concept_id" json:"note_nlp_concept_id"`
	NlpSystem              null.String `gorm:"column:nlp_system" json:"nlp_system"`
	NlpDate                time.Time   `gorm:"column:nlp_date" json:"nlp_date"`
	NlpDatetime            null.Time   `gorm:"column:nlp_datetime" json:"nlp_datetime"`
	TermExists             null.String `gorm:"column:term_exists" json:"term_exists"`
	TermTemporal           null.String `gorm:"column:term_temporal" json:"term_temporal"`
	TermModifiers          null.String `gorm:"column:term_modifiers" json:"term_modifiers"`
	NoteNlpSourceConceptID int         `gorm:"column:note_nlp_source_concept_id" json:"note_nlp_source_concept_id"`
}

// TableName sets the insert table name for this struct type
func (n *NoteNlp) TableName() string {
	return "note_nlp"
}
