package main

import (
  "fmt"
  EmployeeController "backend/api/controller/employee"
  AdminController "backend/api/controller/admin"
  AuthController "backend/api/controller/auth"
  "github.com/gin-gonic/gin"
  "backend/api/db"
  "github.com/joho/godotenv"
  "backend/api/middleware"
)

func main() {
  //Get .env
  err := godotenv.Load(".env")
  if err != nil {
	fmt.Println("Error loading .env file")
  }
  //get InitDB fuction
  db.InitDB()

  router := gin.Default()

  authorized := router.Group("/api", middleware.JwtAuthen())

  //Employee API Method
  router.GET("/employee", EmployeeController.GetEmployee) //GET
  router.GET("/employee/:id", EmployeeController.GetEmployeeByID) //GET BY ID
  router.GET("/employeedb", EmployeeController.GetEmployeeDB) //GET FROM DB
  router.GET("/admin", AdminController.GetAdmin) //POST TO DB

  authorized.POST("/employee", EmployeeController.PostEmployee) //POST
  authorized.POST("/employeedb", EmployeeController.PostEmployeeDB) //POST TO DB
  
  router.POST("/register", AdminController.PostAdmin) //POST TO DB

  authorized.PUT("/employee", EmployeeController.PutEmployee) //PUT
  authorized.PUT("/employeedb", EmployeeController.PutEmployeeDB) //PUT DB
  
  authorized.DELETE("/employee", EmployeeController.DeleteEmployee) //DELETE
  authorized.DELETE("/employeedb/:id", EmployeeController.DeleteEmployeeDB) //DELETE DB

  //Customer API Method
  router.POST("/login", AuthController.Login) //POST LOGIN
  router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}