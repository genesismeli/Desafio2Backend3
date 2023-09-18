package paciente

var (
	QueryGetPacienteById   = `"SELECT * FROM odontologos.pacientes where id= 2"`
	QueryGetAllOdontologos = `SELECT id, matricula, apellido, nombre FROM odontologos.dentistas`
)
