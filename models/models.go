package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	ID           string    `json:"ID" gorm:"text;not null;default:null`
	Name         string    `json:"Name" gorm:"text;not null;default:null`
	Salary       string    `json:"Salary" gorm:"text;not null;default:null`
	Technologies string    `json:"Technologies" gorm:"text;not null;default:null`
	Projects     *Projects `json:"Projects" gorm:"text;not null;default:null`
	Manager      *Manager  `json:"Manager" gorm:"text;not null;default:null`
}

type Projects struct {
	ID                  string `json:"ID" gorm:"text;not null;default:null`
	ProjectName         string `json:"Projectname" gorm:"text;not null;default:null`
	Billing             string `json:"Billing" gorm:"text;not null;default:null`
	Duration            string `json:"Duration" gorm:"text;not null;default:null`
	NoOfEmplyoeeWorking int    `json:"NoOfEmployeeWorking" gorm:"text;not null;default:null`
}

type Manager struct {
	ID   string `json:"ID" gorm:"text;not null;default:null`
	Name string `json:"Name" gorm:"text;not null;default:null`
}
