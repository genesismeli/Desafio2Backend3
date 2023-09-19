package turno

import (
	"context"
	"errors"
	"fmt"
	"log"
)

type Service interface {
	GetByID(ctx context.Context, id int) (Turno, error)
	Create(ctx context.Context, requestTurno RequestTurno) (Turno, error)
	Update(ctx context.Context, requestTurno RequestTurno, id int) (Turno, error)
	Delete(ctx context.Context, id int) error
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

func (s *service) Create(ctx context.Context, requestTurno RequestTurno) (Turno, error) {
	turno := requestToTurno(requestTurno)
	turnoResponse, err := s.repository.Create(ctx, turno)

	if err != nil {
		log.Printf("Error en el servicio. Método CreateTurno: %v", err)
		return Turno{}, fmt.Errorf("error en servicio. Método CreateTurno: %v", err)
	}

	return turnoResponse, nil
}

// Update updates a patient.
func (s *service) Update(ctx context.Context, requestTurno RequestTurno, id int) (Turno, error) {
	turno := requestToTurno(requestTurno)
	turno.ID = id
	response, err := s.repository.Update(ctx, turno)
	if err != nil {
		log.Println("log de error en service de turno", err.Error())
		return Turno{}, errors.New("error en servicio. Metodo Update")
	}

	return response, nil
}

func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		log.Println("log de error borrado de turno", err.Error())
		return errors.New("error en servicio. Metodo Delete")
	}

	return nil
}

func requestToTurno(request RequestTurno) Turno {
	var turno Turno
	turno.Paciente = request.Paciente
	turno.Odontologo = request.Odontologo
	turno.FechaHora = request.FechaHora
	turno.Descripcion = request.Descripcion

	return turno
}
