package model

import "time"

type Submission struct {
	ID              *uint64    `json:"id" db:"id"`
	UUID            *string    `json:"uuid" db:"uuid"`
	ProjectID       *string    `json:"projectID" db:"project_id"`
	SubmittedAt     *time.Time `json:"submittedAt" db:"submitted_at"`
	Project         *Project
	SubmittedFields []SubmittedField
}

type SubmittedField struct {
	ID           *uint64 `json:"id" db:"id"`
	FieldID      *string `json:"fieldID" db:"field_id"`
	SubmissionID *string `json:"submissionID" db:"submission_id"`
	Value        *byte   `json:"value" db:"value"`
	Field        *Field
	Submission   *Submission
}
