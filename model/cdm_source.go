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

type CdmSource struct {
	CdmSourceName                string      `gorm:"column:cdm_source_name;primary_key" json:"cdm_source_name"`
	CdmSourceAbbreviation        null.String `gorm:"column:cdm_source_abbreviation" json:"cdm_source_abbreviation"`
	CdmHolder                    null.String `gorm:"column:cdm_holder" json:"cdm_holder"`
	SourceDescription            null.String `gorm:"column:source_description" json:"source_description"`
	SourceDocumentationReference null.String `gorm:"column:source_documentation_reference" json:"source_documentation_reference"`
	CdmEtlReference              null.String `gorm:"column:cdm_etl_reference" json:"cdm_etl_reference"`
	SourceReleaseDate            null.Time   `gorm:"column:source_release_date" json:"source_release_date"`
	CdmReleaseDate               null.Time   `gorm:"column:cdm_release_date" json:"cdm_release_date"`
	CdmVersion                   null.String `gorm:"column:cdm_version" json:"cdm_version"`
	VocabularyVersion            null.String `gorm:"column:vocabulary_version" json:"vocabulary_version"`
}

// TableName sets the insert table name for this struct type
func (c *CdmSource) TableName() string {
	return "cdm_source"
}
