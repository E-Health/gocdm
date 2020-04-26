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

type Metadatum struct {
	MetadataConceptID     int         `gorm:"column:metadata_concept_id;primary_key" json:"metadata_concept_id"`
	MetadataTypeConceptID int         `gorm:"column:metadata_type_concept_id" json:"metadata_type_concept_id"`
	Name                  string      `gorm:"column:name" json:"name"`
	ValueAsString         null.String `gorm:"column:value_as_string" json:"value_as_string"`
	ValueAsConceptID      null.Int    `gorm:"column:value_as_concept_id" json:"value_as_concept_id"`
	MetadataDate          null.Time   `gorm:"column:metadata_date" json:"metadata_date"`
	MetadataDatetime      null.Time   `gorm:"column:metadata_datetime" json:"metadata_datetime"`
}

// TableName sets the insert table name for this struct type
func (m *Metadatum) TableName() string {
	return "metadata"
}
