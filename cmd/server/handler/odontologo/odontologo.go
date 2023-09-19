package odontologo
<<<<<<< HEAD

import (
<<<<<<< HEAD
=======
	"net/http"
	"strconv"

>>>>>>> 8452c61c3db9073d2fb57a94b640edc38b42eee8
	"github.com/genesismeli/Desafio2Backend3/core/web"
	"github.com/genesismeli/Desafio2Backend3/internal/domain/odontologo"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Controlador struct {
	service odontologo.Service
}

func NewControladorOdontologo(service odontologo.Service) *Controlador {
	return &Controlador{
		service: service,
	}
}

// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /odontologos/:id [get]
func (c *Controlador) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "No existe el ID o es invalido")
			return
		}

		product, err := c.service.GetByID(ctx, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"data": odonto,
		})
	}
}

// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /odontologos/:id [put]
func (c *Controlador) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request odontologo.RequestOdontologo

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

		odonto, err := c.service.Update(ctx, request, idInt)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"data": odonto,
		})

	}
}

// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /odontologos/:id [patch]
func (c *Controlador) UpdateSubject() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request odontologo.RequestUpdateOdontologoSubject
		errBind := ctx.Bind(&request)
		if errBind != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request")
			return
		}

		id := ctx.Param("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request param")
			return
		}

		odonto, err := c.service.UpdateSubject(ctx, idInt, request)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Succses(ctx, http.StatusOK, gin.H{
			"data": odonto,
		})

	}
}

// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /odontologos/:id [delete]
func (c *Controlador) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "No existe el ID o es invalido")
			return
		}

		err = c.service.Delete(ctx, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"mensaje": "odontologo eliminado",
		})
	}
}
<<<<<<< HEAD
=======
=======
>>>>>>> securityLeandro
>>>>>>> 8452c61c3db9073d2fb57a94b640edc38b42eee8
