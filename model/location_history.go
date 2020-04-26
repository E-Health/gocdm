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
// LocationHistory : LocationHistory
type LocationHistory struct {
	LocationHistoryID         int64     `gorm:"column:location_history_id;primary_key" json:"location_history_id"`
	LocationID                int64     `gorm:"column:location_id" json:"location_id"`
	RelationshipTypeConceptID int       `gorm:"column:relationship_type_concept_id" json:"relationship_type_concept_id"`
	DomainID                  string    `gorm:"column:domain_id" json:"domain_id"`
	EntityID                  int64     `gorm:"column:entity_id" json:"entity_id"`
	StartDate                 time.Time `gorm:"column:start_date" json:"start_date"`
	EndDate                   null.Time `gorm:"column:end_date" json:"end_date"`
}

// TableName sets the insert table name for this struct type
func (l *LocationHistory) TableName() string {
	return "location_history"
}
