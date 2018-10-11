package models

import (
	"github.com/jinzhu/gorm"
)

//a struct to rep user account
type Estate struct {
	gorm.Model
	Street string `json:"street"`
	HouseNumber string `json:"housenumber"`
}

func (estate *Estate) Validate() (map[string]interface{}, bool) {
	return nil, false
}

func (estate *Estate) Create() map[string]interface{} {
	return nil
}
