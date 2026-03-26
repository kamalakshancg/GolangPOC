package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kamalakshancg/GolangPOC/internal/db"
	"github.com/kamalakshancg/GolangPOC/internal/handler"
	"github.com/kamalakshancg/GolangPOC/internal/repository"
	"github.com/kamalakshancg/GolangPOC/internal/service"
)

func main() {
	// 1. Init DB
	dsn := db.LoadDBConfig()
	database := db.NewPostgresDB(dsn)

	// 2. Dependency Injection

	userRepo := repository.UserRepo{Db: database}
	orderRepo := repository.OrderRepo{Db: database}

	userService := service.UserService{UserRepo: &userRepo}
	userHandler := handler.UserHandler{UserService: &userService}

	orderService := service.OrderService{OrderRepo: &orderRepo}
	orderHandler := handler.OrderHandler{OrderService: &orderService}

	// 3. Router Setup
	r := gin.Default()
	api := r.Group("/api/user")
	{
		api.GET("/test3", userHandler.Test3)
	}

	api1 := r.Group("/api/order")
	{
		api1.GET("/test1", orderHandler.Test1)
		api1.GET("/test2", orderHandler.Test2)
	}

	r.Run(":8080")
}
