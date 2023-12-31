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

func (r *repository) Create(ctx context.Context, turno Turno) (Turno, error) {
	// Construir la consulta SQL directamente
	sqlQuery := "INSERT INTO odontologos.turnos (paciente_DNI, dentista_matricula, fecha_hora, descripcion) VALUES (?, ?, ?, ?)"

	// Ejecutar la consulta SQL directamente
	result, err := r.db.Exec(sqlQuery, turno.PacienteDNI, turno.OdontologoMatri, "2023-01-15", turno.Descripcion)

	if err != nil {
		return Turno{}, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return Turno{}, err
	}

	turno.ID = int(lastID)

	return turno, nil
}

func (r *repository) Update(ctx context.Context, turno Turno) (Turno, error) {
	statement, err := r.db.Prepare(`UPDATE odontologos.turnos SET paciente_DNI = ?, dentista_matricula = ?, fecha_hora = ?, descripcion = ? WHERE id = ?`)

	if err != nil {
		return Turno{}, err
	}

	defer statement.Close()

	result, err := statement.Exec(
		turno.PacienteDNI,
		turno.OdontologoMatri,
		"2023-09-17",
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
func (r repository) GetByDNI(ctx context.Context, dni string) (Turno, error) {
    row := r.db.QueryRow("SELECT FROM odontologos.turnos where paciente_DNI=?", dni)

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