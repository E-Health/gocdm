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

type SurveyConduct struct {
	SurveyConductID             int64       `gorm:"column:survey_conduct_id;primary_key" json:"survey_conduct_id"`
	PersonID                    int64       `gorm:"column:person_id" json:"person_id"`
	SurveyConceptID             int         `gorm:"column:survey_concept_id" json:"survey_concept_id"`
	SurveyStartDate             null.Time   `gorm:"column:survey_start_date" json:"survey_start_date"`
	SurveyStartDatetime         null.Time   `gorm:"column:survey_start_datetime" json:"survey_start_datetime"`
	SurveyEndDate               null.Time   `gorm:"column:survey_end_date" json:"survey_end_date"`
	SurveyEndDatetime           time.Time   `gorm:"column:survey_end_datetime" json:"survey_end_datetime"`
	ProviderID                  null.Int    `gorm:"column:provider_id" json:"provider_id"`
	AssistedConceptID           int         `gorm:"column:assisted_concept_id" json:"assisted_concept_id"`
	RespondentTypeConceptID     int         `gorm:"column:respondent_type_concept_id" json:"respondent_type_concept_id"`
	TimingConceptID             int         `gorm:"column:timing_concept_id" json:"timing_concept_id"`
	CollectionMethodConceptID   int         `gorm:"column:collection_method_concept_id" json:"collection_method_concept_id"`
	AssistedSourceValue         null.String `gorm:"column:assisted_source_value" json:"assisted_source_value"`
	RespondentTypeSourceValue   null.String `gorm:"column:respondent_type_source_value" json:"respondent_type_source_value"`
	TimingSourceValue           null.String `gorm:"column:timing_source_value" json:"timing_source_value"`
	CollectionMethodSourceValue null.String `gorm:"column:collection_method_source_value" json:"collection_method_source_value"`
	SurveySourceValue           null.String `gorm:"column:survey_source_value" json:"survey_source_value"`
	SurveySourceConceptID       int         `gorm:"column:survey_source_concept_id" json:"survey_source_concept_id"`
	SurveySourceIdentifier      null.String `gorm:"column:survey_source_identifier" json:"survey_source_identifier"`
	ValidatedSurveyConceptID    int         `gorm:"column:validated_survey_concept_id" json:"validated_survey_concept_id"`
	ValidatedSurveySourceValue  null.String `gorm:"column:validated_survey_source_value" json:"validated_survey_source_value"`
	SurveyVersionNumber         null.String `gorm:"column:survey_version_number" json:"survey_version_number"`
	VisitOccurrenceID           null.Int    `gorm:"column:visit_occurrence_id" json:"visit_occurrence_id"`
	VisitDetailID               null.Int    `gorm:"column:visit_detail_id" json:"visit_detail_id"`
	ResponseVisitOccurrenceID   null.Int    `gorm:"column:response_visit_occurrence_id" json:"response_visit_occurrence_id"`
}

// TableName sets the insert table name for this struct type
func (s *SurveyConduct) TableName() string {
	return "survey_conduct"
}
