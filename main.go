package main

import (
	"fmt"
	"go-restful/app"
	"go-restful/controller"
	"go-restful/helper"
	"go-restful/middleware"
	"go-restful/repository"
	"go-restful/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	db := app.NewDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr: "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	fmt.Println("Application is running on port 3000")
	helper.PanicIfError(err)

}