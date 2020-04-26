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

type FactRelationship struct {
	DomainConceptID1      int   `gorm:"column:domain_concept_id_1;primary_key" json:"domain_concept_id_1"`
	FactID1               int64 `gorm:"column:fact_id_1" json:"fact_id_1"`
	DomainConceptID2      int   `gorm:"column:domain_concept_id_2" json:"domain_concept_id_2"`
	FactID2               int64 `gorm:"column:fact_id_2" json:"fact_id_2"`
	RelationshipConceptID int   `gorm:"column:relationship_concept_id" json:"relationship_concept_id"`
}

// TableName sets the insert table name for this struct type
func (f *FactRelationship) TableName() string {
	return "fact_relationship"
}
