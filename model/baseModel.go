package model

import "time"

type baseModel struct {
	ID         uint      `gorm:"not null;primarykey" json:"id" `
	CreatedOn  time.Time `json:"created_on"`
	ModifiedOn time.Time `json:"modified_on"`
}
