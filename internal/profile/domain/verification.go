package domain

import (
	"time"

	"github.com/google/uuid"
)

const (
	VerificationTableName = "verification"
)

// VerificationModel is a model for entity.Verification
type Verification struct {
	Id        uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	PeopleId  uuid.UUID `gorm:"type:uuid;not_null" json:"people_id"`
	Expiresat time.Time `gorm:"type:date;not_null" json:"expiresat"`
	Verified  bool      `gorm:"default:false;null" json:"verified"`
	Profile   *Profile  `gorm:"foreignKey:PeopleId"`
}

func (model *Verification) TableName() string {
	return VerificationTableName
}
