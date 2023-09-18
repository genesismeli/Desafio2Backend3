package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"

	// Alias para el paquete domain

	pacienteHandler "github.com/genesismeli/Desafio2Backend3/cmd/server/handler/paciente"
	pacienteModel "github.com/genesismeli/Desafio2Backend3/internal/domain/paciente"

	_ "github.com/go-sql-driver/mysql"
)

const (
	// puerto en el que se ejecutará el servidor.
	puerto = ":9090"
)

func main() {
	fmt.Println("¡Hola, mundo!")
	// Connect to the database.
	db := connectDB()
	defer db.Close() // Cierra la conexión a la base de datos al finalizar la función main.

	//egine
	router := gin.New()
	PacientesGroup := router.Group("/pacientes")

	pacienteDatabase := pacienteModel.NewRepositoryMySql(db)
	pacienteService := pacienteModel.NewService(pacienteDatabase)
	controlador := pacienteHandler.NewControladorProducto(pacienteService)

	PacientesGroup.GET("/:id", controlador.GetByID())

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
