package domain

import (
	"dating-api/pkg/errs"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

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
	Email       string    `gorm:"type:varchar;not_null" json:"email"`
	Password    string    `gorm:"type:varchar;not_null" json:"password"`
	CreatedAt   time.Time `gorm:"type:date;not_null" json:"created_at"`
}

type ProfileData struct {
	Id          uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Name        string    `gorm:"type:varchar;not_null" json:"name"`
	Age         int       `gorm:"type:int;not_null" json:"age"`
	Sex         string    `gorm:"type:varchar;not_null" json:"sex"`
	Orientation string    `gorm:"type:varchar;not_null" json:"orientation"`
	Status      string    `gorm:"type:varchar;not_null" json:"status"`
	Account     string    `gorm:"type:varchar;not_null" json:"account"`
	CreatedAt   time.Time `gorm:"type:date;not_null" json:"created_at"`
}

type ProfileLogin struct {
	Email    string `gorm:"type:varchar;not_null" json:"email"`
	Password string `gorm:"type:varchar;not_null" json:"password"`
}

func (model *Profile) TableName() string {
	return ProfileTableName
}

func (model *Profile) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	model.Password = string(bytes)
	return nil
}
func (model *Profile) CheckPassword(providedPassword string) errs.Error {
	err := bcrypt.CompareHashAndPassword([]byte(model.Password), []byte(providedPassword))
	if err != nil {
		return errs.Wrap(err)
	}
	return nil
}
