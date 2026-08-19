package main

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/middleware"
	"net/http"

	_ "belajar-golang-restful-api/docs"

	"github.com/joho/godotenv"
)


func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr: "localhost:8080",
		Handler: authMiddleware,
	}
}

// @title Category Restful API
// @version 1.0
// @description API for Category
// @host localhost:8080
// @BasePath /api
func main() {
	err := godotenv.Load()
	helper.PanicIfError(err)

	server := InitializeServer()

	err = server.ListenAndServe()
	helper.PanicIfError(err)
}

