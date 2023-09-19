package odontologo

//Modelo del odontologo

type Odontologo struct {
	ID        int    `json:"id"`
	Matricula string `json:"matricula"`
	Apellido  string `json:"apellido"`
	Nombre    string `json:"nombre"`
}

type RequestOdontologo struct {
	Matricula string `json:"matricula"`
	Apellido  string `json:"apellido"`
	Nombre    string `json:"nombre"`
}

type RequestUpdateOdontologoSubject struct {
	key   string `query:"key"`
	value string `query:"value"`
}
