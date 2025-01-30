package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // Importación del driver MySQL
)

func ConnectDB() (*sql.DB, error) {
	dsn := "root:01toto01@tcp(127.0.0.1:3306)/clinic" // Cambia el usuario, contraseña y base de datos según sea necesario
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %w", err)
	}
	// Prueba la conexión
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	return db, nil
}
