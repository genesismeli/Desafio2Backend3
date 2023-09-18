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
	GetAll(ctx context.Context) ([]Odontologo, error)
	GetByID(ctx context.Context, id int) (Odontologo, error)
	Update(ctx context.Context, requestOdontologo RequestOdontologo, id int) (Odontologo, error)
	Delete(ctx context.Context, id int) error
	UpdateName(ctx context.Context, id int, nombreNuevo string) (Odontologo, error)
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
		return Odontologo{}, errors.New("Error en crear odontologo")
	}

	return response, nil
}

// Funcion que devuelve el odontologo buscado por ID
func (s *service) GetByID(ctx context.Context, id int) (Odontologo, error) {
	odontologo, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("Error en el metodo GetByID", err.Error())
		return Odontologo{}, errors.New("Error en el metodo GetByID")
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
		return Odontologo{}, errors.New("Error en el metodo de update de odontologo")
	}

	return response, nil
}

// Update actualiza alguno de los campos de odontologo
func (s *service) UpdateSubject(ctx context.Context, id int, request RequestUpdateOdontologoSubject) (Odontologo, error) {

	response, err := s.repository.UpdateSubject(ctx, id, request)
	if err != nil {
		log.Println("log de error en service de odontologo", err.Error())
		return Odontologo{}, errors.New("error en servicio. Metodo UpdateName")
	}
	return response, nil
}

// Funcion para eliminar odontologo
func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		log.Println("Error en el service Odontologo: Método Delete", err.Error())
		return errors.New("Error en service: Método Delete")
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