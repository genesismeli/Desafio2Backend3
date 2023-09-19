package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

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
	PacientesGroup.POST("/create", controlador.Create())
	PacientesGroup.PUT("/:id", controlador.Update())
	PacientesGroup.DELETE("/:id", controlador.Delete())
	PacientesGroup.PATCH("/:id", func(c *gin.Context) {
		userID := c.Param("id")
		userIDint, err := strconv.Atoi(userID)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var userReq pacienteModel.RequestPaciente2

		err = c.BindJSON(&userReq)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		backgroundContext := context.Background()
		paciente, _ := pacienteService.GetByID(backgroundContext, userIDint)

		if userReq.Nombre != nil {
			paciente.Nombre = *userReq.Nombre
		}

		if userReq.Apellido != nil && *userReq.Apellido != "" {
			paciente.Apellido = *userReq.Apellido
		}

		if userReq.Domicilio != nil {
			paciente.Domicilio = *userReq.Domicilio
		}

		if userReq.DNI != nil {
			paciente.DNI = *userReq.DNI
		}

		if userReq.FechaAlta != nil {
			paciente.FechaAlta = *userReq.FechaAlta
		}

		paciente, _ = pacienteDatabase.Update(backgroundContext, paciente)
		c.JSON(http.StatusOK, paciente)
	})

	PacientesGroup.PATCH("/patch/:id", controlador.UpdateField())

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
