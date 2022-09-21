package models

type Role struct{
	Name string `json:"name" gorm:"varchar;not_null;unique"`
	
}