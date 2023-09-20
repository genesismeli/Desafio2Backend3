package turno

import (
	"context"
	"errors"
	"log"
)

type Service interface {
	GetByID(ctx context.Context, id int) (Turno, error)
	//Create(ctx context.Context, requestTurno RequestTurno) (Turno, error)
	//Update(ctx context.Context, requestTurno RequestTurno, id int) (Turno, error)
	//Delete(ctx context.Context, id int) error
	//UpdateField(ctx context.Context, requestTurno2 RequestTurno2, id int) (Turno, error)
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
func (s *service) GetByID(ctx context.Context, id int) (Turno, error) {
	turno, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("Error en servicio de turnos:", err.Error())
		return Turno{}, errors.New("error en servicio. Metodo GetByID")
	}

	return turno, nil
}
