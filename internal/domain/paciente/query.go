package paciente

var (
	QueryGetPacienteById   = `"SELECT * FROM odontologos.pacientes"`
	QueryGetAllOdontologos = `SELECT id, matricula, apellido, nombre FROM odontologos.dentistas`
)
