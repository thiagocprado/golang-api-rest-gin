package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thiagocprado/golang-api-rest-gin/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

	r.DELETE("/students/:id", controllers.DeleteStudent)

	r.GET("/students", controllers.GetAllStudents)
	r.GET("/students/cpf/:cpf", controllers.GetStudentByCpf)
	r.GET("/students/:id", controllers.GetStudentByID)
	r.GET("/ping", controllers.Ping)
	r.GET("/pages/index", controllers.ShowIndexPage)

	r.POST("/students", controllers.SaveStudent)

	r.PUT("/students/:id", controllers.UpdateStudent)

	r.NoRoute(controllers.HandleRouteNotFound)

	r.Run()
}
