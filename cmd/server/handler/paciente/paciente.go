package paciente

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/genesismeli/Desafio2Backend3/internal/domain/paciente"
	"github.com/gin-gonic/gin"
)

type Controlador struct {
	service paciente.Service
}

func NewControladorProducto(service paciente.Service) *Controlador {
	return &Controlador{
		service: service,
	}
}

func (c *Controlador) GetByID(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		errors.New("Error linea 26")
		return
	}

	paciente, err := c.service.GetByID(ctx, id)
	if err != nil {
		errors.New("Error linea 32")
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"data": paciente,
	})

}
