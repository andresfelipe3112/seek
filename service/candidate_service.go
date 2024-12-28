
package service


import (
    "candidate-management/models"
    "candidate-management/repository"
)

type CandidateService struct {
    Repo *repository.CandidateRepository
}

func (s *CandidateService) GetAllCandidates() ([]models.Candidate, error) {
    return s.Repo.GetAll()
}

func (s *CandidateService) CreateCandidate(candidate models.Candidate) error {
    return s.Repo.Create(candidate)
}

// GetCandidateByEmail obtiene un candidato por su correo electrónico
func (s *CandidateService) GetCandidateByEmail(email string) (*models.Candidate, error) {
    return s.Repo.GetByEmail(email)
}
