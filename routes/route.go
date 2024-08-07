package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thiagocprado/golang-api-rest-gin/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.GET("/students", controllers.GetAllStudents)
	r.GET("/students/cpf/:cpf", controllers.GetStudentByCpf)
	r.GET("/students/:id", controllers.GetStudentById)
	r.GET("/ping", controllers.Ping)
	r.POST("/students", controllers.SaveStudent)
	r.PUT("/students/:id", controllers.UpdateStudent)
	r.Run()
}
