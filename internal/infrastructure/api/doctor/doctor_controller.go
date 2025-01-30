package doctor

import (
	"github.com/gin-gonic/gin"
	appDoctor "my-clinic-api/internal/application/doctor" // Alias para application/doctor
	domainDoctor "my-clinic-api/internal/domain/doctor"   // Alias para domain/doctor
	"net/http"
	"strconv"
)

type Controller struct {
	createDoctor *appDoctor.CreateDoctor
	listDoctors  *appDoctor.ListDoctors
	updateDoctor *appDoctor.UpdateDoctor
	deleteDoctor *appDoctor.DeleteDoctor
}

func NewController(cd *appDoctor.CreateDoctor, ld *appDoctor.ListDoctors, ud *appDoctor.UpdateDoctor, dd *appDoctor.DeleteDoctor) *Controller {
	return &Controller{
		createDoctor: cd,
		listDoctors:  ld,
		updateDoctor: ud,
		deleteDoctor: dd,
	}
}

func (c *Controller) Create(ctx *gin.Context) {
	var d domainDoctor.Doctor
	if err := ctx.ShouldBindJSON(&d); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.createDoctor.Execute(d); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create doctor"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Doctor created successfully"})
}

func (c *Controller) List(ctx *gin.Context) {
	doctors, err := c.listDoctors.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch doctors"})
		return
	}
	ctx.JSON(http.StatusOK, doctors)
}

func (c *Controller) Update(ctx *gin.Context) {
	var d domainDoctor.Doctor
	if err := ctx.ShouldBindJSON(&d); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Obtenemos el ID de la ruta
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing ID parameter"})
		return
	}
	d.ID = convertID(id, ctx)

	if err := c.updateDoctor.Execute(d); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update doctor"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Doctor updated successfully"})
}

func (c *Controller) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing ID parameter"})
		return
	}
	parsedID := convertID(id, ctx)

	if err := c.deleteDoctor.Execute(parsedID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete doctor"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Doctor deleted successfully"})
}

// Función de utilidad para convertir el ID
func convertID(id string, ctx *gin.Context) int {
	parsedID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		ctx.Abort()
	}
	return parsedID
}
