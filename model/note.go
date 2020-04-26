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

type Note struct {
	NoteID                  int64       `gorm:"column:note_id;primary_key" json:"note_id"`
	PersonID                int64       `gorm:"column:person_id" json:"person_id"`
	NoteEventID             null.Int    `gorm:"column:note_event_id" json:"note_event_id"`
	NoteEventFieldConceptID int         `gorm:"column:note_event_field_concept_id" json:"note_event_field_concept_id"`
	NoteDate                null.Time   `gorm:"column:note_date" json:"note_date"`
	NoteDatetime            time.Time   `gorm:"column:note_datetime" json:"note_datetime"`
	NoteTypeConceptID       int         `gorm:"column:note_type_concept_id" json:"note_type_concept_id"`
	NoteClassConceptID      int         `gorm:"column:note_class_concept_id" json:"note_class_concept_id"`
	NoteTitle               null.String `gorm:"column:note_title" json:"note_title"`
	NoteText                null.String `gorm:"column:note_text" json:"note_text"`
	EncodingConceptID       int         `gorm:"column:encoding_concept_id" json:"encoding_concept_id"`
	LanguageConceptID       int         `gorm:"column:language_concept_id" json:"language_concept_id"`
	ProviderID              null.Int    `gorm:"column:provider_id" json:"provider_id"`
	VisitOccurrenceID       null.Int    `gorm:"column:visit_occurrence_id" json:"visit_occurrence_id"`
	VisitDetailID           null.Int    `gorm:"column:visit_detail_id" json:"visit_detail_id"`
	NoteSourceValue         null.String `gorm:"column:note_source_value" json:"note_source_value"`
}

// TableName sets the insert table name for this struct type
func (n *Note) TableName() string {
	return "note"
}
