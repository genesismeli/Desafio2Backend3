package turno

import (
	"context"
	"errors"
	"log"
)

type Service interface {
	GetByID(ctx context.Context, id int) (Turno, error)
	Create(ctx context.Context, requestTurno RequestTurno) (Turno, error)
	Update(ctx context.Context, requestTurno RequestTurno, id int) (Turno, error)
	Delete(ctx context.Context, id int) error
	UpdateField(ctx context.Context, requestTurno2 RequestTurno2, id int) (Turno, error)
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
	response, err := s.repository.Create(ctx, turno)
	if err != nil {
		log.Println("Error en service Turno: Método Create")
		return Turno{}, errors.New("error en service Turno: Método Create")
	}
	return response, nil
}

func (s *service) Update(ctx context.Context, requestTurno RequestTurno, id int) (Turno, error) {
	turno := requestToTurno(requestTurno)
	turno.ID = id
	response, err := s.repository.Update(ctx, turno)
	if err != nil {
		log.Println("log de error en service de turni", err.Error())
		return Turno{}, errors.New("error en servicio. Metodo Update")
	}

	return response, nil
}

func requestToTurno(requestTurno RequestTurno) Turno {
	var turno Turno
	turno.PacienteDNI = requestTurno.PacienteDNI
	turno.OdontologoMatri = requestTurno.OdontologoMatri
	turno.FechaHora = requestTurno.FechaHora
	turno.Descripcion = requestTurno.Descripcion

	return turno
}

// Delete elimina el turno
func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		log.Println("Error borrado de turno", err.Error())
		return errors.New("error en service de Turnos: Metodo Delete")
	}

	return nil
}

func (s *service) UpdateField(ctx context.Context, requestTurno2 RequestTurno2, id int) (Turno, error) {

	pacienteFromDB, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("Error en servicio de paciente:", err.Error())
		return Turno{}, errors.New("error en servicio. Metodo GetByID desde updateFild")
	}

	paciente := requestToTurno2(requestTurno2, pacienteFromDB)

	response, err := s.repository.Update(ctx, paciente)
	if err != nil {
		log.Println("log de error en service de paciente", err.Error())
		return Turno{}, errors.New("error en servicio. Metodo Update")
	}

	return response, nil

}
func requestToTurno2(userReq RequestTurno2, turno Turno) Turno {

	if userReq.PacienteDNI != nil {
		turno.PacienteDNI = *userReq.PacienteDNI
	}

	if userReq.OdontologoMatri != nil {
		turno.OdontologoMatri = *userReq.OdontologoMatri
	}

	if userReq.FechaHora != nil {
		turno.FechaHora = *userReq.FechaHora
	}

	if userReq.Descripcion != nil {
		turno.Descripcion = *userReq.Descripcion
	}
	return turno
}
