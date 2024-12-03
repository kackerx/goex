package server

import (
	"gorm.io/gorm"

	"goex1/internal/data/model"
)

type Migrate struct {
	db *gorm.DB
}

func NewMigrate(db *gorm.DB) *Migrate {
	return &Migrate{db: db}
}

func (m *Migrate) Start() {
	if err := m.db.Migrator().AutoMigrate(&model.User{}); err != nil {
		panic(err)
	}
}
