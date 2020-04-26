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

type Location struct {
	LocationID          int64       `gorm:"column:location_id;primary_key" json:"location_id"`
	Address1            null.String `gorm:"column:address_1" json:"address_1"`
	Address2            null.String `gorm:"column:address_2" json:"address_2"`
	City                null.String `gorm:"column:city" json:"city"`
	State               null.String `gorm:"column:state" json:"state"`
	Zip                 null.String `gorm:"column:zip" json:"zip"`
	County              null.String `gorm:"column:county" json:"county"`
	Country             null.String `gorm:"column:country" json:"country"`
	LocationSourceValue null.String `gorm:"column:location_source_value" json:"location_source_value"`
	Latitude            null.Float  `gorm:"column:latitude" json:"latitude"`
	Longitude           null.Float  `gorm:"column:longitude" json:"longitude"`
}

// TableName sets the insert table name for this struct type
func (l *Location) TableName() string {
	return "location"
}
