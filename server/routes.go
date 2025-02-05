package server

import (
	"github.com/gin-gonic/gin"
	doctorAPI "my-clinic-api/doctor/infrastructure/api" // Ajustada a la nueva estructura
	appointmentAPI "my-clinic-api/appointment/infrastructure/api" // Ajustada a la nueva estructura
)


func SetupRouter(
	doctorController *doctorAPI.Controller,
	appointmentController *appointmentAPI.Controller,
) *gin.Engine {
	r := gin.Default()

	// Rutas para Doctors
	r.POST("/doctors", doctorController.Create)
	r.GET("/doctors", doctorController.List)
	r.PUT("/doctors/:id", doctorController.Update)
	r.DELETE("/doctors/:id", doctorController.Delete)

	// Rutas para Appointments
	r.POST("/appointments", appointmentController.Create)
	r.GET("/appointments", appointmentController.List)
	r.PUT("/appointments/:id", appointmentController.Update)
	r.DELETE("/appointments/:id", appointmentController.Delete)

	return r
}
