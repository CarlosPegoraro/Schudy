package main

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
    "os"
)

func InitDB() (*sql.DB, error) {
    host := os.Getenv("DB_HOST")
    if host == "" {
        host = "localhost"
    }
    port := os.Getenv("DB_PORT")
    if port == "" {
        port = "5432"
    }
    user := os.Getenv("DB_USER")
    if user == "" {
        user = "postgres"
    }
    password := os.Getenv("DB_PASSWORD")
    if password == "" {
        password = "123"
    }
    dbname := os.Getenv("DB_NAME")
    if dbname == "" {
        dbname = "schudy"
    }

    // String de conexão com Postgres
    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        return nil, err
    }

    // Verifica conexão
    err = db.Ping()
    if err != nil {
        return nil, err
    }

    // Criar a tabela de usuários se não existir (exemplo)
    createTable := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        email VARCHAR(255) UNIQUE NOT NULL,
        password_hash VARCHAR(255) NOT NULL
    );
    `
    _, err = db.Exec(createTable)
    if err != nil {
        return nil, err
    }

    return db, nil
}
