package entity


import "gorm.io/gorm"


type Stats struct {

   gorm.Model

   Status string `json:"stat"`

}