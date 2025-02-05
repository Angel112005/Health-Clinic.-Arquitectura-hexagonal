package main

import (
	"fmt"
	"my-clinic-api/config"
	appDoctor "my-clinic-api/doctor/application"
	doctorAPI "my-clinic-api/doctor/infrastructure/api"
	doctorPersistence "my-clinic-api/doctor/infrastructure/persistence"
	appAppointment "my-clinic-api/appointment/application"
	appointmentAPI "my-clinic-api/appointment/infrastructure/api"
	appointmentPersistence "my-clinic-api/appointment/infrastructure/persistence"
	"my-clinic-api/server" // Importación corregida para rutas
)

func main() {
	// Conexión a la base de datos
	db, err := config.ConnectDB()
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	defer db.Close()

	// Inicializar repositorios
	doctorRepo := doctorPersistence.NewDoctorMySQL(db)
	appointmentRepo := appointmentPersistence.NewAppointmentMySQL(db)

	// Inicializar casos de uso
	createDoctorUseCase := appDoctor.NewCreateDoctor(doctorRepo)
	listDoctorsUseCase := appDoctor.NewListDoctors(doctorRepo)
	updateDoctorUseCase := appDoctor.NewUpdateDoctor(doctorRepo)
	deleteDoctorUseCase := appDoctor.NewDeleteDoctor(doctorRepo)

	createAppointmentUseCase := appAppointment.NewCreateAppointment(appointmentRepo)
	listAppointmentsUseCase := appAppointment.NewListAppointments(appointmentRepo)
	updateAppointmentUseCase := appAppointment.NewUpdateAppointment(appointmentRepo)
	deleteAppointmentUseCase := appAppointment.NewDeleteAppointment(appointmentRepo)

	// Inicializar controladores
	doctorController := doctorAPI.NewController(
		createDoctorUseCase,
		listDoctorsUseCase,
		updateDoctorUseCase,
		deleteDoctorUseCase,
	)

	appointmentController := appointmentAPI.NewController(
		createAppointmentUseCase,
		listAppointmentsUseCase,
		updateAppointmentUseCase,
		deleteAppointmentUseCase,
	)

	// Configurar el router y levantar el servidor
	router := server.SetupRouter(doctorController, appointmentController) // Cambiado
	router.Run(":8080")
}
