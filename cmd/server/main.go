package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"

	pacienteHandler "github.com/genesismeli/Desafio2Backend3/cmd/server/handler/paciente"
	pacienteModel "github.com/genesismeli/Desafio2Backend3/internal/domain/paciente"
	odontologoHandler "github.com/genesismeli/Desafio2Backend3/cmd/server/handler/odontologo"
	odontologoModel "github.com/genesismeli/Desafio2Backend3/internal/domain/odontologo"
	turnoHandler "github.com/genesismeli/Desafio2Backend3/cmd/server/handler/turno"
	turnoModel "github.com/genesismeli/Desafio2Backend3/internal/domain/turno"
	"log"

	"github.com/genesismeli/Desafio2Backend3/core/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

const (
	// puerto en el que se ejecutará el servidor.
	puerto = ":9090"
)

// @title Desafío backend 2
// @version 1.0
// @description 
// @termsOfService

// @contact.name
// @contact.url

// @license.name asdfasdf
// @license.url 
func main() {

	//utilizamos la librería godotenv, con esto recorremos las variables del archivo .env y las setean en el environment.
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Conección a la base de datos
	db := connectDB()
	defer db.Close() // Cierra la conexión a la base de datos al finalizar la función main.

	//egine
	router := gin.New()

	// Router Group para odontologos
	OdontologoGroup := router.Group("/odontologos")

	odontologoDatabase := odontologoModel.NewRepositoryMySql(db)
	odontologoService := odontologoModel.NewService(odontologoDatabase)
	controladorOdontologo := odontologoHandler.NewControladorOdontologo(odontologoService)

	OdontologoGroup.GET("/:id", controladorOdontologo.GetByID())
	OdontologoGroup.POST("/create",middleware.Authenticate(), controladorOdontologo.Create())
	OdontologoGroup.PUT("/:id",middleware.Authenticate(), controladorOdontologo.Update())
	OdontologoGroup.PATCH("/patch/:id",middleware.Authenticate(), controladorOdontologo.UpdateField())
	OdontologoGroup.DELETE("/:id",middleware.Authenticate(), controladorOdontologo.Delete())

	// Router Group para pacientes
	PacientesGroup := router.Group("/pacientes")

	pacienteDatabase := pacienteModel.NewRepositoryMySql(db)
	pacienteService := pacienteModel.NewService(pacienteDatabase)
	controladorPaciente := pacienteHandler.NewControladorProducto(pacienteService)

	PacientesGroup.GET("/:id", controladorPaciente.GetByID())
	PacientesGroup.POST("/create",middleware.Authenticate(), controladorPaciente.Create())
	PacientesGroup.PUT("/:id",middleware.Authenticate(), controladorPaciente.Update())
	PacientesGroup.PATCH("/patch/:id",middleware.Authenticate(), controladorPaciente.UpdateField())
	PacientesGroup.DELETE("/:id",middleware.Authenticate(), controladorPaciente.Delete())
	
	// Router Group para turnos
	turnosGroup := router.Group("/turnos")

	turnoDatabase := turnoModel.NewRepositoryMySql(db)
	turnoService := turnoModel.NewService(turnoDatabase)
	controladorTurno := turnoHandler.NewControladorProducto(turnoService)
	turnosGroup.GET("/:id", controladorTurno.GetByID())
	turnosGroup.POST("/create",middleware.Authenticate(), controladorTurno.Create())
	turnosGroup.PUT("/:id",middleware.Authenticate(), controladorTurno.Update())
	turnosGroup.PATCH("/patch/:id",middleware.Authenticate(), controladorTurno.UpdateField())
	turnosGroup.DELETE("/:id",middleware.Authenticate(), controladorTurno.Delete())


	router.Run("localhost" + puerto)

}
func connectDB() *sql.DB {
	var dbUsername, dbPassword, dbHost, dbPort, dbName string
	dbUsername = "root"
	dbPassword = "admin1234"
	dbHost = "localhost"
	dbPort = "33060"
	dbName = "odontologos"

	// Create the data source.
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUsername, dbPassword, dbHost, dbPort, dbName)

	// Open the connection.
	db, err := sql.Open("mysql", dataSource)

	if err != nil {
		panic(err)
	}

	// Check the connection.
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	return db
}
