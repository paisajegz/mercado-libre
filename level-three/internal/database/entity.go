package db

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint64    `gorm:"column:id;primary_key;auto_increment;"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;" json:"updated_at"`
}

type DNARegister struct {
	Model
	IsMutant bool   `gorm:"column:is_mutant;not null;" json:"is_mutant"`
	Dna      string `gorm:"column:dna;not null;" json:"dna"`
}

func (m *DNARegister) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m *DNARegister) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}

type Stat struct {
	IsMutant bool
	Count    int
}
