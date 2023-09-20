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
