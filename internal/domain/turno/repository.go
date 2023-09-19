package turno

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

func (r *repository) GetByID(ctx context.Context, id int) (Turno, error) {
	row := r.db.QueryRow("SELECT * FROM odontologos.turnos where id=?", id)

	var turno Turno
	err := row.Scan(
		&turno.ID,
		&turno.Paciente.DNI,
		&turno.Odontologo.Matricula,
		&turno.FechaHora,
		&turno.Descripcion,
	)

	if err != nil {
		return Turno{}, err
	}

	return turno, nil
}

func (r *repository) Create(ctx context.Context, turno Turno) (Turno, error) {

	statement, err := r.db.Prepare("INSERT INTO turnos (paciente, odontologo, fechaHora, descripcion) VALUES (?, ?, ?, ?)")

	if err != nil {
		return Turno{}, ErrStatement
	}

	defer statement.Close()

	result, err := statement.Exec(
		turno.Paciente,
		turno.Odontologo,
		turno.FechaHora,
		turno.Descripcion,
	)

	if err != nil {
		return Turno{}, ErrExec
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return Turno{}, ErrLastId
	}

	turno.ID = int(lastID)

	return turno, nil
}

func (r *repository) Update(ctx context.Context, turno Turno) (Turno, error) {
	statement, err := r.db.Prepare("UPDATE odontologos.turnos SET paciente = ?, odontologo = ?, fechaHora = ?, descripcion = ? WHERE id = ?")

	if err != nil {
		return Turno{}, err
	}

	defer statement.Close()

	result, err := statement.Exec(
		turno.Paciente,
		turno.Odontologo,
		turno.FechaHora,
		turno.Descripcion,
		turno.ID,
	)

	if err != nil {
		return Turno{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return Turno{}, err
	}

	if rowsAffected < 1 {
		return Turno{}, ErrNotFound
	}

	return turno, nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	result, err := r.db.Exec(`DELETE FROM odontologos.turnos WHERE id = ?`, id)
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