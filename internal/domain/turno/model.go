package turno

type Turno struct {
	ID              int    `json:"id"`
	PacienteDNI     string `json:"paciente"`
	OdontologoMatri string `json:"odontologo"`
	FechaHora       string `json:"fecha-hora"`
	Descripcion     string `json:"descripcion"`
}

type RequestTurno struct {
	PacienteDNI     string `json:"paciente"`
	OdontologoMatri string `json:"odontologo"`
	FechaHora       string `json:"fechaHora"`
	Descripcion     string `json:"descripcion"`
}

type RequestTurno2 struct {
	ID              *int       `json:"id"`
	PacienteDNI     *string    `json:"paciente"`
	OdontologoMatri *string    `json:"odontologo"`
	FechaHora       *string    `json:"fecha-hora"`
	Descripcion     *string    `json:"descripcion"`
}
