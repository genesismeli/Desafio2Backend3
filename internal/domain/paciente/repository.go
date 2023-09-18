package paciente

import (
	"context"
	"database/sql"
)

type repository struct {
	db *sql.DB
}

// NewRepositoryMySql creates a new repository.
func NewRepositoryMySql(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetByID(ctx context.Context, id int) (Paciente, error) {
	row := r.db.QueryRow("SELECT * FROM odontologos.pacientes where id=?", id)

	var paciente Paciente
	err := row.Scan(
		&paciente.ID,
		&paciente.Nombre,
		&paciente.Apellido,
		&paciente.Domicilio,
		&paciente.DNI,
		&paciente.FechaAlta,
	)

	if err != nil {
		return Paciente{}, err
	}

	return paciente, nil
}

func (r *repository) Create(ctx context.Context, paciente Paciente) (Paciente, error) {

	statement, err := r.db.Prepare("INSERT INTO pacientes (DNI, nombre, apellido, domicilio, fecha_alta) VALUES (?, ?, ?, ?, ?)")

	if err != nil {
		return Paciente{}, ErrStatement
	}

	defer statement.Close()

	result, err := statement.Exec(
		paciente.DNI,
		paciente.Nombre,
		paciente.Apellido,
		paciente.Domicilio,
		paciente.FechaAlta,
	)

	if err != nil {
		return Paciente{}, ErrExec
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return Paciente{}, ErrLastId
	}

	paciente.ID = int(lastID)

	return paciente, nil
}

func (r *repository) Update(ctx context.Context, paciente Paciente) (Paciente, error) {
	statement, err := r.db.Prepare("UPDATE odontologos.pacientes SET DNI = ?, nombre = ?, apellido = ?, domicilio = ?, fecha_alta = ? WHERE id = ?")

	if err != nil {
		return Paciente{}, err
	}

	defer statement.Close()

	result, err := statement.Exec(
		paciente.DNI,
		paciente.Nombre,
		paciente.Apellido,
		paciente.Domicilio,
		paciente.FechaAlta,
		paciente.ID,
	)

	if err != nil {
		return Paciente{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return Paciente{}, err
	}

	if rowsAffected < 1 {
		return Paciente{}, ErrNotFound
	}

	return paciente, nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	result, err := r.db.Exec(`DELETE FROM odontologos.pacientes WHERE id = ?`, id)
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
