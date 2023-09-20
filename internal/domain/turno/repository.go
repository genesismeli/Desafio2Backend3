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
		&turno.OdontologoMatri,
		&turno.PacienteDNI,
		&turno.FechaHora,
		&turno.Descripcion,
	)

	if err != nil {
		return Turno{}, err
	}

	return turno, nil
}


// Create crea un turno
func (r *repository) Create(ctx context.Context, turno Turno) (Turno, error) {

	statement, err := r.db.Prepare(QueryInsertTurno)

	if err != nil {
		return Turno{}, ErrStatement
	}

	defer statement.Close()

	result, err := statement.Exec(
		turno.PacienteDNI,
		turno.OdontologoMatri,
		turno.FechaHora,
		turno.Descripcion,
	)

	if err != nil {
		return Turno{}, ErrExec
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return Turno{}, ErrLastId
	}

	turno.ID = int(lastId)

	return turno, nil
}

// Delete elimina el turno
func (r *repository) Delete(ctx context.Context, id int) error {
	result, err := r.db.Exec(QueryDeleteTurno, id)
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