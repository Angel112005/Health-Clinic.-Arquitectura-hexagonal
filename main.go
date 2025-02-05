package main

import (
	"fmt"
	"my-clinic-api/config"
	appDoctor "my-clinic-api/internal/application/doctor"
	appAppointment "my-clinic-api/internal/application/appointment"
	doctorAPI "my-clinic-api/internal/infrastructure/api/doctor"
	appointmentAPI "my-clinic-api/internal/infrastructure/api/appointment"
	"my-clinic-api/internal/infrastructure/persistence/doctor"
	"my-clinic-api/internal/infrastructure/persistence/appointment"
	api "my-clinic-api/internal/infrastructure/api"
)

func main() {
	// Conexión a la base de datos
	db, err := config.ConnectDB()
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	//Defer para no perder la conexxión
	defer db.Close()

	// Inicializar repositorios
	doctorRepo := doctor.NewDoctorMySQL(db)
	// Inicializar repositorios
	appointmentRepo := appointment.NewAppointmentMySQL(db)

	// Inicializar casos de uso
	createDoctorUseCase := appDoctor.NewCreateDoctor(doctorRepo)
	listDoctorsUseCase := appDoctor.NewListDoctors(doctorRepo)
	updateDoctorUseCase := appDoctor.NewUpdateDoctor(doctorRepo)
	deleteDoctorUseCase := appDoctor.NewDeleteDoctor(doctorRepo)

	// Inicializar controladores
	doctorController := doctorAPI.NewController(
		createDoctorUseCase,
		listDoctorsUseCase,
		updateDoctorUseCase,
		deleteDoctorUseCase,
	)

	// Inicializar casos de uso
	createAppointmentUseCase := appAppointment.NewCreateAppointment(appointmentRepo)
	listAppointmentsUseCase := appAppointment.NewListAppointments(appointmentRepo)
	updateAppointmentUseCase := appAppointment.NewUpdateAppointment(appointmentRepo)
	deleteAppointmentUseCase := appAppointment.NewDeleteAppointment(appointmentRepo)

	// Inicializar controladores
	appointmentController := appointmentAPI.NewController(
		createAppointmentUseCase,
		listAppointmentsUseCase,
		updateAppointmentUseCase,
		deleteAppointmentUseCase,
	)

	// Configurar el router y levantar el servidor
	router := api.SetupRouter(doctorController, appointmentController)
	router.Run(":8080")
}
