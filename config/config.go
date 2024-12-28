package config

import (
    "log"
    "github.com/jmoiron/sqlx"
    _ "github.com/go-sql-driver/mysql"
    "github.com/joho/godotenv"
    "os"
)

func GetDB() *sqlx.DB {
    // Cargar el archivo .env
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Obtener las variables de entorno
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")

    // Crear la cadena de conexión (DSN) usando las variables de entorno
    dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName
    db, err := sqlx.Connect("mysql", dsn)
    if err != nil {
        log.Fatal("Error connecting to the database: ", err)
    }
    return db
}
