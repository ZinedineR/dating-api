package domain

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

const (
	ProfilePreferencesTableName = "profile_preferences"
)

type ProfilePreferences struct {
	Id       uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
	PeopleId uuid.UUID      `gorm:"type:uuid" json:"people_id"`
	Like     pq.StringArray `gorm:"type:uuid[];null" json:"like"`
	Pass     pq.StringArray `gorm:"type:uuid[];null" json:"pass"`
	View     int            `gorm:"type:int;null" json:"view"`
	Jwt      string         `gorm:"type:varchar;not_null" json:"jwt"`
	Profile  *Profile       `gorm:"foreignKey:PeopleId"`
}

type ProfilePreferencesLikePass struct {
	List string `gorm:"type:varchar;null" json:"list"`
}

func (model *ProfilePreferences) TableName() string {
	return ProfilePreferencesTableName
}
