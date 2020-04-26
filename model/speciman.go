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

type Speciman struct {
	SpecimenID               int64       `gorm:"column:specimen_id;primary_key" json:"specimen_id"`
	PersonID                 int64       `gorm:"column:person_id" json:"person_id"`
	SpecimenConceptID        int         `gorm:"column:specimen_concept_id" json:"specimen_concept_id"`
	SpecimenTypeConceptID    int         `gorm:"column:specimen_type_concept_id" json:"specimen_type_concept_id"`
	SpecimenDate             null.Time   `gorm:"column:specimen_date" json:"specimen_date"`
	SpecimenDatetime         time.Time   `gorm:"column:specimen_datetime" json:"specimen_datetime"`
	Quantity                 null.Float  `gorm:"column:quantity" json:"quantity"`
	UnitConceptID            null.Int    `gorm:"column:unit_concept_id" json:"unit_concept_id"`
	AnatomicSiteConceptID    int         `gorm:"column:anatomic_site_concept_id" json:"anatomic_site_concept_id"`
	DiseaseStatusConceptID   int         `gorm:"column:disease_status_concept_id" json:"disease_status_concept_id"`
	SpecimenSourceID         null.String `gorm:"column:specimen_source_id" json:"specimen_source_id"`
	SpecimenSourceValue      null.String `gorm:"column:specimen_source_value" json:"specimen_source_value"`
	UnitSourceValue          null.String `gorm:"column:unit_source_value" json:"unit_source_value"`
	AnatomicSiteSourceValue  null.String `gorm:"column:anatomic_site_source_value" json:"anatomic_site_source_value"`
	DiseaseStatusSourceValue null.String `gorm:"column:disease_status_source_value" json:"disease_status_source_value"`
}

// TableName sets the insert table name for this struct type
func (s *Speciman) TableName() string {
	return "specimen"
}
