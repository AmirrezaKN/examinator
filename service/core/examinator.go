package core

import (
	"github.com/AmirrezaKN/examinator/domain/models"
	"github.com/AmirrezaKN/examinator/service"
	"github.com/AmirrezaKN/examinator/service/ports"
)

type examinatorService struct {
	repo ports.RepositoryPort
}

func NewExaminatorService(repo ports.RepositoryPort) service.ExaminatorService {
	return &examinatorService{
		repo: repo,
	}
}

func (s examinatorService) AddChoice(input models.Choice) int64 {
	return -1
}
