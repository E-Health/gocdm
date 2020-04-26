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

type CareSite struct {
	CareSiteID                int64       `gorm:"column:care_site_id;primary_key" json:"care_site_id"`
	CareSiteName              null.String `gorm:"column:care_site_name" json:"care_site_name"`
	PlaceOfServiceConceptID   int         `gorm:"column:place_of_service_concept_id" json:"place_of_service_concept_id"`
	LocationID                null.Int    `gorm:"column:location_id" json:"location_id"`
	CareSiteSourceValue       null.String `gorm:"column:care_site_source_value" json:"care_site_source_value"`
	PlaceOfServiceSourceValue null.String `gorm:"column:place_of_service_source_value" json:"place_of_service_source_value"`
}

// TableName sets the insert table name for this struct type
func (c *CareSite) TableName() string {
	return "care_site"
}
