package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/titoyudha/golang_todo_list/controller"
	"github.com/titoyudha/golang_todo_list/model"
	"gorm.io/gorm"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestCreateTodo(t *testing.T) {
	r := SetUpRouter()
	r.POST("/api/v1/todo", controller.Create)

	todo := model.Todo{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: gorm.DeletedAt{},
		},
		Title: "test title",
		Email: "test@gmail.com",
	}
	jsonValue, _ := json.Marshal(todo)
	req, _ := http.NewRequest("POST", "/api/v1/todo", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetTodo(t *testing.T) {
	r := SetUpRouter()
	r.GET("/api/v1/todo", controller.Get)
	req, _ := http.NewRequest("GET", "/api/v1/todo", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var todo []model.Todo
	json.Unmarshal(w.Body.Bytes(), &todo)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateTodo(t *testing.T) {
	r := SetUpRouter()
	r.PUT("/api/v1/todo/:id", controller.Update)

	todo := model.Todo{
		Model: gorm.Model{
			ID:        13,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: gorm.DeletedAt{},
		},
		Title: "test update",
		Email: "testupdate@gmail.com",
	}

	jsonValue, _ := json.Marshal(todo)
	req, _ := http.NewRequest("PUT", "/api/v1/todo/13", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	reqNotFound, _ := http.NewRequest("PUT", "/api/v1/todo/13", bytes.NewBuffer(jsonValue))
	r.ServeHTTP(w, reqNotFound)
	assert.Equal(t, http.StatusCreated, w.Code)
}
