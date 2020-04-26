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

type Domain struct {
	DomainID        string `gorm:"column:domain_id;primary_key" json:"domain_id"`
	DomainName      string `gorm:"column:domain_name" json:"domain_name"`
	DomainConceptID int    `gorm:"column:domain_concept_id" json:"domain_concept_id"`
}

// TableName sets the insert table name for this struct type
func (d *Domain) TableName() string {
	return "domain"
}
