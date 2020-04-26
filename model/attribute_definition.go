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

type AttributeDefinition struct {
	AttributeDefinitionID  int         `gorm:"column:attribute_definition_id;primary_key" json:"attribute_definition_id"`
	AttributeName          string      `gorm:"column:attribute_name" json:"attribute_name"`
	AttributeDescription   null.String `gorm:"column:attribute_description" json:"attribute_description"`
	AttributeTypeConceptID int         `gorm:"column:attribute_type_concept_id" json:"attribute_type_concept_id"`
	AttributeSyntax        null.String `gorm:"column:attribute_syntax" json:"attribute_syntax"`
}

// TableName sets the insert table name for this struct type
func (a *AttributeDefinition) TableName() string {
	return "attribute_definition"
}
