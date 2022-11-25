package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/titoyudha/golang_todo_list/config"
	"github.com/titoyudha/golang_todo_list/model"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.ConnectDB()
)

func Create(ctx *gin.Context) {
	var todo model.Todo

	// Processing entity/model
	if err := ctx.ShouldBindJSON(&todo); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	//parsing time
	todo.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	//Saving data into db
	res := db.Save(&todo)
	if res.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to save to DB"})
		return
	}

	//Final Result
	ctx.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusCreated,
		"message": "Success Create Todo",
		"data":    todo,
	})
}

func Get(ctx *gin.Context) {
	var todo []model.Todo

	res := db.Find(&todo)
	if res.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error while get all data"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success Retrieving All Data",
		"Code":    200,
		"data":    todo,
	})
}

func Update(ctx *gin.Context) {
	var todo model.Todo

	param := ctx.Param("id")
	if err := ctx.ShouldBindJSON(&todo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	todos := model.Todo{}
	todoByID := db.Where("id = ?", param).First(&todos)
	if todoByID.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Todo not found"})
		return
	}
	res := db.Save(&todos)
	if res.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Something Wrong when saving todo"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusOK,
		"message": "Success Update",
		"data":    todo,
	})
}

func Delete(ctx *gin.Context) {
	todos := model.Todo{}
	params := ctx.Param("id")

	res := db.Where("id = ?", params).Unscoped().Delete(&todos)
	fmt.Println(res)

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Success Delete",
		"data":    nil,
	})
}
