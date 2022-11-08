package model

import (
	"gorm.io/gorm"
	"time"
)

type Transacao struct {
	Tipo      string         `json:"tipo"`
	Quantia   float64        `json:"quantia"`
	Conta     int            `json:"conta"`
	Moeda     string         `json:"moeda"`
	CreatedAt time.Time      `json:"created"`
	UpdatedAt time.Time      `json:"UpdatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
