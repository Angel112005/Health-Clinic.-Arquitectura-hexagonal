package main

import (
	"fmt"
	"my-clinic-api/dependencies"
	"my-clinic-api/server"
)

func main() {
	// Inicializar dependencias
	appDeps, err := dependencies.InitializeDependencies()
	if err != nil {
		fmt.Println("Error initializing dependencies:", err)
		return
	}

	// Configurar el router y levantar el servidor
	router := server.SetupRouter(appDeps.DoctorController, appDeps.AppointmentController)
	router.Run(":8080")
}
