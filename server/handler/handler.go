package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hyperxpizza/vue-go-crud/server/database"
)

type Employee struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetAllEmployee(c *gin.Context) {
	var employees []Employee

	rows, err := database.Db.Query("SELECT * FROM employees")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	for rows.Next() {
		var employee Employee

		err := rows.Scan(
			&employee.ID,
			&employee.Name,
			&employee.Email,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})

			return
		}

		employees = append(employees, employee)
	}

	c.JSON(http.StatusOK, &employees)
}

func UpdateEmployee(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}

func AddUser(c *gin.Context) {

}
