package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/titoyudha/golang_todo_list/config"
	"github.com/titoyudha/golang_todo_list/controller"
	"github.com/titoyudha/golang_todo_list/model"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.ConnectDB()
)

func main() {
	server := gin.Default()
	defer config.DisconnectDB(db)
	db.AutoMigrate(model.Todo{})

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"Message": "Todo List App"})
	})

	server.POST("/api/v1/todo", controller.Create)
	server.PUT("/api/v1/todo/:id", controller.Update)
	server.GET("/api/v1/todo", controller.Get)
	server.DELETE("/api/v1/todo/:id", controller.Delete)
	fmt.Println("Server is Running")

	server.Run()

}
