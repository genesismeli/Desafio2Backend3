package paciente

import (
	"context"
	"errors"
	"fmt"
	"log"
)

type Service interface {
	GetByID(ctx context.Context, id int) (Paciente, error)
	Create(ctx context.Context, requestPaciente RequestPaciente) (Paciente, error)
	Update(ctx context.Context, requestPaciente RequestPaciente, id int) (Paciente, error)
	Delete(ctx context.Context, id int) error
	UpdateField(ctx context.Context, requestPaciente2 RequestPaciente2, id int) (Paciente, error)
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

func (s *service) Create(ctx context.Context, requestPaciente RequestPaciente) (Paciente, error) {
	paciente := requestToPaciente(requestPaciente)
	pacienteReponse, err := s.repository.Create(ctx, paciente)

	if err != nil {
		log.Printf("Error en el servicio. Método CreatePaciente: %v", err)
		return Paciente{}, fmt.Errorf("error en servicio. Método CreatePaciente: %v", err)
	}

	return pacienteReponse, nil
}

// Update updates a patient.
func (s *service) Update(ctx context.Context, requestPaciente RequestPaciente, id int) (Paciente, error) {
	paciente := requestToPaciente(requestPaciente)
	paciente.ID = id
	response, err := s.repository.Update(ctx, paciente)
	if err != nil {
		log.Println("log de error en service de paciente", err.Error())
		return Paciente{}, errors.New("error en servicio. Metodo Update")
	}

	return response, nil
}

func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		log.Println("log de error borrado de producto", err.Error())
		return errors.New("error en servicio. Metodo Delete")
	}

	return nil
}

func (s *service) UpdateField(ctx context.Context, requestPaciente2 RequestPaciente2, id int) (Paciente, error) {

	pacienteFromDB, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("Error en servicio de paciente:", err.Error())
		return Paciente{}, errors.New("error en servicio. Metodo GetByID desde updateFild")
	}

	paciente := requestToPaciente2(requestPaciente2, pacienteFromDB)

	response, err := s.repository.Update(ctx, paciente)
	if err != nil {
		log.Println("log de error en service de paciente", err.Error())
		return Paciente{}, errors.New("error en servicio. Metodo Update")
	}

	return response, nil

}

func requestToPaciente(request RequestPaciente) Paciente {
	var paciente Paciente
	paciente.Nombre = request.Nombre
	paciente.Apellido = request.Apellido
	paciente.Domicilio = request.Domicilio
	paciente.DNI = request.DNI
	paciente.FechaAlta = request.FechaAlta

	return paciente
}

func requestToPaciente2(userReq RequestPaciente2, paciente Paciente) Paciente {

	if userReq.Nombre != nil {
		paciente.Nombre = *userReq.Nombre
	}

	if userReq.Apellido != nil {
		paciente.Apellido = *userReq.Apellido
	}

	if userReq.Domicilio != nil {
		paciente.Domicilio = *userReq.Domicilio
	}

	if userReq.DNI != nil {
		paciente.DNI = *userReq.DNI
	}

	if userReq.FechaAlta != nil {
		paciente.FechaAlta = *userReq.FechaAlta
	}

	return paciente
}


