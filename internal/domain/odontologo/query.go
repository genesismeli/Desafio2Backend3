package odontologo

var (
	QueryInsertOdontologo       = `INSERT INTO odontologos.dentistas(matricula, apellido, nombre) VALUES(?,?,?)`
	QueryGetAllOdontologos      = `SELECT id, matricula, apellido, nombre FROM odontologos.dentistas`
	QueryDeleteOdontologo       = `DELETE FROM odontologos.dentistas WHERE id = ?`
	QueryGetOdontologoById      = `SELECT id, matricula,  apellido, nombre FROM odontologos.dentistas WHERE id = ?`
	QueryUpdateOdontologo       = `UPDATE odontologos.dentistas SET matricula = ?, apellido = ?, nombre = ? WHERE id = ?`
	QueryUpdateOdontologoNombre = `UPDATE odontologos.dentistas SET nombre = ? WHERE id = ?`
)