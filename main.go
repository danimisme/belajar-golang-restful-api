package main

import (
	"belajar-golang-restful-api/app"
	"belajar-golang-restful-api/controller"
	"belajar-golang-restful-api/exeption"
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/repository"
	"belajar-golang-restful-api/service"
	"net/http"

	_ "belajar-golang-restful-api/docs"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Category Restful API
// @version 1.0
// @description API for Category
// @host localhost:8080
// @BasePath /api
func main() {
	err := godotenv.Load()
	helper.PanicIfError(err)
	db := app.NewDb()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository,db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exeption.ErrorHandler

	router.GET("/swagger/*any", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		httpSwagger.WrapHandler.ServeHTTP(w, r)
	})

	server := http.Server{
		Addr: "localhost:8080",
		Handler: router,
	}

	err = server.ListenAndServe()
	helper.PanicIfError(err)
}

