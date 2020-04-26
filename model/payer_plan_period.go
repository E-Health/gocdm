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

type PayerPlanPeriod struct {
	PayerPlanPeriodID         int64       `gorm:"column:payer_plan_period_id;primary_key" json:"payer_plan_period_id"`
	PersonID                  int64       `gorm:"column:person_id" json:"person_id"`
	ContractPersonID          null.Int    `gorm:"column:contract_person_id" json:"contract_person_id"`
	PayerPlanPeriodStartDate  time.Time   `gorm:"column:payer_plan_period_start_date" json:"payer_plan_period_start_date"`
	PayerPlanPeriodEndDate    time.Time   `gorm:"column:payer_plan_period_end_date" json:"payer_plan_period_end_date"`
	PayerConceptID            int         `gorm:"column:payer_concept_id" json:"payer_concept_id"`
	PlanConceptID             int         `gorm:"column:plan_concept_id" json:"plan_concept_id"`
	ContractConceptID         int         `gorm:"column:contract_concept_id" json:"contract_concept_id"`
	SponsorConceptID          int         `gorm:"column:sponsor_concept_id" json:"sponsor_concept_id"`
	StopReasonConceptID       int         `gorm:"column:stop_reason_concept_id" json:"stop_reason_concept_id"`
	PayerSourceValue          null.String `gorm:"column:payer_source_value" json:"payer_source_value"`
	PayerSourceConceptID      int         `gorm:"column:payer_source_concept_id" json:"payer_source_concept_id"`
	PlanSourceValue           null.String `gorm:"column:plan_source_value" json:"plan_source_value"`
	PlanSourceConceptID       int         `gorm:"column:plan_source_concept_id" json:"plan_source_concept_id"`
	ContractSourceValue       null.String `gorm:"column:contract_source_value" json:"contract_source_value"`
	ContractSourceConceptID   int         `gorm:"column:contract_source_concept_id" json:"contract_source_concept_id"`
	SponsorSourceValue        null.String `gorm:"column:sponsor_source_value" json:"sponsor_source_value"`
	SponsorSourceConceptID    int         `gorm:"column:sponsor_source_concept_id" json:"sponsor_source_concept_id"`
	FamilySourceValue         null.String `gorm:"column:family_source_value" json:"family_source_value"`
	StopReasonSourceValue     null.String `gorm:"column:stop_reason_source_value" json:"stop_reason_source_value"`
	StopReasonSourceConceptID int         `gorm:"column:stop_reason_source_concept_id" json:"stop_reason_source_concept_id"`
}

// TableName sets the insert table name for this struct type
func (p *PayerPlanPeriod) TableName() string {
	return "payer_plan_period"
}
