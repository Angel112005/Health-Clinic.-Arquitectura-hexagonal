package dependencies

import (
	"my-clinic-api/config"
	appDoctor "my-clinic-api/doctor/application"
	doctorAPI "my-clinic-api/doctor/infrastructure/api"
	doctorPersistence "my-clinic-api/doctor/infrastructure/persistence"
	appAppointment "my-clinic-api/appointment/application"
	appointmentAPI "my-clinic-api/appointment/infrastructure/api"
	appointmentPersistence "my-clinic-api/appointment/infrastructure/persistence"
)

type AppDependencies struct {
	DoctorController     *doctorAPI.Controller
	AppointmentController *appointmentAPI.Controller
}

func InitializeDependencies() (*AppDependencies, error) {
	// Conexión a la base de datos
	db, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}

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

	return &AppDependencies{
		DoctorController:     doctorController,
		AppointmentController: appointmentController,
	}, nil
}
