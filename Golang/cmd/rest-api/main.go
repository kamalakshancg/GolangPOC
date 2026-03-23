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
	pocRepo := repository.PocRepo{Db: database}
	pocService := service.PocService{Repo: &pocRepo}
	h := handler.PocHandler{Service: &pocService}

	// 3. Router Setup
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/test1", h.Test1)
		api.GET("/test2", h.Test2)
		api.GET("/test3", h.Test3)
	}

	r.Run(":8080")
}
