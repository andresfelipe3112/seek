package repository

import (
    "candidate-management/models"
    "github.com/jmoiron/sqlx"
    "log"
)

type CandidateRepository struct {
    DB *sqlx.DB
}

// GetAll obtiene todos los candidatos
func (r *CandidateRepository) GetAll() ([]models.Candidate, error) {
    var candidates []models.Candidate
    err := r.DB.Select(&candidates, "SELECT * FROM candidates")
    return candidates, err
}

// Create agrega un nuevo candidato
func (r *CandidateRepository) Create(candidate models.Candidate) error {
    _, err := r.DB.NamedExec(`
        INSERT INTO candidates (name, email, gender, salary_expected) 
        VALUES (:name, :email, :gender, :salary_expected)`, candidate)
    return err
}

// GetByEmail obtiene un candidato por su correo electrónico
func (r *CandidateRepository) GetByEmail(email string) (*models.Candidate, error) {
    var candidate models.Candidate
    err := r.DB.Get(&candidate, "SELECT * FROM candidates WHERE email = ?", email)
    if err != nil {
        if err.Error() == "sql: no rows in result set" {
            return nil, nil // No se encontró el candidato
        }
        log.Printf("Error fetching candidate by email: %v", err)
        return nil, err
    }
    return &candidate, nil
}
