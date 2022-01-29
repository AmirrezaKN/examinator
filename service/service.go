package service

import "github.com/AmirrezaKN/examinator/domain/models"

type ExaminatorService interface {
	AddChoice(input models.Choice) (id int64)
}
