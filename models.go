package main

type User struct {
    ID           int    `json:"id"`
    Username     string
    Email        string `json:"email"`
    Password     string
    BornDate     string
}
