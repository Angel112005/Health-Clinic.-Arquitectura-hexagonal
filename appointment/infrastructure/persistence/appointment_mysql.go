package appointment

import (
	"database/sql"
	"my-clinic-api/appointment/domain"
)

type AppointmentMySQL struct {
	db *sql.DB
}

func NewAppointmentMySQL(db *sql.DB) *AppointmentMySQL {
	return &AppointmentMySQL{db: db}
}

func (r *AppointmentMySQL) Save(a appointment.Appointment) error {
	_, err := r.db.Exec(
		"INSERT INTO appointments (doctor_id, patient_name, appointment_date, status) VALUES (?, ?, ?, ?)",
		a.DoctorID, a.PatientName, a.AppointmentDate, a.Status,
	)
	return err
}


func (r *AppointmentMySQL) FindAll() ([]appointment.Appointment, error) {
	rows, err := r.db.Query("SELECT id, doctor_id, patient_name, appointment_date, status FROM appointments")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var appointments []appointment.Appointment
	for rows.Next() {
		var a appointment.Appointment
		if err := rows.Scan(&a.ID, &a.DoctorID, &a.PatientName, &a.AppointmentDate, &a.Status); err != nil {
			return nil, err
		}
		appointments = append(appointments, a)
	}
	return appointments, nil
}

func (r *AppointmentMySQL) FindByID(id int) (*appointment.Appointment, error) {
	row := r.db.QueryRow("SELECT id, doctor_id, patient_name, appointment_date, status FROM appointments WHERE id = ?", id)
	var a appointment.Appointment
	if err := row.Scan(&a.ID, &a.DoctorID, &a.PatientName, &a.AppointmentDate, &a.Status); err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *AppointmentMySQL) Update(a appointment.Appointment) error {
	_, err := r.db.Exec(
		"UPDATE appointments SET doctor_id = ?, patient_name = ?, appointment_date = ?, status = ? WHERE id = ?",
		a.DoctorID, a.PatientName, a.AppointmentDate, a.Status, a.ID,
	)
	return err
}

func (r *AppointmentMySQL) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM appointments WHERE id = ?", id)
	return err
}
