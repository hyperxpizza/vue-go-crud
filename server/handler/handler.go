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
	var newEmployee Employee

	if err := c.ShouldBindJSON(&newEmployee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	stmt, err := database.Db.Prepare(`UPDATE employees SET name = $1, email =$2 WHERE id = $3;`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	_, err = stmt.Exec(newEmployee.Name, newEmployee.Email, newEmployee.ID)
	if err == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, &newEmployee)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	stmt, err := database.Db.Prepare(`DELETE FROM employees WHERE id = $1`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	_, err = stmt.Exec(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Employee with id: " + id + " deleted.",
	})
}

func AddUser(c *gin.Context) {
	var newEmployee Employee

	if err := c.ShouldBindJSON(&newEmployee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	stmt, err := database.Db.Prepare(`INSERT INTO employees VALUES($1, $2, $3);`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	_, err = stmt.Exec(newEmployee.ID, newEmployee.Name, newEmployee.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, &newEmployee)

}
