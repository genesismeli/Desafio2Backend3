package turno

var (
	QueryInsertTurno           = `INSERT INTO odontologos.turnos(paciente, odontologo, fechaHora, descripcion) VALUES(?,?,?,?)`
	QueryGetAllTurnos          = `SELECT id, paciente, odontologo, fechaHora, descripcion FROM odontologos.turnos`
	QueryDeleteTurno           = `DELETE FROM odontologos.turnos WHERE id = ?`
	QueryGetTurnoById          = `SELECT id, paciente,  odontologo, fechaHora, descripcion FROM odontologos.turnos WHERE id = ?`
	QueryUpdateTurnos          = `UPDATE odontologos.turnos SET paciente = ?, odontologo = ?, fechaHora = ?, descripcion = ? WHERE id = ?`
	QueryUpdateTurnoPaciente   = `UPDATE odontologos.turnos SET paciente = ? WHERE id = ?`
	QueryUpdateTurnoOdontologo = `UPDATE odontologos.turnos SET odontologo = ? WHERE id = ?`
	QueryUpdateTurnoFecha      = `UPDATE odontologos.turnos SET fechaHora = ? WHERE id = ?`
)
