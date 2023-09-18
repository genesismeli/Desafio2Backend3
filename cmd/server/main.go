package main

import (
	"database/sql"
	"fmt"
	"net/http"

	// Alias para el paquete handler
	pacientemodel "github.com/genesismeli/Desafio2Backend3/internal/domain/paciente" // Alias para el paquete domain
	"github.com/gin-gonic/gin"

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

	var paciente pacientemodel.Paciente
	sql := "SELECT * FROM odontologos.pacientes where id=?"

	db.QueryRow(sql, "2").Scan(
		&paciente.ID,
		&paciente.Nombre,
		&paciente.Apellido,
		&paciente.Domicilio,
		&paciente.DNI,
		&paciente.FechaAlta,
	)

	fmt.Println(paciente)

	// Iniciar el engine
	router := gin.New()

	// Define la ruta GET para el endpoint "/"
	router.GET("/pacientes/:id", func(c *gin.Context) {
		c.IndentedJSON(http.StatusCreated, paciente)
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
