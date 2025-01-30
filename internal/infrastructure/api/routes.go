package api

import (
	"github.com/gin-gonic/gin"
	"my-clinic-api/internal/infrastructure/api/doctor"
	"my-clinic-api/internal/infrastructure/api/appointment"
)

func SetupRouter(
	doctorController *doctor.Controller,
	appointmentController *appointment.Controller,
) *gin.Engine {
	r := gin.Default()

	// Rutas CRUD para Doctores
	r.POST("/doctors", doctorController.Create)
	r.GET("/doctors", doctorController.List)
	r.PUT("/doctors/:id", doctorController.Update)
	r.DELETE("/doctors/:id", doctorController.Delete)

	// Rutas CRUD para Citas
	r.POST("/appointments", appointmentController.Create)
	r.GET("/appointments", appointmentController.List)
	r.PUT("/appointments/:id", appointmentController.Update)
	r.DELETE("/appointments/:id", appointmentController.Delete)

	return r
}
