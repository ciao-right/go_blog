package model

import (
	"time"
)

type BaseModel struct {
	ID         uint      `gorm:"primaryKey;not null;autoIncrement" json:"id" `
	CreatedOn  time.Time `json:"created_on" gorm:"autoCreateTime"`
	ModifiedOn time.Time `json:"modified_on" gorm:"autoUpdateTime"`
}

type PageListModel struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type OverrideTimeModel struct {
	CreatedOn  string `json:"created_on" `
	ModifiedOn string `json:"modified_on"`
}
