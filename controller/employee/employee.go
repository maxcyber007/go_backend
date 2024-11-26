package employee

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "backend/api/db"
  "backend/api/model"
)

// func GetEmployee(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 	"message": "Employee GET Method!",
// 	})
// }

func GetEmployee(c *gin.Context) {
	var employees []model.Tbl_employee
	db.Db.Find(&employees)
	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "User Read Success", "employees": employees})
}

func PostEmployee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
	"message": "Employee POST Method!",
	})
}

func PutEmployee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
	"message": "Employee PUT Method!",
	})
}

func DeleteEmployee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
	"message": "Employee DELETE Method!",
	})
}