package models

import (
	"github.com/jinzhu/gorm"
)

//a struct to rep user account
type Project struct {
	gorm.Model
	Name string `json:"name"`
}

func (project *Project) Validate() (map[string]interface{}, bool) {
	return nil, false
}

func (project *Project) Create() map[string]interface{} {
	return nil
}
