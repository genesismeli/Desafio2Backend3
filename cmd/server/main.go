package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"

	// Alias para el paquete domain

	odontologoHandler "github.com/genesismeli/Desafio2Backend3/cmd/server/handler/odontologo"
	odontologoModel "github.com/genesismeli/Desafio2Backend3/internal/domain/odontologo"

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
	OdontologoGroup := router.Group("/odontologos")

	odontologoDatabase := odontologoModel.NewRepositoryMySql(db)
	odontologoService := odontologoModel.NewService(odontologoDatabase)
	controlador := odontologoHandler.NewControladorOdontologo(odontologoService)

	OdontologoGroup.GET("/:id", controlador.GetByID())
	OdontologoGroup.PUT("/:id", controlador.Update())
	OdontologoGroup.PATCH("/:id", controlador.UpdateSubject())
	OdontologoGroup.DELETE("/:id", controlador.Delete())



	router.Run("localhost" + puerto)

}
func connectDB() *sql.DB {
	var dbUsername, dbPassword, dbHost, dbPort, dbName string
	dbUsername = "root"
	dbPassword = "admin1234"
	dbHost = "localhost"
	dbPort = "3306"
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