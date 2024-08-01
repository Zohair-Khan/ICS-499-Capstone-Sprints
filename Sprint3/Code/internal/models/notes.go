package models

import (
	"database/sql"
	"time"
)

// This is used as an interface. The model's functions take care of converting to the database format
// Note: This should probably be a timestamp, so that the implementation is not specific to html
type Note struct {
	ID                     int
	Provider               string
	Username               string // Eventually split this into fname and lname, but for now is fine
	Patient                string
	Service                string
	ServiceDate            time.Time
	StartTime              time.Time
	EndTime                time.Time
	Summary                string
	Progress               string
	Response               string
	AssessmentStatus       string
	RiskFactors            string
	EmergencyInterventions string
	Status                 string
}

// The reason we have this is so we can use the db connection, without
// having to re-establish it. It is passed in from the main func
type NoteModel struct {
	DB *sql.DB
}

func (n *NoteModel) Insert(providerID int,
	patient string,
	service string,
	serviceDate string,
	startTime string,
	endTime string,
	summary string,
	progress string,
	response string,
	assessmentStatus string,
	riskFactors string,
	emergencyInterventions string) (int, error) {
	statement := `insert into Note 
	(providerID, patient, service, serviceDate, startTime, endTime, summary, progress, response, assessmentStatus, riskFactors, emergencyInterventions, status)
	values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, 'Pending')`

	result, err := n.DB.Exec(statement, providerID, patient, service, serviceDate, startTime, endTime, summary, progress, response, assessmentStatus, riskFactors, emergencyInterventions)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (n *NoteModel) Get(id int) (Note, error) {
	//DB.query to get the note based on the id

	statement := `select Note.id, concat(User.fname, ' ', User.lname), User.username, Note.patient, Note.service, Note.serviceDate, Note.startTime, Note.endTime, summary, progress, response, assessmentStatus, riskFactors, emergencyInterventions, status from Note inner join User on Note.providerID = User.id where Note.id = ?`

	row := n.DB.QueryRow(statement, id)

	var note Note

	var startString, endString string

	err := row.Scan(&note.ID, &note.Provider, &note.Username, &note.Patient, &note.Service, &note.ServiceDate, &startString, &endString, &note.Summary, &note.Progress, &note.Response, &note.AssessmentStatus, &note.RiskFactors, &note.EmergencyInterventions, &note.Status)

	if err != nil {
		return Note{}, err
	}

	note.StartTime, err = time.Parse("15:04:05", startString)

	if err != nil {
		return Note{}, nil
	}

	note.EndTime, err = time.Parse("15:04:05", endString)

	if err != nil {
		return Note{}, nil
	}

	return note, nil
}

func (*NoteModel) CheckExistingNote() (int, error) {

	return 0, nil
}

func (n *NoteModel) GetNotesByProvider(username string) ([]Note, error) {

	statement := `select Note.id, concat(User.fname, ' ', User.lname),Note.patient, Note.service, Note.serviceDate, Note.startTime, Note.endTime, summary, progress, response, assessmentStatus, riskFactors, emergencyInterventions, status from Note inner join User on Note.providerID = User.id where User.username = ?`

	rows, err := n.DB.Query(statement, username)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	notes := make([]Note, 0)

	for rows.Next() {
		var note Note
		var startString, endString string

		err := rows.Scan(&note.ID, &note.Provider, &note.Patient, &note.Service, &note.ServiceDate, &startString, &endString, &note.Summary, &note.Progress, &note.Response, &note.AssessmentStatus, &note.RiskFactors, &note.EmergencyInterventions, &note.Status)

		if err != nil {
			return nil, err
		}

		note.StartTime, err = time.Parse("15:04:05", startString)

		if err != nil {
			return nil, err
		}

		note.EndTime, err = time.Parse("15:04:05", endString)

		if err != nil {
			return nil, err
		}

		notes = append(notes, note)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return notes, nil
}
