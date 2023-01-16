package model

import (
	"time"
)

type BaseModel struct {
	ID         uint       `gorm:"primaryKey;not null;autoIncrement" json:"id" `
	CreatedOn  *time.Time `json:"created_on"`
	ModifiedOn *time.Time `json:"modified_on"`
}
