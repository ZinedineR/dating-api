package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const (
	ProfilePreferencesTableName = "profile_preferences"
)

type ProfilePreferences struct {
	Id        uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
	PeopleId  uuid.UUID      `gorm:"type:uuid" json:"people_id"`
	Like      pq.StringArray `gorm:"type:text[];null" json:"like"`
	Pass      pq.StringArray `gorm:"type:text[];null" json:"pass"`
	View      int            `gorm:"type:int;null" json:"view"`
	LastLogin time.Time      `gorm:"type:date;not_null" json:"last_login"`
	Profile   *Profile       `gorm:"foreignKey:PeopleId"`
}

func (model *ProfilePreferences) TableName() string {
	return ProfilePreferencesTableName
}
