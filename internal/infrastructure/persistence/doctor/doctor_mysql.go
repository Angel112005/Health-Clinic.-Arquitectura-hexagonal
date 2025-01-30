package doctor

import (
	"database/sql"
	"my-clinic-api/internal/domain/doctor"
)

type DoctorMySQL struct {
	db *sql.DB
}

func NewDoctorMySQL(db *sql.DB) *DoctorMySQL {
	return &DoctorMySQL{db: db}
}

func (r *DoctorMySQL) Save(d doctor.Doctor) error {
	_, err := r.db.Exec("INSERT INTO doctors (name, specialty, experience) VALUES (?, ?, ?)",
		d.Name, d.Specialty, d.Experience)
	return err
}

func (r *DoctorMySQL) FindAll() ([]doctor.Doctor, error) {
	rows, err := r.db.Query("SELECT id, name, specialty, experience FROM doctors")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var doctors []doctor.Doctor
	for rows.Next() {
		var d doctor.Doctor
		if err := rows.Scan(&d.ID, &d.Name, &d.Specialty, &d.Experience); err != nil {
			return nil, err
		}
		doctors = append(doctors, d)
	}
	return doctors, nil
}

func (r *DoctorMySQL) FindByID(id int) (*doctor.Doctor, error) {
	row := r.db.QueryRow("SELECT id, name, specialty, experience FROM doctors WHERE id = ?", id)
	var d doctor.Doctor
	if err := row.Scan(&d.ID, &d.Name, &d.Specialty, &d.Experience); err != nil {
		return nil, err
	}
	return &d, nil
}


func (r *DoctorMySQL) Update(d doctor.Doctor) error {
	_, err := r.db.Exec(
		"UPDATE doctors SET name = ?, specialty = ?, experience = ? WHERE id = ?",
		d.Name, d.Specialty, d.Experience, d.ID,
	)
	return err
}

func (r *DoctorMySQL) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM doctors WHERE id = ?", id)
	return err
}




// func (r *DoctorMySQL) Update(d doctor.Doctor) error {
// 	_, err := r.db.Exec("UPDATE doctors SET name = ?, specialty = ?, experience = ? WHERE id = ?",
// 		d.Name, d.Specialty, d.Experience, d.ID)
// 	return err
// }

// func (r *DoctorMySQL) Delete(id int) error {
// 	_, err := r.db.Exec("DELETE FROM doctors WHERE id = ?", id)
// 	return err
// }

