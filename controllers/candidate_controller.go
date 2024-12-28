package controllers

import (
    "candidate-management/models"
    "candidate-management/service"
    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
    "net/http"
    "time"
    "log"
    "os"
)

type CandidateController struct {
    Service service.CandidateServiceInterface
}

var secretKey = os.Getenv("SECRET_KEY")

// GenerateToken godoc
// @Summary Generate a JWT token
// @Description Generate a JWT token for authentication
// @Tags Authentication
// @Produce json
// @Success 200 {object} map[string]interface{} "A JWT token"
// @Failure 500 {object} map[string]interface{} "Error message"
// @Router /login/get-token [post]
func GenerateToken(c *gin.Context) {
    // Datos ficticios para generar el token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username": "admin",
        "exp":      time.Now().Add(time.Hour * 24).Unix(),
    })

    tokenString, err := token.SignedString([]byte(secretKey))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

// GetAll godoc
// @Summary Get all candidates
// @Description Fetch all candidates from the database
// @Tags Candidates
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Candidate
// @Failure 500 {object} map[string]interface{} "Error fetching candidates"
// @Router /candidates [get]
func (ctrl *CandidateController) GetAll(c *gin.Context) {
    // Log para ver cuando se empieza a obtener los candidatos
	log.Println("Fetching all candidates...")

    candidates, err := ctrl.Service.GetAllCandidates()
    if err != nil {
        // Log de error si falla la obtención de los candidatos
		log.Printf("Error fetching candidates: %v\n", err)

        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching candidates"})
        return
    }

    // Log para verificar cuántos candidatos fueron encontrados
	log.Printf("Successfully fetched %d candidates.\n", len(candidates))

    c.JSON(http.StatusOK, candidates)
}

// Create godoc
// @Summary Create a new candidate
// @Description Add a new candidate to the database
// @Tags Candidates
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param candidate body models.Candidate true "Candidate details"
// @Success 201 {object} map[string]interface{} "Candidate created successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 500 {object} map[string]interface{} "Error creating candidate"
// @Router /candidates [post]
func (ctrl *CandidateController) Create(c *gin.Context) {
    var candidate models.Candidate

    log.Println("Processing new candidate creation request...")

    if err := c.ShouldBindJSON(&candidate); err != nil {
        log.Printf("Error binding candidate data: %v\n", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    log.Printf("Received candidate data: %+v\n", candidate)

    // Check if the candidate's email already exists
    existingCandidate, err := ctrl.Service.GetCandidateByEmail(candidate.Email)
    if err == nil && existingCandidate != nil {
        log.Printf("Error: Candidate with email %s already exists\n", candidate.Email)
        c.JSON(http.StatusConflict, gin.H{"error": "Candidate with this email already exists"})
        return
    }

    if err := ctrl.Service.CreateCandidate(candidate); err != nil {
        log.Printf("Error creating candidate in database: %v\n", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating candidate"})
        return
    }

    log.Println("Candidate created successfully")

    c.JSON(http.StatusCreated, gin.H{"message": "Candidate created successfully"})
}

