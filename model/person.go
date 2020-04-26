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

type Person struct {
	PersonID                 int64       `gorm:"column:person_id;primary_key" json:"person_id"`
	GenderConceptID          int         `gorm:"column:gender_concept_id" json:"gender_concept_id"`
	YearOfBirth              int         `gorm:"column:year_of_birth" json:"year_of_birth"`
	MonthOfBirth             null.Int    `gorm:"column:month_of_birth" json:"month_of_birth"`
	DayOfBirth               null.Int    `gorm:"column:day_of_birth" json:"day_of_birth"`
	BirthDatetime            null.Time   `gorm:"column:birth_datetime" json:"birth_datetime"`
	DeathDatetime            null.Time   `gorm:"column:death_datetime" json:"death_datetime"`
	RaceConceptID            int         `gorm:"column:race_concept_id" json:"race_concept_id"`
	EthnicityConceptID       int         `gorm:"column:ethnicity_concept_id" json:"ethnicity_concept_id"`
	LocationID               null.Int    `gorm:"column:location_id" json:"location_id"`
	ProviderID               null.Int    `gorm:"column:provider_id" json:"provider_id"`
	CareSiteID               null.Int    `gorm:"column:care_site_id" json:"care_site_id"`
	PersonSourceValue        null.String `gorm:"column:person_source_value" json:"person_source_value"`
	GenderSourceValue        null.String `gorm:"column:gender_source_value" json:"gender_source_value"`
	GenderSourceConceptID    int         `gorm:"column:gender_source_concept_id" json:"gender_source_concept_id"`
	RaceSourceValue          null.String `gorm:"column:race_source_value" json:"race_source_value"`
	RaceSourceConceptID      int         `gorm:"column:race_source_concept_id" json:"race_source_concept_id"`
	EthnicitySourceValue     null.String `gorm:"column:ethnicity_source_value" json:"ethnicity_source_value"`
	EthnicitySourceConceptID int         `gorm:"column:ethnicity_source_concept_id" json:"ethnicity_source_concept_id"`
}

// TableName sets the insert table name for this struct type
func (p *Person) TableName() string {
	return "person"
}
