package repository

import (
	"database/sql"

	"github.com/AmirrezaKN/examinator/domain/models"
	"github.com/AmirrezaKN/examinator/service/ports"
)

type pgRepository struct {
	db *sql.DB
}

func NewPGRepository(db *sql.DB) ports.RepositoryPort {
	return pgRepository{
		db: db,
	}
}

func (r pgRepository) AddChoice(input models.Choice) int64 {
	return -1
}
