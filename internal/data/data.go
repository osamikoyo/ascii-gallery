package data

import (
	"github.com/osamikoyo/ascii-gallery/internal/data/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Storage struct{
	*gorm.DB
}

func Connect() (*Storage, error) {
	db, err := gorm.Open(sqlite.Open("storage/main.db"))
	return &Storage{DB: db}, err
}

func (s *Storage) AddArt(art models.Art) error {
	return s.Create(&art).Error
}