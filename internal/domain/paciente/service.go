package paciente

import (
	"context"
	"errors"
	"log"
)

type Service interface {
	GetByID(ctx context.Context, id int) (Paciente, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

// GetByID returns a patient by its ID.
func (s *service) GetByID(ctx context.Context, id int) (Paciente, error) {
	paciente, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("Error en servicio de paciente:", err.Error())
		return Paciente{}, errors.New("error en servicio. Metodo GetByID")
	}

	return paciente, nil
}
