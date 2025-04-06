package main

import (
    "errors"
    "fmt"
    "net/http"
    "os"
    "strings"
    "time"

    "github.com/golang-jwt/jwt/v5"
)

// Para simplificar, usaremos uma variável de ambiente ou um valor fixo
func getJWTSecret() []byte {
    secret := os.Getenv("JWT_SECRET")
    if secret == "" {
        secret = "secret_jwt_chave_trocar"
    }
    return []byte(secret)
}

// GenerateJWT gera um token com ID e Email do usuário
func GenerateJWT(userID int, email string) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "email":   email,
        "exp":     time.Now().Add(time.Hour * 24).Unix(), // token expira em 24 horas
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(getJWTSecret())
}

// AuthenticateJWT valida o token recebido no cabeçalho Authorization e retorna o user_id
func AuthenticateJWT(w http.ResponseWriter, r *http.Request) (int, error) {
    authHeader := r.Header.Get("Authorization")
    if authHeader == "" {
        http.Error(w, "Token não fornecido", http.StatusUnauthorized)
        return 0, errors.New("missing token")
    }

    parts := strings.Split(authHeader, " ")
    if len(parts) != 2 || parts[0] != "Bearer" {
        http.Error(w, "Formatação incorreta do header Authorization", http.StatusUnauthorized)
        return 0, errors.New("invalid auth header")
    }

    tokenString := parts[1]

    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Verifica se o método de assinatura do token é o esperado
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("método de assinatura inesperado: %v", token.Header["alg"])
        }
        return getJWTSecret(), nil
    })

    if err != nil {
        http.Error(w, "Token inválido", http.StatusUnauthorized)
        return 0, err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        // Recupera o user_id
        userID, ok := claims["user_id"].(float64)
        if !ok {
            http.Error(w, "Token não contém user_id válido", http.StatusUnauthorized)
            return 0, errors.New("invalid user_id in token")
        }
        return int(userID), nil
    }

    http.Error(w, "Token inválido", http.StatusUnauthorized)
    return 0, errors.New("invalid token")
}
