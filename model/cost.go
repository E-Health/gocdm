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

type Cost struct {
	CostID                  int64       `gorm:"column:cost_id;primary_key" json:"cost_id"`
	PersonID                int64       `gorm:"column:person_id" json:"person_id"`
	CostEventID             int64       `gorm:"column:cost_event_id" json:"cost_event_id"`
	CostEventFieldConceptID int         `gorm:"column:cost_event_field_concept_id" json:"cost_event_field_concept_id"`
	CostConceptID           int         `gorm:"column:cost_concept_id" json:"cost_concept_id"`
	CostTypeConceptID       int         `gorm:"column:cost_type_concept_id" json:"cost_type_concept_id"`
	CurrencyConceptID       int         `gorm:"column:currency_concept_id" json:"currency_concept_id"`
	Cost                    null.Float  `gorm:"column:cost" json:"cost"`
	IncurredDate            time.Time   `gorm:"column:incurred_date" json:"incurred_date"`
	BilledDate              null.Time   `gorm:"column:billed_date" json:"billed_date"`
	PaidDate                null.Time   `gorm:"column:paid_date" json:"paid_date"`
	RevenueCodeConceptID    int         `gorm:"column:revenue_code_concept_id" json:"revenue_code_concept_id"`
	DrgConceptID            int         `gorm:"column:drg_concept_id" json:"drg_concept_id"`
	CostSourceValue         null.String `gorm:"column:cost_source_value" json:"cost_source_value"`
	CostSourceConceptID     int         `gorm:"column:cost_source_concept_id" json:"cost_source_concept_id"`
	RevenueCodeSourceValue  null.String `gorm:"column:revenue_code_source_value" json:"revenue_code_source_value"`
	DrgSourceValue          null.String `gorm:"column:drg_source_value" json:"drg_source_value"`
	PayerPlanPeriodID       null.Int    `gorm:"column:payer_plan_period_id" json:"payer_plan_period_id"`
}

// TableName sets the insert table name for this struct type
func (c *Cost) TableName() string {
	return "cost"
}
