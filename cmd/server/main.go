package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"

	pacienteHandler "github.com/genesismeli/Desafio2Backend3/cmd/server/handler/paciente"
	pacienteModel "github.com/genesismeli/Desafio2Backend3/internal/domain/paciente"
	odontologoHandler "github.com/genesismeli/Desafio2Backend3/cmd/server/handler/odontologo"
	odontologoModel "github.com/genesismeli/Desafio2Backend3/internal/domain/odontologo"
	"log"

	"github.com/genesismeli/Desafio2Backend3/core/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

const (
	// puerto en el que se ejecutará el servidor.
	puerto = ":9090"
)

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

	OdontologoGroup.POST("/create", controladorOdontologo.Create())
	OdontologoGroup.GET("/:id", controladorOdontologo.GetByID())
	OdontologoGroup.PUT("/:id", controladorOdontologo.Update())
	OdontologoGroup.PATCH("/:id", controladorOdontologo.UpdateSubject())
	OdontologoGroup.DELETE("/:id", controladorOdontologo.Delete())

	// Router Group para pacientes
	PacientesGroup := router.Group("/pacientes")

	pacienteDatabase := pacienteModel.NewRepositoryMySql(db)
	pacienteService := pacienteModel.NewService(pacienteDatabase)
	controladorPaciente := pacienteHandler.NewControladorProducto(pacienteService)

	PacientesGroup.GET("/:id", middleware.Authenticate(), controladorPaciente.GetByID())
	PacientesGroup.POST("/create", controladorPaciente.Create())
	PacientesGroup.PUT("/:id", controladorPaciente.Update())
	PacientesGroup.DELETE("/:id", controladorPaciente.Delete())
	PacientesGroup.PATCH("/patch/:id", controladorPaciente.UpdateField())

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
