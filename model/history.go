package model

type Historys struct {
	ID     int    `gorm:"primary_key"`
	Name   string `gorm:"type:varchar(255)"`
	History	string	`gorm:"type:varchar(1000)"`
	UserID int    `gorm:"type:int"` 
	User   Users	`gorm:"foreignKey:UserID;references:ID"`
}
