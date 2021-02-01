package models

import (
	"time"
)

type AnalyticEventCount struct {
	ID        uint      `json:"id" gorm:"type:bigint(11);primary_key;AUTO_INCREMENT"`
	EventName string    `json:"event_name" gorm:"type:varchar(50);"`
	EventKey  string    `json:"event_key" gorm:"type:varchar(50);"`
	Count     uint64    `json:"count" gorm:"type:bigint(11);"`
	UpdatedAt time.Time `json:"updated_at" gorm:"Column:updated_at"`
}

func (AnalyticEventCount) TableName() string {
	return "analytic_event_counts"
}

func (AnalyticEventCount) Relationships() []string {
	return []string{}
}

func (AnalyticEventCount) Connection() string {
	return "spire"
}

func (AnalyticEventCount) Indexes() map[string][]string {
	return map[string][]string{
		"event_name":     {"event_name"},
		"event_name_key": {"event_name", "event_key"},
	}
}
