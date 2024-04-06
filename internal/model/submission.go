package model

import "time"

type Submission struct {
	ID              uint64    `json:"id" db:"id"`
	ProjectID       uint64    `json:"projectID" db:"project_id"`
	SubmittedAt     time.Time `json:"submittedAt" db:"submitted_at"`
	Project         *Project
	SubmittedFields *[]SubmittedField
}

type SubmittedField struct {
	ID           uint64 `json:"id" db:"id"`
	FieldID      uint64 `json:"fieldID" db:"field_id"`
	SubmissionID uint64 `json:"submissionID" db:"submission_id"`
	Value        byte   `json:"value" db:"value"`
	Field        *Field
	Submission   *Submission
}
