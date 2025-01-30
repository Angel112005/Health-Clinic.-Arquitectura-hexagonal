package appointment

import (
	"github.com/gin-gonic/gin"
	appAppointment "my-clinic-api/internal/application/appointment"
	domainAppointment "my-clinic-api/internal/domain/appointment"
	"net/http"
	"strconv"
	"fmt"
)

type Controller struct {
	createAppointment *appAppointment.CreateAppointment
	listAppointments  *appAppointment.ListAppointments
	updateAppointment *appAppointment.UpdateAppointment
	deleteAppointment *appAppointment.DeleteAppointment
}

func NewController(
	ca *appAppointment.CreateAppointment,
	la *appAppointment.ListAppointments,
	ua *appAppointment.UpdateAppointment,
	da *appAppointment.DeleteAppointment,
) *Controller {
	return &Controller{
		createAppointment: ca,
		listAppointments:  la,
		updateAppointment: ua,
		deleteAppointment: da,
	}
}

func (c *Controller) Create(ctx *gin.Context) {
    var a domainAppointment.Appointment
    if err := ctx.ShouldBindJSON(&a); err != nil {
        fmt.Println("Error en ShouldBindJSON:", err) // Agregar este log
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    fmt.Println("Datos recibidos:", a) // Log de los datos recibidos
    if err := c.createAppointment.Execute(a); err != nil {
        fmt.Println("Error en createAppointment.Execute:", err) // Log del error en la capa de aplicación
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create appointment"})
        return
    }
    ctx.JSON(http.StatusCreated, gin.H{"message": "Appointment created successfully"})
}


func (c *Controller) List(ctx *gin.Context) {
	appointments, err := c.listAppointments.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch appointments"})
		return
	}
	ctx.JSON(http.StatusOK, appointments)
}

func (c *Controller) Update(ctx *gin.Context) {
	var a domainAppointment.Appointment
	if err := ctx.ShouldBindJSON(&a); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := ctx.Param("id")
	parsedID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	a.ID = parsedID

	if err := c.updateAppointment.Execute(a); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update appointment"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Appointment updated successfully"})
}

func (c *Controller) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	parsedID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := c.deleteAppointment.Execute(parsedID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete appointment"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Appointment deleted successfully"})
}
