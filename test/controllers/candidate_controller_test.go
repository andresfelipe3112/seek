package controllers_test

import (
	"bytes"
	"candidate-management/controllers"
	"candidate-management/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Mock de CandidateService
type MockCandidateService struct{}

func (m *MockCandidateService) GetAllCandidates() ([]models.Candidate, error) {
	return []models.Candidate{
		{ID: 1, Name: "John Doe", Email: "john.doe@example.com"},
		{ID: 2, Name: "Jane Doe", Email: "jane.doe@example.com"},
	}, nil
}

func (m *MockCandidateService) CreateCandidate(candidate models.Candidate) error {
	if candidate.Email == "existing@example.com" {
		return models.ErrCandidateExists
	}
	return nil
}

func (m *MockCandidateService) GetCandidateByEmail(email string) (*models.Candidate, error) {
	if email == "existing@example.com" {
		return &models.Candidate{ID: 1, Name: "Existing User", Email: "existing@example.com"}, nil
	}
	return nil, nil
}

// Test para GetAll
func TestGetAllCandidates(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Configurar el controlador con un mock
	mockService := &MockCandidateService{}
	ctrl := controllers.CandidateController{Service: mockService}

	// Crear un router y agregar la ruta
	router := gin.Default()
	router.GET("/candidates", ctrl.GetAll)

	// Simular la solicitud HTTP
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/candidates", nil)
	router.ServeHTTP(w, req)

	// Verificar los resultados
	assert.Equal(t, http.StatusOK, w.Code)

	var candidates []models.Candidate
	err := json.Unmarshal(w.Body.Bytes(), &candidates)
	assert.NoError(t, err)
	assert.Len(t, candidates, 2)
	assert.Equal(t, "John Doe", candidates[0].Name)
}

// Test para Create
func TestCreateCandidate(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Configurar el controlador con un mock
	mockService := &MockCandidateService{}
	ctrl := controllers.CandidateController{Service: mockService}

	// Crear un router y agregar la ruta
	router := gin.Default()
	router.POST("/candidates", ctrl.Create)

	// Caso exitoso
	newCandidate := models.Candidate{Name: "New User", Email: "new.user@example.com"}
	body, _ := json.Marshal(newCandidate)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/candidates", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "Candidate created successfully", response["message"])

	// Caso de email existente
	existingCandidate := models.Candidate{Name: "Existing User", Email: "existing@example.com"}
	body, _ = json.Marshal(existingCandidate)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/candidates", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusConflict, w.Code)
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "Candidate with this email already exists", response["error"])
}
