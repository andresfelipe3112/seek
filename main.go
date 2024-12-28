package main

import (
    "candidate-management/config"
    "candidate-management/controllers"
    "candidate-management/middleware"
    "candidate-management/repository"
    "candidate-management/service"

    _ "candidate-management/docs" 

    "github.com/gin-gonic/gin"
    ginSwagger "github.com/swaggo/gin-swagger"
    swaggerFiles "github.com/swaggo/files"

)

// @title Candidate Management API
// @version 1.0
// @description This is the API documentation for the Candidate Management application.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
    // Configuración de la base de datos
    db := config.GetDB()

    // Inicialización de capas
    repo := &repository.CandidateRepository{DB: db}
    svc := &service.CandidateService{Repo: repo}
    ctrl := &controllers.CandidateController{Service: svc}

    // Configuración del router
    r := gin.Default()

    // Ruta de autenticación para obtener el token JWT
    r.POST("/login/get-token", controllers.GenerateToken)

    // Rutas protegidas con autenticación
    auth := r.Group("/").Use(middleware.JWTAuthMiddleware())
    auth.GET("/candidates", ctrl.GetAll)
    auth.POST("/candidates", ctrl.Create)

    // Ruta para la documentación de Swagger
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // Iniciar el servidor
    r.Run(":8080")
}
