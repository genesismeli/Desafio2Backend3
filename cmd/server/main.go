package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	pacienteHandler "github.com/genesismeli/Desafio2Backend3/cmd/server/handler/paciente"
	pacienteModel "github.com/genesismeli/Desafio2Backend3/internal/domain/paciente"
	"github.com/gin-gonic/gin"

	// Alias para el paquete handler
	// Alias para el paquete domain

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
	//PacientesGroup := router.Group("/pacientes")

	pacienteDatabase := pacienteModel.NewRepositoryMySql(db)
	pacienteService := pacienteModel.NewService(pacienteDatabase)
	pacienteHandler1 := pacienteHandler.NewControladorProducto(pacienteService)
	fmt.Print(pacienteHandler1)

	ctx := context.Background()

	pacientePrueba, _ := pacienteDatabase.GetByID(ctx, 3)

	// Define la ruta GET para el endpoint "/"
	router.GET("/pacientes/:id", func(c *gin.Context) {
		c.IndentedJSON(http.StatusCreated, pacientePrueba)
	})

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
