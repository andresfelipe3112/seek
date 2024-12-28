package service

import (
	"candidate-management/models"
	"candidate-management/repository"
)

// CandidateServiceInterface define los métodos que debe implementar cualquier servicio de candidato
type CandidateServiceInterface interface {
	GetAllCandidates() ([]models.Candidate, error)
	CreateCandidate(candidate models.Candidate) error
	GetCandidateByEmail(email string) (*models.Candidate, error)
}

// CandidateService implementa CandidateServiceInterface
type CandidateService struct {
	Repo *repository.CandidateRepository
}

// GetAllCandidates devuelve todos los candidatos
func (s *CandidateService) GetAllCandidates() ([]models.Candidate, error) {
	return s.Repo.GetAll()
}

// CreateCandidate crea un nuevo candidato
func (s *CandidateService) CreateCandidate(candidate models.Candidate) error {
	return s.Repo.Create(candidate)
}

// GetCandidateByEmail obtiene un candidato por su correo electrónico
func (s *CandidateService) GetCandidateByEmail(email string) (*models.Candidate, error) {
	return s.Repo.GetByEmail(email)
}
