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
	dbURL := db.LoadDBConfig()
	dbConnection := db.GetDBConnection(dbURL)

	userRepo := repository.UserRepo{Db: dbConnection}
	orderRepo := repository.OrderRepo{Db: dbConnection}

	userService := service.UserService{UserRepo: &userRepo}
	userHandler := handler.UserHandler{UserService: &userService}

	orderService := service.OrderService{OrderRepo: &orderRepo}
	orderHandler := handler.OrderHandler{OrderService: &orderService}

	// 3. Router Setup
	r := gin.Default()
	api := r.Group("/user")
	{
		api.GET("/getUserWithOrders", userHandler.GetUserWithOrdersHandler)
	}

	api1 := r.Group("/order")
	{
		api1.GET("/getOrders", orderHandler.GetOrdersHandler)
	}

	r.Run(":8081")
}
