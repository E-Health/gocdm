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

type Provider struct {
	ProviderID               int64       `gorm:"column:provider_id;primary_key" json:"provider_id"`
	ProviderName             null.String `gorm:"column:provider_name" json:"provider_name"`
	NPI                      null.String `gorm:"column:NPI" json:"NPI"`
	DEA                      null.String `gorm:"column:DEA" json:"DEA"`
	SpecialtyConceptID       int         `gorm:"column:specialty_concept_id" json:"specialty_concept_id"`
	CareSiteID               null.Int    `gorm:"column:care_site_id" json:"care_site_id"`
	YearOfBirth              null.Int    `gorm:"column:year_of_birth" json:"year_of_birth"`
	GenderConceptID          int         `gorm:"column:gender_concept_id" json:"gender_concept_id"`
	ProviderSourceValue      null.String `gorm:"column:provider_source_value" json:"provider_source_value"`
	SpecialtySourceValue     null.String `gorm:"column:specialty_source_value" json:"specialty_source_value"`
	SpecialtySourceConceptID null.Int    `gorm:"column:specialty_source_concept_id" json:"specialty_source_concept_id"`
	GenderSourceValue        null.String `gorm:"column:gender_source_value" json:"gender_source_value"`
	GenderSourceConceptID    int         `gorm:"column:gender_source_concept_id" json:"gender_source_concept_id"`
}

// TableName sets the insert table name for this struct type
func (p *Provider) TableName() string {
	return "provider"
}
