package app

import (
	"belajar-golang-restful-api/controller"
	"belajar-golang-restful-api/exeption"
	"net/http"

	"github.com/julienschmidt/httprouter"
	httpSwagger "github.com/swaggo/http-swagger"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)
	router.GET("/swagger/*any", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		httpSwagger.WrapHandler.ServeHTTP(w, r)
	})

	router.PanicHandler = exeption.ErrorHandler

	return router
}