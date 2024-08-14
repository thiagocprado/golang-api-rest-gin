package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thiagocprado/golang-api-rest-gin/database"
	"github.com/thiagocprado/golang-api-rest-gin/models"
)

func DeleteStudent(c *gin.Context) {
	var student models.Student

	id := c.Params.ByName("id")
	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Student not found!",
		})
		return
	}

	database.DB.Delete(&student, id)

	c.JSON(http.StatusOK, gin.H{
		"msg": "Student deleted!",
	})
}

func GetAllStudents(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)

	c.JSON(http.StatusOK, students)
}

func GetStudentByCpf(c *gin.Context) {
	var student models.Student

	cpf := c.Param("cpf")
	database.DB.Where(&models.Student{CPF: cpf}).First(&student)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Student not found!",
		})
		return
	}

	c.JSON(http.StatusOK, student)
}

func GetStudentByID(c *gin.Context) {
	var student models.Student

	id := c.Params.ByName("id")
	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Student not found!",
		})
		return
	}

	c.JSON(http.StatusOK, student)
}

func HandleRouteNotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "pong",
	})
}

func SaveStudent(c *gin.Context) {
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	if err := models.ValidateStudent(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	database.DB.Create(&student)
	c.JSON(http.StatusOK, student)
}

func ShowIndexPage(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"students": students,
	})
}

func UpdateStudent(c *gin.Context) {
	var student models.Student

	id := c.Params.ByName("id")
	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Student not found!",
		})
		return
	}

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	if err := models.ValidateStudent(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	uID, _ := strconv.ParseUint(id, 10, 64)
	student.ID = uint(uID)

	database.DB.Save(&student)
	c.JSON(http.StatusOK, student)
}
