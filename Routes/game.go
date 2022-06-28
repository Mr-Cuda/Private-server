package game

import (
	"fmt"
	"net/http"

  "http://"
)

//Creating Employee Table using Postgres function query
func CreateEmployeeTable(c *gin.Context) {
	var db = database.DB

	Emp := models.Employee{}
	c.BindJSON(&Emp)

	db.Exec("Call EmployeeTable($1, $2, $3, $4) ", "Name", "Email", "Phone", "Address").Scan(&models.Employee{})

	c.JSON(http.StatusOK, gin.H{"message": "Employee table created successfully"})

}

//This fucntion will add employees to the database
func AddEmployee(c *gin.Context) {
	var db = database.DB

	Emp := new(models.Employee)
	c.BindJSON(&Emp)

	if Emp.Name == "" {
		c.JSON(http.StatusOK, gin.H{"message": "Employee name is required"})
	} else if Emp.Email == "" {
		c.JSON(http.StatusOK, gin.H{"message": "Employee email is required"})
	} else if Emp.Phone == "" {
		c.JSON(http.StatusOK, gin.H{"message": "Employee phone is required"})
	} else if Emp.Address == "" {
		c.JSON(http.StatusOK, gin.H{"message": "Employee address is required"})
	} else {
		fmt.Println("Error here")
		db.Exec("Call AddEmployee($1, $2, $3, $4) ", "Name", "Email", "Phone", "Address").Scan(&models.Employee{})
		c.JSON(http.StatusOK, gin.H{
			"message":  "Employee added successfully",
			"Employee": Emp,
		})

	}

}

func GetEmployees(c *gin.Context) {
	var db = database.DB
	var Emp []models.Employee
	db.Find(&Emp)
	c.JSON(http.StatusOK, gin.H{
		"Employees": Emp,
	})
}
func GetEmployeeById(c *gin.Context) {
	var db = database.DB
	id := c.Param("id")
	var Emp models.Employee
	db.Find(&Emp, id)
	if db.Find(&Emp, id).RecordNotFound() {
		c.JSON(http.StatusOK, gin.H{"message": "No employee found"})
	} else {
		c.JSON(http.StatusOK, gin.H{"Employees": Emp})
	}
}
func DeleteEmployee(c *gin.Context) {
	var db = database.DB
	id := c.Param("id")
	var Emp models.Employee
	db.Find(&Emp, id)
	if Emp.Id == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Employee not found"})
	} else {
		db.Delete(&Emp, id)
		c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
	}
}
func UpdateEmpDetails(c *gin.Context) {
	var db = database.DB
	Emp := new(models.Employee)
	c.BindJSON(&Emp)
	id := c.Param("id")
	var Emp1 models.Employee
	db.Find(&Emp1, id)
	Emp1.Name = Emp.Name
	Emp1.Email = Emp.Email
	Emp1.Phone = Emp.Phone
	Emp1.Address = Emp.Address
	err := db.Save(&Emp1)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Employee details updated successfully"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Employee details not updated"})
	}
} 
