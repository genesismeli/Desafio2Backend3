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
// Summary Get paciente.
// @Tags  domain.paciente
// @Produce json
// @Success 200 {object} web.Responses
// @Router /pacientes/{id} [get]
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

func (c *Controlador) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request paciente.RequestPaciente

		errBind := ctx.Bind(&request)

		if errBind != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request binding")
			return
		}

		id := ctx.Param("id")

		idInt, err := strconv.Atoi(id)

		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request param")
			return
		}

		paciente, err := c.service.Update(ctx, request, idInt)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"data": paciente,
		})

	}
}

func (c *Controlador) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "id invalido")
			return
		}

		err = c.service.Delete(ctx, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"mensaje": "paciente eliminado",
		})
	}
}

func (c *Controlador) UpdateField() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request paciente.RequestPaciente2

		errBind := ctx.BindJSON(&request)

		if errBind != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request binding")
			return
		}

		id := ctx.Param("id")

		idInt, err := strconv.Atoi(id)

		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request param")
			return
		}

		paciente, err := c.service.UpdateField(ctx, request, idInt)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"data": paciente,
		})

	}
}
