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

type Relationship struct {
	RelationshipID        string `gorm:"column:relationship_id;primary_key" json:"relationship_id"`
	RelationshipName      string `gorm:"column:relationship_name" json:"relationship_name"`
	IsHierarchical        string `gorm:"column:is_hierarchical" json:"is_hierarchical"`
	DefinesAncestry       string `gorm:"column:defines_ancestry" json:"defines_ancestry"`
	ReverseRelationshipID string `gorm:"column:reverse_relationship_id" json:"reverse_relationship_id"`
	RelationshipConceptID int    `gorm:"column:relationship_concept_id" json:"relationship_concept_id"`
}

// TableName sets the insert table name for this struct type
func (r *Relationship) TableName() string {
	return "relationship"
}
