package odontologo


//Modelo del odontologo

type Odontologo struct {
	ID        int    `json:"id"`
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	Matricula string `json:"matricula"`
}

type RequestOdontologo struct {
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	Matricula string `json:"matricula"`
}