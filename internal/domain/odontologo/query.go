package odontologo

var (
	QueryInsertOdontologo  = `INSERT INTO my_db.odontologos(nombre, apellido, matricula) VALUES(?,?,?)`
	QueryGetAllOdontologos = `SELECT id, nombre, apellido, matricula FROM my_db.odontologos`
	QueryDeleteOdontologo  = `DELETE FROM my_db.odontologos WHERE id = ?`
	QueryGetOdontologoById = `SELECT id, nombre, apellido, matricula FROM my_db.odontologos WHERE id = ?`
	QueryUpdateOdontologo  = `UPDATE my_db.odontologos SET nombre = ?, apellido = ?, matricula = ? WHERE id = ?`
)