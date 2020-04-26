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

type ConceptRelationship struct {
	ConceptID1     int         `gorm:"column:concept_id_1;primary_key" json:"concept_id_1"`
	ConceptID2     int         `gorm:"column:concept_id_2" json:"concept_id_2"`
	RelationshipID string      `gorm:"column:relationship_id" json:"relationship_id"`
	ValidStartDate time.Time   `gorm:"column:valid_start_date" json:"valid_start_date"`
	ValidEndDate   time.Time   `gorm:"column:valid_end_date" json:"valid_end_date"`
	InvalidReason  null.String `gorm:"column:invalid_reason" json:"invalid_reason"`
}

// TableName sets the insert table name for this struct type
func (c *ConceptRelationship) TableName() string {
	return "concept_relationship"
}
