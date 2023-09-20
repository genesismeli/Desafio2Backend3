package turno

import (
	"net/http"
	"strconv"

	"github.com/genesismeli/Desafio2Backend3/internal/domain/turno"

	"github.com/genesismeli/Desafio2Backend3/core/web"

	"github.com/gin-gonic/gin"
)

type Controlador struct {
	service turno.Service
}

func NewControladorTurno(service turno.Service) *Controlador {
	return &Controlador{
		service: service,
	}
}
// Summary Get turno.
// @Tags   domain.turno
// @Produce json
// @Success 200 {object} web.Response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /turnos/:id [get]
func (c *Controlador) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "id invalido")
			return
		}

		turno, err := c.service.GetByID(ctx, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, turno)
	}

}

// Summary Post turno.
// @Tags   domain.turno
// @Produce json
// @Success 200 {object} web.Response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /turnos/create [post]
func (c *Controlador) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request turno.RequestTurno

		err := ctx.Bind(&request)

		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request")
			return
		}

		turno, err := c.service.Create(ctx, request)
		if err != nil {
            web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
            return
        }

        web.Success(ctx, http.StatusOK, gin.H{
            "data": turno,
        })
    }
}

	


// Summary Put turno.
// @Tags   domain.turno
// @Produce json
// @Success 200 {object} web.Response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /turnos/:id [put]
func (c *Controlador) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request turno.RequestTurno
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

		turno, err := c.service.Update(ctx, request, idInt)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"data": turno,
		})

	}
}
// Summary Delete turno.
// @Tags   domain.turno
// @Produce json
// @Success 200 {object} web.Response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /turnos/:id [delete]
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
			"mensaje": "turno eliminado",
		})
	}
}

// Summary Patch turno.
// @Tags   domain.turno
// @Produce json
// @Success 200 {object} web.Response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /turnos/patch/:id [patch]
func (c *Controlador) UpdateField() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request turno.RequestTurno2
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
