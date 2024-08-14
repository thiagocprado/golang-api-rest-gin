package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/thiagocprado/golang-api-rest-gin/controllers"
	"github.com/thiagocprado/golang-api-rest-gin/database"
	"github.com/thiagocprado/golang-api-rest-gin/models"
)

var ID int

func mockDeleteStudent() {
	var student models.Student
	database.DB.Delete(&student, ID)
}

func mockStudent() {
	student := models.Student{
		Name: "John Doe",
		CPF:  "12345678910",
		RG:   "123456789",
	}

	database.DB.Create(&student)
	ID = int(student.ID)
}

func setupTestRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()

	return routes
}

func TestDeleteStudent(t *testing.T) {
	database.ConnectDB()

	mockStudent()

	r := setupTestRoutes()
	r.DELETE("/students/:id", controllers.DeleteStudent)

	path := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", path, nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGetAllStudents(t *testing.T) {
	database.ConnectDB()

	r := setupTestRoutes()
	r.GET("/students", controllers.GetAllStudents)

	req, _ := http.NewRequest("GET", "/students", nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGetStudentByCpf(t *testing.T) {
	database.ConnectDB()

	mockStudent()
	defer mockDeleteStudent()

	r := setupTestRoutes()
	r.GET("/students/cpf/:cpf", controllers.GetStudentByCpf)

	req, _ := http.NewRequest("GET", "/students/cpf/12345678910", nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGetStudentByID(t *testing.T) {
	database.ConnectDB()

	mockStudent()
	defer mockDeleteStudent()

	r := setupTestRoutes()
	r.GET("/students/:id", controllers.GetStudentByID)

	path := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", path, nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	var studentResp models.Student
	json.Unmarshal(resp.Body.Bytes(), &studentResp)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "John Doe", studentResp.Name)
}

func TestPing(t *testing.T) {
	r := setupTestRoutes()
	r.GET("/ping", controllers.Ping)

	req, _ := http.NewRequest("GET", "/ping", nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	bodyResp, _ := io.ReadAll(resp.Body)
	mockResp := `{"msg":"pong"}`

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, mockResp, string(bodyResp))
}

func TestUpdateStudent(t *testing.T) {
	database.ConnectDB()

	mockStudent()
	defer mockDeleteStudent()

	r := setupTestRoutes()
	r.PUT("/students/:id", controllers.UpdateStudent)

	student := models.Student{CPF: "10987654321", Name: "John Doe", RG: "987654321"}
	studentJson, _ := json.Marshal(student)

	path := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PUT", path, bytes.NewBuffer(studentJson))
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	var studentResp models.Student
	json.Unmarshal(resp.Body.Bytes(), &studentResp)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "10987654321", studentResp.CPF)
	assert.Equal(t, "987654321", studentResp.RG)
}
