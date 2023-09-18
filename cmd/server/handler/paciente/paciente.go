package paciente

import (
	"net/http"
	"strconv"

	"github.com/genesismeli/Desafio2Backend3/core/web"

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

func (c *Controlador) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "id invalido")
			return
		}

		paciente, err := c.service.GetByID(ctx, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, paciente)
	}

}

func (c *Controlador) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request paciente.RequestPaciente

		err := ctx.Bind(&request)

		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request")
			return
		}

		paciente, err := c.service.Create(ctx, request)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"data": paciente,
		})

	}
}
