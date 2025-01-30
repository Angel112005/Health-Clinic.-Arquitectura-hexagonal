package appointment

type Appointment struct {
	ID              int    `json:"id"`
	DoctorID        int    `json:"doctor_id"`
	PatientName     string `json:"patient_name"`
	AppointmentDate string `json:"appointment_date"`
	Status          string `json:"status"`
}
