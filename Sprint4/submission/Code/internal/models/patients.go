package models

import "database/sql"

type Patient struct {
	ID             int
	FirstInitials  string
	LastInitials   string
	DiagnosisCodes []string
	Goals          map[int]string
}

type PatientModel struct {
	DB *sql.DB
}

func (n *PatientModel) GetID(firstInitials string, lastInitials string) (int, error) {
	statement := `select Patient.id from Patient where Patient.firstInitials = ? && Patient.lastInitials = ?`

	row := n.DB.QueryRow(statement, firstInitials, lastInitials)

	var id int

	err := row.Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (n *PatientModel) GetAll() ([]Patient, error) {
	statement := `select Patient.id, Patient.firstInitials, Patient.lastInitials from Patient`

	patients := make([]Patient, 0)

	rows, err := n.DB.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var patient Patient

		err := rows.Scan(&patient.ID, &patient.FirstInitials, &patient.LastInitials)

		if err != nil {
			return nil, err
		}

		patients = append(patients, patient)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return patients, nil
}
