package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "net/http"

    "golang.org/x/crypto/bcrypt"
)

// Handler de Registro
func RegisterHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
        return
    }

    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
        return
    }

    if user.Email == "" || user.Password == "" {
        http.Error(w, "Email e senha são obrigatórios", http.StatusBadRequest)
        return
    }

    // Faz hash da senha
    hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, "Erro ao gerar hash da senha", http.StatusInternalServerError)
        return
    }

    // Salva no DB
    var id int
    query := `INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING id`
    err = db.QueryRow(query, user.Email, string(hash)).Scan(&id)
    if err != nil {
        http.Error(w, "Erro ao inserir usuário", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    fmt.Fprintf(w, "Usuário criado com ID: %d\n", id)
}

// Handler de Login
func LoginHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
        return
    }

    var creds struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
        http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
        return
    }

    // Busca usuário por email
    var user User
    query := `SELECT id, email, password_hash FROM users WHERE email = $1`
    err := db.QueryRow(query, creds.Email).Scan(&user.ID, &user.Email, &user.Password)
    if err != nil {
        if err == sql.ErrNoRows {
            http.Error(w, "Usuário não encontrado", http.StatusUnauthorized)
        } else {
            http.Error(w, "Erro ao buscar usuário", http.StatusInternalServerError)
        }
        return
    }

    // Compara senha
    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
    if err != nil {
        http.Error(w, "Senha incorreta", http.StatusUnauthorized)
        return
    }

    // Se senha for correta, gera JWT
    tokenString, err := GenerateJWT(user.ID, user.Email)
    if err != nil {
        http.Error(w, "Falha ao gerar token", http.StatusInternalServerError)
        return
    }

    // Retorna token em JSON
    resp := map[string]string{"token": tokenString}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}

// Handler de Listagem de usuários
func ListUsersHandler(db *sql.DB, w http.ResponseWriter, r *http.Request, userID int) {
    // userID é só para ilustrar caso queira usar (ex: checar permissões, etc.)
    rows, err := db.Query("SELECT id, email FROM users")
    if err != nil {
        http.Error(w, "Erro ao buscar usuários", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var users []User
    for rows.Next() {
        var usr User
        if err := rows.Scan(&usr.ID, &usr.Email); err != nil {
            log.Println(err)
            continue
        }
        users = append(users, usr)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}
