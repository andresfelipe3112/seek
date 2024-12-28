package models

import "errors"

var ErrCandidateExists = errors.New("candidate already exists")
type Candidate struct {
    ID            int     `db:"id" json:"id"`
    Name          string  `db:"name" json:"name"`
    Email         string  `db:"email" json:"email"`
    Gender        string  `db:"gender" json:"gender"`
    SalaryExpected float64 `db:"salary_expected" json:"salary_expected"`
    CreatedAt     string  `db:"created_at" json:"created_at"`
}
