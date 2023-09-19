package turno

import (
	"github.com/genesismeli/Desafio2Backend3/internal/domain/odontologo"
	"github.com/genesismeli/Desafio2Backend3/internal/domain/paciente"
)

type Turno struct {
	ID          int                   `json:"id"`
	Paciente    paciente.Paciente     `json:"paciente"`
	Odontologo  odontologo.Odontologo `json:"odontologo"`
	FechaHora   string                `json:"fecha-hora"`
	Descripcion string                `json:"descripcion"`
}

type RequestTurno struct {
	Paciente    paciente.Paciente     `json:"paciente"`
	Odontologo  odontologo.Odontologo `json:"odontologo"`
	FechaHora   string                `json:"fechaHora"`
	Descripcion string                `json:"descripcion"`
}

type RequestUpdateTurnoSubject struct {
	key   string `query:"key"`
	value string `query:"value"`
}
