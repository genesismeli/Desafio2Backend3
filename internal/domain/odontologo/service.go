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
}


func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}