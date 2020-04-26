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

type ConceptAncestor struct {
	AncestorConceptID     int `gorm:"column:ancestor_concept_id;primary_key" json:"ancestor_concept_id"`
	DescendantConceptID   int `gorm:"column:descendant_concept_id" json:"descendant_concept_id"`
	MinLevelsOfSeparation int `gorm:"column:min_levels_of_separation" json:"min_levels_of_separation"`
	MaxLevelsOfSeparation int `gorm:"column:max_levels_of_separation" json:"max_levels_of_separation"`
}

// TableName sets the insert table name for this struct type
func (c *ConceptAncestor) TableName() string {
	return "concept_ancestor"
}
