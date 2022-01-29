package ports

import "github.com/AmirrezaKN/examinator/domain/models"

type RepositoryPort interface {
	AddChoice(input models.Choice) (id int64)
}
