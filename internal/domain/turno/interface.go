package turno

import (
	"context"
	"errors"
)

var (
	ErrEmptyList = errors.New("the list is empty")
	ErrNotFound  = errors.New("product not found")
	ErrStatement = errors.New("error preparing statement")
	ErrExec      = errors.New("error exect statement")
	ErrLastId    = errors.New("error getting last id")
)

type Repository interface {
	Create(ctx context.Context, turno Turno) (Turno, error)
	GetByID(ctx context.Context, id int) (Turno, error)
	//Update(ctx context.Context, paciente Turno) (Turno, error)
	//Delete(ctx context.Context, id int) error
	//UpdateField(ctx context.Context, paciente Paciente) (Paciente, error)
	//Update(ctx context.Context, paciente Paciente) (Paciente, error)
	//UpdateField(ctx context.Context, paciente Paciente) (Paciente, error)
	//Delete(ctx context.Context, id int) error
}
