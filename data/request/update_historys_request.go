package request

import "github.com/c0utin/historinha/model"

type UpdateHistoryRequest struct {
	ID     int    `gorm:"primary_key"`
	Name   string `gorm:"type:varchar(255)"`
	History	string	`gorm:"type:varchar(1000)"`
	UserID int    `gorm:"type:int"` 
	User   model.Users	`gorm:"foreignKey:UserID;references:ID"`
}
