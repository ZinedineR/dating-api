package domain

import "github.com/google/uuid"

const (
	ProfileTableName = "profile"
)

type Profile struct {
	Id          uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Name        string    `gorm:"type:varchar;not_null" json:"name"`
	Age         int       `gorm:"type:int;not_null" json:"age"`
	Sex         string    `gorm:"type:varchar;not_null" json:"sex"`
	Orientation string    `gorm:"type:varchar;not_null" json:"orientation"`
	Status      string    `gorm:"type:varchar;not_null" json:"status"`
	Account     string    `gorm:"type:varchar;not_null" json:"account"`
}

func (model *Profile) TableName() string {
	return ProfileTableName
}
