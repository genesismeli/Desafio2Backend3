package odontologo

import (
	"context"
	"database/sql"
	"errors"
)

var (
	ErrNotFound = errors.New("odontologo not found")
	ErrStatement = errors.New("statement error")
    ErrExec      = errors.New("execution error")
    ErrLastId    = errors.New("last ID error")
)


type Repository interface {
	Create(ctx context.Context, odontologo Odontologo) (Odontologo, error)
	GetByID(ctx context.Context, id int) (Odontologo, error)
	Update(ctx context.Context, odontologo Odontologo) (Odontologo, error)
	Delete(ctx context.Context, id int) error
	
}

type repository struct {
	db *sql.DB
}

func NewRepositoryMySql(db *sql.DB) Repository {
	return &repository {
		db: db,
	}
}

//Crea odontologo
func (r *repository) Create(ctx context.Context, odontologo Odontologo) (Odontologo, error) {

	statement, err := r.db.Prepare(QueryInsertOdontologo)

	if err != nil {
		return Odontologo{}, ErrStatement
	}

	defer statement.Close()

	result, err := statement.Exec(
		odontologo.Matricula,
		odontologo.Apellido,
		odontologo.Nombre,
		
		
	)

	if err != nil {
		return Odontologo{}, ErrExec
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return Odontologo{}, ErrLastId
	}

	odontologo.ID = int(lastId)

	return odontologo, nil
}

// GetByID 
func (r *repository) GetByID(ctx context.Context, id int) (Odontologo, error) {
	row := r.db.QueryRow(QueryGetOdontologoById, id)

	var odontologo Odontologo
	err := row.Scan(
		&odontologo.ID,
		&odontologo.Matricula,
		&odontologo.Apellido,
		&odontologo.Nombre,
		
		
	)

	if err != nil {
		return Odontologo{}, err
	}

	return odontologo, nil
}

// Actualizar odontologo
func (r *repository) Update(ctx context.Context, odontologo Odontologo) (Odontologo, error) {
	statement, err := r.db.Prepare(QueryUpdateOdontologo)
	if err != nil {
		return Odontologo{}, err
	}

	defer statement.Close()

	result, err := statement.Exec(
		odontologo.Nombre,
		odontologo.Apellido,
		odontologo.Matricula,
		odontologo.ID,
	)

	if err != nil {
		return Odontologo{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return Odontologo{}, err
	}

	if rowsAffected < 1 {
		return Odontologo{}, ErrNotFound
	}

	return odontologo, nil
}

// elimina odontologo
func (r *repository) Delete(ctx context.Context, id int) error {
	result, err := r.db.Exec(QueryDeleteOdontologo, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		return ErrNotFound
	}

	return nil
}