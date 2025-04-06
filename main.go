package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func main() {
    // Atribuir a porta ou usar padrão 8080
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    // Inicializa a conexão com o banco de dados
    db, err := InitDB()
    if err != nil {
        log.Fatalf("Erro ao conectar no banco: %v\n", err)
    }
    defer db.Close()

    // Criação das rotas
    http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
        RegisterHandler(db, w, r)
    })
    http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
        LoginHandler(db, w, r)
    })
    // Rota protegida - requer token JWT
    http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
        // Primeiro verificamos a autenticação
        userID, err := AuthenticateJWT(w, r)
        if err != nil {
            // Já foi tratado dentro de AuthenticateJWT
            return
        }
        // Se chegou aqui, token é válido
        ListUsersHandler(db, w, r, userID)
    })

    fmt.Printf("Servidor rodando na porta %s...\n", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
