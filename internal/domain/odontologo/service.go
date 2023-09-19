package odontologo

import (
	"context"
	"errors"
	"log"
)

type service struct {
	repository Repository
}

//Para la administración de datos de odontologos
type Service interface {
	Create(ctx context.Context, requestOdontologo RequestOdontologo) (Odontologo, error)
	GetByID(ctx context.Context, id int) (Odontologo, error)
	Update(ctx context.Context, requestOdontologo RequestOdontologo, id int) (Odontologo, error)
	Delete(ctx context.Context, id int) error
	UpdateField(ctx context.Context, requestPaciente2 RequestOdontologo2, id int) (Odontologo, error)
}


func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

//Funcion crear odontologo
func (s *service) Create(ctx context.Context, requestOdontologo RequestOdontologo) (Odontologo, error) {
	odontologo := requestToOdontologo(requestOdontologo)
	response, err := s.repository.Create(ctx, odontologo)
	if err != nil {
		log.Println("Error en crear odontologo")
		return Odontologo{}, errors.New("error en crear odontologo")
	}

	return response, nil
}

// Funcion que devuelve el odontologo buscado por ID
func (s *service) GetByID(ctx context.Context, id int) (Odontologo, error) {
	odontologo, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("Error en el metodo GetByID", err.Error())
		return Odontologo{}, errors.New("error en el metodo GetByID")
	}

	return odontologo, nil
}


// Funcion para actualizar un odontologo.
func (s *service) Update(ctx context.Context, requestOdontologo RequestOdontologo, id int) (Odontologo, error) {
	odontologo := requestToOdontologo(requestOdontologo)
	odontologo.ID = id
	response, err := s.repository.Update(ctx, odontologo)
	if err != nil {
		log.Println("Error en el metodo de update de odontologo", err.Error())
		return Odontologo{}, errors.New("error en el metodo de update de odontologo")
	}

	return response, nil
}

func (s *service) UpdateField(ctx context.Context, requestOdontologo2 RequestOdontologo2, id int) (Odontologo, error) {

	pacienteFromDB, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("Error en servicio de odontologo:", err.Error())
		return Odontologo{}, errors.New("error en servicio. Metodo GetByID desde updateFild")
	}

	paciente := requestToOdontologo2(requestOdontologo2, pacienteFromDB)

	response, err := s.repository.Update(ctx, paciente)
	if err != nil {
		log.Println("log de error en service de odontologo", err.Error())
		return Odontologo{}, errors.New("error en servicio. Metodo Update")
	}

	return response, nil

}

// Funcion para eliminar odontologo
func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		log.Println("Error en el service Odontologo: Método Delete", err.Error())
		return errors.New("error en service: Método Delete")
	}

	return nil
}

func requestToOdontologo(requestOdontologo RequestOdontologo) Odontologo {
	var odontologo Odontologo
	odontologo.Matricula = requestOdontologo.Matricula
	odontologo.Apellido = requestOdontologo.Apellido
	odontologo.Nombre = requestOdontologo.Nombre
	
	return odontologo

}

func requestToOdontologo2(userReq RequestOdontologo2, odontologo Odontologo) Odontologo {

	if userReq.Nombre != nil {
		odontologo.Nombre = *userReq.Nombre
	}

	if userReq.Apellido != nil {
		odontologo.Apellido = *userReq.Apellido
	}

	if userReq.Matricula != nil {
		odontologo.Matricula = *userReq.Matricula
	}

	return odontologo
}
